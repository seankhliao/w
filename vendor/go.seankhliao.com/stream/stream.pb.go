// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: stream.proto

package stream

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BeaconRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DurationMs int64  `protobuf:"varint,1,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	SrcPage    string `protobuf:"bytes,2,opt,name=src_page,json=srcPage,proto3" json:"src_page,omitempty"`
	DstPage    string `protobuf:"bytes,3,opt,name=dst_page,json=dstPage,proto3" json:"dst_page,omitempty"`
	Remote     string `protobuf:"bytes,4,opt,name=remote,proto3" json:"remote,omitempty"`
	UserAgent  string `protobuf:"bytes,5,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	Referrer   string `protobuf:"bytes,6,opt,name=referrer,proto3" json:"referrer,omitempty"`
}

func (x *BeaconRequest) Reset() {
	*x = BeaconRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeaconRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeaconRequest) ProtoMessage() {}

func (x *BeaconRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeaconRequest.ProtoReflect.Descriptor instead.
func (*BeaconRequest) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

func (x *BeaconRequest) GetDurationMs() int64 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *BeaconRequest) GetSrcPage() string {
	if x != nil {
		return x.SrcPage
	}
	return ""
}

func (x *BeaconRequest) GetDstPage() string {
	if x != nil {
		return x.DstPage
	}
	return ""
}

func (x *BeaconRequest) GetRemote() string {
	if x != nil {
		return x.Remote
	}
	return ""
}

func (x *BeaconRequest) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *BeaconRequest) GetReferrer() string {
	if x != nil {
		return x.Referrer
	}
	return ""
}

type CSPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp          string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Remote             string `protobuf:"bytes,2,opt,name=remote,proto3" json:"remote,omitempty"`
	UserAgent          string `protobuf:"bytes,3,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	Referrer           string `protobuf:"bytes,4,opt,name=referrer,proto3" json:"referrer,omitempty"`
	Enforce            string `protobuf:"bytes,5,opt,name=enforce,proto3" json:"enforce,omitempty"`
	BlockedUri         string `protobuf:"bytes,6,opt,name=blocked_uri,json=blockedUri,proto3" json:"blocked_uri,omitempty"`
	SourceFile         string `protobuf:"bytes,7,opt,name=source_file,json=sourceFile,proto3" json:"source_file,omitempty"`
	DocumentUri        string `protobuf:"bytes,8,opt,name=document_uri,json=documentUri,proto3" json:"document_uri,omitempty"`
	ViolatedDirective  string `protobuf:"bytes,9,opt,name=violated_directive,json=violatedDirective,proto3" json:"violated_directive,omitempty"`
	EffectiveDirective string `protobuf:"bytes,10,opt,name=effective_directive,json=effectiveDirective,proto3" json:"effective_directive,omitempty"`
	LineNumber         int64  `protobuf:"varint,11,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	StatusCode         int64  `protobuf:"varint,12,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
}

func (x *CSPRequest) Reset() {
	*x = CSPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSPRequest) ProtoMessage() {}

func (x *CSPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSPRequest.ProtoReflect.Descriptor instead.
func (*CSPRequest) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{1}
}

func (x *CSPRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *CSPRequest) GetRemote() string {
	if x != nil {
		return x.Remote
	}
	return ""
}

func (x *CSPRequest) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *CSPRequest) GetReferrer() string {
	if x != nil {
		return x.Referrer
	}
	return ""
}

func (x *CSPRequest) GetEnforce() string {
	if x != nil {
		return x.Enforce
	}
	return ""
}

func (x *CSPRequest) GetBlockedUri() string {
	if x != nil {
		return x.BlockedUri
	}
	return ""
}

func (x *CSPRequest) GetSourceFile() string {
	if x != nil {
		return x.SourceFile
	}
	return ""
}

func (x *CSPRequest) GetDocumentUri() string {
	if x != nil {
		return x.DocumentUri
	}
	return ""
}

func (x *CSPRequest) GetViolatedDirective() string {
	if x != nil {
		return x.ViolatedDirective
	}
	return ""
}

func (x *CSPRequest) GetEffectiveDirective() string {
	if x != nil {
		return x.EffectiveDirective
	}
	return ""
}

func (x *CSPRequest) GetLineNumber() int64 {
	if x != nil {
		return x.LineNumber
	}
	return 0
}

func (x *CSPRequest) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

type HTTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Method    string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Domain    string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Path      string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Remote    string `protobuf:"bytes,5,opt,name=remote,proto3" json:"remote,omitempty"`
	UserAgent string `protobuf:"bytes,6,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	Referrer  string `protobuf:"bytes,7,opt,name=referrer,proto3" json:"referrer,omitempty"`
}

func (x *HTTPRequest) Reset() {
	*x = HTTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRequest) ProtoMessage() {}

func (x *HTTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRequest.ProtoReflect.Descriptor instead.
func (*HTTPRequest) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{2}
}

func (x *HTTPRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *HTTPRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HTTPRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *HTTPRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *HTTPRequest) GetRemote() string {
	if x != nil {
		return x.Remote
	}
	return ""
}

func (x *HTTPRequest) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *HTTPRequest) GetReferrer() string {
	if x != nil {
		return x.Referrer
	}
	return ""
}

type RepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Owner     string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Repo      string `protobuf:"bytes,3,opt,name=repo,proto3" json:"repo,omitempty"`
}

func (x *RepoRequest) Reset() {
	*x = RepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoRequest) ProtoMessage() {}

func (x *RepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoRequest.ProtoReflect.Descriptor instead.
func (*RepoRequest) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{3}
}

func (x *RepoRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *RepoRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *RepoRequest) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{4}
}

var File_stream_proto protoreflect.FileDescriptor

var file_stream_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0xb9, 0x01, 0x0a, 0x0d, 0x42, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x22, 0x9e, 0x03, 0x0a, 0x0a, 0x43, 0x53, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x55, 0x72, 0x69, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72,
	0x69, 0x12, 0x2d, 0x0a, 0x12, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76,
	0x69, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x2f, 0x0a, 0x13, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x0b, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x22,
	0x08, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xca, 0x01, 0x0a, 0x06, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x2e, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x48, 0x54, 0x54, 0x50, 0x12,
	0x13, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x42, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x12, 0x15, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x43,
	0x53, 0x50, 0x12, 0x12, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x53, 0x50, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70,
	0x6f, 0x12, 0x13, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x6f, 0x2e, 0x73, 0x65, 0x61,
	0x6e, 0x6b, 0x68, 0x6c, 0x69, 0x61, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_proto_rawDescOnce sync.Once
	file_stream_proto_rawDescData = file_stream_proto_rawDesc
)

func file_stream_proto_rawDescGZIP() []byte {
	file_stream_proto_rawDescOnce.Do(func() {
		file_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_proto_rawDescData)
	})
	return file_stream_proto_rawDescData
}

var file_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_stream_proto_goTypes = []interface{}{
	(*BeaconRequest)(nil), // 0: stream.BeaconRequest
	(*CSPRequest)(nil),    // 1: stream.CSPRequest
	(*HTTPRequest)(nil),   // 2: stream.HTTPRequest
	(*RepoRequest)(nil),   // 3: stream.RepoRequest
	(*Result)(nil),        // 4: stream.Result
}
var file_stream_proto_depIdxs = []int32{
	2, // 0: stream.Stream.LogHTTP:input_type -> stream.HTTPRequest
	0, // 1: stream.Stream.LogBeacon:input_type -> stream.BeaconRequest
	1, // 2: stream.Stream.LogCSP:input_type -> stream.CSPRequest
	3, // 3: stream.Stream.LogRepo:input_type -> stream.RepoRequest
	4, // 4: stream.Stream.LogHTTP:output_type -> stream.Result
	4, // 5: stream.Stream.LogBeacon:output_type -> stream.Result
	4, // 6: stream.Stream.LogCSP:output_type -> stream.Result
	4, // 7: stream.Stream.LogRepo:output_type -> stream.Result
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stream_proto_init() }
func file_stream_proto_init() {
	if File_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeaconRequest); i {
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
		file_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSPRequest); i {
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
		file_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPRequest); i {
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
		file_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoRequest); i {
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
		file_stream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_proto_goTypes,
		DependencyIndexes: file_stream_proto_depIdxs,
		MessageInfos:      file_stream_proto_msgTypes,
	}.Build()
	File_stream_proto = out.File
	file_stream_proto_rawDesc = nil
	file_stream_proto_goTypes = nil
	file_stream_proto_depIdxs = nil
}
