// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orders.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OrderOpen struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderOpen) Reset()         { *m = OrderOpen{} }
func (m *OrderOpen) String() string { return proto.CompactTextString(m) }
func (*OrderOpen) ProtoMessage()    {}
func (*OrderOpen) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{0}
}

func (m *OrderOpen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderOpen.Unmarshal(m, b)
}
func (m *OrderOpen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderOpen.Marshal(b, m, deterministic)
}
func (m *OrderOpen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderOpen.Merge(m, src)
}
func (m *OrderOpen) XXX_Size() int {
	return xxx_messageInfo_OrderOpen.Size(m)
}
func (m *OrderOpen) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderOpen.DiscardUnknown(m)
}

var xxx_messageInfo_OrderOpen proto.InternalMessageInfo

type OrderReject struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderReject) Reset()         { *m = OrderReject{} }
func (m *OrderReject) String() string { return proto.CompactTextString(m) }
func (*OrderReject) ProtoMessage()    {}
func (*OrderReject) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{1}
}

func (m *OrderReject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderReject.Unmarshal(m, b)
}
func (m *OrderReject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderReject.Marshal(b, m, deterministic)
}
func (m *OrderReject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderReject.Merge(m, src)
}
func (m *OrderReject) XXX_Size() int {
	return xxx_messageInfo_OrderReject.Size(m)
}
func (m *OrderReject) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderReject.DiscardUnknown(m)
}

var xxx_messageInfo_OrderReject proto.InternalMessageInfo

type OrderCancel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderCancel) Reset()         { *m = OrderCancel{} }
func (m *OrderCancel) String() string { return proto.CompactTextString(m) }
func (*OrderCancel) ProtoMessage()    {}
func (*OrderCancel) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{2}
}

func (m *OrderCancel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderCancel.Unmarshal(m, b)
}
func (m *OrderCancel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderCancel.Marshal(b, m, deterministic)
}
func (m *OrderCancel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderCancel.Merge(m, src)
}
func (m *OrderCancel) XXX_Size() int {
	return xxx_messageInfo_OrderCancel.Size(m)
}
func (m *OrderCancel) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderCancel.DiscardUnknown(m)
}

var xxx_messageInfo_OrderCancel proto.InternalMessageInfo

type OrderConfirmation struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderConfirmation) Reset()         { *m = OrderConfirmation{} }
func (m *OrderConfirmation) String() string { return proto.CompactTextString(m) }
func (*OrderConfirmation) ProtoMessage()    {}
func (*OrderConfirmation) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{3}
}

func (m *OrderConfirmation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderConfirmation.Unmarshal(m, b)
}
func (m *OrderConfirmation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderConfirmation.Marshal(b, m, deterministic)
}
func (m *OrderConfirmation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderConfirmation.Merge(m, src)
}
func (m *OrderConfirmation) XXX_Size() int {
	return xxx_messageInfo_OrderConfirmation.Size(m)
}
func (m *OrderConfirmation) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderConfirmation.DiscardUnknown(m)
}

var xxx_messageInfo_OrderConfirmation proto.InternalMessageInfo

type OrderFulfillment struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderFulfillment) Reset()         { *m = OrderFulfillment{} }
func (m *OrderFulfillment) String() string { return proto.CompactTextString(m) }
func (*OrderFulfillment) ProtoMessage()    {}
func (*OrderFulfillment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{4}
}

func (m *OrderFulfillment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderFulfillment.Unmarshal(m, b)
}
func (m *OrderFulfillment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderFulfillment.Marshal(b, m, deterministic)
}
func (m *OrderFulfillment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderFulfillment.Merge(m, src)
}
func (m *OrderFulfillment) XXX_Size() int {
	return xxx_messageInfo_OrderFulfillment.Size(m)
}
func (m *OrderFulfillment) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderFulfillment.DiscardUnknown(m)
}

var xxx_messageInfo_OrderFulfillment proto.InternalMessageInfo

type OrderComplete struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderComplete) Reset()         { *m = OrderComplete{} }
func (m *OrderComplete) String() string { return proto.CompactTextString(m) }
func (*OrderComplete) ProtoMessage()    {}
func (*OrderComplete) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{5}
}

func (m *OrderComplete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderComplete.Unmarshal(m, b)
}
func (m *OrderComplete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderComplete.Marshal(b, m, deterministic)
}
func (m *OrderComplete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderComplete.Merge(m, src)
}
func (m *OrderComplete) XXX_Size() int {
	return xxx_messageInfo_OrderComplete.Size(m)
}
func (m *OrderComplete) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderComplete.DiscardUnknown(m)
}

var xxx_messageInfo_OrderComplete proto.InternalMessageInfo

type DisputeOpen struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisputeOpen) Reset()         { *m = DisputeOpen{} }
func (m *DisputeOpen) String() string { return proto.CompactTextString(m) }
func (*DisputeOpen) ProtoMessage()    {}
func (*DisputeOpen) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{6}
}

func (m *DisputeOpen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisputeOpen.Unmarshal(m, b)
}
func (m *DisputeOpen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisputeOpen.Marshal(b, m, deterministic)
}
func (m *DisputeOpen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisputeOpen.Merge(m, src)
}
func (m *DisputeOpen) XXX_Size() int {
	return xxx_messageInfo_DisputeOpen.Size(m)
}
func (m *DisputeOpen) XXX_DiscardUnknown() {
	xxx_messageInfo_DisputeOpen.DiscardUnknown(m)
}

var xxx_messageInfo_DisputeOpen proto.InternalMessageInfo

type DisputeUpdate struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisputeUpdate) Reset()         { *m = DisputeUpdate{} }
func (m *DisputeUpdate) String() string { return proto.CompactTextString(m) }
func (*DisputeUpdate) ProtoMessage()    {}
func (*DisputeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{7}
}

func (m *DisputeUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisputeUpdate.Unmarshal(m, b)
}
func (m *DisputeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisputeUpdate.Marshal(b, m, deterministic)
}
func (m *DisputeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisputeUpdate.Merge(m, src)
}
func (m *DisputeUpdate) XXX_Size() int {
	return xxx_messageInfo_DisputeUpdate.Size(m)
}
func (m *DisputeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_DisputeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_DisputeUpdate proto.InternalMessageInfo

type DisputeClose struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisputeClose) Reset()         { *m = DisputeClose{} }
func (m *DisputeClose) String() string { return proto.CompactTextString(m) }
func (*DisputeClose) ProtoMessage()    {}
func (*DisputeClose) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{8}
}

func (m *DisputeClose) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisputeClose.Unmarshal(m, b)
}
func (m *DisputeClose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisputeClose.Marshal(b, m, deterministic)
}
func (m *DisputeClose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisputeClose.Merge(m, src)
}
func (m *DisputeClose) XXX_Size() int {
	return xxx_messageInfo_DisputeClose.Size(m)
}
func (m *DisputeClose) XXX_DiscardUnknown() {
	xxx_messageInfo_DisputeClose.DiscardUnknown(m)
}

var xxx_messageInfo_DisputeClose proto.InternalMessageInfo

type Refund struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Refund) Reset()         { *m = Refund{} }
func (m *Refund) String() string { return proto.CompactTextString(m) }
func (*Refund) ProtoMessage()    {}
func (*Refund) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{9}
}

func (m *Refund) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Refund.Unmarshal(m, b)
}
func (m *Refund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Refund.Marshal(b, m, deterministic)
}
func (m *Refund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Refund.Merge(m, src)
}
func (m *Refund) XXX_Size() int {
	return xxx_messageInfo_Refund.Size(m)
}
func (m *Refund) XXX_DiscardUnknown() {
	xxx_messageInfo_Refund.DiscardUnknown(m)
}

var xxx_messageInfo_Refund proto.InternalMessageInfo

type PaymentSent struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentSent) Reset()         { *m = PaymentSent{} }
func (m *PaymentSent) String() string { return proto.CompactTextString(m) }
func (*PaymentSent) ProtoMessage()    {}
func (*PaymentSent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{10}
}

func (m *PaymentSent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentSent.Unmarshal(m, b)
}
func (m *PaymentSent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentSent.Marshal(b, m, deterministic)
}
func (m *PaymentSent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentSent.Merge(m, src)
}
func (m *PaymentSent) XXX_Size() int {
	return xxx_messageInfo_PaymentSent.Size(m)
}
func (m *PaymentSent) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentSent.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentSent proto.InternalMessageInfo

type PaymentFinalized struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentFinalized) Reset()         { *m = PaymentFinalized{} }
func (m *PaymentFinalized) String() string { return proto.CompactTextString(m) }
func (*PaymentFinalized) ProtoMessage()    {}
func (*PaymentFinalized) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{11}
}

func (m *PaymentFinalized) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentFinalized.Unmarshal(m, b)
}
func (m *PaymentFinalized) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentFinalized.Marshal(b, m, deterministic)
}
func (m *PaymentFinalized) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentFinalized.Merge(m, src)
}
func (m *PaymentFinalized) XXX_Size() int {
	return xxx_messageInfo_PaymentFinalized.Size(m)
}
func (m *PaymentFinalized) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentFinalized.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentFinalized proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OrderOpen)(nil), "OrderOpen")
	proto.RegisterType((*OrderReject)(nil), "OrderReject")
	proto.RegisterType((*OrderCancel)(nil), "OrderCancel")
	proto.RegisterType((*OrderConfirmation)(nil), "OrderConfirmation")
	proto.RegisterType((*OrderFulfillment)(nil), "OrderFulfillment")
	proto.RegisterType((*OrderComplete)(nil), "OrderComplete")
	proto.RegisterType((*DisputeOpen)(nil), "DisputeOpen")
	proto.RegisterType((*DisputeUpdate)(nil), "DisputeUpdate")
	proto.RegisterType((*DisputeClose)(nil), "DisputeClose")
	proto.RegisterType((*Refund)(nil), "Refund")
	proto.RegisterType((*PaymentSent)(nil), "PaymentSent")
	proto.RegisterType((*PaymentFinalized)(nil), "PaymentFinalized")
}

func init() { proto.RegisterFile("orders.proto", fileDescriptor_e0f5d4cf0fc9e41b) }

var fileDescriptor_e0f5d4cf0fc9e41b = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xcf, 0xc1, 0x8a, 0x83, 0x30,
	0x10, 0x06, 0x60, 0x58, 0x16, 0xd9, 0x1d, 0x75, 0x75, 0xed, 0x1b, 0xe4, 0x01, 0x7a, 0xe9, 0x1b,
	0xd4, 0xe2, 0xd5, 0x62, 0xe9, 0xa5, 0xb7, 0x68, 0x46, 0x48, 0x89, 0x93, 0x10, 0xc7, 0x43, 0xfb,
	0xf4, 0x25, 0x69, 0xe8, 0xf1, 0xfb, 0x99, 0xf9, 0x87, 0x81, 0xc2, 0x7a, 0x85, 0x7e, 0xdd, 0x3b,
	0x6f, 0xd9, 0x8a, 0x1c, 0x7e, 0xfb, 0xe0, 0xde, 0x21, 0x89, 0x12, 0xf2, 0x88, 0x01, 0xef, 0x38,
	0xf1, 0x87, 0xad, 0xa4, 0x09, 0x8d, 0xd8, 0xc1, 0xff, 0x9b, 0x96, 0x66, 0xed, 0x17, 0xc9, 0xda,
	0x92, 0x68, 0xa0, 0x8e, 0x61, 0xb7, 0x99, 0x59, 0x1b, 0xb3, 0x20, 0xb1, 0xa8, 0xa0, 0x4c, 0x83,
	0x8b, 0x33, 0xc8, 0x18, 0x8a, 0x4e, 0x7a, 0x75, 0x1b, 0x63, 0x3c, 0x53, 0x41, 0x99, 0x78, 0x75,
	0x4a, 0x32, 0x8a, 0x3f, 0x28, 0x52, 0xd0, 0x1a, 0xbb, 0xa2, 0xf8, 0x81, 0x6c, 0xc0, 0x79, 0x23,
	0x15, 0x36, 0xcf, 0xf2, 0x11, 0x5a, 0x2f, 0xa1, 0xb9, 0x81, 0x3a, 0xb1, 0xd3, 0x24, 0x8d, 0x7e,
	0xa2, 0x3a, 0x7e, 0xdf, 0xbe, 0xdc, 0x38, 0x66, 0xf1, 0x9d, 0xc3, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0xa4, 0xc8, 0x63, 0x38, 0xde, 0x00, 0x00, 0x00,
}