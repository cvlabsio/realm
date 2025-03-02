// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: c2.proto

package c2pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Host_Platform int32

const (
	Host_PLATFORM_UNSPECIFIED Host_Platform = 0
	Host_PLATFORM_WINDOWS     Host_Platform = 1
	Host_PLATFORM_LINUX       Host_Platform = 2
	Host_PLATFORM_MACOS       Host_Platform = 3
	Host_PLATFORM_BSD         Host_Platform = 4
)

// Enum value maps for Host_Platform.
var (
	Host_Platform_name = map[int32]string{
		0: "PLATFORM_UNSPECIFIED",
		1: "PLATFORM_WINDOWS",
		2: "PLATFORM_LINUX",
		3: "PLATFORM_MACOS",
		4: "PLATFORM_BSD",
	}
	Host_Platform_value = map[string]int32{
		"PLATFORM_UNSPECIFIED": 0,
		"PLATFORM_WINDOWS":     1,
		"PLATFORM_LINUX":       2,
		"PLATFORM_MACOS":       3,
		"PLATFORM_BSD":         4,
	}
)

func (x Host_Platform) Enum() *Host_Platform {
	p := new(Host_Platform)
	*p = x
	return p
}

func (x Host_Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Host_Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_c2_proto_enumTypes[0].Descriptor()
}

func (Host_Platform) Type() protoreflect.EnumType {
	return &file_c2_proto_enumTypes[0]
}

func (x Host_Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Host_Platform.Descriptor instead.
func (Host_Platform) EnumDescriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{0, 0}
}

// Host information for the system a beacon is running on.
type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string        `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Name       string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Platform   Host_Platform `protobuf:"varint,3,opt,name=platform,proto3,enum=c2.Host_Platform" json:"platform,omitempty"`
	PrimaryIp  string        `protobuf:"bytes,4,opt,name=primary_ip,json=primaryIp,proto3" json:"primary_ip,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Host) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Host) GetPlatform() Host_Platform {
	if x != nil {
		return x.Platform
	}
	return Host_PLATFORM_UNSPECIFIED
}

func (x *Host) GetPrimaryIp() string {
	if x != nil {
		return x.PrimaryIp
	}
	return ""
}

// Agent information to identify the type of beacon.
type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{1}
}

func (x *Agent) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

// Beacon information that is unique to the current running beacon.
type Beacon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Principal  string `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
	Host       *Host  `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Agent      *Agent `protobuf:"bytes,4,opt,name=agent,proto3" json:"agent,omitempty"`
	Interval   uint64 `protobuf:"varint,5,opt,name=interval,proto3" json:"interval,omitempty"` // Duration until next callback, in seconds.
}

func (x *Beacon) Reset() {
	*x = Beacon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Beacon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Beacon) ProtoMessage() {}

func (x *Beacon) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Beacon.ProtoReflect.Descriptor instead.
func (*Beacon) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{2}
}

func (x *Beacon) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Beacon) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *Beacon) GetHost() *Host {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *Beacon) GetAgent() *Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

func (x *Beacon) GetInterval() uint64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

// Task instructions for the beacon to execute.
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Eldritch   string            `protobuf:"bytes,2,opt,name=eldritch,proto3" json:"eldritch,omitempty"`
	Parameters map[string]string `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FileNames  []string          `protobuf:"bytes,4,rep,name=file_names,json=fileNames,proto3" json:"file_names,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{3}
}

func (x *Task) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetEldritch() string {
	if x != nil {
		return x.Eldritch
	}
	return ""
}

func (x *Task) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *Task) GetFileNames() []string {
	if x != nil {
		return x.FileNames
	}
	return nil
}

// TaskError provides information when task execution fails.
type TaskError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *TaskError) Reset() {
	*x = TaskError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskError) ProtoMessage() {}

func (x *TaskError) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskError.ProtoReflect.Descriptor instead.
func (*TaskError) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{4}
}

func (x *TaskError) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// TaskOutput provides information about a running task.
type TaskOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Output string     `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Error  *TaskError `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// Indicates the UTC timestamp task execution began, set only in the first message for reporting.
	ExecStartedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=exec_started_at,json=execStartedAt,proto3" json:"exec_started_at,omitempty"`
	// Indicates the UTC timestamp task execution completed, set only in last message for reporting.
	ExecFinishedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=exec_finished_at,json=execFinishedAt,proto3" json:"exec_finished_at,omitempty"`
}

func (x *TaskOutput) Reset() {
	*x = TaskOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskOutput) ProtoMessage() {}

func (x *TaskOutput) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskOutput.ProtoReflect.Descriptor instead.
func (*TaskOutput) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{5}
}

func (x *TaskOutput) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskOutput) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *TaskOutput) GetError() *TaskError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *TaskOutput) GetExecStartedAt() *timestamp.Timestamp {
	if x != nil {
		return x.ExecStartedAt
	}
	return nil
}

func (x *TaskOutput) GetExecFinishedAt() *timestamp.Timestamp {
	if x != nil {
		return x.ExecFinishedAt
	}
	return nil
}

// RPC Messages
type ClaimTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beacon *Beacon `protobuf:"bytes,1,opt,name=beacon,proto3" json:"beacon,omitempty"`
}

func (x *ClaimTasksRequest) Reset() {
	*x = ClaimTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimTasksRequest) ProtoMessage() {}

func (x *ClaimTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimTasksRequest.ProtoReflect.Descriptor instead.
func (*ClaimTasksRequest) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{6}
}

func (x *ClaimTasksRequest) GetBeacon() *Beacon {
	if x != nil {
		return x.Beacon
	}
	return nil
}

type ClaimTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *ClaimTasksResponse) Reset() {
	*x = ClaimTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimTasksResponse) ProtoMessage() {}

func (x *ClaimTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimTasksResponse.ProtoReflect.Descriptor instead.
func (*ClaimTasksResponse) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{7}
}

func (x *ClaimTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type ReportTaskOutputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output *TaskOutput `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ReportTaskOutputRequest) Reset() {
	*x = ReportTaskOutputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportTaskOutputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportTaskOutputRequest) ProtoMessage() {}

func (x *ReportTaskOutputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportTaskOutputRequest.ProtoReflect.Descriptor instead.
func (*ReportTaskOutputRequest) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{8}
}

func (x *ReportTaskOutputRequest) GetOutput() *TaskOutput {
	if x != nil {
		return x.Output
	}
	return nil
}

type ReportTaskOutputResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportTaskOutputResponse) Reset() {
	*x = ReportTaskOutputResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportTaskOutputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportTaskOutputResponse) ProtoMessage() {}

func (x *ReportTaskOutputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportTaskOutputResponse.ProtoReflect.Descriptor instead.
func (*ReportTaskOutputResponse) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{9}
}

type DownloadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DownloadFileRequest) Reset() {
	*x = DownloadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileRequest) ProtoMessage() {}

func (x *DownloadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadFileRequest) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{10}
}

func (x *DownloadFileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DownloadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *DownloadFileResponse) Reset() {
	*x = DownloadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileResponse) ProtoMessage() {}

func (x *DownloadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_c2_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileResponse.ProtoReflect.Descriptor instead.
func (*DownloadFileResponse) Descriptor() ([]byte, []int) {
	return file_c2_proto_rawDescGZIP(), []int{11}
}

func (x *DownloadFileResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_c2_proto protoreflect.FileDescriptor

var file_c2_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x63, 0x32, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfe, 0x01, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x63, 0x32, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x70, 0x22, 0x74, 0x0a, 0x08, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f,
	0x52, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x57, 0x49, 0x4e,
	0x44, 0x4f, 0x57, 0x53, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f,
	0x52, 0x4d, 0x5f, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4c,
	0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x4d, 0x41, 0x43, 0x4f, 0x53, 0x10, 0x03, 0x12, 0x10,
	0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x42, 0x53, 0x44, 0x10, 0x04,
	0x22, 0x27, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xa1, 0x01, 0x0a, 0x06, 0x42, 0x65,
	0x61, 0x63, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x63, 0x32, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x63, 0x32, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0xca, 0x01,
	0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6c, 0x64, 0x72, 0x69, 0x74,
	0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6c, 0x64, 0x72, 0x69, 0x74,
	0x63, 0x68, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x32, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1d, 0x0a, 0x09, 0x54, 0x61,
	0x73, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xe3, 0x01, 0x0a, 0x0a, 0x54, 0x61,
	0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x32, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x65, 0x78, 0x65,
	0x63, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0e, 0x65, 0x78, 0x65, 0x63, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x37, 0x0a, 0x11, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x32, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e,
	0x52, 0x06, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x12, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x63, 0x32, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x41,
	0x0a, 0x17, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x32, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x0a,
	0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0xd9, 0x01, 0x0a, 0x02, 0x43, 0x32, 0x12, 0x3d, 0x0a,
	0x0a, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x15, 0x2e, 0x63, 0x32,
	0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x32, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x10,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x1b, 0x2e, 0x63, 0x32, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x63, 0x32, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e,
	0x63, 0x32, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x32, 0x2e, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x63, 0x61, 0x72, 0x72, 0x65, 0x74, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x76, 0x65, 0x72,
	0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x32, 0x2f, 0x63, 0x32,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_c2_proto_rawDescOnce sync.Once
	file_c2_proto_rawDescData = file_c2_proto_rawDesc
)

func file_c2_proto_rawDescGZIP() []byte {
	file_c2_proto_rawDescOnce.Do(func() {
		file_c2_proto_rawDescData = protoimpl.X.CompressGZIP(file_c2_proto_rawDescData)
	})
	return file_c2_proto_rawDescData
}

var file_c2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_c2_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_c2_proto_goTypes = []interface{}{
	(Host_Platform)(0),               // 0: c2.Host.Platform
	(*Host)(nil),                     // 1: c2.Host
	(*Agent)(nil),                    // 2: c2.Agent
	(*Beacon)(nil),                   // 3: c2.Beacon
	(*Task)(nil),                     // 4: c2.Task
	(*TaskError)(nil),                // 5: c2.TaskError
	(*TaskOutput)(nil),               // 6: c2.TaskOutput
	(*ClaimTasksRequest)(nil),        // 7: c2.ClaimTasksRequest
	(*ClaimTasksResponse)(nil),       // 8: c2.ClaimTasksResponse
	(*ReportTaskOutputRequest)(nil),  // 9: c2.ReportTaskOutputRequest
	(*ReportTaskOutputResponse)(nil), // 10: c2.ReportTaskOutputResponse
	(*DownloadFileRequest)(nil),      // 11: c2.DownloadFileRequest
	(*DownloadFileResponse)(nil),     // 12: c2.DownloadFileResponse
	nil,                              // 13: c2.Task.ParametersEntry
	(*timestamp.Timestamp)(nil),      // 14: google.protobuf.Timestamp
}
var file_c2_proto_depIdxs = []int32{
	0,  // 0: c2.Host.platform:type_name -> c2.Host.Platform
	1,  // 1: c2.Beacon.host:type_name -> c2.Host
	2,  // 2: c2.Beacon.agent:type_name -> c2.Agent
	13, // 3: c2.Task.parameters:type_name -> c2.Task.ParametersEntry
	5,  // 4: c2.TaskOutput.error:type_name -> c2.TaskError
	14, // 5: c2.TaskOutput.exec_started_at:type_name -> google.protobuf.Timestamp
	14, // 6: c2.TaskOutput.exec_finished_at:type_name -> google.protobuf.Timestamp
	3,  // 7: c2.ClaimTasksRequest.beacon:type_name -> c2.Beacon
	4,  // 8: c2.ClaimTasksResponse.tasks:type_name -> c2.Task
	6,  // 9: c2.ReportTaskOutputRequest.output:type_name -> c2.TaskOutput
	7,  // 10: c2.C2.ClaimTasks:input_type -> c2.ClaimTasksRequest
	9,  // 11: c2.C2.ReportTaskOutput:input_type -> c2.ReportTaskOutputRequest
	11, // 12: c2.C2.DownloadFile:input_type -> c2.DownloadFileRequest
	8,  // 13: c2.C2.ClaimTasks:output_type -> c2.ClaimTasksResponse
	10, // 14: c2.C2.ReportTaskOutput:output_type -> c2.ReportTaskOutputResponse
	12, // 15: c2.C2.DownloadFile:output_type -> c2.DownloadFileResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_c2_proto_init() }
func file_c2_proto_init() {
	if File_c2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_c2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_c2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
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
		file_c2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Beacon); i {
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
		file_c2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_c2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskError); i {
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
		file_c2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskOutput); i {
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
		file_c2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimTasksRequest); i {
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
		file_c2_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimTasksResponse); i {
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
		file_c2_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportTaskOutputRequest); i {
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
		file_c2_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportTaskOutputResponse); i {
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
		file_c2_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileRequest); i {
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
		file_c2_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileResponse); i {
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
			RawDescriptor: file_c2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_c2_proto_goTypes,
		DependencyIndexes: file_c2_proto_depIdxs,
		EnumInfos:         file_c2_proto_enumTypes,
		MessageInfos:      file_c2_proto_msgTypes,
	}.Build()
	File_c2_proto = out.File
	file_c2_proto_rawDesc = nil
	file_c2_proto_goTypes = nil
	file_c2_proto_depIdxs = nil
}
