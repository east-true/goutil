package parser

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

const (
	BIT16 uint8 = iota
	BIT32
	BIT64
	INT16
	UINT16
	INT32
	UINT32
	INT64
	UINT64
	FLOAT32
	FLOAT64
	INT16ARR
	UINT16ARR
	INT32ARR
	UINT32ARR
	INT64ARR
	UINT64ARR
	FLOAT32ARR
	FLOAT64ARR

	// TODO : add struct of order by, add  ByteParser interface
	LITTLE_LOWER string = "little"
	BIG_LOWER    string = "big"
)

type Parser struct {
	order string
}

func New(order string) *Parser {
	return &Parser{
		order: order,
	}
}

func (p *Parser) ToBit(datum byte) []uint8 {
	var bits []uint8 = make([]uint8, 8)
	for i := range bits {
		bits[i] = datum & 1 // and
		datum >>= 1         // shift bit
	}

	return bits
}

func (p *Parser) ToBitArr(datum []byte) []uint8 {
	var bits []uint8 = make([]uint8, len(datum)*8)
	for i, data := range datum {
		res := p.ToBit(data)
		for j, bit := range res {
			bits[j+(8*i)] = bit
		}
	}

	return bits
}

func (p *Parser) ToInt16(b []byte) int16 {
	if p.order == BIG_LOWER {
		return int16(binary.BigEndian.Uint16(b))
	} else {
		return int16(binary.LittleEndian.Uint16(b))
	}
}

func (p *Parser) ToInt16Arr(b []byte) ([]int16, error) {
	len := len(b)
	if len%2 != 0 {
		return nil, errors.New("not matched units (1word, 2byte, 16bit)")
	}

	var data []int16 = make([]int16, len/2)
	for i := range data {
		data[i] = p.ToInt16(b[2*i : 2*(i+1)])
	}

	return data, nil
}

func (p *Parser) ToUint16(b []byte) uint16 {
	if p.order == BIG_LOWER {
		return binary.BigEndian.Uint16(b)
	} else {
		return binary.LittleEndian.Uint16(b)
	}
}

func (p *Parser) ToUint16Arr(b []byte) ([]uint16, error) {
	len := len(b)
	if len%2 != 0 {
		return nil, errors.New("not matched units (1word, 2byte, 16bit)")
	}

	var data []uint16 = make([]uint16, len/2)
	for i := range data {
		data[i] = p.ToUint16(b[2*i : 2*(i+1)])
	}

	return data, nil
}

func (p *Parser) ToInt32(b []byte) int32 {
	if p.order == BIG_LOWER {
		return int32(binary.BigEndian.Uint32(b))
	} else {
		return int32(binary.LittleEndian.Uint32(b))
	}
}

func (p *Parser) ToInt32Arr(b []byte) ([]int32, error) {
	len := len(b)
	if len%4 != 0 {
		return nil, errors.New("not matched units (2word, 4byte, 32bit)")
	}

	var data []int32 = make([]int32, len/4)
	for i := range data {
		data[i] = p.ToInt32(b[4*i : 4*(i+1)])
	}

	return data, nil
}

func (p *Parser) ToUint32(b []byte) uint32 {
	if p.order == BIG_LOWER {
		return binary.BigEndian.Uint32(b)
	} else {
		return binary.LittleEndian.Uint32(b)
	}
}

func (p *Parser) ToUint32Arr(b []byte) ([]uint32, error) {
	len := len(b)
	if len%4 != 0 {
		return nil, errors.New("not matched units (2word, 4byte, 32bit)")
	}

	var data []uint32 = make([]uint32, len/4)
	for i := range data {
		data[i] = p.ToUint32(b[4*i : 4*(i+1)])
	}

	return data, nil
}

func (p *Parser) ToInt64(b []byte) int64 {
	if p.order == BIG_LOWER {
		return int64(binary.BigEndian.Uint64(b))
	} else {
		return int64(binary.LittleEndian.Uint64(b))
	}
}

func (p *Parser) ToInt64Arr(b []byte) ([]int64, error) {
	len := len(b)
	if len%8 != 0 {
		return nil, errors.New("not matched units (4word, 8byte, 64bit)")
	}

	var data []int64 = make([]int64, len/8)
	for i := range data {
		data[i] = p.ToInt64(b[8*i : 8*(i+1)])
	}

	return data, nil
}

func (p *Parser) ToUint64(b []byte) uint64 {
	if p.order == BIG_LOWER {
		return binary.BigEndian.Uint64(b)
	} else {
		return binary.LittleEndian.Uint64(b)
	}
}

func (p *Parser) ToUint64Arr(b []byte) ([]uint64, error) {
	len := len(b)
	if len%8 != 0 {
		return nil, errors.New("not matched units (4word, 8byte, 64bit)")
	}

	var data []uint64 = make([]uint64, len/8)
	for i := range data {
		data[i] = p.ToUint64(b[8*i : 8*(i+1)])
	}

	return data, nil
}

func (p *Parser) ToFloat32(b []byte) float32 {
	var datum uint32
	if p.order == BIG_LOWER {
		datum = binary.BigEndian.Uint32(b)
	} else {
		datum = binary.LittleEndian.Uint32(b)
	}

	return math.Float32frombits(datum)
}

func (p *Parser) ToFloat32Arr(b []byte) ([]float32, error) {
	len := len(b)
	if len%4 != 0 {
		return nil, errors.New("not matched units (2word, 4byte, 32bit)")
	}

	var data []float32 = make([]float32, len/4)
	for i := range data {
		data[i] = p.ToFloat32(b[4*i : 4*(i+1)])
	}

	return data, nil
}

func (p *Parser) ToFloat64(b []byte) float64 {
	var datum uint64
	if p.order == BIG_LOWER {
		datum = binary.BigEndian.Uint64(b)
	} else {
		datum = binary.LittleEndian.Uint64(b)
	}

	return math.Float64frombits(datum)
}

func (p *Parser) ToFloat64Arr(b []byte) ([]float64, error) {
	len := len(b)
	if len%8 != 0 {
		return nil, errors.New("not matched units (4word, 8byte, 64bit)")
	}

	var data []float64 = make([]float64, len/8)
	for i := range data {
		data[i] = p.ToFloat64(b[8*i : 8*(i+1)])
	}

	return data, nil
}

func (p *Parser) ToAnyOf(dataType uint8, b []byte) (any, error) {
	switch dataType {
	case BIT16:
		return fmt.Sprintf("%016b", p.ToInt16(b)), nil
	case BIT32:
		return fmt.Sprintf("%032b", p.ToInt32(b)), nil
	case BIT64:
		return fmt.Sprintf("%064b", p.ToInt64(b)), nil
	case INT16:
		return p.ToInt16(b), nil
	case UINT16:
		return p.ToUint16(b), nil
	case INT32:
		return p.ToInt32(b), nil
	case UINT32:
		return p.ToUint32(b), nil
	case INT64:
		return p.ToInt64(b), nil
	case UINT64:
		return p.ToUint64(b), nil
	case FLOAT32:
		return p.ToFloat32(b), nil
	case FLOAT64:
		return p.ToFloat64(b), nil
	case INT16ARR:
		data, err := p.ToInt16Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case UINT16ARR:
		data, err := p.ToUint16Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case INT32ARR:
		data, err := p.ToInt32Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case UINT32ARR:
		data, err := p.ToUint32Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case INT64ARR:
		data, err := p.ToInt64Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case UINT64ARR:
		data, err := p.ToUint64Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case FLOAT32ARR:
		data, err := p.ToFloat32Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	case FLOAT64ARR:
		data, err := p.ToFloat64Arr(b)
		wrapper := make([]interface{}, len(data))
		for i := range wrapper {
			wrapper[i] = data[i]
		}
		return wrapper, err
	default:
		return nil, errors.New("Not Supported Data Type")
	}
}
