// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package transaction

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AllTransaction struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllTransaction) Reset()         { *m = AllTransaction{} }
func (m *AllTransaction) String() string { return proto.CompactTextString(m) }
func (*AllTransaction) ProtoMessage()    {}
func (*AllTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

func (m *AllTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllTransaction.Unmarshal(m, b)
}
func (m *AllTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllTransaction.Marshal(b, m, deterministic)
}
func (m *AllTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllTransaction.Merge(m, src)
}
func (m *AllTransaction) XXX_Size() int {
	return xxx_messageInfo_AllTransaction.Size(m)
}
func (m *AllTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_AllTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_AllTransaction proto.InternalMessageInfo

type Transaction struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	AccountID            string   `protobuf:"bytes,2,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	CreatedAt            int64    `protobuf:"varint,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Amount               float32  `protobuf:"fixed32,5,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Notes                string   `protobuf:"bytes,6,opt,name=Notes,proto3" json:"Notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Transaction) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *Transaction) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Transaction) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Transaction) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Transaction) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

type TransactionStatus struct {
	Amount               float32  `protobuf:"fixed32,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
	State                bool     `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionStatus) Reset()         { *m = TransactionStatus{} }
func (m *TransactionStatus) String() string { return proto.CompactTextString(m) }
func (*TransactionStatus) ProtoMessage()    {}
func (*TransactionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{2}
}

func (m *TransactionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionStatus.Unmarshal(m, b)
}
func (m *TransactionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionStatus.Marshal(b, m, deterministic)
}
func (m *TransactionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionStatus.Merge(m, src)
}
func (m *TransactionStatus) XXX_Size() int {
	return xxx_messageInfo_TransactionStatus.Size(m)
}
func (m *TransactionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionStatus proto.InternalMessageInfo

func (m *TransactionStatus) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransactionStatus) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

func init() {
	proto.RegisterType((*AllTransaction)(nil), "transaction.AllTransaction")
	proto.RegisterType((*Transaction)(nil), "transaction.Transaction")
	proto.RegisterType((*TransactionStatus)(nil), "transaction.TransactionStatus")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0xbb, 0xa9, 0x0d, 0x76, 0x02, 0xc5, 0x0e, 0x22, 0x4b, 0x15, 0x09, 0x39, 0xe5, 0x54,
	0x44, 0x9f, 0x60, 0xe9, 0x1e, 0x0c, 0xf8, 0x07, 0xa2, 0x2f, 0xb0, 0xa6, 0x7b, 0x08, 0xd6, 0x6c,
	0xd8, 0x9d, 0x3e, 0x94, 0xcf, 0xe7, 0x0b, 0xc8, 0x6e, 0x84, 0x6e, 0x84, 0x9e, 0x7a, 0xfc, 0x7d,
	0x93, 0xf9, 0x66, 0x98, 0x2c, 0x2c, 0xc9, 0xaa, 0xce, 0xa9, 0x86, 0x5a, 0xd3, 0xad, 0x7b, 0x6b,
	0xc8, 0x60, 0x16, 0xa1, 0xe2, 0x02, 0x16, 0x62, 0xb7, 0x7b, 0x8f, 0xc8, 0x37, 0x83, 0x2c, 0xca,
	0xb8, 0x80, 0xa4, 0x92, 0x9c, 0xe5, 0xac, 0x9c, 0xd7, 0x49, 0x25, 0xf1, 0x06, 0xe6, 0xa2, 0x69,
	0xcc, 0xbe, 0xa3, 0x4a, 0xf2, 0x24, 0xe0, 0x03, 0xf0, 0xd5, 0x8d, 0xd5, 0x8a, 0xf4, 0x56, 0x10,
	0x9f, 0xe6, 0xac, 0x9c, 0xd6, 0x07, 0x80, 0x39, 0x64, 0x52, 0xbb, 0xc6, 0xb6, 0xbd, 0x57, 0xf3,
	0xb3, 0xd0, 0x1d, 0x23, 0xbc, 0x82, 0x54, 0x7c, 0x79, 0x17, 0x9f, 0xe5, 0xac, 0x4c, 0xea, 0xbf,
	0x84, 0x97, 0x30, 0x7b, 0x31, 0xa4, 0x1d, 0x4f, 0x43, 0xcf, 0x10, 0x0a, 0x01, 0xcb, 0x68, 0xd5,
	0x37, 0x52, 0xb4, 0x77, 0x91, 0x82, 0xfd, 0x57, 0xf8, 0x2f, 0x74, 0x58, 0xfa, 0xbc, 0x1e, 0xc2,
	0xfd, 0x0f, 0x83, 0xf8, 0x20, 0x58, 0x41, 0xf6, 0xac, 0x3e, 0xb5, 0xd4, 0xbd, 0x71, 0x2d, 0x21,
	0x5f, 0xc7, 0x07, 0x8c, 0x86, 0xad, 0x6e, 0x8f, 0x55, 0x86, 0x35, 0x8a, 0x09, 0x3e, 0x02, 0x78,
	0xd5, 0xc6, 0xea, 0xed, 0x89, 0xa6, 0x57, 0xc0, 0xa7, 0xd6, 0xd1, 0xf8, 0x4f, 0xe1, 0xf5, 0xa8,
	0x6f, 0x5c, 0x5c, 0x1d, 0x1d, 0x57, 0x4c, 0xee, 0xd8, 0x47, 0x1a, 0x9e, 0xc2, 0xc3, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x6d, 0x3a, 0xd7, 0x05, 0x1f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	MakeDeposit(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionStatus, error)
	MakeCredit(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionStatus, error)
	ListAllTransaction(ctx context.Context, in *AllTransaction, opts ...grpc.CallOption) (Transaction_ListAllTransactionClient, error)
}

type transactionClient struct {
	cc *grpc.ClientConn
}

func NewTransactionClient(cc *grpc.ClientConn) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) MakeDeposit(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionStatus, error) {
	out := new(TransactionStatus)
	err := c.cc.Invoke(ctx, "/transaction.transaction/MakeDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) MakeCredit(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionStatus, error) {
	out := new(TransactionStatus)
	err := c.cc.Invoke(ctx, "/transaction.transaction/MakeCredit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) ListAllTransaction(ctx context.Context, in *AllTransaction, opts ...grpc.CallOption) (Transaction_ListAllTransactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Transaction_serviceDesc.Streams[0], "/transaction.transaction/ListAllTransaction", opts...)
	if err != nil {
		return nil, err
	}
	x := &transactionListAllTransactionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Transaction_ListAllTransactionClient interface {
	Recv() (*Transaction, error)
	grpc.ClientStream
}

type transactionListAllTransactionClient struct {
	grpc.ClientStream
}

func (x *transactionListAllTransactionClient) Recv() (*Transaction, error) {
	m := new(Transaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	MakeDeposit(context.Context, *Transaction) (*TransactionStatus, error)
	MakeCredit(context.Context, *Transaction) (*TransactionStatus, error)
	ListAllTransaction(*AllTransaction, Transaction_ListAllTransactionServer) error
}

// UnimplementedTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (*UnimplementedTransactionServer) MakeDeposit(ctx context.Context, req *Transaction) (*TransactionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeDeposit not implemented")
}
func (*UnimplementedTransactionServer) MakeCredit(ctx context.Context, req *Transaction) (*TransactionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeCredit not implemented")
}
func (*UnimplementedTransactionServer) ListAllTransaction(req *AllTransaction, srv Transaction_ListAllTransactionServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAllTransaction not implemented")
}

func RegisterTransactionServer(s *grpc.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_MakeDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).MakeDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.transaction/MakeDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).MakeDeposit(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_MakeCredit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).MakeCredit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.transaction/MakeCredit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).MakeCredit(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_ListAllTransaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AllTransaction)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransactionServer).ListAllTransaction(m, &transactionListAllTransactionServer{stream})
}

type Transaction_ListAllTransactionServer interface {
	Send(*Transaction) error
	grpc.ServerStream
}

type transactionListAllTransactionServer struct {
	grpc.ServerStream
}

func (x *transactionListAllTransactionServer) Send(m *Transaction) error {
	return x.ServerStream.SendMsg(m)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeDeposit",
			Handler:    _Transaction_MakeDeposit_Handler,
		},
		{
			MethodName: "MakeCredit",
			Handler:    _Transaction_MakeCredit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAllTransaction",
			Handler:       _Transaction_ListAllTransaction_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transaction.proto",
}
