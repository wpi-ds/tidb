// Code generated by protoc-gen-go.
// source: metapb.proto
// DO NOT EDIT!

/*
Package metapb is a generated protocol buffer package.

It is generated from these files:
	metapb.proto

It has these top-level messages:
	Cluster
	Store
	RegionEpoch
	Peer
	Region
*/
package metapb

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

type Cluster struct {
	Id *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// max peer number for a region.
	// pd will do the auto-balance if region peer number mismatches.
	MaxPeerNumber    *uint32 `protobuf:"varint,2,opt,name=max_peer_number" json:"max_peer_number,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Cluster) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Cluster) GetMaxPeerNumber() uint32 {
	if m != nil && m.MaxPeerNumber != nil {
		return *m.MaxPeerNumber
	}
	return 0
}

type Store struct {
	Id               *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Address          *string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Store) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Store) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

type RegionEpoch struct {
	// Conf change version, auto increment when add or remove peer
	ConfVer *uint64 `protobuf:"varint,1,opt,name=conf_ver" json:"conf_ver,omitempty"`
	// Region version, auto increment when split or merge
	Version          *uint64 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RegionEpoch) Reset()                    { *m = RegionEpoch{} }
func (m *RegionEpoch) String() string            { return proto.CompactTextString(m) }
func (*RegionEpoch) ProtoMessage()               {}
func (*RegionEpoch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegionEpoch) GetConfVer() uint64 {
	if m != nil && m.ConfVer != nil {
		return *m.ConfVer
	}
	return 0
}

func (m *RegionEpoch) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type Peer struct {
	Id               *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	StoreId          *uint64 `protobuf:"varint,2,opt,name=store_id" json:"store_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Peer) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
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
	Id *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Region key range [start_key, end_key).
	StartKey         []byte       `protobuf:"bytes,2,opt,name=start_key" json:"start_key,omitempty"`
	EndKey           []byte       `protobuf:"bytes,3,opt,name=end_key" json:"end_key,omitempty"`
	RegionEpoch      *RegionEpoch `protobuf:"bytes,4,opt,name=region_epoch" json:"region_epoch,omitempty"`
	Peers            []*Peer      `protobuf:"bytes,5,rep,name=peers" json:"peers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Region) Reset()                    { *m = Region{} }
func (m *Region) String() string            { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()               {}
func (*Region) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Region) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
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

func (m *Region) GetRegionEpoch() *RegionEpoch {
	if m != nil {
		return m.RegionEpoch
	}
	return nil
}

func (m *Region) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*Cluster)(nil), "metapb.Cluster")
	proto.RegisterType((*Store)(nil), "metapb.Store")
	proto.RegisterType((*RegionEpoch)(nil), "metapb.RegionEpoch")
	proto.RegisterType((*Peer)(nil), "metapb.Peer")
	proto.RegisterType((*Region)(nil), "metapb.Region")
}

var fileDescriptor0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4e, 0xc3, 0x30,
	0x10, 0x44, 0x95, 0x36, 0x69, 0xe9, 0xd6, 0xa8, 0x60, 0x0e, 0x58, 0xe2, 0x52, 0x45, 0x1c, 0xe0,
	0x12, 0x21, 0x7e, 0x01, 0x71, 0x47, 0xf0, 0x01, 0x96, 0x5b, 0x2f, 0x50, 0x41, 0x6c, 0x6b, 0xed,
	0x22, 0xb8, 0xf1, 0xe9, 0xac, 0x1d, 0x40, 0x84, 0xe3, 0x8e, 0xdf, 0xcc, 0x8e, 0x17, 0x44, 0x8f,
	0xc9, 0x84, 0x4d, 0x17, 0xc8, 0x27, 0x2f, 0x67, 0xc3, 0xd4, 0x76, 0x30, 0xbf, 0x79, 0xdd, 0xc7,
	0x84, 0x24, 0x01, 0x26, 0x3b, 0xab, 0xaa, 0x75, 0x75, 0x51, 0xcb, 0x53, 0x58, 0xf5, 0xe6, 0x5d,
	0x07, 0x44, 0xd2, 0x6e, 0xdf, 0x6f, 0x90, 0xd4, 0x84, 0x1f, 0x0e, 0xdb, 0x73, 0x68, 0x1e, 0x92,
	0x27, 0x1c, 0xd1, 0x2b, 0x98, 0x1b, 0x6b, 0x09, 0x63, 0x2c, 0xd4, 0xa2, 0xbd, 0x82, 0xe5, 0x3d,
	0x3e, 0xed, 0xbc, 0xbb, 0x0d, 0x7e, 0xfb, 0x2c, 0x8f, 0xe0, 0x60, 0xeb, 0xdd, 0xa3, 0x7e, 0xe3,
	0x98, 0x5f, 0x07, 0x0f, 0x91, 0x89, 0xe2, 0xa8, 0x39, 0xb7, 0xbe, 0xc3, 0x7f, 0x25, 0xd8, 0x16,
	0xf3, 0x2e, 0xcd, 0xca, 0x40, 0x7d, 0x56, 0x30, 0x1b, 0x82, 0x47, 0xe0, 0x31, 0x2c, 0x62, 0x32,
	0x94, 0xf4, 0x0b, 0x7e, 0x14, 0x52, 0xe4, 0x05, 0xe8, 0x6c, 0x11, 0xa6, 0x45, 0xb8, 0x04, 0x41,
	0xc5, 0xa9, 0x31, 0x77, 0x52, 0x35, 0xab, 0xcb, 0xeb, 0x93, 0xee, 0xfb, 0x2a, 0x7f, 0xeb, 0x9e,
	0x41, 0x93, 0x3f, 0x1e, 0x55, 0xb3, 0x9e, 0x32, 0x23, 0x7e, 0x98, 0x5c, 0xf0, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0xe0, 0xfe, 0x89, 0x51, 0x47, 0x01, 0x00, 0x00,
}