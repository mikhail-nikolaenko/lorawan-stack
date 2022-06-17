// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/metadata.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-flags/annotations"
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LocationSource int32

const (
	// The source of the location is not known or not set.
	LocationSource_SOURCE_UNKNOWN LocationSource = 0
	// The location is determined by GPS.
	LocationSource_SOURCE_GPS LocationSource = 1
	// The location is set in and updated from a registry.
	LocationSource_SOURCE_REGISTRY LocationSource = 3
	// The location is estimated with IP geolocation.
	LocationSource_SOURCE_IP_GEOLOCATION LocationSource = 4
	// The location is estimated with WiFi RSSI geolocation.
	LocationSource_SOURCE_WIFI_RSSI_GEOLOCATION LocationSource = 5
	// The location is estimated with BT/BLE RSSI geolocation.
	LocationSource_SOURCE_BT_RSSI_GEOLOCATION LocationSource = 6
	// The location is estimated with LoRa RSSI geolocation.
	LocationSource_SOURCE_LORA_RSSI_GEOLOCATION LocationSource = 7
	// The location is estimated with LoRa TDOA geolocation.
	LocationSource_SOURCE_LORA_TDOA_GEOLOCATION LocationSource = 8
	// The location is estimated by a combination of geolocation sources.
	LocationSource_SOURCE_COMBINED_GEOLOCATION LocationSource = 9
)

var LocationSource_name = map[int32]string{
	0: "SOURCE_UNKNOWN",
	1: "SOURCE_GPS",
	3: "SOURCE_REGISTRY",
	4: "SOURCE_IP_GEOLOCATION",
	5: "SOURCE_WIFI_RSSI_GEOLOCATION",
	6: "SOURCE_BT_RSSI_GEOLOCATION",
	7: "SOURCE_LORA_RSSI_GEOLOCATION",
	8: "SOURCE_LORA_TDOA_GEOLOCATION",
	9: "SOURCE_COMBINED_GEOLOCATION",
}

var LocationSource_value = map[string]int32{
	"SOURCE_UNKNOWN":               0,
	"SOURCE_GPS":                   1,
	"SOURCE_REGISTRY":              3,
	"SOURCE_IP_GEOLOCATION":        4,
	"SOURCE_WIFI_RSSI_GEOLOCATION": 5,
	"SOURCE_BT_RSSI_GEOLOCATION":   6,
	"SOURCE_LORA_RSSI_GEOLOCATION": 7,
	"SOURCE_LORA_TDOA_GEOLOCATION": 8,
	"SOURCE_COMBINED_GEOLOCATION":  9,
}

func (x LocationSource) String() string {
	return proto.EnumName(LocationSource_name, int32(x))
}

func (LocationSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e1123b3e8fd87092, []int{0}
}

// Contains metadata for a received message. Each antenna that receives
// a message corresponds to one RxMetadata.
type RxMetadata struct {
	GatewayIds   *GatewayIdentifiers   `protobuf:"bytes,1,opt,name=gateway_ids,json=gatewayIds,proto3" json:"gateway_ids,omitempty"`
	PacketBroker *PacketBrokerMetadata `protobuf:"bytes,18,opt,name=packet_broker,json=packetBroker,proto3" json:"packet_broker,omitempty"`
	AntennaIndex uint32                `protobuf:"varint,2,opt,name=antenna_index,json=antennaIndex,proto3" json:"antenna_index,omitempty"`
	Time         *types.Timestamp      `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	// Gateway concentrator timestamp when the Rx finished (microseconds).
	Timestamp uint32 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Gateway's internal fine timestamp when the Rx finished (nanoseconds).
	FineTimestamp uint64 `protobuf:"varint,5,opt,name=fine_timestamp,json=fineTimestamp,proto3" json:"fine_timestamp,omitempty"`
	// Encrypted gateway's internal fine timestamp when the Rx finished (nanoseconds).
	EncryptedFineTimestamp      []byte `protobuf:"bytes,6,opt,name=encrypted_fine_timestamp,json=encryptedFineTimestamp,proto3" json:"encrypted_fine_timestamp,omitempty"`
	EncryptedFineTimestampKeyId string `protobuf:"bytes,7,opt,name=encrypted_fine_timestamp_key_id,json=encryptedFineTimestampKeyId,proto3" json:"encrypted_fine_timestamp_key_id,omitempty"`
	// Received signal strength indicator (dBm).
	// This value equals `channel_rssi`.
	Rssi float32 `protobuf:"fixed32,8,opt,name=rssi,proto3" json:"rssi,omitempty"`
	// Received signal strength indicator of the signal (dBm).
	SignalRssi *types.FloatValue `protobuf:"bytes,16,opt,name=signal_rssi,json=signalRssi,proto3" json:"signal_rssi,omitempty"`
	// Received signal strength indicator of the channel (dBm).
	ChannelRssi float32 `protobuf:"fixed32,9,opt,name=channel_rssi,json=channelRssi,proto3" json:"channel_rssi,omitempty"`
	// Standard deviation of the RSSI during preamble.
	RssiStandardDeviation float32 `protobuf:"fixed32,10,opt,name=rssi_standard_deviation,json=rssiStandardDeviation,proto3" json:"rssi_standard_deviation,omitempty"`
	// Signal-to-noise ratio (dB).
	Snr float32 `protobuf:"fixed32,11,opt,name=snr,proto3" json:"snr,omitempty"`
	// Frequency offset (Hz).
	FrequencyOffset int64 `protobuf:"varint,12,opt,name=frequency_offset,json=frequencyOffset,proto3" json:"frequency_offset,omitempty"`
	// Antenna location; injected by the Gateway Server.
	Location *Location `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"`
	// Gateway downlink path constraint; injected by the Gateway Server.
	DownlinkPathConstraint DownlinkPathConstraint `protobuf:"varint,14,opt,name=downlink_path_constraint,json=downlinkPathConstraint,proto3,enum=ttn.lorawan.v3.DownlinkPathConstraint" json:"downlink_path_constraint,omitempty"`
	// Uplink token to be included in the Tx request in class A downlink; injected by gateway, Gateway Server or fNS.
	UplinkToken []byte `protobuf:"bytes,15,opt,name=uplink_token,json=uplinkToken,proto3" json:"uplink_token,omitempty"`
	// Index of the gateway channel that received the message.
	ChannelIndex uint32 `protobuf:"varint,17,opt,name=channel_index,json=channelIndex,proto3" json:"channel_index,omitempty"`
	// Hopping width; a number describing the number of steps of the LR-FHSS grid.
	HoppingWidth uint32 `protobuf:"varint,19,opt,name=hopping_width,json=hoppingWidth,proto3" json:"hopping_width,omitempty"`
	// Frequency drift in Hz between start and end of an LR-FHSS packet (signed).
	FrequencyDrift int32 `protobuf:"varint,20,opt,name=frequency_drift,json=frequencyDrift,proto3" json:"frequency_drift,omitempty"`
	// Advanced metadata fields
	// - can be used for advanced information or experimental features that are not yet formally defined in the API
	// - field names are written in snake_case
	Advanced             *types.Struct `protobuf:"bytes,99,opt,name=advanced,proto3" json:"advanced,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RxMetadata) Reset()         { *m = RxMetadata{} }
func (m *RxMetadata) String() string { return proto.CompactTextString(m) }
func (*RxMetadata) ProtoMessage()    {}
func (*RxMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1123b3e8fd87092, []int{0}
}
func (m *RxMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RxMetadata.Unmarshal(m, b)
}
func (m *RxMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RxMetadata.Marshal(b, m, deterministic)
}
func (m *RxMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RxMetadata.Merge(m, src)
}
func (m *RxMetadata) XXX_Size() int {
	return xxx_messageInfo_RxMetadata.Size(m)
}
func (m *RxMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RxMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RxMetadata proto.InternalMessageInfo

func (m *RxMetadata) GetGatewayIds() *GatewayIdentifiers {
	if m != nil {
		return m.GatewayIds
	}
	return nil
}

func (m *RxMetadata) GetPacketBroker() *PacketBrokerMetadata {
	if m != nil {
		return m.PacketBroker
	}
	return nil
}

func (m *RxMetadata) GetAntennaIndex() uint32 {
	if m != nil {
		return m.AntennaIndex
	}
	return 0
}

func (m *RxMetadata) GetTime() *types.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *RxMetadata) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *RxMetadata) GetFineTimestamp() uint64 {
	if m != nil {
		return m.FineTimestamp
	}
	return 0
}

func (m *RxMetadata) GetEncryptedFineTimestamp() []byte {
	if m != nil {
		return m.EncryptedFineTimestamp
	}
	return nil
}

func (m *RxMetadata) GetEncryptedFineTimestampKeyId() string {
	if m != nil {
		return m.EncryptedFineTimestampKeyId
	}
	return ""
}

func (m *RxMetadata) GetRssi() float32 {
	if m != nil {
		return m.Rssi
	}
	return 0
}

func (m *RxMetadata) GetSignalRssi() *types.FloatValue {
	if m != nil {
		return m.SignalRssi
	}
	return nil
}

func (m *RxMetadata) GetChannelRssi() float32 {
	if m != nil {
		return m.ChannelRssi
	}
	return 0
}

func (m *RxMetadata) GetRssiStandardDeviation() float32 {
	if m != nil {
		return m.RssiStandardDeviation
	}
	return 0
}

func (m *RxMetadata) GetSnr() float32 {
	if m != nil {
		return m.Snr
	}
	return 0
}

func (m *RxMetadata) GetFrequencyOffset() int64 {
	if m != nil {
		return m.FrequencyOffset
	}
	return 0
}

func (m *RxMetadata) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RxMetadata) GetDownlinkPathConstraint() DownlinkPathConstraint {
	if m != nil {
		return m.DownlinkPathConstraint
	}
	return DownlinkPathConstraint_DOWNLINK_PATH_CONSTRAINT_NONE
}

func (m *RxMetadata) GetUplinkToken() []byte {
	if m != nil {
		return m.UplinkToken
	}
	return nil
}

func (m *RxMetadata) GetChannelIndex() uint32 {
	if m != nil {
		return m.ChannelIndex
	}
	return 0
}

func (m *RxMetadata) GetHoppingWidth() uint32 {
	if m != nil {
		return m.HoppingWidth
	}
	return 0
}

func (m *RxMetadata) GetFrequencyDrift() int32 {
	if m != nil {
		return m.FrequencyDrift
	}
	return 0
}

func (m *RxMetadata) GetAdvanced() *types.Struct {
	if m != nil {
		return m.Advanced
	}
	return nil
}

type Location struct {
	// The North–South position (degrees; -90 to +90), where 0 is the equator, North pole is positive, South pole is negative.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// The East-West position (degrees; -180 to +180), where 0 is the Prime Meridian (Greenwich), East is positive , West is negative.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// The altitude (meters), where 0 is the mean sea level.
	Altitude int32 `protobuf:"varint,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// The accuracy of the location (meters).
	Accuracy int32 `protobuf:"varint,4,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	// Source of the location information.
	Source               LocationSource `protobuf:"varint,5,opt,name=source,proto3,enum=ttn.lorawan.v3.LocationSource" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1123b3e8fd87092, []int{1}
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

func (m *Location) GetAltitude() int32 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Location) GetAccuracy() int32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func (m *Location) GetSource() LocationSource {
	if m != nil {
		return m.Source
	}
	return LocationSource_SOURCE_UNKNOWN
}

type PacketBrokerMetadata struct {
	// Message identifier generated by Packet Broker Router.
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// LoRa Alliance NetID of the Packet Broker Forwarder Member.
	ForwarderNetId []byte `protobuf:"bytes,2,opt,name=forwarder_net_id,json=forwarderNetId,proto3" json:"forwarder_net_id,omitempty"`
	// Tenant ID managed by the Packet Broker Forwarder Member.
	ForwarderTenantId string `protobuf:"bytes,3,opt,name=forwarder_tenant_id,json=forwarderTenantId,proto3" json:"forwarder_tenant_id,omitempty"`
	// Forwarder Cluster ID of the Packet Broker Forwarder.
	ForwarderClusterId string `protobuf:"bytes,4,opt,name=forwarder_cluster_id,json=forwarderClusterId,proto3" json:"forwarder_cluster_id,omitempty"`
	// Forwarder gateway EUI.
	ForwarderGatewayEui []byte `protobuf:"bytes,9,opt,name=forwarder_gateway_eui,json=forwarderGatewayEui,proto3" json:"forwarder_gateway_eui,omitempty"`
	// Forwarder gateway ID.
	ForwarderGatewayId *types.StringValue `protobuf:"bytes,10,opt,name=forwarder_gateway_id,json=forwarderGatewayId,proto3" json:"forwarder_gateway_id,omitempty"`
	// LoRa Alliance NetID of the Packet Broker Home Network Member.
	HomeNetworkNetId []byte `protobuf:"bytes,5,opt,name=home_network_net_id,json=homeNetworkNetId,proto3" json:"home_network_net_id,omitempty"`
	// Tenant ID managed by the Packet Broker Home Network Member.
	// This value is empty if it cannot be determined by the Packet Broker Router.
	HomeNetworkTenantId string `protobuf:"bytes,6,opt,name=home_network_tenant_id,json=homeNetworkTenantId,proto3" json:"home_network_tenant_id,omitempty"`
	// Home Network Cluster ID of the Packet Broker Home Network.
	HomeNetworkClusterId string `protobuf:"bytes,8,opt,name=home_network_cluster_id,json=homeNetworkClusterId,proto3" json:"home_network_cluster_id,omitempty"`
	// Hops that the message passed. Each Packet Broker Router service appends an entry.
	Hops                 []*PacketBrokerRouteHop `protobuf:"bytes,7,rep,name=hops,proto3" json:"hops,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PacketBrokerMetadata) Reset()         { *m = PacketBrokerMetadata{} }
func (m *PacketBrokerMetadata) String() string { return proto.CompactTextString(m) }
func (*PacketBrokerMetadata) ProtoMessage()    {}
func (*PacketBrokerMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1123b3e8fd87092, []int{2}
}
func (m *PacketBrokerMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketBrokerMetadata.Unmarshal(m, b)
}
func (m *PacketBrokerMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketBrokerMetadata.Marshal(b, m, deterministic)
}
func (m *PacketBrokerMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketBrokerMetadata.Merge(m, src)
}
func (m *PacketBrokerMetadata) XXX_Size() int {
	return xxx_messageInfo_PacketBrokerMetadata.Size(m)
}
func (m *PacketBrokerMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketBrokerMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_PacketBrokerMetadata proto.InternalMessageInfo

func (m *PacketBrokerMetadata) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *PacketBrokerMetadata) GetForwarderNetId() []byte {
	if m != nil {
		return m.ForwarderNetId
	}
	return nil
}

func (m *PacketBrokerMetadata) GetForwarderTenantId() string {
	if m != nil {
		return m.ForwarderTenantId
	}
	return ""
}

func (m *PacketBrokerMetadata) GetForwarderClusterId() string {
	if m != nil {
		return m.ForwarderClusterId
	}
	return ""
}

func (m *PacketBrokerMetadata) GetForwarderGatewayEui() []byte {
	if m != nil {
		return m.ForwarderGatewayEui
	}
	return nil
}

func (m *PacketBrokerMetadata) GetForwarderGatewayId() *types.StringValue {
	if m != nil {
		return m.ForwarderGatewayId
	}
	return nil
}

func (m *PacketBrokerMetadata) GetHomeNetworkNetId() []byte {
	if m != nil {
		return m.HomeNetworkNetId
	}
	return nil
}

func (m *PacketBrokerMetadata) GetHomeNetworkTenantId() string {
	if m != nil {
		return m.HomeNetworkTenantId
	}
	return ""
}

func (m *PacketBrokerMetadata) GetHomeNetworkClusterId() string {
	if m != nil {
		return m.HomeNetworkClusterId
	}
	return ""
}

func (m *PacketBrokerMetadata) GetHops() []*PacketBrokerRouteHop {
	if m != nil {
		return m.Hops
	}
	return nil
}

type PacketBrokerRouteHop struct {
	// Time when the service received the message.
	ReceivedAt *types.Timestamp `protobuf:"bytes,1,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
	// Sender of the message, typically the authorized client identifier.
	SenderName string `protobuf:"bytes,2,opt,name=sender_name,json=senderName,proto3" json:"sender_name,omitempty"`
	// Sender IP address or host name.
	SenderAddress string `protobuf:"bytes,3,opt,name=sender_address,json=senderAddress,proto3" json:"sender_address,omitempty"`
	// Receiver of the message.
	ReceiverName string `protobuf:"bytes,4,opt,name=receiver_name,json=receiverName,proto3" json:"receiver_name,omitempty"`
	// Receiver agent.
	ReceiverAgent        string   `protobuf:"bytes,5,opt,name=receiver_agent,json=receiverAgent,proto3" json:"receiver_agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketBrokerRouteHop) Reset()         { *m = PacketBrokerRouteHop{} }
func (m *PacketBrokerRouteHop) String() string { return proto.CompactTextString(m) }
func (*PacketBrokerRouteHop) ProtoMessage()    {}
func (*PacketBrokerRouteHop) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1123b3e8fd87092, []int{3}
}
func (m *PacketBrokerRouteHop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketBrokerRouteHop.Unmarshal(m, b)
}
func (m *PacketBrokerRouteHop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketBrokerRouteHop.Marshal(b, m, deterministic)
}
func (m *PacketBrokerRouteHop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketBrokerRouteHop.Merge(m, src)
}
func (m *PacketBrokerRouteHop) XXX_Size() int {
	return xxx_messageInfo_PacketBrokerRouteHop.Size(m)
}
func (m *PacketBrokerRouteHop) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketBrokerRouteHop.DiscardUnknown(m)
}

var xxx_messageInfo_PacketBrokerRouteHop proto.InternalMessageInfo

func (m *PacketBrokerRouteHop) GetReceivedAt() *types.Timestamp {
	if m != nil {
		return m.ReceivedAt
	}
	return nil
}

func (m *PacketBrokerRouteHop) GetSenderName() string {
	if m != nil {
		return m.SenderName
	}
	return ""
}

func (m *PacketBrokerRouteHop) GetSenderAddress() string {
	if m != nil {
		return m.SenderAddress
	}
	return ""
}

func (m *PacketBrokerRouteHop) GetReceiverName() string {
	if m != nil {
		return m.ReceiverName
	}
	return ""
}

func (m *PacketBrokerRouteHop) GetReceiverAgent() string {
	if m != nil {
		return m.ReceiverAgent
	}
	return ""
}

func init() {
	proto.RegisterEnum("ttn.lorawan.v3.LocationSource", LocationSource_name, LocationSource_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.LocationSource", LocationSource_name, LocationSource_value)
	proto.RegisterType((*RxMetadata)(nil), "ttn.lorawan.v3.RxMetadata")
	golang_proto.RegisterType((*RxMetadata)(nil), "ttn.lorawan.v3.RxMetadata")
	proto.RegisterType((*Location)(nil), "ttn.lorawan.v3.Location")
	golang_proto.RegisterType((*Location)(nil), "ttn.lorawan.v3.Location")
	proto.RegisterType((*PacketBrokerMetadata)(nil), "ttn.lorawan.v3.PacketBrokerMetadata")
	golang_proto.RegisterType((*PacketBrokerMetadata)(nil), "ttn.lorawan.v3.PacketBrokerMetadata")
	proto.RegisterType((*PacketBrokerRouteHop)(nil), "ttn.lorawan.v3.PacketBrokerRouteHop")
	golang_proto.RegisterType((*PacketBrokerRouteHop)(nil), "ttn.lorawan.v3.PacketBrokerRouteHop")
}

func init() { proto.RegisterFile("lorawan-stack/api/metadata.proto", fileDescriptor_e1123b3e8fd87092) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/metadata.proto", fileDescriptor_e1123b3e8fd87092)
}

var fileDescriptor_e1123b3e8fd87092 = []byte{
	// 1429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x3d, 0x6f, 0xdb, 0x46,
	0x1f, 0x0f, 0x2d, 0xd9, 0x96, 0x4e, 0x96, 0xa2, 0x9c, 0x9d, 0x98, 0x76, 0x1c, 0x5b, 0x8f, 0xf3,
	0xbc, 0x28, 0x01, 0x2c, 0x3d, 0xb0, 0x9b, 0x22, 0x68, 0x0b, 0xb4, 0x96, 0xdf, 0x42, 0x24, 0x96,
	0x8c, 0x93, 0x92, 0xb4, 0x5d, 0x88, 0x33, 0x79, 0xa2, 0x18, 0x49, 0x77, 0x2c, 0xef, 0x68, 0x47,
	0x9d, 0x82, 0x8c, 0x9d, 0xba, 0x77, 0xea, 0xea, 0xaf, 0xd0, 0xc5, 0x4b, 0xbf, 0x42, 0xbf, 0x40,
	0x87, 0x02, 0x19, 0x33, 0x66, 0x69, 0xc1, 0xe3, 0x91, 0x7a, 0x73, 0x5a, 0x64, 0x8b, 0x26, 0xf2,
	0xf7, 0xc6, 0xe3, 0xff, 0xfe, 0xf7, 0x17, 0x41, 0xa9, 0xc7, 0x7c, 0x7c, 0x8e, 0xe9, 0x16, 0x17,
	0xd8, 0xea, 0x56, 0xb1, 0xe7, 0x56, 0xfb, 0x44, 0x60, 0x1b, 0x0b, 0x5c, 0xf1, 0x7c, 0x26, 0x18,
	0x2c, 0x08, 0x41, 0x2b, 0x4a, 0x55, 0x39, 0xdb, 0x59, 0xdd, 0x77, 0x5c, 0xd1, 0x09, 0x4e, 0x2b,
	0x16, 0xeb, 0x57, 0x5b, 0x1d, 0xd2, 0xea, 0xb8, 0xd4, 0xe1, 0x06, 0xb5, 0x03, 0x2e, 0x7c, 0x97,
	0xf0, 0xaa, 0x74, 0x59, 0x5b, 0x0e, 0xa1, 0x5b, 0x0e, 0xdb, 0x6a, 0xf7, 0xb0, 0xc3, 0xab, 0x98,
	0x52, 0x26, 0xb0, 0x70, 0x19, 0xe5, 0x51, 0xea, 0xea, 0xee, 0x48, 0x0a, 0xa1, 0x67, 0x6c, 0xe0,
	0xf9, 0xec, 0xe5, 0x60, 0xd4, 0x7c, 0x86, 0x7b, 0xae, 0x8d, 0x05, 0xa9, 0x4e, 0x5d, 0xa8, 0x88,
	0xad, 0x91, 0x08, 0x87, 0x39, 0x2c, 0x32, 0x9f, 0x06, 0x6d, 0x79, 0x27, 0x6f, 0xe4, 0x95, 0x92,
	0xef, 0x7d, 0xd0, 0xba, 0x5f, 0x70, 0x46, 0xaf, 0x58, 0xf6, 0x9a, 0xc3, 0x98, 0xd3, 0x23, 0xc3,
	0x47, 0x71, 0xe1, 0x07, 0x96, 0x50, 0xec, 0xc6, 0x24, 0x2b, 0xdc, 0x3e, 0xe1, 0x02, 0xf7, 0x3d,
	0x25, 0x58, 0x9f, 0x14, 0x9c, 0xfb, 0xd8, 0xf3, 0x88, 0x1f, 0xc7, 0xdf, 0x99, 0xde, 0x0d, 0x42,
	0x83, 0x7e, 0x4c, 0xdf, 0x9d, 0xa6, 0x5d, 0x9b, 0x50, 0xe1, 0xb6, 0xdd, 0x24, 0x63, 0xf3, 0xe7,
	0x0c, 0x00, 0xe8, 0xe5, 0xb1, 0xda, 0x44, 0x78, 0x0c, 0x72, 0x0e, 0x16, 0xe4, 0x1c, 0x0f, 0x4c,
	0xd7, 0xe6, 0xba, 0x56, 0xd2, 0xca, 0xb9, 0xed, 0xcd, 0xca, 0xf8, 0xa6, 0x56, 0x8e, 0x22, 0x89,
	0x31, 0x4c, 0xab, 0x65, 0xde, 0xd5, 0x66, 0x7f, 0xd0, 0x66, 0x8a, 0x1a, 0x02, 0x4e, 0xcc, 0x72,
	0x68, 0x80, 0xbc, 0x87, 0xad, 0x2e, 0x11, 0xe6, 0xa9, 0xcf, 0xba, 0xc4, 0xd7, 0xa1, 0x0c, 0xfc,
	0xf7, 0x64, 0xe0, 0x89, 0x14, 0xd5, 0xa4, 0x26, 0x5e, 0x0b, 0x5a, 0xf0, 0x46, 0x50, 0x78, 0x17,
	0xe4, 0x31, 0x15, 0x84, 0x52, 0x6c, 0xba, 0xd4, 0x26, 0x2f, 0xf5, 0x99, 0x92, 0x56, 0xce, 0xa3,
	0x05, 0x05, 0x1a, 0x21, 0x06, 0x2b, 0x20, 0x1d, 0x16, 0x51, 0x4f, 0xc9, 0xc7, 0xac, 0x56, 0xa2,
	0x02, 0x56, 0xe2, 0x02, 0x56, 0x5a, 0x71, 0x85, 0x91, 0xd4, 0xc1, 0x35, 0x90, 0x4d, 0x8a, 0xae,
	0xa7, 0x65, 0xe0, 0x10, 0x80, 0xff, 0x01, 0x85, 0xb6, 0x4b, 0x89, 0x39, 0x94, 0xcc, 0x96, 0xb4,
	0x72, 0x1a, 0xe5, 0x43, 0x34, 0x89, 0x82, 0x0f, 0x81, 0x4e, 0xa8, 0xe5, 0x0f, 0x3c, 0x41, 0x6c,
	0x73, 0xc2, 0x30, 0x57, 0xd2, 0xca, 0x0b, 0xe8, 0x56, 0xc2, 0x1f, 0x8e, 0x39, 0xf7, 0xc1, 0xc6,
	0xfb, 0x9c, 0x66, 0x97, 0x84, 0x5b, 0xa0, 0xcf, 0x97, 0xb4, 0x72, 0x16, 0xdd, 0xbe, 0x3a, 0xe0,
	0x31, 0x19, 0x18, 0x36, 0x84, 0x20, 0xed, 0x73, 0xee, 0xea, 0x99, 0x92, 0x56, 0x9e, 0x41, 0xf2,
	0x1a, 0x7e, 0x01, 0x72, 0xdc, 0x75, 0x28, 0xee, 0x99, 0x92, 0x2a, 0xca, 0x7a, 0xdc, 0x9e, 0xaa,
	0xc7, 0x61, 0x8f, 0x61, 0xf1, 0x0c, 0xf7, 0x02, 0x82, 0x40, 0xa4, 0x47, 0xa1, 0xfb, 0x5f, 0x60,
	0xc1, 0xea, 0x60, 0x4a, 0x89, 0xb2, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x52, 0xf2, 0x29, 0x58, 0x0e,
	0x29, 0x93, 0x0b, 0x4c, 0x6d, 0xec, 0xdb, 0xa6, 0x4d, 0xce, 0x5c, 0xd9, 0xfc, 0x3a, 0x90, 0xea,
	0x9b, 0x21, 0xdd, 0x54, 0xec, 0x7e, 0x4c, 0xc2, 0x22, 0x48, 0x71, 0xea, 0xeb, 0x39, 0xa9, 0x09,
	0x2f, 0xe1, 0x3d, 0x50, 0x6c, 0xfb, 0xe4, 0xbb, 0x80, 0x50, 0x6b, 0x60, 0xb2, 0x76, 0x9b, 0x13,
	0xa1, 0x2f, 0x94, 0xb4, 0x72, 0x0a, 0x5d, 0x4f, 0xf0, 0x86, 0x84, 0xe1, 0x27, 0x20, 0xd3, 0x63,
	0x56, 0xf4, 0x94, 0xbc, 0x7c, 0x25, 0x7d, 0xb2, 0x93, 0x9e, 0x28, 0x1e, 0x25, 0x4a, 0xf8, 0x02,
	0xe8, 0x36, 0x3b, 0xa7, 0x3d, 0x97, 0x76, 0x4d, 0x0f, 0x8b, 0x8e, 0x69, 0x31, 0xca, 0x85, 0x8f,
	0x5d, 0x2a, 0xf4, 0x42, 0x49, 0x2b, 0x17, 0xb6, 0xff, 0x3b, 0x99, 0xb2, 0xaf, 0xf4, 0x27, 0x58,
	0x74, 0xf6, 0x12, 0xb5, 0x6c, 0xf2, 0xd7, 0xb2, 0xc9, 0x6f, 0xd9, 0x57, 0x2a, 0xc2, 0xca, 0x05,
	0x9e, 0x7c, 0x92, 0x60, 0x5d, 0x42, 0xf5, 0xeb, 0x72, 0xff, 0x73, 0x11, 0xd6, 0x0a, 0x21, 0xb8,
	0x05, 0xf2, 0x71, 0x71, 0xa3, 0x46, 0xbe, 0x11, 0xf6, 0x9d, 0xcc, 0xbe, 0x9f, 0xd2, 0xff, 0xd4,
	0x50, 0x5c, 0xfb, 0xa8, 0xa5, 0xef, 0x82, 0x7c, 0x87, 0x79, 0x9e, 0x4b, 0x1d, 0xf3, 0xdc, 0xb5,
	0x45, 0x47, 0x5f, 0x8c, 0xfa, 0x5e, 0x81, 0xcf, 0x43, 0x0c, 0xfe, 0x0f, 0x0c, 0x6b, 0x65, 0xda,
	0xbe, 0xdb, 0x16, 0xfa, 0x52, 0x49, 0x2b, 0xcf, 0xa2, 0x42, 0x02, 0xef, 0x87, 0x28, 0xdc, 0x01,
	0x19, 0x6c, 0x9f, 0x61, 0x6a, 0x11, 0x5b, 0xb7, 0x64, 0x05, 0x97, 0xa7, 0x9a, 0xa2, 0x29, 0x87,
	0x14, 0x4a, 0x84, 0x9b, 0xaf, 0x66, 0x40, 0x26, 0xae, 0x6b, 0x98, 0xd0, 0xc3, 0xc2, 0x15, 0x81,
	0x4d, 0xe4, 0x78, 0xd0, 0x6a, 0xcb, 0xef, 0x6a, 0x4b, 0x10, 0xae, 0x5c, 0x0b, 0x7f, 0xaf, 0x9e,
	0x7d, 0x75, 0x4f, 0x5d, 0x5c, 0xa2, 0x44, 0x08, 0x1f, 0x80, 0x6c, 0x8f, 0x51, 0x27, 0x72, 0xcd,
	0x4c, 0xbb, 0xda, 0xb1, 0xab, 0x7d, 0x89, 0x86, 0x4a, 0xb8, 0x0a, 0x32, 0xb8, 0xa7, 0x9e, 0x95,
	0x92, 0xef, 0x93, 0xdc, 0x4b, 0xce, 0xb2, 0x02, 0x1f, 0x5b, 0x03, 0x79, 0x72, 0x43, 0x4e, 0xdd,
	0xc3, 0x43, 0x30, 0xc7, 0x59, 0xe0, 0x5b, 0x44, 0x1e, 0xd8, 0xc2, 0xf6, 0xfa, 0xfb, 0xba, 0xa4,
	0x29, 0x55, 0xb5, 0x42, 0xbc, 0xaf, 0x6f, 0x2f, 0x56, 0x66, 0x8a, 0xd7, 0x90, 0x72, 0x7f, 0x96,
	0x79, 0x7b, 0xb1, 0x92, 0xce, 0x68, 0x45, 0x6d, 0xf3, 0xb7, 0x79, 0xb0, 0x74, 0xd5, 0x90, 0x82,
	0x77, 0x00, 0xe8, 0x13, 0xce, 0xb1, 0x43, 0xc2, 0xd3, 0xaa, 0xc9, 0xd3, 0x9a, 0x55, 0x88, 0x61,
	0xc3, 0x5f, 0x34, 0x50, 0x6c, 0x33, 0xff, 0x1c, 0xfb, 0x36, 0xf1, 0x4d, 0x4a, 0x44, 0xa8, 0x0a,
	0x0b, 0xb0, 0x50, 0xfb, 0x51, 0x7b, 0x57, 0x9b, 0xfb, 0x3e, 0xdd, 0x49, 0x79, 0xda, 0x9b, 0x8b,
	0x95, 0xd7, 0x1a, 0xf8, 0xd2, 0x61, 0x15, 0xd1, 0x21, 0x42, 0xfe, 0xd5, 0x54, 0x28, 0x11, 0xe7,
	0xcc, 0xef, 0x56, 0xc7, 0x87, 0xf8, 0xd9, 0x4e, 0xd5, 0xeb, 0x3a, 0x55, 0x31, 0xf0, 0x08, 0xaf,
	0x1c, 0x63, 0x9f, 0x77, 0x70, 0xef, 0xd1, 0xc1, 0xd7, 0xb5, 0x81, 0x20, 0x1c, 0x7e, 0x70, 0xc0,
	0x53, 0xda, 0x8f, 0x22, 0x76, 0x64, 0x00, 0x2a, 0x24, 0x4b, 0xad, 0x13, 0x61, 0xd8, 0xb0, 0x02,
	0x16, 0x87, 0x8b, 0x17, 0x84, 0x62, 0x2a, 0xd7, 0x9f, 0x92, 0x6f, 0x79, 0x23, 0xa1, 0x5a, 0x92,
	0x31, 0x6c, 0xf8, 0x7f, 0xb0, 0x34, 0xd4, 0x5b, 0xbd, 0x80, 0x0b, 0xe2, 0x87, 0x86, 0xb4, 0x34,
	0xc0, 0x84, 0xdb, 0x8b, 0x28, 0xc3, 0x86, 0xbf, 0x6a, 0xe0, 0xe6, 0xd0, 0x12, 0xff, 0xf5, 0x90,
	0x20, 0x9a, 0x39, 0xc3, 0x22, 0x65, 0x3e, 0x82, 0x22, 0x3d, 0x8c, 0x8a, 0x34, 0x2c, 0x89, 0xfa,
	0x1b, 0x3c, 0x08, 0x5c, 0x58, 0x1f, 0x7d, 0xf3, 0xe1, 0x3f, 0xa8, 0x9c, 0x85, 0xb9, 0xed, 0xb5,
	0xab, 0xce, 0x98, 0x4b, 0x9d, 0x68, 0xf2, 0xc2, 0xc9, 0x40, 0xc3, 0x86, 0x97, 0x1a, 0x58, 0xec,
	0xb0, 0x3e, 0x31, 0xd5, 0x02, 0xe3, 0xd6, 0x99, 0xfd, 0x58, 0x5b, 0xa7, 0x18, 0xae, 0xb6, 0x1e,
	0xf9, 0xa2, 0xe6, 0xd9, 0x01, 0xb7, 0xc6, 0xde, 0x60, 0xd8, 0x3f, 0x73, 0xb2, 0x1d, 0x16, 0x47,
	0x1c, 0x49, 0x07, 0x3d, 0x00, 0xcb, 0x63, 0xa6, 0x91, 0x26, 0xca, 0x48, 0xd7, 0xd2, 0x88, 0x6b,
	0xd8, 0x46, 0x0f, 0x41, 0xba, 0xc3, 0x3c, 0xae, 0xcf, 0x97, 0x52, 0xff, 0xf4, 0x79, 0x81, 0x58,
	0x20, 0xc8, 0x23, 0xe6, 0x21, 0xe9, 0xd8, 0xfc, 0x43, 0x1b, 0x3f, 0xd8, 0x31, 0x0d, 0x3f, 0x07,
	0x39, 0x9f, 0x58, 0xc4, 0x3d, 0x23, 0xb6, 0x89, 0x85, 0xfa, 0x12, 0xfa, 0xbb, 0x2f, 0x0a, 0x10,
	0xcb, 0x77, 0x05, 0xdc, 0x00, 0x39, 0x4e, 0xa8, 0x3c, 0xf2, 0xb8, 0x1f, 0x4d, 0xbc, 0x2c, 0x02,
	0x11, 0x54, 0xc7, 0x7d, 0x12, 0x7e, 0x5a, 0x28, 0x01, 0xb6, 0x6d, 0x9f, 0x70, 0xae, 0x0e, 0x55,
	0x3e, 0x42, 0x77, 0x23, 0x30, 0x1c, 0xfe, 0x2a, 0x55, 0x25, 0x45, 0x27, 0x69, 0x21, 0x06, 0xe3,
	0xac, 0x44, 0x84, 0x1d, 0x42, 0x85, 0xec, 0x92, 0x2c, 0x4a, 0xac, 0xbb, 0x21, 0x78, 0xff, 0xa7,
	0x19, 0x50, 0x18, 0x9f, 0x7b, 0x10, 0x82, 0x42, 0xb3, 0xf1, 0x14, 0xed, 0x1d, 0x98, 0x4f, 0xeb,
	0x8f, 0xeb, 0x8d, 0xe7, 0xf5, 0xe2, 0x35, 0x58, 0x00, 0x40, 0x61, 0x47, 0x27, 0xcd, 0xa2, 0x06,
	0x17, 0xc1, 0x75, 0x75, 0x8f, 0x0e, 0x8e, 0x8c, 0x66, 0x0b, 0x7d, 0x53, 0x4c, 0xc1, 0x15, 0x70,
	0x53, 0x81, 0xc6, 0x89, 0x79, 0x74, 0xd0, 0x78, 0xd2, 0xd8, 0xdb, 0x6d, 0x19, 0x8d, 0x7a, 0x31,
	0x0d, 0x4b, 0x60, 0x4d, 0x51, 0xcf, 0x8d, 0x43, 0xc3, 0x44, 0xcd, 0xa6, 0x31, 0xa6, 0x98, 0x85,
	0xeb, 0x60, 0x55, 0x29, 0x6a, 0xad, 0x69, 0x7e, 0x6e, 0x24, 0xe1, 0x49, 0x03, 0xed, 0x4e, 0x2b,
	0xe6, 0x27, 0x15, 0xad, 0xfd, 0xc6, 0xee, 0x98, 0x22, 0x03, 0x37, 0xc0, 0x6d, 0xa5, 0xd8, 0x6b,
	0x1c, 0xd7, 0x8c, 0xfa, 0xc1, 0xfe, 0x98, 0x20, 0xbb, 0x5a, 0x78, 0x73, 0xb1, 0x02, 0x74, 0xed,
	0xfe, 0x5c, 0x24, 0xab, 0x3d, 0xb8, 0xfc, 0x7d, 0x5d, 0xfb, 0xb6, 0xfa, 0x01, 0xc7, 0x40, 0x50,
	0xef, 0xf4, 0x74, 0x4e, 0x36, 0xc2, 0xce, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x19, 0xca,
	0xc7, 0x19, 0x0d, 0x00, 0x00,
}
