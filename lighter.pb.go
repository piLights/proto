// Code generated by protoc-gen-go.
// source: lighter.proto
// DO NOT EDIT!

/*
Package LighterGRPC is a generated protocol buffer package.

It is generated from these files:
	lighter.proto

It has these top-level messages:
	ColorMessage
	StateMessage
	Confirmation
	Request
	ServerConfiguration
	ChangeParameterMessage
	FadeTime
	IPVersion
	Pins
	ScheduledSwitch
	LogEntry
	LogRequest
*/
package LighterGRPC

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type LogLevel int32

const (
	LogLevel_DEBUG LogLevel = 0
	LogLevel_INFO  LogLevel = 1
	LogLevel_WARN  LogLevel = 2
	LogLevel_ERROR LogLevel = 3
	LogLevel_PANIC LogLevel = 4
)

var LogLevel_name = map[int32]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARN",
	3: "ERROR",
	4: "PANIC",
}
var LogLevel_value = map[string]int32{
	"DEBUG": 0,
	"INFO":  1,
	"WARN":  2,
	"ERROR": 3,
	"PANIC": 4,
}

func (x LogLevel) String() string {
	return proto.EnumName(LogLevel_name, int32(x))
}
func (LogLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type IPVersion_Version int32

const (
	IPVersion_DUAL     IPVersion_Version = 0
	IPVersion_IPV4ONLY IPVersion_Version = 1
	IPVersion_IPV6ONLY IPVersion_Version = 2
)

var IPVersion_Version_name = map[int32]string{
	0: "DUAL",
	1: "IPV4ONLY",
	2: "IPV6ONLY",
}
var IPVersion_Version_value = map[string]int32{
	"DUAL":     0,
	"IPV4ONLY": 1,
	"IPV6ONLY": 2,
}

func (x IPVersion_Version) String() string {
	return proto.EnumName(IPVersion_Version_name, int32(x))
}
func (IPVersion_Version) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type LogRequest_LogLevel int32

const (
	LogRequest_TRACE LogRequest_LogLevel = 0
	LogRequest_DEBUG LogRequest_LogLevel = 1
	LogRequest_INFO  LogRequest_LogLevel = 2
	LogRequest_WARN  LogRequest_LogLevel = 3
	LogRequest_ERROR LogRequest_LogLevel = 4
)

var LogRequest_LogLevel_name = map[int32]string{
	0: "TRACE",
	1: "DEBUG",
	2: "INFO",
	3: "WARN",
	4: "ERROR",
}
var LogRequest_LogLevel_value = map[string]int32{
	"TRACE": 0,
	"DEBUG": 1,
	"INFO":  2,
	"WARN":  3,
	"ERROR": 4,
}

func (x LogRequest_LogLevel) String() string {
	return proto.EnumName(LogRequest_LogLevel_name, int32(x))
}
func (LogRequest_LogLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{11, 0} }

type ColorMessage struct {
	Onstate  bool   `protobuf:"varint,1,opt,name=onstate" json:"onstate,omitempty"`
	R        int32  `protobuf:"varint,2,opt,name=r" json:"r,omitempty"`
	G        int32  `protobuf:"varint,3,opt,name=g" json:"g,omitempty"`
	B        int32  `protobuf:"varint,4,opt,name=b" json:"b,omitempty"`
	Opacity  int32  `protobuf:"varint,5,opt,name=opacity" json:"opacity,omitempty"`
	DeviceID string `protobuf:"bytes,6,opt,name=deviceID" json:"deviceID,omitempty"`
	Password string `protobuf:"bytes,7,opt,name=password" json:"password,omitempty"`
}

func (m *ColorMessage) Reset()                    { *m = ColorMessage{} }
func (m *ColorMessage) String() string            { return proto.CompactTextString(m) }
func (*ColorMessage) ProtoMessage()               {}
func (*ColorMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StateMessage struct {
	Onstate  bool   `protobuf:"varint,1,opt,name=onstate" json:"onstate,omitempty"`
	DeviceID string `protobuf:"bytes,2,opt,name=deviceID" json:"deviceID,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *StateMessage) Reset()                    { *m = StateMessage{} }
func (m *StateMessage) String() string            { return proto.CompactTextString(m) }
func (*StateMessage) ProtoMessage()               {}
func (*StateMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Confirmation struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *Confirmation) Reset()                    { *m = Confirmation{} }
func (m *Confirmation) String() string            { return proto.CompactTextString(m) }
func (*Confirmation) ProtoMessage()               {}
func (*Confirmation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Request struct {
	DeviceID string `protobuf:"bytes,1,opt,name=deviceID" json:"deviceID,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ServerConfiguration struct {
	BindTo            string     `protobuf:"bytes,1,opt,name=bindTo" json:"bindTo,omitempty"`
	Pins              *Pins      `protobuf:"bytes,2,opt,name=pins" json:"pins,omitempty"`
	Debug             bool       `protobuf:"varint,3,opt,name=debug" json:"debug,omitempty"`
	UpdateURL         string     `protobuf:"bytes,4,opt,name=updateURL" json:"updateURL,omitempty"`
	ConfigurationFile string     `protobuf:"bytes,5,opt,name=configurationFile" json:"configurationFile,omitempty"`
	Password          string     `protobuf:"bytes,6,opt,name=password" json:"password,omitempty"`
	PiBlaster         string     `protobuf:"bytes,7,opt,name=piBlaster" json:"piBlaster,omitempty"`
	ServerName        string     `protobuf:"bytes,8,opt,name=serverName" json:"serverName,omitempty"`
	IpVersion         *IPVersion `protobuf:"bytes,9,opt,name=ipVersion" json:"ipVersion,omitempty"`
}

func (m *ServerConfiguration) Reset()                    { *m = ServerConfiguration{} }
func (m *ServerConfiguration) String() string            { return proto.CompactTextString(m) }
func (*ServerConfiguration) ProtoMessage()               {}
func (*ServerConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ServerConfiguration) GetPins() *Pins {
	if m != nil {
		return m.Pins
	}
	return nil
}

func (m *ServerConfiguration) GetIpVersion() *IPVersion {
	if m != nil {
		return m.IpVersion
	}
	return nil
}

type ChangeParameterMessage struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	// Types that are valid to be assigned to Parameter:
	//	*ChangeParameterMessage_ServerName
	//	*ChangeParameterMessage_FadeTime
	//	*ChangeParameterMessage_Ipversion
	//	*ChangeParameterMessage_Pins
	Parameter isChangeParameterMessage_Parameter `protobuf_oneof:"parameter"`
}

func (m *ChangeParameterMessage) Reset()                    { *m = ChangeParameterMessage{} }
func (m *ChangeParameterMessage) String() string            { return proto.CompactTextString(m) }
func (*ChangeParameterMessage) ProtoMessage()               {}
func (*ChangeParameterMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isChangeParameterMessage_Parameter interface {
	isChangeParameterMessage_Parameter()
}

type ChangeParameterMessage_ServerName struct {
	ServerName string `protobuf:"bytes,2,opt,name=serverName,oneof"`
}
type ChangeParameterMessage_FadeTime struct {
	FadeTime *FadeTime `protobuf:"bytes,3,opt,name=fadeTime,oneof"`
}
type ChangeParameterMessage_Ipversion struct {
	Ipversion *IPVersion `protobuf:"bytes,4,opt,name=ipversion,oneof"`
}
type ChangeParameterMessage_Pins struct {
	Pins *Pins `protobuf:"bytes,5,opt,name=pins,oneof"`
}

func (*ChangeParameterMessage_ServerName) isChangeParameterMessage_Parameter() {}
func (*ChangeParameterMessage_FadeTime) isChangeParameterMessage_Parameter()   {}
func (*ChangeParameterMessage_Ipversion) isChangeParameterMessage_Parameter()  {}
func (*ChangeParameterMessage_Pins) isChangeParameterMessage_Parameter()       {}

func (m *ChangeParameterMessage) GetParameter() isChangeParameterMessage_Parameter {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func (m *ChangeParameterMessage) GetServerName() string {
	if x, ok := m.GetParameter().(*ChangeParameterMessage_ServerName); ok {
		return x.ServerName
	}
	return ""
}

func (m *ChangeParameterMessage) GetFadeTime() *FadeTime {
	if x, ok := m.GetParameter().(*ChangeParameterMessage_FadeTime); ok {
		return x.FadeTime
	}
	return nil
}

func (m *ChangeParameterMessage) GetIpversion() *IPVersion {
	if x, ok := m.GetParameter().(*ChangeParameterMessage_Ipversion); ok {
		return x.Ipversion
	}
	return nil
}

func (m *ChangeParameterMessage) GetPins() *Pins {
	if x, ok := m.GetParameter().(*ChangeParameterMessage_Pins); ok {
		return x.Pins
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ChangeParameterMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ChangeParameterMessage_OneofMarshaler, _ChangeParameterMessage_OneofUnmarshaler, _ChangeParameterMessage_OneofSizer, []interface{}{
		(*ChangeParameterMessage_ServerName)(nil),
		(*ChangeParameterMessage_FadeTime)(nil),
		(*ChangeParameterMessage_Ipversion)(nil),
		(*ChangeParameterMessage_Pins)(nil),
	}
}

func _ChangeParameterMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ChangeParameterMessage)
	// parameter
	switch x := m.Parameter.(type) {
	case *ChangeParameterMessage_ServerName:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ServerName)
	case *ChangeParameterMessage_FadeTime:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FadeTime); err != nil {
			return err
		}
	case *ChangeParameterMessage_Ipversion:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ipversion); err != nil {
			return err
		}
	case *ChangeParameterMessage_Pins:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pins); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ChangeParameterMessage.Parameter has unexpected type %T", x)
	}
	return nil
}

func _ChangeParameterMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ChangeParameterMessage)
	switch tag {
	case 2: // parameter.serverName
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Parameter = &ChangeParameterMessage_ServerName{x}
		return true, err
	case 3: // parameter.fadeTime
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FadeTime)
		err := b.DecodeMessage(msg)
		m.Parameter = &ChangeParameterMessage_FadeTime{msg}
		return true, err
	case 4: // parameter.ipversion
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IPVersion)
		err := b.DecodeMessage(msg)
		m.Parameter = &ChangeParameterMessage_Ipversion{msg}
		return true, err
	case 5: // parameter.pins
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pins)
		err := b.DecodeMessage(msg)
		m.Parameter = &ChangeParameterMessage_Pins{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ChangeParameterMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ChangeParameterMessage)
	// parameter
	switch x := m.Parameter.(type) {
	case *ChangeParameterMessage_ServerName:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ServerName)))
		n += len(x.ServerName)
	case *ChangeParameterMessage_FadeTime:
		s := proto.Size(x.FadeTime)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ChangeParameterMessage_Ipversion:
		s := proto.Size(x.Ipversion)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ChangeParameterMessage_Pins:
		s := proto.Size(x.Pins)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type FadeTime struct {
	Milliseconds int32 `protobuf:"varint,1,opt,name=milliseconds" json:"milliseconds,omitempty"`
}

func (m *FadeTime) Reset()                    { *m = FadeTime{} }
func (m *FadeTime) String() string            { return proto.CompactTextString(m) }
func (*FadeTime) ProtoMessage()               {}
func (*FadeTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type IPVersion struct {
	Version IPVersion_Version `protobuf:"varint,1,opt,name=version,enum=LighterGRPC.IPVersion_Version" json:"version,omitempty"`
}

func (m *IPVersion) Reset()                    { *m = IPVersion{} }
func (m *IPVersion) String() string            { return proto.CompactTextString(m) }
func (*IPVersion) ProtoMessage()               {}
func (*IPVersion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Pins struct {
	RedPin   int32 `protobuf:"varint,1,opt,name=redPin" json:"redPin,omitempty"`
	GreenPin int32 `protobuf:"varint,2,opt,name=greenPin" json:"greenPin,omitempty"`
	BluePin  int32 `protobuf:"varint,3,opt,name=bluePin" json:"bluePin,omitempty"`
	WhitePin int32 `protobuf:"varint,4,opt,name=whitePin" json:"whitePin,omitempty"`
}

func (m *Pins) Reset()                    { *m = Pins{} }
func (m *Pins) String() string            { return proto.CompactTextString(m) }
func (*Pins) ProtoMessage()               {}
func (*Pins) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ScheduledSwitch struct {
	Time     int64  `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Onstate  bool   `protobuf:"varint,2,opt,name=onstate" json:"onstate,omitempty"`
	DeviceID string `protobuf:"bytes,3,opt,name=deviceID" json:"deviceID,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
}

func (m *ScheduledSwitch) Reset()                    { *m = ScheduledSwitch{} }
func (m *ScheduledSwitch) String() string            { return proto.CompactTextString(m) }
func (*ScheduledSwitch) ProtoMessage()               {}
func (*ScheduledSwitch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type LogEntry struct {
	Time    int64  `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type LogRequest struct {
	LogLevel LogRequest_LogLevel `protobuf:"varint,1,opt,name=logLevel,enum=LighterGRPC.LogRequest_LogLevel" json:"logLevel,omitempty"`
	Amount   int64               `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Password string              `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *LogRequest) Reset()                    { *m = LogRequest{} }
func (m *LogRequest) String() string            { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()               {}
func (*LogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*ColorMessage)(nil), "LighterGRPC.ColorMessage")
	proto.RegisterType((*StateMessage)(nil), "LighterGRPC.StateMessage")
	proto.RegisterType((*Confirmation)(nil), "LighterGRPC.Confirmation")
	proto.RegisterType((*Request)(nil), "LighterGRPC.Request")
	proto.RegisterType((*ServerConfiguration)(nil), "LighterGRPC.ServerConfiguration")
	proto.RegisterType((*ChangeParameterMessage)(nil), "LighterGRPC.ChangeParameterMessage")
	proto.RegisterType((*FadeTime)(nil), "LighterGRPC.FadeTime")
	proto.RegisterType((*IPVersion)(nil), "LighterGRPC.IPVersion")
	proto.RegisterType((*Pins)(nil), "LighterGRPC.Pins")
	proto.RegisterType((*ScheduledSwitch)(nil), "LighterGRPC.ScheduledSwitch")
	proto.RegisterType((*LogEntry)(nil), "LighterGRPC.LogEntry")
	proto.RegisterType((*LogRequest)(nil), "LighterGRPC.LogRequest")
	proto.RegisterEnum("LighterGRPC.LogLevel", LogLevel_name, LogLevel_value)
	proto.RegisterEnum("LighterGRPC.IPVersion_Version", IPVersion_Version_name, IPVersion_Version_value)
	proto.RegisterEnum("LighterGRPC.LogRequest_LogLevel", LogRequest_LogLevel_name, LogRequest_LogLevel_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Lighter service

type LighterClient interface {
	SetColor(ctx context.Context, in *ColorMessage, opts ...grpc.CallOption) (*Confirmation, error)
	CheckConnection(ctx context.Context, in *Request, opts ...grpc.CallOption) (Lighter_CheckConnectionClient, error)
	SwitchState(ctx context.Context, in *StateMessage, opts ...grpc.CallOption) (*Confirmation, error)
	ChangeServerParameter(ctx context.Context, in *ChangeParameterMessage, opts ...grpc.CallOption) (*Confirmation, error)
	ScheduleSwitchState(ctx context.Context, in *ScheduledSwitch, opts ...grpc.CallOption) (*Confirmation, error)
	LoadServerLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (Lighter_LoadServerLogClient, error)
	// Get and set the Configuration of the backend
	LoadServerConfiguration(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ServerConfiguration, error)
	SetServerConfiguration(ctx context.Context, in *ServerConfiguration, opts ...grpc.CallOption) (*Confirmation, error)
}

type lighterClient struct {
	cc *grpc.ClientConn
}

func NewLighterClient(cc *grpc.ClientConn) LighterClient {
	return &lighterClient{cc}
}

func (c *lighterClient) SetColor(ctx context.Context, in *ColorMessage, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/LighterGRPC.Lighter/setColor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lighterClient) CheckConnection(ctx context.Context, in *Request, opts ...grpc.CallOption) (Lighter_CheckConnectionClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Lighter_serviceDesc.Streams[0], c.cc, "/LighterGRPC.Lighter/checkConnection", opts...)
	if err != nil {
		return nil, err
	}
	x := &lighterCheckConnectionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Lighter_CheckConnectionClient interface {
	Recv() (*ColorMessage, error)
	grpc.ClientStream
}

type lighterCheckConnectionClient struct {
	grpc.ClientStream
}

func (x *lighterCheckConnectionClient) Recv() (*ColorMessage, error) {
	m := new(ColorMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lighterClient) SwitchState(ctx context.Context, in *StateMessage, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/LighterGRPC.Lighter/switchState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lighterClient) ChangeServerParameter(ctx context.Context, in *ChangeParameterMessage, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/LighterGRPC.Lighter/changeServerParameter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lighterClient) ScheduleSwitchState(ctx context.Context, in *ScheduledSwitch, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/LighterGRPC.Lighter/scheduleSwitchState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lighterClient) LoadServerLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (Lighter_LoadServerLogClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Lighter_serviceDesc.Streams[1], c.cc, "/LighterGRPC.Lighter/loadServerLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &lighterLoadServerLogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Lighter_LoadServerLogClient interface {
	Recv() (*LogEntry, error)
	grpc.ClientStream
}

type lighterLoadServerLogClient struct {
	grpc.ClientStream
}

func (x *lighterLoadServerLogClient) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lighterClient) LoadServerConfiguration(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ServerConfiguration, error) {
	out := new(ServerConfiguration)
	err := grpc.Invoke(ctx, "/LighterGRPC.Lighter/loadServerConfiguration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lighterClient) SetServerConfiguration(ctx context.Context, in *ServerConfiguration, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/LighterGRPC.Lighter/setServerConfiguration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lighter service

type LighterServer interface {
	SetColor(context.Context, *ColorMessage) (*Confirmation, error)
	CheckConnection(*Request, Lighter_CheckConnectionServer) error
	SwitchState(context.Context, *StateMessage) (*Confirmation, error)
	ChangeServerParameter(context.Context, *ChangeParameterMessage) (*Confirmation, error)
	ScheduleSwitchState(context.Context, *ScheduledSwitch) (*Confirmation, error)
	LoadServerLog(*LogRequest, Lighter_LoadServerLogServer) error
	// Get and set the Configuration of the backend
	LoadServerConfiguration(context.Context, *Request) (*ServerConfiguration, error)
	SetServerConfiguration(context.Context, *ServerConfiguration) (*Confirmation, error)
}

func RegisterLighterServer(s *grpc.Server, srv LighterServer) {
	s.RegisterService(&_Lighter_serviceDesc, srv)
}

func _Lighter_SetColor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ColorMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LighterServer).SetColor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LighterGRPC.Lighter/SetColor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LighterServer).SetColor(ctx, req.(*ColorMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lighter_CheckConnection_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LighterServer).CheckConnection(m, &lighterCheckConnectionServer{stream})
}

type Lighter_CheckConnectionServer interface {
	Send(*ColorMessage) error
	grpc.ServerStream
}

type lighterCheckConnectionServer struct {
	grpc.ServerStream
}

func (x *lighterCheckConnectionServer) Send(m *ColorMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Lighter_SwitchState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LighterServer).SwitchState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LighterGRPC.Lighter/SwitchState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LighterServer).SwitchState(ctx, req.(*StateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lighter_ChangeServerParameter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeParameterMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LighterServer).ChangeServerParameter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LighterGRPC.Lighter/ChangeServerParameter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LighterServer).ChangeServerParameter(ctx, req.(*ChangeParameterMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lighter_ScheduleSwitchState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduledSwitch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LighterServer).ScheduleSwitchState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LighterGRPC.Lighter/ScheduleSwitchState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LighterServer).ScheduleSwitchState(ctx, req.(*ScheduledSwitch))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lighter_LoadServerLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LighterServer).LoadServerLog(m, &lighterLoadServerLogServer{stream})
}

type Lighter_LoadServerLogServer interface {
	Send(*LogEntry) error
	grpc.ServerStream
}

type lighterLoadServerLogServer struct {
	grpc.ServerStream
}

func (x *lighterLoadServerLogServer) Send(m *LogEntry) error {
	return x.ServerStream.SendMsg(m)
}

func _Lighter_LoadServerConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LighterServer).LoadServerConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LighterGRPC.Lighter/LoadServerConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LighterServer).LoadServerConfiguration(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lighter_SetServerConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerConfiguration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LighterServer).SetServerConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LighterGRPC.Lighter/SetServerConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LighterServer).SetServerConfiguration(ctx, req.(*ServerConfiguration))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lighter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LighterGRPC.Lighter",
	HandlerType: (*LighterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "setColor",
			Handler:    _Lighter_SetColor_Handler,
		},
		{
			MethodName: "switchState",
			Handler:    _Lighter_SwitchState_Handler,
		},
		{
			MethodName: "changeServerParameter",
			Handler:    _Lighter_ChangeServerParameter_Handler,
		},
		{
			MethodName: "scheduleSwitchState",
			Handler:    _Lighter_ScheduleSwitchState_Handler,
		},
		{
			MethodName: "loadServerConfiguration",
			Handler:    _Lighter_LoadServerConfiguration_Handler,
		},
		{
			MethodName: "setServerConfiguration",
			Handler:    _Lighter_SetServerConfiguration_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "checkConnection",
			Handler:       _Lighter_CheckConnection_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "loadServerLog",
			Handler:       _Lighter_LoadServerLog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("lighter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 929 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x15, 0x25, 0xca, 0xa2, 0xae, 0xdd, 0x46, 0x19, 0xc7, 0x36, 0x6b, 0x04, 0x81, 0xc1, 0xa2,
	0x68, 0x50, 0x14, 0x6a, 0xe1, 0x04, 0x41, 0x16, 0x05, 0x0a, 0x49, 0x96, 0x53, 0x03, 0xaa, 0x2c,
	0x8c, 0xec, 0x06, 0xd9, 0x95, 0x22, 0x27, 0x14, 0x11, 0x8a, 0x54, 0xf9, 0x90, 0x91, 0xff, 0xe9,
	0x7f, 0x74, 0xd7, 0x7f, 0x2a, 0xd0, 0x45, 0xe7, 0x45, 0x8a, 0xc3, 0x4a, 0xf4, 0x8e, 0x67, 0xee,
	0xeb, 0xdc, 0x39, 0x97, 0x97, 0x84, 0x2f, 0x02, 0xdf, 0x5b, 0xa6, 0x24, 0xee, 0xaf, 0xe3, 0x28,
	0x8d, 0xd0, 0xe1, 0x44, 0xc0, 0x77, 0x78, 0x36, 0xb2, 0xfe, 0xd4, 0xe0, 0x68, 0x14, 0x05, 0x51,
	0xfc, 0x2b, 0x49, 0x12, 0xdb, 0x23, 0xc8, 0x84, 0x4e, 0x14, 0x26, 0xa9, 0x9d, 0x12, 0x53, 0xbb,
	0xd0, 0x5e, 0x1a, 0x38, 0x87, 0xe8, 0x08, 0xb4, 0xd8, 0x6c, 0xd2, 0xb3, 0x36, 0xd6, 0x62, 0x86,
	0x3c, 0xb3, 0x25, 0x90, 0xc7, 0xd0, 0xc2, 0xd4, 0x05, 0x5a, 0xf0, 0x1c, 0x6b, 0xdb, 0xf1, 0xd3,
	0xcf, 0x66, 0x9b, 0x9f, 0xe5, 0x10, 0x9d, 0x83, 0xe1, 0x92, 0x8d, 0xef, 0x90, 0x9b, 0x2b, 0xf3,
	0x80, 0x9a, 0xba, 0xb8, 0xc0, 0xcc, 0xb6, 0xb6, 0x93, 0xe4, 0x21, 0x8a, 0x5d, 0xb3, 0x23, 0x6c,
	0x39, 0xb6, 0x7e, 0x87, 0xa3, 0x39, 0x23, 0xf1, 0x38, 0xcb, 0x72, 0x85, 0x66, 0x4d, 0x85, 0x56,
	0xa5, 0xc2, 0x4b, 0x76, 0x0f, 0xe1, 0x47, 0x3f, 0x5e, 0xd9, 0xa9, 0x1f, 0x85, 0xac, 0x42, 0x92,
	0x39, 0x0e, 0xad, 0x97, 0x57, 0x90, 0xd0, 0x1a, 0x40, 0x07, 0x93, 0x3f, 0x32, 0x92, 0xa4, 0x4a,
	0x31, 0xad, 0xa6, 0x58, 0xb3, 0x52, 0xec, 0xef, 0x26, 0x1c, 0xcf, 0x49, 0xbc, 0x21, 0x31, 0xaf,
	0xe9, 0x65, 0xb1, 0x28, 0x7a, 0x0a, 0x07, 0x0b, 0x3f, 0x74, 0xef, 0x22, 0x99, 0x4d, 0x22, 0xf4,
	0x0d, 0xe8, 0x6b, 0x3f, 0x4c, 0x78, 0x9e, 0xc3, 0xcb, 0xa7, 0xfd, 0x92, 0x82, 0xfd, 0x19, 0x35,
	0x60, 0x6e, 0x46, 0xcf, 0xa0, 0xed, 0x92, 0x45, 0x26, 0x74, 0x31, 0xb0, 0x00, 0xe8, 0x39, 0x74,
	0xb3, 0xb5, 0x4b, 0xef, 0xe6, 0x1e, 0x4f, 0xb8, 0x46, 0x5d, 0xbc, 0x3d, 0x40, 0xdf, 0xc3, 0x53,
	0xa7, 0xcc, 0xe1, 0xda, 0x0f, 0x08, 0x57, 0xad, 0x8b, 0xff, 0x6f, 0x50, 0x9a, 0x3a, 0x50, 0x9b,
	0x62, 0x75, 0xd6, 0xfe, 0x30, 0xb0, 0x13, 0xca, 0x4c, 0x0a, 0xb8, 0x3d, 0x40, 0x2f, 0x00, 0x12,
	0xde, 0xf1, 0xd4, 0x5e, 0x11, 0xd3, 0xe0, 0xe6, 0xd2, 0x09, 0x7a, 0x0d, 0x5d, 0x7f, 0xfd, 0x1b,
	0x89, 0x13, 0x5a, 0xca, 0xec, 0xf2, 0x3e, 0x4f, 0x95, 0x3e, 0x6f, 0x66, 0xd2, 0x8a, 0xb7, 0x8e,
	0xd6, 0x3f, 0x1a, 0x9c, 0x8e, 0x96, 0x76, 0xe8, 0x91, 0x99, 0x1d, 0xd3, 0x34, 0xd4, 0x39, 0x1f,
	0x91, 0x32, 0x55, 0xad, 0x42, 0xf5, 0x42, 0x21, 0xc3, 0xd5, 0xf9, 0xa5, 0xa1, 0xd0, 0x79, 0x05,
	0xc6, 0x47, 0xdb, 0x25, 0x77, 0x3e, 0xb5, 0xb7, 0x38, 0x9b, 0x13, 0x85, 0xcd, 0xb5, 0x34, 0xd2,
	0xb0, 0xc2, 0x11, 0xbd, 0x61, 0x3d, 0x6c, 0x64, 0x0f, 0x7a, 0x5d, 0x0f, 0x34, 0x6c, 0xeb, 0x8a,
	0xbe, 0x95, 0xf2, 0xb6, 0xf7, 0xc8, 0x4b, 0xbd, 0xb9, 0xc3, 0xf0, 0x90, 0x5e, 0x71, 0xde, 0xa7,
	0xd5, 0x07, 0x23, 0x67, 0x81, 0x2c, 0x38, 0x5a, 0xf9, 0x41, 0xe0, 0x27, 0x84, 0x6a, 0xe6, 0x8a,
	0x91, 0x6d, 0x63, 0xe5, 0xcc, 0xda, 0x40, 0xb7, 0xa8, 0x8f, 0xde, 0x42, 0x27, 0x27, 0xca, 0x7c,
	0xbf, 0xbc, 0x7c, 0xb1, 0x9b, 0x68, 0x3f, 0xbf, 0xf4, 0xdc, 0xdd, 0xfa, 0x01, 0x3a, 0x79, 0x12,
	0x03, 0xf4, 0xab, 0xfb, 0xc1, 0xa4, 0xd7, 0xa0, 0xef, 0xbf, 0x41, 0x43, 0x5e, 0xdf, 0x4e, 0x27,
	0x1f, 0x7a, 0x9a, 0x44, 0x6f, 0x38, 0x6a, 0x5a, 0x6b, 0xd0, 0x59, 0x13, 0x6c, 0xb8, 0x63, 0xe2,
	0xd2, 0x47, 0xc9, 0x4e, 0x22, 0x26, 0x94, 0x17, 0x13, 0x12, 0x32, 0x8b, 0x58, 0x2f, 0x05, 0x66,
	0x6f, 0xe1, 0x22, 0xc8, 0x08, 0x33, 0x89, 0x5d, 0x93, 0x43, 0x16, 0xf5, 0xb0, 0xf4, 0x53, 0x6e,
	0x12, 0x8b, 0xa7, 0xc0, 0xd6, 0x03, 0x3c, 0x99, 0x3b, 0x4b, 0xe2, 0x66, 0x01, 0x71, 0xe7, 0x0f,
	0x7e, 0xea, 0x2c, 0x11, 0x02, 0x3d, 0x65, 0x5a, 0xb2, 0xd2, 0x2d, 0xcc, 0x9f, 0xcb, 0x4b, 0xa4,
	0xb9, 0x7f, 0x89, 0xb4, 0x6a, 0xde, 0x6b, 0xbd, 0xf2, 0x5e, 0xbf, 0x05, 0x63, 0x12, 0x79, 0xe3,
	0x30, 0x8d, 0x3f, 0xef, 0xab, 0xb8, 0x12, 0xe3, 0x29, 0x57, 0x42, 0x0e, 0xad, 0xbf, 0x34, 0x00,
	0x1a, 0x9a, 0x2f, 0x96, 0x9f, 0xc0, 0x08, 0x22, 0x6f, 0x42, 0x36, 0x24, 0x90, 0xfa, 0x5c, 0x28,
	0xfa, 0x6c, 0x5d, 0xd9, 0x23, 0xf7, 0xc3, 0x45, 0x04, 0xbb, 0x69, 0x7b, 0x15, 0x65, 0x61, 0xca,
	0xab, 0xb4, 0xb0, 0x44, 0xb5, 0xfb, 0xef, 0x67, 0x4e, 0x5d, 0xc4, 0x77, 0xa1, 0x7d, 0x87, 0x07,
	0xa3, 0x31, 0x15, 0x96, 0x3e, 0x5e, 0x8d, 0x87, 0xf7, 0xef, 0xa8, 0xaa, 0x54, 0xed, 0x9b, 0xe9,
	0xf5, 0x6d, 0xaf, 0xc9, 0x9e, 0xde, 0x0f, 0xf0, 0xb4, 0xd7, 0x62, 0xe6, 0x31, 0xc6, 0xb7, 0xb8,
	0xa7, 0x7f, 0x57, 0x49, 0x20, 0xa2, 0x1a, 0x45, 0x94, 0x56, 0x44, 0x35, 0xb7, 0x51, 0x3c, 0xc1,
	0x6c, 0x30, 0xbd, 0x19, 0xf5, 0xf4, 0xcb, 0x7f, 0x75, 0xe8, 0xc8, 0x1e, 0xd1, 0x10, 0x8c, 0x84,
	0xa4, 0xfc, 0xc3, 0x84, 0xbe, 0x52, 0x3a, 0x2f, 0x7f, 0xac, 0xce, 0xab, 0xa6, 0xed, 0xfe, 0xb6,
	0x1a, 0xe8, 0x1a, 0x9e, 0xd0, 0x21, 0x70, 0x3e, 0xd1, 0xe3, 0x90, 0x38, 0x7c, 0xbf, 0x3e, 0x53,
	0xfc, 0xe5, 0x0d, 0x9e, 0xef, 0x2f, 0x60, 0x35, 0x7e, 0xd4, 0xd0, 0x18, 0x0e, 0x13, 0x3e, 0x44,
	0xfc, 0x0b, 0x54, 0xa1, 0x53, 0xfe, 0x2a, 0xd5, 0xd3, 0xf9, 0x00, 0x27, 0x0e, 0xdf, 0x54, 0x62,
	0xf1, 0x17, 0xfb, 0x0a, 0x7d, 0xad, 0x46, 0xed, 0xdc, 0x66, 0xf5, 0xa9, 0x67, 0x70, 0x9c, 0xc8,
	0x79, 0x9f, 0x97, 0x98, 0x3e, 0x57, 0x99, 0xaa, 0x6f, 0x44, 0x7d, 0xc6, 0x11, 0xfd, 0x69, 0x88,
	0x6c, 0x57, 0x50, 0xa5, 0xb2, 0xa2, 0xb3, 0x3d, 0xe3, 0x77, 0x7e, 0x52, 0x35, 0xf0, 0xe9, 0xe7,
	0x17, 0x37, 0x87, 0xb3, 0x6d, 0x12, 0xf5, 0x43, 0xb7, 0x5b, 0x08, 0x75, 0xc6, 0x77, 0xc4, 0x51,
	0x66, 0xef, 0xe1, 0x94, 0x4e, 0xc6, 0xae, 0x9c, 0x8f, 0x46, 0xd7, 0xb6, 0x3c, 0x3c, 0x83, 0x63,
	0x97, 0xf4, 0xbd, 0x38, 0x0b, 0x7d, 0xe7, 0x13, 0xe9, 0xcb, 0x7f, 0xa6, 0x99, 0xb6, 0x38, 0xe0,
	0xbf, 0x4d, 0xaf, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x43, 0x14, 0xd6, 0x24, 0x47, 0x09, 0x00,
	0x00,
}
