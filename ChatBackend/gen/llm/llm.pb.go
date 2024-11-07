// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: llm.proto

package llmv1

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role    string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_llm_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_llm_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_llm_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type LLMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages  []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	MaxTokens uint32     `protobuf:"varint,2,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
}

func (x *LLMRequest) Reset() {
	*x = LLMRequest{}
	mi := &file_llm_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LLMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LLMRequest) ProtoMessage() {}

func (x *LLMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_llm_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LLMRequest.ProtoReflect.Descriptor instead.
func (*LLMRequest) Descriptor() ([]byte, []int) {
	return file_llm_proto_rawDescGZIP(), []int{1}
}

func (x *LLMRequest) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *LLMRequest) GetMaxTokens() uint32 {
	if x != nil {
		return x.MaxTokens
	}
	return 0
}

type LLMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Error   string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LLMResponse) Reset() {
	*x = LLMResponse{}
	mi := &file_llm_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LLMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LLMResponse) ProtoMessage() {}

func (x *LLMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_llm_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LLMResponse.ProtoReflect.Descriptor instead.
func (*LLMResponse) Descriptor() ([]byte, []int) {
	return file_llm_proto_rawDescGZIP(), []int{2}
}

func (x *LLMResponse) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *LLMResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_llm_proto protoreflect.FileDescriptor

var file_llm_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6c, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6c, 0x6c, 0x6d,
	0x22, 0x37, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x0a, 0x4c, 0x4c, 0x4d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x22, 0x4b, 0x0a, 0x0b, 0x4c, 0x4c, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x3b, 0x0a,
	0x0a, 0x4c, 0x4c, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x4c, 0x4c,
	0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x4c,
	0x4c, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x4d, 0x4d,
	0x61, 0x78, 0x61, 0x4d, 0x4d, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x3b, 0x6c, 0x6c, 0x6d,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_llm_proto_rawDescOnce sync.Once
	file_llm_proto_rawDescData = file_llm_proto_rawDesc
)

func file_llm_proto_rawDescGZIP() []byte {
	file_llm_proto_rawDescOnce.Do(func() {
		file_llm_proto_rawDescData = protoimpl.X.CompressGZIP(file_llm_proto_rawDescData)
	})
	return file_llm_proto_rawDescData
}

var file_llm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_llm_proto_goTypes = []any{
	(*Message)(nil),     // 0: llm.Message
	(*LLMRequest)(nil),  // 1: llm.LLMRequest
	(*LLMResponse)(nil), // 2: llm.LLMResponse
}
var file_llm_proto_depIdxs = []int32{
	0, // 0: llm.LLMRequest.messages:type_name -> llm.Message
	0, // 1: llm.LLMResponse.message:type_name -> llm.Message
	1, // 2: llm.LLMService.Generate:input_type -> llm.LLMRequest
	2, // 3: llm.LLMService.Generate:output_type -> llm.LLMResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_llm_proto_init() }
func file_llm_proto_init() {
	if File_llm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_llm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_llm_proto_goTypes,
		DependencyIndexes: file_llm_proto_depIdxs,
		MessageInfos:      file_llm_proto_msgTypes,
	}.Build()
	File_llm_proto = out.File
	file_llm_proto_rawDesc = nil
	file_llm_proto_goTypes = nil
	file_llm_proto_depIdxs = nil
}