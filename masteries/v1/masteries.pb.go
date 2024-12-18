// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: masteries/v1/masteries.proto

package masteriesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "maxischmaxi/jstreams/summoner/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetChampionMasteriesRequeset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Puuid  string                   `protobuf:"bytes,1,opt,name=puuid,proto3" json:"puuid,omitempty"`
	Region v1.PlatformRoutingValues `protobuf:"varint,2,opt,name=region,proto3,enum=summoner.PlatformRoutingValues" json:"region,omitempty"`
}

func (x *GetChampionMasteriesRequeset) Reset() {
	*x = GetChampionMasteriesRequeset{}
	mi := &file_masteries_v1_masteries_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChampionMasteriesRequeset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChampionMasteriesRequeset) ProtoMessage() {}

func (x *GetChampionMasteriesRequeset) ProtoReflect() protoreflect.Message {
	mi := &file_masteries_v1_masteries_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChampionMasteriesRequeset.ProtoReflect.Descriptor instead.
func (*GetChampionMasteriesRequeset) Descriptor() ([]byte, []int) {
	return file_masteries_v1_masteries_proto_rawDescGZIP(), []int{0}
}

func (x *GetChampionMasteriesRequeset) GetPuuid() string {
	if x != nil {
		return x.Puuid
	}
	return ""
}

func (x *GetChampionMasteriesRequeset) GetRegion() v1.PlatformRoutingValues {
	if x != nil {
		return x.Region
	}
	return v1.PlatformRoutingValues(0)
}

type NextSeasonMilestone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RewardMarks        int32            `protobuf:"varint,1,opt,name=rewardMarks,proto3" json:"rewardMarks,omitempty"`
	Bonus              bool             `protobuf:"varint,2,opt,name=bonus,proto3" json:"bonus,omitempty"`
	TotalGamesRequires int32            `protobuf:"varint,3,opt,name=totalGamesRequires,proto3" json:"totalGamesRequires,omitempty"`
	RequireGradeCounts map[string]int32 `protobuf:"bytes,4,rep,name=requireGradeCounts,proto3" json:"requireGradeCounts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *NextSeasonMilestone) Reset() {
	*x = NextSeasonMilestone{}
	mi := &file_masteries_v1_masteries_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NextSeasonMilestone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextSeasonMilestone) ProtoMessage() {}

func (x *NextSeasonMilestone) ProtoReflect() protoreflect.Message {
	mi := &file_masteries_v1_masteries_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextSeasonMilestone.ProtoReflect.Descriptor instead.
func (*NextSeasonMilestone) Descriptor() ([]byte, []int) {
	return file_masteries_v1_masteries_proto_rawDescGZIP(), []int{1}
}

func (x *NextSeasonMilestone) GetRewardMarks() int32 {
	if x != nil {
		return x.RewardMarks
	}
	return 0
}

func (x *NextSeasonMilestone) GetBonus() bool {
	if x != nil {
		return x.Bonus
	}
	return false
}

func (x *NextSeasonMilestone) GetTotalGamesRequires() int32 {
	if x != nil {
		return x.TotalGamesRequires
	}
	return 0
}

func (x *NextSeasonMilestone) GetRequireGradeCounts() map[string]int32 {
	if x != nil {
		return x.RequireGradeCounts
	}
	return nil
}

type ChampionMastery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Puuid                        string               `protobuf:"bytes,1,opt,name=puuid,proto3" json:"puuid,omitempty"`
	ChampionId                   int32                `protobuf:"varint,2,opt,name=championId,proto3" json:"championId,omitempty"`
	ChampionLevel                int32                `protobuf:"varint,3,opt,name=championLevel,proto3" json:"championLevel,omitempty"`
	ChampionPoints               int32                `protobuf:"varint,4,opt,name=championPoints,proto3" json:"championPoints,omitempty"`
	LastPlayTime                 int64                `protobuf:"varint,5,opt,name=lastPlayTime,proto3" json:"lastPlayTime,omitempty"`
	ChampionPointsSinceLastLevel int32                `protobuf:"varint,6,opt,name=championPointsSinceLastLevel,proto3" json:"championPointsSinceLastLevel,omitempty"`
	ChampionPointsUntilNextLevel int32                `protobuf:"varint,7,opt,name=championPointsUntilNextLevel,proto3" json:"championPointsUntilNextLevel,omitempty"`
	MarkRequiredForNextLevel     int32                `protobuf:"varint,8,opt,name=markRequiredForNextLevel,proto3" json:"markRequiredForNextLevel,omitempty"`
	TokensEarned                 int32                `protobuf:"varint,9,opt,name=tokensEarned,proto3" json:"tokensEarned,omitempty"`
	ChampionSeasonMilestone      int32                `protobuf:"varint,10,opt,name=championSeasonMilestone,proto3" json:"championSeasonMilestone,omitempty"`
	MilestoneGrades              []string             `protobuf:"bytes,11,rep,name=milestoneGrades,proto3" json:"milestoneGrades,omitempty"`
	NextSeasonMilestone          *NextSeasonMilestone `protobuf:"bytes,12,opt,name=nextSeasonMilestone,proto3" json:"nextSeasonMilestone,omitempty"`
}

func (x *ChampionMastery) Reset() {
	*x = ChampionMastery{}
	mi := &file_masteries_v1_masteries_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChampionMastery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChampionMastery) ProtoMessage() {}

func (x *ChampionMastery) ProtoReflect() protoreflect.Message {
	mi := &file_masteries_v1_masteries_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChampionMastery.ProtoReflect.Descriptor instead.
func (*ChampionMastery) Descriptor() ([]byte, []int) {
	return file_masteries_v1_masteries_proto_rawDescGZIP(), []int{2}
}

func (x *ChampionMastery) GetPuuid() string {
	if x != nil {
		return x.Puuid
	}
	return ""
}

func (x *ChampionMastery) GetChampionId() int32 {
	if x != nil {
		return x.ChampionId
	}
	return 0
}

func (x *ChampionMastery) GetChampionLevel() int32 {
	if x != nil {
		return x.ChampionLevel
	}
	return 0
}

func (x *ChampionMastery) GetChampionPoints() int32 {
	if x != nil {
		return x.ChampionPoints
	}
	return 0
}

func (x *ChampionMastery) GetLastPlayTime() int64 {
	if x != nil {
		return x.LastPlayTime
	}
	return 0
}

func (x *ChampionMastery) GetChampionPointsSinceLastLevel() int32 {
	if x != nil {
		return x.ChampionPointsSinceLastLevel
	}
	return 0
}

func (x *ChampionMastery) GetChampionPointsUntilNextLevel() int32 {
	if x != nil {
		return x.ChampionPointsUntilNextLevel
	}
	return 0
}

func (x *ChampionMastery) GetMarkRequiredForNextLevel() int32 {
	if x != nil {
		return x.MarkRequiredForNextLevel
	}
	return 0
}

func (x *ChampionMastery) GetTokensEarned() int32 {
	if x != nil {
		return x.TokensEarned
	}
	return 0
}

func (x *ChampionMastery) GetChampionSeasonMilestone() int32 {
	if x != nil {
		return x.ChampionSeasonMilestone
	}
	return 0
}

func (x *ChampionMastery) GetMilestoneGrades() []string {
	if x != nil {
		return x.MilestoneGrades
	}
	return nil
}

func (x *ChampionMastery) GetNextSeasonMilestone() *NextSeasonMilestone {
	if x != nil {
		return x.NextSeasonMilestone
	}
	return nil
}

type GetChampionMasteriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChampionMasteries []*ChampionMastery `protobuf:"bytes,1,rep,name=championMasteries,proto3" json:"championMasteries,omitempty"`
}

func (x *GetChampionMasteriesResponse) Reset() {
	*x = GetChampionMasteriesResponse{}
	mi := &file_masteries_v1_masteries_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChampionMasteriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChampionMasteriesResponse) ProtoMessage() {}

func (x *GetChampionMasteriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_masteries_v1_masteries_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChampionMasteriesResponse.ProtoReflect.Descriptor instead.
func (*GetChampionMasteriesResponse) Descriptor() ([]byte, []int) {
	return file_masteries_v1_masteries_proto_rawDescGZIP(), []int{3}
}

func (x *GetChampionMasteriesResponse) GetChampionMasteries() []*ChampionMastery {
	if x != nil {
		return x.ChampionMasteries
	}
	return nil
}

type GetChampionMasteriesByChampionRequeset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Puuid      string                   `protobuf:"bytes,1,opt,name=puuid,proto3" json:"puuid,omitempty"`
	ChampionId int32                    `protobuf:"varint,2,opt,name=championId,proto3" json:"championId,omitempty"`
	Region     v1.PlatformRoutingValues `protobuf:"varint,3,opt,name=region,proto3,enum=summoner.PlatformRoutingValues" json:"region,omitempty"`
}

func (x *GetChampionMasteriesByChampionRequeset) Reset() {
	*x = GetChampionMasteriesByChampionRequeset{}
	mi := &file_masteries_v1_masteries_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChampionMasteriesByChampionRequeset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChampionMasteriesByChampionRequeset) ProtoMessage() {}

func (x *GetChampionMasteriesByChampionRequeset) ProtoReflect() protoreflect.Message {
	mi := &file_masteries_v1_masteries_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChampionMasteriesByChampionRequeset.ProtoReflect.Descriptor instead.
func (*GetChampionMasteriesByChampionRequeset) Descriptor() ([]byte, []int) {
	return file_masteries_v1_masteries_proto_rawDescGZIP(), []int{4}
}

func (x *GetChampionMasteriesByChampionRequeset) GetPuuid() string {
	if x != nil {
		return x.Puuid
	}
	return ""
}

func (x *GetChampionMasteriesByChampionRequeset) GetChampionId() int32 {
	if x != nil {
		return x.ChampionId
	}
	return 0
}

func (x *GetChampionMasteriesByChampionRequeset) GetRegion() v1.PlatformRoutingValues {
	if x != nil {
		return x.Region
	}
	return v1.PlatformRoutingValues(0)
}

type GetChampionMasteriesByChampionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChampionMastery *ChampionMastery `protobuf:"bytes,1,opt,name=championMastery,proto3" json:"championMastery,omitempty"`
}

func (x *GetChampionMasteriesByChampionResponse) Reset() {
	*x = GetChampionMasteriesByChampionResponse{}
	mi := &file_masteries_v1_masteries_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChampionMasteriesByChampionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChampionMasteriesByChampionResponse) ProtoMessage() {}

func (x *GetChampionMasteriesByChampionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_masteries_v1_masteries_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChampionMasteriesByChampionResponse.ProtoReflect.Descriptor instead.
func (*GetChampionMasteriesByChampionResponse) Descriptor() ([]byte, []int) {
	return file_masteries_v1_masteries_proto_rawDescGZIP(), []int{5}
}

func (x *GetChampionMasteriesByChampionResponse) GetChampionMastery() *ChampionMastery {
	if x != nil {
		return x.ChampionMastery
	}
	return nil
}

var File_masteries_v1_masteries_proto protoreflect.FileDescriptor

var file_masteries_v1_masteries_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x1a, 0x73, 0x75, 0x6d, 0x6d, 0x6f,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6d,
	0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x75,
	0x6d, 0x6d, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x22, 0xac, 0x02, 0x0a, 0x13, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x62,
	0x6f, 0x6e, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x61, 0x6d,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x73, 0x12, 0x66, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x78,
	0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x45, 0x0a, 0x17,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xd7, 0x04, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e,
	0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x68, 0x61,
	0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x42, 0x0a, 0x1c, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1c, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x42, 0x0a, 0x1c, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x4e, 0x65, 0x78, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1c, 0x63, 0x68, 0x61, 0x6d, 0x70,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x4e, 0x65,
	0x78, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x3a, 0x0a, 0x18, 0x6d, 0x61, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x4e, 0x65, 0x78, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x6d, 0x61, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x4e, 0x65, 0x78, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x45, 0x61, 0x72,
	0x6e, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x17, 0x63, 0x68, 0x61, 0x6d, 0x70,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x69, 0x6c, 0x65,
	0x73, 0x74, 0x6f, 0x6e, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x13, 0x6e,
	0x65, 0x78, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4d,
	0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x52, 0x13, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x22, 0x68, 0x0a,
	0x1c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a,
	0x11, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x79, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x42, 0x79, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6d,
	0x70, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x68,
	0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x6f,
	0x6e, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x22, 0x6e, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e,
	0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x6d, 0x70,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x63,
	0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x79,
	0x52, 0x0f, 0x63, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x79, 0x32, 0x97, 0x02, 0x0a, 0x10, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x79,
	0x50, 0x75, 0x75, 0x69, 0x64, 0x12, 0x27, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x65, 0x74, 0x1a, 0x27,
	0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8f, 0x01, 0x0a, 0x25, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x42, 0x79, 0x50, 0x75, 0x75, 0x69, 0x64, 0x42, 0x79, 0x43, 0x68, 0x61, 0x6d, 0x70,
	0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x65, 0x74, 0x1a, 0x31, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x6d,
	0x61, 0x78, 0x69, 0x73, 0x63, 0x68, 0x6d, 0x61, 0x78, 0x69, 0x2f, 0x6a, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x69, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_masteries_v1_masteries_proto_rawDescOnce sync.Once
	file_masteries_v1_masteries_proto_rawDescData = file_masteries_v1_masteries_proto_rawDesc
)

func file_masteries_v1_masteries_proto_rawDescGZIP() []byte {
	file_masteries_v1_masteries_proto_rawDescOnce.Do(func() {
		file_masteries_v1_masteries_proto_rawDescData = protoimpl.X.CompressGZIP(file_masteries_v1_masteries_proto_rawDescData)
	})
	return file_masteries_v1_masteries_proto_rawDescData
}

var file_masteries_v1_masteries_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_masteries_v1_masteries_proto_goTypes = []any{
	(*GetChampionMasteriesRequeset)(nil),           // 0: masteries.GetChampionMasteriesRequeset
	(*NextSeasonMilestone)(nil),                    // 1: masteries.NextSeasonMilestone
	(*ChampionMastery)(nil),                        // 2: masteries.ChampionMastery
	(*GetChampionMasteriesResponse)(nil),           // 3: masteries.GetChampionMasteriesResponse
	(*GetChampionMasteriesByChampionRequeset)(nil), // 4: masteries.GetChampionMasteriesByChampionRequeset
	(*GetChampionMasteriesByChampionResponse)(nil), // 5: masteries.GetChampionMasteriesByChampionResponse
	nil,                           // 6: masteries.NextSeasonMilestone.RequireGradeCountsEntry
	(v1.PlatformRoutingValues)(0), // 7: summoner.PlatformRoutingValues
}
var file_masteries_v1_masteries_proto_depIdxs = []int32{
	7, // 0: masteries.GetChampionMasteriesRequeset.region:type_name -> summoner.PlatformRoutingValues
	6, // 1: masteries.NextSeasonMilestone.requireGradeCounts:type_name -> masteries.NextSeasonMilestone.RequireGradeCountsEntry
	1, // 2: masteries.ChampionMastery.nextSeasonMilestone:type_name -> masteries.NextSeasonMilestone
	2, // 3: masteries.GetChampionMasteriesResponse.championMasteries:type_name -> masteries.ChampionMastery
	7, // 4: masteries.GetChampionMasteriesByChampionRequeset.region:type_name -> summoner.PlatformRoutingValues
	2, // 5: masteries.GetChampionMasteriesByChampionResponse.championMastery:type_name -> masteries.ChampionMastery
	0, // 6: masteries.MasteriesService.GetChampionMasteriesByPuuid:input_type -> masteries.GetChampionMasteriesRequeset
	4, // 7: masteries.MasteriesService.GetChampionMasteriesByPuuidByChampion:input_type -> masteries.GetChampionMasteriesByChampionRequeset
	3, // 8: masteries.MasteriesService.GetChampionMasteriesByPuuid:output_type -> masteries.GetChampionMasteriesResponse
	5, // 9: masteries.MasteriesService.GetChampionMasteriesByPuuidByChampion:output_type -> masteries.GetChampionMasteriesByChampionResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_masteries_v1_masteries_proto_init() }
func file_masteries_v1_masteries_proto_init() {
	if File_masteries_v1_masteries_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_masteries_v1_masteries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_masteries_v1_masteries_proto_goTypes,
		DependencyIndexes: file_masteries_v1_masteries_proto_depIdxs,
		MessageInfos:      file_masteries_v1_masteries_proto_msgTypes,
	}.Build()
	File_masteries_v1_masteries_proto = out.File
	file_masteries_v1_masteries_proto_rawDesc = nil
	file_masteries_v1_masteries_proto_goTypes = nil
	file_masteries_v1_masteries_proto_depIdxs = nil
}
