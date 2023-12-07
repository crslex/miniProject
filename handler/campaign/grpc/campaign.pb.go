// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: handler/campaign/grpc/campaign.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetCampaignByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCampaignByIDRequest) Reset() {
	*x = GetCampaignByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_campaign_grpc_campaign_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCampaignByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCampaignByIDRequest) ProtoMessage() {}

func (x *GetCampaignByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_handler_campaign_grpc_campaign_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCampaignByIDRequest.ProtoReflect.Descriptor instead.
func (*GetCampaignByIDRequest) Descriptor() ([]byte, []int) {
	return file_handler_campaign_grpc_campaign_proto_rawDescGZIP(), []int{0}
}

func (x *GetCampaignByIDRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCampaignByListIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCampaignByListIDRequest) Reset() {
	*x = GetCampaignByListIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_campaign_grpc_campaign_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCampaignByListIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCampaignByListIDRequest) ProtoMessage() {}

func (x *GetCampaignByListIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_handler_campaign_grpc_campaign_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCampaignByListIDRequest.ProtoReflect.Descriptor instead.
func (*GetCampaignByListIDRequest) Descriptor() ([]byte, []int) {
	return file_handler_campaign_grpc_campaign_proto_rawDescGZIP(), []int{1}
}

func (x *GetCampaignByListIDRequest) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type Campaign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CampaignName string                 `protobuf:"bytes,2,opt,name=campaign_name,json=campaignName,proto3" json:"campaign_name,omitempty"`
	Start        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	Active       bool                   `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Campaign) Reset() {
	*x = Campaign{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_campaign_grpc_campaign_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Campaign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Campaign) ProtoMessage() {}

func (x *Campaign) ProtoReflect() protoreflect.Message {
	mi := &file_handler_campaign_grpc_campaign_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Campaign.ProtoReflect.Descriptor instead.
func (*Campaign) Descriptor() ([]byte, []int) {
	return file_handler_campaign_grpc_campaign_proto_rawDescGZIP(), []int{2}
}

func (x *Campaign) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Campaign) GetCampaignName() string {
	if x != nil {
		return x.CampaignName
	}
	return ""
}

func (x *Campaign) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Campaign) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *Campaign) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type GetCampaignByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Campaign []*Campaign `protobuf:"bytes,1,rep,name=campaign,proto3" json:"campaign,omitempty"`
}

func (x *GetCampaignByIDResponse) Reset() {
	*x = GetCampaignByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_campaign_grpc_campaign_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCampaignByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCampaignByIDResponse) ProtoMessage() {}

func (x *GetCampaignByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_handler_campaign_grpc_campaign_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCampaignByIDResponse.ProtoReflect.Descriptor instead.
func (*GetCampaignByIDResponse) Descriptor() ([]byte, []int) {
	return file_handler_campaign_grpc_campaign_proto_rawDescGZIP(), []int{3}
}

func (x *GetCampaignByIDResponse) GetCampaign() []*Campaign {
	if x != nil {
		return x.Campaign
	}
	return nil
}

var File_handler_campaign_grpc_campaign_proto protoreflect.FileDescriptor

var file_handler_campaign_grpc_campaign_proto_rawDesc = []byte{
	0x0a, 0x24, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x08, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a,
	0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x22, 0x49, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2e, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x32, 0xbe,
	0x01, 0x0a, 0x0f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x12, 0x49, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x22, 0x00, 0x12, 0x60, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x44, 0x12, 0x24, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_handler_campaign_grpc_campaign_proto_rawDescOnce sync.Once
	file_handler_campaign_grpc_campaign_proto_rawDescData = file_handler_campaign_grpc_campaign_proto_rawDesc
)

func file_handler_campaign_grpc_campaign_proto_rawDescGZIP() []byte {
	file_handler_campaign_grpc_campaign_proto_rawDescOnce.Do(func() {
		file_handler_campaign_grpc_campaign_proto_rawDescData = protoimpl.X.CompressGZIP(file_handler_campaign_grpc_campaign_proto_rawDescData)
	})
	return file_handler_campaign_grpc_campaign_proto_rawDescData
}

var file_handler_campaign_grpc_campaign_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_handler_campaign_grpc_campaign_proto_goTypes = []interface{}{
	(*GetCampaignByIDRequest)(nil),     // 0: campaign.GetCampaignByIDRequest
	(*GetCampaignByListIDRequest)(nil), // 1: campaign.GetCampaignByListIDRequest
	(*Campaign)(nil),                   // 2: campaign.Campaign
	(*GetCampaignByIDResponse)(nil),    // 3: campaign.GetCampaignByIDResponse
	(*timestamppb.Timestamp)(nil),      // 4: google.protobuf.Timestamp
}
var file_handler_campaign_grpc_campaign_proto_depIdxs = []int32{
	4, // 0: campaign.Campaign.start:type_name -> google.protobuf.Timestamp
	4, // 1: campaign.Campaign.end:type_name -> google.protobuf.Timestamp
	2, // 2: campaign.GetCampaignByIDResponse.campaign:type_name -> campaign.Campaign
	0, // 3: campaign.CampaignHandler.GetCampaignByID:input_type -> campaign.GetCampaignByIDRequest
	1, // 4: campaign.CampaignHandler.GetCampaignByListID:input_type -> campaign.GetCampaignByListIDRequest
	2, // 5: campaign.CampaignHandler.GetCampaignByID:output_type -> campaign.Campaign
	3, // 6: campaign.CampaignHandler.GetCampaignByListID:output_type -> campaign.GetCampaignByIDResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_handler_campaign_grpc_campaign_proto_init() }
func file_handler_campaign_grpc_campaign_proto_init() {
	if File_handler_campaign_grpc_campaign_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_handler_campaign_grpc_campaign_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCampaignByIDRequest); i {
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
		file_handler_campaign_grpc_campaign_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCampaignByListIDRequest); i {
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
		file_handler_campaign_grpc_campaign_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Campaign); i {
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
		file_handler_campaign_grpc_campaign_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCampaignByIDResponse); i {
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
			RawDescriptor: file_handler_campaign_grpc_campaign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_handler_campaign_grpc_campaign_proto_goTypes,
		DependencyIndexes: file_handler_campaign_grpc_campaign_proto_depIdxs,
		MessageInfos:      file_handler_campaign_grpc_campaign_proto_msgTypes,
	}.Build()
	File_handler_campaign_grpc_campaign_proto = out.File
	file_handler_campaign_grpc_campaign_proto_rawDesc = nil
	file_handler_campaign_grpc_campaign_proto_goTypes = nil
	file_handler_campaign_grpc_campaign_proto_depIdxs = nil
}
