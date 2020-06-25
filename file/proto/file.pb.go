// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package file

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FileShare struct {
	//The uuid of the file share.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The createdAt representing the server time when the file share was created.
	CreatedAt string `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// The updatedAt representing the server time when the file share was updated.
	UpdatedAt string `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// The name of the file share.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the file share.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// The uuid of the project that the file share belongs to.
	TenantId string `protobuf:"bytes,6,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// The uuid of the user that the file share belongs to.
	UserId string `protobuf:"bytes,7,opt,name=userId,proto3" json:"userId,omitempty"`
	// The uuid of the backend that the file share belongs to.
	BackendId string `protobuf:"bytes,8,opt,name=backendId,proto3" json:"backendId,omitempty"`
	// The name of the backend that the file share belongs to.
	Backend string `protobuf:"bytes,9,opt,name=backend,proto3" json:"backend,omitempty"`
	// The size of the file share requested by the user.
	Size int64 `protobuf:"varint,10,opt,name=size,proto3" json:"size,omitempty"`
	// The type of the file share.
	Type string `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty"`
	// The location that file share belongs to.
	Region string `protobuf:"bytes,12,opt,name=region,proto3" json:"region,omitempty"`
	// The locality that file share belongs to.
	AvailabilityZone string `protobuf:"bytes,13,opt,name=availabilityZone,proto3" json:"availabilityZone,omitempty"`
	// The status of the file share.
	Status string `protobuf:"bytes,14,opt,name=status,proto3" json:"status,omitempty"`
	// The uuid of the snapshot from which the file share is created
	SnapshotId string `protobuf:"bytes,15,opt,name=snapshotId,proto3" json:"snapshotId,omitempty"`
	// The protocol of the fileshare. e.g NFS, SMB etc.
	Protocols []string `protobuf:"bytes,16,rep,name=protocols,proto3" json:"protocols,omitempty"`
	// Any tags assigned to the file share.
	Tags []*Tag `protobuf:"bytes,17,rep,name=tags,proto3" json:"tags,omitempty"`
	// Indicates whether the file share is encrypted.
	Encrypted bool `protobuf:"varint,18,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	// EncryptionSettings that was used to protect the file share encryption.
	EncryptionSettings map[string]string `protobuf:"bytes,19,rep,name=encryptionSettings,proto3" json:"encryptionSettings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Metadata should be kept until the semantics between file share and backend storage resource.
	Metadata             *_struct.Struct `protobuf:"bytes,20,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FileShare) Reset()         { *m = FileShare{} }
func (m *FileShare) String() string { return proto.CompactTextString(m) }
func (*FileShare) ProtoMessage()    {}
func (*FileShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}

func (m *FileShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileShare.Unmarshal(m, b)
}
func (m *FileShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileShare.Marshal(b, m, deterministic)
}
func (m *FileShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileShare.Merge(m, src)
}
func (m *FileShare) XXX_Size() int {
	return xxx_messageInfo_FileShare.Size(m)
}
func (m *FileShare) XXX_DiscardUnknown() {
	xxx_messageInfo_FileShare.DiscardUnknown(m)
}

var xxx_messageInfo_FileShare proto.InternalMessageInfo

func (m *FileShare) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FileShare) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *FileShare) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *FileShare) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileShare) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FileShare) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *FileShare) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *FileShare) GetBackendId() string {
	if m != nil {
		return m.BackendId
	}
	return ""
}

func (m *FileShare) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *FileShare) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileShare) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FileShare) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *FileShare) GetAvailabilityZone() string {
	if m != nil {
		return m.AvailabilityZone
	}
	return ""
}

func (m *FileShare) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FileShare) GetSnapshotId() string {
	if m != nil {
		return m.SnapshotId
	}
	return ""
}

func (m *FileShare) GetProtocols() []string {
	if m != nil {
		return m.Protocols
	}
	return nil
}

func (m *FileShare) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FileShare) GetEncrypted() bool {
	if m != nil {
		return m.Encrypted
	}
	return false
}

func (m *FileShare) GetEncryptionSettings() map[string]string {
	if m != nil {
		return m.EncryptionSettings
	}
	return nil
}

func (m *FileShare) GetMetadata() *_struct.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Tag struct {
	// The key of the tag.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value of the tag.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{1}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ListFileShareRequest struct {
	Limit                int32             `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32             `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	SortKeys             []string          `protobuf:"bytes,3,rep,name=sortKeys,proto3" json:"sortKeys,omitempty"`
	SortDirs             []string          `protobuf:"bytes,4,rep,name=sortDirs,proto3" json:"sortDirs,omitempty"`
	Filter               map[string]string `protobuf:"bytes,5,rep,name=Filter,proto3" json:"Filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListFileShareRequest) Reset()         { *m = ListFileShareRequest{} }
func (m *ListFileShareRequest) String() string { return proto.CompactTextString(m) }
func (*ListFileShareRequest) ProtoMessage()    {}
func (*ListFileShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{2}
}

func (m *ListFileShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFileShareRequest.Unmarshal(m, b)
}
func (m *ListFileShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFileShareRequest.Marshal(b, m, deterministic)
}
func (m *ListFileShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFileShareRequest.Merge(m, src)
}
func (m *ListFileShareRequest) XXX_Size() int {
	return xxx_messageInfo_ListFileShareRequest.Size(m)
}
func (m *ListFileShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFileShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFileShareRequest proto.InternalMessageInfo

func (m *ListFileShareRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListFileShareRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListFileShareRequest) GetSortKeys() []string {
	if m != nil {
		return m.SortKeys
	}
	return nil
}

func (m *ListFileShareRequest) GetSortDirs() []string {
	if m != nil {
		return m.SortDirs
	}
	return nil
}

func (m *ListFileShareRequest) GetFilter() map[string]string {
	if m != nil {
		return m.Filter
	}
	return nil
}

type ListFileShareResponse struct {
	Fileshares           []*FileShare `protobuf:"bytes,1,rep,name=fileshares,proto3" json:"fileshares,omitempty"`
	Next                 int32        `protobuf:"varint,2,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListFileShareResponse) Reset()         { *m = ListFileShareResponse{} }
func (m *ListFileShareResponse) String() string { return proto.CompactTextString(m) }
func (*ListFileShareResponse) ProtoMessage()    {}
func (*ListFileShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{3}
}

func (m *ListFileShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFileShareResponse.Unmarshal(m, b)
}
func (m *ListFileShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFileShareResponse.Marshal(b, m, deterministic)
}
func (m *ListFileShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFileShareResponse.Merge(m, src)
}
func (m *ListFileShareResponse) XXX_Size() int {
	return xxx_messageInfo_ListFileShareResponse.Size(m)
}
func (m *ListFileShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFileShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFileShareResponse proto.InternalMessageInfo

func (m *ListFileShareResponse) GetFileshares() []*FileShare {
	if m != nil {
		return m.Fileshares
	}
	return nil
}

func (m *ListFileShareResponse) GetNext() int32 {
	if m != nil {
		return m.Next
	}
	return 0
}

type GetFileShareRequest struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fileshare            *FileShare `protobuf:"bytes,2,opt,name=fileshare,proto3" json:"fileshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetFileShareRequest) Reset()         { *m = GetFileShareRequest{} }
func (m *GetFileShareRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileShareRequest) ProtoMessage()    {}
func (*GetFileShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{4}
}

func (m *GetFileShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileShareRequest.Unmarshal(m, b)
}
func (m *GetFileShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileShareRequest.Marshal(b, m, deterministic)
}
func (m *GetFileShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileShareRequest.Merge(m, src)
}
func (m *GetFileShareRequest) XXX_Size() int {
	return xxx_messageInfo_GetFileShareRequest.Size(m)
}
func (m *GetFileShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileShareRequest proto.InternalMessageInfo

func (m *GetFileShareRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetFileShareRequest) GetFileshare() *FileShare {
	if m != nil {
		return m.Fileshare
	}
	return nil
}

type GetFileShareResponse struct {
	Fileshare            *FileShare `protobuf:"bytes,1,opt,name=fileshare,proto3" json:"fileshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetFileShareResponse) Reset()         { *m = GetFileShareResponse{} }
func (m *GetFileShareResponse) String() string { return proto.CompactTextString(m) }
func (*GetFileShareResponse) ProtoMessage()    {}
func (*GetFileShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{5}
}

func (m *GetFileShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileShareResponse.Unmarshal(m, b)
}
func (m *GetFileShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileShareResponse.Marshal(b, m, deterministic)
}
func (m *GetFileShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileShareResponse.Merge(m, src)
}
func (m *GetFileShareResponse) XXX_Size() int {
	return xxx_messageInfo_GetFileShareResponse.Size(m)
}
func (m *GetFileShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileShareResponse proto.InternalMessageInfo

func (m *GetFileShareResponse) GetFileshare() *FileShare {
	if m != nil {
		return m.Fileshare
	}
	return nil
}

type CreateFileShareRequest struct {
	Fileshare            *FileShare `protobuf:"bytes,1,opt,name=fileshare,proto3" json:"fileshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateFileShareRequest) Reset()         { *m = CreateFileShareRequest{} }
func (m *CreateFileShareRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFileShareRequest) ProtoMessage()    {}
func (*CreateFileShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{6}
}

func (m *CreateFileShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFileShareRequest.Unmarshal(m, b)
}
func (m *CreateFileShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFileShareRequest.Marshal(b, m, deterministic)
}
func (m *CreateFileShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFileShareRequest.Merge(m, src)
}
func (m *CreateFileShareRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFileShareRequest.Size(m)
}
func (m *CreateFileShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFileShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFileShareRequest proto.InternalMessageInfo

func (m *CreateFileShareRequest) GetFileshare() *FileShare {
	if m != nil {
		return m.Fileshare
	}
	return nil
}

type CreateFileShareResponse struct {
	Fileshare            *FileShare `protobuf:"bytes,1,opt,name=fileshare,proto3" json:"fileshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateFileShareResponse) Reset()         { *m = CreateFileShareResponse{} }
func (m *CreateFileShareResponse) String() string { return proto.CompactTextString(m) }
func (*CreateFileShareResponse) ProtoMessage()    {}
func (*CreateFileShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{7}
}

func (m *CreateFileShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFileShareResponse.Unmarshal(m, b)
}
func (m *CreateFileShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFileShareResponse.Marshal(b, m, deterministic)
}
func (m *CreateFileShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFileShareResponse.Merge(m, src)
}
func (m *CreateFileShareResponse) XXX_Size() int {
	return xxx_messageInfo_CreateFileShareResponse.Size(m)
}
func (m *CreateFileShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFileShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFileShareResponse proto.InternalMessageInfo

func (m *CreateFileShareResponse) GetFileshare() *FileShare {
	if m != nil {
		return m.Fileshare
	}
	return nil
}

type UpdateFileShareRequest struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fileshare            *FileShare `protobuf:"bytes,2,opt,name=fileshare,proto3" json:"fileshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateFileShareRequest) Reset()         { *m = UpdateFileShareRequest{} }
func (m *UpdateFileShareRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFileShareRequest) ProtoMessage()    {}
func (*UpdateFileShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{8}
}

func (m *UpdateFileShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFileShareRequest.Unmarshal(m, b)
}
func (m *UpdateFileShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFileShareRequest.Marshal(b, m, deterministic)
}
func (m *UpdateFileShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFileShareRequest.Merge(m, src)
}
func (m *UpdateFileShareRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFileShareRequest.Size(m)
}
func (m *UpdateFileShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFileShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFileShareRequest proto.InternalMessageInfo

func (m *UpdateFileShareRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateFileShareRequest) GetFileshare() *FileShare {
	if m != nil {
		return m.Fileshare
	}
	return nil
}

type UpdateFileShareResponse struct {
	Fileshare            *FileShare `protobuf:"bytes,1,opt,name=fileshare,proto3" json:"fileshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateFileShareResponse) Reset()         { *m = UpdateFileShareResponse{} }
func (m *UpdateFileShareResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateFileShareResponse) ProtoMessage()    {}
func (*UpdateFileShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{9}
}

func (m *UpdateFileShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFileShareResponse.Unmarshal(m, b)
}
func (m *UpdateFileShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFileShareResponse.Marshal(b, m, deterministic)
}
func (m *UpdateFileShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFileShareResponse.Merge(m, src)
}
func (m *UpdateFileShareResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateFileShareResponse.Size(m)
}
func (m *UpdateFileShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFileShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFileShareResponse proto.InternalMessageInfo

func (m *UpdateFileShareResponse) GetFileshare() *FileShare {
	if m != nil {
		return m.Fileshare
	}
	return nil
}

type DeleteFileShareRequest struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fileshare            *FileShare `protobuf:"bytes,2,opt,name=fileshare,proto3" json:"fileshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteFileShareRequest) Reset()         { *m = DeleteFileShareRequest{} }
func (m *DeleteFileShareRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFileShareRequest) ProtoMessage()    {}
func (*DeleteFileShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{10}
}

func (m *DeleteFileShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFileShareRequest.Unmarshal(m, b)
}
func (m *DeleteFileShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFileShareRequest.Marshal(b, m, deterministic)
}
func (m *DeleteFileShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFileShareRequest.Merge(m, src)
}
func (m *DeleteFileShareRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFileShareRequest.Size(m)
}
func (m *DeleteFileShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFileShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFileShareRequest proto.InternalMessageInfo

func (m *DeleteFileShareRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteFileShareRequest) GetFileshare() *FileShare {
	if m != nil {
		return m.Fileshare
	}
	return nil
}

type DeleteFileShareResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFileShareResponse) Reset()         { *m = DeleteFileShareResponse{} }
func (m *DeleteFileShareResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteFileShareResponse) ProtoMessage()    {}
func (*DeleteFileShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{11}
}

func (m *DeleteFileShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFileShareResponse.Unmarshal(m, b)
}
func (m *DeleteFileShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFileShareResponse.Marshal(b, m, deterministic)
}
func (m *DeleteFileShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFileShareResponse.Merge(m, src)
}
func (m *DeleteFileShareResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteFileShareResponse.Size(m)
}
func (m *DeleteFileShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFileShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFileShareResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FileShare)(nil), "FileShare")
	proto.RegisterMapType((map[string]string)(nil), "FileShare.EncryptionSettingsEntry")
	proto.RegisterType((*Tag)(nil), "Tag")
	proto.RegisterType((*ListFileShareRequest)(nil), "ListFileShareRequest")
	proto.RegisterMapType((map[string]string)(nil), "ListFileShareRequest.FilterEntry")
	proto.RegisterType((*ListFileShareResponse)(nil), "ListFileShareResponse")
	proto.RegisterType((*GetFileShareRequest)(nil), "GetFileShareRequest")
	proto.RegisterType((*GetFileShareResponse)(nil), "GetFileShareResponse")
	proto.RegisterType((*CreateFileShareRequest)(nil), "CreateFileShareRequest")
	proto.RegisterType((*CreateFileShareResponse)(nil), "CreateFileShareResponse")
	proto.RegisterType((*UpdateFileShareRequest)(nil), "UpdateFileShareRequest")
	proto.RegisterType((*UpdateFileShareResponse)(nil), "UpdateFileShareResponse")
	proto.RegisterType((*DeleteFileShareRequest)(nil), "DeleteFileShareRequest")
	proto.RegisterType((*DeleteFileShareResponse)(nil), "DeleteFileShareResponse")
}

func init() {
	proto.RegisterFile("file.proto", fileDescriptor_9188e3b7e55e1162)
}

var fileDescriptor_9188e3b7e55e1162 = []byte{
	// 737 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdb, 0x6e, 0xdb, 0x38,
	0x10, 0x8d, 0x7c, 0x8b, 0x3d, 0xce, 0x6d, 0x19, 0xc7, 0xe2, 0x0a, 0xc1, 0xc2, 0xab, 0x27, 0x23,
	0xc0, 0x32, 0x40, 0xf2, 0xb2, 0x59, 0x60, 0x81, 0xec, 0xe6, 0xb2, 0x08, 0xb6, 0x40, 0x01, 0x25,
	0x45, 0x81, 0xbe, 0xd1, 0xd6, 0xd8, 0x21, 0xa2, 0x48, 0xae, 0x48, 0x07, 0x75, 0x7f, 0xaf, 0x1f,
	0xd2, 0x7f, 0xe8, 0x17, 0x14, 0x24, 0x65, 0xf9, 0x26, 0x17, 0x49, 0x91, 0x37, 0xce, 0x39, 0xe4,
	0xe1, 0xd1, 0xcc, 0x11, 0x01, 0x06, 0x22, 0x42, 0x36, 0x4a, 0x13, 0x95, 0x78, 0x87, 0xc3, 0x24,
	0x19, 0x46, 0x78, 0x6c, 0xaa, 0xde, 0x78, 0x70, 0x2c, 0x55, 0x3a, 0xee, 0x2b, 0xcb, 0xfa, 0x5f,
	0xaa, 0xd0, 0xb8, 0x16, 0x11, 0xde, 0xde, 0xf3, 0x14, 0xc9, 0x0e, 0x94, 0x44, 0x48, 0x9d, 0x8e,
	0xd3, 0x6d, 0x04, 0x25, 0x11, 0x92, 0x43, 0x68, 0xf4, 0x53, 0xe4, 0x0a, 0xc3, 0x7f, 0x14, 0x2d,
	0x19, 0x78, 0x06, 0x68, 0x76, 0x3c, 0x0a, 0x33, 0xb6, 0x6c, 0xd9, 0x1c, 0x20, 0x04, 0x2a, 0x31,
	0x7f, 0x44, 0x5a, 0x31, 0x84, 0x59, 0x93, 0x0e, 0x34, 0x43, 0x94, 0xfd, 0x54, 0x8c, 0x94, 0x48,
	0x62, 0x5a, 0x35, 0xd4, 0x3c, 0x44, 0x3c, 0xa8, 0x2b, 0x8c, 0x79, 0xac, 0x6e, 0x42, 0x5a, 0x33,
	0x74, 0x5e, 0x93, 0x36, 0xd4, 0xc6, 0x12, 0xd3, 0x9b, 0x90, 0x6e, 0x1a, 0x26, 0xab, 0xb4, 0x8f,
	0x1e, 0xef, 0x3f, 0x60, 0x1c, 0xde, 0x84, 0xb4, 0x6e, 0x7d, 0xe4, 0x00, 0xa1, 0xb0, 0x99, 0x15,
	0xb4, 0x61, 0xb8, 0x69, 0xa9, 0x1d, 0x4a, 0xf1, 0x19, 0x29, 0x74, 0x9c, 0x6e, 0x39, 0x30, 0x6b,
	0x8d, 0xa9, 0xc9, 0x08, 0x69, 0xd3, 0xba, 0xd6, 0x6b, 0x7d, 0x6f, 0x8a, 0x43, 0x6d, 0x78, 0xcb,
	0xde, 0x6b, 0x2b, 0x72, 0x04, 0x7b, 0xfc, 0x89, 0x8b, 0x88, 0xf7, 0x44, 0x24, 0xd4, 0xe4, 0x43,
	0x12, 0x23, 0xdd, 0x36, 0x3b, 0x56, 0x70, 0xad, 0x21, 0x15, 0x57, 0x63, 0x49, 0x77, 0xac, 0x86,
	0xad, 0xc8, 0x6f, 0x00, 0x32, 0xe6, 0x23, 0x79, 0x9f, 0xe8, 0x2f, 0xde, 0x35, 0xdc, 0x1c, 0xa2,
	0xbf, 0xcd, 0x0c, 0xaa, 0x9f, 0x44, 0x92, 0xee, 0x75, 0xca, 0xfa, 0xdb, 0x72, 0x80, 0x50, 0xa8,
	0x28, 0x3e, 0x94, 0xf4, 0x97, 0x4e, 0xb9, 0xdb, 0x3c, 0xa9, 0xb0, 0x3b, 0x3e, 0x0c, 0x0c, 0xa2,
	0xcf, 0x61, 0xdc, 0x4f, 0x27, 0x23, 0x85, 0x21, 0x25, 0x1d, 0xa7, 0x5b, 0x0f, 0x66, 0x00, 0x09,
	0x80, 0x64, 0x85, 0x48, 0xe2, 0x5b, 0x54, 0x4a, 0xc4, 0x43, 0x49, 0xf7, 0x8d, 0x8a, 0xcf, 0xf2,
	0x3c, 0xb0, 0xab, 0x95, 0x4d, 0x57, 0xb1, 0x4a, 0x27, 0x41, 0xc1, 0x69, 0x72, 0x0a, 0xf5, 0x47,
	0x54, 0x3c, 0xe4, 0x8a, 0xd3, 0x56, 0xc7, 0xe9, 0x36, 0x4f, 0x5c, 0x66, 0xa3, 0xc7, 0xa6, 0xd1,
	0x63, 0xb7, 0x26, 0x7a, 0x41, 0xbe, 0xd1, 0xbb, 0x02, 0x77, 0xcd, 0x1d, 0x64, 0x0f, 0xca, 0x0f,
	0x38, 0xc9, 0xc2, 0xa8, 0x97, 0xa4, 0x05, 0xd5, 0x27, 0x1e, 0x8d, 0x31, 0x4b, 0xa2, 0x2d, 0xfe,
	0x2a, 0xfd, 0xe9, 0xf8, 0x7f, 0x40, 0xf9, 0x8e, 0x0f, 0x9f, 0x7b, 0xc4, 0xff, 0xe6, 0x40, 0xeb,
	0x8d, 0x90, 0x2a, 0xff, 0xd0, 0x00, 0x3f, 0x8e, 0x51, 0x2a, 0xbd, 0x3d, 0x12, 0x8f, 0x42, 0x19,
	0x89, 0x6a, 0x60, 0x0b, 0x3d, 0xbb, 0x64, 0x30, 0x90, 0x68, 0x7f, 0x81, 0x6a, 0x90, 0x55, 0x3a,
	0xab, 0x32, 0x49, 0xd5, 0xff, 0x38, 0x91, 0xb4, 0x6c, 0x46, 0x93, 0xd7, 0x53, 0xee, 0x52, 0xa4,
	0x92, 0x56, 0x66, 0x9c, 0xae, 0xc9, 0x19, 0xd4, 0xae, 0x45, 0xa4, 0x30, 0xa5, 0x55, 0xd3, 0xf1,
	0xdf, 0x59, 0x91, 0x19, 0x66, 0xf7, 0xd8, 0x86, 0x67, 0x07, 0xbc, 0x33, 0x68, 0xce, 0xc1, 0x2f,
	0xea, 0xd1, 0x7b, 0x38, 0x58, 0xba, 0x46, 0x8e, 0x92, 0x58, 0x22, 0x39, 0xb2, 0xcf, 0x85, 0xd4,
	0xa0, 0xa4, 0x8e, 0xb1, 0x04, 0xb3, 0x10, 0x04, 0x73, 0xac, 0xf9, 0xa9, 0xf1, 0xd3, 0xb4, 0x11,
	0x66, 0xed, 0xbf, 0x85, 0xfd, 0xff, 0x70, 0xb5, 0x97, 0xcb, 0x6f, 0x49, 0x17, 0x1a, 0xb9, 0x90,
	0x39, 0xbf, 0x78, 0xcb, 0x8c, 0xf4, 0xcf, 0xa1, 0xb5, 0x28, 0x98, 0x19, 0x5d, 0x50, 0x70, 0x7e,
	0xa4, 0xf0, 0x2f, 0xb4, 0x2f, 0xcc, 0x33, 0xb5, 0xe2, 0xea, 0xf9, 0x1a, 0x17, 0xe0, 0xae, 0x68,
	0xbc, 0xd8, 0x48, 0x00, 0xed, 0x77, 0xe6, 0x45, 0x7c, 0xc5, 0xf6, 0x5c, 0x80, 0xbb, 0xa2, 0xf9,
	0x33, 0xc6, 0x2e, 0x31, 0xc2, 0x57, 0x35, 0xf6, 0x2b, 0xb8, 0x2b, 0x9a, 0xd6, 0xd8, 0xc9, 0xd7,
	0x12, 0x54, 0x34, 0x4a, 0xce, 0x61, 0x7b, 0x21, 0x85, 0xe4, 0xa0, 0x30, 0xfc, 0x5e, 0x9b, 0x15,
	0x86, 0xd5, 0xdf, 0x20, 0x7f, 0xc3, 0xd6, 0x7c, 0x3a, 0x48, 0x8b, 0x15, 0xa4, 0xcf, 0x3b, 0x60,
	0x45, 0x11, 0xf2, 0x37, 0xc8, 0x35, 0xec, 0x2e, 0x8d, 0x95, 0xb8, 0xac, 0x38, 0x2c, 0x1e, 0x65,
	0x6b, 0x12, 0x60, 0x75, 0x96, 0xa6, 0x40, 0x5c, 0x56, 0x3c, 0x6b, 0x8f, 0xb2, 0x35, 0x03, 0xb3,
	0x3a, 0x4b, 0x4d, 0x23, 0x2e, 0x2b, 0x1e, 0x8d, 0x47, 0xd9, 0x9a, 0xfe, 0xfa, 0x1b, 0xbd, 0x9a,
	0x79, 0x64, 0x4f, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x13, 0x89, 0x53, 0x84, 0xfb, 0x07, 0x00,
	0x00,
}