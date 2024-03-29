// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: calculateumur.proto

package proto

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

type RequestUmur struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TanggalLahir string `protobuf:"bytes,1,opt,name=tanggalLahir,proto3" json:"tanggalLahir,omitempty"`
}

func (x *RequestUmur) Reset() {
	*x = RequestUmur{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculateumur_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestUmur) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUmur) ProtoMessage() {}

func (x *RequestUmur) ProtoReflect() protoreflect.Message {
	mi := &file_calculateumur_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUmur.ProtoReflect.Descriptor instead.
func (*RequestUmur) Descriptor() ([]byte, []int) {
	return file_calculateumur_proto_rawDescGZIP(), []int{0}
}

func (x *RequestUmur) GetTanggalLahir() string {
	if x != nil {
		return x.TanggalLahir
	}
	return ""
}

type ResponseUmur struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Umur int32 `protobuf:"varint,1,opt,name=umur,proto3" json:"umur,omitempty"`
}

func (x *ResponseUmur) Reset() {
	*x = ResponseUmur{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculateumur_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseUmur) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseUmur) ProtoMessage() {}

func (x *ResponseUmur) ProtoReflect() protoreflect.Message {
	mi := &file_calculateumur_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseUmur.ProtoReflect.Descriptor instead.
func (*ResponseUmur) Descriptor() ([]byte, []int) {
	return file_calculateumur_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseUmur) GetUmur() int32 {
	if x != nil {
		return x.Umur
	}
	return 0
}

var File_calculateumur_proto protoreflect.FileDescriptor

var file_calculateumur_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x75, 0x6d, 0x75, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x0b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x6d, 0x75, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x74,
	0x61, 0x6e, 0x67, 0x67, 0x61, 0x6c, 0x4c, 0x61, 0x68, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x61, 0x6e, 0x67, 0x67, 0x61, 0x6c, 0x4c, 0x61, 0x68, 0x69, 0x72, 0x22,
	0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x6d, 0x75, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x6d, 0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x75,
	0x6d, 0x75, 0x72, 0x32, 0x43, 0x0a, 0x0a, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x55, 0x6d, 0x75,
	0x72, 0x12, 0x35, 0x0a, 0x0a, 0x48, 0x69, 0x74, 0x75, 0x6e, 0x67, 0x55, 0x6d, 0x75, 0x72, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55,
	0x6d, 0x75, 0x72, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x55, 0x6d, 0x75, 0x72, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculateumur_proto_rawDescOnce sync.Once
	file_calculateumur_proto_rawDescData = file_calculateumur_proto_rawDesc
)

func file_calculateumur_proto_rawDescGZIP() []byte {
	file_calculateumur_proto_rawDescOnce.Do(func() {
		file_calculateumur_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculateumur_proto_rawDescData)
	})
	return file_calculateumur_proto_rawDescData
}

var file_calculateumur_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calculateumur_proto_goTypes = []interface{}{
	(*RequestUmur)(nil),  // 0: proto.RequestUmur
	(*ResponseUmur)(nil), // 1: proto.ResponseUmur
}
var file_calculateumur_proto_depIdxs = []int32{
	0, // 0: proto.HitungUmur.HitungUmur:input_type -> proto.RequestUmur
	1, // 1: proto.HitungUmur.HitungUmur:output_type -> proto.ResponseUmur
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculateumur_proto_init() }
func file_calculateumur_proto_init() {
	if File_calculateumur_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculateumur_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestUmur); i {
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
		file_calculateumur_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseUmur); i {
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
			RawDescriptor: file_calculateumur_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculateumur_proto_goTypes,
		DependencyIndexes: file_calculateumur_proto_depIdxs,
		MessageInfos:      file_calculateumur_proto_msgTypes,
	}.Build()
	File_calculateumur_proto = out.File
	file_calculateumur_proto_rawDesc = nil
	file_calculateumur_proto_goTypes = nil
	file_calculateumur_proto_depIdxs = nil
}
