// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/trafficsim/trafficsim.proto

package trafficsim

import (
	context "context"
	fmt "fmt"
	types "github.com/OpenNetworkingFoundation/gmap-ran/api/types"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Change event type
type Type int32

const (
	// NONE indicates this response does not represent a modification of the Change
	Type_NONE Type = 0
	// ADDED is an event which occurs when a Change is added to the topology
	Type_ADDED Type = 1
	// UPDATED is an event which occurs when a Change is updated
	Type_UPDATED Type = 2
	// REMOVED is an event which occurs when a Change is removed from the configuration
	Type_REMOVED Type = 3
)

var Type_name = map[int32]string{
	0: "NONE",
	1: "ADDED",
	2: "UPDATED",
	3: "REMOVED",
}

var Type_value = map[string]int32{
	"NONE":    0,
	"ADDED":   1,
	"UPDATED": 2,
	"REMOVED": 3,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{0}
}

type ListTowersRequest struct {
}

func (m *ListTowersRequest) Reset()         { *m = ListTowersRequest{} }
func (m *ListTowersRequest) String() string { return proto.CompactTextString(m) }
func (*ListTowersRequest) ProtoMessage()    {}
func (*ListTowersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{0}
}
func (m *ListTowersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListTowersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListTowersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListTowersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTowersRequest.Merge(m, src)
}
func (m *ListTowersRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListTowersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTowersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTowersRequest proto.InternalMessageInfo

type ListRoutesRequest struct {
	// subscribe indicates whether to subscribe to events (e.g. ADD, UPDATE, and REMOVE) that occur
	// after all routes have been streamed to the client
	Subscribe bool `protobuf:"varint,1,opt,name=subscribe,proto3" json:"subscribe,omitempty"`
	// option to request only changes that happen after the call
	WithoutReplay bool `protobuf:"varint,2,opt,name=withoutReplay,proto3" json:"withoutReplay,omitempty"`
}

func (m *ListRoutesRequest) Reset()         { *m = ListRoutesRequest{} }
func (m *ListRoutesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRoutesRequest) ProtoMessage()    {}
func (*ListRoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{1}
}
func (m *ListRoutesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListRoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListRoutesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListRoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRoutesRequest.Merge(m, src)
}
func (m *ListRoutesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListRoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRoutesRequest proto.InternalMessageInfo

func (m *ListRoutesRequest) GetSubscribe() bool {
	if m != nil {
		return m.Subscribe
	}
	return false
}

func (m *ListRoutesRequest) GetWithoutReplay() bool {
	if m != nil {
		return m.WithoutReplay
	}
	return false
}

type ListRoutesResponse struct {
	// route is the route change on which the event occurred
	Route *types.Route `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	// type is a qualification of the type of change being made
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=ran.trafficsim.Type" json:"type,omitempty"`
}

func (m *ListRoutesResponse) Reset()         { *m = ListRoutesResponse{} }
func (m *ListRoutesResponse) String() string { return proto.CompactTextString(m) }
func (*ListRoutesResponse) ProtoMessage()    {}
func (*ListRoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{2}
}
func (m *ListRoutesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListRoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListRoutesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListRoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRoutesResponse.Merge(m, src)
}
func (m *ListRoutesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListRoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRoutesResponse proto.InternalMessageInfo

func (m *ListRoutesResponse) GetRoute() *types.Route {
	if m != nil {
		return m.Route
	}
	return nil
}

func (m *ListRoutesResponse) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_NONE
}

type ListUesRequest struct {
	// subscribe indicates whether to subscribe to events (e.g. ADD, UPDATE, and REMOVE) that occur
	// after all routes have been streamed to the client
	Subscribe bool `protobuf:"varint,1,opt,name=subscribe,proto3" json:"subscribe,omitempty"`
	// option to request only changes that happen after the call
	WithoutReplay bool `protobuf:"varint,2,opt,name=withoutReplay,proto3" json:"withoutReplay,omitempty"`
}

func (m *ListUesRequest) Reset()         { *m = ListUesRequest{} }
func (m *ListUesRequest) String() string { return proto.CompactTextString(m) }
func (*ListUesRequest) ProtoMessage()    {}
func (*ListUesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{3}
}
func (m *ListUesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListUesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListUesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListUesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUesRequest.Merge(m, src)
}
func (m *ListUesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListUesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUesRequest proto.InternalMessageInfo

func (m *ListUesRequest) GetSubscribe() bool {
	if m != nil {
		return m.Subscribe
	}
	return false
}

func (m *ListUesRequest) GetWithoutReplay() bool {
	if m != nil {
		return m.WithoutReplay
	}
	return false
}

type ListUesResponse struct {
	// Ue is the UserEquipment change on which the event occurred
	Ue *types.Ue `protobuf:"bytes,1,opt,name=ue,proto3" json:"ue,omitempty"`
	// type is a qualification of the type of change being made
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=ran.trafficsim.Type" json:"type,omitempty"`
}

func (m *ListUesResponse) Reset()         { *m = ListUesResponse{} }
func (m *ListUesResponse) String() string { return proto.CompactTextString(m) }
func (*ListUesResponse) ProtoMessage()    {}
func (*ListUesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47869854f8356ea4, []int{4}
}
func (m *ListUesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListUesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListUesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListUesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUesResponse.Merge(m, src)
}
func (m *ListUesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListUesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUesResponse proto.InternalMessageInfo

func (m *ListUesResponse) GetUe() *types.Ue {
	if m != nil {
		return m.Ue
	}
	return nil
}

func (m *ListUesResponse) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_NONE
}

func init() {
	proto.RegisterEnum("ran.trafficsim.Type", Type_name, Type_value)
	proto.RegisterType((*ListTowersRequest)(nil), "ran.trafficsim.ListTowersRequest")
	proto.RegisterType((*ListRoutesRequest)(nil), "ran.trafficsim.ListRoutesRequest")
	proto.RegisterType((*ListRoutesResponse)(nil), "ran.trafficsim.ListRoutesResponse")
	proto.RegisterType((*ListUesRequest)(nil), "ran.trafficsim.ListUesRequest")
	proto.RegisterType((*ListUesResponse)(nil), "ran.trafficsim.ListUesResponse")
}

func init() { proto.RegisterFile("api/trafficsim/trafficsim.proto", fileDescriptor_47869854f8356ea4) }

var fileDescriptor_47869854f8356ea4 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xeb, 0xd2, 0xd1, 0xed, 0x4c, 0x94, 0x62, 0xb8, 0xa8, 0x02, 0xca, 0x20, 0xe2, 0xa2,
	0x42, 0x22, 0x19, 0x45, 0x3c, 0xc0, 0xa6, 0x84, 0xab, 0x91, 0x22, 0x2b, 0x81, 0xeb, 0xa4, 0x78,
	0x9d, 0x05, 0x8d, 0x3d, 0xff, 0x51, 0xd5, 0xb7, 0xe0, 0xb1, 0xb8, 0xdc, 0x25, 0x97, 0xa8, 0x7d,
	0x0f, 0x84, 0x62, 0xaf, 0x64, 0x1d, 0x2d, 0x12, 0x12, 0x37, 0x55, 0xfd, 0xe5, 0x3b, 0xbf, 0xe3,
	0xef, 0x1c, 0xc3, 0x51, 0x21, 0x58, 0xa4, 0x65, 0x71, 0x7e, 0xce, 0x26, 0x8a, 0xcd, 0x6e, 0xfc,
	0x0d, 0x85, 0xe4, 0x9a, 0xe3, 0x9e, 0x2c, 0xaa, 0xb0, 0x51, 0xbd, 0xd3, 0x29, 0xd3, 0x17, 0xa6,
	0x0c, 0x27, 0x7c, 0x16, 0x8d, 0x05, 0xad, 0x52, 0xaa, 0xe7, 0x5c, 0x7e, 0x66, 0xd5, 0xf4, 0x2d,
	0x37, 0xd5, 0xa7, 0x42, 0x33, 0x5e, 0x45, 0xd3, 0x59, 0x21, 0x5e, 0xca, 0xa2, 0x8a, 0x2c, 0x7d,
	0x21, 0xa8, 0x72, 0xbf, 0x8e, 0x19, 0x3c, 0x84, 0x07, 0x67, 0x4c, 0xe9, 0x8c, 0xcf, 0xa9, 0x54,
	0x84, 0x5e, 0x1a, 0xaa, 0x74, 0xf0, 0xd1, 0x89, 0x84, 0x1b, 0x4d, 0xd7, 0x22, 0x7e, 0x02, 0x07,
	0xca, 0x94, 0x6a, 0x22, 0x59, 0x49, 0x07, 0xe8, 0x29, 0x1a, 0xee, 0x93, 0x46, 0xc0, 0xcf, 0xe1,
	0xde, 0x9c, 0xe9, 0x0b, 0x6e, 0x34, 0xa1, 0xe2, 0x4b, 0xb1, 0x18, 0xb4, 0xad, 0x63, 0x53, 0x0c,
	0x2e, 0x01, 0xdf, 0x04, 0x2b, 0xc1, 0x2b, 0x45, 0xf1, 0x2b, 0xd8, 0x93, 0xb5, 0x62, 0xa9, 0x87,
	0xa3, 0xc7, 0xe1, 0x66, 0xce, 0xd0, 0xdd, 0xd7, 0x16, 0x11, 0xe7, 0xc4, 0x43, 0xe8, 0xd4, 0xaa,
	0xed, 0xd2, 0x1b, 0x3d, 0xba, 0x5d, 0x91, 0x2d, 0x04, 0x25, 0xd6, 0x11, 0x64, 0xd0, 0xab, 0x5b,
	0xe6, 0xff, 0x37, 0x08, 0x85, 0xfb, 0xbf, 0xa9, 0xd7, 0x29, 0x86, 0xd0, 0x36, 0xeb, 0x08, 0x83,
	0xed, 0x11, 0x72, 0x4a, 0xda, 0xe6, 0x1f, 0x2e, 0xff, 0xe2, 0x0d, 0x74, 0xea, 0x13, 0xde, 0x87,
	0x4e, 0x3a, 0x4e, 0x93, 0x7e, 0x0b, 0x1f, 0xc0, 0xde, 0x49, 0x1c, 0x27, 0x71, 0x1f, 0xe1, 0x43,
	0xe8, 0xe6, 0xef, 0xe3, 0x93, 0x2c, 0x89, 0xfb, 0xed, 0xfa, 0x40, 0x92, 0x77, 0xe3, 0x0f, 0x49,
	0xdc, 0xbf, 0x33, 0xfa, 0x89, 0xa0, 0x9b, 0x39, 0x20, 0x4e, 0x01, 0x9a, 0x05, 0xe3, 0x67, 0xb7,
	0x9b, 0xfd, 0xb1, 0x7c, 0x6f, 0xc7, 0xf8, 0xad, 0xe9, 0x18, 0xe1, 0xdc, 0xf1, 0xdc, 0x0a, 0xb7,
	0xf3, 0x36, 0xde, 0x8d, 0x17, 0xfc, 0xcd, 0xe2, 0x66, 0x77, 0x8c, 0xf0, 0x19, 0x74, 0xaf, 0x07,
	0x8a, 0xfd, 0x6d, 0x05, 0xcd, 0xfe, 0xbc, 0xa3, 0x9d, 0xdf, 0xd7, 0xb4, 0xd3, 0xc1, 0xb7, 0xa5,
	0x8f, 0xae, 0x96, 0x3e, 0xfa, 0xb1, 0xf4, 0xd1, 0xd7, 0x95, 0xdf, 0xba, 0x5a, 0xf9, 0xad, 0xef,
	0x2b, 0xbf, 0x55, 0xde, 0xb5, 0xcf, 0xfe, 0xf5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x98,
	0x86, 0x5d, 0x6d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TrafficClient is the client API for Traffic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrafficClient interface {
	ListTowers(ctx context.Context, in *ListTowersRequest, opts ...grpc.CallOption) (Traffic_ListTowersClient, error)
	ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (Traffic_ListRoutesClient, error)
	ListUes(ctx context.Context, in *ListUesRequest, opts ...grpc.CallOption) (Traffic_ListUesClient, error)
}

type trafficClient struct {
	cc *grpc.ClientConn
}

func NewTrafficClient(cc *grpc.ClientConn) TrafficClient {
	return &trafficClient{cc}
}

func (c *trafficClient) ListTowers(ctx context.Context, in *ListTowersRequest, opts ...grpc.CallOption) (Traffic_ListTowersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Traffic_serviceDesc.Streams[0], "/ran.trafficsim.Traffic/ListTowers", opts...)
	if err != nil {
		return nil, err
	}
	x := &trafficListTowersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Traffic_ListTowersClient interface {
	Recv() (*types.Tower, error)
	grpc.ClientStream
}

type trafficListTowersClient struct {
	grpc.ClientStream
}

func (x *trafficListTowersClient) Recv() (*types.Tower, error) {
	m := new(types.Tower)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trafficClient) ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (Traffic_ListRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Traffic_serviceDesc.Streams[1], "/ran.trafficsim.Traffic/ListRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &trafficListRoutesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Traffic_ListRoutesClient interface {
	Recv() (*ListRoutesResponse, error)
	grpc.ClientStream
}

type trafficListRoutesClient struct {
	grpc.ClientStream
}

func (x *trafficListRoutesClient) Recv() (*ListRoutesResponse, error) {
	m := new(ListRoutesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *trafficClient) ListUes(ctx context.Context, in *ListUesRequest, opts ...grpc.CallOption) (Traffic_ListUesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Traffic_serviceDesc.Streams[2], "/ran.trafficsim.Traffic/ListUes", opts...)
	if err != nil {
		return nil, err
	}
	x := &trafficListUesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Traffic_ListUesClient interface {
	Recv() (*ListUesResponse, error)
	grpc.ClientStream
}

type trafficListUesClient struct {
	grpc.ClientStream
}

func (x *trafficListUesClient) Recv() (*ListUesResponse, error) {
	m := new(ListUesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TrafficServer is the server API for Traffic service.
type TrafficServer interface {
	ListTowers(*ListTowersRequest, Traffic_ListTowersServer) error
	ListRoutes(*ListRoutesRequest, Traffic_ListRoutesServer) error
	ListUes(*ListUesRequest, Traffic_ListUesServer) error
}

// UnimplementedTrafficServer can be embedded to have forward compatible implementations.
type UnimplementedTrafficServer struct {
}

func (*UnimplementedTrafficServer) ListTowers(req *ListTowersRequest, srv Traffic_ListTowersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTowers not implemented")
}
func (*UnimplementedTrafficServer) ListRoutes(req *ListRoutesRequest, srv Traffic_ListRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListRoutes not implemented")
}
func (*UnimplementedTrafficServer) ListUes(req *ListUesRequest, srv Traffic_ListUesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUes not implemented")
}

func RegisterTrafficServer(s *grpc.Server, srv TrafficServer) {
	s.RegisterService(&_Traffic_serviceDesc, srv)
}

func _Traffic_ListTowers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTowersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrafficServer).ListTowers(m, &trafficListTowersServer{stream})
}

type Traffic_ListTowersServer interface {
	Send(*types.Tower) error
	grpc.ServerStream
}

type trafficListTowersServer struct {
	grpc.ServerStream
}

func (x *trafficListTowersServer) Send(m *types.Tower) error {
	return x.ServerStream.SendMsg(m)
}

func _Traffic_ListRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRoutesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrafficServer).ListRoutes(m, &trafficListRoutesServer{stream})
}

type Traffic_ListRoutesServer interface {
	Send(*ListRoutesResponse) error
	grpc.ServerStream
}

type trafficListRoutesServer struct {
	grpc.ServerStream
}

func (x *trafficListRoutesServer) Send(m *ListRoutesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Traffic_ListUes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrafficServer).ListUes(m, &trafficListUesServer{stream})
}

type Traffic_ListUesServer interface {
	Send(*ListUesResponse) error
	grpc.ServerStream
}

type trafficListUesServer struct {
	grpc.ServerStream
}

func (x *trafficListUesServer) Send(m *ListUesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Traffic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ran.trafficsim.Traffic",
	HandlerType: (*TrafficServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTowers",
			Handler:       _Traffic_ListTowers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListRoutes",
			Handler:       _Traffic_ListRoutes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListUes",
			Handler:       _Traffic_ListUes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/trafficsim/trafficsim.proto",
}

func (m *ListTowersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListTowersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListTowersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListRoutesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListRoutesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListRoutesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WithoutReplay {
		i--
		if m.WithoutReplay {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Subscribe {
		i--
		if m.Subscribe {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListRoutesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListRoutesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListRoutesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintTrafficsim(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.Route != nil {
		{
			size, err := m.Route.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTrafficsim(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListUesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListUesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListUesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WithoutReplay {
		i--
		if m.WithoutReplay {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Subscribe {
		i--
		if m.Subscribe {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListUesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListUesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListUesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintTrafficsim(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.Ue != nil {
		{
			size, err := m.Ue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTrafficsim(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTrafficsim(dAtA []byte, offset int, v uint64) int {
	offset -= sovTrafficsim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListTowersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListRoutesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Subscribe {
		n += 2
	}
	if m.WithoutReplay {
		n += 2
	}
	return n
}

func (m *ListRoutesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Route != nil {
		l = m.Route.Size()
		n += 1 + l + sovTrafficsim(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovTrafficsim(uint64(m.Type))
	}
	return n
}

func (m *ListUesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Subscribe {
		n += 2
	}
	if m.WithoutReplay {
		n += 2
	}
	return n
}

func (m *ListUesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ue != nil {
		l = m.Ue.Size()
		n += 1 + l + sovTrafficsim(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovTrafficsim(uint64(m.Type))
	}
	return n
}

func sovTrafficsim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTrafficsim(x uint64) (n int) {
	return sovTrafficsim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListTowersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListTowersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListTowersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListRoutesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListRoutesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListRoutesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscribe", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Subscribe = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithoutReplay", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WithoutReplay = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListRoutesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListRoutesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListRoutesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Route", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTrafficsim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Route == nil {
				m.Route = &types.Route{}
			}
			if err := m.Route.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListUesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListUesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListUesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscribe", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Subscribe = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithoutReplay", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WithoutReplay = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListUesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListUesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListUesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTrafficsim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ue == nil {
				m.Ue = &types.Ue{}
			}
			if err := m.Ue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTrafficsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrafficsim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTrafficsim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrafficsim
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTrafficsim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTrafficsim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTrafficsim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTrafficsim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTrafficsim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrafficsim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTrafficsim = fmt.Errorf("proto: unexpected end of group")
)
