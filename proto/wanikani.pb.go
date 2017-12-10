// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wanikani.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	wanikani.proto

It has these top-level messages:
	Meaning
	Reading
	Radical
	Kanji
	Vocabulary
	Subject
	Assignment
	Progress
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Reading_Type int32

const (
	Reading_ONYOMI  Reading_Type = 1
	Reading_KUNYOMI Reading_Type = 2
	Reading_NANORI  Reading_Type = 3
)

var Reading_Type_name = map[int32]string{
	1: "ONYOMI",
	2: "KUNYOMI",
	3: "NANORI",
}
var Reading_Type_value = map[string]int32{
	"ONYOMI":  1,
	"KUNYOMI": 2,
	"NANORI":  3,
}

func (x Reading_Type) Enum() *Reading_Type {
	p := new(Reading_Type)
	*p = x
	return p
}
func (x Reading_Type) String() string {
	return proto1.EnumName(Reading_Type_name, int32(x))
}
func (x *Reading_Type) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Reading_Type_value, data, "Reading_Type")
	if err != nil {
		return err
	}
	*x = Reading_Type(value)
	return nil
}
func (Reading_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Vocabulary_PartOfSpeech int32

const (
	Vocabulary_NOUN              Vocabulary_PartOfSpeech = 1
	Vocabulary_NUMERAL           Vocabulary_PartOfSpeech = 2
	Vocabulary_INTRANSITIVE_VERB Vocabulary_PartOfSpeech = 3
	Vocabulary_ICHIDAN_VERB      Vocabulary_PartOfSpeech = 4
	Vocabulary_TRANSITIVE_VERB   Vocabulary_PartOfSpeech = 5
	Vocabulary_NO_ADJECTIVE      Vocabulary_PartOfSpeech = 6
	Vocabulary_GODAN_VERB        Vocabulary_PartOfSpeech = 7
	Vocabulary_NA_ADJECTIVE      Vocabulary_PartOfSpeech = 8
	Vocabulary_I_ADJECTIVE       Vocabulary_PartOfSpeech = 9
	Vocabulary_SUFFIX            Vocabulary_PartOfSpeech = 10
	Vocabulary_ADVERB            Vocabulary_PartOfSpeech = 11
	Vocabulary_SURU_VERB         Vocabulary_PartOfSpeech = 12
	Vocabulary_PREFIX            Vocabulary_PartOfSpeech = 13
	Vocabulary_PROPER_NOUN       Vocabulary_PartOfSpeech = 14
	Vocabulary_EXPRESSION        Vocabulary_PartOfSpeech = 15
	Vocabulary_ADJECTIVE         Vocabulary_PartOfSpeech = 16
	Vocabulary_INTERJECTION      Vocabulary_PartOfSpeech = 17
	Vocabulary_COUNTER           Vocabulary_PartOfSpeech = 18
	Vocabulary_PRONOUN           Vocabulary_PartOfSpeech = 19
	Vocabulary_CONJUNCTION       Vocabulary_PartOfSpeech = 20
)

var Vocabulary_PartOfSpeech_name = map[int32]string{
	1:  "NOUN",
	2:  "NUMERAL",
	3:  "INTRANSITIVE_VERB",
	4:  "ICHIDAN_VERB",
	5:  "TRANSITIVE_VERB",
	6:  "NO_ADJECTIVE",
	7:  "GODAN_VERB",
	8:  "NA_ADJECTIVE",
	9:  "I_ADJECTIVE",
	10: "SUFFIX",
	11: "ADVERB",
	12: "SURU_VERB",
	13: "PREFIX",
	14: "PROPER_NOUN",
	15: "EXPRESSION",
	16: "ADJECTIVE",
	17: "INTERJECTION",
	18: "COUNTER",
	19: "PRONOUN",
	20: "CONJUNCTION",
}
var Vocabulary_PartOfSpeech_value = map[string]int32{
	"NOUN":              1,
	"NUMERAL":           2,
	"INTRANSITIVE_VERB": 3,
	"ICHIDAN_VERB":      4,
	"TRANSITIVE_VERB":   5,
	"NO_ADJECTIVE":      6,
	"GODAN_VERB":        7,
	"NA_ADJECTIVE":      8,
	"I_ADJECTIVE":       9,
	"SUFFIX":            10,
	"ADVERB":            11,
	"SURU_VERB":         12,
	"PREFIX":            13,
	"PROPER_NOUN":       14,
	"EXPRESSION":        15,
	"ADJECTIVE":         16,
	"INTERJECTION":      17,
	"COUNTER":           18,
	"PRONOUN":           19,
	"CONJUNCTION":       20,
}

func (x Vocabulary_PartOfSpeech) Enum() *Vocabulary_PartOfSpeech {
	p := new(Vocabulary_PartOfSpeech)
	*p = x
	return p
}
func (x Vocabulary_PartOfSpeech) String() string {
	return proto1.EnumName(Vocabulary_PartOfSpeech_name, int32(x))
}
func (x *Vocabulary_PartOfSpeech) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Vocabulary_PartOfSpeech_value, data, "Vocabulary_PartOfSpeech")
	if err != nil {
		return err
	}
	*x = Vocabulary_PartOfSpeech(value)
	return nil
}
func (Vocabulary_PartOfSpeech) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type Subject_Type int32

const (
	Subject_RADICAL    Subject_Type = 1
	Subject_KANJI      Subject_Type = 2
	Subject_VOCABULARY Subject_Type = 3
)

var Subject_Type_name = map[int32]string{
	1: "RADICAL",
	2: "KANJI",
	3: "VOCABULARY",
}
var Subject_Type_value = map[string]int32{
	"RADICAL":    1,
	"KANJI":      2,
	"VOCABULARY": 3,
}

func (x Subject_Type) Enum() *Subject_Type {
	p := new(Subject_Type)
	*p = x
	return p
}
func (x Subject_Type) String() string {
	return proto1.EnumName(Subject_Type_name, int32(x))
}
func (x *Subject_Type) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(Subject_Type_value, data, "Subject_Type")
	if err != nil {
		return err
	}
	*x = Subject_Type(value)
	return nil
}
func (Subject_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type Meaning struct {
	Meaning          *string `protobuf:"bytes,1,opt,name=meaning" json:"meaning,omitempty"`
	IsPrimary        *bool   `protobuf:"varint,2,opt,name=is_primary" json:"is_primary,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Meaning) Reset()                    { *m = Meaning{} }
func (m *Meaning) String() string            { return proto1.CompactTextString(m) }
func (*Meaning) ProtoMessage()               {}
func (*Meaning) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Meaning) GetMeaning() string {
	if m != nil && m.Meaning != nil {
		return *m.Meaning
	}
	return ""
}

func (m *Meaning) GetIsPrimary() bool {
	if m != nil && m.IsPrimary != nil {
		return *m.IsPrimary
	}
	return false
}

type Reading struct {
	Reading          *string       `protobuf:"bytes,1,opt,name=reading" json:"reading,omitempty"`
	IsPrimary        *bool         `protobuf:"varint,2,opt,name=is_primary" json:"is_primary,omitempty"`
	Type             *Reading_Type `protobuf:"varint,3,opt,name=type,enum=proto.Reading_Type" json:"type,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Reading) Reset()                    { *m = Reading{} }
func (m *Reading) String() string            { return proto1.CompactTextString(m) }
func (*Reading) ProtoMessage()               {}
func (*Reading) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reading) GetReading() string {
	if m != nil && m.Reading != nil {
		return *m.Reading
	}
	return ""
}

func (m *Reading) GetIsPrimary() bool {
	if m != nil && m.IsPrimary != nil {
		return *m.IsPrimary
	}
	return false
}

func (m *Reading) GetType() Reading_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Reading_ONYOMI
}

type Radical struct {
	Japanese         *string    `protobuf:"bytes,1,opt,name=japanese" json:"japanese,omitempty"`
	Meanings         []*Meaning `protobuf:"bytes,2,rep,name=meanings" json:"meanings,omitempty"`
	CharacterImage   *string    `protobuf:"bytes,3,opt,name=character_image" json:"character_image,omitempty"`
	Mnemonic         *string    `protobuf:"bytes,4,opt,name=mnemonic" json:"mnemonic,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Radical) Reset()                    { *m = Radical{} }
func (m *Radical) String() string            { return proto1.CompactTextString(m) }
func (*Radical) ProtoMessage()               {}
func (*Radical) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Radical) GetJapanese() string {
	if m != nil && m.Japanese != nil {
		return *m.Japanese
	}
	return ""
}

func (m *Radical) GetMeanings() []*Meaning {
	if m != nil {
		return m.Meanings
	}
	return nil
}

func (m *Radical) GetCharacterImage() string {
	if m != nil && m.CharacterImage != nil {
		return *m.CharacterImage
	}
	return ""
}

func (m *Radical) GetMnemonic() string {
	if m != nil && m.Mnemonic != nil {
		return *m.Mnemonic
	}
	return ""
}

type Kanji struct {
	Japanese            *string    `protobuf:"bytes,1,opt,name=japanese" json:"japanese,omitempty"`
	Readings            []*Reading `protobuf:"bytes,2,rep,name=readings" json:"readings,omitempty"`
	Meanings            []*Meaning `protobuf:"bytes,3,rep,name=meanings" json:"meanings,omitempty"`
	MeaningMnemonic     *string    `protobuf:"bytes,4,opt,name=meaning_mnemonic" json:"meaning_mnemonic,omitempty"`
	MeaningHint         *string    `protobuf:"bytes,5,opt,name=meaning_hint" json:"meaning_hint,omitempty"`
	ReadingMnemonic     *string    `protobuf:"bytes,6,opt,name=reading_mnemonic" json:"reading_mnemonic,omitempty"`
	ReadingHint         *string    `protobuf:"bytes,7,opt,name=reading_hint" json:"reading_hint,omitempty"`
	ComponentSubjectIds []int32    `protobuf:"varint,8,rep,name=component_subject_ids" json:"component_subject_ids,omitempty"`
	XXX_unrecognized    []byte     `json:"-"`
}

func (m *Kanji) Reset()                    { *m = Kanji{} }
func (m *Kanji) String() string            { return proto1.CompactTextString(m) }
func (*Kanji) ProtoMessage()               {}
func (*Kanji) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Kanji) GetJapanese() string {
	if m != nil && m.Japanese != nil {
		return *m.Japanese
	}
	return ""
}

func (m *Kanji) GetReadings() []*Reading {
	if m != nil {
		return m.Readings
	}
	return nil
}

func (m *Kanji) GetMeanings() []*Meaning {
	if m != nil {
		return m.Meanings
	}
	return nil
}

func (m *Kanji) GetMeaningMnemonic() string {
	if m != nil && m.MeaningMnemonic != nil {
		return *m.MeaningMnemonic
	}
	return ""
}

func (m *Kanji) GetMeaningHint() string {
	if m != nil && m.MeaningHint != nil {
		return *m.MeaningHint
	}
	return ""
}

func (m *Kanji) GetReadingMnemonic() string {
	if m != nil && m.ReadingMnemonic != nil {
		return *m.ReadingMnemonic
	}
	return ""
}

func (m *Kanji) GetReadingHint() string {
	if m != nil && m.ReadingHint != nil {
		return *m.ReadingHint
	}
	return ""
}

func (m *Kanji) GetComponentSubjectIds() []int32 {
	if m != nil {
		return m.ComponentSubjectIds
	}
	return nil
}

type Vocabulary struct {
	Japanese            *string                   `protobuf:"bytes,1,opt,name=japanese" json:"japanese,omitempty"`
	Readings            []*Reading                `protobuf:"bytes,2,rep,name=readings" json:"readings,omitempty"`
	Meanings            []*Meaning                `protobuf:"bytes,3,rep,name=meanings" json:"meanings,omitempty"`
	MeaningExplanation  *string                   `protobuf:"bytes,4,opt,name=meaning_explanation" json:"meaning_explanation,omitempty"`
	ReadingExplanation  *string                   `protobuf:"bytes,5,opt,name=reading_explanation" json:"reading_explanation,omitempty"`
	Sentences           []*Vocabulary_Sentence    `protobuf:"bytes,6,rep,name=sentences" json:"sentences,omitempty"`
	ComponentSubjectIds []int32                   `protobuf:"varint,7,rep,name=component_subject_ids" json:"component_subject_ids,omitempty"`
	PartsOfSpeech       []Vocabulary_PartOfSpeech `protobuf:"varint,8,rep,name=parts_of_speech,enum=proto.Vocabulary_PartOfSpeech" json:"parts_of_speech,omitempty"`
	Audio               *string                   `protobuf:"bytes,9,opt,name=audio" json:"audio,omitempty"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *Vocabulary) Reset()                    { *m = Vocabulary{} }
func (m *Vocabulary) String() string            { return proto1.CompactTextString(m) }
func (*Vocabulary) ProtoMessage()               {}
func (*Vocabulary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Vocabulary) GetJapanese() string {
	if m != nil && m.Japanese != nil {
		return *m.Japanese
	}
	return ""
}

func (m *Vocabulary) GetReadings() []*Reading {
	if m != nil {
		return m.Readings
	}
	return nil
}

func (m *Vocabulary) GetMeanings() []*Meaning {
	if m != nil {
		return m.Meanings
	}
	return nil
}

func (m *Vocabulary) GetMeaningExplanation() string {
	if m != nil && m.MeaningExplanation != nil {
		return *m.MeaningExplanation
	}
	return ""
}

func (m *Vocabulary) GetReadingExplanation() string {
	if m != nil && m.ReadingExplanation != nil {
		return *m.ReadingExplanation
	}
	return ""
}

func (m *Vocabulary) GetSentences() []*Vocabulary_Sentence {
	if m != nil {
		return m.Sentences
	}
	return nil
}

func (m *Vocabulary) GetComponentSubjectIds() []int32 {
	if m != nil {
		return m.ComponentSubjectIds
	}
	return nil
}

func (m *Vocabulary) GetPartsOfSpeech() []Vocabulary_PartOfSpeech {
	if m != nil {
		return m.PartsOfSpeech
	}
	return nil
}

func (m *Vocabulary) GetAudio() string {
	if m != nil && m.Audio != nil {
		return *m.Audio
	}
	return ""
}

type Vocabulary_Sentence struct {
	Japanese         *string `protobuf:"bytes,1,opt,name=japanese" json:"japanese,omitempty"`
	English          *string `protobuf:"bytes,2,opt,name=english" json:"english,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Vocabulary_Sentence) Reset()                    { *m = Vocabulary_Sentence{} }
func (m *Vocabulary_Sentence) String() string            { return proto1.CompactTextString(m) }
func (*Vocabulary_Sentence) ProtoMessage()               {}
func (*Vocabulary_Sentence) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *Vocabulary_Sentence) GetJapanese() string {
	if m != nil && m.Japanese != nil {
		return *m.Japanese
	}
	return ""
}

func (m *Vocabulary_Sentence) GetEnglish() string {
	if m != nil && m.English != nil {
		return *m.English
	}
	return ""
}

type Subject struct {
	Id               *int32      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Level            *int32      `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
	Slug             *string     `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty"`
	DocumentUrl      *string     `protobuf:"bytes,4,opt,name=document_url" json:"document_url,omitempty"`
	Radical          *Radical    `protobuf:"bytes,5,opt,name=radical" json:"radical,omitempty"`
	Kanji            *Kanji      `protobuf:"bytes,6,opt,name=kanji" json:"kanji,omitempty"`
	Vocabulary       *Vocabulary `protobuf:"bytes,7,opt,name=vocabulary" json:"vocabulary,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Subject) Reset()                    { *m = Subject{} }
func (m *Subject) String() string            { return proto1.CompactTextString(m) }
func (*Subject) ProtoMessage()               {}
func (*Subject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Subject) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Subject) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *Subject) GetSlug() string {
	if m != nil && m.Slug != nil {
		return *m.Slug
	}
	return ""
}

func (m *Subject) GetDocumentUrl() string {
	if m != nil && m.DocumentUrl != nil {
		return *m.DocumentUrl
	}
	return ""
}

func (m *Subject) GetRadical() *Radical {
	if m != nil {
		return m.Radical
	}
	return nil
}

func (m *Subject) GetKanji() *Kanji {
	if m != nil {
		return m.Kanji
	}
	return nil
}

func (m *Subject) GetVocabulary() *Vocabulary {
	if m != nil {
		return m.Vocabulary
	}
	return nil
}

type Assignment struct {
	Id               *int32        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Level            *int32        `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
	SubjectId        *int32        `protobuf:"varint,3,opt,name=subject_id" json:"subject_id,omitempty"`
	SubjectType      *Subject_Type `protobuf:"varint,4,opt,name=subject_type,enum=proto.Subject_Type" json:"subject_type,omitempty"`
	AvailableAt      *int32        `protobuf:"varint,5,opt,name=available_at" json:"available_at,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Assignment) Reset()                    { *m = Assignment{} }
func (m *Assignment) String() string            { return proto1.CompactTextString(m) }
func (*Assignment) ProtoMessage()               {}
func (*Assignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Assignment) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Assignment) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *Assignment) GetSubjectId() int32 {
	if m != nil && m.SubjectId != nil {
		return *m.SubjectId
	}
	return 0
}

func (m *Assignment) GetSubjectType() Subject_Type {
	if m != nil && m.SubjectType != nil {
		return *m.SubjectType
	}
	return Subject_RADICAL
}

func (m *Assignment) GetAvailableAt() int32 {
	if m != nil && m.AvailableAt != nil {
		return *m.AvailableAt
	}
	return 0
}

type Progress struct {
	Id               *int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	MeaningWrong     *bool  `protobuf:"varint,2,opt,name=meaning_wrong" json:"meaning_wrong,omitempty"`
	ReadingWrong     *bool  `protobuf:"varint,3,opt,name=reading_wrong" json:"reading_wrong,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Progress) Reset()                    { *m = Progress{} }
func (m *Progress) String() string            { return proto1.CompactTextString(m) }
func (*Progress) ProtoMessage()               {}
func (*Progress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Progress) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Progress) GetMeaningWrong() bool {
	if m != nil && m.MeaningWrong != nil {
		return *m.MeaningWrong
	}
	return false
}

func (m *Progress) GetReadingWrong() bool {
	if m != nil && m.ReadingWrong != nil {
		return *m.ReadingWrong
	}
	return false
}

func init() {
	proto1.RegisterType((*Meaning)(nil), "proto.Meaning")
	proto1.RegisterType((*Reading)(nil), "proto.Reading")
	proto1.RegisterType((*Radical)(nil), "proto.Radical")
	proto1.RegisterType((*Kanji)(nil), "proto.Kanji")
	proto1.RegisterType((*Vocabulary)(nil), "proto.Vocabulary")
	proto1.RegisterType((*Vocabulary_Sentence)(nil), "proto.Vocabulary.Sentence")
	proto1.RegisterType((*Subject)(nil), "proto.Subject")
	proto1.RegisterType((*Assignment)(nil), "proto.Assignment")
	proto1.RegisterType((*Progress)(nil), "proto.Progress")
	proto1.RegisterEnum("proto.Reading_Type", Reading_Type_name, Reading_Type_value)
	proto1.RegisterEnum("proto.Vocabulary_PartOfSpeech", Vocabulary_PartOfSpeech_name, Vocabulary_PartOfSpeech_value)
	proto1.RegisterEnum("proto.Subject_Type", Subject_Type_name, Subject_Type_value)
}

func init() { proto1.RegisterFile("wanikani.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 835 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x86, 0x7e, 0x68, 0x52, 0xa3, 0x3f, 0x7a, 0x65, 0xa3, 0x44, 0x8c, 0xb6, 0x2a, 0x81, 0x02,
	0x2a, 0x8a, 0xe8, 0xa0, 0x4b, 0xcf, 0x8c, 0xc4, 0xb4, 0xf4, 0x0f, 0x29, 0x90, 0xa6, 0x91, 0x9c,
	0x88, 0x35, 0xb5, 0x91, 0xd6, 0xa1, 0x96, 0x04, 0x49, 0x39, 0x35, 0x50, 0xf4, 0x89, 0xfa, 0x20,
	0x7d, 0x88, 0xf6, 0x5d, 0x8a, 0x5d, 0x2e, 0x6d, 0x29, 0x8e, 0x7b, 0xcb, 0x89, 0xcb, 0x99, 0x6f,
	0x67, 0xe6, 0x9b, 0x9d, 0xf9, 0x60, 0xf0, 0x09, 0x33, 0xfa, 0x11, 0x33, 0x3a, 0xcd, 0xf2, 0xb4,
	0x4c, 0x91, 0x22, 0x3e, 0xe6, 0x14, 0xd4, 0x2b, 0x82, 0x19, 0x65, 0x6b, 0x34, 0x04, 0x75, 0x5b,
	0x1d, 0x8d, 0xc6, 0xb8, 0x31, 0xe9, 0x20, 0x04, 0x40, 0x8b, 0x28, 0xcb, 0xe9, 0x16, 0xe7, 0x0f,
	0x46, 0x73, 0xdc, 0x98, 0x68, 0xe6, 0x9f, 0xa0, 0xfa, 0x04, 0xaf, 0x24, 0x3e, 0xaf, 0x8e, 0x2f,
	0xe3, 0xd1, 0x0f, 0xd0, 0x2e, 0x1f, 0x32, 0x62, 0xb4, 0xc6, 0x8d, 0xc9, 0x60, 0x36, 0xaa, 0x92,
	0x4f, 0x65, 0x88, 0xe9, 0xf5, 0x43, 0x46, 0xcc, 0x9f, 0xa1, 0xcd, 0xbf, 0x08, 0xe0, 0xc8, 0x73,
	0xdf, 0x7b, 0x57, 0x8e, 0xde, 0x40, 0x5d, 0x50, 0x2f, 0xc2, 0xea, 0xa7, 0xc9, 0x1d, 0xae, 0xe5,
	0x7a, 0xbe, 0xa3, 0xb7, 0xcc, 0x0d, 0xa8, 0x3e, 0x5e, 0xd1, 0x18, 0x27, 0x48, 0x07, 0xed, 0x0e,
	0x67, 0x98, 0x91, 0x82, 0xc8, 0x02, 0xc6, 0xa0, 0x49, 0x06, 0x85, 0xd1, 0x1c, 0xb7, 0x26, 0xdd,
	0xd9, 0x40, 0x26, 0xac, 0x39, 0x7e, 0x03, 0xc3, 0x78, 0x83, 0x73, 0x1c, 0x97, 0x24, 0x8f, 0xe8,
	0x16, 0xaf, 0xab, 0xca, 0x3a, 0x3c, 0xd8, 0x96, 0x91, 0x6d, 0xca, 0x68, 0x6c, 0xb4, 0xb9, 0xc5,
	0xfc, 0xa7, 0x01, 0xca, 0x05, 0x66, 0x77, 0xf4, 0xcb, 0x89, 0x24, 0xf5, 0xcf, 0x13, 0xd5, 0xcd,
	0xd9, 0x2f, 0xa5, 0xf5, 0xc5, 0x52, 0x0c, 0xd0, 0x25, 0x22, 0x3a, 0xcc, 0x8c, 0x4e, 0xa0, 0x57,
	0x7b, 0x36, 0x94, 0x95, 0x86, 0x22, 0xac, 0x06, 0xe8, 0x32, 0xe7, 0x13, 0xfe, 0xa8, 0xc6, 0xd7,
	0x1e, 0x81, 0x57, 0x85, 0xf5, 0x5b, 0x38, 0x8d, 0xd3, 0x6d, 0x96, 0x32, 0xc2, 0xca, 0xa8, 0xd8,
	0xdd, 0xde, 0x91, 0xb8, 0x8c, 0xe8, 0xaa, 0x30, 0xb4, 0x71, 0x6b, 0xa2, 0x98, 0x7f, 0x29, 0x00,
	0x37, 0x69, 0x8c, 0x6f, 0x77, 0x09, 0xce, 0x1f, 0xbe, 0x12, 0xc7, 0x33, 0x18, 0xd5, 0x4c, 0xc8,
	0xef, 0x59, 0x82, 0x19, 0x2e, 0x69, 0xca, 0x24, 0xcd, 0x33, 0x18, 0xd5, 0x65, 0xef, 0x3b, 0x2b,
	0xb6, 0xaf, 0xa1, 0x53, 0x10, 0x56, 0x12, 0x16, 0x93, 0xc2, 0x38, 0x12, 0xc1, 0x5f, 0xc9, 0xe0,
	0x4f, 0x55, 0x4f, 0x03, 0x09, 0x79, 0x99, 0xac, 0xca, 0xc9, 0xa2, 0x5f, 0x60, 0x98, 0xe1, 0xbc,
	0x2c, 0xa2, 0xf4, 0x43, 0x54, 0x64, 0x84, 0xc4, 0x1b, 0xd1, 0x85, 0xc1, 0xec, 0xbb, 0xe7, 0x31,
	0x97, 0x38, 0x2f, 0xbd, 0x0f, 0x81, 0x40, 0xa1, 0x3e, 0x28, 0x78, 0xb7, 0xa2, 0xa9, 0xd1, 0xe1,
	0x55, 0xbd, 0x7a, 0x0d, 0xda, 0x63, 0xca, 0xe7, 0x1d, 0x1b, 0x82, 0x4a, 0xd8, 0x3a, 0xa1, 0xc5,
	0x46, 0x0c, 0x7f, 0xc7, 0xfc, 0xbb, 0x09, 0xbd, 0x83, 0x70, 0x1a, 0xb4, 0x5d, 0x2f, 0x74, 0xab,
	0x01, 0x77, 0xc3, 0x2b, 0xdb, 0xb7, 0x2e, 0xf5, 0x26, 0x3a, 0x85, 0x63, 0xc7, 0xbd, 0xf6, 0x2d,
	0x37, 0x70, 0xae, 0x9d, 0x1b, 0x3b, 0xba, 0xb1, 0xfd, 0x37, 0x7a, 0x0b, 0xe9, 0xd0, 0x73, 0xe6,
	0xbf, 0x39, 0x0b, 0xcb, 0xad, 0x2c, 0x6d, 0x34, 0x82, 0xe1, 0xe7, 0x30, 0x85, 0xc3, 0x5c, 0x2f,
	0xb2, 0x16, 0xe7, 0xf6, 0x9c, 0x9b, 0xf5, 0x23, 0x34, 0x00, 0xf8, 0xd5, 0x7b, 0xbc, 0xa6, 0x0a,
	0x84, 0xb5, 0x87, 0xd0, 0xd0, 0x10, 0xba, 0xce, 0x9e, 0xa1, 0xc3, 0x77, 0x2c, 0x08, 0xdf, 0xbe,
	0x75, 0xde, 0xe9, 0xc0, 0xcf, 0xd6, 0x42, 0x5c, 0xed, 0xa2, 0x3e, 0x74, 0x82, 0xd0, 0x0f, 0xab,
	0x48, 0x3d, 0xee, 0x5a, 0xfa, 0x36, 0x87, 0xf5, 0x79, 0x8c, 0xa5, 0xef, 0x2d, 0x6d, 0x3f, 0x12,
	0x9c, 0x06, 0x3c, 0xad, 0xfd, 0x6e, 0xe9, 0xdb, 0x41, 0xe0, 0x78, 0xae, 0x3e, 0xe4, 0x77, 0x9f,
	0x52, 0xe8, 0x82, 0x8e, 0x7b, 0x6d, 0xfb, 0xc2, 0xe2, 0xb9, 0xfa, 0x31, 0x6f, 0xc2, 0xdc, 0x0b,
	0xb9, 0x4d, 0x47, 0xfc, 0x67, 0xe9, 0x7b, 0x22, 0xd4, 0x88, 0xc7, 0x9e, 0x7b, 0xee, 0x79, 0xe8,
	0x56, 0xd0, 0x13, 0xf3, 0xdf, 0x06, 0xa8, 0x41, 0xf5, 0xae, 0x08, 0xa0, 0x49, 0x57, 0xa2, 0xe7,
	0x0a, 0x7f, 0xa0, 0x84, 0xdc, 0x93, 0x44, 0x74, 0x5c, 0x41, 0x3d, 0x68, 0x17, 0xc9, 0x6e, 0x2d,
	0x97, 0xfa, 0x04, 0x7a, 0xab, 0x34, 0xde, 0x6d, 0xf9, 0x50, 0xec, 0xf2, 0x44, 0xce, 0xdd, 0xf7,
	0xa0, 0xe6, 0x95, 0x84, 0x88, 0x59, 0xdb, 0x9b, 0x6b, 0x29, 0x2c, 0x67, 0xa0, 0x7c, 0xe4, 0x8b,
	0x2f, 0xd6, 0xab, 0x3b, 0xeb, 0x49, 0x77, 0x25, 0x06, 0x3f, 0x02, 0xdc, 0x3f, 0x0e, 0x8b, 0x58,
	0xb5, 0xee, 0xec, 0xf8, 0xd9, 0x14, 0x99, 0x53, 0x29, 0x6a, 0x5d, 0x50, 0x7d, 0x6b, 0xe1, 0xcc,
	0xad, 0x4b, 0xbd, 0x81, 0x3a, 0xa0, 0x5c, 0x58, 0xee, 0x39, 0xd7, 0xb4, 0x01, 0xc0, 0x8d, 0x37,
	0xb7, 0xde, 0x84, 0x97, 0x96, 0xff, 0x5e, 0x6f, 0x99, 0x7f, 0x00, 0x58, 0x45, 0x41, 0xd7, 0x8c,
	0x17, 0xfb, 0x7f, 0x0c, 0x11, 0xc0, 0xd3, 0x7c, 0x0b, 0x9e, 0x0a, 0xfa, 0x09, 0x7a, 0xb5, 0x4d,
	0x88, 0x6d, 0xfb, 0x40, 0x6c, 0x65, 0xdb, 0x84, 0xd8, 0xf2, 0x96, 0xe0, 0x7b, 0x4c, 0x13, 0x7c,
	0x9b, 0x90, 0x08, 0x57, 0xda, 0xa2, 0x98, 0x0b, 0xd0, 0x96, 0x79, 0xba, 0xce, 0x49, 0x51, 0x1c,
	0xe4, 0x3e, 0x85, 0x7e, 0xbd, 0xbf, 0x9f, 0xf2, 0x94, 0xad, 0xa5, 0xa8, 0x9f, 0x42, 0xbf, 0xde,
	0xdc, 0xca, 0xcc, 0xcb, 0xd0, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xda, 0xfb, 0xd6, 0xd0, 0x63,
	0x06, 0x00, 0x00,
}
