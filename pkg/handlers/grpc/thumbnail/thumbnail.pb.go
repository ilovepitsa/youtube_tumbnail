// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: pkg/handlers/grpc/thumbnail/thumbnail.proto

package thumbnail

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

type ThumbnailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ThumbnailRequest) Reset() {
	*x = ThumbnailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_handlers_grpc_thumbnail_thumbnail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThumbnailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThumbnailRequest) ProtoMessage() {}

func (x *ThumbnailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_handlers_grpc_thumbnail_thumbnail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThumbnailRequest.ProtoReflect.Descriptor instead.
func (*ThumbnailRequest) Descriptor() ([]byte, []int) {
	return file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDescGZIP(), []int{0}
}

func (x *ThumbnailRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ThumbnailResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ThumbnailResponce) Reset() {
	*x = ThumbnailResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_handlers_grpc_thumbnail_thumbnail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThumbnailResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThumbnailResponce) ProtoMessage() {}

func (x *ThumbnailResponce) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_handlers_grpc_thumbnail_thumbnail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThumbnailResponce.ProtoReflect.Descriptor instead.
func (*ThumbnailResponce) Descriptor() ([]byte, []int) {
	return file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDescGZIP(), []int{1}
}

func (x *ThumbnailResponce) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_handlers_grpc_thumbnail_thumbnail_proto protoreflect.FileDescriptor

var file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x2f, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x22, 0x22, 0x0a, 0x10, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x11,
	0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xb0, 0x01, 0x0a, 0x0a, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x2e, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x22,
	0x00, 0x12, 0x55, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x73, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x1b, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x4a, 0x5a, 0x48, 0x79, 0x6f, 0x75, 0x74,
	0x75, 0x62, 0x65, 0x2d, 0x74, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDescOnce sync.Once
	file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDescData = file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDesc
)

func file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDescGZIP() []byte {
	file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDescOnce.Do(func() {
		file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDescData)
	})
	return file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDescData
}

var file_pkg_handlers_grpc_thumbnail_thumbnail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_handlers_grpc_thumbnail_thumbnail_proto_goTypes = []any{
	(*ThumbnailRequest)(nil),  // 0: thumbnail.ThumbnailRequest
	(*ThumbnailResponce)(nil), // 1: thumbnail.ThumbnailResponce
}
var file_pkg_handlers_grpc_thumbnail_thumbnail_proto_depIdxs = []int32{
	0, // 0: thumbnail.Thumbnails.GetThumbnail:input_type -> thumbnail.ThumbnailRequest
	0, // 1: thumbnail.Thumbnails.GetThumbnailsAsync:input_type -> thumbnail.ThumbnailRequest
	1, // 2: thumbnail.Thumbnails.GetThumbnail:output_type -> thumbnail.ThumbnailResponce
	1, // 3: thumbnail.Thumbnails.GetThumbnailsAsync:output_type -> thumbnail.ThumbnailResponce
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_handlers_grpc_thumbnail_thumbnail_proto_init() }
func file_pkg_handlers_grpc_thumbnail_thumbnail_proto_init() {
	if File_pkg_handlers_grpc_thumbnail_thumbnail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_handlers_grpc_thumbnail_thumbnail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ThumbnailRequest); i {
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
		file_pkg_handlers_grpc_thumbnail_thumbnail_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ThumbnailResponce); i {
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
			RawDescriptor: file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_handlers_grpc_thumbnail_thumbnail_proto_goTypes,
		DependencyIndexes: file_pkg_handlers_grpc_thumbnail_thumbnail_proto_depIdxs,
		MessageInfos:      file_pkg_handlers_grpc_thumbnail_thumbnail_proto_msgTypes,
	}.Build()
	File_pkg_handlers_grpc_thumbnail_thumbnail_proto = out.File
	file_pkg_handlers_grpc_thumbnail_thumbnail_proto_rawDesc = nil
	file_pkg_handlers_grpc_thumbnail_thumbnail_proto_goTypes = nil
	file_pkg_handlers_grpc_thumbnail_thumbnail_proto_depIdxs = nil
}
