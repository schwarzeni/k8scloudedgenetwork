// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: proxy.pb

//option go_package = "cloudedgenetwork/proxyserver/pb";

package pb

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hello      bool        `protobuf:"varint,2,opt,name=hello,proto3" json:"hello,omitempty"`
	TargetAddr string      `protobuf:"bytes,3,opt,name=target_addr,json=targetAddr,proto3" json:"target_addr,omitempty"`
	TargetNode string      `protobuf:"bytes,4,opt,name=target_node,json=targetNode,proto3" json:"target_node,omitempty"`
	Header     *HTTPHeader `protobuf:"bytes,5,opt,name=header,proto3" json:"header,omitempty"`
	HttpMethod string      `protobuf:"bytes,6,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	Url        string      `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Body       []byte      `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	// edge --> cloud
	Nodeid string `protobuf:"bytes,9,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_pb_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_pb_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proxy_pb_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Request) GetHello() bool {
	if x != nil {
		return x.Hello
	}
	return false
}

func (x *Request) GetTargetAddr() string {
	if x != nil {
		return x.TargetAddr
	}
	return ""
}

func (x *Request) GetTargetNode() string {
	if x != nil {
		return x.TargetNode
	}
	return ""
}

func (x *Request) GetHeader() *HTTPHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Request) GetHttpMethod() string {
	if x != nil {
		return x.HttpMethod
	}
	return ""
}

func (x *Request) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Request) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Request) GetNodeid() string {
	if x != nil {
		return x.Nodeid
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hello  bool        `protobuf:"varint,2,opt,name=hello,proto3" json:"hello,omitempty"`
	Nodeid string      `protobuf:"bytes,3,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	Header *HTTPHeader `protobuf:"bytes,4,opt,name=header,proto3" json:"header,omitempty"`
	Body   []byte      `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// edge --> cloud
	Statuscode int64 `protobuf:"varint,10,opt,name=statuscode,proto3" json:"statuscode,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_pb_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_pb_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proxy_pb_rawDescGZIP(), []int{1}
}

func (x *Response) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Response) GetHello() bool {
	if x != nil {
		return x.Hello
	}
	return false
}

func (x *Response) GetNodeid() string {
	if x != nil {
		return x.Nodeid
	}
	return ""
}

func (x *Response) GetHeader() *HTTPHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Response) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Response) GetStatuscode() int64 {
	if x != nil {
		return x.Statuscode
	}
	return 0
}

type HTTPHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []*HTTPHeaderValue `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *HTTPHeader) Reset() {
	*x = HTTPHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_pb_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPHeader) ProtoMessage() {}

func (x *HTTPHeader) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_pb_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPHeader.ProtoReflect.Descriptor instead.
func (*HTTPHeader) Descriptor() ([]byte, []int) {
	return file_proxy_pb_rawDescGZIP(), []int{2}
}

func (x *HTTPHeader) GetItem() []*HTTPHeaderValue {
	if x != nil {
		return x.Item
	}
	return nil
}

type HTTPHeaderValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *HTTPHeaderValue) Reset() {
	*x = HTTPHeaderValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_pb_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPHeaderValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPHeaderValue) ProtoMessage() {}

func (x *HTTPHeaderValue) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_pb_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPHeaderValue.ProtoReflect.Descriptor instead.
func (*HTTPHeaderValue) Descriptor() ([]byte, []int) {
	return file_proxy_pb_rawDescGZIP(), []int{3}
}

func (x *HTTPHeaderValue) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_proxy_pb protoreflect.FileDescriptor

var file_proxy_pb_rawDesc = []byte{
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x62, 0x12, 0x02, 0x70, 0x62, 0x22, 0xf8,
	0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x68, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x35, 0x0a, 0x0a, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x27,
	0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x27, 0x0a, 0x0f, 0x48, 0x54, 0x54, 0x50, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x32, 0x7a, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x32, 0x45, 0x64, 0x67, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x32, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x45, 0x64, 0x67, 0x65, 0x32, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_pb_rawDescOnce sync.Once
	file_proxy_pb_rawDescData = file_proxy_pb_rawDesc
)

func file_proxy_pb_rawDescGZIP() []byte {
	file_proxy_pb_rawDescOnce.Do(func() {
		file_proxy_pb_rawDescData = protoimpl.X.CompressGZIP(file_proxy_pb_rawDescData)
	})
	return file_proxy_pb_rawDescData
}

var file_proxy_pb_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proxy_pb_goTypes = []interface{}{
	(*Request)(nil),         // 0: pb.Request
	(*Response)(nil),        // 1: pb.Response
	(*HTTPHeader)(nil),      // 2: pb.HTTPHeader
	(*HTTPHeaderValue)(nil), // 3: pb.HTTPHeaderValue
}
var file_proxy_pb_depIdxs = []int32{
	2, // 0: pb.Request.header:type_name -> pb.HTTPHeader
	2, // 1: pb.Response.header:type_name -> pb.HTTPHeader
	3, // 2: pb.HTTPHeader.item:type_name -> pb.HTTPHeaderValue
	1, // 3: pb.ProxyHttpService.ProxyCloud2Edge:input_type -> pb.Response
	0, // 4: pb.ProxyHttpService.ProxyEdge2Cloud:input_type -> pb.Request
	0, // 5: pb.ProxyHttpService.ProxyCloud2Edge:output_type -> pb.Request
	1, // 6: pb.ProxyHttpService.ProxyEdge2Cloud:output_type -> pb.Response
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proxy_pb_init() }
func file_proxy_pb_init() {
	if File_proxy_pb != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_pb_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proxy_pb_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proxy_pb_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPHeader); i {
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
		file_proxy_pb_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPHeaderValue); i {
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
			RawDescriptor: file_proxy_pb_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxy_pb_goTypes,
		DependencyIndexes: file_proxy_pb_depIdxs,
		MessageInfos:      file_proxy_pb_msgTypes,
	}.Build()
	File_proxy_pb = out.File
	file_proxy_pb_rawDesc = nil
	file_proxy_pb_goTypes = nil
	file_proxy_pb_depIdxs = nil
}
