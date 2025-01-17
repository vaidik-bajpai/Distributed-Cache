// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: common/api/api.proto

package api

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

type PutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutReq) Reset() {
	*x = PutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutReq) ProtoMessage() {}

func (x *PutReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutReq.ProtoReflect.Descriptor instead.
func (*PutReq) Descriptor() ([]byte, []int) {
	return file_common_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *PutReq) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PutReq) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type PutRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutRes) Reset() {
	*x = PutRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRes) ProtoMessage() {}

func (x *PutRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRes.ProtoReflect.Descriptor instead.
func (*PutRes) Descriptor() ([]byte, []int) {
	return file_common_api_api_proto_rawDescGZIP(), []int{1}
}

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_common_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetReq) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type GetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetRes) Reset() {
	*x = GetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRes) ProtoMessage() {}

func (x *GetRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRes.ProtoReflect.Descriptor instead.
func (*GetRes) Descriptor() ([]byte, []int) {
	return file_common_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetRes) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type JoinReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeID string `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	Addr   string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *JoinReq) Reset() {
	*x = JoinReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinReq) ProtoMessage() {}

func (x *JoinReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinReq.ProtoReflect.Descriptor instead.
func (*JoinReq) Descriptor() ([]byte, []int) {
	return file_common_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *JoinReq) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *JoinReq) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type JoinRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *JoinRes) Reset() {
	*x = JoinRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRes) ProtoMessage() {}

func (x *JoinRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRes.ProtoReflect.Descriptor instead.
func (*JoinRes) Descriptor() ([]byte, []int) {
	return file_common_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *JoinRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_common_api_api_proto protoreflect.FileDescriptor

var file_common_api_api_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x30, 0x0a, 0x06, 0x50,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x08, 0x0a,
	0x06, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x22, 0x1a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x1e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x35, 0x0a, 0x07, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x23, 0x0a, 0x07, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32,
	0x6d, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x1f, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12,
	0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x4a, 0x6f,
	0x69, 0x6e, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x69,
	0x64, 0x69, 0x6b, 0x2d, 0x62, 0x61, 0x6a, 0x70, 0x61, 0x69, 0x2f, 0x64, 0x2d, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_api_api_proto_rawDescOnce sync.Once
	file_common_api_api_proto_rawDescData = file_common_api_api_proto_rawDesc
)

func file_common_api_api_proto_rawDescGZIP() []byte {
	file_common_api_api_proto_rawDescOnce.Do(func() {
		file_common_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_api_api_proto_rawDescData)
	})
	return file_common_api_api_proto_rawDescData
}

var file_common_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_api_api_proto_goTypes = []any{
	(*PutReq)(nil),  // 0: api.PutReq
	(*PutRes)(nil),  // 1: api.PutRes
	(*GetReq)(nil),  // 2: api.GetReq
	(*GetRes)(nil),  // 3: api.GetRes
	(*JoinReq)(nil), // 4: api.JoinReq
	(*JoinRes)(nil), // 5: api.JoinRes
}
var file_common_api_api_proto_depIdxs = []int32{
	0, // 0: api.Cache.Put:input_type -> api.PutReq
	2, // 1: api.Cache.Get:input_type -> api.GetReq
	4, // 2: api.Cache.Join:input_type -> api.JoinReq
	1, // 3: api.Cache.Put:output_type -> api.PutRes
	3, // 4: api.Cache.Get:output_type -> api.GetRes
	5, // 5: api.Cache.Join:output_type -> api.JoinRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_api_api_proto_init() }
func file_common_api_api_proto_init() {
	if File_common_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_api_api_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PutReq); i {
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
		file_common_api_api_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PutRes); i {
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
		file_common_api_api_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetReq); i {
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
		file_common_api_api_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetRes); i {
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
		file_common_api_api_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*JoinReq); i {
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
		file_common_api_api_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*JoinRes); i {
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
			RawDescriptor: file_common_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_api_api_proto_goTypes,
		DependencyIndexes: file_common_api_api_proto_depIdxs,
		MessageInfos:      file_common_api_api_proto_msgTypes,
	}.Build()
	File_common_api_api_proto = out.File
	file_common_api_api_proto_rawDesc = nil
	file_common_api_api_proto_goTypes = nil
	file_common_api_api_proto_depIdxs = nil
}
