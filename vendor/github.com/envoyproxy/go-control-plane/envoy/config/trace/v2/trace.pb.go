// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v2/trace.proto

package envoy_config_trace_v2

import (
	fmt "fmt"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type LightstepConfig_PropagationMode int32

const (
	LightstepConfig_ENVOY         LightstepConfig_PropagationMode = 0
	LightstepConfig_LIGHTSTEP     LightstepConfig_PropagationMode = 1
	LightstepConfig_B3            LightstepConfig_PropagationMode = 2
	LightstepConfig_TRACE_CONTEXT LightstepConfig_PropagationMode = 3
)

var LightstepConfig_PropagationMode_name = map[int32]string{
	0: "ENVOY",
	1: "LIGHTSTEP",
	2: "B3",
	3: "TRACE_CONTEXT",
}

var LightstepConfig_PropagationMode_value = map[string]int32{
	"ENVOY":         0,
	"LIGHTSTEP":     1,
	"B3":            2,
	"TRACE_CONTEXT": 3,
}

func (x LightstepConfig_PropagationMode) String() string {
	return proto.EnumName(LightstepConfig_PropagationMode_name, int32(x))
}

func (LightstepConfig_PropagationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{1, 0}
}

type ZipkinConfig_CollectorEndpointVersion int32

const (
	ZipkinConfig_HTTP_JSON_V1 ZipkinConfig_CollectorEndpointVersion = 0 // Deprecated: Do not use.
	ZipkinConfig_HTTP_JSON    ZipkinConfig_CollectorEndpointVersion = 1
	ZipkinConfig_HTTP_PROTO   ZipkinConfig_CollectorEndpointVersion = 2
	ZipkinConfig_GRPC         ZipkinConfig_CollectorEndpointVersion = 3
)

var ZipkinConfig_CollectorEndpointVersion_name = map[int32]string{
	0: "HTTP_JSON_V1",
	1: "HTTP_JSON",
	2: "HTTP_PROTO",
	3: "GRPC",
}

var ZipkinConfig_CollectorEndpointVersion_value = map[string]int32{
	"HTTP_JSON_V1": 0,
	"HTTP_JSON":    1,
	"HTTP_PROTO":   2,
	"GRPC":         3,
}

func (x ZipkinConfig_CollectorEndpointVersion) String() string {
	return proto.EnumName(ZipkinConfig_CollectorEndpointVersion_name, int32(x))
}

func (ZipkinConfig_CollectorEndpointVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{2, 0}
}

type OpenCensusConfig_TraceContext int32

const (
	OpenCensusConfig_NONE                OpenCensusConfig_TraceContext = 0
	OpenCensusConfig_TRACE_CONTEXT       OpenCensusConfig_TraceContext = 1
	OpenCensusConfig_GRPC_TRACE_BIN      OpenCensusConfig_TraceContext = 2
	OpenCensusConfig_CLOUD_TRACE_CONTEXT OpenCensusConfig_TraceContext = 3
	OpenCensusConfig_B3                  OpenCensusConfig_TraceContext = 4
)

var OpenCensusConfig_TraceContext_name = map[int32]string{
	0: "NONE",
	1: "TRACE_CONTEXT",
	2: "GRPC_TRACE_BIN",
	3: "CLOUD_TRACE_CONTEXT",
	4: "B3",
}

var OpenCensusConfig_TraceContext_value = map[string]int32{
	"NONE":                0,
	"TRACE_CONTEXT":       1,
	"GRPC_TRACE_BIN":      2,
	"CLOUD_TRACE_CONTEXT": 3,
	"B3":                  4,
}

func (x OpenCensusConfig_TraceContext) String() string {
	return proto.EnumName(OpenCensusConfig_TraceContext_name, int32(x))
}

func (OpenCensusConfig_TraceContext) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{5, 0}
}

type Tracing struct {
	Http                 *Tracing_Http `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Tracing) Reset()         { *m = Tracing{} }
func (m *Tracing) String() string { return proto.CompactTextString(m) }
func (*Tracing) ProtoMessage()    {}
func (*Tracing) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{0}
}

func (m *Tracing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing.Unmarshal(m, b)
}
func (m *Tracing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing.Marshal(b, m, deterministic)
}
func (m *Tracing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing.Merge(m, src)
}
func (m *Tracing) XXX_Size() int {
	return xxx_messageInfo_Tracing.Size(m)
}
func (m *Tracing) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing proto.InternalMessageInfo

func (m *Tracing) GetHttp() *Tracing_Http {
	if m != nil {
		return m.Http
	}
	return nil
}

type Tracing_Http struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*Tracing_Http_Config
	//	*Tracing_Http_TypedConfig
	ConfigType           isTracing_Http_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Tracing_Http) Reset()         { *m = Tracing_Http{} }
func (m *Tracing_Http) String() string { return proto.CompactTextString(m) }
func (*Tracing_Http) ProtoMessage()    {}
func (*Tracing_Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{0, 0}
}

func (m *Tracing_Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing_Http.Unmarshal(m, b)
}
func (m *Tracing_Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing_Http.Marshal(b, m, deterministic)
}
func (m *Tracing_Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing_Http.Merge(m, src)
}
func (m *Tracing_Http) XXX_Size() int {
	return xxx_messageInfo_Tracing_Http.Size(m)
}
func (m *Tracing_Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing_Http proto.InternalMessageInfo

func (m *Tracing_Http) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTracing_Http_ConfigType interface {
	isTracing_Http_ConfigType()
}

type Tracing_Http_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type Tracing_Http_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Tracing_Http_Config) isTracing_Http_ConfigType() {}

func (*Tracing_Http_TypedConfig) isTracing_Http_ConfigType() {}

func (m *Tracing_Http) GetConfigType() isTracing_Http_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *Tracing_Http) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*Tracing_Http_Config); ok {
		return x.Config
	}
	return nil
}

func (m *Tracing_Http) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Tracing_Http_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Tracing_Http) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Tracing_Http_Config)(nil),
		(*Tracing_Http_TypedConfig)(nil),
	}
}

type LightstepConfig struct {
	CollectorCluster     string                            `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	AccessTokenFile      string                            `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile,proto3" json:"access_token_file,omitempty"`
	PropagationModes     []LightstepConfig_PropagationMode `protobuf:"varint,3,rep,packed,name=propagation_modes,json=propagationModes,proto3,enum=envoy.config.trace.v2.LightstepConfig_PropagationMode" json:"propagation_modes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *LightstepConfig) Reset()         { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()    {}
func (*LightstepConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{1}
}

func (m *LightstepConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightstepConfig.Unmarshal(m, b)
}
func (m *LightstepConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightstepConfig.Marshal(b, m, deterministic)
}
func (m *LightstepConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightstepConfig.Merge(m, src)
}
func (m *LightstepConfig) XXX_Size() int {
	return xxx_messageInfo_LightstepConfig.Size(m)
}
func (m *LightstepConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LightstepConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LightstepConfig proto.InternalMessageInfo

func (m *LightstepConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *LightstepConfig) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

func (m *LightstepConfig) GetPropagationModes() []LightstepConfig_PropagationMode {
	if m != nil {
		return m.PropagationModes
	}
	return nil
}

type ZipkinConfig struct {
	CollectorCluster         string                                `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	CollectorEndpoint        string                                `protobuf:"bytes,2,opt,name=collector_endpoint,json=collectorEndpoint,proto3" json:"collector_endpoint,omitempty"`
	TraceId_128Bit           bool                                  `protobuf:"varint,3,opt,name=trace_id_128bit,json=traceId128bit,proto3" json:"trace_id_128bit,omitempty"`
	SharedSpanContext        *wrappers.BoolValue                   `protobuf:"bytes,4,opt,name=shared_span_context,json=sharedSpanContext,proto3" json:"shared_span_context,omitempty"`
	CollectorEndpointVersion ZipkinConfig_CollectorEndpointVersion `protobuf:"varint,5,opt,name=collector_endpoint_version,json=collectorEndpointVersion,proto3,enum=envoy.config.trace.v2.ZipkinConfig_CollectorEndpointVersion" json:"collector_endpoint_version,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                              `json:"-"`
	XXX_unrecognized         []byte                                `json:"-"`
	XXX_sizecache            int32                                 `json:"-"`
}

func (m *ZipkinConfig) Reset()         { *m = ZipkinConfig{} }
func (m *ZipkinConfig) String() string { return proto.CompactTextString(m) }
func (*ZipkinConfig) ProtoMessage()    {}
func (*ZipkinConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{2}
}

func (m *ZipkinConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZipkinConfig.Unmarshal(m, b)
}
func (m *ZipkinConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZipkinConfig.Marshal(b, m, deterministic)
}
func (m *ZipkinConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZipkinConfig.Merge(m, src)
}
func (m *ZipkinConfig) XXX_Size() int {
	return xxx_messageInfo_ZipkinConfig.Size(m)
}
func (m *ZipkinConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ZipkinConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ZipkinConfig proto.InternalMessageInfo

func (m *ZipkinConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *ZipkinConfig) GetCollectorEndpoint() string {
	if m != nil {
		return m.CollectorEndpoint
	}
	return ""
}

func (m *ZipkinConfig) GetTraceId_128Bit() bool {
	if m != nil {
		return m.TraceId_128Bit
	}
	return false
}

func (m *ZipkinConfig) GetSharedSpanContext() *wrappers.BoolValue {
	if m != nil {
		return m.SharedSpanContext
	}
	return nil
}

func (m *ZipkinConfig) GetCollectorEndpointVersion() ZipkinConfig_CollectorEndpointVersion {
	if m != nil {
		return m.CollectorEndpointVersion
	}
	return ZipkinConfig_HTTP_JSON_V1
}

type DynamicOtConfig struct {
	Library              string          `protobuf:"bytes,1,opt,name=library,proto3" json:"library,omitempty"`
	Config               *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DynamicOtConfig) Reset()         { *m = DynamicOtConfig{} }
func (m *DynamicOtConfig) String() string { return proto.CompactTextString(m) }
func (*DynamicOtConfig) ProtoMessage()    {}
func (*DynamicOtConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{3}
}

func (m *DynamicOtConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicOtConfig.Unmarshal(m, b)
}
func (m *DynamicOtConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicOtConfig.Marshal(b, m, deterministic)
}
func (m *DynamicOtConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicOtConfig.Merge(m, src)
}
func (m *DynamicOtConfig) XXX_Size() int {
	return xxx_messageInfo_DynamicOtConfig.Size(m)
}
func (m *DynamicOtConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicOtConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicOtConfig proto.InternalMessageInfo

func (m *DynamicOtConfig) GetLibrary() string {
	if m != nil {
		return m.Library
	}
	return ""
}

func (m *DynamicOtConfig) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

type DatadogConfig struct {
	CollectorCluster     string   `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatadogConfig) Reset()         { *m = DatadogConfig{} }
func (m *DatadogConfig) String() string { return proto.CompactTextString(m) }
func (*DatadogConfig) ProtoMessage()    {}
func (*DatadogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{4}
}

func (m *DatadogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatadogConfig.Unmarshal(m, b)
}
func (m *DatadogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatadogConfig.Marshal(b, m, deterministic)
}
func (m *DatadogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatadogConfig.Merge(m, src)
}
func (m *DatadogConfig) XXX_Size() int {
	return xxx_messageInfo_DatadogConfig.Size(m)
}
func (m *DatadogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DatadogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DatadogConfig proto.InternalMessageInfo

func (m *DatadogConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *DatadogConfig) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type OpenCensusConfig struct {
	TraceConfig                *v1.TraceConfig                 `protobuf:"bytes,1,opt,name=trace_config,json=traceConfig,proto3" json:"trace_config,omitempty"`
	StdoutExporterEnabled      bool                            `protobuf:"varint,2,opt,name=stdout_exporter_enabled,json=stdoutExporterEnabled,proto3" json:"stdout_exporter_enabled,omitempty"`
	StackdriverExporterEnabled bool                            `protobuf:"varint,3,opt,name=stackdriver_exporter_enabled,json=stackdriverExporterEnabled,proto3" json:"stackdriver_exporter_enabled,omitempty"`
	StackdriverProjectId       string                          `protobuf:"bytes,4,opt,name=stackdriver_project_id,json=stackdriverProjectId,proto3" json:"stackdriver_project_id,omitempty"`
	StackdriverAddress         string                          `protobuf:"bytes,10,opt,name=stackdriver_address,json=stackdriverAddress,proto3" json:"stackdriver_address,omitempty"`
	StackdriverGrpcService     *core.GrpcService               `protobuf:"bytes,13,opt,name=stackdriver_grpc_service,json=stackdriverGrpcService,proto3" json:"stackdriver_grpc_service,omitempty"`
	ZipkinExporterEnabled      bool                            `protobuf:"varint,5,opt,name=zipkin_exporter_enabled,json=zipkinExporterEnabled,proto3" json:"zipkin_exporter_enabled,omitempty"`
	ZipkinUrl                  string                          `protobuf:"bytes,6,opt,name=zipkin_url,json=zipkinUrl,proto3" json:"zipkin_url,omitempty"`
	OcagentExporterEnabled     bool                            `protobuf:"varint,11,opt,name=ocagent_exporter_enabled,json=ocagentExporterEnabled,proto3" json:"ocagent_exporter_enabled,omitempty"`
	OcagentAddress             string                          `protobuf:"bytes,12,opt,name=ocagent_address,json=ocagentAddress,proto3" json:"ocagent_address,omitempty"`
	OcagentGrpcService         *core.GrpcService               `protobuf:"bytes,14,opt,name=ocagent_grpc_service,json=ocagentGrpcService,proto3" json:"ocagent_grpc_service,omitempty"`
	IncomingTraceContext       []OpenCensusConfig_TraceContext `protobuf:"varint,8,rep,packed,name=incoming_trace_context,json=incomingTraceContext,proto3,enum=envoy.config.trace.v2.OpenCensusConfig_TraceContext" json:"incoming_trace_context,omitempty"`
	OutgoingTraceContext       []OpenCensusConfig_TraceContext `protobuf:"varint,9,rep,packed,name=outgoing_trace_context,json=outgoingTraceContext,proto3,enum=envoy.config.trace.v2.OpenCensusConfig_TraceContext" json:"outgoing_trace_context,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                        `json:"-"`
	XXX_unrecognized           []byte                          `json:"-"`
	XXX_sizecache              int32                           `json:"-"`
}

func (m *OpenCensusConfig) Reset()         { *m = OpenCensusConfig{} }
func (m *OpenCensusConfig) String() string { return proto.CompactTextString(m) }
func (*OpenCensusConfig) ProtoMessage()    {}
func (*OpenCensusConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{5}
}

func (m *OpenCensusConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenCensusConfig.Unmarshal(m, b)
}
func (m *OpenCensusConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenCensusConfig.Marshal(b, m, deterministic)
}
func (m *OpenCensusConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenCensusConfig.Merge(m, src)
}
func (m *OpenCensusConfig) XXX_Size() int {
	return xxx_messageInfo_OpenCensusConfig.Size(m)
}
func (m *OpenCensusConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenCensusConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OpenCensusConfig proto.InternalMessageInfo

func (m *OpenCensusConfig) GetTraceConfig() *v1.TraceConfig {
	if m != nil {
		return m.TraceConfig
	}
	return nil
}

func (m *OpenCensusConfig) GetStdoutExporterEnabled() bool {
	if m != nil {
		return m.StdoutExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetStackdriverExporterEnabled() bool {
	if m != nil {
		return m.StackdriverExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetStackdriverProjectId() string {
	if m != nil {
		return m.StackdriverProjectId
	}
	return ""
}

func (m *OpenCensusConfig) GetStackdriverAddress() string {
	if m != nil {
		return m.StackdriverAddress
	}
	return ""
}

func (m *OpenCensusConfig) GetStackdriverGrpcService() *core.GrpcService {
	if m != nil {
		return m.StackdriverGrpcService
	}
	return nil
}

func (m *OpenCensusConfig) GetZipkinExporterEnabled() bool {
	if m != nil {
		return m.ZipkinExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetZipkinUrl() string {
	if m != nil {
		return m.ZipkinUrl
	}
	return ""
}

func (m *OpenCensusConfig) GetOcagentExporterEnabled() bool {
	if m != nil {
		return m.OcagentExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetOcagentAddress() string {
	if m != nil {
		return m.OcagentAddress
	}
	return ""
}

func (m *OpenCensusConfig) GetOcagentGrpcService() *core.GrpcService {
	if m != nil {
		return m.OcagentGrpcService
	}
	return nil
}

func (m *OpenCensusConfig) GetIncomingTraceContext() []OpenCensusConfig_TraceContext {
	if m != nil {
		return m.IncomingTraceContext
	}
	return nil
}

func (m *OpenCensusConfig) GetOutgoingTraceContext() []OpenCensusConfig_TraceContext {
	if m != nil {
		return m.OutgoingTraceContext
	}
	return nil
}

type TraceServiceConfig struct {
	GrpcService          *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TraceServiceConfig) Reset()         { *m = TraceServiceConfig{} }
func (m *TraceServiceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceServiceConfig) ProtoMessage()    {}
func (*TraceServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{6}
}

func (m *TraceServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceServiceConfig.Unmarshal(m, b)
}
func (m *TraceServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceServiceConfig.Marshal(b, m, deterministic)
}
func (m *TraceServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceServiceConfig.Merge(m, src)
}
func (m *TraceServiceConfig) XXX_Size() int {
	return xxx_messageInfo_TraceServiceConfig.Size(m)
}
func (m *TraceServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceServiceConfig proto.InternalMessageInfo

func (m *TraceServiceConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.trace.v2.LightstepConfig_PropagationMode", LightstepConfig_PropagationMode_name, LightstepConfig_PropagationMode_value)
	proto.RegisterEnum("envoy.config.trace.v2.ZipkinConfig_CollectorEndpointVersion", ZipkinConfig_CollectorEndpointVersion_name, ZipkinConfig_CollectorEndpointVersion_value)
	proto.RegisterEnum("envoy.config.trace.v2.OpenCensusConfig_TraceContext", OpenCensusConfig_TraceContext_name, OpenCensusConfig_TraceContext_value)
	proto.RegisterType((*Tracing)(nil), "envoy.config.trace.v2.Tracing")
	proto.RegisterType((*Tracing_Http)(nil), "envoy.config.trace.v2.Tracing.Http")
	proto.RegisterType((*LightstepConfig)(nil), "envoy.config.trace.v2.LightstepConfig")
	proto.RegisterType((*ZipkinConfig)(nil), "envoy.config.trace.v2.ZipkinConfig")
	proto.RegisterType((*DynamicOtConfig)(nil), "envoy.config.trace.v2.DynamicOtConfig")
	proto.RegisterType((*DatadogConfig)(nil), "envoy.config.trace.v2.DatadogConfig")
	proto.RegisterType((*OpenCensusConfig)(nil), "envoy.config.trace.v2.OpenCensusConfig")
	proto.RegisterType((*TraceServiceConfig)(nil), "envoy.config.trace.v2.TraceServiceConfig")
}

func init() { proto.RegisterFile("envoy/config/trace/v2/trace.proto", fileDescriptor_0785d24fc8ab55c7) }

var fileDescriptor_0785d24fc8ab55c7 = []byte{
	// 1178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x41, 0x73, 0xdb, 0x44,
	0x14, 0x8e, 0x1c, 0x27, 0x71, 0x9e, 0xed, 0x44, 0xd9, 0xa6, 0x8d, 0x30, 0x6d, 0x27, 0x75, 0x99,
	0xd2, 0x61, 0x18, 0x79, 0xe2, 0x96, 0x52, 0x06, 0x0e, 0x54, 0xae, 0xdb, 0xa4, 0x14, 0xdb, 0xa3,
	0xb8, 0x99, 0xc2, 0x45, 0x6c, 0xa4, 0xad, 0xaa, 0x56, 0xd1, 0x2e, 0xab, 0xb5, 0xa9, 0x7b, 0x62,
	0x38, 0xc2, 0x8d, 0xdf, 0xc0, 0x81, 0x1f, 0xc0, 0x81, 0xe1, 0x17, 0x70, 0xe1, 0xc0, 0x7f, 0xe0,
	0xc2, 0x2f, 0x60, 0xa6, 0x27, 0x46, 0xbb, 0xab, 0x44, 0x91, 0x9b, 0x99, 0xce, 0xf4, 0x26, 0xbd,
	0xef, 0xfb, 0xde, 0xbe, 0xfd, 0xf6, 0xed, 0x5b, 0xb8, 0x42, 0x92, 0x29, 0x9d, 0x75, 0x7c, 0x9a,
	0x3c, 0x89, 0xc2, 0x8e, 0xe0, 0xd8, 0x27, 0x9d, 0x69, 0x57, 0x7d, 0xd8, 0x8c, 0x53, 0x41, 0xd1,
	0x79, 0x49, 0xb1, 0x15, 0xc5, 0x56, 0xc8, 0xb4, 0xdb, 0x7a, 0x4f, 0x29, 0x31, 0x8b, 0x32, 0x81,
	0x4f, 0x39, 0xe9, 0x84, 0x9c, 0xf9, 0x5e, 0x4a, 0xf8, 0x34, 0xca, 0xc5, 0xad, 0x77, 0x42, 0x4a,
	0xc3, 0x98, 0x74, 0xe4, 0xdf, 0xe1, 0xe4, 0x49, 0x07, 0x27, 0x33, 0x0d, 0x5d, 0x2c, 0x43, 0xa9,
	0xe0, 0x13, 0x5f, 0x68, 0xf4, 0x72, 0x19, 0xfd, 0x8e, 0x63, 0xc6, 0x08, 0x4f, 0x35, 0xfe, 0x21,
	0x65, 0x24, 0xf1, 0x49, 0x92, 0x4e, 0x52, 0xc5, 0xc9, 0x8b, 0xdf, 0x51, 0x1f, 0x9e, 0xae, 0x57,
	0xb1, 0xaf, 0xea, 0x62, 0x93, 0x84, 0x0a, 0x2c, 0x22, 0x9a, 0xa4, 0x9d, 0x80, 0x30, 0x4e, 0x7c,
	0xf9, 0xa3, 0x49, 0x97, 0x26, 0x01, 0xc3, 0xa7, 0x38, 0xa9, 0xc0, 0x62, 0x92, 0xaf, 0xb8, 0x35,
	0xc5, 0x71, 0x14, 0x60, 0x41, 0x3a, 0xf9, 0x87, 0x02, 0xda, 0xff, 0x1a, 0xb0, 0x32, 0xe6, 0xd8,
	0x8f, 0x92, 0x10, 0x7d, 0x0c, 0xd5, 0xa7, 0x42, 0x30, 0xcb, 0xd8, 0x36, 0xae, 0xd7, 0xbb, 0x57,
	0xed, 0xd7, 0x7a, 0x67, 0x6b, 0xb6, 0xbd, 0x2b, 0x04, 0x73, 0xa5, 0xa0, 0xf5, 0x8b, 0x01, 0xd5,
	0xec, 0x17, 0xbd, 0x0b, 0xd5, 0x04, 0x1f, 0x11, 0x99, 0x61, 0xd5, 0x59, 0x79, 0xe5, 0x54, 0x79,
	0x65, 0xdb, 0x70, 0x65, 0x10, 0x7d, 0x04, 0xcb, 0x2a, 0x97, 0x55, 0x91, 0x0b, 0x6c, 0xd9, 0xca,
	0x26, 0x3b, 0xb7, 0xc9, 0xde, 0x97, 0x26, 0x3a, 0x15, 0xcb, 0xd8, 0x5d, 0x70, 0x35, 0x19, 0x7d,
	0x02, 0x0d, 0x31, 0x63, 0x24, 0xd0, 0xa6, 0x58, 0x8b, 0x52, 0xbc, 0x39, 0x27, 0xbe, 0x93, 0xcc,
	0x76, 0x17, 0xdc, 0xba, 0xe4, 0xf6, 0x24, 0xd5, 0x69, 0x42, 0x5d, 0x89, 0xbc, 0x2c, 0xda, 0xfe,
	0xbd, 0x02, 0xeb, 0x0f, 0xa3, 0xf0, 0xa9, 0x48, 0x05, 0x61, 0x8a, 0x82, 0x6e, 0xc2, 0x86, 0x4f,
	0xe3, 0x98, 0xf8, 0x82, 0x72, 0xcf, 0x8f, 0x27, 0xa9, 0x20, 0xbc, 0x5c, 0xbe, 0x79, 0xcc, 0xe8,
	0x29, 0x02, 0xba, 0x01, 0x1b, 0xd8, 0xf7, 0x49, 0x9a, 0x7a, 0x82, 0x3e, 0x27, 0x89, 0xf7, 0x24,
	0x8a, 0x89, 0xdc, 0x55, 0x41, 0xb5, 0xae, 0x18, 0xe3, 0x8c, 0x70, 0x2f, 0x8a, 0x09, 0xe2, 0xb0,
	0xc1, 0x38, 0x65, 0x38, 0x94, 0x07, 0xe4, 0x1d, 0xd1, 0x80, 0xa4, 0xd6, 0xe2, 0xf6, 0xe2, 0xf5,
	0xb5, 0xee, 0xad, 0x33, 0xbc, 0x2e, 0x55, 0x6b, 0x8f, 0x4e, 0xf4, 0x5f, 0xd2, 0x80, 0x38, 0xcd,
	0x57, 0x0e, 0xfc, 0x6c, 0xac, 0xb4, 0x97, 0x7e, 0x30, 0x2a, 0xa6, 0xe1, 0x9a, 0xec, 0x34, 0x9e,
	0xb6, 0xef, 0xc1, 0x7a, 0x49, 0x83, 0x56, 0x61, 0xa9, 0x3f, 0x38, 0x18, 0x7e, 0x65, 0x2e, 0xa0,
	0x26, 0xac, 0x3e, 0xdc, 0xbb, 0xbf, 0x3b, 0xde, 0x1f, 0xf7, 0x47, 0xa6, 0x81, 0x96, 0xa1, 0xe2,
	0xdc, 0x30, 0x2b, 0x68, 0x03, 0x9a, 0x63, 0xf7, 0x4e, 0xaf, 0xef, 0xf5, 0x86, 0x83, 0x71, 0xff,
	0xf1, 0xd8, 0x5c, 0x6c, 0xff, 0xb5, 0x08, 0x8d, 0xaf, 0x23, 0xf6, 0x3c, 0x4a, 0xde, 0xca, 0xb7,
	0x5b, 0x80, 0x4e, 0x54, 0x24, 0x09, 0x18, 0x8d, 0x12, 0x51, 0x36, 0xee, 0x24, 0x71, 0x5f, 0x33,
	0xd0, 0x35, 0x58, 0x57, 0x17, 0x23, 0x0a, 0xbc, 0x9d, 0xee, 0xed, 0xc3, 0x48, 0xc8, 0x36, 0xa8,
	0xb9, 0x4d, 0x19, 0xde, 0x0b, 0x54, 0x10, 0x3d, 0x80, 0x73, 0xe9, 0x53, 0xcc, 0x49, 0xe0, 0xa5,
	0x0c, 0x27, 0x59, 0xc7, 0x08, 0xf2, 0x42, 0x58, 0x55, 0xd9, 0x32, 0xad, 0xb9, 0x96, 0x71, 0x28,
	0x8d, 0x0f, 0x70, 0x3c, 0x21, 0xee, 0x86, 0x92, 0xed, 0x33, 0x9c, 0x6d, 0x30, 0x13, 0xa1, 0x97,
	0xd0, 0x9a, 0xaf, 0xd5, 0x9b, 0x12, 0x9e, 0x46, 0x34, 0xb1, 0x96, 0xb6, 0x8d, 0xeb, 0x6b, 0xdd,
	0xcf, 0xce, 0x38, 0xb7, 0xa2, 0x55, 0x76, 0xaf, 0xbc, 0x9d, 0x03, 0x95, 0xc3, 0xb5, 0xfc, 0x33,
	0x90, 0xb6, 0x07, 0xd6, 0x59, 0x2a, 0xd4, 0x82, 0xc6, 0xee, 0x78, 0x3c, 0xf2, 0x1e, 0xec, 0x0f,
	0x07, 0xde, 0xc1, 0x8e, 0xb9, 0xd0, 0xaa, 0xfd, 0xfa, 0xdf, 0x6f, 0x3f, 0x55, 0x8c, 0x9a, 0x91,
	0x1d, 0xe8, 0x31, 0x66, 0x1a, 0x68, 0x0d, 0x40, 0xfe, 0x8e, 0xdc, 0xe1, 0x78, 0x68, 0x56, 0x50,
	0x0d, 0xaa, 0xf7, 0xdd, 0x51, 0xcf, 0x5c, 0x6c, 0x13, 0x58, 0xbf, 0x3b, 0x4b, 0xf0, 0x51, 0xe4,
	0x0f, 0x85, 0x3e, 0xd1, 0x2b, 0xb0, 0x12, 0x47, 0x87, 0x1c, 0xf3, 0x59, 0xf9, 0x1c, 0xf3, 0x38,
	0xea, 0xbc, 0xe1, 0x0d, 0xce, 0xef, 0x6e, 0xfb, 0x5b, 0x68, 0xde, 0xc5, 0x02, 0x07, 0x34, 0x7c,
	0xab, 0xb6, 0xf9, 0x00, 0x1a, 0x7a, 0x32, 0x7b, 0x72, 0xbc, 0x94, 0x1a, 0xa6, 0xae, 0xc1, 0x01,
	0x3e, 0x22, 0xed, 0x7f, 0x56, 0xc0, 0x1c, 0x32, 0x92, 0xf4, 0xe4, 0x78, 0xd5, 0xcb, 0xee, 0x41,
	0xa3, 0x38, 0x58, 0xf5, 0x84, 0xbb, 0x66, 0x9f, 0xcc, 0x61, 0xb5, 0x85, 0xfc, 0x04, 0x77, 0xe4,
	0x94, 0x23, 0x4a, 0xed, 0xd6, 0xc5, 0xc9, 0x0f, 0xba, 0x05, 0x5b, 0xa9, 0x08, 0xe8, 0x44, 0x78,
	0xe4, 0x05, 0xa3, 0x5c, 0x90, 0xac, 0x39, 0xf0, 0x61, 0x4c, 0x02, 0x59, 0x56, 0xcd, 0x3d, 0xaf,
	0xe0, 0xbe, 0x46, 0xfb, 0x0a, 0x44, 0x9f, 0xc3, 0xc5, 0x54, 0x60, 0xff, 0x79, 0xc0, 0xa3, 0x69,
	0xa6, 0x29, 0x8b, 0x55, 0x3f, 0xb7, 0x0a, 0x9c, 0x72, 0x86, 0x9b, 0x70, 0xa1, 0x98, 0x81, 0x71,
	0xfa, 0x8c, 0xf8, 0xc2, 0x8b, 0x02, 0xd9, 0xdf, 0xab, 0xee, 0x66, 0x01, 0x1d, 0x29, 0x70, 0x2f,
	0x40, 0x1d, 0x38, 0x57, 0x54, 0xe1, 0x20, 0xe0, 0x24, 0x4d, 0x2d, 0x90, 0x12, 0x54, 0x80, 0xee,
	0x28, 0x04, 0x3d, 0x06, 0xab, 0x28, 0x28, 0xbe, 0x8b, 0x56, 0x53, 0xfa, 0x76, 0x59, 0x77, 0x3d,
	0x66, 0x51, 0xd6, 0xec, 0xd9, 0xf3, 0x69, 0xdf, 0xe7, 0xcc, 0xdf, 0x57, 0x2c, 0xb7, 0x58, 0x66,
	0x21, 0x9e, 0x59, 0xf7, 0x52, 0x5e, 0x8c, 0xf9, 0xdd, 0x2f, 0x29, 0xeb, 0x14, 0x5c, 0xde, 0xf8,
	0x25, 0x00, 0xad, 0x9b, 0xf0, 0xd8, 0x5a, 0x96, 0x95, 0xaf, 0xaa, 0xc8, 0x23, 0x1e, 0xa3, 0xdb,
	0x60, 0x51, 0x1f, 0x87, 0x24, 0x79, 0xcd, 0x91, 0xd4, 0x65, 0xde, 0x0b, 0x1a, 0x2f, 0x27, 0x7e,
	0x1f, 0xd6, 0x73, 0x65, 0xee, 0x4b, 0x43, 0x66, 0x5f, 0xd3, 0xe1, 0xdc, 0x93, 0x11, 0x6c, 0xe6,
	0xc4, 0x53, 0x7e, 0xac, 0xbd, 0x91, 0x1f, 0x48, 0x6b, 0x8b, 0x5e, 0x3c, 0x83, 0x0b, 0x51, 0xe2,
	0xd3, 0xa3, 0x28, 0x09, 0xbd, 0xe3, 0xd6, 0x94, 0xc3, 0xaa, 0x26, 0x5f, 0x84, 0x9b, 0x67, 0x4c,
	0x96, 0x72, 0x6b, 0x1f, 0x37, 0x6a, 0xa6, 0x75, 0x37, 0xf3, 0x9c, 0xc5, 0x68, 0xb6, 0x16, 0x9d,
	0x88, 0x90, 0xce, 0xaf, 0xb5, 0xfa, 0x36, 0x6b, 0xe5, 0x39, 0x8b, 0xd1, 0xf6, 0x37, 0xd0, 0x38,
	0xb5, 0x76, 0x0d, 0xaa, 0x83, 0xe1, 0xa0, 0x6f, 0x2e, 0xcc, 0xbf, 0x2a, 0x06, 0x42, 0xb0, 0x96,
	0xcd, 0x23, 0x4f, 0xc5, 0x9d, 0xbd, 0x81, 0x59, 0x41, 0x5b, 0x70, 0xae, 0xf7, 0x70, 0xf8, 0xe8,
	0xae, 0x57, 0x7a, 0x82, 0xf4, 0xeb, 0x54, 0x7d, 0x50, 0xad, 0xad, 0x98, 0xb5, 0x36, 0x06, 0x24,
	0xd7, 0xd1, 0x7e, 0xea, 0xcb, 0xf9, 0x05, 0x34, 0x4e, 0x9d, 0x8f, 0xf1, 0x26, 0xe7, 0xe3, 0xd4,
	0x5e, 0x39, 0x4b, 0x3f, 0xca, 0x07, 0xb4, 0x1e, 0x16, 0xc2, 0x9f, 0xfe, 0xf1, 0xfd, 0x9f, 0x7f,
	0x2f, 0x57, 0x4c, 0x03, 0xae, 0x46, 0x54, 0xa5, 0x60, 0x9c, 0xbe, 0x98, 0xbd, 0xde, 0x2d, 0x07,
	0x64, 0x3d, 0xa3, 0x6c, 0x8c, 0x8c, 0x8c, 0xc3, 0x65, 0x39, 0x4f, 0x6e, 0xfc, 0x1f, 0x00, 0x00,
	0xff, 0xff, 0xa8, 0xc0, 0x9e, 0xf5, 0xa4, 0x0a, 0x00, 0x00,
}