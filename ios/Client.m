#import "Client.h"
#import "proto/Wanikani+Convenience.h"

NS_ASSUME_NONNULL_BEGIN

static const char *kURLBase = "https://www.wanikani.com/api/v2";
static const char *kProgressURL = "https://www.wanikani.com/json/progress";
static const char *kStudyMaterialsURLBase = "https://www.wanikani.com/study_materials";
static const char *kReviewSessionURL = "https://www.wanikani.com/review/session";

NSErrorDomain WKClientErrorDomain = @"WKClientErrorDomain";

static NSError *MakeError(int code, NSString *msg) {
  return [NSError errorWithDomain:WKClientErrorDomain
                             code:code
                         userInfo:@{NSLocalizedDescriptionKey:msg}];
}

static const NSTimeInterval kCSRFTokenValidity = 2 * 60 * 60;  // 2 hours.
static NSRegularExpression *kCSRFTokenRE;

typedef void(^PartialResponseHandler)(id _Nullable data, NSError * _Nullable error);

@implementation Client {
  NSString *_apiToken;
  NSString *_cookie;
  NSURLSession *_urlSession;
  NSDateFormatter *_dateFormatter;
  NSString *_csrfToken;
  NSDate *_csrfTokenUpdated;
}

- (instancetype)initWithApiToken:(NSString *)apiToken
                          cookie:(NSString *)cookie {
  static dispatch_once_t onceToken;
  dispatch_once(&onceToken, ^{
    kCSRFTokenRE = [NSRegularExpression regularExpressionWithPattern:
                    @"<meta name=\"csrf-token\" content=\"([^\"]*)" options:0 error:nil];
  });
  
  if (self = [super init]) {
    _apiToken = apiToken;
    _cookie = cookie;
    _urlSession = [NSURLSession sharedSession];
    _dateFormatter = [[NSDateFormatter alloc] init];
    _dateFormatter.dateFormat = @"yyyy-MM-dd'T'HH:mm:ss.SSSSSS'Z'";
    _dateFormatter.timeZone = [NSTimeZone timeZoneForSecondsFromGMT:0];
  }
  return self;
}

#pragma mark - Authorization

- (NSMutableURLRequest *)authorizeAPIRequest:(NSURL *)url {
  NSMutableURLRequest *req = [NSMutableURLRequest requestWithURL:url];
  [req setValue:[NSString stringWithFormat:@"Token token=%@", _apiToken]
      forHTTPHeaderField:@"Authorization"];
  return req;
}

- (NSMutableURLRequest *)authorizeUserRequest:(NSURL *)url {
  NSMutableURLRequest *req = [NSMutableURLRequest requestWithURL:url];
  [req setValue:[NSString stringWithFormat:@"_wanikani_session=%@", _cookie]
       forHTTPHeaderField:@"Cookie"];
  return req;
}

#pragma mark - Query utilities

- (NSString *)currentISO8601Time {
  return [_dateFormatter stringFromDate:[NSDate date]];
}

- (void)startPagedQueryFor:(NSURL *)url
                   handler:(PartialResponseHandler)handler {
  NSURLRequest *req = [self authorizeAPIRequest:url];
  NSLog(@"Request: %@", url);
  NSURLSessionDataTask *task =
      [_urlSession dataTaskWithRequest:req
                     completionHandler:^(NSData * _Nullable data,
                                         NSURLResponse * _Nullable response,
                                         NSError * _Nullable error) {
    [self parseJsonResponse:data
                      error:error
                    handler:handler];
  }];
  [task resume];
}

- (void)parseJsonResponse:(NSData *)data
                    error:(NSError *)error
                  handler:(PartialResponseHandler)handler {
  if (error != nil) {
    handler(nil, error);
    return;
  }
  NSDictionary *dict = [NSJSONSerialization JSONObjectWithData:data options:0 error:&error];
  if (error != nil) {
    handler(nil, error);
    return;
  }
  if (dict[@"error"] != nil) {
    handler(nil, MakeError(0, dict[@"error"]));
    return;
  }
  
  handler(dict[@"data"], nil);
  
  // Get the next page if we have one.
  if (!dict[@"pages"][@"next_url"]) {
    // This is a single-page query, so don't call the handler again.
    return;
  }
  if (dict[@"pages"][@"next_url"] != [NSNull null]) {
    NSString *nextURLString = dict[@"pages"][@"next_url"];
    NSLog(@"Request: %@", nextURLString);
    NSURLRequest *req = [self authorizeAPIRequest:[NSURL URLWithString:nextURLString]];
    NSURLSessionDataTask *task =
        [_urlSession dataTaskWithRequest:req
                      completionHandler:^(NSData * _Nullable data,
                                          NSURLResponse * _Nullable response,
                                          NSError * _Nullable error) {
      [self parseJsonResponse:data 
                        error:error
                      handler:handler];
    }];
    [task resume];
  } else {
    handler(nil, nil);
  }
}

#pragma mark - Assignments

- (void)getAssignmentsModifiedAfter:(NSString *)date
                            handler:(AssignmentHandler)handler {
  NSMutableArray<WKAssignment *> *ret = [NSMutableArray array];

  NSURLComponents *url =
      [NSURLComponents componentsWithString:[NSString stringWithFormat:@"%s/assignments",
                                             kURLBase]];
  if (date && date.length) {
    [url setQueryItems:@[
        [NSURLQueryItem queryItemWithName:@"updated_after" value:date],
    ]];
  }
  
  [self startPagedQueryFor:url.URL handler:^(NSArray *data, NSError *error) {
    if (error) {
      handler(error, nil);
      return;
    } else if (!data) {
      handler(nil, ret);
      return;
    }
    
    for (NSDictionary *d in data) {
      WKAssignment *assignment = [[WKAssignment alloc] init];
      assignment.id_p = [d[@"id"] intValue];
      assignment.level = [d[@"data"][@"level"] intValue];
      assignment.subjectId = [d[@"data"][@"subject_id"] intValue];
      assignment.srsStage = [d[@"data"][@"srs_stage"] intValue];
      
      if (d[@"data"][@"available_at"] != [NSNull null]) {
        assignment.availableAt =
            [[_dateFormatter dateFromString:d[@"data"][@"available_at"]] timeIntervalSince1970];
      }
      
      if (d[@"data"][@"started_at"] != [NSNull null]) {
        assignment.startedAt =
            [[_dateFormatter dateFromString:d[@"data"][@"started_at"]] timeIntervalSince1970];
      }
      
      if (d[@"data"][@"passed_at"] != [NSNull null]) {
        assignment.passedAt =
            [[_dateFormatter dateFromString:d[@"data"][@"passed_at"]] timeIntervalSince1970];
      }
      
      NSString *subjectType = d[@"data"][@"subject_type"];
      if ([subjectType isEqualToString:@"radical"]) {
        assignment.subjectType = WKSubject_Type_Radical;
      } else if ([subjectType isEqualToString:@"kanji"]) {
        assignment.subjectType = WKSubject_Type_Kanji;
      } else if ([subjectType isEqualToString:@"vocabulary"]) {
        assignment.subjectType = WKSubject_Type_Vocabulary;
      } else {
        NSAssert(false, @"Unknown subject type %@", subjectType);
      }
      [ret addObject:assignment];
    }
  }];
}

#pragma mark - Progress

- (void)fetchCSRFToken:(ProgressHandler)handler {
  NSURLRequest *req = [self authorizeUserRequest:[NSURL URLWithString:@(kReviewSessionURL)]];
  NSURLSessionDataTask *task =
    [_urlSession dataTaskWithRequest:req
                   completionHandler:^(NSData * _Nullable data,
                                       NSURLResponse * _Nullable response,
                                       NSError * _Nullable error) {
    if (error != nil) {
      handler(error);
      return;
    }
    NSHTTPURLResponse *httpResponse = (NSHTTPURLResponse *)response;
    if (httpResponse.statusCode != 200) {
      handler(MakeError((int)httpResponse.statusCode,
                        [NSString stringWithFormat:@"HTTP error %ld for %@",
                         (long)httpResponse.statusCode,
                         response.URL.absoluteString]));
      return;
    }
                                        
    NSString *body = [[NSString alloc] initWithData:data encoding:NSUTF8StringEncoding];
    NSTextCheckingResult *result = [kCSRFTokenRE firstMatchInString:body
                                                            options:0
                                                              range:NSMakeRange(0, body.length)];
    if (!result || result.range.location == NSNotFound) {
      NSLog(@"Page contents: %@", body);
      handler(MakeError(0, @"Progress token not found in page"));
      return;
    }
                                                  
    _csrfToken = [body substringWithRange:[result rangeAtIndex:1]];
    _csrfTokenUpdated = [NSDate date];
    NSLog(@"Got CSRF token %@", _csrfToken);
    handler(nil);
  }];
  [task resume];
}

- (bool)isCSRFTokenValid {
  return _csrfToken.length &&
      - [_csrfTokenUpdated timeIntervalSinceNow] < kCSRFTokenValidity;
}

- (void)ensureValidCSRFTokenAndThen:(void (^)(NSError * _Nullable))handler {
  // Fetch a new CSRF token if it's invalid or expired.
  if ([self isCSRFTokenValid]) {
    handler(nil);
  } else {
    [self fetchCSRFToken:handler];
  }
}

- (void)sendProgress:(NSArray<WKProgress *> *)progress
             handler:(ProgressHandler _Nullable)handler {
  if (progress.count == 0) {
    handler(nil);
    return;
  }
  
  [self ensureValidCSRFTokenAndThen:^(NSError * _Nullable error) {
    if (error != nil) {
      handler(error);
      return;
    }
    
    // Encode the data to send in the request.
    NSMutableArray<NSString *> *formParameters = [NSMutableArray array];
    for (WKProgress *p in progress) {
      [formParameters addObject:p.formParameters];
    }
    NSData *data = [[formParameters componentsJoinedByString:@"&"] dataUsingEncoding:NSUTF8StringEncoding];
    
    // Add the CSRF token and the data to the request.
    NSMutableURLRequest *req = [self authorizeUserRequest:[NSURL URLWithString:@(kProgressURL)]];
    [req addValue:_csrfToken forHTTPHeaderField:@"X-CSRF-Token"];
    [req addValue:@"application/x-www-form-urlencoded" forHTTPHeaderField:@"Content-Type"];
    [req addValue:[@(data.length) stringValue] forHTTPHeaderField:@"Content-Length"];
    req.HTTPMethod = @"PUT";
    req.HTTPBody = data;

    // Start the request.
    NSLog(@"PUT %@ to %@", [[NSString alloc] initWithData:data encoding:NSUTF8StringEncoding], req.URL);
    NSURLSessionDataTask *task =
      [_urlSession dataTaskWithRequest:req
                     completionHandler:^(NSData * _Nullable data,
                                         NSURLResponse * _Nullable response,
                                         NSError * _Nullable error) {
        if (handler) {
          handler(error);
        }
    }];
    [task resume];
  }];
}

#pragma mark - Study Materials

- (void)getStudyMaterialsModifiedAfter:(NSString *)date
                               handler:(StudyMaterialsHandler)handler {
  NSMutableArray<WKStudyMaterials *> *ret = [NSMutableArray array];

  NSURLComponents *url =
      [NSURLComponents componentsWithString:[NSString stringWithFormat:@"%s/study_materials",
                                             kURLBase]];
  if (date && date.length) {
    [url setQueryItems:@[
        [NSURLQueryItem queryItemWithName:@"updated_after" value:date],
    ]];
  }
  
  [self startPagedQueryFor:url.URL handler:^(NSArray *data, NSError *error) {
    if (error) {
      handler(error, nil);
      return;
    } else if (!data) {
      handler(nil, ret);
      return;
    }
    
    for (NSDictionary *d in data) {
      WKStudyMaterials *studyMaterials = [[WKStudyMaterials alloc] init];
      studyMaterials.id_p = [d[@"id"] intValue];
      studyMaterials.subjectId = [d[@"data"][@"subject_id"] intValue];
      studyMaterials.subjectType = d[@"data"][@"subject_type"];

      if (d[@"data"][@"meaning_note"] != [NSNull null]) {
        studyMaterials.meaningNote = d[@"data"][@"meaning_note"];
      }
      if (d[@"data"][@"reading_note"] != [NSNull null]) {
        studyMaterials.meaningNote = d[@"data"][@"reading_note"];
      }
      if (d[@"data"][@"meaning_synonyms"] != [NSNull null]) {
        studyMaterials.meaningSynonymsArray = d[@"data"][@"meaning_synonyms"];
      }
      [ret addObject:studyMaterials];
    }
  }];
}

- (void)updateStudyMaterial:(WKStudyMaterials *)material
                    handler:(UpdateStudyMaterialHandler)handler {
  [self ensureValidCSRFTokenAndThen:^(NSError * _Nullable error) {
    if (error != nil) {
      handler(error);
      return;
    }
    
    // Encode the data to send in the request.
    NSMutableDictionary *payload = [NSMutableDictionary dictionary];
    [payload setObject:@(material.subjectId) forKey:@"subject_id"];
    [payload setObject:material.subjectType forKey:@"subject_type"];
    [payload setObject:material.meaningSynonymsArray forKey:@"meaning_synonyms"];
    NSData *data = [NSJSONSerialization dataWithJSONObject:payload options:0 error:nil];
    
    // Add the CSRF token and the data to the request.
    NSString *urlString = [NSString stringWithFormat:@"%s/%d",
                           kStudyMaterialsURLBase, material.subjectId];
    NSMutableURLRequest *req = [self authorizeUserRequest:[NSURL URLWithString:urlString]];
    [req addValue:_csrfToken forHTTPHeaderField:@"X-CSRF-Token"];
    [req addValue:@"application/json" forHTTPHeaderField:@"Content-Type"];
    [req addValue:[@(data.length) stringValue] forHTTPHeaderField:@"Content-Length"];
    req.HTTPMethod = @"PUT";
    req.HTTPBody = data;
    
    // Start the request.
    NSLog(@"PUT %@ to %@", [[NSString alloc] initWithData:data encoding:NSUTF8StringEncoding], req.URL);
    NSURLSessionDataTask *task =
        [_urlSession dataTaskWithRequest:req
                       completionHandler:^(NSData * _Nullable data,
                                           NSURLResponse * _Nullable response,
                                           NSError * _Nullable error) {
                         if (handler) {
                           handler(error);
                         }
                       }];
    [task resume];
  }];
}

#pragma mark - User Info

- (void)getUserInfo:(UserInfoHandler)handler {
  NSURLComponents *url =
  [NSURLComponents componentsWithString:[NSString stringWithFormat:@"%s/user", kURLBase]];
  
  [self startPagedQueryFor:url.URL handler:^(NSDictionary *data, NSError *error) {
    if (error) {
      handler(error, nil);
      return;
    }
    
    WKUser *ret = [[WKUser alloc] init];
    ret.username = data[@"username"];
    ret.level = [data[@"level"] intValue];
    ret.maxLevelGrantedBySubscription = [data[@"max_level_granted_by_subscription"] intValue];
    ret.profileURL = data[@"profile_url"];
    ret.subscribed = [data[@"subscribed"] boolValue];
    
    if (data[@"started_at"] != [NSNull null]) {
      ret.startedAt = [[_dateFormatter dateFromString:data[@"started_at"]] timeIntervalSince1970];
    }
    
    handler(nil, ret);
  }];
}

@end

NS_ASSUME_NONNULL_END
