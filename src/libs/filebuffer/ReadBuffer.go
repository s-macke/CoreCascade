package filebuffer

import (
	"encoding/binary"
	"math"
	"os"
)

type ReadBuffer struct {
	b      []byte
	offset int
}

func NewReadBuffer(b []byte) *ReadBuffer {
	return &ReadBuffer{
		b:      b,
		offset: 0,
	}
}

func NewReadBufferFromFile(filename string) *ReadBuffer {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return &ReadBuffer{
		b:      b,
		offset: 0,
	}
}

func (rb *ReadBuffer) Length() int {
	return len(rb.b)
}

func (rb *ReadBuffer) ReadNextByte() byte {
	if rb.offset >= len(rb.b) {
		return 0
	}
	b := rb.b[rb.offset]
	rb.offset++
	return b
}

func (rb *ReadBuffer) ReadSlice(n int) []byte {
	//fmt.Println("ReadSlice", n, rb.offset, len(rb.b))
	b := rb.b[rb.offset : rb.offset+n]
	rb.offset = rb.offset + n
	return b
}

func (rb *ReadBuffer) ReadSliceAsString(n int) string {
	bytes := rb.ReadSlice(n)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == 0 {
			return string(bytes[:i])
		}
	}
	return string(bytes)
}

func clen(n []byte) int {
	for i := 0; i < len(n); i++ {
		if n[i] == 0 {
			return i
		}
	}
	return len(n)
}
func (rb *ReadBuffer) ReadSliceAsNullTerminatedString(n int) string {
	bytes := rb.ReadSlice(n)
	length := clen(bytes)
	return string(bytes[:length])
}

func (rb *ReadBuffer) NewReadBuffer(n int) *ReadBuffer {
	b := rb.b[rb.offset : rb.offset+n]
	rb.offset = rb.offset + n
	return NewReadBuffer(b)
}

func (rb *ReadBuffer) NewReadBufferAt(offset int) *ReadBuffer {
	b := rb.b[offset:]
	return NewReadBuffer(b)
}

func (rb *ReadBuffer) SkipNBytes(n int) {
	rb.offset = rb.offset + n
}

func (rb *ReadBuffer) EOF() bool {
	return rb.offset >= len(rb.b)
}

func (rb *ReadBuffer) ReadInt(bytes int) int {
	var value = 0
	for i := 0; i < bytes; i++ {
		value |= int(rb.ReadNextByte()) << (8 * i)
	}
	return value
}

func (rb *ReadBuffer) ReadFloat64() float64 {
	if rb.offset+8 > len(rb.b) {
		panic("ReadFloat64: not enough bytes in buffer")
	}
	buf := rb.ReadSlice(8)
	vint := binary.LittleEndian.Uint64(buf[:8])
	value := math.Float64frombits(vint)
	return value
}

func (rb *ReadBuffer) ReadFloat32() float32 {
	if rb.offset+4 > len(rb.b) {
		panic("ReadFloat32: not enough bytes in buffer")
	}
	buf := rb.ReadSlice(4)
	vint := binary.LittleEndian.Uint32(buf[:4])
	value := math.Float32frombits(vint)
	return value
}
