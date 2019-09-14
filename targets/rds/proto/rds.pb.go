// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/targets/rds/proto/rds.proto

package proto

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

type IPConfig_IPType int32

const (
	// Default IP of the resource.
	//  - Private IP for instance resource
	//  - Forwarding rule IP for forwarding rule.
	IPConfig_DEFAULT IPConfig_IPType = 0
	// Instance's external IP.
	IPConfig_PUBLIC IPConfig_IPType = 1
	// First IP address from the first Alias IP range. For example, for
	// alias IP range "192.168.12.0/24", 192.168.12.0 will be returned.
	// Supported only on GCE.
	IPConfig_ALIAS IPConfig_IPType = 2
)

var IPConfig_IPType_name = map[int32]string{
	0: "DEFAULT",
	1: "PUBLIC",
	2: "ALIAS",
}

var IPConfig_IPType_value = map[string]int32{
	"DEFAULT": 0,
	"PUBLIC":  1,
	"ALIAS":   2,
}

func (x IPConfig_IPType) Enum() *IPConfig_IPType {
	p := new(IPConfig_IPType)
	*p = x
	return p
}

func (x IPConfig_IPType) String() string {
	return proto.EnumName(IPConfig_IPType_name, int32(x))
}

func (x *IPConfig_IPType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IPConfig_IPType_value, data, "IPConfig_IPType")
	if err != nil {
		return err
	}
	*x = IPConfig_IPType(value)
	return nil
}

func (IPConfig_IPType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df4e3c8c1fe12b66, []int{2, 0}
}

type ListResourcesRequest struct {
	// Provider is the resource list provider, for example: "gcp", "aws", etc.
	Provider *string `protobuf:"bytes,1,req,name=provider" json:"provider,omitempty"`
	// Provider specific resource path. For example: for GCP, it could be
	// "gce_instances/<project>", "regional_forwarding_rules/<project>", etc.
	ResourcePath *string `protobuf:"bytes,2,opt,name=resource_path,json=resourcePath" json:"resource_path,omitempty"`
	// Filters for the resources list. Filters are ANDed: all filters should
	// succeed for an item to included in the result list.
	Filter []*Filter `protobuf:"bytes,3,rep,name=filter" json:"filter,omitempty"`
	// Optional. If resource has an IP (and a NIC) address, following
	// fields determine which IP address will be included in the results.
	IpConfig             *IPConfig `protobuf:"bytes,4,opt,name=ip_config,json=ipConfig" json:"ip_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResourcesRequest) Reset()         { *m = ListResourcesRequest{} }
func (m *ListResourcesRequest) String() string { return proto.CompactTextString(m) }
func (*ListResourcesRequest) ProtoMessage()    {}
func (*ListResourcesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4e3c8c1fe12b66, []int{0}
}

func (m *ListResourcesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResourcesRequest.Unmarshal(m, b)
}
func (m *ListResourcesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResourcesRequest.Marshal(b, m, deterministic)
}
func (m *ListResourcesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourcesRequest.Merge(m, src)
}
func (m *ListResourcesRequest) XXX_Size() int {
	return xxx_messageInfo_ListResourcesRequest.Size(m)
}
func (m *ListResourcesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourcesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourcesRequest proto.InternalMessageInfo

func (m *ListResourcesRequest) GetProvider() string {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return ""
}

func (m *ListResourcesRequest) GetResourcePath() string {
	if m != nil && m.ResourcePath != nil {
		return *m.ResourcePath
	}
	return ""
}

func (m *ListResourcesRequest) GetFilter() []*Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListResourcesRequest) GetIpConfig() *IPConfig {
	if m != nil {
		return m.IpConfig
	}
	return nil
}

type Filter struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4e3c8c1fe12b66, []int{1}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Filter) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type IPConfig struct {
	// NIC index
	NicIndex             *int32           `protobuf:"varint,1,opt,name=nic_index,json=nicIndex,def=0" json:"nic_index,omitempty"`
	IpType               *IPConfig_IPType `protobuf:"varint,3,opt,name=ip_type,json=ipType,enum=cloudprober.targets.rds.IPConfig_IPType" json:"ip_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IPConfig) Reset()         { *m = IPConfig{} }
func (m *IPConfig) String() string { return proto.CompactTextString(m) }
func (*IPConfig) ProtoMessage()    {}
func (*IPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4e3c8c1fe12b66, []int{2}
}

func (m *IPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPConfig.Unmarshal(m, b)
}
func (m *IPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPConfig.Marshal(b, m, deterministic)
}
func (m *IPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPConfig.Merge(m, src)
}
func (m *IPConfig) XXX_Size() int {
	return xxx_messageInfo_IPConfig.Size(m)
}
func (m *IPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IPConfig proto.InternalMessageInfo

const Default_IPConfig_NicIndex int32 = 0

func (m *IPConfig) GetNicIndex() int32 {
	if m != nil && m.NicIndex != nil {
		return *m.NicIndex
	}
	return Default_IPConfig_NicIndex
}

func (m *IPConfig) GetIpType() IPConfig_IPType {
	if m != nil && m.IpType != nil {
		return *m.IpType
	}
	return IPConfig_DEFAULT
}

type Resource struct {
	// Resource name.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// Resource's IP address, selected based on the request's ip_config.
	Ip *string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	// Resource's port, if any.
	Port *int32 `protobuf:"varint,5,opt,name=port" json:"port,omitempty"`
	// Resource's labels, if any.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Id associated with the resource, if any.
	Id *string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	// Optional info associated with the resource. Some resource type may make use
	// of it.
	Info                 []byte   `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4e3c8c1fe12b66, []int{3}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Resource) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Resource) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *Resource) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Resource) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Resource) GetInfo() []byte {
	if m != nil {
		return m.Info
	}
	return nil
}

type ListResourcesResponse struct {
	Resources            []*Resource `protobuf:"bytes,1,rep,name=resources" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListResourcesResponse) Reset()         { *m = ListResourcesResponse{} }
func (m *ListResourcesResponse) String() string { return proto.CompactTextString(m) }
func (*ListResourcesResponse) ProtoMessage()    {}
func (*ListResourcesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4e3c8c1fe12b66, []int{4}
}

func (m *ListResourcesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResourcesResponse.Unmarshal(m, b)
}
func (m *ListResourcesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResourcesResponse.Marshal(b, m, deterministic)
}
func (m *ListResourcesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourcesResponse.Merge(m, src)
}
func (m *ListResourcesResponse) XXX_Size() int {
	return xxx_messageInfo_ListResourcesResponse.Size(m)
}
func (m *ListResourcesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourcesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourcesResponse proto.InternalMessageInfo

func (m *ListResourcesResponse) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterEnum("cloudprober.targets.rds.IPConfig_IPType", IPConfig_IPType_name, IPConfig_IPType_value)
	proto.RegisterType((*ListResourcesRequest)(nil), "cloudprober.targets.rds.ListResourcesRequest")
	proto.RegisterType((*Filter)(nil), "cloudprober.targets.rds.Filter")
	proto.RegisterType((*IPConfig)(nil), "cloudprober.targets.rds.IPConfig")
	proto.RegisterType((*Resource)(nil), "cloudprober.targets.rds.Resource")
	proto.RegisterMapType((map[string]string)(nil), "cloudprober.targets.rds.Resource.LabelsEntry")
	proto.RegisterType((*ListResourcesResponse)(nil), "cloudprober.targets.rds.ListResourcesResponse")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/targets/rds/proto/rds.proto", fileDescriptor_df4e3c8c1fe12b66)
}

var fileDescriptor_df4e3c8c1fe12b66 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x9d, 0xd3, 0x35, 0x6b, 0xbe, 0x6e, 0x53, 0xb1, 0x86, 0xb0, 0x76, 0x80, 0x2c, 0x5c, 0x72,
	0x80, 0x74, 0xea, 0x05, 0xb6, 0x03, 0xa8, 0x6c, 0x9d, 0x54, 0xa9, 0x87, 0xca, 0x6c, 0x12, 0xb7,
	0x2a, 0x4d, 0xdc, 0xd6, 0x22, 0x8b, 0x8d, 0xed, 0x56, 0xe4, 0x07, 0xf0, 0x2b, 0xf8, 0x53, 0xfc,
	0x0c, 0x7e, 0x06, 0x8a, 0x93, 0x6c, 0x03, 0xad, 0x1a, 0xa7, 0x3e, 0x7f, 0x7a, 0x9f, 0xdf, 0x7b,
	0x7d, 0x0e, 0x9c, 0x2d, 0xb9, 0x59, 0xad, 0xe7, 0x51, 0x22, 0x6e, 0xfb, 0x4b, 0x21, 0x96, 0x19,
	0xeb, 0x27, 0x99, 0x58, 0xa7, 0x52, 0x89, 0x39, 0x53, 0x7d, 0x13, 0xab, 0x25, 0x33, 0xba, 0xaf,
	0x52, 0xdd, 0x97, 0x4a, 0x18, 0x51, 0xa2, 0xc8, 0x22, 0xfc, 0xe2, 0x01, 0x31, 0xaa, 0x89, 0x91,
	0x4a, 0x75, 0xf0, 0x0b, 0xc1, 0xd1, 0x84, 0x6b, 0x43, 0x99, 0x16, 0x6b, 0x95, 0x30, 0x4d, 0xd9,
	0xb7, 0x35, 0xd3, 0x06, 0x1f, 0x43, 0x47, 0x2a, 0xb1, 0xe1, 0x29, 0x53, 0x04, 0xf9, 0x4e, 0xe8,
	0xd1, 0xbb, 0x33, 0x7e, 0x0d, 0x07, 0xaa, 0xe6, 0xcf, 0x64, 0x6c, 0x56, 0xc4, 0xf1, 0x51, 0xe8,
	0xd1, 0xfd, 0x66, 0x38, 0x8d, 0xcd, 0x0a, 0xbf, 0x03, 0x77, 0xc1, 0x33, 0xc3, 0x14, 0x69, 0xf9,
	0xad, 0xb0, 0x3b, 0x78, 0x15, 0x6d, 0xf1, 0x10, 0x5d, 0x59, 0x1a, 0xad, 0xe9, 0xf8, 0x03, 0x78,
	0x5c, 0xce, 0x12, 0x91, 0x2f, 0xf8, 0x92, 0xec, 0xfa, 0x28, 0xec, 0x0e, 0x4e, 0xb6, 0xee, 0x8e,
	0xa7, 0x17, 0x96, 0x48, 0x3b, 0x5c, 0x56, 0x28, 0x38, 0x05, 0xb7, 0xba, 0x11, 0xf7, 0xa0, 0xf5,
	0x95, 0x15, 0xb5, 0xfd, 0x12, 0xe2, 0x23, 0x68, 0x6f, 0xe2, 0x6c, 0xcd, 0x88, 0x63, 0x67, 0xd5,
	0x21, 0xf8, 0x89, 0xa0, 0xd3, 0x5c, 0x84, 0x5f, 0x82, 0x97, 0xf3, 0x64, 0xc6, 0xf3, 0x94, 0x7d,
	0x27, 0xc8, 0x47, 0x61, 0xfb, 0x1c, 0x9d, 0xd2, 0x4e, 0xce, 0x93, 0x71, 0x39, 0xc2, 0x43, 0xd8,
	0xe3, 0x72, 0x66, 0x0a, 0xc9, 0x48, 0xcb, 0x47, 0xe1, 0xe1, 0x20, 0x7c, 0xd2, 0x5c, 0x34, 0x9e,
	0x5e, 0x17, 0x92, 0x51, 0x97, 0xcb, 0xf2, 0x37, 0x78, 0x03, 0x6e, 0x35, 0xc1, 0x5d, 0xd8, 0xbb,
	0x1c, 0x5d, 0x0d, 0x6f, 0x26, 0xd7, 0xbd, 0x1d, 0x0c, 0xe0, 0x4e, 0x6f, 0x3e, 0x4d, 0xc6, 0x17,
	0x3d, 0x84, 0x3d, 0x68, 0x0f, 0x27, 0xe3, 0xe1, 0xe7, 0x9e, 0x13, 0xfc, 0x46, 0xd0, 0x69, 0xea,
	0xc1, 0x18, 0x76, 0xf3, 0xf8, 0x96, 0xd5, 0x99, 0x2c, 0xc6, 0x87, 0xe0, 0x70, 0x59, 0x77, 0xe0,
	0x70, 0x59, 0x72, 0xa4, 0x50, 0x86, 0xb4, 0x4b, 0xf3, 0xd4, 0x62, 0x3c, 0x02, 0x37, 0x8b, 0xe7,
	0x2c, 0xd3, 0xc4, 0xb5, 0x6d, 0xbc, 0xdd, 0x6a, 0xba, 0x91, 0x8a, 0x26, 0x96, 0x3f, 0xca, 0x8d,
	0x2a, 0x68, 0xbd, 0x6c, 0xa5, 0x52, 0x9b, 0xbb, 0x94, 0x4a, 0x4b, 0x29, 0x9e, 0x2f, 0x84, 0xad,
	0x69, 0x9f, 0x5a, 0x7c, 0x7c, 0x06, 0xdd, 0x07, 0xab, 0xf7, 0x25, 0xa0, 0x47, 0x4a, 0x40, 0x77,
	0x25, 0x9c, 0x3b, 0xef, 0x51, 0xf0, 0x05, 0x9e, 0xff, 0xf3, 0x18, 0xb5, 0x14, 0xb9, 0x66, 0xf8,
	0x23, 0x78, 0xcd, 0xe3, 0xd2, 0x04, 0xd9, 0x04, 0x27, 0x4f, 0x26, 0xa0, 0xf7, 0x3b, 0x83, 0x1f,
	0x08, 0x9e, 0x35, 0xf3, 0x4b, 0xae, 0x13, 0xb1, 0x61, 0xaa, 0xc0, 0x12, 0x0e, 0xfe, 0xd2, 0xc3,
	0xdb, 0xff, 0x96, 0xc7, 0x3e, 0x92, 0xe3, 0xe8, 0x7f, 0xe9, 0x55, 0x8c, 0x60, 0xe7, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x54, 0x9d, 0x7e, 0x73, 0xc4, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceDiscoveryClient is the client API for ResourceDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceDiscoveryClient interface {
	// ListResources returns the list of resources matching the URI provided in
	// the request.
	ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error)
}

type resourceDiscoveryClient struct {
	cc *grpc.ClientConn
}

func NewResourceDiscoveryClient(cc *grpc.ClientConn) ResourceDiscoveryClient {
	return &resourceDiscoveryClient{cc}
}

func (c *resourceDiscoveryClient) ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error) {
	out := new(ListResourcesResponse)
	err := c.cc.Invoke(ctx, "/cloudprober.targets.rds.ResourceDiscovery/ListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceDiscoveryServer is the server API for ResourceDiscovery service.
type ResourceDiscoveryServer interface {
	// ListResources returns the list of resources matching the URI provided in
	// the request.
	ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error)
}

// UnimplementedResourceDiscoveryServer can be embedded to have forward compatible implementations.
type UnimplementedResourceDiscoveryServer struct {
}

func (*UnimplementedResourceDiscoveryServer) ListResources(ctx context.Context, req *ListResourcesRequest) (*ListResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResources not implemented")
}

func RegisterResourceDiscoveryServer(s *grpc.Server, srv ResourceDiscoveryServer) {
	s.RegisterService(&_ResourceDiscovery_serviceDesc, srv)
}

func _ResourceDiscovery_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceDiscoveryServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprober.targets.rds.ResourceDiscovery/ListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceDiscoveryServer).ListResources(ctx, req.(*ListResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceDiscovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudprober.targets.rds.ResourceDiscovery",
	HandlerType: (*ResourceDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListResources",
			Handler:    _ResourceDiscovery_ListResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/google/cloudprober/targets/rds/proto/rds.proto",
}
