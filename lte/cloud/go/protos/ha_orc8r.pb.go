//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: lte/protos/ha_orc8r.proto

package protos

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

type GetEnodebOffloadStateResponse_EnodebOffloadState int32

const (
	GetEnodebOffloadStateResponse_NO_OP                             GetEnodebOffloadStateResponse_EnodebOffloadState = 0
	GetEnodebOffloadStateResponse_PRIMARY_CONNECTED                 GetEnodebOffloadStateResponse_EnodebOffloadState = 1
	GetEnodebOffloadStateResponse_PRIMARY_CONNECTED_AND_SERVING_UES GetEnodebOffloadStateResponse_EnodebOffloadState = 2
)

// Enum value maps for GetEnodebOffloadStateResponse_EnodebOffloadState.
var (
	GetEnodebOffloadStateResponse_EnodebOffloadState_name = map[int32]string{
		0: "NO_OP",
		1: "PRIMARY_CONNECTED",
		2: "PRIMARY_CONNECTED_AND_SERVING_UES",
	}
	GetEnodebOffloadStateResponse_EnodebOffloadState_value = map[string]int32{
		"NO_OP":                             0,
		"PRIMARY_CONNECTED":                 1,
		"PRIMARY_CONNECTED_AND_SERVING_UES": 2,
	}
)

func (x GetEnodebOffloadStateResponse_EnodebOffloadState) Enum() *GetEnodebOffloadStateResponse_EnodebOffloadState {
	p := new(GetEnodebOffloadStateResponse_EnodebOffloadState)
	*p = x
	return p
}

func (x GetEnodebOffloadStateResponse_EnodebOffloadState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetEnodebOffloadStateResponse_EnodebOffloadState) Descriptor() protoreflect.EnumDescriptor {
	return file_lte_protos_ha_orc8r_proto_enumTypes[0].Descriptor()
}

func (GetEnodebOffloadStateResponse_EnodebOffloadState) Type() protoreflect.EnumType {
	return &file_lte_protos_ha_orc8r_proto_enumTypes[0]
}

func (x GetEnodebOffloadStateResponse_EnodebOffloadState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetEnodebOffloadStateResponse_EnodebOffloadState.Descriptor instead.
func (GetEnodebOffloadStateResponse_EnodebOffloadState) EnumDescriptor() ([]byte, []int) {
	return file_lte_protos_ha_orc8r_proto_rawDescGZIP(), []int{1, 0}
}

type GetEnodebOffloadStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEnodebOffloadStateRequest) Reset() {
	*x = GetEnodebOffloadStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_ha_orc8r_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnodebOffloadStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnodebOffloadStateRequest) ProtoMessage() {}

func (x *GetEnodebOffloadStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_ha_orc8r_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnodebOffloadStateRequest.ProtoReflect.Descriptor instead.
func (*GetEnodebOffloadStateRequest) Descriptor() ([]byte, []int) {
	return file_lte_protos_ha_orc8r_proto_rawDescGZIP(), []int{0}
}

type GetEnodebOffloadStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map from ENB ID to offload state
	EnodebOffloadStates map[uint32]GetEnodebOffloadStateResponse_EnodebOffloadState `protobuf:"bytes,1,rep,name=enodeb_offload_states,json=enodebOffloadStates,proto3" json:"enodeb_offload_states,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=magma.lte.GetEnodebOffloadStateResponse_EnodebOffloadState"`
}

func (x *GetEnodebOffloadStateResponse) Reset() {
	*x = GetEnodebOffloadStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_ha_orc8r_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnodebOffloadStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnodebOffloadStateResponse) ProtoMessage() {}

func (x *GetEnodebOffloadStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_ha_orc8r_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnodebOffloadStateResponse.ProtoReflect.Descriptor instead.
func (*GetEnodebOffloadStateResponse) Descriptor() ([]byte, []int) {
	return file_lte_protos_ha_orc8r_proto_rawDescGZIP(), []int{1}
}

func (x *GetEnodebOffloadStateResponse) GetEnodebOffloadStates() map[uint32]GetEnodebOffloadStateResponse_EnodebOffloadState {
	if x != nil {
		return x.EnodebOffloadStates
	}
	return nil
}

var File_lte_protos_ha_orc8r_proto protoreflect.FileDescriptor

var file_lte_protos_ha_orc8r_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6c, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x61, 0x5f,
	0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x6f,
	0x64, 0x65, 0x62, 0x4f, 0x66, 0x66, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xfb, 0x02, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x45, 0x6e,
	0x6f, 0x64, 0x65, 0x62, 0x4f, 0x66, 0x66, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x15, 0x65, 0x6e, 0x6f, 0x64,
	0x65, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e,
	0x6c, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x4f, 0x66, 0x66,
	0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x45, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x4f, 0x66, 0x66, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x65, 0x6e, 0x6f, 0x64,
	0x65, 0x62, 0x4f, 0x66, 0x66, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x1a,
	0x83, 0x01, 0x0a, 0x18, 0x45, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x4f, 0x66, 0x66, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x51,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x6f,
	0x64, 0x65, 0x62, 0x4f, 0x66, 0x66, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x4f, 0x66,
	0x66, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5d, 0x0a, 0x12, 0x45, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x4f,
	0x66, 0x66, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4e,
	0x4f, 0x5f, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x49, 0x4d, 0x41, 0x52,
	0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x25, 0x0a,
	0x21, 0x50, 0x52, 0x49, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x55,
	0x45, 0x53, 0x10, 0x02, 0x32, 0x72, 0x0a, 0x02, 0x48, 0x61, 0x12, 0x6c, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x45, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x4f, 0x66, 0x66, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x4f, 0x66, 0x66, 0x6c, 0x6f, 0x61, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x6f, 0x64,
	0x65, 0x62, 0x4f, 0x66, 0x66, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2f, 0x6c, 0x74, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lte_protos_ha_orc8r_proto_rawDescOnce sync.Once
	file_lte_protos_ha_orc8r_proto_rawDescData = file_lte_protos_ha_orc8r_proto_rawDesc
)

func file_lte_protos_ha_orc8r_proto_rawDescGZIP() []byte {
	file_lte_protos_ha_orc8r_proto_rawDescOnce.Do(func() {
		file_lte_protos_ha_orc8r_proto_rawDescData = protoimpl.X.CompressGZIP(file_lte_protos_ha_orc8r_proto_rawDescData)
	})
	return file_lte_protos_ha_orc8r_proto_rawDescData
}

var file_lte_protos_ha_orc8r_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lte_protos_ha_orc8r_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_lte_protos_ha_orc8r_proto_goTypes = []interface{}{
	(GetEnodebOffloadStateResponse_EnodebOffloadState)(0), // 0: magma.lte.GetEnodebOffloadStateResponse.EnodebOffloadState
	(*GetEnodebOffloadStateRequest)(nil),                  // 1: magma.lte.GetEnodebOffloadStateRequest
	(*GetEnodebOffloadStateResponse)(nil),                 // 2: magma.lte.GetEnodebOffloadStateResponse
	nil,                                                   // 3: magma.lte.GetEnodebOffloadStateResponse.EnodebOffloadStatesEntry
}
var file_lte_protos_ha_orc8r_proto_depIdxs = []int32{
	3, // 0: magma.lte.GetEnodebOffloadStateResponse.enodeb_offload_states:type_name -> magma.lte.GetEnodebOffloadStateResponse.EnodebOffloadStatesEntry
	0, // 1: magma.lte.GetEnodebOffloadStateResponse.EnodebOffloadStatesEntry.value:type_name -> magma.lte.GetEnodebOffloadStateResponse.EnodebOffloadState
	1, // 2: magma.lte.Ha.GetEnodebOffloadState:input_type -> magma.lte.GetEnodebOffloadStateRequest
	2, // 3: magma.lte.Ha.GetEnodebOffloadState:output_type -> magma.lte.GetEnodebOffloadStateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_lte_protos_ha_orc8r_proto_init() }
func file_lte_protos_ha_orc8r_proto_init() {
	if File_lte_protos_ha_orc8r_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lte_protos_ha_orc8r_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnodebOffloadStateRequest); i {
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
		file_lte_protos_ha_orc8r_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnodebOffloadStateResponse); i {
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
			RawDescriptor: file_lte_protos_ha_orc8r_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lte_protos_ha_orc8r_proto_goTypes,
		DependencyIndexes: file_lte_protos_ha_orc8r_proto_depIdxs,
		EnumInfos:         file_lte_protos_ha_orc8r_proto_enumTypes,
		MessageInfos:      file_lte_protos_ha_orc8r_proto_msgTypes,
	}.Build()
	File_lte_protos_ha_orc8r_proto = out.File
	file_lte_protos_ha_orc8r_proto_rawDesc = nil
	file_lte_protos_ha_orc8r_proto_goTypes = nil
	file_lte_protos_ha_orc8r_proto_depIdxs = nil
}
