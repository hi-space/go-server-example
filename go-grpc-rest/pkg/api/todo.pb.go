// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: todo.proto

package api

import (
	context "context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ToDo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Reminder    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=reminder,proto3" json:"reminder,omitempty"`
}

func (x *ToDo) Reset() {
	*x = ToDo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo) ProtoMessage() {}

func (x *ToDo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo.ProtoReflect.Descriptor instead.
func (*ToDo) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *ToDo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ToDo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ToDo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ToDo) GetReminder() *timestamp.Timestamp {
	if x != nil {
		return x.Reminder
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api  string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	ToDo *ToDo  `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *CreateRequest) GetToDo() *ToDo {
	if x != nil {
		return x.ToDo
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *CreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{3}
}

func (x *ReadRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *ReadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api  string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	ToDo *ToDo  `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{4}
}

func (x *ReadResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *ReadResponse) GetToDo() *ToDo {
	if x != nil {
		return x.ToDo
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api  string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	ToDo *ToDo  `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[5]
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
	return file_todo_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *UpdateRequest) GetToDo() *ToDo {
	if x != nil {
		return x.ToDo
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api     string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Updated *ToDo  `protobuf:"bytes,2,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *UpdateResponse) GetUpdated() *ToDo {
	if x != nil {
		return x.Updated
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api     string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Deleted string `protobuf:"bytes,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *DeleteResponse) GetDeleted() string {
	if x != nil {
		return x.Deleted
	}
	return ""
}

type ReadAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
}

func (x *ReadAllRequest) Reset() {
	*x = ReadAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllRequest) ProtoMessage() {}

func (x *ReadAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllRequest.ProtoReflect.Descriptor instead.
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{9}
}

func (x *ReadAllRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

type ReadAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api   string  `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	ToDos []*ToDo `protobuf:"bytes,2,rep,name=toDos,proto3" json:"toDos,omitempty"`
}

func (x *ReadAllResponse) Reset() {
	*x = ReadAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllResponse) ProtoMessage() {}

func (x *ReadAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllResponse.ProtoReflect.Descriptor instead.
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{10}
}

func (x *ReadAllResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *ReadAllResponse) GetToDos() []*ToDo {
	if x != nil {
		return x.ToDos
	}
	return nil
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01,
	0x0a, 0x04, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x0a, 0x08, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x72, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x6f, 0x44,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x44,
	0x6f, 0x52, 0x04, 0x74, 0x6f, 0x44, 0x6f, 0x22, 0x32, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x0b, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70,
	0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0c,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x1c,
	0x0a, 0x04, 0x74, 0x6f, 0x44, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x44, 0x6f, 0x22, 0x3f, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12,
	0x1c, 0x0a, 0x04, 0x74, 0x6f, 0x44, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x44, 0x6f, 0x22, 0x46, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70,
	0x69, 0x12, 0x22, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x31, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70,
	0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x22, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x22, 0x43, 0x0a, 0x0f, 0x52, 0x65,
	0x61, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12,
	0x1e, 0x0a, 0x05, 0x74, 0x6f, 0x44, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x05, 0x74, 0x6f, 0x44, 0x6f, 0x73, 0x32,
	0xf7, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x44, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x44, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x40, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x0f, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x1a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x7b, 0x74, 0x6f, 0x44, 0x6f,
	0x2e, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x46, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x2a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x48, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x61, 0x6c, 0x6c, 0x42, 0xdd, 0x01, 0x5a, 0x07, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x92, 0x41, 0xd0, 0x01, 0x12, 0x6a, 0x0a, 0x0c, 0x54, 0x6f, 0x44,
	0x6f, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x55, 0x0a, 0x14, 0x67, 0x6f, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x72, 0x65, 0x73, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x2d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x72, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x68, 0x69, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x6a, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d,
	0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x3b, 0x0a, 0x03,
	0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20,
	0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_todo_proto_goTypes = []interface{}{
	(*ToDo)(nil),                // 0: v1.ToDo
	(*CreateRequest)(nil),       // 1: v1.CreateRequest
	(*CreateResponse)(nil),      // 2: v1.CreateResponse
	(*ReadRequest)(nil),         // 3: v1.ReadRequest
	(*ReadResponse)(nil),        // 4: v1.ReadResponse
	(*UpdateRequest)(nil),       // 5: v1.UpdateRequest
	(*UpdateResponse)(nil),      // 6: v1.UpdateResponse
	(*DeleteRequest)(nil),       // 7: v1.DeleteRequest
	(*DeleteResponse)(nil),      // 8: v1.DeleteResponse
	(*ReadAllRequest)(nil),      // 9: v1.ReadAllRequest
	(*ReadAllResponse)(nil),     // 10: v1.ReadAllResponse
	(*timestamp.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_todo_proto_depIdxs = []int32{
	11, // 0: v1.ToDo.reminder:type_name -> google.protobuf.Timestamp
	0,  // 1: v1.CreateRequest.toDo:type_name -> v1.ToDo
	0,  // 2: v1.ReadResponse.toDo:type_name -> v1.ToDo
	0,  // 3: v1.UpdateRequest.toDo:type_name -> v1.ToDo
	0,  // 4: v1.UpdateResponse.updated:type_name -> v1.ToDo
	0,  // 5: v1.ReadAllResponse.toDos:type_name -> v1.ToDo
	1,  // 6: v1.ToDoService.Create:input_type -> v1.CreateRequest
	3,  // 7: v1.ToDoService.Read:input_type -> v1.ReadRequest
	5,  // 8: v1.ToDoService.Update:input_type -> v1.UpdateRequest
	7,  // 9: v1.ToDoService.Delete:input_type -> v1.DeleteRequest
	9,  // 10: v1.ToDoService.ReadAll:input_type -> v1.ReadAllRequest
	2,  // 11: v1.ToDoService.Create:output_type -> v1.CreateResponse
	4,  // 12: v1.ToDoService.Read:output_type -> v1.ReadResponse
	6,  // 13: v1.ToDoService.Update:output_type -> v1.UpdateResponse
	8,  // 14: v1.ToDoService.Delete:output_type -> v1.DeleteResponse
	10, // 15: v1.ToDoService.ReadAll:output_type -> v1.ReadAllResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDo); i {
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
		file_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_todo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_todo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_todo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllRequest); i {
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
		file_todo_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllResponse); i {
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
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ToDoServiceClient is the client API for ToDoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ToDoServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type toDoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToDoServiceClient(cc grpc.ClientConnInterface) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

func (c *toDoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoServiceServer is the server API for ToDoService service.
type ToDoServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
}

// UnimplementedToDoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedToDoServiceServer struct {
}

func (*UnimplementedToDoServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedToDoServiceServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedToDoServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedToDoServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedToDoServiceServer) ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}

func RegisterToDoServiceServer(s *grpc.Server, srv ToDoServiceServer) {
	s.RegisterService(&_ToDoService_serviceDesc, srv)
}

func _ToDoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToDoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ToDoService",
	HandlerType: (*ToDoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ToDoService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ToDoService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ToDoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ToDoService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _ToDoService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
