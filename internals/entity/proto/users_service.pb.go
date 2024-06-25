// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: users_service.proto

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

type VerifyPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *VerifyPasswordRequest) Reset() {
	*x = VerifyPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPasswordRequest) ProtoMessage() {}

func (x *VerifyPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPasswordRequest.ProtoReflect.Descriptor instead.
func (*VerifyPasswordRequest) Descriptor() ([]byte, []int) {
	return file_users_service_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyPasswordRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *VerifyPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type VerifyPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Match     bool           `protobuf:"varint,1,opt,name=match,proto3" json:"match,omitempty"`
	ClaimData *UserClaimData `protobuf:"bytes,2,opt,name=claim_data,json=claimData,proto3" json:"claim_data,omitempty"`
}

func (x *VerifyPasswordResponse) Reset() {
	*x = VerifyPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPasswordResponse) ProtoMessage() {}

func (x *VerifyPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPasswordResponse.ProtoReflect.Descriptor instead.
func (*VerifyPasswordResponse) Descriptor() ([]byte, []int) {
	return file_users_service_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyPasswordResponse) GetMatch() bool {
	if x != nil {
		return x.Match
	}
	return false
}

func (x *VerifyPasswordResponse) GetClaimData() *UserClaimData {
	if x != nil {
		return x.ClaimData
	}
	return nil
}

type UserGroupData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserGroupId string `protobuf:"bytes,1,opt,name=user_group_id,json=userGroupId,proto3" json:"user_group_id,omitempty"`
}

func (x *UserGroupData) Reset() {
	*x = UserGroupData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroupData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupData) ProtoMessage() {}

func (x *UserGroupData) ProtoReflect() protoreflect.Message {
	mi := &file_users_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupData.ProtoReflect.Descriptor instead.
func (*UserGroupData) Descriptor() ([]byte, []int) {
	return file_users_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserGroupData) GetUserGroupId() string {
	if x != nil {
		return x.UserGroupId
	}
	return ""
}

type UserClaimData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FirstName      string           `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName       string           `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	LocationTypeId int32            `protobuf:"varint,4,opt,name=location_type_id,json=locationTypeId,proto3" json:"location_type_id,omitempty"`
	LocationId     int32            `protobuf:"varint,5,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Email          string           `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	UserGroups     []*UserGroupData `protobuf:"bytes,7,rep,name=user_groups,json=userGroups,proto3" json:"user_groups,omitempty"`
}

func (x *UserClaimData) Reset() {
	*x = UserClaimData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserClaimData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserClaimData) ProtoMessage() {}

func (x *UserClaimData) ProtoReflect() protoreflect.Message {
	mi := &file_users_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserClaimData.ProtoReflect.Descriptor instead.
func (*UserClaimData) Descriptor() ([]byte, []int) {
	return file_users_service_proto_rawDescGZIP(), []int{3}
}

func (x *UserClaimData) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserClaimData) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserClaimData) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserClaimData) GetLocationTypeId() int32 {
	if x != nil {
		return x.LocationTypeId
	}
	return 0
}

func (x *UserClaimData) GetLocationId() int32 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

func (x *UserClaimData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserClaimData) GetUserGroups() []*UserGroupData {
	if x != nil {
		return x.UserGroups
	}
	return nil
}

type ClaimDataByUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ClaimDataByUserIDRequest) Reset() {
	*x = ClaimDataByUserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimDataByUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimDataByUserIDRequest) ProtoMessage() {}

func (x *ClaimDataByUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimDataByUserIDRequest.ProtoReflect.Descriptor instead.
func (*ClaimDataByUserIDRequest) Descriptor() ([]byte, []int) {
	return file_users_service_proto_rawDescGZIP(), []int{4}
}

func (x *ClaimDataByUserIDRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_users_service_proto protoreflect.FileDescriptor

var file_users_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x15,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x63, 0x0a, 0x16, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x22, 0xfc, 0x01, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x35, 0x0a, 0x0b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x22, 0x33, 0x0a, 0x18, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xba, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x1d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x41, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x47, 0x6f, 0x6e, 0x6f, 0x6e, 0x65, 0x74, 0x6c, 0x6c, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_service_proto_rawDescOnce sync.Once
	file_users_service_proto_rawDescData = file_users_service_proto_rawDesc
)

func file_users_service_proto_rawDescGZIP() []byte {
	file_users_service_proto_rawDescOnce.Do(func() {
		file_users_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_service_proto_rawDescData)
	})
	return file_users_service_proto_rawDescData
}

var file_users_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_users_service_proto_goTypes = []interface{}{
	(*VerifyPasswordRequest)(nil),    // 0: proto.VerifyPasswordRequest
	(*VerifyPasswordResponse)(nil),   // 1: proto.VerifyPasswordResponse
	(*UserGroupData)(nil),            // 2: proto.userGroupData
	(*UserClaimData)(nil),            // 3: proto.userClaimData
	(*ClaimDataByUserIDRequest)(nil), // 4: proto.ClaimDataByUserIDRequest
}
var file_users_service_proto_depIdxs = []int32{
	3, // 0: proto.VerifyPasswordResponse.claim_data:type_name -> proto.userClaimData
	2, // 1: proto.userClaimData.user_groups:type_name -> proto.userGroupData
	0, // 2: proto.UserService.VerifyPasswordAndGetClaimData:input_type -> proto.VerifyPasswordRequest
	4, // 3: proto.UserService.GetClaimDataByUserID:input_type -> proto.ClaimDataByUserIDRequest
	1, // 4: proto.UserService.VerifyPasswordAndGetClaimData:output_type -> proto.VerifyPasswordResponse
	3, // 5: proto.UserService.GetClaimDataByUserID:output_type -> proto.userClaimData
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_users_service_proto_init() }
func file_users_service_proto_init() {
	if File_users_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPasswordRequest); i {
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
		file_users_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPasswordResponse); i {
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
		file_users_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGroupData); i {
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
		file_users_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserClaimData); i {
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
		file_users_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimDataByUserIDRequest); i {
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
			RawDescriptor: file_users_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_service_proto_goTypes,
		DependencyIndexes: file_users_service_proto_depIdxs,
		MessageInfos:      file_users_service_proto_msgTypes,
	}.Build()
	File_users_service_proto = out.File
	file_users_service_proto_rawDesc = nil
	file_users_service_proto_goTypes = nil
	file_users_service_proto_depIdxs = nil
}