// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: payment.proto

package payment

import (
	context "context"
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

type Status int32

const (
	Status_Unknown Status = 0
	Status_Init    Status = 1
	Status_Success Status = 2
	Status_Failed  Status = 3
	Status_Refund  Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Unknown",
		1: "Init",
		2: "Success",
		3: "Failed",
		4: "Refund",
	}
	Status_value = map[string]int32{
		"Unknown": 0,
		"Init":    1,
		"Success": 2,
		"Failed":  3,
		"Refund":  4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_payment_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_payment_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

type PrepayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount     string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`                             //金额 单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	OutTradeNo string `protobuf:"bytes,2,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"` // 唯一订单号商户订单号，64个字符以内、可包含字母、数字、下划线
	Subject    string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`                           //订单标题
}

func (x *PrepayReq) Reset() {
	*x = PrepayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepayReq) ProtoMessage() {}

func (x *PrepayReq) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepayReq.ProtoReflect.Descriptor instead.
func (*PrepayReq) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PrepayReq) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *PrepayReq) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

func (x *PrepayReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type PrepayResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayUrl string `protobuf:"bytes,1,opt,name=pay_url,json=payUrl,proto3" json:"pay_url,omitempty"` //支付链接
}

func (x *PrepayResp) Reset() {
	*x = PrepayResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepayResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepayResp) ProtoMessage() {}

func (x *PrepayResp) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepayResp.ProtoReflect.Descriptor instead.
func (*PrepayResp) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PrepayResp) GetPayUrl() string {
	if x != nil {
		return x.PayUrl
	}
	return ""
}

type FinishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutTradeNo string `protobuf:"bytes,1,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"` // 唯一订单号商户订单号
	TradeNo    string `protobuf:"bytes,2,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`            // 唯一订单号商户订单号
}

func (x *FinishReq) Reset() {
	*x = FinishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishReq) ProtoMessage() {}

func (x *FinishReq) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishReq.ProtoReflect.Descriptor instead.
func (*FinishReq) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{2}
}

func (x *FinishReq) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

func (x *FinishReq) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

type FinishResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinishResp) Reset() {
	*x = FinishResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishResp) ProtoMessage() {}

func (x *FinishResp) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishResp.ProtoReflect.Descriptor instead.
func (*FinishResp) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{3}
}

type GetByOutTradeNoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutTradeNo string `protobuf:"bytes,1,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"` // 唯一订单号商户订单号
}

func (x *GetByOutTradeNoReq) Reset() {
	*x = GetByOutTradeNoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByOutTradeNoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByOutTradeNoReq) ProtoMessage() {}

func (x *GetByOutTradeNoReq) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByOutTradeNoReq.ProtoReflect.Descriptor instead.
func (*GetByOutTradeNoReq) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{4}
}

func (x *GetByOutTradeNoReq) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

type GetByOutTradeNoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	OutTradeNo  string `protobuf:"bytes,2,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
	TradeNo     string `protobuf:"bytes,3,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`
	Status      Status `protobuf:"varint,4,opt,name=status,proto3,enum=payment.Status" json:"status,omitempty"`
	Amount      string `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GetByOutTradeNoResp) Reset() {
	*x = GetByOutTradeNoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByOutTradeNoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByOutTradeNoResp) ProtoMessage() {}

func (x *GetByOutTradeNoResp) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByOutTradeNoResp.ProtoReflect.Descriptor instead.
func (*GetByOutTradeNoResp) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{5}
}

func (x *GetByOutTradeNoResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetByOutTradeNoResp) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

func (x *GetByOutTradeNoResp) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *GetByOutTradeNoResp) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Unknown
}

func (x *GetByOutTradeNoResp) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_payment_proto protoreflect.FileDescriptor

var file_payment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5f, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x70,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0c, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x25, 0x0a, 0x0a, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x79, 0x55, 0x72, 0x6c,
	0x22, 0x48, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a,
	0x0c, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x36, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x20,
	0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f,
	0x22, 0xb5, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x75,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x44, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x10, 0x04, 0x32, 0xca,
	0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x33, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x12, 0x12, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x52, 0x65, 0x71, 0x1a,
	0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x12, 0x12, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x1b,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4f, 0x75,
	0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4f, 0x75, 0x74, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x69, 0x2d, 0x47, 0x75,
	0x6f, 0x2d, 0x63, 0x72, 0x75, 0x73, 0x68, 0x65, 0x64, 0x2d, 0x68, 0x69, 0x73, 0x2d, 0x74, 0x65,
	0x61, 0x6d, 0x2f, 0x43, 0x75, 0x69, 0x47, 0x75, 0x6f, 0x4d, 0x61, 0x6c, 0x6c, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData = file_payment_proto_rawDesc
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_proto_rawDescData)
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_payment_proto_goTypes = []interface{}{
	(Status)(0),                 // 0: payment.Status
	(*PrepayReq)(nil),           // 1: payment.PrepayReq
	(*PrepayResp)(nil),          // 2: payment.PrepayResp
	(*FinishReq)(nil),           // 3: payment.FinishReq
	(*FinishResp)(nil),          // 4: payment.FinishResp
	(*GetByOutTradeNoReq)(nil),  // 5: payment.GetByOutTradeNoReq
	(*GetByOutTradeNoResp)(nil), // 6: payment.GetByOutTradeNoResp
}
var file_payment_proto_depIdxs = []int32{
	0, // 0: payment.GetByOutTradeNoResp.status:type_name -> payment.Status
	1, // 1: payment.PaymentService.Prepay:input_type -> payment.PrepayReq
	3, // 2: payment.PaymentService.Finish:input_type -> payment.FinishReq
	5, // 3: payment.PaymentService.GetByOutTradeNo:input_type -> payment.GetByOutTradeNoReq
	2, // 4: payment.PaymentService.Prepay:output_type -> payment.PrepayResp
	4, // 5: payment.PaymentService.Finish:output_type -> payment.FinishResp
	6, // 6: payment.PaymentService.GetByOutTradeNo:output_type -> payment.GetByOutTradeNoResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepayReq); i {
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
		file_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepayResp); i {
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
		file_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishReq); i {
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
		file_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishResp); i {
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
		file_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByOutTradeNoReq); i {
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
		file_payment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByOutTradeNoResp); i {
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
			RawDescriptor: file_payment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		EnumInfos:         file_payment_proto_enumTypes,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_rawDesc = nil
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type PaymentService interface {
	Prepay(ctx context.Context, req *PrepayReq) (res *PrepayResp, err error)
	Finish(ctx context.Context, req *FinishReq) (res *FinishResp, err error)
	GetByOutTradeNo(ctx context.Context, req *GetByOutTradeNoReq) (res *GetByOutTradeNoResp, err error)
}
