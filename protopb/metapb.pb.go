// Code generated by protoc-gen-go.
// source: metapb.proto
// DO NOT EDIT!

/*
Package protopb is a generated protocol buffer package.

It is generated from these files:
	metapb.proto
	pdpb.proto

It has these top-level messages:
	Node
	Store
	Peer
	Region
	TsoRequest
	Timestamp
	TsoResponse
	BootstrapRequest
	BootstrapResponse
	IsBootstrappedRequest
	IsBootstrappedResponse
	AllocIdRequest
	AllocIdResponse
	GetMetaRequest
	GetMetaResponse
	PutMetaRequest
	PutMetaResponse
	DeleteMetaRequest
	DeleteMetaResponse
	RequestHeader
	ResponseHeader
	Request
	Response
	BootstrappedError
	Error
*/
package protopb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Node struct {
	NodeId           *uint64 `protobuf:"varint,1,opt,name=node_id" json:"node_id,omitempty"`
	Address          *string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Node) GetNodeId() uint64 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

func (m *Node) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

type Store struct {
	StoreId          *uint64 `protobuf:"varint,1,opt,name=store_id" json:"store_id,omitempty"`
	NodeId           *uint64 `protobuf:"varint,2,opt,name=node_id" json:"node_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Store) GetStoreId() uint64 {
	if m != nil && m.StoreId != nil {
		return *m.StoreId
	}
	return 0
}

func (m *Store) GetNodeId() uint64 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

type Peer struct {
	PeerId           *uint64 `protobuf:"varint,1,opt,name=peer_id" json:"peer_id,omitempty"`
	NodeId           *uint64 `protobuf:"varint,2,opt,name=node_id" json:"node_id,omitempty"`
	StoreId          *uint64 `protobuf:"varint,3,opt,name=store_id" json:"store_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Peer) GetPeerId() uint64 {
	if m != nil && m.PeerId != nil {
		return *m.PeerId
	}
	return 0
}

func (m *Peer) GetNodeId() uint64 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

func (m *Peer) GetStoreId() uint64 {
	if m != nil && m.StoreId != nil {
		return *m.StoreId
	}
	return 0
}

type Region struct {
	RegionId *uint64 `protobuf:"varint,1,opt,name=region_id" json:"region_id,omitempty"`
	// Region key range [start_key, end_key).
	StartKey []byte  `protobuf:"bytes,2,opt,name=start_key" json:"start_key,omitempty"`
	EndKey   []byte  `protobuf:"bytes,3,opt,name=end_key" json:"end_key,omitempty"`
	Peers    []*Peer `protobuf:"bytes,4,rep,name=peers" json:"peers,omitempty"`
	// The peer id is generated by outer placement driver incrementally,
	// and we can guarantee that the next generated one is larger than
	// the max_peer_id.
	// If we receive a message which the peer is not in region peers
	// and the peer id is <= max_peer_id, we can be sure that the
	// peer which sends this message is a tombstone peer and can
	// skip the message directly.
	MaxPeerId        *uint64 `protobuf:"varint,5,opt,name=max_peer_id" json:"max_peer_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Region) Reset()                    { *m = Region{} }
func (m *Region) String() string            { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()               {}
func (*Region) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Region) GetRegionId() uint64 {
	if m != nil && m.RegionId != nil {
		return *m.RegionId
	}
	return 0
}

func (m *Region) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *Region) GetEndKey() []byte {
	if m != nil {
		return m.EndKey
	}
	return nil
}

func (m *Region) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *Region) GetMaxPeerId() uint64 {
	if m != nil && m.MaxPeerId != nil {
		return *m.MaxPeerId
	}
	return 0
}

func init() {
	proto.RegisterType((*Node)(nil), "metapb.Node")
	proto.RegisterType((*Store)(nil), "metapb.Store")
	proto.RegisterType((*Peer)(nil), "metapb.Peer")
	proto.RegisterType((*Region)(nil), "metapb.Region")
}

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4d, 0x2d, 0x49,
	0x2c, 0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x34, 0xb8, 0x58,
	0xfc, 0xf2, 0x53, 0x52, 0x85, 0xf8, 0xb9, 0xd8, 0xf3, 0x80, 0x74, 0x7c, 0x66, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x0b, 0x48, 0x20, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x09, 0x28,
	0xc0, 0xa9, 0xa4, 0xc5, 0xc5, 0x1a, 0x5c, 0x92, 0x5f, 0x94, 0x2a, 0x24, 0xc0, 0xc5, 0x51, 0x0c,
	0x62, 0xa0, 0xa8, 0x85, 0x69, 0x06, 0xa9, 0x65, 0x51, 0xb2, 0xe2, 0x62, 0x09, 0x48, 0x4d, 0x2d,
	0x02, 0x49, 0x14, 0x00, 0x69, 0xdc, 0x2a, 0x51, 0x0c, 0x63, 0x06, 0xeb, 0x2d, 0xe4, 0x62, 0x0b,
	0x4a, 0x4d, 0xcf, 0xcc, 0xcf, 0x13, 0x12, 0xe4, 0xe2, 0x2c, 0x02, 0xb3, 0x10, 0xfa, 0x81, 0x42,
	0xc5, 0x25, 0x89, 0x45, 0x25, 0xf1, 0xd9, 0xa9, 0x95, 0x60, 0x13, 0x78, 0x40, 0x46, 0xa6, 0xe6,
	0xa5, 0x80, 0x05, 0x98, 0xc1, 0x02, 0xd2, 0x5c, 0xac, 0x20, 0x4b, 0x8b, 0x25, 0x58, 0x14, 0x98,
	0x35, 0xb8, 0x8d, 0x78, 0xf4, 0xa0, 0x1e, 0x07, 0xbb, 0x48, 0x98, 0x8b, 0x3b, 0x37, 0xb1, 0x22,
	0x1e, 0xe6, 0x2a, 0x56, 0x90, 0xa9, 0x4e, 0x9c, 0x51, 0xec, 0xe0, 0x50, 0x29, 0x48, 0x02, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xa2, 0xc6, 0xa8, 0x90, 0x26, 0x01, 0x00, 0x00,
}
