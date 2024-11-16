// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: audio.proto

package audiov1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AudioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *AudioRequest) Reset() {
	*x = AudioRequest{}
	mi := &file_audio_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AudioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioRequest) ProtoMessage() {}

func (x *AudioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audio_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioRequest.ProtoReflect.Descriptor instead.
func (*AudioRequest) Descriptor() ([]byte, []int) {
	return file_audio_proto_rawDescGZIP(), []int{0}
}

func (x *AudioRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type AudioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AudioResponse) Reset() {
	*x = AudioResponse{}
	mi := &file_audio_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AudioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioResponse) ProtoMessage() {}

func (x *AudioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audio_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioResponse.ProtoReflect.Descriptor instead.
func (*AudioResponse) Descriptor() ([]byte, []int) {
	return file_audio_proto_rawDescGZIP(), []int{1}
}

func (x *AudioResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_audio_proto protoreflect.FileDescriptor

var file_audio_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x22, 0x27, 0x0a, 0x0d, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0x46, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x36, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x12, 0x13, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x4d, 0x4d, 0x61, 0x78, 0x61,
	0x4d, 0x4d, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x64, 0x69,
	0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_audio_proto_rawDescOnce sync.Once
	file_audio_proto_rawDescData = file_audio_proto_rawDesc
)

func file_audio_proto_rawDescGZIP() []byte {
	file_audio_proto_rawDescOnce.Do(func() {
		file_audio_proto_rawDescData = protoimpl.X.CompressGZIP(file_audio_proto_rawDescData)
	})
	return file_audio_proto_rawDescData
}

var file_audio_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_audio_proto_goTypes = []any{
	(*AudioRequest)(nil),  // 0: audio.AudioRequest
	(*AudioResponse)(nil), // 1: audio.AudioResponse
}
var file_audio_proto_depIdxs = []int32{
	0, // 0: audio.AudioService.Recognize:input_type -> audio.AudioRequest
	1, // 1: audio.AudioService.Recognize:output_type -> audio.AudioResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_audio_proto_init() }
func file_audio_proto_init() {
	if File_audio_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_audio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_audio_proto_goTypes,
		DependencyIndexes: file_audio_proto_depIdxs,
		MessageInfos:      file_audio_proto_msgTypes,
	}.Build()
	File_audio_proto = out.File
	file_audio_proto_rawDesc = nil
	file_audio_proto_goTypes = nil
	file_audio_proto_depIdxs = nil
}