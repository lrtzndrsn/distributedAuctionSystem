// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: auction.proto

package auction

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

type CallElectionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallElectionMessage) Reset() {
	*x = CallElectionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallElectionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallElectionMessage) ProtoMessage() {}

func (x *CallElectionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallElectionMessage.ProtoReflect.Descriptor instead.
func (*CallElectionMessage) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{0}
}

type CallElectionResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallElectionResponseMessage) Reset() {
	*x = CallElectionResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallElectionResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallElectionResponseMessage) ProtoMessage() {}

func (x *CallElectionResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallElectionResponseMessage.ProtoReflect.Descriptor instead.
func (*CallElectionResponseMessage) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{1}
}

type AssertCoordinatorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port string `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *AssertCoordinatorMessage) Reset() {
	*x = AssertCoordinatorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssertCoordinatorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssertCoordinatorMessage) ProtoMessage() {}

func (x *AssertCoordinatorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssertCoordinatorMessage.ProtoReflect.Descriptor instead.
func (*AssertCoordinatorMessage) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{2}
}

func (x *AssertCoordinatorMessage) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type AssertCoordinatorResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port string `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *AssertCoordinatorResponseMessage) Reset() {
	*x = AssertCoordinatorResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssertCoordinatorResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssertCoordinatorResponseMessage) ProtoMessage() {}

func (x *AssertCoordinatorResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssertCoordinatorResponseMessage.ProtoReflect.Descriptor instead.
func (*AssertCoordinatorResponseMessage) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{3}
}

func (x *AssertCoordinatorResponseMessage) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type SendBidMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueIdentifier int64 `protobuf:"varint,1,opt,name=unique_identifier,json=uniqueIdentifier,proto3" json:"unique_identifier,omitempty"`
	Bid              int64 `protobuf:"varint,2,opt,name=bid,proto3" json:"bid,omitempty"`
	FromCoordinator  bool  `protobuf:"varint,3,opt,name=from_coordinator,json=fromCoordinator,proto3" json:"from_coordinator,omitempty"`
	EndTime          int64 `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	IsFirstBid       bool  `protobuf:"varint,5,opt,name=is_first_bid,json=isFirstBid,proto3" json:"is_first_bid,omitempty"`
}

func (x *SendBidMessage) Reset() {
	*x = SendBidMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendBidMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBidMessage) ProtoMessage() {}

func (x *SendBidMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBidMessage.ProtoReflect.Descriptor instead.
func (*SendBidMessage) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{4}
}

func (x *SendBidMessage) GetUniqueIdentifier() int64 {
	if x != nil {
		return x.UniqueIdentifier
	}
	return 0
}

func (x *SendBidMessage) GetBid() int64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *SendBidMessage) GetFromCoordinator() bool {
	if x != nil {
		return x.FromCoordinator
	}
	return false
}

func (x *SendBidMessage) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *SendBidMessage) GetIsFirstBid() bool {
	if x != nil {
		return x.IsFirstBid
	}
	return false
}

type ResponseBidMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseBidMessage) Reset() {
	*x = ResponseBidMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBidMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBidMessage) ProtoMessage() {}

func (x *ResponseBidMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBidMessage.ProtoReflect.Descriptor instead.
func (*ResponseBidMessage) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseBidMessage) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type RequestResultMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestResultMessage) Reset() {
	*x = RequestResultMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestResultMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestResultMessage) ProtoMessage() {}

func (x *RequestResultMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestResultMessage.ProtoReflect.Descriptor instead.
func (*RequestResultMessage) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{6}
}

type ResultResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result               int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	CurrentHighestBidder int64 `protobuf:"varint,2,opt,name=current_highest_bidder,json=currentHighestBidder,proto3" json:"current_highest_bidder,omitempty"`
}

func (x *ResultResponseMessage) Reset() {
	*x = ResultResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponseMessage) ProtoMessage() {}

func (x *ResultResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponseMessage.ProtoReflect.Descriptor instead.
func (*ResultResponseMessage) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{7}
}

func (x *ResultResponseMessage) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *ResultResponseMessage) GetCurrentHighestBidder() int64 {
	if x != nil {
		return x.CurrentHighestBidder
	}
	return 0
}

var File_auction_proto protoreflect.FileDescriptor

var file_auction_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x61, 0x6c, 0x6c,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x1d, 0x0a, 0x1b, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e,
	0x0a, 0x18, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x36,
	0x0a, 0x20, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xb7, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x42,
	0x69, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x42, 0x69, 0x64,
	0x22, 0x2c, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x69, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x16,
	0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x65, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x32, 0xc6, 0x02,
	0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x03, 0x42, 0x69, 0x64,
	0x12, 0x17, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42,
	0x69, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x69, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x1d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x1e, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x52, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x24, 0x2e,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x61, 0x0a, 0x11, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x21, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x29, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auction_proto_rawDescOnce sync.Once
	file_auction_proto_rawDescData = file_auction_proto_rawDesc
)

func file_auction_proto_rawDescGZIP() []byte {
	file_auction_proto_rawDescOnce.Do(func() {
		file_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_auction_proto_rawDescData)
	})
	return file_auction_proto_rawDescData
}

var file_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_auction_proto_goTypes = []interface{}{
	(*CallElectionMessage)(nil),              // 0: auction.CallElectionMessage
	(*CallElectionResponseMessage)(nil),      // 1: auction.CallElectionResponseMessage
	(*AssertCoordinatorMessage)(nil),         // 2: auction.AssertCoordinatorMessage
	(*AssertCoordinatorResponseMessage)(nil), // 3: auction.AssertCoordinatorResponseMessage
	(*SendBidMessage)(nil),                   // 4: auction.SendBidMessage
	(*ResponseBidMessage)(nil),               // 5: auction.ResponseBidMessage
	(*RequestResultMessage)(nil),             // 6: auction.RequestResultMessage
	(*ResultResponseMessage)(nil),            // 7: auction.ResultResponseMessage
}
var file_auction_proto_depIdxs = []int32{
	4, // 0: auction.Auction.Bid:input_type -> auction.SendBidMessage
	6, // 1: auction.Auction.Result:input_type -> auction.RequestResultMessage
	0, // 2: auction.Auction.CallElection:input_type -> auction.CallElectionMessage
	2, // 3: auction.Auction.AssertCoordinator:input_type -> auction.AssertCoordinatorMessage
	5, // 4: auction.Auction.Bid:output_type -> auction.ResponseBidMessage
	7, // 5: auction.Auction.Result:output_type -> auction.ResultResponseMessage
	1, // 6: auction.Auction.CallElection:output_type -> auction.CallElectionResponseMessage
	3, // 7: auction.Auction.AssertCoordinator:output_type -> auction.AssertCoordinatorResponseMessage
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auction_proto_init() }
func file_auction_proto_init() {
	if File_auction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallElectionMessage); i {
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
		file_auction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallElectionResponseMessage); i {
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
		file_auction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssertCoordinatorMessage); i {
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
		file_auction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssertCoordinatorResponseMessage); i {
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
		file_auction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendBidMessage); i {
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
		file_auction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBidMessage); i {
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
		file_auction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestResultMessage); i {
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
		file_auction_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultResponseMessage); i {
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
			RawDescriptor: file_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auction_proto_goTypes,
		DependencyIndexes: file_auction_proto_depIdxs,
		MessageInfos:      file_auction_proto_msgTypes,
	}.Build()
	File_auction_proto = out.File
	file_auction_proto_rawDesc = nil
	file_auction_proto_goTypes = nil
	file_auction_proto_depIdxs = nil
}