// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/service303.proto

package protos // import "magma/orc8r/cloud/go/protos"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServiceInfo_ServiceState int32

const (
	ServiceInfo_UNKNOWN  ServiceInfo_ServiceState = 0
	ServiceInfo_STARTING ServiceInfo_ServiceState = 1
	ServiceInfo_ALIVE    ServiceInfo_ServiceState = 2
	ServiceInfo_STOPPING ServiceInfo_ServiceState = 3
	ServiceInfo_STOPPED  ServiceInfo_ServiceState = 4
)

var ServiceInfo_ServiceState_name = map[int32]string{
	0: "UNKNOWN",
	1: "STARTING",
	2: "ALIVE",
	3: "STOPPING",
	4: "STOPPED",
}
var ServiceInfo_ServiceState_value = map[string]int32{
	"UNKNOWN":  0,
	"STARTING": 1,
	"ALIVE":    2,
	"STOPPING": 3,
	"STOPPED":  4,
}

func (x ServiceInfo_ServiceState) String() string {
	return proto.EnumName(ServiceInfo_ServiceState_name, int32(x))
}
func (ServiceInfo_ServiceState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service303_950751da80ee2d66, []int{2, 0}
}

// Gives information about whether the application is usable. Though the
// process may have started, the connections may not be set up. APP_HEALTHY in
// this case means the application is entirely usable
type ServiceInfo_ApplicationHealth int32

const (
	ServiceInfo_APP_UNKNOWN   ServiceInfo_ApplicationHealth = 0
	ServiceInfo_APP_UNHEALTHY ServiceInfo_ApplicationHealth = 1
	ServiceInfo_APP_HEALTHY   ServiceInfo_ApplicationHealth = 2
)

var ServiceInfo_ApplicationHealth_name = map[int32]string{
	0: "APP_UNKNOWN",
	1: "APP_UNHEALTHY",
	2: "APP_HEALTHY",
}
var ServiceInfo_ApplicationHealth_value = map[string]int32{
	"APP_UNKNOWN":   0,
	"APP_UNHEALTHY": 1,
	"APP_HEALTHY":   2,
}

func (x ServiceInfo_ApplicationHealth) String() string {
	return proto.EnumName(ServiceInfo_ApplicationHealth_name, int32(x))
}
func (ServiceInfo_ApplicationHealth) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service303_950751da80ee2d66, []int{2, 1}
}

type ReloadConfigResponse_ReloadConfigResult int32

const (
	ReloadConfigResponse_RELOAD_UNKNOWN     ReloadConfigResponse_ReloadConfigResult = 0
	ReloadConfigResponse_RELOAD_SUCCESS     ReloadConfigResponse_ReloadConfigResult = 1
	ReloadConfigResponse_RELOAD_FAILURE     ReloadConfigResponse_ReloadConfigResult = 2
	ReloadConfigResponse_RELOAD_UNSUPPORTED ReloadConfigResponse_ReloadConfigResult = 3
)

var ReloadConfigResponse_ReloadConfigResult_name = map[int32]string{
	0: "RELOAD_UNKNOWN",
	1: "RELOAD_SUCCESS",
	2: "RELOAD_FAILURE",
	3: "RELOAD_UNSUPPORTED",
}
var ReloadConfigResponse_ReloadConfigResult_value = map[string]int32{
	"RELOAD_UNKNOWN":     0,
	"RELOAD_SUCCESS":     1,
	"RELOAD_FAILURE":     2,
	"RELOAD_UNSUPPORTED": 3,
}

func (x ReloadConfigResponse_ReloadConfigResult) String() string {
	return proto.EnumName(ReloadConfigResponse_ReloadConfigResult_name, int32(x))
}
func (ReloadConfigResponse_ReloadConfigResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service303_950751da80ee2d66, []int{5, 0}
}

type EnodebdStatus struct {
	// For bools, parameter is not set if result can't be determined
	EnodebConnected      *wrappers.BoolValue  `protobuf:"bytes,1,opt,name=enodeb_connected,json=enodebConnected,proto3" json:"enodeb_connected,omitempty"`
	OpstateEnabled       *wrappers.BoolValue  `protobuf:"bytes,2,opt,name=opstate_enabled,json=opstateEnabled,proto3" json:"opstate_enabled,omitempty"`
	RfTxOn               *wrappers.BoolValue  `protobuf:"bytes,3,opt,name=rf_tx_on,json=rfTxOn,proto3" json:"rf_tx_on,omitempty"`
	GpsConnected         *wrappers.BoolValue  `protobuf:"bytes,4,opt,name=gps_connected,json=gpsConnected,proto3" json:"gps_connected,omitempty"`
	PtpConnected         *wrappers.BoolValue  `protobuf:"bytes,5,opt,name=ptp_connected,json=ptpConnected,proto3" json:"ptp_connected,omitempty"`
	MmeConnected         *wrappers.BoolValue  `protobuf:"bytes,6,opt,name=mme_connected,json=mmeConnected,proto3" json:"mme_connected,omitempty"`
	EnodebConfigured     *wrappers.BoolValue  `protobuf:"bytes,7,opt,name=enodeb_configured,json=enodebConfigured,proto3" json:"enodeb_configured,omitempty"`
	GpsLatitude          *wrappers.FloatValue `protobuf:"bytes,8,opt,name=gps_latitude,json=gpsLatitude,proto3" json:"gps_latitude,omitempty"`
	GpsLongitude         *wrappers.FloatValue `protobuf:"bytes,9,opt,name=gps_longitude,json=gpsLongitude,proto3" json:"gps_longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EnodebdStatus) Reset()         { *m = EnodebdStatus{} }
func (m *EnodebdStatus) String() string { return proto.CompactTextString(m) }
func (*EnodebdStatus) ProtoMessage()    {}
func (*EnodebdStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_service303_950751da80ee2d66, []int{0}
}
func (m *EnodebdStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnodebdStatus.Unmarshal(m, b)
}
func (m *EnodebdStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnodebdStatus.Marshal(b, m, deterministic)
}
func (dst *EnodebdStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnodebdStatus.Merge(dst, src)
}
func (m *EnodebdStatus) XXX_Size() int {
	return xxx_messageInfo_EnodebdStatus.Size(m)
}
func (m *EnodebdStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EnodebdStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EnodebdStatus proto.InternalMessageInfo

func (m *EnodebdStatus) GetEnodebConnected() *wrappers.BoolValue {
	if m != nil {
		return m.EnodebConnected
	}
	return nil
}

func (m *EnodebdStatus) GetOpstateEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.OpstateEnabled
	}
	return nil
}

func (m *EnodebdStatus) GetRfTxOn() *wrappers.BoolValue {
	if m != nil {
		return m.RfTxOn
	}
	return nil
}

func (m *EnodebdStatus) GetGpsConnected() *wrappers.BoolValue {
	if m != nil {
		return m.GpsConnected
	}
	return nil
}

func (m *EnodebdStatus) GetPtpConnected() *wrappers.BoolValue {
	if m != nil {
		return m.PtpConnected
	}
	return nil
}

func (m *EnodebdStatus) GetMmeConnected() *wrappers.BoolValue {
	if m != nil {
		return m.MmeConnected
	}
	return nil
}

func (m *EnodebdStatus) GetEnodebConfigured() *wrappers.BoolValue {
	if m != nil {
		return m.EnodebConfigured
	}
	return nil
}

func (m *EnodebdStatus) GetGpsLatitude() *wrappers.FloatValue {
	if m != nil {
		return m.GpsLatitude
	}
	return nil
}

func (m *EnodebdStatus) GetGpsLongitude() *wrappers.FloatValue {
	if m != nil {
		return m.GpsLongitude
	}
	return nil
}

type ServiceStatus struct {
	// Metadata from the services that will be sent to the cloud through checkin
	Meta                 map[string]string `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServiceStatus) Reset()         { *m = ServiceStatus{} }
func (m *ServiceStatus) String() string { return proto.CompactTextString(m) }
func (*ServiceStatus) ProtoMessage()    {}
func (*ServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_service303_950751da80ee2d66, []int{1}
}
func (m *ServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceStatus.Unmarshal(m, b)
}
func (m *ServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceStatus.Marshal(b, m, deterministic)
}
func (dst *ServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceStatus.Merge(dst, src)
}
func (m *ServiceStatus) XXX_Size() int {
	return xxx_messageInfo_ServiceStatus.Size(m)
}
func (m *ServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceStatus proto.InternalMessageInfo

func (m *ServiceStatus) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

type ServiceInfo struct {
	Name    string                        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string                        `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	State   ServiceInfo_ServiceState      `protobuf:"varint,3,opt,name=state,proto3,enum=magma.orc8r.ServiceInfo_ServiceState" json:"state,omitempty"`
	Status  *ServiceStatus                `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Health  ServiceInfo_ApplicationHealth `protobuf:"varint,5,opt,name=health,proto3,enum=magma.orc8r.ServiceInfo_ApplicationHealth" json:"health,omitempty"`
	// Time when the service was started (in seconds since epoch)
	StartTimeSecs        uint64   `protobuf:"varint,6,opt,name=start_time_secs,json=startTimeSecs,proto3" json:"start_time_secs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceInfo) Reset()         { *m = ServiceInfo{} }
func (m *ServiceInfo) String() string { return proto.CompactTextString(m) }
func (*ServiceInfo) ProtoMessage()    {}
func (*ServiceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_service303_950751da80ee2d66, []int{2}
}
func (m *ServiceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInfo.Unmarshal(m, b)
}
func (m *ServiceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInfo.Marshal(b, m, deterministic)
}
func (dst *ServiceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInfo.Merge(dst, src)
}
func (m *ServiceInfo) XXX_Size() int {
	return xxx_messageInfo_ServiceInfo.Size(m)
}
func (m *ServiceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInfo proto.InternalMessageInfo

func (m *ServiceInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceInfo) GetState() ServiceInfo_ServiceState {
	if m != nil {
		return m.State
	}
	return ServiceInfo_UNKNOWN
}

func (m *ServiceInfo) GetStatus() *ServiceStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ServiceInfo) GetHealth() ServiceInfo_ApplicationHealth {
	if m != nil {
		return m.Health
	}
	return ServiceInfo_APP_UNKNOWN
}

func (m *ServiceInfo) GetStartTimeSecs() uint64 {
	if m != nil {
		return m.StartTimeSecs
	}
	return 0
}

type LogLevelMessage struct {
	Level                LogLevel `protobuf:"varint,1,opt,name=level,proto3,enum=magma.orc8r.LogLevel" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevelMessage) Reset()         { *m = LogLevelMessage{} }
func (m *LogLevelMessage) String() string { return proto.CompactTextString(m) }
func (*LogLevelMessage) ProtoMessage()    {}
func (*LogLevelMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_service303_950751da80ee2d66, []int{3}
}
func (m *LogLevelMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevelMessage.Unmarshal(m, b)
}
func (m *LogLevelMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevelMessage.Marshal(b, m, deterministic)
}
func (dst *LogLevelMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevelMessage.Merge(dst, src)
}
func (m *LogLevelMessage) XXX_Size() int {
	return xxx_messageInfo_LogLevelMessage.Size(m)
}
func (m *LogLevelMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevelMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevelMessage proto.InternalMessageInfo

func (m *LogLevelMessage) GetLevel() LogLevel {
	if m != nil {
		return m.Level
	}
	return LogLevel_DEBUG
}

type LogVerbosity struct {
	Verbosity            int32    `protobuf:"varint,1,opt,name=verbosity,proto3" json:"verbosity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogVerbosity) Reset()         { *m = LogVerbosity{} }
func (m *LogVerbosity) String() string { return proto.CompactTextString(m) }
func (*LogVerbosity) ProtoMessage()    {}
func (*LogVerbosity) Descriptor() ([]byte, []int) {
	return fileDescriptor_service303_950751da80ee2d66, []int{4}
}
func (m *LogVerbosity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogVerbosity.Unmarshal(m, b)
}
func (m *LogVerbosity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogVerbosity.Marshal(b, m, deterministic)
}
func (dst *LogVerbosity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogVerbosity.Merge(dst, src)
}
func (m *LogVerbosity) XXX_Size() int {
	return xxx_messageInfo_LogVerbosity.Size(m)
}
func (m *LogVerbosity) XXX_DiscardUnknown() {
	xxx_messageInfo_LogVerbosity.DiscardUnknown(m)
}

var xxx_messageInfo_LogVerbosity proto.InternalMessageInfo

func (m *LogVerbosity) GetVerbosity() int32 {
	if m != nil {
		return m.Verbosity
	}
	return 0
}

type ReloadConfigResponse struct {
	Result               ReloadConfigResponse_ReloadConfigResult `protobuf:"varint,1,opt,name=result,proto3,enum=magma.orc8r.ReloadConfigResponse_ReloadConfigResult" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ReloadConfigResponse) Reset()         { *m = ReloadConfigResponse{} }
func (m *ReloadConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ReloadConfigResponse) ProtoMessage()    {}
func (*ReloadConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service303_950751da80ee2d66, []int{5}
}
func (m *ReloadConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadConfigResponse.Unmarshal(m, b)
}
func (m *ReloadConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadConfigResponse.Marshal(b, m, deterministic)
}
func (dst *ReloadConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadConfigResponse.Merge(dst, src)
}
func (m *ReloadConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ReloadConfigResponse.Size(m)
}
func (m *ReloadConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadConfigResponse proto.InternalMessageInfo

func (m *ReloadConfigResponse) GetResult() ReloadConfigResponse_ReloadConfigResult {
	if m != nil {
		return m.Result
	}
	return ReloadConfigResponse_RELOAD_UNKNOWN
}

type State struct {
	// Type determines how the value is deserialized and validated on the cloud service side
	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	DeviceID string `protobuf:"bytes,2,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	// Value contains the operational state json-serialized.
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_service303_950751da80ee2d66, []int{6}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *State) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *State) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetOperationalStatesResponse struct {
	States               []*State `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperationalStatesResponse) Reset()         { *m = GetOperationalStatesResponse{} }
func (m *GetOperationalStatesResponse) String() string { return proto.CompactTextString(m) }
func (*GetOperationalStatesResponse) ProtoMessage()    {}
func (*GetOperationalStatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service303_950751da80ee2d66, []int{7}
}
func (m *GetOperationalStatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationalStatesResponse.Unmarshal(m, b)
}
func (m *GetOperationalStatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationalStatesResponse.Marshal(b, m, deterministic)
}
func (dst *GetOperationalStatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationalStatesResponse.Merge(dst, src)
}
func (m *GetOperationalStatesResponse) XXX_Size() int {
	return xxx_messageInfo_GetOperationalStatesResponse.Size(m)
}
func (m *GetOperationalStatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationalStatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationalStatesResponse proto.InternalMessageInfo

func (m *GetOperationalStatesResponse) GetStates() []*State {
	if m != nil {
		return m.States
	}
	return nil
}

func init() {
	proto.RegisterType((*EnodebdStatus)(nil), "magma.orc8r.EnodebdStatus")
	proto.RegisterType((*ServiceStatus)(nil), "magma.orc8r.ServiceStatus")
	proto.RegisterMapType((map[string]string)(nil), "magma.orc8r.ServiceStatus.MetaEntry")
	proto.RegisterType((*ServiceInfo)(nil), "magma.orc8r.ServiceInfo")
	proto.RegisterType((*LogLevelMessage)(nil), "magma.orc8r.LogLevelMessage")
	proto.RegisterType((*LogVerbosity)(nil), "magma.orc8r.LogVerbosity")
	proto.RegisterType((*ReloadConfigResponse)(nil), "magma.orc8r.ReloadConfigResponse")
	proto.RegisterType((*State)(nil), "magma.orc8r.State")
	proto.RegisterType((*GetOperationalStatesResponse)(nil), "magma.orc8r.GetOperationalStatesResponse")
	proto.RegisterEnum("magma.orc8r.ServiceInfo_ServiceState", ServiceInfo_ServiceState_name, ServiceInfo_ServiceState_value)
	proto.RegisterEnum("magma.orc8r.ServiceInfo_ApplicationHealth", ServiceInfo_ApplicationHealth_name, ServiceInfo_ApplicationHealth_value)
	proto.RegisterEnum("magma.orc8r.ReloadConfigResponse_ReloadConfigResult", ReloadConfigResponse_ReloadConfigResult_name, ReloadConfigResponse_ReloadConfigResult_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Service303Client is the client API for Service303 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Service303Client interface {
	// Returns the service level info like name, version, state, status, etc.
	//
	GetServiceInfo(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ServiceInfo, error)
	// Request to stop the service gracefully.
	//
	StopService(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	// Collects metrics from the service
	GetMetrics(ctx context.Context, in *Void, opts ...grpc.CallOption) (*MetricsContainer, error)
	// Set logging level
	SetLogLevel(ctx context.Context, in *LogLevelMessage, opts ...grpc.CallOption) (*Void, error)
	// Set logging verbosity The larger, the more verbose. default 0
	SetLogVerbosity(ctx context.Context, in *LogVerbosity, opts ...grpc.CallOption) (*Void, error)
	// Requests service reloads config files loaded on startup (<servicename>.yml)
	ReloadServiceConfig(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ReloadConfigResponse, error)
	// Returns the  operational states of devices managed by this service.
	GetOperationalStates(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetOperationalStatesResponse, error)
}

type service303Client struct {
	cc *grpc.ClientConn
}

func NewService303Client(cc *grpc.ClientConn) Service303Client {
	return &service303Client{cc}
}

func (c *service303Client) GetServiceInfo(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ServiceInfo, error) {
	out := new(ServiceInfo)
	err := c.cc.Invoke(ctx, "/magma.orc8r.Service303/GetServiceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service303Client) StopService(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.Service303/StopService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service303Client) GetMetrics(ctx context.Context, in *Void, opts ...grpc.CallOption) (*MetricsContainer, error) {
	out := new(MetricsContainer)
	err := c.cc.Invoke(ctx, "/magma.orc8r.Service303/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service303Client) SetLogLevel(ctx context.Context, in *LogLevelMessage, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.Service303/SetLogLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service303Client) SetLogVerbosity(ctx context.Context, in *LogVerbosity, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.Service303/SetLogVerbosity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service303Client) ReloadServiceConfig(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ReloadConfigResponse, error) {
	out := new(ReloadConfigResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.Service303/ReloadServiceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service303Client) GetOperationalStates(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetOperationalStatesResponse, error) {
	out := new(GetOperationalStatesResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.Service303/GetOperationalStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service303Server is the server API for Service303 service.
type Service303Server interface {
	// Returns the service level info like name, version, state, status, etc.
	//
	GetServiceInfo(context.Context, *Void) (*ServiceInfo, error)
	// Request to stop the service gracefully.
	//
	StopService(context.Context, *Void) (*Void, error)
	// Collects metrics from the service
	GetMetrics(context.Context, *Void) (*MetricsContainer, error)
	// Set logging level
	SetLogLevel(context.Context, *LogLevelMessage) (*Void, error)
	// Set logging verbosity The larger, the more verbose. default 0
	SetLogVerbosity(context.Context, *LogVerbosity) (*Void, error)
	// Requests service reloads config files loaded on startup (<servicename>.yml)
	ReloadServiceConfig(context.Context, *Void) (*ReloadConfigResponse, error)
	// Returns the  operational states of devices managed by this service.
	GetOperationalStates(context.Context, *Void) (*GetOperationalStatesResponse, error)
}

func RegisterService303Server(s *grpc.Server, srv Service303Server) {
	s.RegisterService(&_Service303_serviceDesc, srv)
}

func _Service303_GetServiceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service303Server).GetServiceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.Service303/GetServiceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service303Server).GetServiceInfo(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service303_StopService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service303Server).StopService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.Service303/StopService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service303Server).StopService(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service303_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service303Server).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.Service303/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service303Server).GetMetrics(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service303_SetLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLevelMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service303Server).SetLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.Service303/SetLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service303Server).SetLogLevel(ctx, req.(*LogLevelMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service303_SetLogVerbosity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogVerbosity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service303Server).SetLogVerbosity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.Service303/SetLogVerbosity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service303Server).SetLogVerbosity(ctx, req.(*LogVerbosity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service303_ReloadServiceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service303Server).ReloadServiceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.Service303/ReloadServiceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service303Server).ReloadServiceConfig(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service303_GetOperationalStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service303Server).GetOperationalStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.Service303/GetOperationalStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service303Server).GetOperationalStates(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service303_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.Service303",
	HandlerType: (*Service303Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceInfo",
			Handler:    _Service303_GetServiceInfo_Handler,
		},
		{
			MethodName: "StopService",
			Handler:    _Service303_StopService_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _Service303_GetMetrics_Handler,
		},
		{
			MethodName: "SetLogLevel",
			Handler:    _Service303_SetLogLevel_Handler,
		},
		{
			MethodName: "SetLogVerbosity",
			Handler:    _Service303_SetLogVerbosity_Handler,
		},
		{
			MethodName: "ReloadServiceConfig",
			Handler:    _Service303_ReloadServiceConfig_Handler,
		},
		{
			MethodName: "GetOperationalStates",
			Handler:    _Service303_GetOperationalStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/service303.proto",
}

func init() {
	proto.RegisterFile("orc8r/protos/service303.proto", fileDescriptor_service303_950751da80ee2d66)
}

var fileDescriptor_service303_950751da80ee2d66 = []byte{
	// 977 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0xed, 0x6e, 0xe3, 0x44,
	0x14, 0x6d, 0x9a, 0x8f, 0x6d, 0x6e, 0x9a, 0x8f, 0x0e, 0x05, 0x65, 0xd3, 0x16, 0x2d, 0x16, 0xa0,
	0x65, 0x41, 0x29, 0x4a, 0x17, 0x51, 0x81, 0xb4, 0xdd, 0xb4, 0xf5, 0xb6, 0x05, 0xb7, 0x89, 0xec,
	0xb4, 0x08, 0xfe, 0x44, 0xae, 0x7d, 0xeb, 0x5a, 0xd8, 0x1e, 0xcb, 0x33, 0x09, 0xdb, 0xbf, 0xbc,
	0x12, 0xef, 0xc1, 0x2b, 0xf0, 0x0e, 0x3c, 0x01, 0xf2, 0xcc, 0xe4, 0xc3, 0x5b, 0x97, 0xfc, 0xca,
	0xcc, 0xb9, 0xf7, 0x9c, 0xb9, 0x33, 0xf7, 0x5c, 0xc5, 0xb0, 0x47, 0x13, 0xe7, 0x30, 0xd9, 0x8f,
	0x13, 0xca, 0x29, 0xdb, 0x67, 0x98, 0x4c, 0x7d, 0x07, 0x0f, 0xbe, 0x3d, 0xe8, 0x0a, 0x84, 0xd4,
	0x42, 0xdb, 0x0b, 0xed, 0xae, 0x48, 0xea, 0x3c, 0xcf, 0xe4, 0x3a, 0x34, 0x0c, 0x69, 0x24, 0xf3,
	0x3a, 0x3b, 0x99, 0x50, 0x88, 0x3c, 0xf1, 0x1d, 0xe6, 0xaa, 0xe0, 0xa7, 0x1e, 0xa5, 0x5e, 0x80,
	0x32, 0x7a, 0x3b, 0xb9, 0xdb, 0xff, 0x23, 0xb1, 0xe3, 0x18, 0x13, 0x26, 0xe3, 0xda, 0x3f, 0x25,
	0xa8, 0xeb, 0x11, 0x75, 0xf1, 0xd6, 0xb5, 0xb8, 0xcd, 0x27, 0x8c, 0xe8, 0xd0, 0x42, 0x01, 0x8c,
	0x1d, 0x1a, 0x45, 0xe8, 0x70, 0x74, 0xdb, 0x85, 0x17, 0x85, 0x97, 0xb5, 0x5e, 0xa7, 0x2b, 0xc5,
	0xba, 0x33, 0xb1, 0xee, 0x31, 0xa5, 0xc1, 0x8d, 0x1d, 0x4c, 0xd0, 0x6c, 0x4a, 0xce, 0xc9, 0x8c,
	0x42, 0x4e, 0xa0, 0x49, 0x63, 0xc6, 0x6d, 0x8e, 0x63, 0x8c, 0xec, 0xdb, 0x00, 0xdd, 0xf6, 0xfa,
	0x4a, 0x95, 0x86, 0xa2, 0xe8, 0x92, 0x41, 0x5e, 0xc3, 0x46, 0x72, 0x37, 0xe6, 0xef, 0xc7, 0x34,
	0x6a, 0x17, 0x57, 0xb2, 0x2b, 0xc9, 0xdd, 0xe8, 0xfd, 0x20, 0x22, 0x47, 0x50, 0xf7, 0x62, 0xb6,
	0x54, 0x7e, 0x69, 0x25, 0x75, 0xd3, 0x8b, 0xd9, 0xa2, 0xf6, 0x23, 0xa8, 0xc7, 0x3c, 0x5e, 0x12,
	0x28, 0xaf, 0x16, 0x88, 0x79, 0x9c, 0x11, 0x08, 0x43, 0x5c, 0x12, 0xa8, 0xac, 0x16, 0x08, 0x43,
	0x5c, 0x08, 0x9c, 0xc1, 0xd6, 0xa2, 0x09, 0x77, 0xbe, 0x37, 0x49, 0xd0, 0x6d, 0x3f, 0x5b, 0x29,
	0xd2, 0x9a, 0x77, 0x41, 0x71, 0xc8, 0x1b, 0x48, 0xaf, 0x36, 0x0e, 0x6c, 0xee, 0xf3, 0x89, 0x8b,
	0xed, 0x0d, 0xa1, 0xb1, 0xf3, 0x48, 0xe3, 0x5d, 0x40, 0x6d, 0x2e, 0x45, 0x6a, 0x5e, 0xcc, 0x0c,
	0x95, 0x4f, 0xde, 0xca, 0xb7, 0x0c, 0x68, 0xe4, 0x49, 0x81, 0xea, 0x6a, 0x81, 0xf4, 0x44, 0x63,
	0x46, 0xd0, 0xfe, 0x2c, 0x40, 0xdd, 0x92, 0xde, 0x56, 0x0e, 0x3b, 0x84, 0x52, 0x88, 0xdc, 0x6e,
	0xaf, 0xbf, 0x28, 0xbe, 0xac, 0xf5, 0x3e, 0xef, 0x2e, 0xf9, 0xbc, 0x9b, 0xc9, 0xec, 0x5e, 0x22,
	0xb7, 0xf5, 0x88, 0x27, 0x0f, 0xa6, 0x60, 0x74, 0xbe, 0x87, 0xea, 0x1c, 0x22, 0x2d, 0x28, 0xfe,
	0x8e, 0x0f, 0xc2, 0x9b, 0x55, 0x33, 0x5d, 0x92, 0x6d, 0x28, 0x4f, 0xd3, 0x0a, 0x84, 0xd3, 0xaa,
	0xa6, 0xdc, 0xfc, 0xb0, 0x7e, 0x58, 0xd0, 0xfe, 0x2a, 0x42, 0x4d, 0x49, 0x5f, 0x44, 0x77, 0x94,
	0x10, 0x28, 0x45, 0x76, 0x88, 0x8a, 0x2c, 0xd6, 0xa4, 0x0d, 0xcf, 0xa6, 0x98, 0x30, 0x9f, 0x46,
	0x8a, 0x3f, 0xdb, 0x92, 0x1f, 0xa1, 0x2c, 0x6c, 0x29, 0x3c, 0xd8, 0xe8, 0x7d, 0x91, 0x57, 0x71,
	0x2a, 0xbb, 0x5c, 0x3d, 0x9a, 0x92, 0x43, 0x7a, 0x50, 0x61, 0xe2, 0x36, 0x73, 0x1b, 0x3e, 0x79,
	0x5f, 0x53, 0x65, 0x92, 0x63, 0xa8, 0xdc, 0xa3, 0x1d, 0xf0, 0x7b, 0xe1, 0xbc, 0x46, 0xef, 0xd5,
	0x93, 0x27, 0xf6, 0xe3, 0x38, 0xf0, 0x1d, 0x9b, 0xfb, 0x34, 0x3a, 0x17, 0x0c, 0x53, 0x31, 0xc9,
	0x97, 0xd0, 0x64, 0xdc, 0x4e, 0xf8, 0x98, 0xfb, 0x21, 0x8e, 0x19, 0x3a, 0x4c, 0xb8, 0xb0, 0x64,
	0xd6, 0x05, 0x3c, 0xf2, 0x43, 0xb4, 0xd0, 0x61, 0xda, 0x00, 0x36, 0x97, 0xcb, 0x26, 0x35, 0x78,
	0x76, 0x7d, 0xf5, 0xf3, 0xd5, 0xe0, 0x97, 0xab, 0xd6, 0x1a, 0xd9, 0x84, 0x0d, 0x6b, 0xd4, 0x37,
	0x47, 0x17, 0x57, 0x67, 0xad, 0x02, 0xa9, 0x42, 0xb9, 0x6f, 0x5c, 0xdc, 0xe8, 0xad, 0x75, 0x19,
	0x18, 0x0c, 0x87, 0x69, 0xa0, 0x98, 0x72, 0xc4, 0x4e, 0x3f, 0x6d, 0x95, 0xb4, 0x73, 0xd8, 0x7a,
	0x54, 0x15, 0x69, 0x42, 0xad, 0x3f, 0x1c, 0x8e, 0x17, 0xca, 0x5b, 0x50, 0x97, 0xc0, 0xb9, 0xde,
	0x37, 0x46, 0xe7, 0xbf, 0xb6, 0x0a, 0xb3, 0x9c, 0x19, 0xb0, 0xae, 0xbd, 0x81, 0xa6, 0x41, 0x3d,
	0x03, 0xa7, 0x18, 0x5c, 0x22, 0x63, 0xb6, 0x87, 0xe4, 0x6b, 0x28, 0x07, 0xe9, 0x5e, 0x74, 0xae,
	0xd1, 0xfb, 0x38, 0xf3, 0x30, 0xb3, 0x64, 0x53, 0xe6, 0x68, 0xdf, 0xc0, 0xa6, 0x41, 0xbd, 0x1b,
	0x4c, 0x6e, 0x29, 0xf3, 0xf9, 0x03, 0xd9, 0x85, 0xea, 0x74, 0xb6, 0x11, 0x02, 0x65, 0x73, 0x01,
	0x68, 0x7f, 0x17, 0x60, 0xdb, 0xc4, 0x80, 0xda, 0xae, 0x9c, 0x1f, 0x13, 0x59, 0x4c, 0x23, 0x86,
	0xc4, 0x80, 0x4a, 0x82, 0x6c, 0x12, 0x70, 0x75, 0xe8, 0xeb, 0xcc, 0xa1, 0x79, 0x94, 0x0f, 0xc1,
	0x49, 0xc0, 0x4d, 0xa5, 0xa1, 0xdd, 0x03, 0x79, 0x1c, 0x25, 0x04, 0x1a, 0xa6, 0x6e, 0x0c, 0xfa,
	0xa7, 0x4b, 0x4f, 0xb4, 0xc0, 0xac, 0xeb, 0x93, 0x13, 0xdd, 0xb2, 0x5a, 0x85, 0x25, 0xec, 0x5d,
	0xff, 0xc2, 0xb8, 0x36, 0xd3, 0x5e, 0x7c, 0x02, 0x64, 0xce, 0xb5, 0xae, 0x87, 0xc3, 0x81, 0x39,
	0xd2, 0x4f, 0x5b, 0x45, 0xed, 0x12, 0xca, 0xb2, 0xa5, 0x04, 0x4a, 0xfc, 0x21, 0x9e, 0xbb, 0x3d,
	0x5d, 0x93, 0x0e, 0x6c, 0xb8, 0x28, 0x6c, 0x74, 0xaa, 0xec, 0x3e, 0xdf, 0x2f, 0xe6, 0x28, 0xf5,
	0xfb, 0xa6, 0x9a, 0x23, 0xed, 0x27, 0xd8, 0x3d, 0x43, 0x3e, 0x88, 0x31, 0x11, 0x8d, 0xb5, 0x03,
	0x21, 0xce, 0xe6, 0xcf, 0xf4, 0x4a, 0x1a, 0x1d, 0x59, 0xbb, 0x20, 0x06, 0x9b, 0x64, 0x4d, 0x2b,
	0x66, 0x42, 0x65, 0xf4, 0xfe, 0x2d, 0x02, 0x58, 0xf3, 0x3f, 0x3c, 0x72, 0x04, 0x8d, 0x33, 0xe4,
	0xcb, 0x03, 0xba, 0x95, 0x21, 0xdf, 0x50, 0xdf, 0xed, 0xb4, 0x9f, 0x1a, 0x02, 0x6d, 0x8d, 0x7c,
	0x07, 0x35, 0x8b, 0xd3, 0x58, 0x81, 0x79, 0xec, 0xc7, 0x90, 0xb6, 0x46, 0xde, 0x02, 0x9c, 0x21,
	0xbf, 0x94, 0x7f, 0x99, 0x79, 0xac, 0xbd, 0x0c, 0xa4, 0x12, 0x4f, 0x68, 0xc4, 0x6d, 0x3f, 0xc2,
	0x44, 0x28, 0xd4, 0x2c, 0xe4, 0x33, 0xe3, 0x91, 0xdd, 0x5c, 0x3f, 0x2a, 0xf3, 0xe6, 0xd7, 0xd0,
	0x87, 0xa6, 0x54, 0x58, 0xf8, 0xf4, 0xf9, 0x87, 0x2a, 0xf3, 0x50, 0xbe, 0xc4, 0x25, 0x7c, 0x24,
	0x2d, 0xa5, 0xee, 0x2f, 0x9d, 0x95, 0x77, 0x9f, 0xcf, 0x56, 0x5a, 0x57, 0x5b, 0x23, 0x37, 0xb0,
	0x9d, 0xd7, 0xe8, 0x3c, 0xbd, 0xaf, 0x32, 0xd0, 0xff, 0xd9, 0x43, 0x5b, 0x3b, 0xde, 0xfb, 0x6d,
	0x47, 0x64, 0xef, 0xcb, 0x0f, 0x16, 0x27, 0xa0, 0x13, 0x77, 0xdf, 0xa3, 0xea, 0xcb, 0xe5, 0xb6,
	0x22, 0x7e, 0x0f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xde, 0x4c, 0x45, 0xb1, 0x17, 0x09, 0x00,
	0x00,
}
