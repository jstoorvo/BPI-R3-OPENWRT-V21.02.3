// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/dataqna/v1alpha/annotated_string.proto

package dataqna

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Semantic markup types.
type AnnotatedString_SemanticMarkupType int32

const (
	// No markup type was specified.
	AnnotatedString_MARKUP_TYPE_UNSPECIFIED AnnotatedString_SemanticMarkupType = 0
	// Markup for a substring denoting a metric.
	AnnotatedString_METRIC AnnotatedString_SemanticMarkupType = 1
	// Markup for a substring denoting a dimension.
	AnnotatedString_DIMENSION AnnotatedString_SemanticMarkupType = 2
	// Markup for a substring denoting a filter.
	AnnotatedString_FILTER AnnotatedString_SemanticMarkupType = 3
	// Markup for an unused substring.
	AnnotatedString_UNUSED AnnotatedString_SemanticMarkupType = 4
	// Markup for a substring containing blocked phrases.
	AnnotatedString_BLOCKED AnnotatedString_SemanticMarkupType = 5
	// Markup for a substring that contains terms for row.
	AnnotatedString_ROW AnnotatedString_SemanticMarkupType = 6
)

// Enum value maps for AnnotatedString_SemanticMarkupType.
var (
	AnnotatedString_SemanticMarkupType_name = map[int32]string{
		0: "MARKUP_TYPE_UNSPECIFIED",
		1: "METRIC",
		2: "DIMENSION",
		3: "FILTER",
		4: "UNUSED",
		5: "BLOCKED",
		6: "ROW",
	}
	AnnotatedString_SemanticMarkupType_value = map[string]int32{
		"MARKUP_TYPE_UNSPECIFIED": 0,
		"METRIC":                  1,
		"DIMENSION":               2,
		"FILTER":                  3,
		"UNUSED":                  4,
		"BLOCKED":                 5,
		"ROW":                     6,
	}
)

func (x AnnotatedString_SemanticMarkupType) Enum() *AnnotatedString_SemanticMarkupType {
	p := new(AnnotatedString_SemanticMarkupType)
	*p = x
	return p
}

func (x AnnotatedString_SemanticMarkupType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnnotatedString_SemanticMarkupType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dataqna_v1alpha_annotated_string_proto_enumTypes[0].Descriptor()
}

func (AnnotatedString_SemanticMarkupType) Type() protoreflect.EnumType {
	return &file_google_cloud_dataqna_v1alpha_annotated_string_proto_enumTypes[0]
}

func (x AnnotatedString_SemanticMarkupType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnnotatedString_SemanticMarkupType.Descriptor instead.
func (AnnotatedString_SemanticMarkupType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDescGZIP(), []int{0, 0}
}

// Describes string annotation from both semantic and formatting perspectives.
// Example:
//
// User Query:
//
//   top countries by population in Africa
//
//   0   4         14 17         28 31    37
//
// Table Data:
//
// + "country" - dimension
// + "population" - metric
// + "Africa" - value in the "continent" column
//
// text_formatted = `"top countries by population in Africa"`
//
// html_formatted =
//   `"top <b>countries</b> by <b>population</b> in <i>Africa</i>"`
//
// ```
// markups = [
//   {DIMENSION, 4, 12}, // 'countries'
//   {METRIC, 17, 26}, // 'population'
//   {FILTER, 31, 36}  // 'Africa'
// ]
// ```
//
// Note that html formattings for 'DIMENSION' and 'METRIC' are the same, while
// semantic markups are different.
type AnnotatedString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Text version of the string.
	TextFormatted string `protobuf:"bytes,1,opt,name=text_formatted,json=textFormatted,proto3" json:"text_formatted,omitempty"`
	// HTML version of the string annotation.
	HtmlFormatted string `protobuf:"bytes,2,opt,name=html_formatted,json=htmlFormatted,proto3" json:"html_formatted,omitempty"`
	// Semantic version of the string annotation.
	Markups []*AnnotatedString_SemanticMarkup `protobuf:"bytes,3,rep,name=markups,proto3" json:"markups,omitempty"`
}

func (x *AnnotatedString) Reset() {
	*x = AnnotatedString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataqna_v1alpha_annotated_string_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotatedString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotatedString) ProtoMessage() {}

func (x *AnnotatedString) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataqna_v1alpha_annotated_string_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotatedString.ProtoReflect.Descriptor instead.
func (*AnnotatedString) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDescGZIP(), []int{0}
}

func (x *AnnotatedString) GetTextFormatted() string {
	if x != nil {
		return x.TextFormatted
	}
	return ""
}

func (x *AnnotatedString) GetHtmlFormatted() string {
	if x != nil {
		return x.HtmlFormatted
	}
	return ""
}

func (x *AnnotatedString) GetMarkups() []*AnnotatedString_SemanticMarkup {
	if x != nil {
		return x.Markups
	}
	return nil
}

// Semantic markup denotes a substring (by index and length) with markup
// information.
type AnnotatedString_SemanticMarkup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The semantic type of the markup substring.
	Type AnnotatedString_SemanticMarkupType `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.dataqna.v1alpha.AnnotatedString_SemanticMarkupType" json:"type,omitempty"`
	// Unicode character index of the query.
	StartCharIndex int32 `protobuf:"varint,2,opt,name=start_char_index,json=startCharIndex,proto3" json:"start_char_index,omitempty"`
	// The length (number of unicode characters) of the markup substring.
	Length int32 `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *AnnotatedString_SemanticMarkup) Reset() {
	*x = AnnotatedString_SemanticMarkup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataqna_v1alpha_annotated_string_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotatedString_SemanticMarkup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotatedString_SemanticMarkup) ProtoMessage() {}

func (x *AnnotatedString_SemanticMarkup) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataqna_v1alpha_annotated_string_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotatedString_SemanticMarkup.ProtoReflect.Descriptor instead.
func (*AnnotatedString_SemanticMarkup) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AnnotatedString_SemanticMarkup) GetType() AnnotatedString_SemanticMarkupType {
	if x != nil {
		return x.Type
	}
	return AnnotatedString_MARKUP_TYPE_UNSPECIFIED
}

func (x *AnnotatedString_SemanticMarkup) GetStartCharIndex() int32 {
	if x != nil {
		return x.StartCharIndex
	}
	return 0
}

func (x *AnnotatedString_SemanticMarkup) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

var File_google_cloud_dataqna_v1alpha_annotated_string_proto protoreflect.FileDescriptor

var file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x22, 0xde, 0x03, 0x0a, 0x0f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x74, 0x65, 0x78, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x74, 0x6d, 0x6c, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x74, 0x65, 0x64, 0x12, 0x56, 0x0a, 0x07, 0x6d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x4d, 0x61,
	0x72, 0x6b, 0x75, 0x70, 0x52, 0x07, 0x6d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x73, 0x1a, 0xa8, 0x01,
	0x0a, 0x0e, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x4d, 0x61, 0x72, 0x6b, 0x75, 0x70,
	0x12, 0x54, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x71, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65,
	0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x4d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x63, 0x68, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x7a, 0x0a, 0x12, 0x53, 0x65, 0x6d, 0x61,
	0x6e, 0x74, 0x69, 0x63, 0x4d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x4d, 0x41, 0x52, 0x4b, 0x55, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x45, 0x54, 0x52, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x49, 0x4d, 0x45, 0x4e,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52,
	0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x55, 0x53, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x52,
	0x4f, 0x57, 0x10, 0x06, 0x42, 0xdf, 0x01, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e,
	0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x14, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x64,
	0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0xaa, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x51, 0x6e, 0x41, 0x2e, 0x56, 0x31,
	0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x51, 0x6e, 0x41, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0xea, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x51, 0x6e, 0x41, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDescOnce sync.Once
	file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDescData = file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDesc
)

func file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDescGZIP() []byte {
	file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDescOnce.Do(func() {
		file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDescData)
	})
	return file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDescData
}

var file_google_cloud_dataqna_v1alpha_annotated_string_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_dataqna_v1alpha_annotated_string_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_dataqna_v1alpha_annotated_string_proto_goTypes = []interface{}{
	(AnnotatedString_SemanticMarkupType)(0), // 0: google.cloud.dataqna.v1alpha.AnnotatedString.SemanticMarkupType
	(*AnnotatedString)(nil),                 // 1: google.cloud.dataqna.v1alpha.AnnotatedString
	(*AnnotatedString_SemanticMarkup)(nil),  // 2: google.cloud.dataqna.v1alpha.AnnotatedString.SemanticMarkup
}
var file_google_cloud_dataqna_v1alpha_annotated_string_proto_depIdxs = []int32{
	2, // 0: google.cloud.dataqna.v1alpha.AnnotatedString.markups:type_name -> google.cloud.dataqna.v1alpha.AnnotatedString.SemanticMarkup
	0, // 1: google.cloud.dataqna.v1alpha.AnnotatedString.SemanticMarkup.type:type_name -> google.cloud.dataqna.v1alpha.AnnotatedString.SemanticMarkupType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_dataqna_v1alpha_annotated_string_proto_init() }
func file_google_cloud_dataqna_v1alpha_annotated_string_proto_init() {
	if File_google_cloud_dataqna_v1alpha_annotated_string_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dataqna_v1alpha_annotated_string_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotatedString); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_dataqna_v1alpha_annotated_string_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotatedString_SemanticMarkup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dataqna_v1alpha_annotated_string_proto_goTypes,
		DependencyIndexes: file_google_cloud_dataqna_v1alpha_annotated_string_proto_depIdxs,
		EnumInfos:         file_google_cloud_dataqna_v1alpha_annotated_string_proto_enumTypes,
		MessageInfos:      file_google_cloud_dataqna_v1alpha_annotated_string_proto_msgTypes,
	}.Build()
	File_google_cloud_dataqna_v1alpha_annotated_string_proto = out.File
	file_google_cloud_dataqna_v1alpha_annotated_string_proto_rawDesc = nil
	file_google_cloud_dataqna_v1alpha_annotated_string_proto_goTypes = nil
	file_google_cloud_dataqna_v1alpha_annotated_string_proto_depIdxs = nil
}
