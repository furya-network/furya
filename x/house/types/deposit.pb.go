// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furya/house/deposit.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Deposit represents the deposit against a market held by an account.
type Deposit struct {
	// creator is the bech32-encoded address of the depositor.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	// market_uid is the uid of market against which deposit is being
	// made.
	MarketUID string `protobuf:"bytes,2,opt,name=market_uid,proto3" json:"market_uid"`
	// participation_index index corresponding to the book participation
	ParticipationIndex uint64 `protobuf:"varint,3,opt,name=participation_index,json=participationIndex,proto3" json:"participation_index,omitempty" yaml:"participation_index"`
	// amount is the amount being deposited.
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount" yaml:"amount"`
	// fee is deducted from the amount on deposition.
	Fee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=fee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"fee" yaml:"fee"`
	// liquidity is the liquidity being provided to the house after fee deduction.
	Liquidity github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=liquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"liquidity" yaml:"liquidity"`
	// withdrawal_count is the total count of the withdrawal attempts
	WithdrawalCount uint64 `protobuf:"varint,7,opt,name=withdrawal_count,json=withdrawalCount,proto3" json:"withdrawal_count,omitempty" yaml:"withdrawals"`
	// total_withdrawal_amount is the total amount of withdrawal attempts
	TotalWithdrawalAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=total_withdrawal_amount,json=totalWithdrawalAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_withdrawal_amount" yaml:"total_withdrawal_amount"`
}

func (m *Deposit) Reset()      { *m = Deposit{} }
func (*Deposit) ProtoMessage() {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f2840908fc45a1, []int{0}
}
func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return m.Size()
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Deposit)(nil), "furyanetwork.furya.house.Deposit")
}

func init() { proto.RegisterFile("furya/house/deposit.proto", fileDescriptor_c6f2840908fc45a1) }

var fileDescriptor_c6f2840908fc45a1 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x6d, 0xda, 0x26, 0xe4, 0xc4, 0x8f, 0xea, 0x28, 0xd4, 0x2a, 0x92, 0xaf, 0xba, 0x01,
	0x65, 0xa0, 0xf6, 0xc0, 0x56, 0x90, 0x50, 0x4d, 0x97, 0x0c, 0xfc, 0xd0, 0x49, 0xa8, 0x12, 0x4b,
	0xb8, 0xda, 0x57, 0xe7, 0x94, 0xd8, 0x67, 0x7c, 0x67, 0xa5, 0xf9, 0x0f, 0x3a, 0x32, 0x32, 0xe6,
	0x0f, 0xe0, 0x0f, 0xe9, 0xd8, 0x11, 0x31, 0x9c, 0x50, 0xb2, 0x20, 0xc6, 0xfc, 0x05, 0x28, 0x77,
	0x26, 0x31, 0x12, 0x0c, 0x99, 0xfc, 0xfc, 0xde, 0xf7, 0x7d, 0xde, 0xd7, 0x4f, 0x7e, 0x60, 0x5f,
	0xa6, 0x2c, 0x1c, 0x88, 0x4a, 0xb2, 0x30, 0x61, 0x85, 0x90, 0x5c, 0x05, 0x45, 0x29, 0x94, 0x80,
	0x7b, 0x32, 0x65, 0x39, 0x53, 0x63, 0x51, 0x0e, 0x03, 0x99, 0xb2, 0xc0, 0x68, 0x0e, 0xf6, 0x52,
	0x91, 0x0a, 0x23, 0x08, 0x97, 0x91, 0xd5, 0xe2, 0xaf, 0x3b, 0xa0, 0x7d, 0x6a, 0xbb, 0xe1, 0x53,
	0xd0, 0x8e, 0x4b, 0x46, 0x95, 0x28, 0x3d, 0xf7, 0xd0, 0xed, 0x76, 0x22, 0xb8, 0xd0, 0xe8, 0xde,
	0x84, 0x66, 0xa3, 0x63, 0x5c, 0x17, 0x30, 0xf9, 0x23, 0x81, 0xcf, 0x01, 0xc8, 0x68, 0x39, 0x64,
	0xaa, 0x5f, 0xf1, 0xc4, 0xbb, 0x65, 0x1a, 0x1e, 0xcf, 0x34, 0xea, 0xbc, 0x36, 0xd9, 0xf7, 0xbd,
	0xd3, 0x5f, 0x1a, 0x35, 0x24, 0xa4, 0x11, 0xc3, 0xb7, 0xe0, 0x41, 0x41, 0x4b, 0xc5, 0x63, 0x5e,
	0x50, 0xc5, 0x45, 0xde, 0xe7, 0x79, 0xc2, 0x2e, 0xbd, 0xad, 0x43, 0xb7, 0xbb, 0x1d, 0xf9, 0x0b,
	0x8d, 0x0e, 0xec, 0xd8, 0x7f, 0x88, 0x30, 0x81, 0x7f, 0x65, 0x7b, 0xcb, 0x24, 0x3c, 0x03, 0x2d,
	0x9a, 0x89, 0x2a, 0x57, 0xde, 0xb6, 0x71, 0xf2, 0xf2, 0x5a, 0x23, 0xe7, 0xbb, 0x46, 0x4f, 0x52,
	0xae, 0x06, 0xd5, 0x79, 0x10, 0x8b, 0x2c, 0x8c, 0x85, 0xcc, 0x84, 0xac, 0x1f, 0x47, 0x32, 0x19,
	0x86, 0x6a, 0x52, 0x30, 0x19, 0xf4, 0x72, 0xb5, 0xd0, 0xe8, 0xae, 0x9d, 0x68, 0x29, 0x98, 0xd4,
	0x38, 0xf8, 0x06, 0x6c, 0x5d, 0x30, 0xe6, 0xed, 0x18, 0xea, 0x8b, 0x8d, 0xa9, 0xc0, 0x52, 0x2f,
	0x18, 0xc3, 0x64, 0x09, 0x82, 0x1f, 0x41, 0x67, 0xc4, 0x3f, 0x55, 0x3c, 0xe1, 0x6a, 0xe2, 0xb5,
	0x0c, 0x35, 0xda, 0x98, 0xba, 0x6b, 0xa9, 0x2b, 0x10, 0x26, 0x6b, 0x28, 0x3c, 0x01, 0xbb, 0x63,
	0xae, 0x06, 0x49, 0x49, 0xc7, 0x74, 0xd4, 0x8f, 0xcd, 0x52, 0xda, 0x66, 0xb1, 0x8f, 0x16, 0x1a,
	0x41, 0xdb, 0xba, 0x56, 0x48, 0x4c, 0xee, 0xaf, 0xdf, 0x5e, 0x99, 0x8f, 0xbe, 0x72, 0xc1, 0xbe,
	0x12, 0x8a, 0x8e, 0xfa, 0x0d, 0x52, 0xbd, 0xdf, 0xdb, 0xc6, 0xf3, 0xbb, 0x8d, 0x3d, 0xfb, 0x76,
	0xf0, 0x7f, 0xb0, 0x98, 0x3c, 0x34, 0x95, 0xb3, 0x55, 0xe1, 0xc4, 0xe4, 0x8f, 0xef, 0x5c, 0x4d,
	0x91, 0xf3, 0x65, 0x8a, 0x9c, 0x9f, 0x53, 0xe4, 0x44, 0xd1, 0xf5, 0xcc, 0x77, 0x6f, 0x66, 0xbe,
	0xfb, 0x63, 0xe6, 0xbb, 0x9f, 0xe7, 0xbe, 0x73, 0x33, 0xf7, 0x9d, 0x6f, 0x73, 0xdf, 0xf9, 0xd0,
	0x6d, 0x18, 0x91, 0x29, 0x3b, 0xaa, 0x0f, 0x60, 0x19, 0x87, 0x97, 0xf5, 0x99, 0x18, 0x3b, 0xe7,
	0x2d, 0xf3, 0xe7, 0x3f, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xad, 0xcf, 0x24, 0xa6, 0x40, 0x03,
	0x00, 0x00,
}

func (m *Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalWithdrawalAmount.Size()
		i -= size
		if _, err := m.TotalWithdrawalAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.WithdrawalCount != 0 {
		i = encodeVarintDeposit(dAtA, i, uint64(m.WithdrawalCount))
		i--
		dAtA[i] = 0x38
	}
	{
		size := m.Liquidity.Size()
		i -= size
		if _, err := m.Liquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Fee.Size()
		i -= size
		if _, err := m.Fee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.ParticipationIndex != 0 {
		i = encodeVarintDeposit(dAtA, i, uint64(m.ParticipationIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.MarketUID) > 0 {
		i -= len(m.MarketUID)
		copy(dAtA[i:], m.MarketUID)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.MarketUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeposit(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeposit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	l = len(m.MarketUID)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	if m.ParticipationIndex != 0 {
		n += 1 + sovDeposit(uint64(m.ParticipationIndex))
	}
	l = m.Amount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	l = m.Fee.Size()
	n += 1 + l + sovDeposit(uint64(l))
	l = m.Liquidity.Size()
	n += 1 + l + sovDeposit(uint64(l))
	if m.WithdrawalCount != 0 {
		n += 1 + sovDeposit(uint64(m.WithdrawalCount))
	}
	l = m.TotalWithdrawalAmount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	return n
}

func sovDeposit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeposit(x uint64) (n int) {
	return sovDeposit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeposit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationIndex", wireType)
			}
			m.ParticipationIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipationIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Liquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalCount", wireType)
			}
			m.WithdrawalCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawalCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWithdrawalAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalWithdrawalAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeposit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeposit
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
func skipDeposit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDeposit
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
			if length < 0 {
				return 0, ErrInvalidLengthDeposit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeposit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeposit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeposit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeposit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeposit = fmt.Errorf("proto: unexpected end of group")
)
