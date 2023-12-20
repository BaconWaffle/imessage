// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: ids.proto

package idsproto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KeyTransparencyLoggableData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NgmPublicIdentity []byte  `protobuf:"bytes,1,opt,name=ngmPublicIdentity,proto3,oneof" json:"ngmPublicIdentity,omitempty"`
	NgmVersion        *uint32 `protobuf:"varint,2,opt,name=ngmVersion,proto3,oneof" json:"ngmVersion,omitempty"`
	KtVersion         *uint32 `protobuf:"varint,3,opt,name=ktVersion,proto3,oneof" json:"ktVersion,omitempty"`
}

func (x *KeyTransparencyLoggableData) Reset() {
	*x = KeyTransparencyLoggableData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ids_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyTransparencyLoggableData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyTransparencyLoggableData) ProtoMessage() {}

func (x *KeyTransparencyLoggableData) ProtoReflect() protoreflect.Message {
	mi := &file_ids_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyTransparencyLoggableData.ProtoReflect.Descriptor instead.
func (*KeyTransparencyLoggableData) Descriptor() ([]byte, []int) {
	return file_ids_proto_rawDescGZIP(), []int{0}
}

func (x *KeyTransparencyLoggableData) GetNgmPublicIdentity() []byte {
	if x != nil {
		return x.NgmPublicIdentity
	}
	return nil
}

func (x *KeyTransparencyLoggableData) GetNgmVersion() uint32 {
	if x != nil && x.NgmVersion != nil {
		return *x.NgmVersion
	}
	return 0
}

func (x *KeyTransparencyLoggableData) GetKtVersion() uint32 {
	if x != nil && x.KtVersion != nil {
		return *x.KtVersion
	}
	return 0
}

type NgmPublicIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey []byte `protobuf:"bytes,1,opt,name=publicKey,proto3,oneof" json:"publicKey,omitempty"`
}

func (x *NgmPublicIdentity) Reset() {
	*x = NgmPublicIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ids_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NgmPublicIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NgmPublicIdentity) ProtoMessage() {}

func (x *NgmPublicIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_ids_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NgmPublicIdentity.ProtoReflect.Descriptor instead.
func (*NgmPublicIdentity) Descriptor() ([]byte, []int) {
	return file_ids_proto_rawDescGZIP(), []int{1}
}

func (x *NgmPublicIdentity) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type PublicDevicePrekey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prekey          []byte  `protobuf:"bytes,1,opt,name=prekey,proto3" json:"prekey,omitempty"`
	PrekeySignature []byte  `protobuf:"bytes,2,opt,name=prekeySignature,proto3" json:"prekeySignature,omitempty"`
	Timestamp       float64 `protobuf:"fixed64,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PublicDevicePrekey) Reset() {
	*x = PublicDevicePrekey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ids_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicDevicePrekey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicDevicePrekey) ProtoMessage() {}

func (x *PublicDevicePrekey) ProtoReflect() protoreflect.Message {
	mi := &file_ids_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicDevicePrekey.ProtoReflect.Descriptor instead.
func (*PublicDevicePrekey) Descriptor() ([]byte, []int) {
	return file_ids_proto_rawDescGZIP(), []int{2}
}

func (x *PublicDevicePrekey) GetPrekey() []byte {
	if x != nil {
		return x.Prekey
	}
	return nil
}

func (x *PublicDevicePrekey) GetPrekeySignature() []byte {
	if x != nil {
		return x.PrekeySignature
	}
	return nil
}

func (x *PublicDevicePrekey) GetTimestamp() float64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type OuterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncryptedPayload []byte `protobuf:"bytes,1,opt,name=encryptedPayload,proto3" json:"encryptedPayload,omitempty"`
	EphemeralPubKey  []byte `protobuf:"bytes,2,opt,name=ephemeralPubKey,proto3" json:"ephemeralPubKey,omitempty"`
	Signature        []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	KeyValidator     []byte `protobuf:"bytes,99,opt,name=keyValidator,proto3" json:"keyValidator,omitempty"`
}

func (x *OuterMessage) Reset() {
	*x = OuterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ids_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterMessage) ProtoMessage() {}

func (x *OuterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ids_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessage.ProtoReflect.Descriptor instead.
func (*OuterMessage) Descriptor() ([]byte, []int) {
	return file_ids_proto_rawDescGZIP(), []int{3}
}

func (x *OuterMessage) GetEncryptedPayload() []byte {
	if x != nil {
		return x.EncryptedPayload
	}
	return nil
}

func (x *OuterMessage) GetEphemeralPubKey() []byte {
	if x != nil {
		return x.EphemeralPubKey
	}
	return nil
}

func (x *OuterMessage) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *OuterMessage) GetKeyValidator() []byte {
	if x != nil {
		return x.KeyValidator
	}
	return nil
}

type InnerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Counter      uint32 `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	KtGossipData []byte `protobuf:"bytes,3,opt,name=ktGossipData,proto3" json:"ktGossipData,omitempty"`
	DebugInfo    []byte `protobuf:"bytes,99,opt,name=debugInfo,proto3" json:"debugInfo,omitempty"`
}

func (x *InnerMessage) Reset() {
	*x = InnerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ids_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerMessage) ProtoMessage() {}

func (x *InnerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ids_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerMessage.ProtoReflect.Descriptor instead.
func (*InnerMessage) Descriptor() ([]byte, []int) {
	return file_ids_proto_rawDescGZIP(), []int{4}
}

func (x *InnerMessage) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *InnerMessage) GetCounter() uint32 {
	if x != nil {
		return x.Counter
	}
	return 0
}

func (x *InnerMessage) GetKtGossipData() []byte {
	if x != nil {
		return x.KtGossipData
	}
	return nil
}

func (x *InnerMessage) GetDebugInfo() []byte {
	if x != nil {
		return x.DebugInfo
	}
	return nil
}

var File_ids_proto protoreflect.FileDescriptor

//go:embed ids.pb.raw
var file_ids_proto_rawDesc []byte

var (
	file_ids_proto_rawDescOnce sync.Once
	file_ids_proto_rawDescData = file_ids_proto_rawDesc
)

func file_ids_proto_rawDescGZIP() []byte {
	file_ids_proto_rawDescOnce.Do(func() {
		file_ids_proto_rawDescData = protoimpl.X.CompressGZIP(file_ids_proto_rawDescData)
	})
	return file_ids_proto_rawDescData
}

var file_ids_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ids_proto_goTypes = []interface{}{
	(*KeyTransparencyLoggableData)(nil), // 0: KeyTransparencyLoggableData
	(*NgmPublicIdentity)(nil),           // 1: NgmPublicIdentity
	(*PublicDevicePrekey)(nil),          // 2: PublicDevicePrekey
	(*OuterMessage)(nil),                // 3: OuterMessage
	(*InnerMessage)(nil),                // 4: InnerMessage
}
var file_ids_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ids_proto_init() }
func file_ids_proto_init() {
	if File_ids_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ids_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyTransparencyLoggableData); i {
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
		file_ids_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NgmPublicIdentity); i {
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
		file_ids_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicDevicePrekey); i {
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
		file_ids_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterMessage); i {
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
		file_ids_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerMessage); i {
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
	file_ids_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_ids_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ids_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ids_proto_goTypes,
		DependencyIndexes: file_ids_proto_depIdxs,
		MessageInfos:      file_ids_proto_msgTypes,
	}.Build()
	File_ids_proto = out.File
	file_ids_proto_rawDesc = nil
	file_ids_proto_goTypes = nil
	file_ids_proto_depIdxs = nil
}