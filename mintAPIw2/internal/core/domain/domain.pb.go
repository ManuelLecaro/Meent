// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: internal/core/domain/domain.proto

package domain

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

// Información sobre un evento.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       uint64 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	TotalTicket uint64 `protobuf:"varint,5,opt,name=totalTicket,proto3" json:"totalTicket,omitempty"`
	RealOwner   string `protobuf:"bytes,6,opt,name=realOwner,proto3" json:"realOwner,omitempty"`
	TotalReward uint64 `protobuf:"varint,7,opt,name=totalReward,proto3" json:"totalReward,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_core_domain_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_internal_core_domain_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_internal_core_domain_domain_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Event) GetTotalTicket() uint64 {
	if x != nil {
		return x.TotalTicket
	}
	return 0
}

func (x *Event) GetRealOwner() string {
	if x != nil {
		return x.RealOwner
	}
	return ""
}

func (x *Event) GetTotalReward() uint64 {
	if x != nil {
		return x.TotalReward
	}
	return 0
}

var File_internal_core_domain_domain_proto protoreflect.FileDescriptor

var file_internal_core_domain_domain_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xc5, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x61, 0x6e, 0x75, 0x65, 0x6c, 0x4c, 0x65, 0x63, 0x61, 0x72, 0x6f, 0x2f, 0x6d,
	0x69, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_core_domain_domain_proto_rawDescOnce sync.Once
	file_internal_core_domain_domain_proto_rawDescData = file_internal_core_domain_domain_proto_rawDesc
)

func file_internal_core_domain_domain_proto_rawDescGZIP() []byte {
	file_internal_core_domain_domain_proto_rawDescOnce.Do(func() {
		file_internal_core_domain_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_core_domain_domain_proto_rawDescData)
	})
	return file_internal_core_domain_domain_proto_rawDescData
}

var file_internal_core_domain_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_core_domain_domain_proto_goTypes = []interface{}{
	(*Event)(nil), // 0: domain.Event
}
var file_internal_core_domain_domain_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_core_domain_domain_proto_init() }
func file_internal_core_domain_domain_proto_init() {
	if File_internal_core_domain_domain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_core_domain_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_internal_core_domain_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_core_domain_domain_proto_goTypes,
		DependencyIndexes: file_internal_core_domain_domain_proto_depIdxs,
		MessageInfos:      file_internal_core_domain_domain_proto_msgTypes,
	}.Build()
	File_internal_core_domain_domain_proto = out.File
	file_internal_core_domain_domain_proto_rawDesc = nil
	file_internal_core_domain_domain_proto_goTypes = nil
	file_internal_core_domain_domain_proto_depIdxs = nil
}
