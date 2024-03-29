// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exchange.proto

package function

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{1}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type RegisterRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{2}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type RegisterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{3}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(dst, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

type IsAuthenticatedRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAuthenticatedRequest) Reset()         { *m = IsAuthenticatedRequest{} }
func (m *IsAuthenticatedRequest) String() string { return proto.CompactTextString(m) }
func (*IsAuthenticatedRequest) ProtoMessage()    {}
func (*IsAuthenticatedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{4}
}
func (m *IsAuthenticatedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthenticatedRequest.Unmarshal(m, b)
}
func (m *IsAuthenticatedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthenticatedRequest.Marshal(b, m, deterministic)
}
func (dst *IsAuthenticatedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthenticatedRequest.Merge(dst, src)
}
func (m *IsAuthenticatedRequest) XXX_Size() int {
	return xxx_messageInfo_IsAuthenticatedRequest.Size(m)
}
func (m *IsAuthenticatedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthenticatedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthenticatedRequest proto.InternalMessageInfo

func (m *IsAuthenticatedRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *IsAuthenticatedRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type IsAuthenticatedResponse struct {
	Authenticated        bool     `protobuf:"varint,1,opt,name=authenticated,proto3" json:"authenticated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAuthenticatedResponse) Reset()         { *m = IsAuthenticatedResponse{} }
func (m *IsAuthenticatedResponse) String() string { return proto.CompactTextString(m) }
func (*IsAuthenticatedResponse) ProtoMessage()    {}
func (*IsAuthenticatedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{5}
}
func (m *IsAuthenticatedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthenticatedResponse.Unmarshal(m, b)
}
func (m *IsAuthenticatedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthenticatedResponse.Marshal(b, m, deterministic)
}
func (dst *IsAuthenticatedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthenticatedResponse.Merge(dst, src)
}
func (m *IsAuthenticatedResponse) XXX_Size() int {
	return xxx_messageInfo_IsAuthenticatedResponse.Size(m)
}
func (m *IsAuthenticatedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthenticatedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthenticatedResponse proto.InternalMessageInfo

func (m *IsAuthenticatedResponse) GetAuthenticated() bool {
	if m != nil {
		return m.Authenticated
	}
	return false
}

type CreateConversationRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ConversationName     string   `protobuf:"bytes,3,opt,name=conversation_name,json=conversationName,proto3" json:"conversation_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConversationRequest) Reset()         { *m = CreateConversationRequest{} }
func (m *CreateConversationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateConversationRequest) ProtoMessage()    {}
func (*CreateConversationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{6}
}
func (m *CreateConversationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConversationRequest.Unmarshal(m, b)
}
func (m *CreateConversationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConversationRequest.Marshal(b, m, deterministic)
}
func (dst *CreateConversationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConversationRequest.Merge(dst, src)
}
func (m *CreateConversationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateConversationRequest.Size(m)
}
func (m *CreateConversationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConversationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConversationRequest proto.InternalMessageInfo

func (m *CreateConversationRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateConversationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CreateConversationRequest) GetConversationName() string {
	if m != nil {
		return m.ConversationName
	}
	return ""
}

type CreateConversationResponse struct {
	ConversationId       string   `protobuf:"bytes,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConversationResponse) Reset()         { *m = CreateConversationResponse{} }
func (m *CreateConversationResponse) String() string { return proto.CompactTextString(m) }
func (*CreateConversationResponse) ProtoMessage()    {}
func (*CreateConversationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{7}
}
func (m *CreateConversationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConversationResponse.Unmarshal(m, b)
}
func (m *CreateConversationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConversationResponse.Marshal(b, m, deterministic)
}
func (dst *CreateConversationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConversationResponse.Merge(dst, src)
}
func (m *CreateConversationResponse) XXX_Size() int {
	return xxx_messageInfo_CreateConversationResponse.Size(m)
}
func (m *CreateConversationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConversationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConversationResponse proto.InternalMessageInfo

func (m *CreateConversationResponse) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

type JoinConversationRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ConversationId       string   `protobuf:"bytes,3,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinConversationRequest) Reset()         { *m = JoinConversationRequest{} }
func (m *JoinConversationRequest) String() string { return proto.CompactTextString(m) }
func (*JoinConversationRequest) ProtoMessage()    {}
func (*JoinConversationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{8}
}
func (m *JoinConversationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinConversationRequest.Unmarshal(m, b)
}
func (m *JoinConversationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinConversationRequest.Marshal(b, m, deterministic)
}
func (dst *JoinConversationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinConversationRequest.Merge(dst, src)
}
func (m *JoinConversationRequest) XXX_Size() int {
	return xxx_messageInfo_JoinConversationRequest.Size(m)
}
func (m *JoinConversationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinConversationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinConversationRequest proto.InternalMessageInfo

func (m *JoinConversationRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *JoinConversationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *JoinConversationRequest) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

type JoinConversationResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinConversationResponse) Reset()         { *m = JoinConversationResponse{} }
func (m *JoinConversationResponse) String() string { return proto.CompactTextString(m) }
func (*JoinConversationResponse) ProtoMessage()    {}
func (*JoinConversationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{9}
}
func (m *JoinConversationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinConversationResponse.Unmarshal(m, b)
}
func (m *JoinConversationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinConversationResponse.Marshal(b, m, deterministic)
}
func (dst *JoinConversationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinConversationResponse.Merge(dst, src)
}
func (m *JoinConversationResponse) XXX_Size() int {
	return xxx_messageInfo_JoinConversationResponse.Size(m)
}
func (m *JoinConversationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinConversationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinConversationResponse proto.InternalMessageInfo

type IsConversationMemberRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ConversationId       string   `protobuf:"bytes,3,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsConversationMemberRequest) Reset()         { *m = IsConversationMemberRequest{} }
func (m *IsConversationMemberRequest) String() string { return proto.CompactTextString(m) }
func (*IsConversationMemberRequest) ProtoMessage()    {}
func (*IsConversationMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{10}
}
func (m *IsConversationMemberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsConversationMemberRequest.Unmarshal(m, b)
}
func (m *IsConversationMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsConversationMemberRequest.Marshal(b, m, deterministic)
}
func (dst *IsConversationMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsConversationMemberRequest.Merge(dst, src)
}
func (m *IsConversationMemberRequest) XXX_Size() int {
	return xxx_messageInfo_IsConversationMemberRequest.Size(m)
}
func (m *IsConversationMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsConversationMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsConversationMemberRequest proto.InternalMessageInfo

func (m *IsConversationMemberRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *IsConversationMemberRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *IsConversationMemberRequest) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

type IsConversationMemberResponse struct {
	Member               bool     `protobuf:"varint,1,opt,name=member,proto3" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsConversationMemberResponse) Reset()         { *m = IsConversationMemberResponse{} }
func (m *IsConversationMemberResponse) String() string { return proto.CompactTextString(m) }
func (*IsConversationMemberResponse) ProtoMessage()    {}
func (*IsConversationMemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{11}
}
func (m *IsConversationMemberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsConversationMemberResponse.Unmarshal(m, b)
}
func (m *IsConversationMemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsConversationMemberResponse.Marshal(b, m, deterministic)
}
func (dst *IsConversationMemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsConversationMemberResponse.Merge(dst, src)
}
func (m *IsConversationMemberResponse) XXX_Size() int {
	return xxx_messageInfo_IsConversationMemberResponse.Size(m)
}
func (m *IsConversationMemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsConversationMemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsConversationMemberResponse proto.InternalMessageInfo

func (m *IsConversationMemberResponse) GetMember() bool {
	if m != nil {
		return m.Member
	}
	return false
}

type SendMessageRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ConversationId       string   `protobuf:"bytes,3,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{12}
}
func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (dst *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(dst, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SendMessageRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SendMessageRequest) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

func (m *SendMessageRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type SendMessageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageResponse) Reset()         { *m = SendMessageResponse{} }
func (m *SendMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SendMessageResponse) ProtoMessage()    {}
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{13}
}
func (m *SendMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageResponse.Unmarshal(m, b)
}
func (m *SendMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageResponse.Marshal(b, m, deterministic)
}
func (dst *SendMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResponse.Merge(dst, src)
}
func (m *SendMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SendMessageResponse.Size(m)
}
func (m *SendMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageResponse proto.InternalMessageInfo

type ReadMessagesRequest struct {
	UserId         string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token          string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ConversationId string `protobuf:"bytes,3,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	// RFC3339 UTC
	FromTimestamp        string   `protobuf:"bytes,4,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	ToTimestamp          string   `protobuf:"bytes,5,opt,name=to_timestamp,json=toTimestamp,proto3" json:"to_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadMessagesRequest) Reset()         { *m = ReadMessagesRequest{} }
func (m *ReadMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*ReadMessagesRequest) ProtoMessage()    {}
func (*ReadMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{14}
}
func (m *ReadMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMessagesRequest.Unmarshal(m, b)
}
func (m *ReadMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMessagesRequest.Marshal(b, m, deterministic)
}
func (dst *ReadMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMessagesRequest.Merge(dst, src)
}
func (m *ReadMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_ReadMessagesRequest.Size(m)
}
func (m *ReadMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMessagesRequest proto.InternalMessageInfo

func (m *ReadMessagesRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ReadMessagesRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ReadMessagesRequest) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

func (m *ReadMessagesRequest) GetFromTimestamp() string {
	if m != nil {
		return m.FromTimestamp
	}
	return ""
}

func (m *ReadMessagesRequest) GetToTimestamp() string {
	if m != nil {
		return m.ToTimestamp
	}
	return ""
}

type ReadMessagesResponse struct {
	Messages             []*ReadMessagesResponse_Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ReadMessagesResponse) Reset()         { *m = ReadMessagesResponse{} }
func (m *ReadMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*ReadMessagesResponse) ProtoMessage()    {}
func (*ReadMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{15}
}
func (m *ReadMessagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMessagesResponse.Unmarshal(m, b)
}
func (m *ReadMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMessagesResponse.Marshal(b, m, deterministic)
}
func (dst *ReadMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMessagesResponse.Merge(dst, src)
}
func (m *ReadMessagesResponse) XXX_Size() int {
	return xxx_messageInfo_ReadMessagesResponse.Size(m)
}
func (m *ReadMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMessagesResponse proto.InternalMessageInfo

func (m *ReadMessagesResponse) GetMessages() []*ReadMessagesResponse_Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type ReadMessagesResponse_Message struct {
	// RFC3339 UTC
	Timestamp            string   `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadMessagesResponse_Message) Reset()         { *m = ReadMessagesResponse_Message{} }
func (m *ReadMessagesResponse_Message) String() string { return proto.CompactTextString(m) }
func (*ReadMessagesResponse_Message) ProtoMessage()    {}
func (*ReadMessagesResponse_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_8331bae46b75aada, []int{15, 0}
}
func (m *ReadMessagesResponse_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMessagesResponse_Message.Unmarshal(m, b)
}
func (m *ReadMessagesResponse_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMessagesResponse_Message.Marshal(b, m, deterministic)
}
func (dst *ReadMessagesResponse_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMessagesResponse_Message.Merge(dst, src)
}
func (m *ReadMessagesResponse_Message) XXX_Size() int {
	return xxx_messageInfo_ReadMessagesResponse_Message.Size(m)
}
func (m *ReadMessagesResponse_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMessagesResponse_Message.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMessagesResponse_Message proto.InternalMessageInfo

func (m *ReadMessagesResponse_Message) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *ReadMessagesResponse_Message) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ReadMessagesResponse_Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "protos.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "protos.LoginResponse")
	proto.RegisterType((*RegisterRequest)(nil), "protos.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "protos.RegisterResponse")
	proto.RegisterType((*IsAuthenticatedRequest)(nil), "protos.IsAuthenticatedRequest")
	proto.RegisterType((*IsAuthenticatedResponse)(nil), "protos.IsAuthenticatedResponse")
	proto.RegisterType((*CreateConversationRequest)(nil), "protos.CreateConversationRequest")
	proto.RegisterType((*CreateConversationResponse)(nil), "protos.CreateConversationResponse")
	proto.RegisterType((*JoinConversationRequest)(nil), "protos.JoinConversationRequest")
	proto.RegisterType((*JoinConversationResponse)(nil), "protos.JoinConversationResponse")
	proto.RegisterType((*IsConversationMemberRequest)(nil), "protos.IsConversationMemberRequest")
	proto.RegisterType((*IsConversationMemberResponse)(nil), "protos.IsConversationMemberResponse")
	proto.RegisterType((*SendMessageRequest)(nil), "protos.SendMessageRequest")
	proto.RegisterType((*SendMessageResponse)(nil), "protos.SendMessageResponse")
	proto.RegisterType((*ReadMessagesRequest)(nil), "protos.ReadMessagesRequest")
	proto.RegisterType((*ReadMessagesResponse)(nil), "protos.ReadMessagesResponse")
	proto.RegisterType((*ReadMessagesResponse_Message)(nil), "protos.ReadMessagesResponse.Message")
}

func init() { proto.RegisterFile("exchange.proto", fileDescriptor_exchange_8331bae46b75aada) }

var fileDescriptor_exchange_8331bae46b75aada = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0x49, 0x63, 0xf7, 0xcf, 0x69, 0x77, 0x5b, 0xa7, 0xb5, 0x1b, 0xb7, 0xbd, 0xa8, 0x43,
	0xc5, 0x42, 0x21, 0x17, 0x0a, 0x5e, 0xfa, 0x87, 0x22, 0x12, 0xb1, 0x5e, 0x44, 0xaf, 0xbc, 0x29,
	0xd3, 0xe4, 0x98, 0x06, 0x99, 0x99, 0xdd, 0xcc, 0x44, 0x7d, 0x02, 0x9f, 0x48, 0xf0, 0xf5, 0x64,
	0x27, 0x93, 0x64, 0xc2, 0x46, 0x50, 0xe8, 0x5e, 0x2d, 0xe7, 0x9b, 0x6f, 0xcf, 0xf7, 0xcb, 0x39,
	0xc3, 0xc0, 0x14, 0x7f, 0x24, 0xb7, 0x4c, 0x64, 0x18, 0x2e, 0x0a, 0xa9, 0x25, 0x19, 0x98, 0x1f,
	0x45, 0x2f, 0x60, 0xf7, 0xbd, 0xcc, 0x72, 0x11, 0xe3, 0xb2, 0x44, 0xa5, 0xc9, 0x31, 0x8c, 0x4b,
	0x85, 0xc5, 0xb5, 0x60, 0x1c, 0x03, 0xef, 0xd4, 0x3b, 0x1f, 0xc7, 0xa3, 0x95, 0xf0, 0x81, 0x71,
	0xa4, 0x2f, 0x60, 0x62, 0xcd, 0x6a, 0x21, 0x85, 0x42, 0x72, 0x08, 0xdb, 0x5a, 0x7e, 0x45, 0x61,
	0x9d, 0x55, 0x41, 0x66, 0x30, 0x34, 0x3d, 0xf2, 0x34, 0xd8, 0x32, 0xfa, 0x60, 0x55, 0x46, 0x29,
	0x0d, 0x61, 0x2f, 0xc6, 0x2c, 0x57, 0x1a, 0x8b, 0x7f, 0xca, 0x23, 0xb0, 0xdf, 0xfa, 0xab, 0x48,
	0xfa, 0x16, 0x8e, 0x22, 0xf5, 0xba, 0xd4, 0xb7, 0x28, 0x74, 0x9e, 0x30, 0x8d, 0x69, 0xdd, 0xea,
	0x3f, 0x61, 0x5e, 0xc2, 0x6c, 0xad, 0x91, 0xfd, 0xac, 0x33, 0x98, 0x30, 0xf7, 0xc0, 0x74, 0x1c,
	0xc5, 0x5d, 0x91, 0x7e, 0x87, 0x87, 0x97, 0x05, 0x32, 0x8d, 0x97, 0x52, 0x7c, 0xc3, 0x42, 0x31,
	0x9d, 0xcb, 0x66, 0x8e, 0x4e, 0xac, 0xe7, 0xc6, 0xb6, 0x94, 0x5b, 0x2e, 0xe5, 0x05, 0xdc, 0x4f,
	0x9c, 0x2e, 0xd5, 0x38, 0x7c, 0xe3, 0xd8, 0x77, 0x0f, 0xcc, 0x58, 0xde, 0xc0, 0xbc, 0x2f, 0xd8,
	0xc2, 0x3f, 0x81, 0xbd, 0x4e, 0xab, 0x86, 0x60, 0xea, 0xca, 0x51, 0x4a, 0x97, 0x30, 0x7b, 0x27,
	0x73, 0x71, 0x07, 0xf4, 0x3d, 0x91, 0x7e, 0x6f, 0xe4, 0x1c, 0x82, 0xf5, 0x48, 0xbb, 0xd8, 0x12,
	0x8e, 0x23, 0xe5, 0x9e, 0x5c, 0x21, 0xbf, 0x69, 0x2f, 0xca, 0xa6, 0x90, 0x9e, 0xc3, 0x49, 0x7f,
	0xac, 0x1d, 0xe7, 0x11, 0x0c, 0xb8, 0x51, 0xec, 0x25, 0xb0, 0x15, 0xfd, 0xe9, 0x01, 0xf9, 0x88,
	0x22, 0xbd, 0x42, 0xa5, 0x58, 0x86, 0x1b, 0xc6, 0x24, 0x01, 0x0c, 0x13, 0x29, 0x34, 0x0a, 0x1d,
	0xdc, 0x33, 0x86, 0xba, 0xa4, 0x0f, 0xe0, 0xa0, 0xc3, 0x61, 0xc7, 0xf9, 0xdb, 0x83, 0x83, 0x18,
	0x59, 0xad, 0xab, 0x4d, 0x03, 0x3e, 0x86, 0xe9, 0x97, 0x42, 0xf2, 0x6b, 0x9d, 0x73, 0x54, 0x9a,
	0xf1, 0x85, 0xe5, 0x9c, 0xac, 0xd4, 0x4f, 0xb5, 0x48, 0x1e, 0xc1, 0xae, 0x96, 0x8e, 0x69, 0xdb,
	0x98, 0x76, 0xb4, 0x6c, 0x2c, 0xf4, 0x97, 0x07, 0x87, 0x5d, 0x72, 0xbb, 0x8a, 0x57, 0x30, 0xe2,
	0x56, 0x0b, 0xbc, 0x53, 0xff, 0x7c, 0xe7, 0xe9, 0x59, 0xf5, 0x9a, 0xa9, 0xb0, 0xcf, 0x1f, 0xd6,
	0x23, 0x69, 0xfe, 0x35, 0xff, 0x0c, 0x43, 0x2b, 0x92, 0x13, 0x18, 0xb7, 0x14, 0xd5, 0x24, 0x5a,
	0xe1, 0xaf, 0xaf, 0x86, 0xbb, 0x07, 0xbf, 0xb3, 0x87, 0x9b, 0xea, 0x45, 0x7d, 0xf6, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x4d, 0xbc, 0x12, 0xd7, 0x6a, 0x05, 0x00, 0x00,
}
