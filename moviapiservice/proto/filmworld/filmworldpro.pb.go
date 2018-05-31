// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/filmworld/filmworldpro.proto

package go_micro_srv_filmworldservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestAllMovies struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestAllMovies) Reset()         { *m = RequestAllMovies{} }
func (m *RequestAllMovies) String() string { return proto.CompactTextString(m) }
func (*RequestAllMovies) ProtoMessage()    {}
func (*RequestAllMovies) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{0}
}
func (m *RequestAllMovies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAllMovies.Unmarshal(m, b)
}
func (m *RequestAllMovies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAllMovies.Marshal(b, m, deterministic)
}
func (dst *RequestAllMovies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAllMovies.Merge(dst, src)
}
func (m *RequestAllMovies) XXX_Size() int {
	return xxx_messageInfo_RequestAllMovies.Size(m)
}
func (m *RequestAllMovies) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAllMovies.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAllMovies proto.InternalMessageInfo

type RequestMovie struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMovie) Reset()         { *m = RequestMovie{} }
func (m *RequestMovie) String() string { return proto.CompactTextString(m) }
func (*RequestMovie) ProtoMessage()    {}
func (*RequestMovie) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{1}
}
func (m *RequestMovie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMovie.Unmarshal(m, b)
}
func (m *RequestMovie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMovie.Marshal(b, m, deterministic)
}
func (dst *RequestMovie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMovie.Merge(dst, src)
}
func (m *RequestMovie) XXX_Size() int {
	return xxx_messageInfo_RequestMovie.Size(m)
}
func (m *RequestMovie) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMovie.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMovie proto.InternalMessageInfo

func (m *RequestMovie) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ResponseAllMovies struct {
	Movie                *MovieSimpleProto `protobuf:"bytes,1,opt,name=Movie" json:"Movie,omitempty"`
	RecordCount          int32             `protobuf:"varint,2,opt,name=RecordCount" json:"RecordCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResponseAllMovies) Reset()         { *m = ResponseAllMovies{} }
func (m *ResponseAllMovies) String() string { return proto.CompactTextString(m) }
func (*ResponseAllMovies) ProtoMessage()    {}
func (*ResponseAllMovies) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{2}
}
func (m *ResponseAllMovies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseAllMovies.Unmarshal(m, b)
}
func (m *ResponseAllMovies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseAllMovies.Marshal(b, m, deterministic)
}
func (dst *ResponseAllMovies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseAllMovies.Merge(dst, src)
}
func (m *ResponseAllMovies) XXX_Size() int {
	return xxx_messageInfo_ResponseAllMovies.Size(m)
}
func (m *ResponseAllMovies) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseAllMovies.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseAllMovies proto.InternalMessageInfo

func (m *ResponseAllMovies) GetMovie() *MovieSimpleProto {
	if m != nil {
		return m.Movie
	}
	return nil
}

func (m *ResponseAllMovies) GetRecordCount() int32 {
	if m != nil {
		return m.RecordCount
	}
	return 0
}

type ResponseMovie struct {
	Movie                *MovieProto `protobuf:"bytes,1,opt,name=Movie" json:"Movie,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ResponseMovie) Reset()         { *m = ResponseMovie{} }
func (m *ResponseMovie) String() string { return proto.CompactTextString(m) }
func (*ResponseMovie) ProtoMessage()    {}
func (*ResponseMovie) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{3}
}
func (m *ResponseMovie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMovie.Unmarshal(m, b)
}
func (m *ResponseMovie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMovie.Marshal(b, m, deterministic)
}
func (dst *ResponseMovie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMovie.Merge(dst, src)
}
func (m *ResponseMovie) XXX_Size() int {
	return xxx_messageInfo_ResponseMovie.Size(m)
}
func (m *ResponseMovie) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMovie.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMovie proto.InternalMessageInfo

func (m *ResponseMovie) GetMovie() *MovieProto {
	if m != nil {
		return m.Movie
	}
	return nil
}

type ResponseMovieStream struct {
	Movie                *MovieSimpleProto `protobuf:"bytes,1,opt,name=Movie" json:"Movie,omitempty"`
	RecordCount          int32             `protobuf:"varint,2,opt,name=RecordCount" json:"RecordCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResponseMovieStream) Reset()         { *m = ResponseMovieStream{} }
func (m *ResponseMovieStream) String() string { return proto.CompactTextString(m) }
func (*ResponseMovieStream) ProtoMessage()    {}
func (*ResponseMovieStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{4}
}
func (m *ResponseMovieStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMovieStream.Unmarshal(m, b)
}
func (m *ResponseMovieStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMovieStream.Marshal(b, m, deterministic)
}
func (dst *ResponseMovieStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMovieStream.Merge(dst, src)
}
func (m *ResponseMovieStream) XXX_Size() int {
	return xxx_messageInfo_ResponseMovieStream.Size(m)
}
func (m *ResponseMovieStream) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMovieStream.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMovieStream proto.InternalMessageInfo

func (m *ResponseMovieStream) GetMovie() *MovieSimpleProto {
	if m != nil {
		return m.Movie
	}
	return nil
}

func (m *ResponseMovieStream) GetRecordCount() int32 {
	if m != nil {
		return m.RecordCount
	}
	return 0
}

type MovieSimpleProto struct {
	Id                   int32         `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Title                string        `protobuf:"bytes,2,opt,name=Title" json:"Title,omitempty"`
	ReleaseDate          string        `protobuf:"bytes,3,opt,name=ReleaseDate" json:"ReleaseDate,omitempty"`
	MovieType            string        `protobuf:"bytes,4,opt,name=MovieType" json:"MovieType,omitempty"`
	Language             string        `protobuf:"bytes,5,opt,name=Language" json:"Language,omitempty"`
	Price                float32       `protobuf:"fixed32,6,opt,name=Price" json:"Price,omitempty"`
	Genre                string        `protobuf:"bytes,7,opt,name=Genre" json:"Genre,omitempty"`
	Director             string        `protobuf:"bytes,8,opt,name=Director" json:"Director,omitempty"`
	Actors               []*ActorProto `protobuf:"bytes,9,rep,name=Actors" json:"Actors,omitempty"`
	Provider             string        `protobuf:"bytes,10,opt,name=provider" json:"provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MovieSimpleProto) Reset()         { *m = MovieSimpleProto{} }
func (m *MovieSimpleProto) String() string { return proto.CompactTextString(m) }
func (*MovieSimpleProto) ProtoMessage()    {}
func (*MovieSimpleProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{5}
}
func (m *MovieSimpleProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieSimpleProto.Unmarshal(m, b)
}
func (m *MovieSimpleProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieSimpleProto.Marshal(b, m, deterministic)
}
func (dst *MovieSimpleProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieSimpleProto.Merge(dst, src)
}
func (m *MovieSimpleProto) XXX_Size() int {
	return xxx_messageInfo_MovieSimpleProto.Size(m)
}
func (m *MovieSimpleProto) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieSimpleProto.DiscardUnknown(m)
}

var xxx_messageInfo_MovieSimpleProto proto.InternalMessageInfo

func (m *MovieSimpleProto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MovieSimpleProto) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MovieSimpleProto) GetReleaseDate() string {
	if m != nil {
		return m.ReleaseDate
	}
	return ""
}

func (m *MovieSimpleProto) GetMovieType() string {
	if m != nil {
		return m.MovieType
	}
	return ""
}

func (m *MovieSimpleProto) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *MovieSimpleProto) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *MovieSimpleProto) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

func (m *MovieSimpleProto) GetDirector() string {
	if m != nil {
		return m.Director
	}
	return ""
}

func (m *MovieSimpleProto) GetActors() []*ActorProto {
	if m != nil {
		return m.Actors
	}
	return nil
}

func (m *MovieSimpleProto) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

type MovieProto struct {
	Id                   int32                `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=Title" json:"Title,omitempty"`
	ReleaseDate          string               `protobuf:"bytes,3,opt,name=ReleaseDate" json:"ReleaseDate,omitempty"`
	TypeId               int32                `protobuf:"varint,4,opt,name=TypeId" json:"TypeId,omitempty"`
	LanguageId           int32                `protobuf:"varint,5,opt,name=LanguageId" json:"LanguageId,omitempty"`
	Price                float32              `protobuf:"fixed32,6,opt,name=Price" json:"Price,omitempty"`
	GenreId              int32                `protobuf:"varint,7,opt,name=GenreId" json:"GenreId,omitempty"`
	DirectorId           int32                `protobuf:"varint,8,opt,name=DirectorId" json:"DirectorId,omitempty"`
	Actors               []*ActorProto        `protobuf:"bytes,9,rep,name=Actors" json:"Actors,omitempty"`
	MovieDirector        *DirectorProto       `protobuf:"bytes,10,opt,name=MovieDirector" json:"MovieDirector,omitempty"`
	MovieType            *MovieTypeProto      `protobuf:"bytes,11,opt,name=MovieType" json:"MovieType,omitempty"`
	Language             *MovieLanguagesProto `protobuf:"bytes,12,opt,name=Language" json:"Language,omitempty"`
	Genre                *MovieGenresProto    `protobuf:"bytes,13,opt,name=Genre" json:"Genre,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MovieProto) Reset()         { *m = MovieProto{} }
func (m *MovieProto) String() string { return proto.CompactTextString(m) }
func (*MovieProto) ProtoMessage()    {}
func (*MovieProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{6}
}
func (m *MovieProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieProto.Unmarshal(m, b)
}
func (m *MovieProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieProto.Marshal(b, m, deterministic)
}
func (dst *MovieProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieProto.Merge(dst, src)
}
func (m *MovieProto) XXX_Size() int {
	return xxx_messageInfo_MovieProto.Size(m)
}
func (m *MovieProto) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieProto.DiscardUnknown(m)
}

var xxx_messageInfo_MovieProto proto.InternalMessageInfo

func (m *MovieProto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MovieProto) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MovieProto) GetReleaseDate() string {
	if m != nil {
		return m.ReleaseDate
	}
	return ""
}

func (m *MovieProto) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *MovieProto) GetLanguageId() int32 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

func (m *MovieProto) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *MovieProto) GetGenreId() int32 {
	if m != nil {
		return m.GenreId
	}
	return 0
}

func (m *MovieProto) GetDirectorId() int32 {
	if m != nil {
		return m.DirectorId
	}
	return 0
}

func (m *MovieProto) GetActors() []*ActorProto {
	if m != nil {
		return m.Actors
	}
	return nil
}

func (m *MovieProto) GetMovieDirector() *DirectorProto {
	if m != nil {
		return m.MovieDirector
	}
	return nil
}

func (m *MovieProto) GetMovieType() *MovieTypeProto {
	if m != nil {
		return m.MovieType
	}
	return nil
}

func (m *MovieProto) GetLanguage() *MovieLanguagesProto {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *MovieProto) GetGenre() *MovieGenresProto {
	if m != nil {
		return m.Genre
	}
	return nil
}

type ActorProto struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=LastName" json:"LastName,omitempty"`
	GenderId             int32    `protobuf:"varint,5,opt,name=GenderId" json:"GenderId,omitempty"`
	Gender               string   `protobuf:"bytes,6,opt,name=Gender" json:"Gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActorProto) Reset()         { *m = ActorProto{} }
func (m *ActorProto) String() string { return proto.CompactTextString(m) }
func (*ActorProto) ProtoMessage()    {}
func (*ActorProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{7}
}
func (m *ActorProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActorProto.Unmarshal(m, b)
}
func (m *ActorProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActorProto.Marshal(b, m, deterministic)
}
func (dst *ActorProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActorProto.Merge(dst, src)
}
func (m *ActorProto) XXX_Size() int {
	return xxx_messageInfo_ActorProto.Size(m)
}
func (m *ActorProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ActorProto.DiscardUnknown(m)
}

var xxx_messageInfo_ActorProto proto.InternalMessageInfo

func (m *ActorProto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ActorProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActorProto) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *ActorProto) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *ActorProto) GetGenderId() int32 {
	if m != nil {
		return m.GenderId
	}
	return 0
}

func (m *ActorProto) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

type DirectorProto struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=LastName" json:"LastName,omitempty"`
	GenderId             int32    `protobuf:"varint,5,opt,name=GenderId" json:"GenderId,omitempty"`
	Gender               string   `protobuf:"bytes,6,opt,name=Gender" json:"Gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirectorProto) Reset()         { *m = DirectorProto{} }
func (m *DirectorProto) String() string { return proto.CompactTextString(m) }
func (*DirectorProto) ProtoMessage()    {}
func (*DirectorProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{8}
}
func (m *DirectorProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectorProto.Unmarshal(m, b)
}
func (m *DirectorProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectorProto.Marshal(b, m, deterministic)
}
func (dst *DirectorProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectorProto.Merge(dst, src)
}
func (m *DirectorProto) XXX_Size() int {
	return xxx_messageInfo_DirectorProto.Size(m)
}
func (m *DirectorProto) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectorProto.DiscardUnknown(m)
}

var xxx_messageInfo_DirectorProto proto.InternalMessageInfo

func (m *DirectorProto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DirectorProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DirectorProto) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *DirectorProto) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *DirectorProto) GetGenderId() int32 {
	if m != nil {
		return m.GenderId
	}
	return 0
}

func (m *DirectorProto) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

type MovieTypeProto struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MovieTypeProto) Reset()         { *m = MovieTypeProto{} }
func (m *MovieTypeProto) String() string { return proto.CompactTextString(m) }
func (*MovieTypeProto) ProtoMessage()    {}
func (*MovieTypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{9}
}
func (m *MovieTypeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieTypeProto.Unmarshal(m, b)
}
func (m *MovieTypeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieTypeProto.Marshal(b, m, deterministic)
}
func (dst *MovieTypeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieTypeProto.Merge(dst, src)
}
func (m *MovieTypeProto) XXX_Size() int {
	return xxx_messageInfo_MovieTypeProto.Size(m)
}
func (m *MovieTypeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieTypeProto.DiscardUnknown(m)
}

var xxx_messageInfo_MovieTypeProto proto.InternalMessageInfo

func (m *MovieTypeProto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MovieTypeProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MovieLanguagesProto struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Language             string   `protobuf:"bytes,2,opt,name=Language" json:"Language,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MovieLanguagesProto) Reset()         { *m = MovieLanguagesProto{} }
func (m *MovieLanguagesProto) String() string { return proto.CompactTextString(m) }
func (*MovieLanguagesProto) ProtoMessage()    {}
func (*MovieLanguagesProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{10}
}
func (m *MovieLanguagesProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieLanguagesProto.Unmarshal(m, b)
}
func (m *MovieLanguagesProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieLanguagesProto.Marshal(b, m, deterministic)
}
func (dst *MovieLanguagesProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieLanguagesProto.Merge(dst, src)
}
func (m *MovieLanguagesProto) XXX_Size() int {
	return xxx_messageInfo_MovieLanguagesProto.Size(m)
}
func (m *MovieLanguagesProto) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieLanguagesProto.DiscardUnknown(m)
}

var xxx_messageInfo_MovieLanguagesProto proto.InternalMessageInfo

func (m *MovieLanguagesProto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MovieLanguagesProto) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

type MovieGenresProto struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Genre                string   `protobuf:"bytes,2,opt,name=Genre" json:"Genre,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MovieGenresProto) Reset()         { *m = MovieGenresProto{} }
func (m *MovieGenresProto) String() string { return proto.CompactTextString(m) }
func (*MovieGenresProto) ProtoMessage()    {}
func (*MovieGenresProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_filmworldpro_ac66f8145b208ca7, []int{11}
}
func (m *MovieGenresProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieGenresProto.Unmarshal(m, b)
}
func (m *MovieGenresProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieGenresProto.Marshal(b, m, deterministic)
}
func (dst *MovieGenresProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieGenresProto.Merge(dst, src)
}
func (m *MovieGenresProto) XXX_Size() int {
	return xxx_messageInfo_MovieGenresProto.Size(m)
}
func (m *MovieGenresProto) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieGenresProto.DiscardUnknown(m)
}

var xxx_messageInfo_MovieGenresProto proto.InternalMessageInfo

func (m *MovieGenresProto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MovieGenresProto) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestAllMovies)(nil), "go.micro.srv.filmworldservice.RequestAllMovies")
	proto.RegisterType((*RequestMovie)(nil), "go.micro.srv.filmworldservice.RequestMovie")
	proto.RegisterType((*ResponseAllMovies)(nil), "go.micro.srv.filmworldservice.ResponseAllMovies")
	proto.RegisterType((*ResponseMovie)(nil), "go.micro.srv.filmworldservice.ResponseMovie")
	proto.RegisterType((*ResponseMovieStream)(nil), "go.micro.srv.filmworldservice.ResponseMovieStream")
	proto.RegisterType((*MovieSimpleProto)(nil), "go.micro.srv.filmworldservice.MovieSimpleProto")
	proto.RegisterType((*MovieProto)(nil), "go.micro.srv.filmworldservice.MovieProto")
	proto.RegisterType((*ActorProto)(nil), "go.micro.srv.filmworldservice.ActorProto")
	proto.RegisterType((*DirectorProto)(nil), "go.micro.srv.filmworldservice.DirectorProto")
	proto.RegisterType((*MovieTypeProto)(nil), "go.micro.srv.filmworldservice.MovieTypeProto")
	proto.RegisterType((*MovieLanguagesProto)(nil), "go.micro.srv.filmworldservice.MovieLanguagesProto")
	proto.RegisterType((*MovieGenresProto)(nil), "go.micro.srv.filmworldservice.MovieGenresProto")
}

func init() {
	proto.RegisterFile("proto/filmworld/filmworldpro.proto", fileDescriptor_filmworldpro_ac66f8145b208ca7)
}

var fileDescriptor_filmworldpro_ac66f8145b208ca7 = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdf, 0x6a, 0x13, 0x4f,
	0x14, 0xee, 0x6e, 0xb3, 0x69, 0x73, 0xda, 0x94, 0x76, 0xfa, 0xe3, 0xc7, 0x52, 0xb4, 0x84, 0xbd,
	0xaa, 0x58, 0xd3, 0x12, 0xbd, 0xf0, 0x4e, 0xa2, 0x6d, 0xc3, 0xa2, 0x96, 0x32, 0x29, 0x78, 0x1d,
	0xb3, 0xc7, 0xb0, 0xb0, 0xc9, 0xac, 0xb3, 0xdb, 0x48, 0x41, 0x05, 0xaf, 0x7c, 0x0b, 0xc1, 0xe7,
	0xf0, 0x41, 0x7c, 0x1d, 0x99, 0x33, 0x3b, 0xfb, 0x27, 0xad, 0x76, 0x95, 0x22, 0x5e, 0x65, 0xce,
	0xbf, 0xef, 0xec, 0x7c, 0xe7, 0x9b, 0x99, 0x80, 0x17, 0x4b, 0x91, 0x8a, 0x83, 0x37, 0x61, 0x34,
	0x7d, 0x27, 0x64, 0x14, 0x14, 0xab, 0x58, 0x8a, 0x2e, 0x05, 0xd9, 0xdd, 0x89, 0xe8, 0x4e, 0xc3,
	0xb1, 0x14, 0xdd, 0x44, 0xce, 0xbb, 0x79, 0x42, 0x82, 0x72, 0x1e, 0x8e, 0xd1, 0x63, 0xb0, 0xc9,
	0xf1, 0xed, 0x05, 0x26, 0x69, 0x3f, 0x8a, 0x5e, 0x8a, 0x79, 0x88, 0x89, 0xb7, 0x0b, 0xeb, 0x99,
	0x8f, 0x1c, 0x6c, 0x03, 0x6c, 0x3f, 0x70, 0xad, 0x8e, 0xb5, 0xe7, 0x70, 0xdb, 0x0f, 0xbc, 0xf7,
	0xb0, 0xc5, 0x31, 0x89, 0xc5, 0x2c, 0xc1, 0xbc, 0x88, 0x1d, 0x83, 0x43, 0x2b, 0xca, 0x5b, 0xeb,
	0x1d, 0x74, 0x7f, 0xd9, 0xb7, 0x4b, 0xb9, 0xc3, 0x70, 0x1a, 0x47, 0x78, 0xa6, 0xbe, 0x93, 0xeb,
	0x6a, 0xd6, 0x81, 0x35, 0x8e, 0x63, 0x21, 0x83, 0x67, 0xe2, 0x62, 0x96, 0xba, 0x36, 0x35, 0x2d,
	0xbb, 0xbc, 0x33, 0x68, 0x9b, 0xee, 0xba, 0xe4, 0x49, 0xb5, 0xf3, 0xbd, 0x3a, 0x9d, 0xcb, 0x3d,
	0xbd, 0x8f, 0xb0, 0x5d, 0x41, 0x1c, 0xa6, 0x12, 0x47, 0xd3, 0xbf, 0xb7, 0xa3, 0x6f, 0x36, 0x6c,
	0x2e, 0x56, 0x2f, 0x92, 0xce, 0xfe, 0x03, 0xe7, 0x3c, 0x4c, 0x23, 0x24, 0x80, 0x16, 0xd7, 0x86,
	0x06, 0x8f, 0x70, 0x94, 0xe0, 0xd1, 0x28, 0x45, 0x77, 0x99, 0x62, 0x65, 0x17, 0xbb, 0x03, 0x2d,
	0xc2, 0x3e, 0xbf, 0x8c, 0xd1, 0x6d, 0x50, 0xbc, 0x70, 0xb0, 0x1d, 0x58, 0x7d, 0x31, 0x9a, 0x4d,
	0x2e, 0x46, 0x13, 0x74, 0x1d, 0x0a, 0xe6, 0xb6, 0xea, 0x78, 0x26, 0xc3, 0x31, 0xba, 0xcd, 0x8e,
	0xb5, 0x67, 0x73, 0x6d, 0x28, 0xef, 0x00, 0x67, 0x12, 0xdd, 0x15, 0xfd, 0x1d, 0x64, 0x28, 0x9c,
	0xa3, 0x50, 0xe2, 0x38, 0x15, 0xd2, 0x5d, 0xd5, 0x38, 0xc6, 0x66, 0x7d, 0x68, 0xf6, 0xd5, 0x22,
	0x71, 0x5b, 0x9d, 0xe5, 0x1a, 0x03, 0xa2, 0x64, 0x4d, 0x61, 0x56, 0xa8, 0xe0, 0x63, 0x29, 0xe6,
	0x61, 0x80, 0xd2, 0x05, 0x0d, 0x6f, 0x6c, 0xef, 0x7b, 0x03, 0xa0, 0x98, 0xe9, 0xad, 0xf1, 0xf6,
	0x3f, 0x34, 0x15, 0x43, 0x7e, 0x40, 0xa4, 0x39, 0x3c, 0xb3, 0xd8, 0x2e, 0x80, 0x61, 0xc8, 0x0f,
	0x88, 0x33, 0x87, 0x97, 0x3c, 0x3f, 0x61, 0xcd, 0x85, 0x15, 0x22, 0xca, 0x0f, 0x88, 0x37, 0x87,
	0x1b, 0x53, 0xe1, 0x19, 0xa6, 0xfc, 0x80, 0xb8, 0x73, 0x78, 0xc9, 0x73, 0x1b, 0xec, 0x71, 0x68,
	0x13, 0x41, 0xf9, 0x84, 0x80, 0x04, 0xbd, 0x7f, 0x03, 0x92, 0x49, 0xd7, 0x60, 0x55, 0x08, 0xf6,
	0xbc, 0x2c, 0xab, 0x35, 0xc2, 0x7b, 0x50, 0xe7, 0x80, 0xa8, 0x7c, 0x0d, 0x58, 0x52, 0xe1, 0x69,
	0x49, 0x85, 0xeb, 0x84, 0xd5, 0xab, 0x83, 0x65, 0x6a, 0x12, 0x0d, 0x58, 0x28, 0xf7, 0xd8, 0x68,
	0xb4, 0x5d, 0xff, 0xe4, 0x52, 0x41, 0x86, 0xa4, 0xab, 0xbd, 0x2f, 0x16, 0x40, 0x41, 0xe7, 0x15,
	0x65, 0x31, 0x68, 0x9c, 0x8e, 0xa6, 0x46, 0x58, 0xb4, 0x56, 0xa7, 0xed, 0x24, 0x94, 0x49, 0x4a,
	0x01, 0xad, 0xaa, 0xc2, 0xa1, 0x4f, 0x5b, 0x16, 0x6c, 0x98, 0xd3, 0x56, 0xc4, 0x06, 0x38, 0x0b,
	0x50, 0xe6, 0xaa, 0xca, 0x6d, 0xa5, 0x45, 0xbd, 0x26, 0x51, 0xb5, 0x78, 0x66, 0x79, 0x5f, 0x2d,
	0x68, 0x57, 0xa6, 0xf4, 0x0f, 0x7e, 0xe3, 0x23, 0xd8, 0xa8, 0x0e, 0xbe, 0xce, 0x37, 0x7a, 0x7d,
	0xd8, 0xbe, 0x66, 0xc4, 0x57, 0x4a, 0xcb, 0xd7, 0x97, 0x5d, 0xbd, 0xbe, 0xbc, 0xc7, 0xd9, 0xa5,
	0x5a, 0x1a, 0xec, 0x75, 0x97, 0x83, 0x16, 0x8a, 0x5d, 0xba, 0xcc, 0x7a, 0x9f, 0x97, 0x61, 0xf3,
	0x24, 0x8c, 0xa6, 0xaf, 0x94, 0x48, 0x86, 0x5a, 0x24, 0xec, 0x93, 0x05, 0xee, 0x00, 0x8b, 0x57,
	0x72, 0x18, 0xce, 0x26, 0x91, 0x79, 0x2a, 0x6e, 0x52, 0xd8, 0xe2, 0x13, 0xbb, 0x73, 0x78, 0x63,
	0xc1, 0xc2, 0xfb, 0xea, 0x2d, 0x1d, 0x5a, 0x2c, 0x86, 0xad, 0x01, 0xea, 0x47, 0xf9, 0xe9, 0x25,
	0xfd, 0xf8, 0x01, 0xbb, 0x5f, 0xaf, 0x37, 0xa5, 0xef, 0xec, 0xd7, 0xec, 0xab, 0x1f, 0xc6, 0x25,
	0xf6, 0x01, 0x58, 0x65, 0xd3, 0x7f, 0xb8, 0xdd, 0xde, 0xef, 0xb4, 0xd5, 0x4d, 0xd4, 0x86, 0x5f,
	0x37, 0xe9, 0x3f, 0xcc, 0xc3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xf2, 0x88, 0x64, 0xe9,
	0x08, 0x00, 0x00,
}
