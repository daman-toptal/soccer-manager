// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: player/player.proto

package player

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	golang "protobuf-v1/golang"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlayerType int32

const (
	PlayerType_PT_UNSPECIFIED PlayerType = 0
	PlayerType_PT_GOAL_KEEPER PlayerType = 1
	PlayerType_PT_DEFENDER    PlayerType = 2
	PlayerType_PT_MID_FIELDER PlayerType = 3
	PlayerType_PT_ATTACKER    PlayerType = 4
)

// Enum value maps for PlayerType.
var (
	PlayerType_name = map[int32]string{
		0: "PT_UNSPECIFIED",
		1: "PT_GOAL_KEEPER",
		2: "PT_DEFENDER",
		3: "PT_MID_FIELDER",
		4: "PT_ATTACKER",
	}
	PlayerType_value = map[string]int32{
		"PT_UNSPECIFIED": 0,
		"PT_GOAL_KEEPER": 1,
		"PT_DEFENDER":    2,
		"PT_MID_FIELDER": 3,
		"PT_ATTACKER":    4,
	}
)

func (x PlayerType) Enum() *PlayerType {
	p := new(PlayerType)
	*p = x
	return p
}

func (x PlayerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerType) Descriptor() protoreflect.EnumDescriptor {
	return file_player_player_proto_enumTypes[0].Descriptor()
}

func (PlayerType) Type() protoreflect.EnumType {
	return &file_player_player_proto_enumTypes[0]
}

func (x PlayerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerType.Descriptor instead.
func (PlayerType) EnumDescriptor() ([]byte, []int) {
	return file_player_player_proto_rawDescGZIP(), []int{0}
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age       int32                  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Type      PlayerType             `protobuf:"varint,5,opt,name=type,proto3,enum=protobuf.player.PlayerType" json:"type,omitempty"`
	Country   string                 `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	TeamId    string                 `protobuf:"bytes,7,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Value     int64                  `protobuf:"varint,8,opt,name=value,proto3" json:"value,omitempty"`
	IsListed  bool                   `protobuf:"varint,9,opt,name=is_listed,json=isListed,proto3" json:"is_listed,omitempty"`
	AskValue  *wrapperspb.Int64Value `protobuf:"bytes,10,opt,name=ask_value,json=askValue,proto3" json:"ask_value,omitempty"`
	Currency  golang.Currency        `protobuf:"varint,11,opt,name=currency,proto3,enum=protobuf.Currency" json:"currency,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_player_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_player_player_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Player) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Player) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Player) GetType() PlayerType {
	if x != nil {
		return x.Type
	}
	return PlayerType_PT_UNSPECIFIED
}

func (x *Player) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Player) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *Player) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Player) GetIsListed() bool {
	if x != nil {
		return x.IsListed
	}
	return false
}

func (x *Player) GetAskValue() *wrapperspb.Int64Value {
	if x != nil {
		return x.AskValue
	}
	return nil
}

func (x *Player) GetCurrency() golang.Currency {
	if x != nil {
		return x.Currency
	}
	return golang.Currency(0)
}

type Players struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total   int32     `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Players []*Player `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *Players) Reset() {
	*x = Players{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Players) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Players) ProtoMessage() {}

func (x *Players) ProtoReflect() protoreflect.Message {
	mi := &file_player_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Players.ProtoReflect.Descriptor instead.
func (*Players) Descriptor() ([]byte, []int) {
	return file_player_player_proto_rawDescGZIP(), []int{1}
}

func (x *Players) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Players) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_player_player_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetByTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *GetByTeamRequest) Reset() {
	*x = GetByTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByTeamRequest) ProtoMessage() {}

func (x *GetByTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByTeamRequest.ProtoReflect.Descriptor instead.
func (*GetByTeamRequest) Descriptor() ([]byte, []int) {
	return file_player_player_proto_rawDescGZIP(), []int{3}
}

func (x *GetByTeamRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

type GetListedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetListedRequest) Reset() {
	*x = GetListedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListedRequest) ProtoMessage() {}

func (x *GetListedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListedRequest.ProtoReflect.Descriptor instead.
func (*GetListedRequest) Descriptor() ([]byte, []int) {
	return file_player_player_proto_rawDescGZIP(), []int{4}
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Country   string                 `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	IsListed  *wrapperspb.BoolValue  `protobuf:"bytes,5,opt,name=is_listed,json=isListed,proto3" json:"is_listed,omitempty"`
	AskValue  *wrapperspb.Int64Value `protobuf:"bytes,6,opt,name=ask_value,json=askValue,proto3" json:"ask_value,omitempty"`
	Value     *wrapperspb.Int64Value `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_player_player_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UpdateRequest) GetIsListed() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsListed
	}
	return nil
}

func (x *UpdateRequest) GetAskValue() *wrapperspb.Int64Value {
	if x != nil {
		return x.AskValue
	}
	return nil
}

func (x *UpdateRequest) GetValue() *wrapperspb.Int64Value {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_player_player_proto protoreflect.FileDescriptor

var file_player_player_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x02, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12,
	0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x2e, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x22, 0x52, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x9b, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x69, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x2a, 0x6a, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x0e, 0x50, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x54, 0x5f, 0x47, 0x4f, 0x41, 0x4c, 0x5f, 0x4b,
	0x45, 0x45, 0x50, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x54, 0x5f, 0x44, 0x45,
	0x46, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x54, 0x5f, 0x4d,
	0x49, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b,
	0x50, 0x54, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x45, 0x52, 0x10, 0x04, 0x32, 0xa3, 0x02,
	0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x42, 0x1b, 0x5a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d,
	0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_player_player_proto_rawDescOnce sync.Once
	file_player_player_proto_rawDescData = file_player_player_proto_rawDesc
)

func file_player_player_proto_rawDescGZIP() []byte {
	file_player_player_proto_rawDescOnce.Do(func() {
		file_player_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_player_proto_rawDescData)
	})
	return file_player_player_proto_rawDescData
}

var file_player_player_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_player_player_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_player_player_proto_goTypes = []interface{}{
	(PlayerType)(0),               // 0: protobuf.player.PlayerType
	(*Player)(nil),                // 1: protobuf.player.Player
	(*Players)(nil),               // 2: protobuf.player.Players
	(*GetRequest)(nil),            // 3: protobuf.player.GetRequest
	(*GetByTeamRequest)(nil),      // 4: protobuf.player.GetByTeamRequest
	(*GetListedRequest)(nil),      // 5: protobuf.player.GetListedRequest
	(*UpdateRequest)(nil),         // 6: protobuf.player.UpdateRequest
	(*wrapperspb.Int64Value)(nil), // 7: google.protobuf.Int64Value
	(golang.Currency)(0),          // 8: protobuf.Currency
	(*wrapperspb.BoolValue)(nil),  // 9: google.protobuf.BoolValue
}
var file_player_player_proto_depIdxs = []int32{
	0,  // 0: protobuf.player.Player.type:type_name -> protobuf.player.PlayerType
	7,  // 1: protobuf.player.Player.ask_value:type_name -> google.protobuf.Int64Value
	8,  // 2: protobuf.player.Player.currency:type_name -> protobuf.Currency
	1,  // 3: protobuf.player.Players.players:type_name -> protobuf.player.Player
	9,  // 4: protobuf.player.UpdateRequest.is_listed:type_name -> google.protobuf.BoolValue
	7,  // 5: protobuf.player.UpdateRequest.ask_value:type_name -> google.protobuf.Int64Value
	7,  // 6: protobuf.player.UpdateRequest.value:type_name -> google.protobuf.Int64Value
	3,  // 7: protobuf.player.PlayerService.Get:input_type -> protobuf.player.GetRequest
	6,  // 8: protobuf.player.PlayerService.Update:input_type -> protobuf.player.UpdateRequest
	4,  // 9: protobuf.player.PlayerService.GetByTeam:input_type -> protobuf.player.GetByTeamRequest
	5,  // 10: protobuf.player.PlayerService.GetListed:input_type -> protobuf.player.GetListedRequest
	1,  // 11: protobuf.player.PlayerService.Get:output_type -> protobuf.player.Player
	1,  // 12: protobuf.player.PlayerService.Update:output_type -> protobuf.player.Player
	2,  // 13: protobuf.player.PlayerService.GetByTeam:output_type -> protobuf.player.Players
	2,  // 14: protobuf.player.PlayerService.GetListed:output_type -> protobuf.player.Players
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_player_player_proto_init() }
func file_player_player_proto_init() {
	if File_player_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_player_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_player_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Players); i {
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
		file_player_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_player_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByTeamRequest); i {
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
		file_player_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListedRequest); i {
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
		file_player_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
			RawDescriptor: file_player_player_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_player_player_proto_goTypes,
		DependencyIndexes: file_player_player_proto_depIdxs,
		EnumInfos:         file_player_player_proto_enumTypes,
		MessageInfos:      file_player_player_proto_msgTypes,
	}.Build()
	File_player_player_proto = out.File
	file_player_player_proto_rawDesc = nil
	file_player_player_proto_goTypes = nil
	file_player_player_proto_depIdxs = nil
}
