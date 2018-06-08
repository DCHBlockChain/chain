/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ledger/rwset/kvrwset/kv_rwset.proto

/*
Package kvrwset is a generated protocol buffer package.

It is generated from these files:
	ledger/rwset/kvrwset/kv_rwset.proto

It has these top-level messages:
	KVRWSet
	HashedRWSet
	KVRead
	KVWrite
	KVMetadataWrite
	KVReadHash
	KVWriteHash
	KVMetadataWriteHash
	KVMetadataEntry
	Version
	RangeQueryInfo
	QueryReads
	QueryReadsMerkleSummary
*/
package kvrwset

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

// KVRWSet encapsulates the read-write set for a chaincode that operates upon a KV or Document data model
// This structure is used for both the public data and the private data
type KVRWSet struct {
	Reads            []*KVRead          `protobuf:"bytes,1,rep,name=reads" json:"reads,omitempty"`
	RangeQueriesInfo []*RangeQueryInfo  `protobuf:"bytes,2,rep,name=range_queries_info,json=rangeQueriesInfo" json:"range_queries_info,omitempty"`
	Writes           []*KVWrite         `protobuf:"bytes,3,rep,name=writes" json:"writes,omitempty"`
	MetadataWrites   []*KVMetadataWrite `protobuf:"bytes,4,rep,name=metadata_writes,json=metadataWrites" json:"metadata_writes,omitempty"`
}

func (m *KVRWSet) Reset()                    { *m = KVRWSet{} }
func (m *KVRWSet) String() string            { return proto.CompactTextString(m) }
func (*KVRWSet) ProtoMessage()               {}
func (*KVRWSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KVRWSet) GetReads() []*KVRead {
	if m != nil {
		return m.Reads
	}
	return nil
}

func (m *KVRWSet) GetRangeQueriesInfo() []*RangeQueryInfo {
	if m != nil {
		return m.RangeQueriesInfo
	}
	return nil
}

func (m *KVRWSet) GetWrites() []*KVWrite {
	if m != nil {
		return m.Writes
	}
	return nil
}

func (m *KVRWSet) GetMetadataWrites() []*KVMetadataWrite {
	if m != nil {
		return m.MetadataWrites
	}
	return nil
}

// HashedRWSet encapsulates hashed representation of a private read-write set for KV or Document data model
type HashedRWSet struct {
	HashedReads    []*KVReadHash          `protobuf:"bytes,1,rep,name=hashed_reads,json=hashedReads" json:"hashed_reads,omitempty"`
	HashedWrites   []*KVWriteHash         `protobuf:"bytes,2,rep,name=hashed_writes,json=hashedWrites" json:"hashed_writes,omitempty"`
	MetadataWrites []*KVMetadataWriteHash `protobuf:"bytes,3,rep,name=metadata_writes,json=metadataWrites" json:"metadata_writes,omitempty"`
}

func (m *HashedRWSet) Reset()                    { *m = HashedRWSet{} }
func (m *HashedRWSet) String() string            { return proto.CompactTextString(m) }
func (*HashedRWSet) ProtoMessage()               {}
func (*HashedRWSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HashedRWSet) GetHashedReads() []*KVReadHash {
	if m != nil {
		return m.HashedReads
	}
	return nil
}

func (m *HashedRWSet) GetHashedWrites() []*KVWriteHash {
	if m != nil {
		return m.HashedWrites
	}
	return nil
}

func (m *HashedRWSet) GetMetadataWrites() []*KVMetadataWriteHash {
	if m != nil {
		return m.MetadataWrites
	}
	return nil
}

// KVRead captures a read operation performed during transaction simulation
// A 'nil' version indicates a non-existing key read by the transaction
type KVRead struct {
	Key     string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Version *Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *KVRead) Reset()                    { *m = KVRead{} }
func (m *KVRead) String() string            { return proto.CompactTextString(m) }
func (*KVRead) ProtoMessage()               {}
func (*KVRead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KVRead) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVRead) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

// KVWrite captures a write (update/delete) operation performed during transaction simulation
type KVWrite struct {
	Key      string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	IsDelete bool   `protobuf:"varint,2,opt,name=is_delete,json=isDelete" json:"is_delete,omitempty"`
	Value    []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KVWrite) Reset()                    { *m = KVWrite{} }
func (m *KVWrite) String() string            { return proto.CompactTextString(m) }
func (*KVWrite) ProtoMessage()               {}
func (*KVWrite) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *KVWrite) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVWrite) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

func (m *KVWrite) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// KVMetadataWrite captures all the entries in the metadata associated with a key
type KVMetadataWrite struct {
	Key     string             `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Entries []*KVMetadataEntry `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
}

func (m *KVMetadataWrite) Reset()                    { *m = KVMetadataWrite{} }
func (m *KVMetadataWrite) String() string            { return proto.CompactTextString(m) }
func (*KVMetadataWrite) ProtoMessage()               {}
func (*KVMetadataWrite) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *KVMetadataWrite) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVMetadataWrite) GetEntries() []*KVMetadataEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// KVReadHash is similar to the KVRead in spirit. However, it captures the hash of the key instead of the key itself
// version is kept as is for now. However, if the version also needs to be privacy-protected, it would need to be the
// hash of the version and hence of 'bytes' type
type KVReadHash struct {
	KeyHash []byte   `protobuf:"bytes,1,opt,name=key_hash,json=keyHash,proto3" json:"key_hash,omitempty"`
	Version *Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *KVReadHash) Reset()                    { *m = KVReadHash{} }
func (m *KVReadHash) String() string            { return proto.CompactTextString(m) }
func (*KVReadHash) ProtoMessage()               {}
func (*KVReadHash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *KVReadHash) GetKeyHash() []byte {
	if m != nil {
		return m.KeyHash
	}
	return nil
}

func (m *KVReadHash) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

// KVWriteHash is similar to the KVWrite. It captures a write (update/delete) operation performed during transaction simulation
type KVWriteHash struct {
	KeyHash   []byte `protobuf:"bytes,1,opt,name=key_hash,json=keyHash,proto3" json:"key_hash,omitempty"`
	IsDelete  bool   `protobuf:"varint,2,opt,name=is_delete,json=isDelete" json:"is_delete,omitempty"`
	ValueHash []byte `protobuf:"bytes,3,opt,name=value_hash,json=valueHash,proto3" json:"value_hash,omitempty"`
}

func (m *KVWriteHash) Reset()                    { *m = KVWriteHash{} }
func (m *KVWriteHash) String() string            { return proto.CompactTextString(m) }
func (*KVWriteHash) ProtoMessage()               {}
func (*KVWriteHash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *KVWriteHash) GetKeyHash() []byte {
	if m != nil {
		return m.KeyHash
	}
	return nil
}

func (m *KVWriteHash) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

func (m *KVWriteHash) GetValueHash() []byte {
	if m != nil {
		return m.ValueHash
	}
	return nil
}

// KVMetadataWriteHash captures all the upserts to the metadata associated with a key hash
type KVMetadataWriteHash struct {
	KeyHash []byte             `protobuf:"bytes,1,opt,name=key_hash,json=keyHash,proto3" json:"key_hash,omitempty"`
	Entries []*KVMetadataEntry `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
}

func (m *KVMetadataWriteHash) Reset()                    { *m = KVMetadataWriteHash{} }
func (m *KVMetadataWriteHash) String() string            { return proto.CompactTextString(m) }
func (*KVMetadataWriteHash) ProtoMessage()               {}
func (*KVMetadataWriteHash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *KVMetadataWriteHash) GetKeyHash() []byte {
	if m != nil {
		return m.KeyHash
	}
	return nil
}

func (m *KVMetadataWriteHash) GetEntries() []*KVMetadataEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// KVMetadataEntry captures a 'name'ed entry in the metadata of a key/key-hash.
type KVMetadataEntry struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KVMetadataEntry) Reset()                    { *m = KVMetadataEntry{} }
func (m *KVMetadataEntry) String() string            { return proto.CompactTextString(m) }
func (*KVMetadataEntry) ProtoMessage()               {}
func (*KVMetadataEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *KVMetadataEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KVMetadataEntry) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Version encapsulates the version of a Key
// A version of a committed key is maintained as the height of the transaction that committed the key.
// The height is represenetd as a tuple <blockNum, txNum> where the txNum is the position of the transaction
// (starting with 0) within block
type Version struct {
	BlockNum uint64 `protobuf:"varint,1,opt,name=block_num,json=blockNum" json:"block_num,omitempty"`
	TxNum    uint64 `protobuf:"varint,2,opt,name=tx_num,json=txNum" json:"tx_num,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Version) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *Version) GetTxNum() uint64 {
	if m != nil {
		return m.TxNum
	}
	return 0
}

// RangeQueryInfo encapsulates the details of a range query performed by a transaction during simulation.
// This helps protect transactions from phantom reads by varifying during validation whether any new items
// got committed within the given range between transaction simuation and validation
// (in addition to regular checks for updates/deletes of the existing items).
// readInfo field contains either the KVReads (for the items read by the range query) or a merkle-tree hash
// if the KVReads exceeds a pre-configured numbers
type RangeQueryInfo struct {
	StartKey     string `protobuf:"bytes,1,opt,name=start_key,json=startKey" json:"start_key,omitempty"`
	EndKey       string `protobuf:"bytes,2,opt,name=end_key,json=endKey" json:"end_key,omitempty"`
	ItrExhausted bool   `protobuf:"varint,3,opt,name=itr_exhausted,json=itrExhausted" json:"itr_exhausted,omitempty"`
	// Types that are valid to be assigned to ReadsInfo:
	//	*RangeQueryInfo_RawReads
	//	*RangeQueryInfo_ReadsMerkleHashes
	ReadsInfo isRangeQueryInfo_ReadsInfo `protobuf_oneof:"reads_info"`
}

func (m *RangeQueryInfo) Reset()                    { *m = RangeQueryInfo{} }
func (m *RangeQueryInfo) String() string            { return proto.CompactTextString(m) }
func (*RangeQueryInfo) ProtoMessage()               {}
func (*RangeQueryInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type isRangeQueryInfo_ReadsInfo interface{ isRangeQueryInfo_ReadsInfo() }

type RangeQueryInfo_RawReads struct {
	RawReads *QueryReads `protobuf:"bytes,4,opt,name=raw_reads,json=rawReads,oneof"`
}
type RangeQueryInfo_ReadsMerkleHashes struct {
	ReadsMerkleHashes *QueryReadsMerkleSummary `protobuf:"bytes,5,opt,name=reads_merkle_hashes,json=readsMerkleHashes,oneof"`
}

func (*RangeQueryInfo_RawReads) isRangeQueryInfo_ReadsInfo()          {}
func (*RangeQueryInfo_ReadsMerkleHashes) isRangeQueryInfo_ReadsInfo() {}

func (m *RangeQueryInfo) GetReadsInfo() isRangeQueryInfo_ReadsInfo {
	if m != nil {
		return m.ReadsInfo
	}
	return nil
}

func (m *RangeQueryInfo) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *RangeQueryInfo) GetEndKey() string {
	if m != nil {
		return m.EndKey
	}
	return ""
}

func (m *RangeQueryInfo) GetItrExhausted() bool {
	if m != nil {
		return m.ItrExhausted
	}
	return false
}

func (m *RangeQueryInfo) GetRawReads() *QueryReads {
	if x, ok := m.GetReadsInfo().(*RangeQueryInfo_RawReads); ok {
		return x.RawReads
	}
	return nil
}

func (m *RangeQueryInfo) GetReadsMerkleHashes() *QueryReadsMerkleSummary {
	if x, ok := m.GetReadsInfo().(*RangeQueryInfo_ReadsMerkleHashes); ok {
		return x.ReadsMerkleHashes
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RangeQueryInfo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RangeQueryInfo_OneofMarshaler, _RangeQueryInfo_OneofUnmarshaler, _RangeQueryInfo_OneofSizer, []interface{}{
		(*RangeQueryInfo_RawReads)(nil),
		(*RangeQueryInfo_ReadsMerkleHashes)(nil),
	}
}

func _RangeQueryInfo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RangeQueryInfo)
	// reads_info
	switch x := m.ReadsInfo.(type) {
	case *RangeQueryInfo_RawReads:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RawReads); err != nil {
			return err
		}
	case *RangeQueryInfo_ReadsMerkleHashes:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReadsMerkleHashes); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RangeQueryInfo.ReadsInfo has unexpected type %T", x)
	}
	return nil
}

func _RangeQueryInfo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RangeQueryInfo)
	switch tag {
	case 4: // reads_info.raw_reads
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(QueryReads)
		err := b.DecodeMessage(msg)
		m.ReadsInfo = &RangeQueryInfo_RawReads{msg}
		return true, err
	case 5: // reads_info.reads_merkle_hashes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(QueryReadsMerkleSummary)
		err := b.DecodeMessage(msg)
		m.ReadsInfo = &RangeQueryInfo_ReadsMerkleHashes{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RangeQueryInfo_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RangeQueryInfo)
	// reads_info
	switch x := m.ReadsInfo.(type) {
	case *RangeQueryInfo_RawReads:
		s := proto.Size(x.RawReads)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RangeQueryInfo_ReadsMerkleHashes:
		s := proto.Size(x.ReadsMerkleHashes)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// QueryReads encapsulates the KVReads for the items read by a transaction as a result of a query execution
type QueryReads struct {
	KvReads []*KVRead `protobuf:"bytes,1,rep,name=kv_reads,json=kvReads" json:"kv_reads,omitempty"`
}

func (m *QueryReads) Reset()                    { *m = QueryReads{} }
func (m *QueryReads) String() string            { return proto.CompactTextString(m) }
func (*QueryReads) ProtoMessage()               {}
func (*QueryReads) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *QueryReads) GetKvReads() []*KVRead {
	if m != nil {
		return m.KvReads
	}
	return nil
}

// QueryReadsMerkleSummary encapsulates the Merkle-tree hashes for the QueryReads
// This allows to reduce the size of RWSet in the presence of query results
// by storing certain hashes instead of actual results.
// maxDegree field refers to the maximum number of children in the tree at any level
// maxLevel field contains the lowest level which has lesser nodes than maxDegree (starting from leaf level)
type QueryReadsMerkleSummary struct {
	MaxDegree      uint32   `protobuf:"varint,1,opt,name=max_degree,json=maxDegree" json:"max_degree,omitempty"`
	MaxLevel       uint32   `protobuf:"varint,2,opt,name=max_level,json=maxLevel" json:"max_level,omitempty"`
	MaxLevelHashes [][]byte `protobuf:"bytes,3,rep,name=max_level_hashes,json=maxLevelHashes,proto3" json:"max_level_hashes,omitempty"`
}

func (m *QueryReadsMerkleSummary) Reset()                    { *m = QueryReadsMerkleSummary{} }
func (m *QueryReadsMerkleSummary) String() string            { return proto.CompactTextString(m) }
func (*QueryReadsMerkleSummary) ProtoMessage()               {}
func (*QueryReadsMerkleSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *QueryReadsMerkleSummary) GetMaxDegree() uint32 {
	if m != nil {
		return m.MaxDegree
	}
	return 0
}

func (m *QueryReadsMerkleSummary) GetMaxLevel() uint32 {
	if m != nil {
		return m.MaxLevel
	}
	return 0
}

func (m *QueryReadsMerkleSummary) GetMaxLevelHashes() [][]byte {
	if m != nil {
		return m.MaxLevelHashes
	}
	return nil
}

func init() {
	proto.RegisterType((*KVRWSet)(nil), "sdk.kvrwset.KVRWSet")
	proto.RegisterType((*HashedRWSet)(nil), "sdk.kvrwset.HashedRWSet")
	proto.RegisterType((*KVRead)(nil), "sdk.kvrwset.KVRead")
	proto.RegisterType((*KVWrite)(nil), "sdk.kvrwset.KVWrite")
	proto.RegisterType((*KVMetadataWrite)(nil), "sdk.kvrwset.KVMetadataWrite")
	proto.RegisterType((*KVReadHash)(nil), "sdk.kvrwset.KVReadHash")
	proto.RegisterType((*KVWriteHash)(nil), "sdk.kvrwset.KVWriteHash")
	proto.RegisterType((*KVMetadataWriteHash)(nil), "sdk.kvrwset.KVMetadataWriteHash")
	proto.RegisterType((*KVMetadataEntry)(nil), "sdk.kvrwset.KVMetadataEntry")
	proto.RegisterType((*Version)(nil), "sdk.kvrwset.Version")
	proto.RegisterType((*RangeQueryInfo)(nil), "sdk.kvrwset.RangeQueryInfo")
	proto.RegisterType((*QueryReads)(nil), "sdk.kvrwset.QueryReads")
	proto.RegisterType((*QueryReadsMerkleSummary)(nil), "sdk.kvrwset.QueryReadsMerkleSummary")
}

func init() { proto.RegisterFile("ledger/rwset/kvrwset/kv_rwset.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdf, 0x6b, 0xdb, 0x48,
	0x10, 0x8e, 0x7f, 0xcb, 0x63, 0x3b, 0xf1, 0x6d, 0x72, 0x44, 0xc7, 0xdd, 0x81, 0x51, 0x38, 0x30,
	0x79, 0xb0, 0xc1, 0x07, 0xc7, 0x85, 0xe3, 0x1e, 0x5a, 0xe2, 0x92, 0x92, 0x26, 0xd0, 0x0d, 0x24,
	0xd0, 0x17, 0xb1, 0x8e, 0x26, 0xb6, 0xb0, 0x25, 0xa5, 0xab, 0x95, 0x6d, 0x3d, 0x95, 0xfe, 0x75,
	0xfd, 0x47, 0xfa, 0x87, 0x94, 0x9d, 0x95, 0x63, 0xc5, 0x75, 0x0c, 0xed, 0x93, 0xb5, 0xf3, 0xcd,
	0x37, 0x3b, 0xdf, 0x37, 0xde, 0x5d, 0x38, 0x99, 0xa1, 0x37, 0x46, 0xd9, 0x97, 0x8b, 0x18, 0x55,
	0x7f, 0x3a, 0x5f, 0xfd, 0xba, 0xf4, 0xd1, 0x7b, 0x94, 0x91, 0x8a, 0x58, 0x2d, 0x8b, 0x3b, 0x5f,
	0x0b, 0x50, 0xbb, 0xbc, 0xe5, 0x77, 0x37, 0xa8, 0xd8, 0x5f, 0x50, 0x91, 0x28, 0xbc, 0xd8, 0x2e,
	0x74, 0x4a, 0xdd, 0xc6, 0xe0, 0xa0, 0x97, 0x25, 0xf5, 0x2e, 0x6f, 0x39, 0x0a, 0x8f, 0x1b, 0x94,
	0x0d, 0x81, 0x49, 0x11, 0x8e, 0xd1, 0xfd, 0x98, 0xa0, 0xf4, 0x31, 0x76, 0xfd, 0xf0, 0x21, 0xb2,
	0x8b, 0xc4, 0x39, 0x7e, 0xe2, 0x70, 0x9d, 0xf2, 0x3e, 0x41, 0x99, 0xbe, 0x0d, 0x1f, 0x22, 0xde,
	0x96, 0xab, 0xb5, 0x8f, 0xb1, 0x8e, 0xb0, 0x2e, 0x54, 0x17, 0xd2, 0x57, 0x18, 0xdb, 0x25, 0xa2,
	0xb6, 0x73, 0xdb, 0xdd, 0x69, 0x80, 0x67, 0x38, 0x7b, 0x05, 0x07, 0x01, 0x2a, 0xe1, 0x09, 0x25,
	0xdc, 0x8c, 0x52, 0x26, 0x8a, 0x9d, 0xa3, 0x5c, 0x65, 0x19, 0x86, 0xba, 0x1f, 0xe4, 0x97, 0xb1,
	0xf3, 0xa5, 0x00, 0x8d, 0x0b, 0x11, 0x4f, 0xd0, 0x33, 0x52, 0xff, 0x81, 0xe6, 0x84, 0x96, 0x6e,
	0x5e, 0xf1, 0xe1, 0x86, 0x62, 0xcd, 0xe0, 0x0d, 0x93, 0xc8, 0x49, 0xfb, 0x19, 0xb4, 0x32, 0x5e,
	0xd6, 0x88, 0x91, 0x7d, 0xb4, 0xd9, 0x3b, 0x31, 0xb3, 0x2d, 0x4c, 0x0b, 0x6c, 0xf8, 0xbd, 0x0a,
	0x23, 0xfc, 0x8f, 0x97, 0x54, 0x50, 0x91, 0x4d, 0x25, 0x6f, 0xa0, 0x6a, 0x9a, 0x63, 0x6d, 0x28,
	0x4d, 0x31, 0xb5, 0x0b, 0x9d, 0x42, 0xb7, 0xce, 0xf5, 0x27, 0x3b, 0x85, 0xda, 0x1c, 0x65, 0xec,
	0x47, 0xa1, 0x5d, 0xec, 0x14, 0x9e, 0x79, 0x7a, 0x6b, 0xe2, 0x7c, 0x95, 0xe0, 0x5c, 0xeb, 0xb9,
	0x53, 0xcd, 0x2d, 0x85, 0x7e, 0x87, 0xba, 0x1f, 0xbb, 0x1e, 0xce, 0x50, 0x21, 0x95, 0xb2, 0xb8,
	0xe5, 0xc7, 0xe7, 0xb4, 0x66, 0x47, 0x50, 0x99, 0x8b, 0x59, 0x82, 0x76, 0xa9, 0x53, 0xe8, 0x36,
	0xb9, 0x59, 0x38, 0x77, 0x70, 0xb0, 0xd1, 0xfe, 0x96, 0xba, 0x03, 0xa8, 0x61, 0xa8, 0xf4, 0x5f,
	0x20, 0x33, 0x6e, 0xdb, 0x04, 0x87, 0xa1, 0x92, 0x29, 0x5f, 0x25, 0x3a, 0x37, 0x00, 0xeb, 0x69,
	0xb0, 0xdf, 0xc0, 0x9a, 0x62, 0xea, 0x6a, 0x67, 0xa9, 0x70, 0x93, 0xd7, 0xa6, 0x98, 0x12, 0xf4,
	0x23, 0xea, 0x3d, 0x68, 0xe4, 0x26, 0xb5, 0xab, 0xea, 0x4e, 0x2b, 0xfe, 0x04, 0x20, 0xf5, 0x86,
	0x69, 0xfc, 0xa8, 0x53, 0x44, 0x73, 0x1d, 0x0f, 0x0e, 0xb7, 0x8c, 0x74, 0xd7, 0x6e, 0x3f, 0x63,
	0xd0, 0x7f, 0x79, 0xe7, 0x09, 0x63, 0x0c, 0xca, 0xa1, 0x08, 0x30, 0xb3, 0x9e, 0xbe, 0xd7, 0x63,
	0x2b, 0xe6, 0xc7, 0xf6, 0x3f, 0xd4, 0x32, 0x73, 0xb4, 0xd2, 0xd1, 0x2c, 0xba, 0x9f, 0xba, 0x61,
	0x12, 0x10, 0xb3, 0xcc, 0x2d, 0x0a, 0x5c, 0x27, 0x01, 0xfb, 0x15, 0xaa, 0x6a, 0x49, 0x48, 0x91,
	0x90, 0x8a, 0x5a, 0x5e, 0x27, 0x81, 0xf3, 0xb9, 0x08, 0xfb, 0xcf, 0x4f, 0xba, 0x2e, 0x13, 0x2b,
	0x21, 0x95, 0xbb, 0x9e, 0xbd, 0x45, 0x81, 0x4b, 0x4c, 0xd9, 0xb1, 0xd6, 0xe7, 0x11, 0x54, 0x24,
	0xa8, 0x8a, 0xa1, 0xa7, 0x81, 0x13, 0x68, 0xf9, 0x4a, 0xba, 0xb8, 0x9c, 0x88, 0x24, 0x56, 0xe8,
	0x91, 0x99, 0x16, 0x6f, 0xfa, 0x4a, 0x0e, 0x57, 0x31, 0x36, 0x80, 0xba, 0x14, 0x8b, 0xec, 0xc8,
	0x96, 0x69, 0xc6, 0xeb, 0x23, 0x4b, 0x1d, 0xd0, 0x29, 0xbd, 0xd8, 0xe3, 0x96, 0x14, 0x0b, 0x73,
	0x62, 0x39, 0x1c, 0x52, 0xbe, 0x1b, 0xa0, 0x9c, 0xce, 0xcc, 0xa4, 0x30, 0xb6, 0x2b, 0xc4, 0xee,
	0x6c, 0x61, 0x5f, 0x51, 0xde, 0x4d, 0x12, 0x04, 0x42, 0xa6, 0x17, 0x7b, 0xfc, 0x17, 0xb9, 0x8e,
	0xd2, 0x15, 0x12, 0xbf, 0x6e, 0x02, 0x98, 0x9a, 0xfa, 0xe6, 0x73, 0xfe, 0x05, 0x58, 0xb3, 0xd9,
	0x29, 0x58, 0xfa, 0xae, 0xdd, 0x75, 0x8f, 0xd6, 0xa6, 0x73, 0xca, 0x75, 0x3e, 0xc1, 0xf1, 0x0b,
	0xfb, 0xea, 0x7f, 0x56, 0x20, 0x96, 0xae, 0x87, 0x63, 0x89, 0x66, 0x8e, 0x2d, 0x5e, 0x0f, 0xc4,
	0xf2, 0x9c, 0x02, 0xda, 0x64, 0x0d, 0xcf, 0x70, 0x8e, 0x33, 0x72, 0xb2, 0xc5, 0xad, 0x40, 0x2c,
	0xdf, 0xe9, 0x35, 0xeb, 0x42, 0xfb, 0x09, 0x5c, 0xe9, 0xd5, 0x57, 0x4d, 0x93, 0xef, 0xaf, 0x72,
	0x32, 0x21, 0x11, 0x0c, 0x22, 0x39, 0xee, 0x4d, 0xd2, 0x47, 0x94, 0xe6, 0xd9, 0xe8, 0x3d, 0x88,
	0x91, 0xf4, 0xef, 0xcd, 0x33, 0x11, 0xf7, 0xb2, 0xa0, 0x69, 0x3f, 0x93, 0xf1, 0xe1, 0x6c, 0xec,
	0xab, 0x49, 0x32, 0xea, 0xdd, 0x47, 0x41, 0x3f, 0x47, 0xed, 0x1b, 0x6a, 0xdf, 0x50, 0xfb, 0xdb,
	0x9e, 0xa1, 0x51, 0x95, 0xc0, 0xbf, 0xbf, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xe1, 0xb5, 0x07,
	0xa5, 0x06, 0x00, 0x00,
}
