package iproute2

import (
	"encoding/binary"
	"errors"
	"reflect"
	"unsafe"

	"golang.org/x/sys/unix"
)

// size of some structures.
const (
	SizeofInetDiagReq = int(unsafe.Sizeof(InetDiagReq{}))
	SizeofInetDiagMsg = int(unsafe.Sizeof(InetDiagMsg{}))
	SizeofIfInfoMsg   = unix.SizeofIfInfomsg
	SizeofNdMsg       = unix.SizeofNdMsg
)

// An InetDiagReq is a request message for sock diag netlink.
type InetDiagReq struct {
	Family   uint8
	Protocol uint8
	Ext      uint8
	Pad      uint8
	States   uint32
	InetDiagSockID
}

// An InetDiagMsg is a response message for sock diag netlink.
type InetDiagMsg struct {
	Family  uint8
	State   uint8
	Timer   uint8
	Retrans uint8

	InetDiagSockID

	Expires uint32
	RQueue  uint32
	WQueue  uint32
	UID     uint32
	Inode   uint32
}

// An InetDiagSockID contains some info about a socket.
type InetDiagSockID struct {
	Sport   uint16 // big endian
	Dport   uint16 // big endian
	Saddr   [16]byte
	Daddr   [16]byte
	Ifindex uint32
	Cookie  [2]uint32
}

// MarshalBinary marshals an inet diag request message to byte slice.
func (req *InetDiagReq) MarshalBinary() (data []byte, err error) {
	data = struct2bytes(unsafe.Pointer(req), SizeofInetDiagReq)
	be, offset := binary.BigEndian, 8
	be.PutUint16(data[offset:], req.Sport)
	be.PutUint16(data[offset+2:], req.Dport)
	return data, nil
}

// UnmarshalBinary unmarshals an inet diag response message from byte slice.
func (msg *InetDiagMsg) UnmarshalBinary(data []byte) error {
	length := SizeofInetDiagMsg
	if len(data) < length {
		return errors.New("InetDiagMsg: not enough data to unmarshal")
	}

	dataSlice := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	newMsg := (*InetDiagMsg)(unsafe.Pointer(dataSlice.Data))
	*msg = *newMsg

	be, offset := binary.BigEndian, 4
	msg.Sport = be.Uint16(data[offset:])
	msg.Dport = be.Uint16(data[offset+2:])

	return nil
}

// NdMsg is a neighbour message,
// that's an alias of golang.org/x/sys/unix.NdMsg
type NdMsg unix.NdMsg

// MarshalBinary marshals a neighbour message to byte slice.
func (m *NdMsg) MarshalBinary() ([]byte, error) {
	return struct2bytes(unsafe.Pointer(m), SizeofNdMsg), nil
}

// A NdAttrCacheInfo is the cache info in the neighbour/fdb message.
type NdAttrCacheInfo struct {
	Confirmed uint32
	Used      uint32
	Updated   uint32
	RefCount  uint32
}

// UnmarshalBinary unmarshals a neighbour attribute's cache info
// from byte slice.
func (c *NdAttrCacheInfo) UnmarshalBinary(data []byte) error {
	sizeof := int(unsafe.Sizeof(*c))
	if len(data) < sizeof {
		return errors.New("NdAttrCacheInfo: not enough data to unmarshal")
	}

	newCacheInfo := (*NdAttrCacheInfo)(unsafe.Pointer(&data[0]))
	*c = *newCacheInfo
	return nil
}

// IfInfoMsg is an interface information message,
// that's an alias of golang.org/x/sys/unix.IfInfoMsg
type IfInfoMsg unix.IfInfoMsg

// MarshalBinary marshals an interface informat message
// to byte slice.
func (m *IfInfoMsg) MarshalBinary() ([]byte, error) {
	return struct2bytes(unsafe.Pointer(m), SizeofIfInfoMsg), nil
}

// UnmarshalBinary unmarshals an interface information message
// from byte slice.
func (m *IfInfoMsg) UnmarshalBinary(data []byte) error {
	if len(data) < SizeofIfInfoMsg {
		return errors.New("IfInfoMsg: not enough data to unmarshal")
	}

	newIfiMsg := (*IfInfoMsg)(unsafe.Pointer(&data[0]))
	*m = *newIfiMsg
	return nil
}

// struct2bytes converts forcely a struct to byte slice in place.
func struct2bytes(p unsafe.Pointer, length int) []byte {
	var dataSlice reflect.SliceHeader
	dataSlice.Len = length
	dataSlice.Cap = length
	dataSlice.Data = uintptr(p)
	return *(*[]byte)(unsafe.Pointer(&dataSlice))
}
