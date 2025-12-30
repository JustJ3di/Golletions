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

// At returns the value at the given index.
// It returns an error if the index is out of bounds.
func (zl *Ziplist) At(index int) (any, error) {
	// 1. Read the current number of items from the header (bytes 6-10)
	countIdx := 6
	if len(zl.bytes) < 10 {
		return nil, fmt.Errorf("ziplist is empty or malformed")
	}
	currentCount := binary.LittleEndian.Uint32(zl.bytes[countIdx : countIdx+4])

	// 2. Bounds Check
	if index < 0 || uint32(index) >= currentCount {
		return nil, fmt.Errorf("index out of range")
	}

	// 3. Traverse the list to find the start offset of the requested index
	// Start after the header (10 bytes)
	cursor := 10
	for i := 0; i < index; i++ {
		size, err := zl.getElementSize(cursor)
		if err != nil {
			return nil, err
		}
		cursor += size
	}

	// 4. Decode the value at the current cursor
	if cursor >= len(zl.bytes) {
		return nil, fmt.Errorf("cursor out of bounds")
	}

	typ := zl.bytes[cursor]
	// Advance cursor past the type byte
	dataStart := cursor + 1

	switch typ {
	case TYPE_UINT8:
		return zl.bytes[dataStart], nil

	case TYPE_UINT16:
		return binary.LittleEndian.Uint16(zl.bytes[dataStart : dataStart+2]), nil

	case TYPE_UINT32:
		return binary.LittleEndian.Uint32(zl.bytes[dataStart : dataStart+4]), nil

	case TYPE_UINT64:
		return binary.LittleEndian.Uint64(zl.bytes[dataStart : dataStart+8]), nil

	case TYPE_INT8:
		// Go doesn't have a direct LittleEndian.Uint8, just cast the byte
		return int8(zl.bytes[dataStart]), nil

	case TYPE_INT16:
		return int16(binary.LittleEndian.Uint16(zl.bytes[dataStart : dataStart+2])), nil

	case TYPE_INT32:
		return int32(binary.LittleEndian.Uint32(zl.bytes[dataStart : dataStart+4])), nil

	case TYPE_INT64:
		return int64(binary.LittleEndian.Uint64(zl.bytes[dataStart : dataStart+8])), nil

	case TYPE_FLOAT32:
		bits := binary.LittleEndian.Uint32(zl.bytes[dataStart : dataStart+4])
		return math.Float32frombits(bits), nil

	case TYPE_FLOAT64:
		bits := binary.LittleEndian.Uint64(zl.bytes[dataStart : dataStart+8])
		return math.Float64frombits(bits), nil

	case TYPE_BOOL:
		val := zl.bytes[dataStart]
		return val == 1, nil

	case TYPE_STRING:
		// String format: [Type] [Len (4 bytes)] [String Data]
		strLen := binary.LittleEndian.Uint32(zl.bytes[dataStart : dataStart+4])
		strBodyStart := dataStart + 4
		strBodyEnd := strBodyStart + int(strLen)

		if strBodyEnd > len(zl.bytes) {
			return nil, fmt.Errorf("malformed string length")
		}
		return string(zl.bytes[strBodyStart:strBodyEnd]), nil

	case TYPE_END:
		return nil, fmt.Errorf("accessed end of list unexpectedly")

	default:
		return nil, fmt.Errorf("unknown type at index %d: %x", index, typ)
	}
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
