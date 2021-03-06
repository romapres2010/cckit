// Code generated by protoc-gen-go. DO NOT EDIT.
// source: state.proto

package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommercialPaper_State int32

const (
	CommercialPaper_ISSUED   CommercialPaper_State = 0
	CommercialPaper_TRADING  CommercialPaper_State = 1
	CommercialPaper_REDEEMED CommercialPaper_State = 2
)

var CommercialPaper_State_name = map[int32]string{
	0: "ISSUED",
	1: "TRADING",
	2: "REDEEMED",
}
var CommercialPaper_State_value = map[string]int32{
	"ISSUED":   0,
	"TRADING":  1,
	"REDEEMED": 2,
}

func (x CommercialPaper_State) String() string {
	return proto.EnumName(CommercialPaper_State_name, int32(x))
}
func (CommercialPaper_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// Commercial Paper state entry
type CommercialPaper struct {
	// Issuer and Paper number comprises composite primary key of Commercial paper entry
	Issuer       string                     `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
	PaperNumber  string                     `protobuf:"bytes,2,opt,name=paper_number,json=paperNumber" json:"paper_number,omitempty"`
	Owner        string                     `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	IssueDate    *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=issue_date,json=issueDate" json:"issue_date,omitempty"`
	MaturityDate *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=maturity_date,json=maturityDate" json:"maturity_date,omitempty"`
	FaceValue    int32                      `protobuf:"varint,6,opt,name=face_value,json=faceValue" json:"face_value,omitempty"`
	State        CommercialPaper_State      `protobuf:"varint,7,opt,name=state,enum=cckit.examples.cpaper_extended.schema.CommercialPaper_State" json:"state,omitempty"`
	// Additional unique field for entry
	ExternalId string `protobuf:"bytes,8,opt,name=external_id,json=externalId" json:"external_id,omitempty"`
}

func (m *CommercialPaper) Reset()                    { *m = CommercialPaper{} }
func (m *CommercialPaper) String() string            { return proto.CompactTextString(m) }
func (*CommercialPaper) ProtoMessage()               {}
func (*CommercialPaper) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CommercialPaper) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *CommercialPaper) GetPaperNumber() string {
	if m != nil {
		return m.PaperNumber
	}
	return ""
}

func (m *CommercialPaper) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CommercialPaper) GetIssueDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.IssueDate
	}
	return nil
}

func (m *CommercialPaper) GetMaturityDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.MaturityDate
	}
	return nil
}

func (m *CommercialPaper) GetFaceValue() int32 {
	if m != nil {
		return m.FaceValue
	}
	return 0
}

func (m *CommercialPaper) GetState() CommercialPaper_State {
	if m != nil {
		return m.State
	}
	return CommercialPaper_ISSUED
}

func (m *CommercialPaper) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

// CommercialPaperId identifier part
type CommercialPaperId struct {
	Issuer      string `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
	PaperNumber string `protobuf:"bytes,2,opt,name=paper_number,json=paperNumber" json:"paper_number,omitempty"`
}

func (m *CommercialPaperId) Reset()                    { *m = CommercialPaperId{} }
func (m *CommercialPaperId) String() string            { return proto.CompactTextString(m) }
func (*CommercialPaperId) ProtoMessage()               {}
func (*CommercialPaperId) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CommercialPaperId) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *CommercialPaperId) GetPaperNumber() string {
	if m != nil {
		return m.PaperNumber
	}
	return ""
}

// Container for returning multiple entities
type CommercialPaperList struct {
	Items []*CommercialPaper `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *CommercialPaperList) Reset()                    { *m = CommercialPaperList{} }
func (m *CommercialPaperList) String() string            { return proto.CompactTextString(m) }
func (*CommercialPaperList) ProtoMessage()               {}
func (*CommercialPaperList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *CommercialPaperList) GetItems() []*CommercialPaper {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*CommercialPaper)(nil), "cckit.examples.cpaper_extended.schema.CommercialPaper")
	proto.RegisterType((*CommercialPaperId)(nil), "cckit.examples.cpaper_extended.schema.CommercialPaperId")
	proto.RegisterType((*CommercialPaperList)(nil), "cckit.examples.cpaper_extended.schema.CommercialPaperList")
	proto.RegisterEnum("cckit.examples.cpaper_extended.schema.CommercialPaper_State", CommercialPaper_State_name, CommercialPaper_State_value)
}

func init() { proto.RegisterFile("state.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xd1, 0x8b, 0xd3, 0x40,
	0x10, 0xc6, 0xcd, 0xd5, 0xe4, 0xda, 0x49, 0xd5, 0xba, 0x8a, 0x2c, 0x07, 0x72, 0xb1, 0x20, 0xe4,
	0x69, 0x0f, 0x2a, 0x08, 0x82, 0x20, 0x6a, 0x82, 0x04, 0xce, 0x22, 0xdb, 0xd3, 0x07, 0x5f, 0xc2,
	0x36, 0x99, 0x3b, 0x17, 0xb3, 0x49, 0xd8, 0xdd, 0x68, 0xfd, 0x77, 0xfc, 0x4b, 0x25, 0xbb, 0xe6,
	0xa5, 0x2f, 0x8a, 0x3e, 0xce, 0x37, 0xf3, 0xfd, 0x66, 0xf8, 0x06, 0x62, 0x63, 0x85, 0x45, 0xd6,
	0xeb, 0xce, 0x76, 0xe4, 0x69, 0x55, 0x7d, 0x95, 0x96, 0xe1, 0x41, 0xa8, 0xbe, 0x41, 0xc3, 0xaa,
	0x5e, 0xf4, 0xa8, 0x4b, 0x3c, 0x58, 0x6c, 0x6b, 0xac, 0x99, 0xa9, 0xbe, 0xa0, 0x12, 0x67, 0xe7,
	0x37, 0x5d, 0x77, 0xd3, 0xe0, 0x85, 0x33, 0xed, 0x87, 0xeb, 0x0b, 0x2b, 0x15, 0x1a, 0x2b, 0x54,
	0xef, 0x39, 0xeb, 0x9f, 0x33, 0xb8, 0xf7, 0xb6, 0x53, 0x0a, 0x75, 0x25, 0x45, 0xf3, 0x61, 0x84,
	0x90, 0x47, 0x10, 0x49, 0x63, 0x06, 0xd4, 0x34, 0x48, 0x82, 0x74, 0xc1, 0x7f, 0x57, 0xe4, 0x09,
	0x2c, 0xfd, 0x96, 0x76, 0x50, 0x7b, 0xd4, 0xf4, 0xc4, 0x75, 0x63, 0xa7, 0x6d, 0x9d, 0x44, 0x1e,
	0x42, 0xd8, 0x7d, 0x6f, 0x51, 0xd3, 0x99, 0xeb, 0xf9, 0x82, 0xbc, 0x00, 0x70, 0x88, 0xb2, 0x16,
	0x16, 0xe9, 0xed, 0x24, 0x48, 0xe3, 0xcd, 0x19, 0xf3, 0xa7, 0xb1, 0xe9, 0x34, 0x76, 0x35, 0x9d,
	0xc6, 0x17, 0x6e, 0x3a, 0x13, 0x16, 0xc9, 0x2b, 0xb8, 0xa3, 0x84, 0x1d, 0xb4, 0xb4, 0x3f, 0xbc,
	0x3b, 0xfc, 0xa3, 0x7b, 0x39, 0x19, 0x1c, 0xe0, 0x31, 0xc0, 0xb5, 0xa8, 0xb0, 0xfc, 0x26, 0x9a,
	0x01, 0x69, 0x94, 0x04, 0x69, 0xc8, 0x17, 0xa3, 0xf2, 0x69, 0x14, 0x08, 0x87, 0xd0, 0xc5, 0x4a,
	0x4f, 0x93, 0x20, 0xbd, 0xbb, 0x79, 0xc9, 0xfe, 0x2a, 0x57, 0x76, 0x14, 0x19, 0xdb, 0x8d, 0x0c,
	0xee, 0x51, 0xe4, 0x1c, 0xe2, 0x71, 0x5e, 0xb7, 0xa2, 0x29, 0x65, 0x4d, 0xe7, 0x2e, 0x0a, 0x98,
	0xa4, 0xa2, 0x5e, 0x33, 0x08, 0x9d, 0x81, 0x00, 0x44, 0xc5, 0x6e, 0xf7, 0x31, 0xcf, 0x56, 0xb7,
	0x48, 0x0c, 0xa7, 0x57, 0xfc, 0x75, 0x56, 0x6c, 0xdf, 0xad, 0x02, 0xb2, 0x84, 0x39, 0xcf, 0xb3,
	0x3c, 0x7f, 0x9f, 0x67, 0xab, 0x93, 0xf5, 0x16, 0xee, 0x1f, 0x2d, 0x2c, 0xea, 0xff, 0xf8, 0xd2,
	0xba, 0x82, 0x07, 0x47, 0xbc, 0x4b, 0x69, 0x2c, 0xb9, 0x84, 0x50, 0x5a, 0x54, 0x86, 0x06, 0xc9,
	0x2c, 0x8d, 0x37, 0xcf, 0xff, 0x2d, 0x0b, 0xee, 0x21, 0x6f, 0xe6, 0x9f, 0x23, 0x3f, 0xb0, 0x8f,
	0xdc, 0x93, 0x9e, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x68, 0x2f, 0x49, 0x95, 0xc1, 0x02, 0x00,
	0x00,
}
