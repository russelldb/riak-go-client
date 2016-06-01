// Code generated by protoc-gen-go.
// source: riak_ts.proto
// DO NOT EDIT!

/*
Package riak_ts is a generated protocol buffer package.

It is generated from these files:
	riak_ts.proto

It has these top-level messages:
	TsQueryReq
	TsQueryResp
	TsGetReq
	TsGetResp
	TsPutReq
	TsPutResp
	TsDelReq
	TsDelResp
	TsInterpolation
	TsColumnDescription
	TsRow
	TsCell
	TsListKeysReq
	TsListKeysResp
	TsCoverageReq
	TsCoverageResp
	TsCoverageEntry
	TsRange
*/
package riak_ts

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import riak "github.com/basho/riak-go-client/rpb/riak"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TsColumnType int32

const (
	TsColumnType_VARCHAR   TsColumnType = 0
	TsColumnType_SINT64    TsColumnType = 1
	TsColumnType_DOUBLE    TsColumnType = 2
	TsColumnType_TIMESTAMP TsColumnType = 3
	TsColumnType_BOOLEAN   TsColumnType = 4
)

var TsColumnType_name = map[int32]string{
	0: "VARCHAR",
	1: "SINT64",
	2: "DOUBLE",
	3: "TIMESTAMP",
	4: "BOOLEAN",
}
var TsColumnType_value = map[string]int32{
	"VARCHAR":   0,
	"SINT64":    1,
	"DOUBLE":    2,
	"TIMESTAMP": 3,
	"BOOLEAN":   4,
}

func (x TsColumnType) Enum() *TsColumnType {
	p := new(TsColumnType)
	*p = x
	return p
}
func (x TsColumnType) String() string {
	return proto.EnumName(TsColumnType_name, int32(x))
}
func (x *TsColumnType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TsColumnType_value, data, "TsColumnType")
	if err != nil {
		return err
	}
	*x = TsColumnType(value)
	return nil
}
func (TsColumnType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Dispatch a query to Riak
type TsQueryReq struct {
	// left optional to support parameterized queries in the future
	Query            *TsInterpolation `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Stream           *bool            `protobuf:"varint,2,opt,name=stream,def=0" json:"stream,omitempty"`
	CoverContext     []byte           `protobuf:"bytes,3,opt,name=cover_context" json:"cover_context,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *TsQueryReq) Reset()                    { *m = TsQueryReq{} }
func (m *TsQueryReq) String() string            { return proto.CompactTextString(m) }
func (*TsQueryReq) ProtoMessage()               {}
func (*TsQueryReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_TsQueryReq_Stream bool = false

func (m *TsQueryReq) GetQuery() *TsInterpolation {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *TsQueryReq) GetStream() bool {
	if m != nil && m.Stream != nil {
		return *m.Stream
	}
	return Default_TsQueryReq_Stream
}

func (m *TsQueryReq) GetCoverContext() []byte {
	if m != nil {
		return m.CoverContext
	}
	return nil
}

type TsQueryResp struct {
	Columns          []*TsColumnDescription `protobuf:"bytes,1,rep,name=columns" json:"columns,omitempty"`
	Rows             []*TsRow               `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
	Done             *bool                  `protobuf:"varint,3,opt,name=done,def=1" json:"done,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *TsQueryResp) Reset()                    { *m = TsQueryResp{} }
func (m *TsQueryResp) String() string            { return proto.CompactTextString(m) }
func (*TsQueryResp) ProtoMessage()               {}
func (*TsQueryResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_TsQueryResp_Done bool = true

func (m *TsQueryResp) GetColumns() []*TsColumnDescription {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *TsQueryResp) GetRows() []*TsRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *TsQueryResp) GetDone() bool {
	if m != nil && m.Done != nil {
		return *m.Done
	}
	return Default_TsQueryResp_Done
}

type TsGetReq struct {
	Table            []byte    `protobuf:"bytes,1,req,name=table" json:"table,omitempty"`
	Key              []*TsCell `protobuf:"bytes,2,rep,name=key" json:"key,omitempty"`
	Timeout          *uint32   `protobuf:"varint,3,opt,name=timeout" json:"timeout,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *TsGetReq) Reset()                    { *m = TsGetReq{} }
func (m *TsGetReq) String() string            { return proto.CompactTextString(m) }
func (*TsGetReq) ProtoMessage()               {}
func (*TsGetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TsGetReq) GetTable() []byte {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *TsGetReq) GetKey() []*TsCell {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *TsGetReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

type TsGetResp struct {
	Columns          []*TsColumnDescription `protobuf:"bytes,1,rep,name=columns" json:"columns,omitempty"`
	Rows             []*TsRow               `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *TsGetResp) Reset()                    { *m = TsGetResp{} }
func (m *TsGetResp) String() string            { return proto.CompactTextString(m) }
func (*TsGetResp) ProtoMessage()               {}
func (*TsGetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TsGetResp) GetColumns() []*TsColumnDescription {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *TsGetResp) GetRows() []*TsRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

type TsPutReq struct {
	Table []byte `protobuf:"bytes,1,req,name=table" json:"table,omitempty"`
	// optional: omitting it should use table order
	Columns          []*TsColumnDescription `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
	Rows             []*TsRow               `protobuf:"bytes,3,rep,name=rows" json:"rows,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *TsPutReq) Reset()                    { *m = TsPutReq{} }
func (m *TsPutReq) String() string            { return proto.CompactTextString(m) }
func (*TsPutReq) ProtoMessage()               {}
func (*TsPutReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TsPutReq) GetTable() []byte {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *TsPutReq) GetColumns() []*TsColumnDescription {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *TsPutReq) GetRows() []*TsRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

type TsPutResp struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TsPutResp) Reset()                    { *m = TsPutResp{} }
func (m *TsPutResp) String() string            { return proto.CompactTextString(m) }
func (*TsPutResp) ProtoMessage()               {}
func (*TsPutResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type TsDelReq struct {
	Table            []byte    `protobuf:"bytes,1,req,name=table" json:"table,omitempty"`
	Key              []*TsCell `protobuf:"bytes,2,rep,name=key" json:"key,omitempty"`
	Vclock           []byte    `protobuf:"bytes,3,opt,name=vclock" json:"vclock,omitempty"`
	Timeout          *uint32   `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *TsDelReq) Reset()                    { *m = TsDelReq{} }
func (m *TsDelReq) String() string            { return proto.CompactTextString(m) }
func (*TsDelReq) ProtoMessage()               {}
func (*TsDelReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TsDelReq) GetTable() []byte {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *TsDelReq) GetKey() []*TsCell {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *TsDelReq) GetVclock() []byte {
	if m != nil {
		return m.Vclock
	}
	return nil
}

func (m *TsDelReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

type TsDelResp struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TsDelResp) Reset()                    { *m = TsDelResp{} }
func (m *TsDelResp) String() string            { return proto.CompactTextString(m) }
func (*TsDelResp) ProtoMessage()               {}
func (*TsDelResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type TsInterpolation struct {
	Base             []byte          `protobuf:"bytes,1,req,name=base" json:"base,omitempty"`
	Interpolations   []*riak.RpbPair `protobuf:"bytes,2,rep,name=interpolations" json:"interpolations,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *TsInterpolation) Reset()                    { *m = TsInterpolation{} }
func (m *TsInterpolation) String() string            { return proto.CompactTextString(m) }
func (*TsInterpolation) ProtoMessage()               {}
func (*TsInterpolation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TsInterpolation) GetBase() []byte {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *TsInterpolation) GetInterpolations() []*riak.RpbPair {
	if m != nil {
		return m.Interpolations
	}
	return nil
}

type TsColumnDescription struct {
	Name             []byte        `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Type             *TsColumnType `protobuf:"varint,2,req,name=type,enum=TsColumnType" json:"type,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *TsColumnDescription) Reset()                    { *m = TsColumnDescription{} }
func (m *TsColumnDescription) String() string            { return proto.CompactTextString(m) }
func (*TsColumnDescription) ProtoMessage()               {}
func (*TsColumnDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TsColumnDescription) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TsColumnDescription) GetType() TsColumnType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return TsColumnType_VARCHAR
}

type TsRow struct {
	Cells            []*TsCell `protobuf:"bytes,1,rep,name=cells" json:"cells,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *TsRow) Reset()                    { *m = TsRow{} }
func (m *TsRow) String() string            { return proto.CompactTextString(m) }
func (*TsRow) ProtoMessage()               {}
func (*TsRow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TsRow) GetCells() []*TsCell {
	if m != nil {
		return m.Cells
	}
	return nil
}

type TsCell struct {
	VarcharValue     []byte   `protobuf:"bytes,1,opt,name=varchar_value" json:"varchar_value,omitempty"`
	Sint64Value      *int64   `protobuf:"zigzag64,2,opt,name=sint64_value" json:"sint64_value,omitempty"`
	TimestampValue   *int64   `protobuf:"zigzag64,3,opt,name=timestamp_value" json:"timestamp_value,omitempty"`
	BooleanValue     *bool    `protobuf:"varint,4,opt,name=boolean_value" json:"boolean_value,omitempty"`
	DoubleValue      *float64 `protobuf:"fixed64,5,opt,name=double_value" json:"double_value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TsCell) Reset()                    { *m = TsCell{} }
func (m *TsCell) String() string            { return proto.CompactTextString(m) }
func (*TsCell) ProtoMessage()               {}
func (*TsCell) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *TsCell) GetVarcharValue() []byte {
	if m != nil {
		return m.VarcharValue
	}
	return nil
}

func (m *TsCell) GetSint64Value() int64 {
	if m != nil && m.Sint64Value != nil {
		return *m.Sint64Value
	}
	return 0
}

func (m *TsCell) GetTimestampValue() int64 {
	if m != nil && m.TimestampValue != nil {
		return *m.TimestampValue
	}
	return 0
}

func (m *TsCell) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

func (m *TsCell) GetDoubleValue() float64 {
	if m != nil && m.DoubleValue != nil {
		return *m.DoubleValue
	}
	return 0
}

type TsListKeysReq struct {
	Table            []byte  `protobuf:"bytes,1,req,name=table" json:"table,omitempty"`
	Timeout          *uint32 `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TsListKeysReq) Reset()                    { *m = TsListKeysReq{} }
func (m *TsListKeysReq) String() string            { return proto.CompactTextString(m) }
func (*TsListKeysReq) ProtoMessage()               {}
func (*TsListKeysReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *TsListKeysReq) GetTable() []byte {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *TsListKeysReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

type TsListKeysResp struct {
	Keys             []*TsRow `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Done             *bool    `protobuf:"varint,2,opt,name=done" json:"done,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TsListKeysResp) Reset()                    { *m = TsListKeysResp{} }
func (m *TsListKeysResp) String() string            { return proto.CompactTextString(m) }
func (*TsListKeysResp) ProtoMessage()               {}
func (*TsListKeysResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *TsListKeysResp) GetKeys() []*TsRow {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *TsListKeysResp) GetDone() bool {
	if m != nil && m.Done != nil {
		return *m.Done
	}
	return false
}

// Request a segmented coverage plan for this query
type TsCoverageReq struct {
	// left optional to support parameterized queries in the future
	Query            *TsInterpolation `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Table            []byte           `protobuf:"bytes,2,req,name=table" json:"table,omitempty"`
	ReplaceCover     []byte           `protobuf:"bytes,3,opt,name=replace_cover" json:"replace_cover,omitempty"`
	UnavailableCover [][]byte         `protobuf:"bytes,4,rep,name=unavailable_cover" json:"unavailable_cover,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *TsCoverageReq) Reset()                    { *m = TsCoverageReq{} }
func (m *TsCoverageReq) String() string            { return proto.CompactTextString(m) }
func (*TsCoverageReq) ProtoMessage()               {}
func (*TsCoverageReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *TsCoverageReq) GetQuery() *TsInterpolation {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *TsCoverageReq) GetTable() []byte {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *TsCoverageReq) GetReplaceCover() []byte {
	if m != nil {
		return m.ReplaceCover
	}
	return nil
}

func (m *TsCoverageReq) GetUnavailableCover() [][]byte {
	if m != nil {
		return m.UnavailableCover
	}
	return nil
}

// Segmented TS coverage plan response
type TsCoverageResp struct {
	Entries          []*TsCoverageEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *TsCoverageResp) Reset()                    { *m = TsCoverageResp{} }
func (m *TsCoverageResp) String() string            { return proto.CompactTextString(m) }
func (*TsCoverageResp) ProtoMessage()               {}
func (*TsCoverageResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *TsCoverageResp) GetEntries() []*TsCoverageEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Segment of a TS coverage plan
type TsCoverageEntry struct {
	Ip               []byte   `protobuf:"bytes,1,req,name=ip" json:"ip,omitempty"`
	Port             *uint32  `protobuf:"varint,2,req,name=port" json:"port,omitempty"`
	CoverContext     []byte   `protobuf:"bytes,3,req,name=cover_context" json:"cover_context,omitempty"`
	Range            *TsRange `protobuf:"bytes,4,opt,name=range" json:"range,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TsCoverageEntry) Reset()                    { *m = TsCoverageEntry{} }
func (m *TsCoverageEntry) String() string            { return proto.CompactTextString(m) }
func (*TsCoverageEntry) ProtoMessage()               {}
func (*TsCoverageEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *TsCoverageEntry) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *TsCoverageEntry) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *TsCoverageEntry) GetCoverContext() []byte {
	if m != nil {
		return m.CoverContext
	}
	return nil
}

func (m *TsCoverageEntry) GetRange() *TsRange {
	if m != nil {
		return m.Range
	}
	return nil
}

// Each prospective subquery has a range of valid time values
type TsRange struct {
	FieldName           []byte `protobuf:"bytes,1,req,name=field_name" json:"field_name,omitempty"`
	LowerBound          *int64 `protobuf:"zigzag64,2,req,name=lower_bound" json:"lower_bound,omitempty"`
	LowerBoundInclusive *bool  `protobuf:"varint,3,req,name=lower_bound_inclusive" json:"lower_bound_inclusive,omitempty"`
	UpperBound          *int64 `protobuf:"zigzag64,4,req,name=upper_bound" json:"upper_bound,omitempty"`
	UpperBoundInclusive *bool  `protobuf:"varint,5,req,name=upper_bound_inclusive" json:"upper_bound_inclusive,omitempty"`
	Desc                []byte `protobuf:"bytes,6,req,name=desc" json:"desc,omitempty"`
	XXX_unrecognized    []byte `json:"-"`
}

func (m *TsRange) Reset()                    { *m = TsRange{} }
func (m *TsRange) String() string            { return proto.CompactTextString(m) }
func (*TsRange) ProtoMessage()               {}
func (*TsRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *TsRange) GetFieldName() []byte {
	if m != nil {
		return m.FieldName
	}
	return nil
}

func (m *TsRange) GetLowerBound() int64 {
	if m != nil && m.LowerBound != nil {
		return *m.LowerBound
	}
	return 0
}

func (m *TsRange) GetLowerBoundInclusive() bool {
	if m != nil && m.LowerBoundInclusive != nil {
		return *m.LowerBoundInclusive
	}
	return false
}

func (m *TsRange) GetUpperBound() int64 {
	if m != nil && m.UpperBound != nil {
		return *m.UpperBound
	}
	return 0
}

func (m *TsRange) GetUpperBoundInclusive() bool {
	if m != nil && m.UpperBoundInclusive != nil {
		return *m.UpperBoundInclusive
	}
	return false
}

func (m *TsRange) GetDesc() []byte {
	if m != nil {
		return m.Desc
	}
	return nil
}

func init() {
	proto.RegisterType((*TsQueryReq)(nil), "TsQueryReq")
	proto.RegisterType((*TsQueryResp)(nil), "TsQueryResp")
	proto.RegisterType((*TsGetReq)(nil), "TsGetReq")
	proto.RegisterType((*TsGetResp)(nil), "TsGetResp")
	proto.RegisterType((*TsPutReq)(nil), "TsPutReq")
	proto.RegisterType((*TsPutResp)(nil), "TsPutResp")
	proto.RegisterType((*TsDelReq)(nil), "TsDelReq")
	proto.RegisterType((*TsDelResp)(nil), "TsDelResp")
	proto.RegisterType((*TsInterpolation)(nil), "TsInterpolation")
	proto.RegisterType((*TsColumnDescription)(nil), "TsColumnDescription")
	proto.RegisterType((*TsRow)(nil), "TsRow")
	proto.RegisterType((*TsCell)(nil), "TsCell")
	proto.RegisterType((*TsListKeysReq)(nil), "TsListKeysReq")
	proto.RegisterType((*TsListKeysResp)(nil), "TsListKeysResp")
	proto.RegisterType((*TsCoverageReq)(nil), "TsCoverageReq")
	proto.RegisterType((*TsCoverageResp)(nil), "TsCoverageResp")
	proto.RegisterType((*TsCoverageEntry)(nil), "TsCoverageEntry")
	proto.RegisterType((*TsRange)(nil), "TsRange")
	proto.RegisterEnum("TsColumnType", TsColumnType_name, TsColumnType_value)
}

var fileDescriptor0 = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x4f, 0xdb, 0x4c,
	0x10, 0x7e, 0xe3, 0x38, 0x1f, 0x4c, 0x3e, 0x5f, 0x03, 0x25, 0x6d, 0x55, 0x95, 0xba, 0xaa, 0x84,
	0x7a, 0x48, 0x25, 0x8a, 0x38, 0x70, 0x6a, 0x02, 0x51, 0x41, 0x05, 0x92, 0x06, 0x97, 0x4b, 0xa5,
	0x46, 0x1b, 0x67, 0x01, 0x8b, 0x8d, 0xd7, 0xec, 0xae, 0x43, 0xa3, 0xfe, 0x8c, 0xfe, 0xe1, 0xce,
	0x6e, 0xec, 0xe0, 0xa2, 0x1c, 0x90, 0x7a, 0x5b, 0xcf, 0xc7, 0x33, 0xcf, 0xcc, 0x3c, 0x63, 0xa8,
	0x89, 0x80, 0xdc, 0x8e, 0x94, 0x6c, 0x47, 0x82, 0x2b, 0xfe, 0x02, 0xf4, 0xe7, 0xe2, 0xed, 0x7e,
	0x07, 0xf0, 0xe4, 0xd7, 0x98, 0x8a, 0xf9, 0x90, 0xde, 0x39, 0xaf, 0xa1, 0x70, 0xa7, 0xdf, 0xad,
	0xdc, 0x76, 0x6e, 0xa7, 0xb2, 0xdb, 0x6c, 0x7b, 0xf2, 0x24, 0x54, 0x54, 0x44, 0x9c, 0x11, 0x15,
	0xf0, 0xd0, 0xd9, 0x84, 0xa2, 0x54, 0x82, 0x92, 0x69, 0xcb, 0xc2, 0x88, 0xf2, 0x41, 0xe1, 0x8a,
	0x30, 0x49, 0xd1, 0x5c, 0xf3, 0xf9, 0x8c, 0x8a, 0x91, 0xcf, 0x31, 0xfc, 0xa7, 0x6a, 0xe5, 0xd1,
	0x5b, 0x75, 0x7f, 0x40, 0x65, 0x09, 0x2e, 0x23, 0xe7, 0x1d, 0x94, 0x7c, 0xce, 0xe2, 0x69, 0x28,
	0x11, 0x3f, 0x8f, 0xf8, 0x1b, 0x88, 0x7f, 0x68, 0x2c, 0x47, 0x54, 0xfa, 0x22, 0x88, 0x4c, 0x8d,
	0x0d, 0xb0, 0x05, 0xbf, 0x97, 0x58, 0x41, 0xc7, 0x14, 0x31, 0x66, 0xc8, 0xef, 0x1d, 0x07, 0xec,
	0x09, 0x0f, 0xa9, 0x41, 0x2e, 0x1f, 0xd8, 0x4a, 0xc4, 0xd4, 0xfd, 0x04, 0x65, 0x4f, 0x7e, 0xa6,
	0x4a, 0x53, 0xaf, 0x41, 0x41, 0x91, 0x31, 0xa3, 0x08, 0x6d, 0xed, 0x54, 0x11, 0x24, 0x7f, 0x4b,
	0xe7, 0x09, 0x46, 0x49, 0xd7, 0xa1, 0x8c, 0x39, 0x0d, 0x28, 0xa9, 0x60, 0x4a, 0x79, 0xbc, 0x60,
	0x58, 0x73, 0x8f, 0x61, 0x2d, 0x41, 0xf8, 0x47, 0x7e, 0xee, 0xa5, 0xe6, 0x32, 0x88, 0x57, 0x71,
	0xc9, 0xe0, 0x5a, 0x4f, 0xc0, 0xcd, 0xff, 0x85, 0x5b, 0xd1, 0x0c, 0x0d, 0xae, 0x8c, 0xdc, 0x81,
	0x2e, 0x72, 0x44, 0xd9, 0x93, 0x1b, 0xae, 0x43, 0x71, 0xe6, 0x33, 0xee, 0xdf, 0x2e, 0x36, 0x92,
	0x1d, 0x80, 0x6d, 0x06, 0x60, 0xe0, 0x0d, 0x22, 0xc2, 0x77, 0xa0, 0xf1, 0x78, 0xe1, 0x55, 0xb0,
	0xc7, 0x44, 0xa6, 0x45, 0xb6, 0xa1, 0x1e, 0x64, 0xdd, 0x69, 0x43, 0xe5, 0xf6, 0x30, 0x1a, 0x0f,
	0x48, 0x20, 0x70, 0x25, 0xeb, 0xab, 0x7a, 0x43, 0x98, 0x90, 0x4c, 0x53, 0x98, 0x97, 0x60, 0xab,
	0x79, 0x44, 0x31, 0xd9, 0xda, 0xa9, 0xef, 0xd6, 0x96, 0xd3, 0xf0, 0xd0, 0xe8, 0xa2, 0x06, 0x17,
	0x1b, 0x7f, 0x06, 0x05, 0x1f, 0x7b, 0x48, 0x97, 0x91, 0xf6, 0xe4, 0xfe, 0x82, 0x62, 0xd2, 0x1d,
	0xca, 0x6e, 0x46, 0x84, 0x7f, 0x43, 0xc4, 0x68, 0x46, 0x58, 0x4c, 0x8d, 0x6c, 0xf5, 0x28, 0xaa,
	0x12, 0x69, 0xee, 0xef, 0x25, 0x56, 0x2d, 0x55, 0xc7, 0xd9, 0x82, 0x86, 0x6e, 0x5d, 0x2a, 0x32,
	0x8d, 0x12, 0x47, 0xde, 0x38, 0x10, 0x65, 0xcc, 0x39, 0xa3, 0x24, 0x4c, 0xcc, 0x7a, 0x32, 0x65,
	0x8d, 0x32, 0xe1, 0x31, 0x0e, 0x38, 0xb1, 0x16, 0xd0, 0x9a, 0x73, 0x3f, 0x40, 0xcd, 0x93, 0xa7,
	0x81, 0x54, 0x5f, 0xe8, 0x5c, 0xae, 0x58, 0x43, 0x66, 0xc0, 0x96, 0x19, 0xf0, 0x1e, 0xd4, 0xb3,
	0x09, 0x28, 0x33, 0xdc, 0x33, 0x6e, 0x2a, 0x6d, 0x2b, 0xd5, 0x77, 0x35, 0xd1, 0xb7, 0xb9, 0x2b,
	0xf7, 0x4e, 0x97, 0x39, 0xd4, 0x27, 0x45, 0xae, 0xe9, 0x93, 0x2e, 0x73, 0xc9, 0xc3, 0x32, 0x3c,
	0xb0, 0x29, 0x41, 0x23, 0x46, 0x7c, 0x3a, 0x32, 0x97, 0x99, 0xec, 0xff, 0x39, 0xfc, 0x1f, 0x87,
	0x64, 0x46, 0x02, 0xa6, 0x63, 0x13, 0x97, 0x8d, 0x44, 0xaa, 0xee, 0x47, 0x4d, 0xf4, 0xa1, 0x24,
	0x12, 0x7d, 0x03, 0x25, 0x1a, 0x2a, 0x11, 0xd0, 0x94, 0x6b, 0xb3, 0xfd, 0x10, 0xd1, 0x43, 0xcf,
	0x1c, 0x7f, 0x1f, 0x8d, 0x47, 0x26, 0x07, 0xc0, 0x0a, 0xa2, 0x64, 0x1a, 0xd8, 0x54, 0xc4, 0x85,
	0x32, 0x9c, 0x6a, 0xab, 0xfe, 0x12, 0x3a, 0x68, 0x0b, 0x0a, 0x82, 0x84, 0xd7, 0x8b, 0xb9, 0x6b,
	0x2d, 0xe1, 0x40, 0xf4, 0xb7, 0xfb, 0x3b, 0x07, 0xa5, 0xe4, 0x8d, 0xe7, 0x0f, 0x57, 0x01, 0x65,
	0x93, 0x51, 0x46, 0x46, 0xeb, 0x50, 0x61, 0xfc, 0x1e, 0xf1, 0xc6, 0x3c, 0x0e, 0x27, 0xa6, 0x88,
	0xe3, 0xbc, 0x82, 0xcd, 0x8c, 0x71, 0x14, 0x84, 0x3e, 0x8b, 0x65, 0x30, 0xa3, 0xa6, 0x58, 0x59,
	0xe7, 0xc4, 0x51, 0xb4, 0xcc, 0xb1, 0xd3, 0x9c, 0x8c, 0x31, 0x93, 0x53, 0x30, 0x39, 0x7a, 0x35,
	0xa8, 0xe5, 0x56, 0x51, 0x57, 0x7d, 0xdf, 0x87, 0x6a, 0x56, 0xaf, 0x4e, 0x05, 0x4a, 0x97, 0x9d,
	0xe1, 0xe1, 0x71, 0x67, 0xd8, 0xfc, 0x0f, 0x59, 0x16, 0x2f, 0x4e, 0xce, 0xbd, 0xfd, 0xbd, 0x66,
	0x4e, 0xbf, 0x8f, 0xfa, 0xdf, 0xba, 0xa7, 0xbd, 0xa6, 0x85, 0xdb, 0x59, 0xf3, 0x4e, 0xce, 0x7a,
	0x17, 0x5e, 0xe7, 0x6c, 0xd0, 0xcc, 0xeb, 0x9c, 0x6e, 0xbf, 0x7f, 0xda, 0xeb, 0x9c, 0x37, 0xed,
	0xee, 0x5b, 0xd8, 0xf2, 0xf9, 0xb4, 0x8d, 0x67, 0x76, 0xc3, 0xdb, 0x0f, 0xbf, 0xe6, 0x71, 0x7c,
	0xd5, 0x2d, 0x0f, 0xf1, 0x13, 0xcf, 0xbf, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x40, 0x89, 0x00,
	0xc6, 0xc3, 0x05, 0x00, 0x00,
}
