// Code generated by protoc-gen-go. DO NOT EDIT.
// source: options/cors.proto

/*
Package options is a generated protocol buffer package.

It is generated from these files:
	options/cors.proto

It has these top-level messages:
	CORS
*/
package options

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CORS struct {
	MaxAge             int32    `protobuf:"varint,1,opt,name=max_age,json=maxAge" json:"max_age,omitempty"`
	AllowCredentials   bool     `protobuf:"varint,2,opt,name=allow_credentials,json=allowCredentials" json:"allow_credentials,omitempty"`
	OptionsPassthrough bool     `protobuf:"varint,3,opt,name=options_passthrough,json=optionsPassthrough" json:"options_passthrough,omitempty"`
	AllowOrigin        []string `protobuf:"bytes,4,rep,name=allow_origin,json=allowOrigin" json:"allow_origin,omitempty"`
	AllowMethod        []string `protobuf:"bytes,5,rep,name=allow_method,json=allowMethod" json:"allow_method,omitempty"`
	AllowHeader        []string `protobuf:"bytes,6,rep,name=allow_header,json=allowHeader" json:"allow_header,omitempty"`
	ExposeHeader       []string `protobuf:"bytes,7,rep,name=expose_header,json=exposeHeader" json:"expose_header,omitempty"`
	Debug              bool     `protobuf:"varint,8,opt,name=debug" json:"debug,omitempty"`
}

func (m *CORS) Reset()                    { *m = CORS{} }
func (m *CORS) String() string            { return proto.CompactTextString(m) }
func (*CORS) ProtoMessage()               {}
func (*CORS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CORS) GetMaxAge() int32 {
	if m != nil {
		return m.MaxAge
	}
	return 0
}

func (m *CORS) GetAllowCredentials() bool {
	if m != nil {
		return m.AllowCredentials
	}
	return false
}

func (m *CORS) GetOptionsPassthrough() bool {
	if m != nil {
		return m.OptionsPassthrough
	}
	return false
}

func (m *CORS) GetAllowOrigin() []string {
	if m != nil {
		return m.AllowOrigin
	}
	return nil
}

func (m *CORS) GetAllowMethod() []string {
	if m != nil {
		return m.AllowMethod
	}
	return nil
}

func (m *CORS) GetAllowHeader() []string {
	if m != nil {
		return m.AllowHeader
	}
	return nil
}

func (m *CORS) GetExposeHeader() []string {
	if m != nil {
		return m.ExposeHeader
	}
	return nil
}

func (m *CORS) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

var E_Cors = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*CORS)(nil),
	Field:         1043,
	Name:          "gengo.grpc.gateway.cors.cors",
	Tag:           "bytes,1043,opt,name=cors",
	Filename:      "options/cors.proto",
}

func init() {
	proto.RegisterType((*CORS)(nil), "gengo.grpc.gateway.cors.CORS")
	proto.RegisterExtension(E_Cors)
}

func init() { proto.RegisterFile("options/cors.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xd9, 0xef, 0x2d, 0x9b, 0xa0, 0x51, 0x58, 0x10, 0xc4, 0xaa, 0x97, 0x81, 0x90, 0x82,
	0xde, 0xbc, 0xe9, 0x2e, 0x5e, 0x64, 0xd2, 0xdd, 0xbc, 0x8c, 0xac, 0x7d, 0xa6, 0x81, 0xae, 0x2f,
	0x24, 0x99, 0x9b, 0x7f, 0x87, 0x7f, 0x81, 0xff, 0xa9, 0x34, 0x59, 0xe9, 0x2e, 0x1e, 0xfb, 0x79,
	0x9f, 0xf7, 0xf8, 0xf6, 0x1b, 0x42, 0x51, 0x3b, 0x85, 0xa5, 0x8d, 0x53, 0x34, 0x96, 0x6b, 0x83,
	0x0e, 0xe9, 0x54, 0x42, 0x29, 0x91, 0x4b, 0xa3, 0x53, 0x2e, 0x85, 0x83, 0x9d, 0xf8, 0xe6, 0xd5,
	0xf8, 0x32, 0x92, 0x88, 0xb2, 0x80, 0xd8, 0x6b, 0xeb, 0xed, 0x67, 0x9c, 0x81, 0x4d, 0x8d, 0xd2,
	0x0e, 0x4d, 0x58, 0xbd, 0xfd, 0x6d, 0x93, 0xee, 0x7c, 0x91, 0x2c, 0xe9, 0x94, 0x0c, 0x36, 0x62,
	0xbf, 0x12, 0x12, 0x58, 0x2b, 0x6a, 0xcd, 0x7a, 0x49, 0x7f, 0x23, 0xf6, 0xcf, 0x12, 0xe8, 0x3d,
	0x39, 0x13, 0x45, 0x81, 0xbb, 0x55, 0x6a, 0x20, 0x83, 0xd2, 0x29, 0x51, 0x58, 0xd6, 0x8e, 0x5a,
	0xb3, 0x61, 0x72, 0xea, 0x07, 0xf3, 0x86, 0xd3, 0x98, 0x9c, 0x1f, 0xf2, 0xad, 0xb4, 0xb0, 0xd6,
	0xe5, 0x06, 0xb7, 0x32, 0x67, 0x1d, 0xaf, 0xd7, 0xd1, 0xdf, 0x9b, 0x09, 0xbd, 0x21, 0x93, 0x70,
	0x1d, 0x8d, 0x92, 0xaa, 0x64, 0xdd, 0xa8, 0x33, 0x1b, 0x25, 0x63, 0xcf, 0x16, 0x1e, 0x35, 0xca,
	0x06, 0x5c, 0x8e, 0x19, 0xeb, 0x1d, 0x29, 0x6f, 0x1e, 0x35, 0x4a, 0x0e, 0x22, 0x03, 0xc3, 0xfa,
	0x47, 0xca, 0xab, 0x47, 0xf4, 0x8e, 0x9c, 0xc0, 0x5e, 0xa3, 0x85, 0xda, 0x19, 0x78, 0x67, 0x12,
	0xe0, 0x41, 0xba, 0x20, 0xbd, 0x0c, 0xd6, 0x5b, 0xc9, 0x86, 0x3e, 0x70, 0xf8, 0x78, 0x5a, 0x92,
	0x6e, 0xd5, 0x26, 0xbd, 0xe6, 0xa1, 0x4e, 0x5e, 0xd7, 0xc9, 0x97, 0x60, 0xbe, 0x54, 0x0a, 0x8b,
	0xf0, 0x5f, 0xec, 0xa7, 0x5a, 0x1b, 0x3f, 0x5c, 0xf1, 0x7f, 0xde, 0x83, 0x57, 0x4d, 0x27, 0xfe,
	0xd8, 0xcb, 0xe8, 0x63, 0x70, 0xa8, 0x63, 0xdd, 0xf7, 0xf7, 0x1e, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x41, 0x8a, 0xa1, 0x3d, 0xdb, 0x01, 0x00, 0x00,
}
