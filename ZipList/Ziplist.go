package ziplist

import (
	"encoding/binary"
	"fmt"
	"math"
)

const (
	TYPE_END = 0xff

	//unsigned integers
	TYPE_UINT8  = 0x00
	TYPE_UINT16 = 0x01
	TYPE_UINT32 = 0x02
	TYPE_UINT64 = 0x03

	//Signed Integers
	TYPE_INT8  = 0x04
	TYPE_INT16 = 0x05
	TYPE_INT32 = 0x06
	TYPE_INT64 = 0x07

	//Floats
	TYPE_FLOAT32 = 0x08
	TYPE_FLOAT64 = 0x09

	//Base
	TYPE_BOOL   = 0x0a
	TYPE_STRING = 0x0b

	TYPE_LEN        = 0x0d
	TYPE_TOTAL_BYTE = 0x0f
)

type Ziplist struct {
	bytes  []byte
	cursor uint32
}

func New(capacity uint32) *Ziplist {
	zl := &Ziplist{bytes: make([]byte, 0, capacity)}
	//SET HEADER
	zl.bytes = append(zl.bytes, TYPE_TOTAL_BYTE)
	zl.bytes = binary.LittleEndian.AppendUint32(zl.bytes, 0)
	zl.bytes = append(zl.bytes, TYPE_LEN)
	zl.bytes = binary.LittleEndian.AppendUint32(zl.bytes, 0)
	zl.bytes = append(zl.bytes, TYPE_END)
	currentLen := uint32(len(zl.bytes))
	binary.LittleEndian.PutUint32(zl.bytes[1:5], currentLen)

	return zl
}

func (zl *Ziplist) getElementSize(offset int) (int, error) {
	if offset >= len(zl.bytes) {
		return 0, fmt.Errorf("out of bounds")
	}

	typ := zl.bytes[offset]

	switch typ {
	case TYPE_END:
		return 0, nil

	case TYPE_UINT8, TYPE_INT8, TYPE_BOOL:
		return 1 + 1, nil
	case TYPE_UINT16, TYPE_INT16:
		return 1 + 2, nil
	case TYPE_UINT32, TYPE_INT32, TYPE_FLOAT32:
		return 1 + 4, nil
	case TYPE_UINT64, TYPE_INT64, TYPE_FLOAT64:
		return 1 + 8, nil

	case TYPE_STRING:

		if offset+5 > len(zl.bytes) {
			return 0, fmt.Errorf("malformed string header")
		}
		strLen := binary.LittleEndian.Uint32(zl.bytes[offset+1 : offset+5])
		return 1 + 4 + int(strLen), nil

	default:
		return 0, fmt.Errorf("unknown type: %x", typ)
	}
}

func (zl *Ziplist) Remove(index int) error {

	countIdx := 6
	currentCount := binary.LittleEndian.Uint32(zl.bytes[countIdx : countIdx+4])

	if index < 0 || uint32(index) >= currentCount {
		return fmt.Errorf("index out of range")
	}

	currentPos := 10

	for i := 0; i < index; i++ {
		size, err := zl.getElementSize(currentPos)
		if err != nil {
			return err
		}
		currentPos += size
	}

	sizeToRemove, err := zl.getElementSize(currentPos)
	if err != nil {
		return err
	}

	copy(zl.bytes[currentPos:], zl.bytes[currentPos+sizeToRemove:])

	zl.bytes = zl.bytes[:len(zl.bytes)-sizeToRemove]

	binary.LittleEndian.PutUint32(zl.bytes[countIdx:countIdx+4], currentCount-1)

	totalLenIdx := 1
	binary.LittleEndian.PutUint32(zl.bytes[totalLenIdx:totalLenIdx+4], uint32(len(zl.bytes)))

	return nil
}

func (zl *Ziplist) Clear() {
	zl.bytes = zl.bytes[:0]
	zl.bytes = append(zl.bytes, TYPE_TOTAL_BYTE)
	zl.bytes = binary.LittleEndian.AppendUint32(zl.bytes, 0)
	zl.bytes = append(zl.bytes, TYPE_LEN)
	zl.bytes = binary.LittleEndian.AppendUint32(zl.bytes, 0)
	zl.bytes = append(zl.bytes, TYPE_END)
	currentLen := uint32(len(zl.bytes))
	binary.LittleEndian.PutUint32(zl.bytes[1:5], currentLen)
}

func (zl *Ziplist) Push(value any) error {

	if len(zl.bytes) > 0 && zl.bytes[len(zl.bytes)-1] == TYPE_END {
		zl.bytes = zl.bytes[:len(zl.bytes)-1]
	}

	switch v := value.(type) {

	case uint8:
		zl.bytes = append(zl.bytes, TYPE_UINT8)
		zl.bytes = append(zl.bytes, v)

	case uint16:
		zl.bytes = append(zl.bytes, TYPE_UINT16)
		zl.bytes = binary.LittleEndian.AppendUint16(zl.bytes, v)

	case uint32:
		zl.bytes = append(zl.bytes, TYPE_UINT32)
		zl.bytes = binary.LittleEndian.AppendUint32(zl.bytes, v)

	case int32:
		zl.bytes = append(zl.bytes, TYPE_INT32)
		zl.bytes = binary.LittleEndian.AppendUint32(zl.bytes, uint32(v))

	case float64:
		zl.bytes = append(zl.bytes, TYPE_FLOAT64)
		bits := math.Float64bits(v)
		zl.bytes = binary.LittleEndian.AppendUint64(zl.bytes, bits)

	case string:
		zl.bytes = append(zl.bytes, TYPE_STRING)
		strLen := uint32(len(v))
		zl.bytes = binary.LittleEndian.AppendUint32(zl.bytes, strLen)
		zl.bytes = append(zl.bytes, []byte(v)...)

	case bool:
		zl.bytes = append(zl.bytes, TYPE_BOOL)
		if v {
			zl.bytes = append(zl.bytes, 1)
		} else {
			zl.bytes = append(zl.bytes, 0)
		}

	default:
		zl.bytes = append(zl.bytes, TYPE_END)
		return fmt.Errorf("tipo non supportato: %T", v)
	}

	zl.bytes = append(zl.bytes, TYPE_END)

	zl.updateHeader()

	return nil
}

func (zl *Ziplist) updateHeader() {
	// Offset 0: TYPE_TOTAL_BYTE (1 byte)
	// Offset 1: Total Byte (4 byte)
	// Offset 5: TYPE_LEN (1 byte)
	// Offset 6: Valore Len (4 byte)
	totalLen := uint32(len(zl.bytes))
	binary.LittleEndian.PutUint32(zl.bytes[1:5], totalLen)
	currentCount := binary.LittleEndian.Uint32(zl.bytes[6:10])
	binary.LittleEndian.PutUint32(zl.bytes[6:10], currentCount+1)
}
