// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package common_v1 is a generated protocol buffer package.

It is generated from these files:
	common.proto

It has these top-level messages:
	RpcResult
	DataplaneCommon
	Aes256Key
	Aes256IV
	XtsAes256Key
	EncryptedXtsAesKey
	XtsKeyHMAC
	VolumeKeySet
	VolumeCrypto
	DaemonStatus
	ConfigUpdateReply
	ConfigBool
	ConfigGetBoolReply
	ConfigBoolList
	ConfigUint32
	ConfigGetUint32Reply
	ConfigUint32List
	ConfigUint64
	ConfigGetUint64Reply
	ConfigUint64List
	ConfigString
	ConfigGetStringReply
	ConfigStringList
	ConfigKey
	ConfigListQuery
*/
package common_v1

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

type StatusCode int32

const (
	// Status code not set.
	StatusCode_STATUS_NONE StatusCode = 0
	// Generic 'no problem' response.
	StatusCode_STATUS_OK StatusCode = 200
	// Accepted, but might not take effect immediately.
	StatusCode_STATUS_ACCEPTED StatusCode = 202
	// Resource created.
	StatusCode_STATUS_CREATED StatusCode = 201
	// Generic 'huh?'.
	StatusCode_STATUS_BAD_REQUEST StatusCode = 400
	// Not permitted. We're unlikely to use this, it indicates authorization will help.
	StatusCode_STATUS_UNAUTHORIZED StatusCode = 401
	// Not allowed, authorization isn't the issue.
	StatusCode_STATUS_FORBIDDEN StatusCode = 403
	// Resource not found.
	StatusCode_STATUS_NOT_FOUND StatusCode = 404
	// Operation would cause a conflict somewhere.
	StatusCode_STATUS_CONFLICT StatusCode = 409
	// General serious failure.
	StatusCode_STATUS_SERVER_ERROR StatusCode = 500
)

var StatusCode_name = map[int32]string{
	0:   "STATUS_NONE",
	200: "STATUS_OK",
	202: "STATUS_ACCEPTED",
	201: "STATUS_CREATED",
	400: "STATUS_BAD_REQUEST",
	401: "STATUS_UNAUTHORIZED",
	403: "STATUS_FORBIDDEN",
	404: "STATUS_NOT_FOUND",
	409: "STATUS_CONFLICT",
	500: "STATUS_SERVER_ERROR",
}
var StatusCode_value = map[string]int32{
	"STATUS_NONE":         0,
	"STATUS_OK":           200,
	"STATUS_ACCEPTED":     202,
	"STATUS_CREATED":      201,
	"STATUS_BAD_REQUEST":  400,
	"STATUS_UNAUTHORIZED": 401,
	"STATUS_FORBIDDEN":    403,
	"STATUS_NOT_FOUND":    404,
	"STATUS_CONFLICT":     409,
	"STATUS_SERVER_ERROR": 500,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}
func (StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ConfigGetStatus int32

const (
	ConfigGetStatus_GET_NONE    ConfigGetStatus = 0
	ConfigGetStatus_GET_OK      ConfigGetStatus = 1
	ConfigGetStatus_GET_MISSING ConfigGetStatus = 2
	ConfigGetStatus_GET_FAILED  ConfigGetStatus = 3
)

var ConfigGetStatus_name = map[int32]string{
	0: "GET_NONE",
	1: "GET_OK",
	2: "GET_MISSING",
	3: "GET_FAILED",
}
var ConfigGetStatus_value = map[string]int32{
	"GET_NONE":    0,
	"GET_OK":      1,
	"GET_MISSING": 2,
	"GET_FAILED":  3,
}

func (x ConfigGetStatus) String() string {
	return proto.EnumName(ConfigGetStatus_name, int32(x))
}
func (ConfigGetStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ConfigUpdateStatus int32

const (
	ConfigUpdateStatus_SET_NONE      ConfigUpdateStatus = 0
	ConfigUpdateStatus_SET_OK        ConfigUpdateStatus = 1
	ConfigUpdateStatus_SET_IMMUTABLE ConfigUpdateStatus = 2
	ConfigUpdateStatus_SET_FAILED    ConfigUpdateStatus = 3
	ConfigUpdateStatus_SET_READ_ONLY ConfigUpdateStatus = 4
)

var ConfigUpdateStatus_name = map[int32]string{
	0: "SET_NONE",
	1: "SET_OK",
	2: "SET_IMMUTABLE",
	3: "SET_FAILED",
	4: "SET_READ_ONLY",
}
var ConfigUpdateStatus_value = map[string]int32{
	"SET_NONE":      0,
	"SET_OK":        1,
	"SET_IMMUTABLE": 2,
	"SET_FAILED":    3,
	"SET_READ_ONLY": 4,
}

func (x ConfigUpdateStatus) String() string {
	return proto.EnumName(ConfigUpdateStatus_name, int32(x))
}
func (ConfigUpdateStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type VolumeCrypto_VolumeCryptoState int32

const (
	VolumeCrypto_VCS_NONE            VolumeCrypto_VolumeCryptoState = 0
	VolumeCrypto_VCS_FULLY_ENCRYPTED VolumeCrypto_VolumeCryptoState = 1
)

var VolumeCrypto_VolumeCryptoState_name = map[int32]string{
	0: "VCS_NONE",
	1: "VCS_FULLY_ENCRYPTED",
}
var VolumeCrypto_VolumeCryptoState_value = map[string]int32{
	"VCS_NONE":            0,
	"VCS_FULLY_ENCRYPTED": 1,
}

func (x VolumeCrypto_VolumeCryptoState) String() string {
	return proto.EnumName(VolumeCrypto_VolumeCryptoState_name, int32(x))
}
func (VolumeCrypto_VolumeCryptoState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

type DaemonStatus_ProgramState int32

const (
	DaemonStatus_DS_NONE     DaemonStatus_ProgramState = 0
	DaemonStatus_DS_STARTING DaemonStatus_ProgramState = 1
	DaemonStatus_DS_READY    DaemonStatus_ProgramState = 2
	DaemonStatus_DS_STOPPING DaemonStatus_ProgramState = 3
	DaemonStatus_DS_STOPPED  DaemonStatus_ProgramState = 4
	DaemonStatus_DS_ERROR    DaemonStatus_ProgramState = 5
	DaemonStatus_DS_CRITICAL DaemonStatus_ProgramState = 6
)

var DaemonStatus_ProgramState_name = map[int32]string{
	0: "DS_NONE",
	1: "DS_STARTING",
	2: "DS_READY",
	3: "DS_STOPPING",
	4: "DS_STOPPED",
	5: "DS_ERROR",
	6: "DS_CRITICAL",
}
var DaemonStatus_ProgramState_value = map[string]int32{
	"DS_NONE":     0,
	"DS_STARTING": 1,
	"DS_READY":    2,
	"DS_STOPPING": 3,
	"DS_STOPPED":  4,
	"DS_ERROR":    5,
	"DS_CRITICAL": 6,
}

func (x DaemonStatus_ProgramState) String() string {
	return proto.EnumName(DaemonStatus_ProgramState_name, int32(x))
}
func (DaemonStatus_ProgramState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

type RpcResult struct {
	// Whether or not the operation was successful.
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// If the operation was not successful, this is an explanatory message.
	Reason string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	// An HTTP-like status code.
	Code StatusCode `protobuf:"varint,3,opt,name=code,enum=common.v1.StatusCode" json:"code,omitempty"`
}

func (m *RpcResult) Reset()                    { *m = RpcResult{} }
func (m *RpcResult) String() string            { return proto.CompactTextString(m) }
func (*RpcResult) ProtoMessage()               {}
func (*RpcResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RpcResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RpcResult) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *RpcResult) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_STATUS_NONE
}

type DataplaneCommon struct {
	// UUID. An opaque string used by the controlplane as a unique ID.
	Uuid string `protobuf:"bytes,3,opt,name=uuid" json:"uuid,omitempty"`
	// The config 'revision', an opaque integer used by the controlplane to identify a configuration iteration.
	// This is deliberately matched to the revision field in etcd ResponseHeader.revision.
	Revision int64 `protobuf:"varint,4,opt,name=revision" json:"revision,omitempty"`
	// Latch for the versioning system. If true, updates will fail unless the generation of the change
	// is newer than the existing value.
	Versioned bool `protobuf:"varint,5,opt,name=versioned" json:"versioned,omitempty"`
}

func (m *DataplaneCommon) Reset()                    { *m = DataplaneCommon{} }
func (m *DataplaneCommon) String() string            { return proto.CompactTextString(m) }
func (*DataplaneCommon) ProtoMessage()               {}
func (*DataplaneCommon) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DataplaneCommon) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *DataplaneCommon) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *DataplaneCommon) GetVersioned() bool {
	if m != nil {
		return m.Versioned
	}
	return false
}

type Aes256Key struct {
	// 32 bytes (256 bits) of AES-256 key as raw bytes.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *Aes256Key) Reset()                    { *m = Aes256Key{} }
func (m *Aes256Key) String() string            { return proto.CompactTextString(m) }
func (*Aes256Key) ProtoMessage()               {}
func (*Aes256Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Aes256Key) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type Aes256IV struct {
	// 32 bytes (256 bits) of IV as raw bytes.
	Iv []byte `protobuf:"bytes,1,opt,name=iv,proto3" json:"iv,omitempty"`
}

func (m *Aes256IV) Reset()                    { *m = Aes256IV{} }
func (m *Aes256IV) String() string            { return proto.CompactTextString(m) }
func (*Aes256IV) ProtoMessage()               {}
func (*Aes256IV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Aes256IV) GetIv() []byte {
	if m != nil {
		return m.Iv
	}
	return nil
}

type XtsAes256Key struct {
	// 64 bytes (512 bits) of key data for XTS-AES 256 as raw bytes.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *XtsAes256Key) Reset()                    { *m = XtsAes256Key{} }
func (m *XtsAes256Key) String() string            { return proto.CompactTextString(m) }
func (*XtsAes256Key) ProtoMessage()               {}
func (*XtsAes256Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *XtsAes256Key) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type EncryptedXtsAesKey struct {
	// 64 bytes (512 bits) + up to 16 bytes of key data + padding for an
	// encrypted XTS-AES 256 key, as raw data.
	Ekey []byte `protobuf:"bytes,1,opt,name=ekey,proto3" json:"ekey,omitempty"`
}

func (m *EncryptedXtsAesKey) Reset()                    { *m = EncryptedXtsAesKey{} }
func (m *EncryptedXtsAesKey) String() string            { return proto.CompactTextString(m) }
func (*EncryptedXtsAesKey) ProtoMessage()               {}
func (*EncryptedXtsAesKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EncryptedXtsAesKey) GetEkey() []byte {
	if m != nil {
		return m.Ekey
	}
	return nil
}

type XtsKeyHMAC struct {
	// 64 bytes (512 bits) of HMAC data.
	Hmac []byte `protobuf:"bytes,1,opt,name=hmac,proto3" json:"hmac,omitempty"`
}

func (m *XtsKeyHMAC) Reset()                    { *m = XtsKeyHMAC{} }
func (m *XtsKeyHMAC) String() string            { return proto.CompactTextString(m) }
func (*XtsKeyHMAC) ProtoMessage()               {}
func (*XtsKeyHMAC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *XtsKeyHMAC) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

type VolumeKeySet struct {
	// The AES-XTS 256 volume master key.
	Vmk *XtsAes256Key `protobuf:"bytes,1,opt,name=vmk" json:"vmk,omitempty"`
}

func (m *VolumeKeySet) Reset()                    { *m = VolumeKeySet{} }
func (m *VolumeKeySet) String() string            { return proto.CompactTextString(m) }
func (*VolumeKeySet) ProtoMessage()               {}
func (*VolumeKeySet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *VolumeKeySet) GetVmk() *XtsAes256Key {
	if m != nil {
		return m.Vmk
	}
	return nil
}

type VolumeCrypto struct {
	// The crypto stepping version. Current version is 1.
	CryptoVersion uint32 `protobuf:"varint,2,opt,name=crypto_version,json=cryptoVersion" json:"crypto_version,omitempty"`
	// The current crypto state for this volume.
	CryptoState VolumeCrypto_VolumeCryptoState `protobuf:"varint,3,opt,name=crypto_state,json=cryptoState,enum=common.v1.VolumeCrypto_VolumeCryptoState" json:"crypto_state,omitempty"`
	// True iff the keyset contains valid key data.
	// Should only be set on create, and is likely to trigger validation errors
	// if set at other times.
	//
	// The dataplane servers will not reflect key data.
	KeysetPresent bool `protobuf:"varint,5,opt,name=keyset_present,json=keysetPresent" json:"keyset_present,omitempty"`
	// The set of keys required to use the volume.
	Keyset *VolumeKeySet `protobuf:"bytes,4,opt,name=keyset" json:"keyset,omitempty"`
}

func (m *VolumeCrypto) Reset()                    { *m = VolumeCrypto{} }
func (m *VolumeCrypto) String() string            { return proto.CompactTextString(m) }
func (*VolumeCrypto) ProtoMessage()               {}
func (*VolumeCrypto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *VolumeCrypto) GetCryptoVersion() uint32 {
	if m != nil {
		return m.CryptoVersion
	}
	return 0
}

func (m *VolumeCrypto) GetCryptoState() VolumeCrypto_VolumeCryptoState {
	if m != nil {
		return m.CryptoState
	}
	return VolumeCrypto_VCS_NONE
}

func (m *VolumeCrypto) GetKeysetPresent() bool {
	if m != nil {
		return m.KeysetPresent
	}
	return false
}

func (m *VolumeCrypto) GetKeyset() *VolumeKeySet {
	if m != nil {
		return m.Keyset
	}
	return nil
}

type DaemonStatus struct {
	ProgramState        DaemonStatus_ProgramState `protobuf:"varint,1,opt,name=program_state,json=programState,enum=common.v1.DaemonStatus_ProgramState" json:"program_state,omitempty"`
	ProgramStateMessage string                    `protobuf:"bytes,2,opt,name=program_state_message,json=programStateMessage" json:"program_state_message,omitempty"`
}

func (m *DaemonStatus) Reset()                    { *m = DaemonStatus{} }
func (m *DaemonStatus) String() string            { return proto.CompactTextString(m) }
func (*DaemonStatus) ProtoMessage()               {}
func (*DaemonStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DaemonStatus) GetProgramState() DaemonStatus_ProgramState {
	if m != nil {
		return m.ProgramState
	}
	return DaemonStatus_DS_NONE
}

func (m *DaemonStatus) GetProgramStateMessage() string {
	if m != nil {
		return m.ProgramStateMessage
	}
	return ""
}

type ConfigUpdateReply struct {
	Status ConfigUpdateStatus `protobuf:"varint,1,opt,name=status,enum=common.v1.ConfigUpdateStatus" json:"status,omitempty"`
	Result *RpcResult         `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *ConfigUpdateReply) Reset()                    { *m = ConfigUpdateReply{} }
func (m *ConfigUpdateReply) String() string            { return proto.CompactTextString(m) }
func (*ConfigUpdateReply) ProtoMessage()               {}
func (*ConfigUpdateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ConfigUpdateReply) GetStatus() ConfigUpdateStatus {
	if m != nil {
		return m.Status
	}
	return ConfigUpdateStatus_SET_NONE
}

func (m *ConfigUpdateReply) GetResult() *RpcResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type ConfigBool struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value bool   `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *ConfigBool) Reset()                    { *m = ConfigBool{} }
func (m *ConfigBool) String() string            { return proto.CompactTextString(m) }
func (*ConfigBool) ProtoMessage()               {}
func (*ConfigBool) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ConfigBool) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigBool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type ConfigGetBoolReply struct {
	Status ConfigGetStatus `protobuf:"varint,1,opt,name=status,enum=common.v1.ConfigGetStatus" json:"status,omitempty"`
	Result *RpcResult      `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Item   *ConfigBool     `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
}

func (m *ConfigGetBoolReply) Reset()                    { *m = ConfigGetBoolReply{} }
func (m *ConfigGetBoolReply) String() string            { return proto.CompactTextString(m) }
func (*ConfigGetBoolReply) ProtoMessage()               {}
func (*ConfigGetBoolReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ConfigGetBoolReply) GetStatus() ConfigGetStatus {
	if m != nil {
		return m.Status
	}
	return ConfigGetStatus_GET_NONE
}

func (m *ConfigGetBoolReply) GetResult() *RpcResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ConfigGetBoolReply) GetItem() *ConfigBool {
	if m != nil {
		return m.Item
	}
	return nil
}

type ConfigBoolList struct {
	Items []*ConfigBool `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ConfigBoolList) Reset()                    { *m = ConfigBoolList{} }
func (m *ConfigBoolList) String() string            { return proto.CompactTextString(m) }
func (*ConfigBoolList) ProtoMessage()               {}
func (*ConfigBoolList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ConfigBoolList) GetItems() []*ConfigBool {
	if m != nil {
		return m.Items
	}
	return nil
}

type ConfigUint32 struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value uint32 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *ConfigUint32) Reset()                    { *m = ConfigUint32{} }
func (m *ConfigUint32) String() string            { return proto.CompactTextString(m) }
func (*ConfigUint32) ProtoMessage()               {}
func (*ConfigUint32) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ConfigUint32) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigUint32) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ConfigGetUint32Reply struct {
	Status ConfigGetStatus `protobuf:"varint,1,opt,name=status,enum=common.v1.ConfigGetStatus" json:"status,omitempty"`
	Result *RpcResult      `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Item   *ConfigUint32   `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
}

func (m *ConfigGetUint32Reply) Reset()                    { *m = ConfigGetUint32Reply{} }
func (m *ConfigGetUint32Reply) String() string            { return proto.CompactTextString(m) }
func (*ConfigGetUint32Reply) ProtoMessage()               {}
func (*ConfigGetUint32Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ConfigGetUint32Reply) GetStatus() ConfigGetStatus {
	if m != nil {
		return m.Status
	}
	return ConfigGetStatus_GET_NONE
}

func (m *ConfigGetUint32Reply) GetResult() *RpcResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ConfigGetUint32Reply) GetItem() *ConfigUint32 {
	if m != nil {
		return m.Item
	}
	return nil
}

type ConfigUint32List struct {
	Items []*ConfigUint32 `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ConfigUint32List) Reset()                    { *m = ConfigUint32List{} }
func (m *ConfigUint32List) String() string            { return proto.CompactTextString(m) }
func (*ConfigUint32List) ProtoMessage()               {}
func (*ConfigUint32List) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *ConfigUint32List) GetItems() []*ConfigUint32 {
	if m != nil {
		return m.Items
	}
	return nil
}

type ConfigUint64 struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value uint64 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *ConfigUint64) Reset()                    { *m = ConfigUint64{} }
func (m *ConfigUint64) String() string            { return proto.CompactTextString(m) }
func (*ConfigUint64) ProtoMessage()               {}
func (*ConfigUint64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *ConfigUint64) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigUint64) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ConfigGetUint64Reply struct {
	Status ConfigGetStatus `protobuf:"varint,1,opt,name=status,enum=common.v1.ConfigGetStatus" json:"status,omitempty"`
	Result *RpcResult      `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Item   *ConfigUint64   `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
}

func (m *ConfigGetUint64Reply) Reset()                    { *m = ConfigGetUint64Reply{} }
func (m *ConfigGetUint64Reply) String() string            { return proto.CompactTextString(m) }
func (*ConfigGetUint64Reply) ProtoMessage()               {}
func (*ConfigGetUint64Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *ConfigGetUint64Reply) GetStatus() ConfigGetStatus {
	if m != nil {
		return m.Status
	}
	return ConfigGetStatus_GET_NONE
}

func (m *ConfigGetUint64Reply) GetResult() *RpcResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ConfigGetUint64Reply) GetItem() *ConfigUint64 {
	if m != nil {
		return m.Item
	}
	return nil
}

type ConfigUint64List struct {
	Items []*ConfigUint64 `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ConfigUint64List) Reset()                    { *m = ConfigUint64List{} }
func (m *ConfigUint64List) String() string            { return proto.CompactTextString(m) }
func (*ConfigUint64List) ProtoMessage()               {}
func (*ConfigUint64List) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *ConfigUint64List) GetItems() []*ConfigUint64 {
	if m != nil {
		return m.Items
	}
	return nil
}

type ConfigString struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ConfigString) Reset()                    { *m = ConfigString{} }
func (m *ConfigString) String() string            { return proto.CompactTextString(m) }
func (*ConfigString) ProtoMessage()               {}
func (*ConfigString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *ConfigString) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ConfigGetStringReply struct {
	Status ConfigGetStatus `protobuf:"varint,1,opt,name=status,enum=common.v1.ConfigGetStatus" json:"status,omitempty"`
	Result *RpcResult      `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Item   *ConfigString   `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
}

func (m *ConfigGetStringReply) Reset()                    { *m = ConfigGetStringReply{} }
func (m *ConfigGetStringReply) String() string            { return proto.CompactTextString(m) }
func (*ConfigGetStringReply) ProtoMessage()               {}
func (*ConfigGetStringReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *ConfigGetStringReply) GetStatus() ConfigGetStatus {
	if m != nil {
		return m.Status
	}
	return ConfigGetStatus_GET_NONE
}

func (m *ConfigGetStringReply) GetResult() *RpcResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ConfigGetStringReply) GetItem() *ConfigString {
	if m != nil {
		return m.Item
	}
	return nil
}

type ConfigStringList struct {
	Items []*ConfigString `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ConfigStringList) Reset()                    { *m = ConfigStringList{} }
func (m *ConfigStringList) String() string            { return proto.CompactTextString(m) }
func (*ConfigStringList) ProtoMessage()               {}
func (*ConfigStringList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *ConfigStringList) GetItems() []*ConfigString {
	if m != nil {
		return m.Items
	}
	return nil
}

type ConfigKey struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *ConfigKey) Reset()                    { *m = ConfigKey{} }
func (m *ConfigKey) String() string            { return proto.CompactTextString(m) }
func (*ConfigKey) ProtoMessage()               {}
func (*ConfigKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *ConfigKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type ConfigListQuery struct {
}

func (m *ConfigListQuery) Reset()                    { *m = ConfigListQuery{} }
func (m *ConfigListQuery) String() string            { return proto.CompactTextString(m) }
func (*ConfigListQuery) ProtoMessage()               {}
func (*ConfigListQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

func init() {
	proto.RegisterType((*RpcResult)(nil), "common.v1.RpcResult")
	proto.RegisterType((*DataplaneCommon)(nil), "common.v1.DataplaneCommon")
	proto.RegisterType((*Aes256Key)(nil), "common.v1.Aes256Key")
	proto.RegisterType((*Aes256IV)(nil), "common.v1.Aes256IV")
	proto.RegisterType((*XtsAes256Key)(nil), "common.v1.XtsAes256Key")
	proto.RegisterType((*EncryptedXtsAesKey)(nil), "common.v1.EncryptedXtsAesKey")
	proto.RegisterType((*XtsKeyHMAC)(nil), "common.v1.XtsKeyHMAC")
	proto.RegisterType((*VolumeKeySet)(nil), "common.v1.VolumeKeySet")
	proto.RegisterType((*VolumeCrypto)(nil), "common.v1.VolumeCrypto")
	proto.RegisterType((*DaemonStatus)(nil), "common.v1.DaemonStatus")
	proto.RegisterType((*ConfigUpdateReply)(nil), "common.v1.ConfigUpdateReply")
	proto.RegisterType((*ConfigBool)(nil), "common.v1.ConfigBool")
	proto.RegisterType((*ConfigGetBoolReply)(nil), "common.v1.ConfigGetBoolReply")
	proto.RegisterType((*ConfigBoolList)(nil), "common.v1.ConfigBoolList")
	proto.RegisterType((*ConfigUint32)(nil), "common.v1.ConfigUint32")
	proto.RegisterType((*ConfigGetUint32Reply)(nil), "common.v1.ConfigGetUint32Reply")
	proto.RegisterType((*ConfigUint32List)(nil), "common.v1.ConfigUint32List")
	proto.RegisterType((*ConfigUint64)(nil), "common.v1.ConfigUint64")
	proto.RegisterType((*ConfigGetUint64Reply)(nil), "common.v1.ConfigGetUint64Reply")
	proto.RegisterType((*ConfigUint64List)(nil), "common.v1.ConfigUint64List")
	proto.RegisterType((*ConfigString)(nil), "common.v1.ConfigString")
	proto.RegisterType((*ConfigGetStringReply)(nil), "common.v1.ConfigGetStringReply")
	proto.RegisterType((*ConfigStringList)(nil), "common.v1.ConfigStringList")
	proto.RegisterType((*ConfigKey)(nil), "common.v1.ConfigKey")
	proto.RegisterType((*ConfigListQuery)(nil), "common.v1.ConfigListQuery")
	proto.RegisterEnum("common.v1.StatusCode", StatusCode_name, StatusCode_value)
	proto.RegisterEnum("common.v1.ConfigGetStatus", ConfigGetStatus_name, ConfigGetStatus_value)
	proto.RegisterEnum("common.v1.ConfigUpdateStatus", ConfigUpdateStatus_name, ConfigUpdateStatus_value)
	proto.RegisterEnum("common.v1.VolumeCrypto_VolumeCryptoState", VolumeCrypto_VolumeCryptoState_name, VolumeCrypto_VolumeCryptoState_value)
	proto.RegisterEnum("common.v1.DaemonStatus_ProgramState", DaemonStatus_ProgramState_name, DaemonStatus_ProgramState_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1061 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x71, 0x9a, 0x6d, 0x4e, 0x93, 0xd4, 0x9d, 0xb6, 0x34, 0x5a, 0x51, 0x29, 0xb2, 0x40,
	0x6a, 0xbb, 0x50, 0x44, 0x36, 0x1b, 0x09, 0x24, 0x2e, 0x5c, 0xdb, 0xed, 0x5a, 0x4d, 0xe3, 0xee,
	0xd8, 0xa9, 0xb6, 0xdc, 0x58, 0x26, 0x19, 0xba, 0x51, 0x13, 0x3b, 0xb2, 0x9d, 0x88, 0xf0, 0x06,
	0xdc, 0x81, 0x40, 0x42, 0x5c, 0x72, 0xc1, 0xbb, 0x00, 0x6f, 0xc2, 0x3d, 0x0f, 0x80, 0xe6, 0xc7,
	0xe9, 0xa4, 0x3f, 0x74, 0xaf, 0xd0, 0xde, 0xcd, 0xf9, 0xce, 0x77, 0x8e, 0xbf, 0xf3, 0x4d, 0x66,
	0x26, 0x50, 0xe9, 0xc7, 0xe3, 0x71, 0x1c, 0x1d, 0x4e, 0x92, 0x38, 0x8b, 0x51, 0x59, 0x44, 0xb3,
	0xcf, 0xf4, 0x37, 0x50, 0xc6, 0x93, 0x3e, 0x26, 0xe9, 0x74, 0x94, 0xa1, 0x3a, 0x3c, 0x49, 0xa7,
	0xfd, 0x3e, 0x49, 0xd3, 0xba, 0xd2, 0x50, 0xf6, 0x56, 0x71, 0x1e, 0xa2, 0xf7, 0xa1, 0x94, 0x90,
	0x30, 0x8d, 0xa3, 0x7a, 0xa1, 0xa1, 0xec, 0x95, 0xb1, 0x88, 0xd0, 0x3e, 0x14, 0xfb, 0xf1, 0x80,
	0xd4, 0xd5, 0x86, 0xb2, 0x57, 0x6b, 0x6e, 0x1f, 0x2e, 0x1a, 0x1f, 0x7a, 0x59, 0x98, 0x4d, 0x53,
	0x33, 0x1e, 0x10, 0xcc, 0x28, 0x7a, 0x00, 0xeb, 0x56, 0x98, 0x85, 0x93, 0x51, 0x18, 0x11, 0x93,
	0xd1, 0x10, 0x82, 0xe2, 0x74, 0x3a, 0x1c, 0xb0, 0xea, 0x32, 0x66, 0x6b, 0xf4, 0x14, 0x56, 0x13,
	0x32, 0x1b, 0xa6, 0xc3, 0x38, 0xaa, 0x17, 0x1b, 0xca, 0x9e, 0x8a, 0x17, 0x31, 0xfa, 0x00, 0xca,
	0x33, 0x92, 0xd0, 0x25, 0x19, 0xd4, 0x57, 0x98, 0xc2, 0x1b, 0x40, 0xdf, 0x85, 0xb2, 0x41, 0xd2,
	0xe6, 0x8b, 0xf6, 0x29, 0x99, 0x23, 0x0d, 0xd4, 0x6b, 0x32, 0x67, 0x63, 0x54, 0x30, 0x5d, 0xea,
	0x4f, 0x61, 0x95, 0xa7, 0x9d, 0x0b, 0x54, 0x83, 0xc2, 0x70, 0x26, 0x92, 0x85, 0xe1, 0x4c, 0x6f,
	0x40, 0xe5, 0x75, 0x96, 0xfe, 0x57, 0xf5, 0x1e, 0x20, 0x3b, 0xea, 0x27, 0xf3, 0x49, 0x46, 0x06,
	0x9c, 0x4a, 0x79, 0x08, 0x8a, 0xe4, 0x86, 0xc8, 0xd6, 0x7a, 0x03, 0xe0, 0x75, 0x46, 0xb3, 0x2f,
	0xcf, 0x0c, 0x93, 0x32, 0xde, 0x8c, 0xc3, 0x7e, 0xce, 0xa0, 0x6b, 0xfd, 0x73, 0xa8, 0x5c, 0xc4,
	0xa3, 0xe9, 0x98, 0x9c, 0x92, 0xb9, 0x47, 0x32, 0xb4, 0x0f, 0xea, 0x6c, 0x7c, 0xcd, 0x28, 0x6b,
	0xcd, 0x1d, 0xc9, 0x43, 0x59, 0x13, 0xa6, 0x1c, 0xfd, 0x97, 0x42, 0x5e, 0x6b, 0x52, 0x2d, 0x31,
	0xfa, 0x08, 0x6a, 0x4c, 0x55, 0x1c, 0x08, 0x23, 0xd8, 0x06, 0x55, 0x71, 0x95, 0xa3, 0x17, 0x1c,
	0x44, 0x1d, 0xa8, 0x08, 0x5a, 0x9a, 0x85, 0x59, 0xbe, 0x5f, 0xfb, 0xd2, 0xb7, 0xe4, 0xae, 0x4b,
	0x01, 0xdd, 0x48, 0x82, 0xd7, 0xfa, 0x37, 0x01, 0xfd, 0xe8, 0x35, 0x99, 0xa7, 0x24, 0x0b, 0x26,
	0x09, 0x49, 0x49, 0x94, 0x89, 0xcd, 0xa8, 0x72, 0xf4, 0x9c, 0x83, 0xe8, 0x53, 0x28, 0x71, 0x80,
	0x6d, 0xe4, 0xf2, 0x68, 0xb2, 0x01, 0x58, 0xd0, 0xf4, 0x2f, 0x60, 0xe3, 0xce, 0x97, 0x51, 0x05,
	0x56, 0x2f, 0x4c, 0x2f, 0xe8, 0xba, 0x5d, 0x5b, 0x7b, 0x0f, 0xed, 0xc0, 0x26, 0x8d, 0x8e, 0x7b,
	0x9d, 0xce, 0x65, 0x60, 0x77, 0x4d, 0x7c, 0x79, 0xee, 0xdb, 0x96, 0xa6, 0xe8, 0xdf, 0x17, 0xa0,
	0x62, 0x85, 0x64, 0x1c, 0x47, 0xfc, 0x97, 0x87, 0x1c, 0xa8, 0x4e, 0x92, 0xf8, 0x2a, 0x09, 0xc7,
	0x62, 0x66, 0x85, 0xcd, 0xfc, 0xa1, 0x24, 0x42, 0xe6, 0x1f, 0x9e, 0x73, 0x32, 0x1f, 0xb7, 0x32,
	0x91, 0x22, 0xd4, 0x84, 0xed, 0xa5, 0x56, 0xc1, 0x98, 0xa4, 0x69, 0x78, 0x45, 0xc4, 0x61, 0xd8,
	0x94, 0xc9, 0x67, 0x3c, 0xa5, 0x7f, 0x07, 0x15, 0xb9, 0x23, 0x5a, 0x83, 0x27, 0xd6, 0x62, 0x8a,
	0x75, 0x58, 0xb3, 0xbc, 0xc0, 0xf3, 0x0d, 0xec, 0x3b, 0xdd, 0x13, 0x4d, 0xa1, 0x43, 0x5a, 0x5e,
	0x80, 0x6d, 0xc3, 0xba, 0xd4, 0x0a, 0x8b, 0xb4, 0x7b, 0x7e, 0x4e, 0xd3, 0x2a, 0xaa, 0x01, 0xe4,
	0x80, 0x6d, 0x69, 0x45, 0x41, 0xb7, 0x31, 0x76, 0xb1, 0xb6, 0x22, 0xe8, 0x26, 0x76, 0x7c, 0xc7,
	0x34, 0x3a, 0x5a, 0x49, 0xff, 0x16, 0x36, 0xcc, 0x38, 0xfa, 0x66, 0x78, 0xd5, 0x9b, 0x0c, 0xe8,
	0x34, 0x64, 0x32, 0x9a, 0xa3, 0x17, 0x50, 0x4a, 0xd9, 0xa4, 0xc2, 0x88, 0x5d, 0xc9, 0x08, 0x99,
	0xcd, 0xed, 0xc0, 0x82, 0x8c, 0x3e, 0xa6, 0x27, 0x9f, 0xde, 0x0e, 0x6c, 0xd8, 0xb5, 0xe6, 0x96,
	0x54, 0xb6, 0xb8, 0x39, 0xb0, 0xe0, 0xe8, 0x2d, 0x00, 0xde, 0xeb, 0x28, 0x8e, 0x47, 0xf2, 0x31,
	0x2a, 0xb3, 0x63, 0x84, 0xb6, 0x60, 0x65, 0x16, 0x8e, 0xa6, 0xdc, 0xb9, 0x55, 0xcc, 0x03, 0xfd,
	0x37, 0x05, 0x10, 0x2f, 0x3b, 0x21, 0x19, 0xad, 0xe4, 0x8a, 0x9b, 0xb7, 0x14, 0x3f, 0xbd, 0xa3,
	0xf8, 0x84, 0x64, 0x0f, 0xca, 0x55, 0x1f, 0x97, 0x4b, 0xaf, 0xaf, 0x61, 0x46, 0xc6, 0x62, 0xb4,
	0xed, 0x3b, 0xfd, 0x99, 0x16, 0x46, 0xd1, 0xbf, 0x84, 0xda, 0x0d, 0xd6, 0x19, 0xa6, 0x19, 0x7a,
	0x06, 0x2b, 0x34, 0x43, 0xd5, 0xa9, 0x0f, 0x57, 0x73, 0x8e, 0xde, 0x86, 0x8a, 0x30, 0x79, 0x18,
	0x65, 0xcf, 0x9b, 0x8f, 0x59, 0x53, 0xcd, 0xad, 0xf9, 0x5d, 0x81, 0xad, 0xc5, 0xac, 0xbc, 0xf6,
	0xff, 0x32, 0xe7, 0xd9, 0x92, 0x39, 0x3b, 0x77, 0x7f, 0x2e, 0x5c, 0x0d, 0xb7, 0xc7, 0x00, 0x4d,
	0x46, 0x99, 0x41, 0x9f, 0x2c, 0x1b, 0xf4, 0x60, 0x87, 0xfb, 0x2c, 0x6a, 0xb7, 0x1e, 0xb3, 0xa8,
	0xf8, 0xa0, 0x45, 0xed, 0xd6, 0xbb, 0x64, 0x51, 0xbb, 0x75, 0x9f, 0x45, 0xed, 0xd6, 0xdb, 0x5b,
	0xd4, 0x6e, 0xdd, 0xb1, 0xc8, 0xcb, 0x92, 0x61, 0x74, 0xf5, 0x98, 0x45, 0xe5, 0x7b, 0x2d, 0xe2,
	0xb5, 0xef, 0x8a, 0x45, 0x42, 0xcd, 0x2d, 0x8b, 0x38, 0xfa, 0x76, 0x16, 0x89, 0x0e, 0xc2, 0xa2,
	0x5d, 0x28, 0x73, 0xf8, 0xd6, 0x3b, 0xce, 0xfd, 0xd1, 0x37, 0x60, 0x9d, 0xa7, 0x69, 0xef, 0x57,
	0x53, 0x92, 0xcc, 0x0f, 0xfe, 0x56, 0x00, 0x6e, 0xfe, 0xad, 0xd0, 0xdb, 0xd4, 0xf3, 0x0d, 0xbf,
	0xb7, 0xb8, 0xac, 0x6b, 0x50, 0x16, 0x80, 0x7b, 0xaa, 0xfd, 0xa1, 0xa0, 0x2d, 0x58, 0x17, 0xb1,
	0x61, 0x9a, 0x36, 0x7b, 0x7e, 0xfe, 0x52, 0xd0, 0x26, 0xd4, 0x04, 0x6a, 0x62, 0xdb, 0xa0, 0xe0,
	0x9f, 0x0a, 0xda, 0x01, 0x24, 0xc0, 0x23, 0xc3, 0x0a, 0xb0, 0xfd, 0xaa, 0x67, 0x7b, 0xbe, 0xf6,
	0x83, 0x8a, 0xea, 0xb0, 0x29, 0x12, 0xbd, 0xae, 0xd1, 0xf3, 0x5f, 0xba, 0xd8, 0xf9, 0xca, 0xb6,
	0xb4, 0x1f, 0x55, 0xb4, 0x0d, 0x9a, 0xc8, 0x1c, 0xbb, 0xf8, 0xc8, 0xb1, 0x2c, 0xbb, 0xab, 0xfd,
	0x24, 0xc3, 0x5d, 0xd7, 0x0f, 0x8e, 0xdd, 0x5e, 0xd7, 0xd2, 0x7e, 0x56, 0x25, 0x2d, 0xa6, 0xdb,
	0x3d, 0xee, 0x38, 0xa6, 0xaf, 0xfd, 0x2a, 0x77, 0xf7, 0x6c, 0x7c, 0x61, 0x63, 0xf1, 0x52, 0xfc,
	0xa3, 0x1e, 0x74, 0xf2, 0xf1, 0x17, 0xdb, 0x4a, 0xdf, 0x92, 0x13, 0xdb, 0xcf, 0x87, 0x05, 0x28,
	0xd1, 0xc8, 0x3d, 0xd5, 0x14, 0xea, 0x04, 0x5d, 0x9f, 0x39, 0x9e, 0x47, 0x9f, 0xa1, 0x02, 0x7d,
	0x86, 0x28, 0x70, 0x6c, 0x38, 0x1d, 0xdb, 0xd2, 0xd4, 0x83, 0x41, 0x7e, 0x6d, 0xcb, 0x2f, 0x07,
	0x6d, 0xe8, 0x2d, 0x35, 0xf4, 0xf2, 0x86, 0x1b, 0x50, 0xa5, 0x6b, 0xe7, 0xec, 0xac, 0xe7, 0x1b,
	0x47, 0x1d, 0x9b, 0xb7, 0xf4, 0xa4, 0x96, 0x39, 0x85, 0xbe, 0x84, 0x81, 0xdb, 0xed, 0x5c, 0x6a,
	0xc5, 0xaf, 0x4b, 0xec, 0x4f, 0xeb, 0xf3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x43, 0x7b, 0x95,
	0x03, 0xc4, 0x0a, 0x00, 0x00,
}
