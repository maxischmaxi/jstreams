// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: items/v1/items.proto

package itemsv1

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

type GetBaseItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatchVersion string `protobuf:"bytes,1,opt,name=patchVersion,proto3" json:"patchVersion,omitempty"`
}

func (x *GetBaseItemsRequest) Reset() {
	*x = GetBaseItemsRequest{}
	mi := &file_items_v1_items_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBaseItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBaseItemsRequest) ProtoMessage() {}

func (x *GetBaseItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_items_v1_items_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBaseItemsRequest.ProtoReflect.Descriptor instead.
func (*GetBaseItemsRequest) Descriptor() ([]byte, []int) {
	return file_items_v1_items_proto_rawDescGZIP(), []int{0}
}

func (x *GetBaseItemsRequest) GetPatchVersion() string {
	if x != nil {
		return x.PatchVersion
	}
	return ""
}

type GetBaseItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string           `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Version string           `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Basic   *Item            `protobuf:"bytes,3,opt,name=basic,proto3" json:"basic,omitempty"`
	Data    map[string]*Item `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetBaseItemsResponse) Reset() {
	*x = GetBaseItemsResponse{}
	mi := &file_items_v1_items_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBaseItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBaseItemsResponse) ProtoMessage() {}

func (x *GetBaseItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_items_v1_items_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBaseItemsResponse.ProtoReflect.Descriptor instead.
func (*GetBaseItemsResponse) Descriptor() ([]byte, []int) {
	return file_items_v1_items_proto_rawDescGZIP(), []int{1}
}

func (x *GetBaseItemsResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetBaseItemsResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetBaseItemsResponse) GetBasic() *Item {
	if x != nil {
		return x.Basic
	}
	return nil
}

func (x *GetBaseItemsResponse) GetData() map[string]*Item {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetItemImageByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId       int32  `protobuf:"varint,1,opt,name=itemId,proto3" json:"itemId,omitempty"`
	PatchVersion string `protobuf:"bytes,2,opt,name=patchVersion,proto3" json:"patchVersion,omitempty"`
}

func (x *GetItemImageByIdRequest) Reset() {
	*x = GetItemImageByIdRequest{}
	mi := &file_items_v1_items_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemImageByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemImageByIdRequest) ProtoMessage() {}

func (x *GetItemImageByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_items_v1_items_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemImageByIdRequest.ProtoReflect.Descriptor instead.
func (*GetItemImageByIdRequest) Descriptor() ([]byte, []int) {
	return file_items_v1_items_proto_rawDescGZIP(), []int{2}
}

func (x *GetItemImageByIdRequest) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *GetItemImageByIdRequest) GetPatchVersion() string {
	if x != nil {
		return x.PatchVersion
	}
	return ""
}

type GetItemImageByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetItemImageByIdResponse) Reset() {
	*x = GetItemImageByIdResponse{}
	mi := &file_items_v1_items_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemImageByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemImageByIdResponse) ProtoMessage() {}

func (x *GetItemImageByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_items_v1_items_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemImageByIdResponse.ProtoReflect.Descriptor instead.
func (*GetItemImageByIdResponse) Descriptor() ([]byte, []int) {
	return file_items_v1_items_proto_rawDescGZIP(), []int{3}
}

func (x *GetItemImageByIdResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetItemInformationByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatchVersion string `protobuf:"bytes,1,opt,name=patchVersion,proto3" json:"patchVersion,omitempty"`
	ItemId       int32  `protobuf:"varint,2,opt,name=itemId,proto3" json:"itemId,omitempty"`
}

func (x *GetItemInformationByIdRequest) Reset() {
	*x = GetItemInformationByIdRequest{}
	mi := &file_items_v1_items_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemInformationByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemInformationByIdRequest) ProtoMessage() {}

func (x *GetItemInformationByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_items_v1_items_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemInformationByIdRequest.ProtoReflect.Descriptor instead.
func (*GetItemInformationByIdRequest) Descriptor() ([]byte, []int) {
	return file_items_v1_items_proto_rawDescGZIP(), []int{4}
}

func (x *GetItemInformationByIdRequest) GetPatchVersion() string {
	if x != nil {
		return x.PatchVersion
	}
	return ""
}

func (x *GetItemInformationByIdRequest) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

type GetItemInformationByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetItemInformationByIdResponse) Reset() {
	*x = GetItemInformationByIdResponse{}
	mi := &file_items_v1_items_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemInformationByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemInformationByIdResponse) ProtoMessage() {}

func (x *GetItemInformationByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_items_v1_items_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemInformationByIdResponse.ProtoReflect.Descriptor instead.
func (*GetItemInformationByIdResponse) Descriptor() ([]byte, []int) {
	return file_items_v1_items_proto_rawDescGZIP(), []int{5}
}

func (x *GetItemInformationByIdResponse) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rune             *Rune              `protobuf:"bytes,2,opt,name=rune,proto3" json:"rune,omitempty"`
	Gold             *Gold              `protobuf:"bytes,3,opt,name=gold,proto3" json:"gold,omitempty"`
	Group            string             `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	Description      string             `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Colloq           string             `protobuf:"bytes,6,opt,name=colloq,proto3" json:"colloq,omitempty"`
	Plaintext        string             `protobuf:"bytes,7,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	Consumed         bool               `protobuf:"varint,8,opt,name=consumed,proto3" json:"consumed,omitempty"`
	Stacks           int32              `protobuf:"varint,9,opt,name=stacks,proto3" json:"stacks,omitempty"`
	Depth            int32              `protobuf:"varint,10,opt,name=depth,proto3" json:"depth,omitempty"`
	ConsumeOnFull    bool               `protobuf:"varint,11,opt,name=consumeOnFull,proto3" json:"consumeOnFull,omitempty"`
	SpecialRecipe    int32              `protobuf:"varint,12,opt,name=specialRecipe,proto3" json:"specialRecipe,omitempty"`
	InStore          bool               `protobuf:"varint,13,opt,name=inStore,proto3" json:"inStore,omitempty"`
	HideFromAll      bool               `protobuf:"varint,14,opt,name=hideFromAll,proto3" json:"hideFromAll,omitempty"`
	RequiredChampion string             `protobuf:"bytes,15,opt,name=requiredChampion,proto3" json:"requiredChampion,omitempty"`
	RequiredAlly     string             `protobuf:"bytes,16,opt,name=requiredAlly,proto3" json:"requiredAlly,omitempty"`
	Stats            map[string]float32 `protobuf:"bytes,17,rep,name=stats,proto3" json:"stats,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	Maps             map[string]bool    `protobuf:"bytes,18,rep,name=maps,proto3" json:"maps,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Tags             []string           `protobuf:"bytes,19,rep,name=tags,proto3" json:"tags,omitempty"`
	Effect           map[string]string  `protobuf:"bytes,20,rep,name=effect,proto3" json:"effect,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_items_v1_items_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_items_v1_items_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_items_v1_items_proto_rawDescGZIP(), []int{6}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetRune() *Rune {
	if x != nil {
		return x.Rune
	}
	return nil
}

func (x *Item) GetGold() *Gold {
	if x != nil {
		return x.Gold
	}
	return nil
}

func (x *Item) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Item) GetColloq() string {
	if x != nil {
		return x.Colloq
	}
	return ""
}

func (x *Item) GetPlaintext() string {
	if x != nil {
		return x.Plaintext
	}
	return ""
}

func (x *Item) GetConsumed() bool {
	if x != nil {
		return x.Consumed
	}
	return false
}

func (x *Item) GetStacks() int32 {
	if x != nil {
		return x.Stacks
	}
	return 0
}

func (x *Item) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *Item) GetConsumeOnFull() bool {
	if x != nil {
		return x.ConsumeOnFull
	}
	return false
}

func (x *Item) GetSpecialRecipe() int32 {
	if x != nil {
		return x.SpecialRecipe
	}
	return 0
}

func (x *Item) GetInStore() bool {
	if x != nil {
		return x.InStore
	}
	return false
}

func (x *Item) GetHideFromAll() bool {
	if x != nil {
		return x.HideFromAll
	}
	return false
}

func (x *Item) GetRequiredChampion() string {
	if x != nil {
		return x.RequiredChampion
	}
	return ""
}

func (x *Item) GetRequiredAlly() string {
	if x != nil {
		return x.RequiredAlly
	}
	return ""
}

func (x *Item) GetStats() map[string]float32 {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Item) GetMaps() map[string]bool {
	if x != nil {
		return x.Maps
	}
	return nil
}

func (x *Item) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Item) GetEffect() map[string]string {
	if x != nil {
		return x.Effect
	}
	return nil
}

type Gold struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base        int32 `protobuf:"varint,1,opt,name=base,proto3" json:"base,omitempty"`
	Total       int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Sell        int32 `protobuf:"varint,3,opt,name=sell,proto3" json:"sell,omitempty"`
	Purchasable bool  `protobuf:"varint,4,opt,name=purchasable,proto3" json:"purchasable,omitempty"`
}

func (x *Gold) Reset() {
	*x = Gold{}
	mi := &file_items_v1_items_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Gold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gold) ProtoMessage() {}

func (x *Gold) ProtoReflect() protoreflect.Message {
	mi := &file_items_v1_items_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gold.ProtoReflect.Descriptor instead.
func (*Gold) Descriptor() ([]byte, []int) {
	return file_items_v1_items_proto_rawDescGZIP(), []int{7}
}

func (x *Gold) GetBase() int32 {
	if x != nil {
		return x.Base
	}
	return 0
}

func (x *Gold) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Gold) GetSell() int32 {
	if x != nil {
		return x.Sell
	}
	return 0
}

func (x *Gold) GetPurchasable() bool {
	if x != nil {
		return x.Purchasable
	}
	return false
}

type Rune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isrune bool   `protobuf:"varint,1,opt,name=isrune,proto3" json:"isrune,omitempty"`
	Tier   int32  `protobuf:"varint,2,opt,name=tier,proto3" json:"tier,omitempty"`
	Type   string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Rune) Reset() {
	*x = Rune{}
	mi := &file_items_v1_items_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Rune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rune) ProtoMessage() {}

func (x *Rune) ProtoReflect() protoreflect.Message {
	mi := &file_items_v1_items_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rune.ProtoReflect.Descriptor instead.
func (*Rune) Descriptor() ([]byte, []int) {
	return file_items_v1_items_proto_rawDescGZIP(), []int{8}
}

func (x *Rune) GetIsrune() bool {
	if x != nil {
		return x.Isrune
	}
	return false
}

func (x *Rune) GetTier() int32 {
	if x != nil {
		return x.Tier
	}
	return 0
}

func (x *Rune) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_items_v1_items_proto protoreflect.FileDescriptor

var file_items_v1_items_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x39, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x74, 0x63, 0x68, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xe8, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x44, 0x0a,
	0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x55, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x74, 0x63, 0x68, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x5b, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x61, 0x74, 0x63, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0xb8, 0x06, 0x0a, 0x04, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x65,
	0x52, 0x04, 0x72, 0x75, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x47, 0x6f, 0x6c,
	0x64, 0x52, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x6c, 0x6f, 0x71, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x6c, 0x6c, 0x6f, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70,
	0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x12,
	0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x4f, 0x6e, 0x46, 0x75, 0x6c, 0x6c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x4f,
	0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x6e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x69, 0x64, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x41, 0x6c, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x69, 0x64, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x41, 0x6c, 0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6d, 0x70, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6d, 0x70,
	0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x41,
	0x6c, 0x6c, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x41, 0x6c, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x61, 0x70, 0x73, 0x18, 0x12, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x2e, 0x4d, 0x61, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x61, 0x70, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x14,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x1a, 0x38, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x37, 0x0a, 0x09, 0x4d, 0x61, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x45, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x66, 0x0a, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x6c, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x65, 0x6c, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x46, 0x0a, 0x04, 0x52,
	0x75, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x72, 0x75, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x72, 0x75, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x32, 0x99, 0x02, 0x0a, 0x0c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x1a, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x27, 0x5a, 0x25, 0x6d, 0x61, 0x78, 0x69, 0x73, 0x63, 0x68, 0x6d, 0x61, 0x78, 0x69, 0x2f, 0x6a,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_items_v1_items_proto_rawDescOnce sync.Once
	file_items_v1_items_proto_rawDescData = file_items_v1_items_proto_rawDesc
)

func file_items_v1_items_proto_rawDescGZIP() []byte {
	file_items_v1_items_proto_rawDescOnce.Do(func() {
		file_items_v1_items_proto_rawDescData = protoimpl.X.CompressGZIP(file_items_v1_items_proto_rawDescData)
	})
	return file_items_v1_items_proto_rawDescData
}

var file_items_v1_items_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_items_v1_items_proto_goTypes = []any{
	(*GetBaseItemsRequest)(nil),            // 0: items.GetBaseItemsRequest
	(*GetBaseItemsResponse)(nil),           // 1: items.GetBaseItemsResponse
	(*GetItemImageByIdRequest)(nil),        // 2: items.GetItemImageByIdRequest
	(*GetItemImageByIdResponse)(nil),       // 3: items.GetItemImageByIdResponse
	(*GetItemInformationByIdRequest)(nil),  // 4: items.GetItemInformationByIdRequest
	(*GetItemInformationByIdResponse)(nil), // 5: items.GetItemInformationByIdResponse
	(*Item)(nil),                           // 6: items.Item
	(*Gold)(nil),                           // 7: items.Gold
	(*Rune)(nil),                           // 8: items.Rune
	nil,                                    // 9: items.GetBaseItemsResponse.DataEntry
	nil,                                    // 10: items.Item.StatsEntry
	nil,                                    // 11: items.Item.MapsEntry
	nil,                                    // 12: items.Item.EffectEntry
}
var file_items_v1_items_proto_depIdxs = []int32{
	6,  // 0: items.GetBaseItemsResponse.basic:type_name -> items.Item
	9,  // 1: items.GetBaseItemsResponse.data:type_name -> items.GetBaseItemsResponse.DataEntry
	6,  // 2: items.GetItemInformationByIdResponse.item:type_name -> items.Item
	8,  // 3: items.Item.rune:type_name -> items.Rune
	7,  // 4: items.Item.gold:type_name -> items.Gold
	10, // 5: items.Item.stats:type_name -> items.Item.StatsEntry
	11, // 6: items.Item.maps:type_name -> items.Item.MapsEntry
	12, // 7: items.Item.effect:type_name -> items.Item.EffectEntry
	6,  // 8: items.GetBaseItemsResponse.DataEntry.value:type_name -> items.Item
	2,  // 9: items.ItemsService.GetItemImageById:input_type -> items.GetItemImageByIdRequest
	4,  // 10: items.ItemsService.GetItemInformationById:input_type -> items.GetItemInformationByIdRequest
	0,  // 11: items.ItemsService.GetBaseItems:input_type -> items.GetBaseItemsRequest
	3,  // 12: items.ItemsService.GetItemImageById:output_type -> items.GetItemImageByIdResponse
	5,  // 13: items.ItemsService.GetItemInformationById:output_type -> items.GetItemInformationByIdResponse
	1,  // 14: items.ItemsService.GetBaseItems:output_type -> items.GetBaseItemsResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_items_v1_items_proto_init() }
func file_items_v1_items_proto_init() {
	if File_items_v1_items_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_items_v1_items_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_items_v1_items_proto_goTypes,
		DependencyIndexes: file_items_v1_items_proto_depIdxs,
		MessageInfos:      file_items_v1_items_proto_msgTypes,
	}.Build()
	File_items_v1_items_proto = out.File
	file_items_v1_items_proto_rawDesc = nil
	file_items_v1_items_proto_goTypes = nil
	file_items_v1_items_proto_depIdxs = nil
}
