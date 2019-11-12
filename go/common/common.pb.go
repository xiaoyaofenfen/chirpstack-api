// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Modulation int32

const (
	// LoRa
	Modulation_LORA Modulation = 0
	// FSK
	Modulation_FSK Modulation = 1
)

var Modulation_name = map[int32]string{
	0: "LORA",
	1: "FSK",
}

var Modulation_value = map[string]int32{
	"LORA": 0,
	"FSK":  1,
}

func (x Modulation) String() string {
	return proto.EnumName(Modulation_name, int32(x))
}

func (Modulation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type Region int32

const (
	// EU868
	Region_EU868 Region = 0
	// US915
	Region_US915 Region = 2
	// CN779
	Region_CN779 Region = 3
	// EU433
	Region_EU433 Region = 4
	// AU915
	Region_AU915 Region = 5
	// CN470
	Region_CN470 Region = 6
	// AS923
	Region_AS923 Region = 7
	// KR920
	Region_KR920 Region = 8
	// IN865
	Region_IN865 Region = 9
	// RU864
	Region_RU864 Region = 10
	// Experimental 2.4GHz.
	Region_EXP24GHZ Region = 99
)

var Region_name = map[int32]string{
	0:  "EU868",
	2:  "US915",
	3:  "CN779",
	4:  "EU433",
	5:  "AU915",
	6:  "CN470",
	7:  "AS923",
	8:  "KR920",
	9:  "IN865",
	10: "RU864",
	99: "EXP24GHZ",
}

var Region_value = map[string]int32{
	"EU868":    0,
	"US915":    2,
	"CN779":    3,
	"EU433":    4,
	"AU915":    5,
	"CN470":    6,
	"AS923":    7,
	"KR920":    8,
	"IN865":    9,
	"RU864":    10,
	"EXP24GHZ": 99,
}

func (x Region) String() string {
	return proto.EnumName(Region_name, int32(x))
}

func (Region) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

type LocationSource int32

const (
	// Unknown.
	LocationSource_UNKNOWN LocationSource = 0
	// GPS.
	LocationSource_GPS LocationSource = 1
	// Manually configured.
	LocationSource_CONFIG LocationSource = 2
	// Geo resolver.
	LocationSource_GEO_RESOLVER LocationSource = 3
)

var LocationSource_name = map[int32]string{
	0: "UNKNOWN",
	1: "GPS",
	2: "CONFIG",
	3: "GEO_RESOLVER",
}

var LocationSource_value = map[string]int32{
	"UNKNOWN":      0,
	"GPS":          1,
	"CONFIG":       2,
	"GEO_RESOLVER": 3,
}

func (x LocationSource) String() string {
	return proto.EnumName(LocationSource_name, int32(x))
}

func (LocationSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

type KeyEnvelope struct {
	// KEK label.
	KekLabel string `protobuf:"bytes,1,opt,name=kek_label,json=kekLabel,proto3" json:"kek_label,omitempty"`
	// AES key (when the kek_label is set, this key is encrypted using a key
	// known to the join-server and application-server.
	// For more information please refer to the LoRaWAN Backend Interface
	// 'Key Transport Security' section.
	AesKey               []byte   `protobuf:"bytes,2,opt,name=aes_key,json=aesKey,proto3" json:"aes_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyEnvelope) Reset()         { *m = KeyEnvelope{} }
func (m *KeyEnvelope) String() string { return proto.CompactTextString(m) }
func (*KeyEnvelope) ProtoMessage()    {}
func (*KeyEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *KeyEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyEnvelope.Unmarshal(m, b)
}
func (m *KeyEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyEnvelope.Marshal(b, m, deterministic)
}
func (m *KeyEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyEnvelope.Merge(m, src)
}
func (m *KeyEnvelope) XXX_Size() int {
	return xxx_messageInfo_KeyEnvelope.Size(m)
}
func (m *KeyEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_KeyEnvelope proto.InternalMessageInfo

func (m *KeyEnvelope) GetKekLabel() string {
	if m != nil {
		return m.KekLabel
	}
	return ""
}

func (m *KeyEnvelope) GetAesKey() []byte {
	if m != nil {
		return m.AesKey
	}
	return nil
}

type Location struct {
	// Latitude.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude.
	Altitude float64 `protobuf:"fixed64,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Location source.
	Source LocationSource `protobuf:"varint,4,opt,name=source,proto3,enum=common.LocationSource" json:"source,omitempty"`
	// Accuracy (in meters).
	Accuracy             uint32   `protobuf:"varint,5,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Location) GetSource() LocationSource {
	if m != nil {
		return m.Source
	}
	return LocationSource_UNKNOWN
}

func (m *Location) GetAccuracy() uint32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func init() {
	proto.RegisterEnum("common.Modulation", Modulation_name, Modulation_value)
	proto.RegisterEnum("common.Region", Region_name, Region_value)
	proto.RegisterEnum("common.LocationSource", LocationSource_name, LocationSource_value)
	proto.RegisterType((*KeyEnvelope)(nil), "common.KeyEnvelope")
	proto.RegisterType((*Location)(nil), "common.Location")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xcd, 0xf2, 0x61, 0xec, 0x09, 0x8d, 0x46, 0x5b, 0xa9, 0x45, 0x6d, 0xa5, 0xa2, 0x9c, 0x10,
	0x55, 0x20, 0x05, 0xc2, 0xc7, 0x31, 0xa5, 0x0e, 0x8d, 0x4c, 0xed, 0x68, 0x2d, 0xb7, 0x55, 0x2e,
	0x68, 0xd9, 0xac, 0x88, 0x65, 0x87, 0x45, 0xc6, 0xae, 0xc4, 0xbd, 0x3f, 0xa7, 0x3f, 0xb2, 0x5a,
	0xdb, 0xa1, 0xca, 0xc9, 0xef, 0xcd, 0x7b, 0x6f, 0x3c, 0xb3, 0x1a, 0x78, 0x2d, 0xd4, 0xd3, 0x93,
	0xda, 0xf6, 0x8b, 0x4f, 0x6f, 0x97, 0xa8, 0x54, 0x51, 0xa3, 0x60, 0xe7, 0x73, 0x38, 0x75, 0xe4,
	0xc1, 0xde, 0xfe, 0x96, 0xb1, 0xda, 0x49, 0xfa, 0x1e, 0xac, 0x48, 0x46, 0xab, 0x98, 0xaf, 0x65,
	0xdc, 0x22, 0x6d, 0xd2, 0xb1, 0x98, 0x19, 0xc9, 0x68, 0xa9, 0x39, 0x7d, 0x0b, 0x0d, 0x2e, 0xf7,
	0xab, 0x48, 0x1e, 0x5a, 0x95, 0x36, 0xe9, 0x34, 0x99, 0xc1, 0xe5, 0xde, 0x91, 0x87, 0xf3, 0xbf,
	0x04, 0xcc, 0xa5, 0x12, 0x3c, 0x0d, 0xd5, 0x96, 0xbe, 0x03, 0x33, 0xe6, 0x69, 0x98, 0x66, 0x0f,
	0x32, 0xef, 0x40, 0xd8, 0x91, 0xd3, 0x0f, 0x60, 0xc5, 0x6a, 0xbb, 0x29, 0xc4, 0x4a, 0x2e, 0xfe,
	0x2f, 0xe8, 0x24, 0x8f, 0xcb, 0x64, 0xb5, 0x48, 0x3e, 0x73, 0xda, 0x03, 0x63, 0xaf, 0xb2, 0x44,
	0xc8, 0x56, 0xad, 0x4d, 0x3a, 0x67, 0x83, 0x37, 0xbd, 0x72, 0x9d, 0xe7, 0xff, 0xfa, 0xb9, 0xca,
	0x4a, 0x57, 0xde, 0x4b, 0x88, 0x2c, 0xe1, 0xe2, 0xd0, 0xaa, 0xb7, 0x49, 0xe7, 0x15, 0x3b, 0xf2,
	0xee, 0x47, 0x80, 0xef, 0xea, 0x21, 0x8b, 0x8b, 0x79, 0x4d, 0xa8, 0x2d, 0x3d, 0x76, 0x8d, 0x27,
	0xb4, 0x01, 0xd5, 0x1b, 0xdf, 0x41, 0xd2, 0xfd, 0x43, 0xc0, 0x60, 0x72, 0xa3, 0x55, 0x0b, 0xea,
	0x76, 0x30, 0x1d, 0x4f, 0xf1, 0x44, 0xc3, 0xc0, 0x9f, 0x7d, 0xbe, 0xc2, 0x8a, 0x86, 0x73, 0x77,
	0x32, 0x99, 0x61, 0xb5, 0x30, 0x8c, 0x86, 0x43, 0xac, 0x69, 0x78, 0x1d, 0x68, 0x43, 0xbd, 0x30,
	0x8c, 0x26, 0x97, 0x68, 0xe4, 0x55, 0x7f, 0x36, 0x18, 0x62, 0x43, 0x43, 0x87, 0xcd, 0x06, 0x97,
	0x68, 0x6a, 0x78, 0xeb, 0x4e, 0xc7, 0x57, 0x68, 0x69, 0xc8, 0x82, 0xe9, 0x78, 0x84, 0x40, 0x9b,
	0x60, 0xda, 0xbf, 0xee, 0x06, 0xa3, 0xc5, 0xb7, 0x7b, 0x14, 0xdd, 0xaf, 0x70, 0xf6, 0x72, 0x3b,
	0x7a, 0x0a, 0x8d, 0xc0, 0x75, 0x5c, 0xef, 0xa7, 0x5b, 0x8c, 0xbb, 0xb8, 0xf3, 0x91, 0x50, 0x00,
	0x63, 0xee, 0xb9, 0x37, 0xb7, 0x0b, 0xac, 0x50, 0x84, 0xe6, 0xc2, 0xf6, 0x56, 0xcc, 0xf6, 0xbd,
	0xe5, 0x0f, 0x9b, 0x61, 0xf5, 0xcb, 0xc5, 0xfd, 0xa7, 0x4d, 0x98, 0x3e, 0x66, 0x6b, 0xfd, 0x62,
	0xfd, 0x75, 0xa2, 0x04, 0xe7, 0x49, 0x5f, 0x3c, 0x86, 0xc9, 0x6e, 0x9f, 0x72, 0x11, 0x5d, 0xf0,
	0x5d, 0xd8, 0xdf, 0xa8, 0xf2, 0x3c, 0xd6, 0x46, 0x7e, 0x1f, 0xc3, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x74, 0xaf, 0xf4, 0xe6, 0x36, 0x02, 0x00, 0x00,
}
