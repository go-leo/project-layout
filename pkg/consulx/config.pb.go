// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pkg/consulx/config.proto

package consulx

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
		mi := &file_pkg_consulx_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_consulx_config_proto_msgTypes[0]
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
	return file_pkg_consulx_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetConfigs() map[string]*Options {
	if x != nil {
		return x.Configs
	}
	return nil
}

// Config is used to pass multiple configuration options to Sarama's constructors.
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Scheme     *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	PathPrefix *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	Datacenter *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	HttpAuth   *Options_HttpBasicAuth  `protobuf:"bytes,5,opt,name=http_auth,json=httpAuth,proto3" json:"http_auth,omitempty"`
	WaitTime   *durationpb.Duration    `protobuf:"bytes,6,opt,name=wait_time,json=waitTime,proto3" json:"wait_time,omitempty"`
	Token      *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	TokenFile  *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=token_file,json=tokenFile,proto3" json:"token_file,omitempty"`
	Namespace  *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Partition  *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=partition,proto3" json:"partition,omitempty"`
	TlsConfig  *Options_TLSConfig      `protobuf:"bytes,11,opt,name=tls_config,json=tlsConfig,proto3" json:"tls_config,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_consulx_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_consulx_config_proto_msgTypes[1]
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
	return file_pkg_consulx_config_proto_rawDescGZIP(), []int{1}
}

func (x *Options) GetAddress() *wrapperspb.StringValue {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Options) GetScheme() *wrapperspb.StringValue {
	if x != nil {
		return x.Scheme
	}
	return nil
}

func (x *Options) GetPathPrefix() *wrapperspb.StringValue {
	if x != nil {
		return x.PathPrefix
	}
	return nil
}

func (x *Options) GetDatacenter() *wrapperspb.StringValue {
	if x != nil {
		return x.Datacenter
	}
	return nil
}

func (x *Options) GetHttpAuth() *Options_HttpBasicAuth {
	if x != nil {
		return x.HttpAuth
	}
	return nil
}

func (x *Options) GetWaitTime() *durationpb.Duration {
	if x != nil {
		return x.WaitTime
	}
	return nil
}

func (x *Options) GetToken() *wrapperspb.StringValue {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *Options) GetTokenFile() *wrapperspb.StringValue {
	if x != nil {
		return x.TokenFile
	}
	return nil
}

func (x *Options) GetNamespace() *wrapperspb.StringValue {
	if x != nil {
		return x.Namespace
	}
	return nil
}

func (x *Options) GetPartition() *wrapperspb.StringValue {
	if x != nil {
		return x.Partition
	}
	return nil
}

func (x *Options) GetTlsConfig() *Options_TLSConfig {
	if x != nil {
		return x.TlsConfig
	}
	return nil
}

type Options_HttpBasicAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Options_HttpBasicAuth) Reset() {
	*x = Options_HttpBasicAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_consulx_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options_HttpBasicAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options_HttpBasicAuth) ProtoMessage() {}

func (x *Options_HttpBasicAuth) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_consulx_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options_HttpBasicAuth.ProtoReflect.Descriptor instead.
func (*Options_HttpBasicAuth) Descriptor() ([]byte, []int) {
	return file_pkg_consulx_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Options_HttpBasicAuth) GetUsername() *wrapperspb.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *Options_HttpBasicAuth) GetPassword() *wrapperspb.StringValue {
	if x != nil {
		return x.Password
	}
	return nil
}

type Options_TLSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address            *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CaFile             *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=ca_file,json=caFile,proto3" json:"ca_file,omitempty"`
	CaPath             *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=ca_path,json=caPath,proto3" json:"ca_path,omitempty"`
	CaPem              []byte                  `protobuf:"bytes,4,opt,name=ca_pem,json=caPem,proto3" json:"ca_pem,omitempty"`
	CertFile           *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty"`
	CertPem            []byte                  `protobuf:"bytes,6,opt,name=cert_pem,json=certPem,proto3" json:"cert_pem,omitempty"`
	KeyFile            *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty"`
	KeyPem             []byte                  `protobuf:"bytes,8,opt,name=key_pem,json=keyPem,proto3" json:"key_pem,omitempty"`
	InsecureSkipVerify *wrapperspb.BoolValue   `protobuf:"bytes,9,opt,name=insecure_skip_verify,json=insecureSkipVerify,proto3" json:"insecure_skip_verify,omitempty"`
}

func (x *Options_TLSConfig) Reset() {
	*x = Options_TLSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_consulx_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options_TLSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options_TLSConfig) ProtoMessage() {}

func (x *Options_TLSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_consulx_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options_TLSConfig.ProtoReflect.Descriptor instead.
func (*Options_TLSConfig) Descriptor() ([]byte, []int) {
	return file_pkg_consulx_config_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Options_TLSConfig) GetAddress() *wrapperspb.StringValue {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Options_TLSConfig) GetCaFile() *wrapperspb.StringValue {
	if x != nil {
		return x.CaFile
	}
	return nil
}

func (x *Options_TLSConfig) GetCaPath() *wrapperspb.StringValue {
	if x != nil {
		return x.CaPath
	}
	return nil
}

func (x *Options_TLSConfig) GetCaPem() []byte {
	if x != nil {
		return x.CaPem
	}
	return nil
}

func (x *Options_TLSConfig) GetCertFile() *wrapperspb.StringValue {
	if x != nil {
		return x.CertFile
	}
	return nil
}

func (x *Options_TLSConfig) GetCertPem() []byte {
	if x != nil {
		return x.CertPem
	}
	return nil
}

func (x *Options_TLSConfig) GetKeyFile() *wrapperspb.StringValue {
	if x != nil {
		return x.KeyFile
	}
	return nil
}

func (x *Options_TLSConfig) GetKeyPem() []byte {
	if x != nil {
		return x.KeyPem
	}
	return nil
}

func (x *Options_TLSConfig) GetInsecureSkipVerify() *wrapperspb.BoolValue {
	if x != nil {
		return x.InsecureSkipVerify
	}
	return nil
}

var File_pkg_consulx_config_proto protoreflect.FileDescriptor

var file_pkg_consulx_config_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x78, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6c, 0x65, 0x6f, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x78, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4,
	0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6c, 0x65, 0x6f,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x78, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x1a, 0x57, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x6c, 0x65, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x78, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xea, 0x09, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x36, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12,
	0x3d, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x3c,
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x09,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x6c, 0x65, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x78, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x36, 0x0a, 0x09, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x77, 0x61, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x3b, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x3a, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x65, 0x6f, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x78, 0x2e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x09, 0x74, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x83, 0x01, 0x0a, 0x0d,
	0x48, 0x74, 0x74, 0x70, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x38, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x1a, 0xbe, 0x03, 0x0a, 0x09, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x36, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x61, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x63, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x35,
	0x0a, 0x07, 0x63, 0x61, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x63,
	0x61, 0x50, 0x61, 0x74, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x5f, 0x70, 0x65, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x61, 0x50, 0x65, 0x6d, 0x12, 0x39, 0x0a, 0x09,
	0x63, 0x65, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x63,
	0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74, 0x5f,
	0x70, 0x65, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x65, 0x72, 0x74, 0x50,
	0x65, 0x6d, 0x12, 0x37, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6b,
	0x65, 0x79, 0x5f, 0x70, 0x65, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6b, 0x65,
	0x79, 0x50, 0x65, 0x6d, 0x12, 0x4c, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12,
	0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x78, 0x3b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_consulx_config_proto_rawDescOnce sync.Once
	file_pkg_consulx_config_proto_rawDescData = file_pkg_consulx_config_proto_rawDesc
)

func file_pkg_consulx_config_proto_rawDescGZIP() []byte {
	file_pkg_consulx_config_proto_rawDescOnce.Do(func() {
		file_pkg_consulx_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_consulx_config_proto_rawDescData)
	})
	return file_pkg_consulx_config_proto_rawDescData
}

var file_pkg_consulx_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_consulx_config_proto_goTypes = []interface{}{
	(*Config)(nil),                 // 0: leo.config.consulx.Config
	(*Options)(nil),                // 1: leo.config.consulx.Options
	nil,                            // 2: leo.config.consulx.Config.ConfigsEntry
	(*Options_HttpBasicAuth)(nil),  // 3: leo.config.consulx.Options.HttpBasicAuth
	(*Options_TLSConfig)(nil),      // 4: leo.config.consulx.Options.TLSConfig
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
	(*durationpb.Duration)(nil),    // 6: google.protobuf.Duration
	(*wrapperspb.BoolValue)(nil),   // 7: google.protobuf.BoolValue
}
var file_pkg_consulx_config_proto_depIdxs = []int32{
	2,  // 0: leo.config.consulx.Config.configs:type_name -> leo.config.consulx.Config.ConfigsEntry
	5,  // 1: leo.config.consulx.Options.address:type_name -> google.protobuf.StringValue
	5,  // 2: leo.config.consulx.Options.scheme:type_name -> google.protobuf.StringValue
	5,  // 3: leo.config.consulx.Options.path_prefix:type_name -> google.protobuf.StringValue
	5,  // 4: leo.config.consulx.Options.datacenter:type_name -> google.protobuf.StringValue
	3,  // 5: leo.config.consulx.Options.http_auth:type_name -> leo.config.consulx.Options.HttpBasicAuth
	6,  // 6: leo.config.consulx.Options.wait_time:type_name -> google.protobuf.Duration
	5,  // 7: leo.config.consulx.Options.token:type_name -> google.protobuf.StringValue
	5,  // 8: leo.config.consulx.Options.token_file:type_name -> google.protobuf.StringValue
	5,  // 9: leo.config.consulx.Options.namespace:type_name -> google.protobuf.StringValue
	5,  // 10: leo.config.consulx.Options.partition:type_name -> google.protobuf.StringValue
	4,  // 11: leo.config.consulx.Options.tls_config:type_name -> leo.config.consulx.Options.TLSConfig
	1,  // 12: leo.config.consulx.Config.ConfigsEntry.value:type_name -> leo.config.consulx.Options
	5,  // 13: leo.config.consulx.Options.HttpBasicAuth.username:type_name -> google.protobuf.StringValue
	5,  // 14: leo.config.consulx.Options.HttpBasicAuth.password:type_name -> google.protobuf.StringValue
	5,  // 15: leo.config.consulx.Options.TLSConfig.address:type_name -> google.protobuf.StringValue
	5,  // 16: leo.config.consulx.Options.TLSConfig.ca_file:type_name -> google.protobuf.StringValue
	5,  // 17: leo.config.consulx.Options.TLSConfig.ca_path:type_name -> google.protobuf.StringValue
	5,  // 18: leo.config.consulx.Options.TLSConfig.cert_file:type_name -> google.protobuf.StringValue
	5,  // 19: leo.config.consulx.Options.TLSConfig.key_file:type_name -> google.protobuf.StringValue
	7,  // 20: leo.config.consulx.Options.TLSConfig.insecure_skip_verify:type_name -> google.protobuf.BoolValue
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_pkg_consulx_config_proto_init() }
func file_pkg_consulx_config_proto_init() {
	if File_pkg_consulx_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_consulx_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_consulx_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_consulx_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options_HttpBasicAuth); i {
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
		file_pkg_consulx_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options_TLSConfig); i {
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
			RawDescriptor: file_pkg_consulx_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_consulx_config_proto_goTypes,
		DependencyIndexes: file_pkg_consulx_config_proto_depIdxs,
		MessageInfos:      file_pkg_consulx_config_proto_msgTypes,
	}.Build()
	File_pkg_consulx_config_proto = out.File
	file_pkg_consulx_config_proto_rawDesc = nil
	file_pkg_consulx_config_proto_goTypes = nil
	file_pkg_consulx_config_proto_depIdxs = nil
}
