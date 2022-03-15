/// This file contains service for the Node RPC Server

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: v1/response.proto

// Package Highway is used for defining a Highway node and its accessible API Endpoints

package highwayv1

import (
	_ "go.buf.build/grpc/go/sonr-io/sonr/object"
	registry "go.buf.build/grpc/go/sonr-io/sonr/registry"
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

// MsgAccessNameResponse represents a response to a request for a name
type MsgAccessNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Code of the response
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Message of the response
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Data of the response
	Peer *registry.Peer `protobuf:"bytes,3,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (x *MsgAccessNameResponse) Reset() {
	*x = MsgAccessNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAccessNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAccessNameResponse) ProtoMessage() {}

func (x *MsgAccessNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgAccessNameResponse.ProtoReflect.Descriptor instead.
func (*MsgAccessNameResponse) Descriptor() ([]byte, []int) {
	return file_v1_response_proto_rawDescGZIP(), []int{0}
}

func (x *MsgAccessNameResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MsgAccessNameResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MsgAccessNameResponse) GetPeer() *registry.Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

type MsgCheckNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// boolean response to know if a name has been taken
	NameAvailable bool `protobuf:"varint,1,opt,name=nameAvailable,proto3" json:"nameAvailable,omitempty"`
}

func (x *MsgCheckNameResponse) Reset() {
	*x = MsgCheckNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCheckNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCheckNameResponse) ProtoMessage() {}

func (x *MsgCheckNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCheckNameResponse.ProtoReflect.Descriptor instead.
func (*MsgCheckNameResponse) Descriptor() ([]byte, []int) {
	return file_v1_response_proto_rawDescGZIP(), []int{1}
}

func (x *MsgCheckNameResponse) GetNameAvailable() bool {
	if x != nil {
		return x.NameAvailable
	}
	return false
}

type MsgGenerateCredsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// boolean response to know if token was attached to the did successfully
	TokenAttached bool `protobuf:"varint,1,opt,name=tokenAttached,proto3" json:"tokenAttached,omitempty"`
}

func (x *MsgGenerateCredsResponse) Reset() {
	*x = MsgGenerateCredsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgGenerateCredsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgGenerateCredsResponse) ProtoMessage() {}

func (x *MsgGenerateCredsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgGenerateCredsResponse.ProtoReflect.Descriptor instead.
func (*MsgGenerateCredsResponse) Descriptor() ([]byte, []int) {
	return file_v1_response_proto_rawDescGZIP(), []int{2}
}

func (x *MsgGenerateCredsResponse) GetTokenAttached() bool {
	if x != nil {
		return x.TokenAttached
	}
	return false
}

// AccessServiceResponse represents a response to a request for a service
type MsgAccessServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Code of the response
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Message of the response
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Data of the response
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MsgAccessServiceResponse) Reset() {
	*x = MsgAccessServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAccessServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAccessServiceResponse) ProtoMessage() {}

func (x *MsgAccessServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgAccessServiceResponse.ProtoReflect.Descriptor instead.
func (*MsgAccessServiceResponse) Descriptor() ([]byte, []int) {
	return file_v1_response_proto_rawDescGZIP(), []int{3}
}

func (x *MsgAccessServiceResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MsgAccessServiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MsgAccessServiceResponse) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgUploadBlobResponse represents a response to a request to upload a blob
type MsgUploadBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Code of the response
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Message of the response
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// DID of the response
	Did string `protobuf:"bytes,3,opt,name=did,proto3" json:"did,omitempty"`
	// Pinned is true if the blob is pinned to IPFS
	Pinned bool `protobuf:"varint,4,opt,name=pinned,proto3" json:"pinned,omitempty"`
}

func (x *MsgUploadBlobResponse) Reset() {
	*x = MsgUploadBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUploadBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUploadBlobResponse) ProtoMessage() {}

func (x *MsgUploadBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgUploadBlobResponse.ProtoReflect.Descriptor instead.
func (*MsgUploadBlobResponse) Descriptor() ([]byte, []int) {
	return file_v1_response_proto_rawDescGZIP(), []int{4}
}

func (x *MsgUploadBlobResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MsgUploadBlobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MsgUploadBlobResponse) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgUploadBlobResponse) GetPinned() bool {
	if x != nil {
		return x.Pinned
	}
	return false
}

// MsgDownloadBlobResponse represents a response to a request to download a blob
type MsgDownloadBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Code of the response
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Message of the response
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// DID of the response
	Did string `protobuf:"bytes,3,opt,name=did,proto3" json:"did,omitempty"`
	// Path of downloaded blob
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *MsgDownloadBlobResponse) Reset() {
	*x = MsgDownloadBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_response_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDownloadBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDownloadBlobResponse) ProtoMessage() {}

func (x *MsgDownloadBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_response_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDownloadBlobResponse.ProtoReflect.Descriptor instead.
func (*MsgDownloadBlobResponse) Descriptor() ([]byte, []int) {
	return file_v1_response_proto_rawDescGZIP(), []int{5}
}

func (x *MsgDownloadBlobResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MsgDownloadBlobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MsgDownloadBlobResponse) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgDownloadBlobResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// MsgSyncBlobResponse represents a response to a request to sync a blob
type MsgSyncBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Code of the response
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Message of the response
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// DID of the response
	Did string `protobuf:"bytes,3,opt,name=did,proto3" json:"did,omitempty"`
}

func (x *MsgSyncBlobResponse) Reset() {
	*x = MsgSyncBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_response_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSyncBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSyncBlobResponse) ProtoMessage() {}

func (x *MsgSyncBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_response_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSyncBlobResponse.ProtoReflect.Descriptor instead.
func (*MsgSyncBlobResponse) Descriptor() ([]byte, []int) {
	return file_v1_response_proto_rawDescGZIP(), []int{6}
}

func (x *MsgSyncBlobResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MsgSyncBlobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MsgSyncBlobResponse) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

// MsgDeleteBlobResponse represents a response to a request to delete a blob
type MsgDeleteBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Code of the response
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Message of the response
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// DID of the response
	Did string `protobuf:"bytes,3,opt,name=did,proto3" json:"did,omitempty"`
}

func (x *MsgDeleteBlobResponse) Reset() {
	*x = MsgDeleteBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_response_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDeleteBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeleteBlobResponse) ProtoMessage() {}

func (x *MsgDeleteBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_response_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeleteBlobResponse.ProtoReflect.Descriptor instead.
func (*MsgDeleteBlobResponse) Descriptor() ([]byte, []int) {
	return file_v1_response_proto_rawDescGZIP(), []int{7}
}

func (x *MsgDeleteBlobResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MsgDeleteBlobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MsgDeleteBlobResponse) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

// MsgParseDidResponse represents a response to a request to parse a DID
type MsgParseDidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Code of the response
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Message of the response
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// DID of the response
	Did string `protobuf:"bytes,3,opt,name=did,proto3" json:"did,omitempty"`
}

func (x *MsgParseDidResponse) Reset() {
	*x = MsgParseDidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_response_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgParseDidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgParseDidResponse) ProtoMessage() {}

func (x *MsgParseDidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_response_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgParseDidResponse.ProtoReflect.Descriptor instead.
func (*MsgParseDidResponse) Descriptor() ([]byte, []int) {
	return file_v1_response_proto_rawDescGZIP(), []int{8}
}

func (x *MsgParseDidResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MsgParseDidResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MsgParseDidResponse) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

// MsgResolveDidResponse represents a response to a request to resolve a DID
type MsgResolveDidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Code of the response
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Message of the response
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// DID of the response
	DidDocument string `protobuf:"bytes,3,opt,name=did_document,json=didDocument,proto3" json:"did_document,omitempty"` // optional
}

func (x *MsgResolveDidResponse) Reset() {
	*x = MsgResolveDidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_response_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgResolveDidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgResolveDidResponse) ProtoMessage() {}

func (x *MsgResolveDidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_response_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgResolveDidResponse.ProtoReflect.Descriptor instead.
func (*MsgResolveDidResponse) Descriptor() ([]byte, []int) {
	return file_v1_response_proto_rawDescGZIP(), []int{9}
}

func (x *MsgResolveDidResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MsgResolveDidResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MsgResolveDidResponse) GetDidDocument() string {
	if x != nil {
		return x.DidDocument
	}
	return ""
}

var File_v1_response_proto protoreflect.FileDescriptor

var file_v1_response_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x68, 0x69, 0x67, 0x68,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x75, 0x0a, 0x15, 0x4d, 0x73, 0x67, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x14, 0x4d, 0x73, 0x67, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x40, 0x0a, 0x18, 0x4d, 0x73, 0x67, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x22, 0xdc, 0x01, 0x0a, 0x18, 0x4d, 0x73, 0x67, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x55, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x68, 0x69,
	0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6f, 0x0a, 0x15, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x22, 0x6d, 0x0a, 0x17, 0x4d, 0x73, 0x67, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x55, 0x0a, 0x13, 0x4d, 0x73, 0x67, 0x53, 0x79, 0x6e,
	0x63, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x22, 0x57, 0x0a,
	0x15, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x13, 0x4d, 0x73, 0x67, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x44, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x22, 0x68, 0x0a,
	0x15, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44, 0x69, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x64, 0x5f, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x64, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x6f, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x3b, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_response_proto_rawDescOnce sync.Once
	file_v1_response_proto_rawDescData = file_v1_response_proto_rawDesc
)

func file_v1_response_proto_rawDescGZIP() []byte {
	file_v1_response_proto_rawDescOnce.Do(func() {
		file_v1_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_response_proto_rawDescData)
	})
	return file_v1_response_proto_rawDescData
}

var file_v1_response_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_v1_response_proto_goTypes = []interface{}{
	(*MsgAccessNameResponse)(nil),    // 0: sonrio.highway.v1.MsgAccessNameResponse
	(*MsgCheckNameResponse)(nil),     // 1: sonrio.highway.v1.MsgCheckNameResponse
	(*MsgGenerateCredsResponse)(nil), // 2: sonrio.highway.v1.MsgGenerateCredsResponse
	(*MsgAccessServiceResponse)(nil), // 3: sonrio.highway.v1.MsgAccessServiceResponse
	(*MsgUploadBlobResponse)(nil),    // 4: sonrio.highway.v1.MsgUploadBlobResponse
	(*MsgDownloadBlobResponse)(nil),  // 5: sonrio.highway.v1.MsgDownloadBlobResponse
	(*MsgSyncBlobResponse)(nil),      // 6: sonrio.highway.v1.MsgSyncBlobResponse
	(*MsgDeleteBlobResponse)(nil),    // 7: sonrio.highway.v1.MsgDeleteBlobResponse
	(*MsgParseDidResponse)(nil),      // 8: sonrio.highway.v1.MsgParseDidResponse
	(*MsgResolveDidResponse)(nil),    // 9: sonrio.highway.v1.MsgResolveDidResponse
	nil,                              // 10: sonrio.highway.v1.MsgAccessServiceResponse.MetadataEntry
	(*registry.Peer)(nil),            // 11: sonrio.sonr.registry.Peer
}
var file_v1_response_proto_depIdxs = []int32{
	11, // 0: sonrio.highway.v1.MsgAccessNameResponse.peer:type_name -> sonrio.sonr.registry.Peer
	10, // 1: sonrio.highway.v1.MsgAccessServiceResponse.metadata:type_name -> sonrio.highway.v1.MsgAccessServiceResponse.MetadataEntry
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_v1_response_proto_init() }
func file_v1_response_proto_init() {
	if File_v1_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAccessNameResponse); i {
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
		file_v1_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCheckNameResponse); i {
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
		file_v1_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgGenerateCredsResponse); i {
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
		file_v1_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAccessServiceResponse); i {
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
		file_v1_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUploadBlobResponse); i {
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
		file_v1_response_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDownloadBlobResponse); i {
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
		file_v1_response_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSyncBlobResponse); i {
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
		file_v1_response_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDeleteBlobResponse); i {
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
		file_v1_response_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgParseDidResponse); i {
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
		file_v1_response_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgResolveDidResponse); i {
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
			RawDescriptor: file_v1_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_response_proto_goTypes,
		DependencyIndexes: file_v1_response_proto_depIdxs,
		MessageInfos:      file_v1_response_proto_msgTypes,
	}.Build()
	File_v1_response_proto = out.File
	file_v1_response_proto_rawDesc = nil
	file_v1_response_proto_goTypes = nil
	file_v1_response_proto_depIdxs = nil
}