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
// source: orc8r/cloud/go/services/device/protos/device.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	protos "magma/orc8r/lib/go/protos"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PhysicalEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Globally unique identifier per type (MAC/SN)
	DeviceID string `protobuf:"bytes,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	// Used to deserialize info
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Any other information (manufacturer, location, owner, etc)
	Info []byte `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *PhysicalEntity) Reset() {
	*x = PhysicalEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhysicalEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhysicalEntity) ProtoMessage() {}

func (x *PhysicalEntity) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhysicalEntity.ProtoReflect.Descriptor instead.
func (*PhysicalEntity) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_device_protos_device_proto_rawDescGZIP(), []int{0}
}

func (x *PhysicalEntity) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *PhysicalEntity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PhysicalEntity) GetInfo() []byte {
	if x != nil {
		return x.Info
	}
	return nil
}

type RegisterOrUpdateDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkID string            `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	Entities  []*PhysicalEntity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *RegisterOrUpdateDevicesRequest) Reset() {
	*x = RegisterOrUpdateDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterOrUpdateDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOrUpdateDevicesRequest) ProtoMessage() {}

func (x *RegisterOrUpdateDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOrUpdateDevicesRequest.ProtoReflect.Descriptor instead.
func (*RegisterOrUpdateDevicesRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_device_protos_device_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterOrUpdateDevicesRequest) GetNetworkID() string {
	if x != nil {
		return x.NetworkID
	}
	return ""
}

func (x *RegisterOrUpdateDevicesRequest) GetEntities() []*PhysicalEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

type DeviceID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceID string `protobuf:"bytes,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *DeviceID) Reset() {
	*x = DeviceID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceID) ProtoMessage() {}

func (x *DeviceID) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceID.ProtoReflect.Descriptor instead.
func (*DeviceID) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_device_protos_device_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceID) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *DeviceID) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetDeviceInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkID string      `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	DeviceIDs []*DeviceID `protobuf:"bytes,2,rep,name=deviceIDs,proto3" json:"deviceIDs,omitempty"`
}

func (x *GetDeviceInfoRequest) Reset() {
	*x = GetDeviceInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceInfoRequest) ProtoMessage() {}

func (x *GetDeviceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceInfoRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceInfoRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_device_protos_device_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeviceInfoRequest) GetNetworkID() string {
	if x != nil {
		return x.NetworkID
	}
	return ""
}

func (x *GetDeviceInfoRequest) GetDeviceIDs() []*DeviceID {
	if x != nil {
		return x.DeviceIDs
	}
	return nil
}

type GetDeviceInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A map of device IDs to corresponding PhysicalEntity structure
	DeviceMap map[string]*PhysicalEntity `protobuf:"bytes,1,rep,name=deviceMap,proto3" json:"deviceMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetDeviceInfoResponse) Reset() {
	*x = GetDeviceInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceInfoResponse) ProtoMessage() {}

func (x *GetDeviceInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceInfoResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceInfoResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_device_protos_device_proto_rawDescGZIP(), []int{4}
}

func (x *GetDeviceInfoResponse) GetDeviceMap() map[string]*PhysicalEntity {
	if x != nil {
		return x.DeviceMap
	}
	return nil
}

type DeleteDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkID string      `protobuf:"bytes,1,opt,name=networkID,proto3" json:"networkID,omitempty"`
	DeviceIDs []*DeviceID `protobuf:"bytes,2,rep,name=deviceIDs,proto3" json:"deviceIDs,omitempty"`
}

func (x *DeleteDevicesRequest) Reset() {
	*x = DeleteDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDevicesRequest) ProtoMessage() {}

func (x *DeleteDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDevicesRequest.ProtoReflect.Descriptor instead.
func (*DeleteDevicesRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_cloud_go_services_device_protos_device_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDevicesRequest) GetNetworkID() string {
	if x != nil {
		return x.NetworkID
	}
	return ""
}

func (x *DeleteDevicesRequest) GetDeviceIDs() []*DeviceID {
	if x != nil {
		return x.DeviceIDs
	}
	return nil
}

var File_orc8r_cloud_go_services_device_protos_device_proto protoreflect.FileDescriptor

var file_orc8r_cloud_go_services_device_protos_device_proto_rawDesc = []byte{
	0x0a, 0x32, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38,
	0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x19, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0e, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x7e, 0x0a, 0x1e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x3e, 0x0a, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x08, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x70, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x52, 0x09, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x56, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63,
	0x38, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x1a, 0x60, 0x0a, 0x0e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x70, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x44, 0x12, 0x3a, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63,
	0x38, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x44, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x73, 0x32, 0xf6, 0x02,
	0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x32, 0x2e, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x32, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72,
	0x63, 0x38, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x66,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x28, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e,
	0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f,
	0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orc8r_cloud_go_services_device_protos_device_proto_rawDescOnce sync.Once
	file_orc8r_cloud_go_services_device_protos_device_proto_rawDescData = file_orc8r_cloud_go_services_device_protos_device_proto_rawDesc
)

func file_orc8r_cloud_go_services_device_protos_device_proto_rawDescGZIP() []byte {
	file_orc8r_cloud_go_services_device_protos_device_proto_rawDescOnce.Do(func() {
		file_orc8r_cloud_go_services_device_protos_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_cloud_go_services_device_protos_device_proto_rawDescData)
	})
	return file_orc8r_cloud_go_services_device_protos_device_proto_rawDescData
}

var file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_orc8r_cloud_go_services_device_protos_device_proto_goTypes = []interface{}{
	(*PhysicalEntity)(nil),                 // 0: magma.orc8r.device.PhysicalEntity
	(*RegisterOrUpdateDevicesRequest)(nil), // 1: magma.orc8r.device.RegisterOrUpdateDevicesRequest
	(*DeviceID)(nil),                       // 2: magma.orc8r.device.DeviceID
	(*GetDeviceInfoRequest)(nil),           // 3: magma.orc8r.device.GetDeviceInfoRequest
	(*GetDeviceInfoResponse)(nil),          // 4: magma.orc8r.device.GetDeviceInfoResponse
	(*DeleteDevicesRequest)(nil),           // 5: magma.orc8r.device.DeleteDevicesRequest
	nil,                                    // 6: magma.orc8r.device.GetDeviceInfoResponse.DeviceMapEntry
	(*protos.Void)(nil),                    // 7: magma.orc8r.Void
}
var file_orc8r_cloud_go_services_device_protos_device_proto_depIdxs = []int32{
	0, // 0: magma.orc8r.device.RegisterOrUpdateDevicesRequest.entities:type_name -> magma.orc8r.device.PhysicalEntity
	2, // 1: magma.orc8r.device.GetDeviceInfoRequest.deviceIDs:type_name -> magma.orc8r.device.DeviceID
	6, // 2: magma.orc8r.device.GetDeviceInfoResponse.deviceMap:type_name -> magma.orc8r.device.GetDeviceInfoResponse.DeviceMapEntry
	2, // 3: magma.orc8r.device.DeleteDevicesRequest.deviceIDs:type_name -> magma.orc8r.device.DeviceID
	0, // 4: magma.orc8r.device.GetDeviceInfoResponse.DeviceMapEntry.value:type_name -> magma.orc8r.device.PhysicalEntity
	1, // 5: magma.orc8r.device.Device.RegisterDevices:input_type -> magma.orc8r.device.RegisterOrUpdateDevicesRequest
	1, // 6: magma.orc8r.device.Device.UpdateDevices:input_type -> magma.orc8r.device.RegisterOrUpdateDevicesRequest
	3, // 7: magma.orc8r.device.Device.GetDeviceInfo:input_type -> magma.orc8r.device.GetDeviceInfoRequest
	5, // 8: magma.orc8r.device.Device.DeleteDevices:input_type -> magma.orc8r.device.DeleteDevicesRequest
	7, // 9: magma.orc8r.device.Device.RegisterDevices:output_type -> magma.orc8r.Void
	7, // 10: magma.orc8r.device.Device.UpdateDevices:output_type -> magma.orc8r.Void
	4, // 11: magma.orc8r.device.Device.GetDeviceInfo:output_type -> magma.orc8r.device.GetDeviceInfoResponse
	7, // 12: magma.orc8r.device.Device.DeleteDevices:output_type -> magma.orc8r.Void
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_orc8r_cloud_go_services_device_protos_device_proto_init() }
func file_orc8r_cloud_go_services_device_protos_device_proto_init() {
	if File_orc8r_cloud_go_services_device_protos_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhysicalEntity); i {
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
		file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterOrUpdateDevicesRequest); i {
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
		file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceID); i {
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
		file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceInfoRequest); i {
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
		file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceInfoResponse); i {
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
		file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDevicesRequest); i {
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
			RawDescriptor: file_orc8r_cloud_go_services_device_protos_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orc8r_cloud_go_services_device_protos_device_proto_goTypes,
		DependencyIndexes: file_orc8r_cloud_go_services_device_protos_device_proto_depIdxs,
		MessageInfos:      file_orc8r_cloud_go_services_device_protos_device_proto_msgTypes,
	}.Build()
	File_orc8r_cloud_go_services_device_protos_device_proto = out.File
	file_orc8r_cloud_go_services_device_protos_device_proto_rawDesc = nil
	file_orc8r_cloud_go_services_device_protos_device_proto_goTypes = nil
	file_orc8r_cloud_go_services_device_protos_device_proto_depIdxs = nil
}
