// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: jdpedrie/llmrpg/v1/llmrpg.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlayRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GameId        string                 `protobuf:"bytes,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	Choice        string                 `protobuf:"bytes,2,opt,name=choice" json:"choice,omitempty"`
	Outcome       string                 `protobuf:"bytes,3,opt,name=outcome" json:"outcome,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlayRequest) Reset() {
	*x = PlayRequest{}
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRequest) ProtoMessage() {}

func (x *PlayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRequest.ProtoReflect.Descriptor instead.
func (*PlayRequest) Descriptor() ([]byte, []int) {
	return file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescGZIP(), []int{0}
}

func (x *PlayRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *PlayRequest) GetChoice() string {
	if x != nil {
		return x.Choice
	}
	return ""
}

func (x *PlayRequest) GetOutcome() string {
	if x != nil {
		return x.Outcome
	}
	return ""
}

type PlayResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Resp:
	//
	//	*PlayResponse_Game
	//	*PlayResponse_Message
	Resp          isPlayResponse_Resp `protobuf_oneof:"resp"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlayResponse) Reset() {
	*x = PlayResponse{}
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse) ProtoMessage() {}

func (x *PlayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse.ProtoReflect.Descriptor instead.
func (*PlayResponse) Descriptor() ([]byte, []int) {
	return file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescGZIP(), []int{1}
}

func (x *PlayResponse) GetResp() isPlayResponse_Resp {
	if x != nil {
		return x.Resp
	}
	return nil
}

func (x *PlayResponse) GetGame() *Game {
	if x != nil {
		if x, ok := x.Resp.(*PlayResponse_Game); ok {
			return x.Game
		}
	}
	return nil
}

func (x *PlayResponse) GetMessage() string {
	if x != nil {
		if x, ok := x.Resp.(*PlayResponse_Message); ok {
			return x.Message
		}
	}
	return ""
}

type isPlayResponse_Resp interface {
	isPlayResponse_Resp()
}

type PlayResponse_Game struct {
	Game *Game `protobuf:"bytes,1,opt,name=game,oneof"`
}

type PlayResponse_Message struct {
	Message string `protobuf:"bytes,2,opt,name=message,oneof"`
}

func (*PlayResponse_Game) isPlayResponse_Resp() {}

func (*PlayResponse_Message) isPlayResponse_Resp() {}

type Game struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Id                   string                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description          string                 `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Characters           []*Character           `protobuf:"bytes,4,rep,name=characters" json:"characters,omitempty"`
	StartingMessage      string                 `protobuf:"bytes,5,opt,name=starting_message,json=startingMessage" json:"starting_message,omitempty"`
	Scenario             string                 `protobuf:"bytes,6,opt,name=scenario" json:"scenario,omitempty"`
	Objectives           string                 `protobuf:"bytes,7,opt,name=objectives" json:"objectives,omitempty"`
	Inventory            []*InventoryItem       `protobuf:"bytes,8,rep,name=inventory" json:"inventory,omitempty"`
	Skills               []string               `protobuf:"bytes,9,rep,name=skills" json:"skills,omitempty"`
	Characteristics      []string               `protobuf:"bytes,10,rep,name=characteristics" json:"characteristics,omitempty"`
	Relationship         []string               `protobuf:"bytes,11,rep,name=relationship" json:"relationship,omitempty"`
	IsTemplate           bool                   `protobuf:"varint,12,opt,name=is_template,json=isTemplate" json:"is_template,omitempty"`
	IsRunning            bool                   `protobuf:"varint,13,opt,name=is_running,json=isRunning" json:"is_running,omitempty"`
	PlaythroughStartTime *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=playthrough_start_time,json=playthroughStartTime" json:"playthrough_start_time,omitempty"`
	PlaythroughEndTime   *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=playthrough_end_time,json=playthroughEndTime" json:"playthrough_end_time,omitempty"`
	LastActivityTime     *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=last_activity_time,json=lastActivityTime" json:"last_activity_time,omitempty"`
	CreateTime           *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Game) Reset() {
	*x = Game{}
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescGZIP(), []int{2}
}

func (x *Game) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Game) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Game) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Game) GetCharacters() []*Character {
	if x != nil {
		return x.Characters
	}
	return nil
}

func (x *Game) GetStartingMessage() string {
	if x != nil {
		return x.StartingMessage
	}
	return ""
}

func (x *Game) GetScenario() string {
	if x != nil {
		return x.Scenario
	}
	return ""
}

func (x *Game) GetObjectives() string {
	if x != nil {
		return x.Objectives
	}
	return ""
}

func (x *Game) GetInventory() []*InventoryItem {
	if x != nil {
		return x.Inventory
	}
	return nil
}

func (x *Game) GetSkills() []string {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *Game) GetCharacteristics() []string {
	if x != nil {
		return x.Characteristics
	}
	return nil
}

func (x *Game) GetRelationship() []string {
	if x != nil {
		return x.Relationship
	}
	return nil
}

func (x *Game) GetIsTemplate() bool {
	if x != nil {
		return x.IsTemplate
	}
	return false
}

func (x *Game) GetIsRunning() bool {
	if x != nil {
		return x.IsRunning
	}
	return false
}

func (x *Game) GetPlaythroughStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PlaythroughStartTime
	}
	return nil
}

func (x *Game) GetPlaythroughEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PlaythroughEndTime
	}
	return nil
}

func (x *Game) GetLastActivityTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastActivityTime
	}
	return nil
}

func (x *Game) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type Character struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description     string                 `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Context         []string               `protobuf:"bytes,4,rep,name=context" json:"context,omitempty"`
	Active          bool                   `protobuf:"varint,5,opt,name=active" json:"active,omitempty"`
	MainCharacter   bool                   `protobuf:"varint,6,opt,name=main_character,json=mainCharacter" json:"main_character,omitempty"`
	Skills          []*CharacterAttribute  `protobuf:"bytes,7,rep,name=skills" json:"skills,omitempty"`
	Characteristics []*CharacterAttribute  `protobuf:"bytes,8,rep,name=characteristics" json:"characteristics,omitempty"`
	Relationship    []*CharacterAttribute  `protobuf:"bytes,9,rep,name=relationship" json:"relationship,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Character) Reset() {
	*x = Character{}
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescGZIP(), []int{3}
}

func (x *Character) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Character) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Character) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Character) GetContext() []string {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *Character) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Character) GetMainCharacter() bool {
	if x != nil {
		return x.MainCharacter
	}
	return false
}

func (x *Character) GetSkills() []*CharacterAttribute {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *Character) GetCharacteristics() []*CharacterAttribute {
	if x != nil {
		return x.Characteristics
	}
	return nil
}

func (x *Character) GetRelationship() []*CharacterAttribute {
	if x != nil {
		return x.Relationship
	}
	return nil
}

type CharacterAttribute struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value         int32                  `protobuf:"varint,3,opt,name=value" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CharacterAttribute) Reset() {
	*x = CharacterAttribute{}
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CharacterAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharacterAttribute) ProtoMessage() {}

func (x *CharacterAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharacterAttribute.ProtoReflect.Descriptor instead.
func (*CharacterAttribute) Descriptor() ([]byte, []int) {
	return file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescGZIP(), []int{4}
}

func (x *CharacterAttribute) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CharacterAttribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CharacterAttribute) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type History struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	GameId        string                 `protobuf:"bytes,2,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	Text          string                 `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
	Choice        string                 `protobuf:"bytes,4,opt,name=choice" json:"choice,omitempty"`
	Outcome       string                 `protobuf:"bytes,5,opt,name=outcome" json:"outcome,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *History) Reset() {
	*x = History{}
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *History) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*History) ProtoMessage() {}

func (x *History) ProtoReflect() protoreflect.Message {
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use History.ProtoReflect.Descriptor instead.
func (*History) Descriptor() ([]byte, []int) {
	return file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescGZIP(), []int{5}
}

func (x *History) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *History) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *History) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *History) GetChoice() string {
	if x != nil {
		return x.Choice
	}
	return ""
}

func (x *History) GetOutcome() string {
	if x != nil {
		return x.Outcome
	}
	return ""
}

type InventoryItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Active        bool                   `protobuf:"varint,4,opt,name=active" json:"active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InventoryItem) Reset() {
	*x = InventoryItem{}
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InventoryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryItem) ProtoMessage() {}

func (x *InventoryItem) ProtoReflect() protoreflect.Message {
	mi := &file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryItem.ProtoReflect.Descriptor instead.
func (*InventoryItem) Descriptor() ([]byte, []int) {
	return file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescGZIP(), []int{6}
}

func (x *InventoryItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InventoryItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InventoryItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InventoryItem) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

var File_jdpedrie_llmrpg_v1_llmrpg_proto protoreflect.FileDescriptor

var file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x6a, 0x64, 0x70, 0x65, 0x64, 0x72, 0x69, 0x65, 0x2f, 0x6c, 0x6c, 0x6d, 0x72, 0x70,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6c, 0x6d, 0x72, 0x70, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65,
	0x22, 0x52, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04,
	0x72, 0x65, 0x73, 0x70, 0x22, 0xe0, 0x05, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x12, 0x50, 0x0a, 0x16, 0x70, 0x6c, 0x61, 0x79, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x70, 0x6c,
	0x61, 0x79, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x14, 0x70, 0x6c, 0x61, 0x79, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67,
	0x68, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x70, 0x6c,
	0x61, 0x79, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x48, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xd8, 0x02, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x06, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x22, 0x4e, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x78, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0x6d, 0x0a, 0x0d,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x32, 0x3e, 0x0a, 0x0d, 0x4c,
	0x4c, 0x4d, 0x52, 0x50, 0x47, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x04,
	0x50, 0x6c, 0x61, 0x79, 0x12, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3d, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x64, 0x70, 0x65, 0x64, 0x72,
	0x69, 0x65, 0x2f, 0x6c, 0x6c, 0x6d, 0x72, 0x70, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6a, 0x64, 0x70, 0x65, 0x64, 0x72, 0x69, 0x65, 0x2f, 0x6c, 0x6c, 0x6d, 0x72,
	0x70, 0x67, 0x2f, 0x76, 0x31, 0x92, 0x03, 0x02, 0x08, 0x02, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var (
	file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescOnce sync.Once
	file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescData []byte
)

func file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescGZIP() []byte {
	file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescOnce.Do(func() {
		file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDesc), len(file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDesc)))
	})
	return file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDescData
}

var file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_jdpedrie_llmrpg_v1_llmrpg_proto_goTypes = []any{
	(*PlayRequest)(nil),           // 0: v1.PlayRequest
	(*PlayResponse)(nil),          // 1: v1.PlayResponse
	(*Game)(nil),                  // 2: v1.Game
	(*Character)(nil),             // 3: v1.Character
	(*CharacterAttribute)(nil),    // 4: v1.CharacterAttribute
	(*History)(nil),               // 5: v1.History
	(*InventoryItem)(nil),         // 6: v1.InventoryItem
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_jdpedrie_llmrpg_v1_llmrpg_proto_depIdxs = []int32{
	2,  // 0: v1.PlayResponse.game:type_name -> v1.Game
	3,  // 1: v1.Game.characters:type_name -> v1.Character
	6,  // 2: v1.Game.inventory:type_name -> v1.InventoryItem
	7,  // 3: v1.Game.playthrough_start_time:type_name -> google.protobuf.Timestamp
	7,  // 4: v1.Game.playthrough_end_time:type_name -> google.protobuf.Timestamp
	7,  // 5: v1.Game.last_activity_time:type_name -> google.protobuf.Timestamp
	7,  // 6: v1.Game.create_time:type_name -> google.protobuf.Timestamp
	4,  // 7: v1.Character.skills:type_name -> v1.CharacterAttribute
	4,  // 8: v1.Character.characteristics:type_name -> v1.CharacterAttribute
	4,  // 9: v1.Character.relationship:type_name -> v1.CharacterAttribute
	0,  // 10: v1.LLMRPGService.Play:input_type -> v1.PlayRequest
	1,  // 11: v1.LLMRPGService.Play:output_type -> v1.PlayResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_jdpedrie_llmrpg_v1_llmrpg_proto_init() }
func file_jdpedrie_llmrpg_v1_llmrpg_proto_init() {
	if File_jdpedrie_llmrpg_v1_llmrpg_proto != nil {
		return
	}
	file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes[1].OneofWrappers = []any{
		(*PlayResponse_Game)(nil),
		(*PlayResponse_Message)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDesc), len(file_jdpedrie_llmrpg_v1_llmrpg_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jdpedrie_llmrpg_v1_llmrpg_proto_goTypes,
		DependencyIndexes: file_jdpedrie_llmrpg_v1_llmrpg_proto_depIdxs,
		MessageInfos:      file_jdpedrie_llmrpg_v1_llmrpg_proto_msgTypes,
	}.Build()
	File_jdpedrie_llmrpg_v1_llmrpg_proto = out.File
	file_jdpedrie_llmrpg_v1_llmrpg_proto_goTypes = nil
	file_jdpedrie_llmrpg_v1_llmrpg_proto_depIdxs = nil
}
