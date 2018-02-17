// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Nsmf_PDUSession.proto

/*
	Package ngc is a generated protocol buffer package.

	It is generated from these files:
		Nsmf_PDUSession.proto

	It has these top-level messages:
		CreateSMContextRequest
		CreateSMContextResponse
*/
package ngc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Define message for Create SM Context
// Input, Required: SUPI or PEI, DNN, S-NSSAI(s), PDU Session ID, AMF ID (GUAMI), N1 SM container.
// Input, Optional: PEI, UE location information, UE Time Zone, AN type, H-SMF identifier/address, old PDU Session ID (if the AMF also received an old PDU Session ID from the UE as specified in clause 4.3.5.2), Subscription For PDU Session Status Notification, indication that the SUPI has not been authenticated..
type CreateSMContextRequest struct {
	SUPI               string `protobuf:"bytes,2,opt,name=SUPI,proto3" json:"SUPI,omitempty"`
	DNN                string `protobuf:"bytes,3,opt,name=DNN,proto3" json:"DNN,omitempty"`
	S_NSSAI            uint32 `protobuf:"varint,4,opt,name=S_NSSAI,json=SNSSAI,proto3" json:"S_NSSAI,omitempty"`
	PDU_Session_ID     uint64 `protobuf:"varint,5,opt,name=PDU_Session_ID,json=PDUSessionID,proto3" json:"PDU_Session_ID,omitempty"`
	AMF_ID             uint32 `protobuf:"varint,6,opt,name=AMF_ID,json=AMFID,proto3" json:"AMF_ID,omitempty"`
	N1_SM_Container    []byte `protobuf:"bytes,7,opt,name=N1_SM_Container,json=N1SMContainer,proto3" json:"N1_SM_Container,omitempty"`
	PEI                string `protobuf:"bytes,8,opt,name=PEI,proto3" json:"PEI,omitempty"`
	UE_Location        string `protobuf:"bytes,9,opt,name=UE_Location,json=UELocation,proto3" json:"UE_Location,omitempty"`
	UE_Time_Zone       string `protobuf:"bytes,10,opt,name=UE_Time_Zone,json=UETimeZone,proto3" json:"UE_Time_Zone,omitempty"`
	AN_Type            string `protobuf:"bytes,11,opt,name=AN_Type,json=ANType,proto3" json:"AN_Type,omitempty"`
	H_SMF_ID           uint32 `protobuf:"varint,12,opt,name=H_SMF_ID,json=HSMFID,proto3" json:"H_SMF_ID,omitempty"`
	Old_PDU_Session_ID uint64 `protobuf:"varint,13,opt,name=Old_PDU_Session_ID,json=OldPDUSessionID,proto3" json:"Old_PDU_Session_ID,omitempty"`
}

func (m *CreateSMContextRequest) Reset()      { *m = CreateSMContextRequest{} }
func (*CreateSMContextRequest) ProtoMessage() {}
func (*CreateSMContextRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorNsmf_PDUSession, []int{0}
}

func (m *CreateSMContextRequest) GetSUPI() string {
	if m != nil {
		return m.SUPI
	}
	return ""
}

func (m *CreateSMContextRequest) GetDNN() string {
	if m != nil {
		return m.DNN
	}
	return ""
}

func (m *CreateSMContextRequest) GetS_NSSAI() uint32 {
	if m != nil {
		return m.S_NSSAI
	}
	return 0
}

func (m *CreateSMContextRequest) GetPDU_Session_ID() uint64 {
	if m != nil {
		return m.PDU_Session_ID
	}
	return 0
}

func (m *CreateSMContextRequest) GetAMF_ID() uint32 {
	if m != nil {
		return m.AMF_ID
	}
	return 0
}

func (m *CreateSMContextRequest) GetN1_SM_Container() []byte {
	if m != nil {
		return m.N1_SM_Container
	}
	return nil
}

func (m *CreateSMContextRequest) GetPEI() string {
	if m != nil {
		return m.PEI
	}
	return ""
}

func (m *CreateSMContextRequest) GetUE_Location() string {
	if m != nil {
		return m.UE_Location
	}
	return ""
}

func (m *CreateSMContextRequest) GetUE_Time_Zone() string {
	if m != nil {
		return m.UE_Time_Zone
	}
	return ""
}

func (m *CreateSMContextRequest) GetAN_Type() string {
	if m != nil {
		return m.AN_Type
	}
	return ""
}

func (m *CreateSMContextRequest) GetH_SMF_ID() uint32 {
	if m != nil {
		return m.H_SMF_ID
	}
	return 0
}

func (m *CreateSMContextRequest) GetOld_PDU_Session_ID() uint64 {
	if m != nil {
		return m.Old_PDU_Session_ID
	}
	return 0
}

// Output, Required: Result Indication.
// Output, Optional: Cause, PDU Session ID, N2 SM information, N1 SM container.
type CreateSMContextResponse struct {
	Result_Indication string `protobuf:"bytes,1,opt,name=Result_Indication,json=ResultIndication,proto3" json:"Result_Indication,omitempty"`
	Cause             uint32 `protobuf:"varint,2,opt,name=Cause,proto3" json:"Cause,omitempty"`
	PDU_Session_ID    uint64 `protobuf:"varint,3,opt,name=PDU_Session_ID,json=PDUSessionID,proto3" json:"PDU_Session_ID,omitempty"`
	N2_SM_Info        []byte `protobuf:"bytes,4,opt,name=N2_SM_Info,json=N2SMInfo,proto3" json:"N2_SM_Info,omitempty"`
	N1_SM_Container   []byte `protobuf:"bytes,5,opt,name=N1_SM_Container,json=N1SMContainer,proto3" json:"N1_SM_Container,omitempty"`
}

func (m *CreateSMContextResponse) Reset()      { *m = CreateSMContextResponse{} }
func (*CreateSMContextResponse) ProtoMessage() {}
func (*CreateSMContextResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorNsmf_PDUSession, []int{1}
}

func (m *CreateSMContextResponse) GetResult_Indication() string {
	if m != nil {
		return m.Result_Indication
	}
	return ""
}

func (m *CreateSMContextResponse) GetCause() uint32 {
	if m != nil {
		return m.Cause
	}
	return 0
}

func (m *CreateSMContextResponse) GetPDU_Session_ID() uint64 {
	if m != nil {
		return m.PDU_Session_ID
	}
	return 0
}

func (m *CreateSMContextResponse) GetN2_SM_Info() []byte {
	if m != nil {
		return m.N2_SM_Info
	}
	return nil
}

func (m *CreateSMContextResponse) GetN1_SM_Container() []byte {
	if m != nil {
		return m.N1_SM_Container
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateSMContextRequest)(nil), "ngc.CreateSMContextRequest")
	proto.RegisterType((*CreateSMContextResponse)(nil), "ngc.CreateSMContextResponse")
}
func (this *CreateSMContextRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateSMContextRequest)
	if !ok {
		that2, ok := that.(CreateSMContextRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SUPI != that1.SUPI {
		return false
	}
	if this.DNN != that1.DNN {
		return false
	}
	if this.S_NSSAI != that1.S_NSSAI {
		return false
	}
	if this.PDU_Session_ID != that1.PDU_Session_ID {
		return false
	}
	if this.AMF_ID != that1.AMF_ID {
		return false
	}
	if !bytes.Equal(this.N1_SM_Container, that1.N1_SM_Container) {
		return false
	}
	if this.PEI != that1.PEI {
		return false
	}
	if this.UE_Location != that1.UE_Location {
		return false
	}
	if this.UE_Time_Zone != that1.UE_Time_Zone {
		return false
	}
	if this.AN_Type != that1.AN_Type {
		return false
	}
	if this.H_SMF_ID != that1.H_SMF_ID {
		return false
	}
	if this.Old_PDU_Session_ID != that1.Old_PDU_Session_ID {
		return false
	}
	return true
}
func (this *CreateSMContextResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateSMContextResponse)
	if !ok {
		that2, ok := that.(CreateSMContextResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Result_Indication != that1.Result_Indication {
		return false
	}
	if this.Cause != that1.Cause {
		return false
	}
	if this.PDU_Session_ID != that1.PDU_Session_ID {
		return false
	}
	if !bytes.Equal(this.N2_SM_Info, that1.N2_SM_Info) {
		return false
	}
	if !bytes.Equal(this.N1_SM_Container, that1.N1_SM_Container) {
		return false
	}
	return true
}
func (this *CreateSMContextRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 16)
	s = append(s, "&ngc.CreateSMContextRequest{")
	s = append(s, "SUPI: "+fmt.Sprintf("%#v", this.SUPI)+",\n")
	s = append(s, "DNN: "+fmt.Sprintf("%#v", this.DNN)+",\n")
	s = append(s, "S_NSSAI: "+fmt.Sprintf("%#v", this.S_NSSAI)+",\n")
	s = append(s, "PDU_Session_ID: "+fmt.Sprintf("%#v", this.PDU_Session_ID)+",\n")
	s = append(s, "AMF_ID: "+fmt.Sprintf("%#v", this.AMF_ID)+",\n")
	s = append(s, "N1_SM_Container: "+fmt.Sprintf("%#v", this.N1_SM_Container)+",\n")
	s = append(s, "PEI: "+fmt.Sprintf("%#v", this.PEI)+",\n")
	s = append(s, "UE_Location: "+fmt.Sprintf("%#v", this.UE_Location)+",\n")
	s = append(s, "UE_Time_Zone: "+fmt.Sprintf("%#v", this.UE_Time_Zone)+",\n")
	s = append(s, "AN_Type: "+fmt.Sprintf("%#v", this.AN_Type)+",\n")
	s = append(s, "H_SMF_ID: "+fmt.Sprintf("%#v", this.H_SMF_ID)+",\n")
	s = append(s, "Old_PDU_Session_ID: "+fmt.Sprintf("%#v", this.Old_PDU_Session_ID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateSMContextResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&ngc.CreateSMContextResponse{")
	s = append(s, "Result_Indication: "+fmt.Sprintf("%#v", this.Result_Indication)+",\n")
	s = append(s, "Cause: "+fmt.Sprintf("%#v", this.Cause)+",\n")
	s = append(s, "PDU_Session_ID: "+fmt.Sprintf("%#v", this.PDU_Session_ID)+",\n")
	s = append(s, "N2_SM_Info: "+fmt.Sprintf("%#v", this.N2_SM_Info)+",\n")
	s = append(s, "N1_SM_Container: "+fmt.Sprintf("%#v", this.N1_SM_Container)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringNsmf_PDUSession(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Nsmf_PDUSession service

type Nsmf_PDUSessionClient interface {
	// Service operation name: Nsmf_PDUSession_CreateSMContext.
	// Description: It creates an AMF-SMF association to support a PDU Session.
	// Input, Required: SUPI or PEI, DNN, S-NSSAI(s), PDU Session ID, AMF ID (GUAMI), N1 SM container.
	// Input, Optional: PEI, UE location information, UE Time Zone, AN type, H-SMF identifier/address, old PDU Session ID (if the AMF also received an old PDU Session ID from the UE as specified in clause 4.3.5.2), Subscription For PDU Session Status Notification, indication that the SUPI has not been authenticated..
	// Output, Required: Result Indication.
	// Output, Optional: Cause, PDU Session ID, N2 SM information, N1 SM container.
	CreateSMContext(ctx context.Context, in *CreateSMContextRequest, opts ...grpc.CallOption) (*CreateSMContextResponse, error)
}

type nsmf_PDUSessionClient struct {
	cc *grpc.ClientConn
}

func NewNsmf_PDUSessionClient(cc *grpc.ClientConn) Nsmf_PDUSessionClient {
	return &nsmf_PDUSessionClient{cc}
}

func (c *nsmf_PDUSessionClient) CreateSMContext(ctx context.Context, in *CreateSMContextRequest, opts ...grpc.CallOption) (*CreateSMContextResponse, error) {
	out := new(CreateSMContextResponse)
	err := grpc.Invoke(ctx, "/ngc.Nsmf_PDUSession/CreateSMContext", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Nsmf_PDUSession service

type Nsmf_PDUSessionServer interface {
	// Service operation name: Nsmf_PDUSession_CreateSMContext.
	// Description: It creates an AMF-SMF association to support a PDU Session.
	// Input, Required: SUPI or PEI, DNN, S-NSSAI(s), PDU Session ID, AMF ID (GUAMI), N1 SM container.
	// Input, Optional: PEI, UE location information, UE Time Zone, AN type, H-SMF identifier/address, old PDU Session ID (if the AMF also received an old PDU Session ID from the UE as specified in clause 4.3.5.2), Subscription For PDU Session Status Notification, indication that the SUPI has not been authenticated..
	// Output, Required: Result Indication.
	// Output, Optional: Cause, PDU Session ID, N2 SM information, N1 SM container.
	CreateSMContext(context.Context, *CreateSMContextRequest) (*CreateSMContextResponse, error)
}

func RegisterNsmf_PDUSessionServer(s *grpc.Server, srv Nsmf_PDUSessionServer) {
	s.RegisterService(&_Nsmf_PDUSession_serviceDesc, srv)
}

func _Nsmf_PDUSession_CreateSMContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSMContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Nsmf_PDUSessionServer).CreateSMContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ngc.Nsmf_PDUSession/CreateSMContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Nsmf_PDUSessionServer).CreateSMContext(ctx, req.(*CreateSMContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nsmf_PDUSession_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ngc.Nsmf_PDUSession",
	HandlerType: (*Nsmf_PDUSessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSMContext",
			Handler:    _Nsmf_PDUSession_CreateSMContext_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Nsmf_PDUSession.proto",
}

func (m *CreateSMContextRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateSMContextRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SUPI) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(len(m.SUPI)))
		i += copy(dAtA[i:], m.SUPI)
	}
	if len(m.DNN) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(len(m.DNN)))
		i += copy(dAtA[i:], m.DNN)
	}
	if m.S_NSSAI != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(m.S_NSSAI))
	}
	if m.PDU_Session_ID != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(m.PDU_Session_ID))
	}
	if m.AMF_ID != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(m.AMF_ID))
	}
	if len(m.N1_SM_Container) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(len(m.N1_SM_Container)))
		i += copy(dAtA[i:], m.N1_SM_Container)
	}
	if len(m.PEI) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(len(m.PEI)))
		i += copy(dAtA[i:], m.PEI)
	}
	if len(m.UE_Location) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(len(m.UE_Location)))
		i += copy(dAtA[i:], m.UE_Location)
	}
	if len(m.UE_Time_Zone) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(len(m.UE_Time_Zone)))
		i += copy(dAtA[i:], m.UE_Time_Zone)
	}
	if len(m.AN_Type) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(len(m.AN_Type)))
		i += copy(dAtA[i:], m.AN_Type)
	}
	if m.H_SMF_ID != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(m.H_SMF_ID))
	}
	if m.Old_PDU_Session_ID != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(m.Old_PDU_Session_ID))
	}
	return i, nil
}

func (m *CreateSMContextResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateSMContextResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Result_Indication) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(len(m.Result_Indication)))
		i += copy(dAtA[i:], m.Result_Indication)
	}
	if m.Cause != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(m.Cause))
	}
	if m.PDU_Session_ID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(m.PDU_Session_ID))
	}
	if len(m.N2_SM_Info) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(len(m.N2_SM_Info)))
		i += copy(dAtA[i:], m.N2_SM_Info)
	}
	if len(m.N1_SM_Container) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintNsmf_PDUSession(dAtA, i, uint64(len(m.N1_SM_Container)))
		i += copy(dAtA[i:], m.N1_SM_Container)
	}
	return i, nil
}

func encodeVarintNsmf_PDUSession(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CreateSMContextRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.SUPI)
	if l > 0 {
		n += 1 + l + sovNsmf_PDUSession(uint64(l))
	}
	l = len(m.DNN)
	if l > 0 {
		n += 1 + l + sovNsmf_PDUSession(uint64(l))
	}
	if m.S_NSSAI != 0 {
		n += 1 + sovNsmf_PDUSession(uint64(m.S_NSSAI))
	}
	if m.PDU_Session_ID != 0 {
		n += 1 + sovNsmf_PDUSession(uint64(m.PDU_Session_ID))
	}
	if m.AMF_ID != 0 {
		n += 1 + sovNsmf_PDUSession(uint64(m.AMF_ID))
	}
	l = len(m.N1_SM_Container)
	if l > 0 {
		n += 1 + l + sovNsmf_PDUSession(uint64(l))
	}
	l = len(m.PEI)
	if l > 0 {
		n += 1 + l + sovNsmf_PDUSession(uint64(l))
	}
	l = len(m.UE_Location)
	if l > 0 {
		n += 1 + l + sovNsmf_PDUSession(uint64(l))
	}
	l = len(m.UE_Time_Zone)
	if l > 0 {
		n += 1 + l + sovNsmf_PDUSession(uint64(l))
	}
	l = len(m.AN_Type)
	if l > 0 {
		n += 1 + l + sovNsmf_PDUSession(uint64(l))
	}
	if m.H_SMF_ID != 0 {
		n += 1 + sovNsmf_PDUSession(uint64(m.H_SMF_ID))
	}
	if m.Old_PDU_Session_ID != 0 {
		n += 1 + sovNsmf_PDUSession(uint64(m.Old_PDU_Session_ID))
	}
	return n
}

func (m *CreateSMContextResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Result_Indication)
	if l > 0 {
		n += 1 + l + sovNsmf_PDUSession(uint64(l))
	}
	if m.Cause != 0 {
		n += 1 + sovNsmf_PDUSession(uint64(m.Cause))
	}
	if m.PDU_Session_ID != 0 {
		n += 1 + sovNsmf_PDUSession(uint64(m.PDU_Session_ID))
	}
	l = len(m.N2_SM_Info)
	if l > 0 {
		n += 1 + l + sovNsmf_PDUSession(uint64(l))
	}
	l = len(m.N1_SM_Container)
	if l > 0 {
		n += 1 + l + sovNsmf_PDUSession(uint64(l))
	}
	return n
}

func sovNsmf_PDUSession(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNsmf_PDUSession(x uint64) (n int) {
	return sovNsmf_PDUSession(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CreateSMContextRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateSMContextRequest{`,
		`SUPI:` + fmt.Sprintf("%v", this.SUPI) + `,`,
		`DNN:` + fmt.Sprintf("%v", this.DNN) + `,`,
		`S_NSSAI:` + fmt.Sprintf("%v", this.S_NSSAI) + `,`,
		`PDU_Session_ID:` + fmt.Sprintf("%v", this.PDU_Session_ID) + `,`,
		`AMF_ID:` + fmt.Sprintf("%v", this.AMF_ID) + `,`,
		`N1_SM_Container:` + fmt.Sprintf("%v", this.N1_SM_Container) + `,`,
		`PEI:` + fmt.Sprintf("%v", this.PEI) + `,`,
		`UE_Location:` + fmt.Sprintf("%v", this.UE_Location) + `,`,
		`UE_Time_Zone:` + fmt.Sprintf("%v", this.UE_Time_Zone) + `,`,
		`AN_Type:` + fmt.Sprintf("%v", this.AN_Type) + `,`,
		`H_SMF_ID:` + fmt.Sprintf("%v", this.H_SMF_ID) + `,`,
		`Old_PDU_Session_ID:` + fmt.Sprintf("%v", this.Old_PDU_Session_ID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CreateSMContextResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateSMContextResponse{`,
		`Result_Indication:` + fmt.Sprintf("%v", this.Result_Indication) + `,`,
		`Cause:` + fmt.Sprintf("%v", this.Cause) + `,`,
		`PDU_Session_ID:` + fmt.Sprintf("%v", this.PDU_Session_ID) + `,`,
		`N2_SM_Info:` + fmt.Sprintf("%v", this.N2_SM_Info) + `,`,
		`N1_SM_Container:` + fmt.Sprintf("%v", this.N1_SM_Container) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringNsmf_PDUSession(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CreateSMContextRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNsmf_PDUSession
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateSMContextRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateSMContextRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SUPI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SUPI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DNN", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DNN = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field S_NSSAI", wireType)
			}
			m.S_NSSAI = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.S_NSSAI |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PDU_Session_ID", wireType)
			}
			m.PDU_Session_ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PDU_Session_ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AMF_ID", wireType)
			}
			m.AMF_ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AMF_ID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field N1_SM_Container", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.N1_SM_Container = append(m.N1_SM_Container[:0], dAtA[iNdEx:postIndex]...)
			if m.N1_SM_Container == nil {
				m.N1_SM_Container = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PEI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PEI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UE_Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UE_Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UE_Time_Zone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UE_Time_Zone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AN_Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AN_Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field H_SMF_ID", wireType)
			}
			m.H_SMF_ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.H_SMF_ID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Old_PDU_Session_ID", wireType)
			}
			m.Old_PDU_Session_ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Old_PDU_Session_ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNsmf_PDUSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateSMContextResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNsmf_PDUSession
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateSMContextResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateSMContextResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result_Indication", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result_Indication = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cause", wireType)
			}
			m.Cause = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cause |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PDU_Session_ID", wireType)
			}
			m.PDU_Session_ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PDU_Session_ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field N2_SM_Info", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.N2_SM_Info = append(m.N2_SM_Info[:0], dAtA[iNdEx:postIndex]...)
			if m.N2_SM_Info == nil {
				m.N2_SM_Info = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field N1_SM_Container", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.N1_SM_Container = append(m.N1_SM_Container[:0], dAtA[iNdEx:postIndex]...)
			if m.N1_SM_Container == nil {
				m.N1_SM_Container = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNsmf_PDUSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNsmf_PDUSession
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNsmf_PDUSession(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNsmf_PDUSession
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNsmf_PDUSession
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNsmf_PDUSession
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNsmf_PDUSession
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNsmf_PDUSession(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNsmf_PDUSession = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNsmf_PDUSession   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("Nsmf_PDUSession.proto", fileDescriptorNsmf_PDUSession) }

var fileDescriptorNsmf_PDUSession = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xeb, 0xf5, 0xcf, 0xba, 0x77, 0x29, 0x1d, 0x16, 0x63, 0x16, 0x4c, 0x26, 0xaa, 0x10,
	0xaa, 0x34, 0x54, 0x69, 0xe5, 0x13, 0x84, 0x26, 0xd3, 0x22, 0x11, 0x2f, 0x8a, 0x97, 0x0b, 0x17,
	0x2b, 0xb4, 0xde, 0x14, 0xa9, 0x73, 0x4a, 0x93, 0x4a, 0x70, 0xe3, 0x23, 0xf0, 0x31, 0xf8, 0x28,
	0x3b, 0xee, 0xc8, 0x91, 0x86, 0x0b, 0xc7, 0xf1, 0x0d, 0x90, 0x9d, 0x4d, 0x43, 0x5d, 0x6e, 0x7e,
	0x9f, 0x27, 0x91, 0x7f, 0xcf, 0xe3, 0x17, 0xf6, 0x59, 0x7e, 0x75, 0x21, 0x42, 0x37, 0xe6, 0x32,
	0xcf, 0xd3, 0x4c, 0x8d, 0x16, 0xcb, 0xac, 0xc8, 0x70, 0x53, 0x5d, 0x4e, 0x07, 0x7f, 0xb7, 0xe0,
	0xf9, 0x64, 0x29, 0x93, 0x42, 0xf2, 0x60, 0x92, 0xa9, 0x42, 0x7e, 0x29, 0x22, 0xf9, 0x79, 0x25,
	0xf3, 0x02, 0x63, 0x68, 0xf1, 0x38, 0xf4, 0xc9, 0x96, 0x8d, 0x86, 0x3b, 0x91, 0x39, 0xe3, 0x3d,
	0x68, 0xba, 0x8c, 0x91, 0xa6, 0x91, 0xf4, 0x11, 0x1f, 0xc0, 0x36, 0x17, 0x8c, 0x73, 0xc7, 0x27,
	0x2d, 0x1b, 0x0d, 0x7b, 0x51, 0x87, 0x9b, 0x09, 0xbf, 0x86, 0x27, 0xa1, 0x1b, 0x8b, 0xbb, 0x3b,
	0x85, 0xef, 0x92, 0xb6, 0x8d, 0x86, 0xad, 0xc8, 0x7a, 0x00, 0xf1, 0x5d, 0xbc, 0x0f, 0x1d, 0x27,
	0x38, 0xd1, 0x6e, 0xc7, 0xfc, 0xdd, 0x76, 0x82, 0x13, 0xdf, 0xc5, 0x6f, 0xa0, 0xcf, 0x8e, 0x05,
	0x0f, 0x84, 0x66, 0x4a, 0x52, 0x25, 0x97, 0x64, 0xdb, 0x46, 0x43, 0x2b, 0xea, 0xb1, 0xe3, 0x0a,
	0xd4, 0x88, 0x9a, 0x27, 0xf4, 0x7c, 0xd2, 0xad, 0x78, 0x42, 0xcf, 0xc7, 0xaf, 0x60, 0x37, 0xf6,
	0xc4, 0x87, 0x6c, 0x9a, 0x14, 0x69, 0xa6, 0xc8, 0x8e, 0x71, 0x20, 0xf6, 0xee, 0x15, 0x6c, 0x83,
	0x15, 0x7b, 0xe2, 0x3c, 0xbd, 0x92, 0xe2, 0x63, 0xa6, 0x24, 0x81, 0xfb, 0x2f, 0xb4, 0xa4, 0x15,
	0x1d, 0xc9, 0x61, 0xe2, 0xfc, 0xeb, 0x42, 0x92, 0x5d, 0x63, 0x76, 0x1c, 0xa6, 0x27, 0x4c, 0xa0,
	0x7b, 0x2a, 0x78, 0x85, 0x6b, 0x55, 0x61, 0x4f, 0xb9, 0xe1, 0x3d, 0x02, 0x7c, 0x36, 0x9f, 0x89,
	0x8d, 0xc0, 0x3d, 0x13, 0xb8, 0x7f, 0x36, 0x9f, 0xfd, 0x9f, 0x79, 0x70, 0x8d, 0xe0, 0xe0, 0x51,
	0xe7, 0xf9, 0x22, 0x53, 0xb9, 0xc4, 0x47, 0xf0, 0x34, 0x92, 0xf9, 0x6a, 0x5e, 0x08, 0x5f, 0xcd,
	0xd2, 0xbb, 0x10, 0xc8, 0x50, 0xec, 0x55, 0xc6, 0x83, 0x8e, 0x9f, 0x41, 0x7b, 0x92, 0xac, 0x72,
	0x69, 0x9e, 0xa8, 0x17, 0x55, 0x43, 0x4d, 0xf1, 0xcd, 0x9a, 0xe2, 0x0f, 0x01, 0xd8, 0x58, 0x37,
	0xec, 0xab, 0x8b, 0xcc, 0x3c, 0x9d, 0x15, 0x75, 0xd9, 0x98, 0x07, 0x7a, 0xae, 0xeb, 0xbf, 0x5d,
	0xd3, 0xff, 0x38, 0x81, 0xfe, 0xc6, 0x72, 0x61, 0x06, 0xfd, 0x8d, 0x70, 0xf8, 0xe5, 0x48, 0x5d,
	0x4e, 0x47, 0xf5, 0x6b, 0xf6, 0xe2, 0xb0, 0xde, 0xac, 0xfa, 0x18, 0x34, 0xde, 0xbf, 0xbd, 0x59,
	0xd3, 0xc6, 0xcf, 0x35, 0x6d, 0xdc, 0xae, 0x29, 0xfa, 0x56, 0x52, 0xf4, 0xa3, 0xa4, 0xe8, 0xba,
	0xa4, 0xe8, 0xa6, 0xa4, 0xe8, 0x57, 0x49, 0xd1, 0x9f, 0x92, 0x36, 0x6e, 0x4b, 0x8a, 0xbe, 0xff,
	0xa6, 0x8d, 0x4f, 0x1d, 0xb3, 0xdb, 0xef, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x26, 0x21, 0xc4,
	0x41, 0xf4, 0x02, 0x00, 0x00,
}
