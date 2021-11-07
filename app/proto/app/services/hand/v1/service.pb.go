// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/services/hand/v1/service.proto

package hand

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

type Hand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParticipatePlayerIds []uint64               `protobuf:"varint,2,rep,packed,name=participate_player_ids,json=participatePlayerIds,proto3" json:"participate_player_ids,omitempty"`
	Timestamp            *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Hand) Reset() {
	*x = Hand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_hand_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hand) ProtoMessage() {}

func (x *Hand) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_hand_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hand.ProtoReflect.Descriptor instead.
func (*Hand) Descriptor() ([]byte, []int) {
	return file_app_services_hand_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Hand) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Hand) GetParticipatePlayerIds() []uint64 {
	if x != nil {
		return x.ParticipatePlayerIds
	}
	return nil
}

func (x *Hand) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type HandScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParticipatePlayerIds []uint64               `protobuf:"varint,2,rep,packed,name=participate_player_ids,json=participatePlayerIds,proto3" json:"participate_player_ids,omitempty"`
	Timestamp            *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// halfGameScores are half game cores.
	// for example, map<2, HalfGameScore{}> expresses 2th half game score.
	HalfGameScores map[uint32]*HandScore_HalfGameScore `protobuf:"bytes,4,rep,name=half_game_scores,json=halfGameScores,proto3" json:"half_game_scores,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HandScore) Reset() {
	*x = HandScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_hand_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandScore) ProtoMessage() {}

func (x *HandScore) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_hand_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandScore.ProtoReflect.Descriptor instead.
func (*HandScore) Descriptor() ([]byte, []int) {
	return file_app_services_hand_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *HandScore) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HandScore) GetParticipatePlayerIds() []uint64 {
	if x != nil {
		return x.ParticipatePlayerIds
	}
	return nil
}

func (x *HandScore) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *HandScore) GetHalfGameScores() map[uint32]*HandScore_HalfGameScore {
	if x != nil {
		return x.HalfGameScores
	}
	return nil
}

type CreateHandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp    *timestamppb.Timestamp           `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PlayerScores []*CreateHandRequest_PlayerScore `protobuf:"bytes,2,rep,name=player_scores,json=playerScores,proto3" json:"player_scores,omitempty"` // TODO: tip
}

func (x *CreateHandRequest) Reset() {
	*x = CreateHandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_hand_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHandRequest) ProtoMessage() {}

func (x *CreateHandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_hand_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHandRequest.ProtoReflect.Descriptor instead.
func (*CreateHandRequest) Descriptor() ([]byte, []int) {
	return file_app_services_hand_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateHandRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *CreateHandRequest) GetPlayerScores() []*CreateHandRequest_PlayerScore {
	if x != nil {
		return x.PlayerScores
	}
	return nil
}

type CreateHandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hand *Hand `protobuf:"bytes,1,opt,name=hand,proto3" json:"hand,omitempty"`
}

func (x *CreateHandResponse) Reset() {
	*x = CreateHandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_hand_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHandResponse) ProtoMessage() {}

func (x *CreateHandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_hand_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHandResponse.ProtoReflect.Descriptor instead.
func (*CreateHandResponse) Descriptor() ([]byte, []int) {
	return file_app_services_hand_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateHandResponse) GetHand() *Hand {
	if x != nil {
		return x.Hand
	}
	return nil
}

type FetchHandsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchHandsRequest) Reset() {
	*x = FetchHandsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_hand_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchHandsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchHandsRequest) ProtoMessage() {}

func (x *FetchHandsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_hand_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchHandsRequest.ProtoReflect.Descriptor instead.
func (*FetchHandsRequest) Descriptor() ([]byte, []int) {
	return file_app_services_hand_v1_service_proto_rawDescGZIP(), []int{4}
}

type FetchHandsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hand []*Hand `protobuf:"bytes,1,rep,name=hand,proto3" json:"hand,omitempty"`
}

func (x *FetchHandsResponse) Reset() {
	*x = FetchHandsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_hand_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchHandsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchHandsResponse) ProtoMessage() {}

func (x *FetchHandsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_hand_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchHandsResponse.ProtoReflect.Descriptor instead.
func (*FetchHandsResponse) Descriptor() ([]byte, []int) {
	return file_app_services_hand_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *FetchHandsResponse) GetHand() []*Hand {
	if x != nil {
		return x.Hand
	}
	return nil
}

type FetchHandScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FetchHandScoreRequest) Reset() {
	*x = FetchHandScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_hand_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchHandScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchHandScoreRequest) ProtoMessage() {}

func (x *FetchHandScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_hand_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchHandScoreRequest.ProtoReflect.Descriptor instead.
func (*FetchHandScoreRequest) Descriptor() ([]byte, []int) {
	return file_app_services_hand_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *FetchHandScoreRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FetchHandScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HandScore *HandScore `protobuf:"bytes,1,opt,name=hand_score,json=handScore,proto3" json:"hand_score,omitempty"`
}

func (x *FetchHandScoreResponse) Reset() {
	*x = FetchHandScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_hand_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchHandScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchHandScoreResponse) ProtoMessage() {}

func (x *FetchHandScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_hand_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchHandScoreResponse.ProtoReflect.Descriptor instead.
func (*FetchHandScoreResponse) Descriptor() ([]byte, []int) {
	return file_app_services_hand_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *FetchHandScoreResponse) GetHandScore() *HandScore {
	if x != nil {
		return x.HandScore
	}
	return nil
}

type HandScore_HalfGameScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerScores []*HandScore_HalfGameScore_PlayerScore `protobuf:"bytes,1,rep,name=player_scores,json=playerScores,proto3" json:"player_scores,omitempty"`
}

func (x *HandScore_HalfGameScore) Reset() {
	*x = HandScore_HalfGameScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_hand_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandScore_HalfGameScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandScore_HalfGameScore) ProtoMessage() {}

func (x *HandScore_HalfGameScore) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_hand_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandScore_HalfGameScore.ProtoReflect.Descriptor instead.
func (*HandScore_HalfGameScore) Descriptor() ([]byte, []int) {
	return file_app_services_hand_v1_service_proto_rawDescGZIP(), []int{1, 1}
}

func (x *HandScore_HalfGameScore) GetPlayerScores() []*HandScore_HalfGameScore_PlayerScore {
	if x != nil {
		return x.PlayerScores
	}
	return nil
}

type HandScore_HalfGameScore_PlayerScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId uint64 `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Score    int32  `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	Ranking  uint32 `protobuf:"varint,3,opt,name=ranking,proto3" json:"ranking,omitempty"`
}

func (x *HandScore_HalfGameScore_PlayerScore) Reset() {
	*x = HandScore_HalfGameScore_PlayerScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_hand_v1_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandScore_HalfGameScore_PlayerScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandScore_HalfGameScore_PlayerScore) ProtoMessage() {}

func (x *HandScore_HalfGameScore_PlayerScore) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_hand_v1_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandScore_HalfGameScore_PlayerScore.ProtoReflect.Descriptor instead.
func (*HandScore_HalfGameScore_PlayerScore) Descriptor() ([]byte, []int) {
	return file_app_services_hand_v1_service_proto_rawDescGZIP(), []int{1, 1, 0}
}

func (x *HandScore_HalfGameScore_PlayerScore) GetPlayerId() uint64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *HandScore_HalfGameScore_PlayerScore) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *HandScore_HalfGameScore_PlayerScore) GetRanking() uint32 {
	if x != nil {
		return x.Ranking
	}
	return 0
}

type CreateHandRequest_PlayerScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId   uint64 `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Score      int32  `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	GameNumber uint32 `protobuf:"varint,3,opt,name=game_number,json=gameNumber,proto3" json:"game_number,omitempty"`
}

func (x *CreateHandRequest_PlayerScore) Reset() {
	*x = CreateHandRequest_PlayerScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_services_hand_v1_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHandRequest_PlayerScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHandRequest_PlayerScore) ProtoMessage() {}

func (x *CreateHandRequest_PlayerScore) ProtoReflect() protoreflect.Message {
	mi := &file_app_services_hand_v1_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHandRequest_PlayerScore.ProtoReflect.Descriptor instead.
func (*CreateHandRequest_PlayerScore) Descriptor() ([]byte, []int) {
	return file_app_services_hand_v1_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CreateHandRequest_PlayerScore) GetPlayerId() uint64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *CreateHandRequest_PlayerScore) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *CreateHandRequest_PlayerScore) GetGameNumber() uint32 {
	if x != nil {
		return x.GameNumber
	}
	return 0
}

var File_app_services_hand_v1_service_proto protoreflect.FileDescriptor

var file_app_services_hand_v1_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x68,
	0x61, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x04,
	0x48, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0xaa, 0x04, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x04, 0x52, 0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x5d, 0x0a, 0x10, 0x68, 0x61, 0x6c, 0x66, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x61, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x61,
	0x6c, 0x66, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0e, 0x68, 0x61, 0x6c, 0x66, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x1a, 0x70, 0x0a, 0x13, 0x48, 0x61, 0x6c, 0x66, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x43, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x61, 0x6c, 0x66, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0xcb, 0x01, 0x0a, 0x0d, 0x48, 0x61, 0x6c, 0x66, 0x47, 0x61, 0x6d, 0x65,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x61, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x61,
	0x6c, 0x66, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x1a, 0x5a, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x22, 0x8a, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x58, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x0c, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x1a, 0x61, 0x0a, 0x0b, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x44,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x68, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x52, 0x04,
	0x68, 0x61, 0x6e, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x48, 0x61, 0x6e,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x12, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x04, 0x68, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x61, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x52, 0x04, 0x68, 0x61, 0x6e, 0x64, 0x22,
	0x27, 0x0a, 0x15, 0x46, 0x65, 0x74, 0x63, 0x68, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x16, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x68, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61,
	0x6e, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x09, 0x68, 0x61, 0x6e, 0x64, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x32, 0xbc, 0x02, 0x0a, 0x0b, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x6e, 0x64,
	0x12, 0x27, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x48, 0x61, 0x6e, 0x64,
	0x73, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x48, 0x61,
	0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x0e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x48, 0x61, 0x6e,
	0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x48, 0x61, 0x6e, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x61, 0x6e, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_services_hand_v1_service_proto_rawDescOnce sync.Once
	file_app_services_hand_v1_service_proto_rawDescData = file_app_services_hand_v1_service_proto_rawDesc
)

func file_app_services_hand_v1_service_proto_rawDescGZIP() []byte {
	file_app_services_hand_v1_service_proto_rawDescOnce.Do(func() {
		file_app_services_hand_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_services_hand_v1_service_proto_rawDescData)
	})
	return file_app_services_hand_v1_service_proto_rawDescData
}

var file_app_services_hand_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_app_services_hand_v1_service_proto_goTypes = []interface{}{
	(*Hand)(nil),                                // 0: app.services.hand.v1.Hand
	(*HandScore)(nil),                           // 1: app.services.hand.v1.HandScore
	(*CreateHandRequest)(nil),                   // 2: app.services.hand.v1.CreateHandRequest
	(*CreateHandResponse)(nil),                  // 3: app.services.hand.v1.CreateHandResponse
	(*FetchHandsRequest)(nil),                   // 4: app.services.hand.v1.FetchHandsRequest
	(*FetchHandsResponse)(nil),                  // 5: app.services.hand.v1.FetchHandsResponse
	(*FetchHandScoreRequest)(nil),               // 6: app.services.hand.v1.FetchHandScoreRequest
	(*FetchHandScoreResponse)(nil),              // 7: app.services.hand.v1.FetchHandScoreResponse
	nil,                                         // 8: app.services.hand.v1.HandScore.HalfGameScoresEntry
	(*HandScore_HalfGameScore)(nil),             // 9: app.services.hand.v1.HandScore.HalfGameScore
	(*HandScore_HalfGameScore_PlayerScore)(nil), // 10: app.services.hand.v1.HandScore.HalfGameScore.PlayerScore
	(*CreateHandRequest_PlayerScore)(nil),       // 11: app.services.hand.v1.CreateHandRequest.PlayerScore
	(*timestamppb.Timestamp)(nil),               // 12: google.protobuf.Timestamp
}
var file_app_services_hand_v1_service_proto_depIdxs = []int32{
	12, // 0: app.services.hand.v1.Hand.timestamp:type_name -> google.protobuf.Timestamp
	12, // 1: app.services.hand.v1.HandScore.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 2: app.services.hand.v1.HandScore.half_game_scores:type_name -> app.services.hand.v1.HandScore.HalfGameScoresEntry
	12, // 3: app.services.hand.v1.CreateHandRequest.timestamp:type_name -> google.protobuf.Timestamp
	11, // 4: app.services.hand.v1.CreateHandRequest.player_scores:type_name -> app.services.hand.v1.CreateHandRequest.PlayerScore
	0,  // 5: app.services.hand.v1.CreateHandResponse.hand:type_name -> app.services.hand.v1.Hand
	0,  // 6: app.services.hand.v1.FetchHandsResponse.hand:type_name -> app.services.hand.v1.Hand
	1,  // 7: app.services.hand.v1.FetchHandScoreResponse.hand_score:type_name -> app.services.hand.v1.HandScore
	9,  // 8: app.services.hand.v1.HandScore.HalfGameScoresEntry.value:type_name -> app.services.hand.v1.HandScore.HalfGameScore
	10, // 9: app.services.hand.v1.HandScore.HalfGameScore.player_scores:type_name -> app.services.hand.v1.HandScore.HalfGameScore.PlayerScore
	2,  // 10: app.services.hand.v1.HandService.CreateHand:input_type -> app.services.hand.v1.CreateHandRequest
	4,  // 11: app.services.hand.v1.HandService.FetchHands:input_type -> app.services.hand.v1.FetchHandsRequest
	6,  // 12: app.services.hand.v1.HandService.FetchHandScore:input_type -> app.services.hand.v1.FetchHandScoreRequest
	3,  // 13: app.services.hand.v1.HandService.CreateHand:output_type -> app.services.hand.v1.CreateHandResponse
	5,  // 14: app.services.hand.v1.HandService.FetchHands:output_type -> app.services.hand.v1.FetchHandsResponse
	7,  // 15: app.services.hand.v1.HandService.FetchHandScore:output_type -> app.services.hand.v1.FetchHandScoreResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_app_services_hand_v1_service_proto_init() }
func file_app_services_hand_v1_service_proto_init() {
	if File_app_services_hand_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_services_hand_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hand); i {
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
		file_app_services_hand_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandScore); i {
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
		file_app_services_hand_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHandRequest); i {
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
		file_app_services_hand_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHandResponse); i {
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
		file_app_services_hand_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchHandsRequest); i {
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
		file_app_services_hand_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchHandsResponse); i {
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
		file_app_services_hand_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchHandScoreRequest); i {
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
		file_app_services_hand_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchHandScoreResponse); i {
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
		file_app_services_hand_v1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandScore_HalfGameScore); i {
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
		file_app_services_hand_v1_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandScore_HalfGameScore_PlayerScore); i {
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
		file_app_services_hand_v1_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHandRequest_PlayerScore); i {
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
			RawDescriptor: file_app_services_hand_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_services_hand_v1_service_proto_goTypes,
		DependencyIndexes: file_app_services_hand_v1_service_proto_depIdxs,
		MessageInfos:      file_app_services_hand_v1_service_proto_msgTypes,
	}.Build()
	File_app_services_hand_v1_service_proto = out.File
	file_app_services_hand_v1_service_proto_rawDesc = nil
	file_app_services_hand_v1_service_proto_goTypes = nil
	file_app_services_hand_v1_service_proto_depIdxs = nil
}