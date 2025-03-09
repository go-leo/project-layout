// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pkg/nacosx/config.proto

package nacosx

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs map[string]*Options `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_nacosx_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_nacosx_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_pkg_nacosx_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetConfigs() map[string]*Options {
	if x != nil {
		return x.Configs
	}
	return nil
}

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port      *wrapperspb.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Namespace *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_nacosx_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_nacosx_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_pkg_nacosx_config_proto_rawDescGZIP(), []int{1}
}

func (x *Options) GetAddress() *wrapperspb.StringValue {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Options) GetPort() *wrapperspb.Int32Value {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *Options) GetNamespace() *wrapperspb.StringValue {
	if x != nil {
		return x.Namespace
	}
	return nil
}

var File_pkg_nacosx_config_proto protoreflect.FileDescriptor

var file_pkg_nacosx_config_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x61, 0x63, 0x6f, 0x73, 0x78, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6c, 0x65, 0x6f, 0x2e, 0x6e,
	0x61, 0x63, 0x6f, 0x73, 0x78, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a,
	0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x6e,
	0x61, 0x63, 0x6f, 0x73, 0x78, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x1a, 0x56, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x65, 0x6f,
	0x2e, 0x6e, 0x61, 0x63, 0x6f, 0x73, 0x78, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xae, 0x01, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x61, 0x63, 0x6f, 0x73,
	0x78, 0x3b, 0x6e, 0x61, 0x63, 0x6f, 0x73, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_nacosx_config_proto_rawDescOnce sync.Once
	file_pkg_nacosx_config_proto_rawDescData = file_pkg_nacosx_config_proto_rawDesc
)

func file_pkg_nacosx_config_proto_rawDescGZIP() []byte {
	file_pkg_nacosx_config_proto_rawDescOnce.Do(func() {
		file_pkg_nacosx_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_nacosx_config_proto_rawDescData)
	})
	return file_pkg_nacosx_config_proto_rawDescData
}

var file_pkg_nacosx_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_nacosx_config_proto_goTypes = []interface{}{
	(*Config)(nil),                 // 0: leo.nacosx.config.Config
	(*Options)(nil),                // 1: leo.nacosx.config.Options
	nil,                            // 2: leo.nacosx.config.Config.ConfigsEntry
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),  // 4: google.protobuf.Int32Value
}
var file_pkg_nacosx_config_proto_depIdxs = []int32{
	2, // 0: leo.nacosx.config.Config.configs:type_name -> leo.nacosx.config.Config.ConfigsEntry
	3, // 1: leo.nacosx.config.Options.address:type_name -> google.protobuf.StringValue
	4, // 2: leo.nacosx.config.Options.port:type_name -> google.protobuf.Int32Value
	3, // 3: leo.nacosx.config.Options.namespace:type_name -> google.protobuf.StringValue
	1, // 4: leo.nacosx.config.Config.ConfigsEntry.value:type_name -> leo.nacosx.config.Options
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_nacosx_config_proto_init() }
func file_pkg_nacosx_config_proto_init() {
	if File_pkg_nacosx_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_nacosx_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_pkg_nacosx_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
			RawDescriptor: file_pkg_nacosx_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_nacosx_config_proto_goTypes,
		DependencyIndexes: file_pkg_nacosx_config_proto_depIdxs,
		MessageInfos:      file_pkg_nacosx_config_proto_msgTypes,
	}.Build()
	File_pkg_nacosx_config_proto = out.File
	file_pkg_nacosx_config_proto_rawDesc = nil
	file_pkg_nacosx_config_proto_goTypes = nil
	file_pkg_nacosx_config_proto_depIdxs = nil
}
