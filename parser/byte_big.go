package parser

import (
	"encoding/binary"
	"errors"
	"math"
)

type BigByteParser struct {
}

func (p *BigByteParser) ToBit(datum byte) []uint8 {
	var bits []uint8 = make([]uint8, 8)
	for i := range bits {
		bits[i] = datum & 1 // and
		datum >>= 1         // shift bit
	}

	return bits
}

func (p *BigByteParser) ToBitArr(datum []byte) []uint8 {
	var bits []uint8 = make([]uint8, len(datum)*8)
	for i, data := range datum {
		res := p.ToBit(data)
		for j, bit := range res {
			bits[j+(8*i)] = bit
		}
	}

	return bits
}

func (p *BigByteParser) ToInt16(b []byte) int16 {
	return int16(binary.BigEndian.Uint16(b))
}

func (p *BigByteParser) ToInt16Arr(b []byte) ([]int16, error) {
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

func (p *BigByteParser) ToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

func (p *BigByteParser) ToUint16Arr(b []byte) ([]uint16, error) {
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

func (p *BigByteParser) ToInt32(b []byte) int32 {
	return int32(binary.BigEndian.Uint32(b))
}

func (p *BigByteParser) ToInt32Arr(b []byte) ([]int32, error) {
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

func (p *BigByteParser) ToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func (p *BigByteParser) ToUint32Arr(b []byte) ([]uint32, error) {
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

func (p *BigByteParser) ToInt64(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(b))
}

func (p *BigByteParser) ToInt64Arr(b []byte) ([]int64, error) {
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

func (p *BigByteParser) ToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func (p *BigByteParser) ToUint64Arr(b []byte) ([]uint64, error) {
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

func (p *BigByteParser) ToFloat32(b []byte) float32 {
	datum := binary.BigEndian.Uint32(b)
	return math.Float32frombits(datum)
}

func (p *BigByteParser) ToFloat32Arr(b []byte) ([]float32, error) {
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

func (p *BigByteParser) ToFloat64(b []byte) float64 {
	datum := binary.BigEndian.Uint64(b)
	return math.Float64frombits(datum)
}

func (p *BigByteParser) ToFloat64Arr(b []byte) ([]float64, error) {
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

func (p *BigByteParser) ToInterface(dataType uint8, b []byte) ([]interface{}, error) {
	var wrapper []interface{}
	switch dataType {
	case INT16:
		return append(wrapper, p.ToInt16(b)), nil
	case UINT16:
		return append(wrapper, p.ToUint16(b)), nil
	case INT32:
		return append(wrapper, p.ToInt32(b)), nil
	case UINT32:
		return append(wrapper, p.ToUint32(b)), nil
	case INT64:
		return append(wrapper, p.ToInt64(b)), nil
	case UINT64:
		return append(wrapper, p.ToUint64(b)), nil
	case FLOAT32:
		return append(wrapper, p.ToFloat32(b)), nil
	case FLOAT64:
		return append(wrapper, p.ToFloat64(b)), nil
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
