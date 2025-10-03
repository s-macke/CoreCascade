package filebuffer

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

type WriteBuffer struct {
	b []byte
}

func NewWriteBuffer(guessedCapacity int) *WriteBuffer {
	return &WriteBuffer{
		b: make([]byte, 0, guessedCapacity),
	}
}

func (wb *WriteBuffer) GetBuffer() []byte {
	return wb.b
}

func (wb *WriteBuffer) Length() int {
	return len(wb.b)
}

func (wb *WriteBuffer) WriteByte(b byte) {
	wb.b = append(wb.b, b)
}

func (wb *WriteBuffer) WriteWord(w uint16) {
	wb.b = append(wb.b, byte(w&0xFF), byte((w>>8)&0xFF))
}

func (wb *WriteBuffer) WriteInt32(v int32) {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], uint32(v))
	wb.WriteSlice(buf[:])
}

func (wb *WriteBuffer) WriteString(s string) {
	for _, b := range []byte(s) {
		wb.b = append(wb.b, b)
	}
}

func (wb *WriteBuffer) WriteFloat64(v float64) {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(v))
	wb.WriteSlice(buf[:])
}

func (wb *WriteBuffer) WriteFloat32(v float32) {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], math.Float32bits(v))
	wb.WriteSlice(buf[:])
}

func (wb *WriteBuffer) WriteZeroes(sLen int) {
	wb.b = append(wb.b, make([]byte, sLen)...)
}

func (wb *WriteBuffer) WriteSlice(s []byte) {
	wb.b = append(wb.b, s...)
}

func (wb *WriteBuffer) StoreToFile(filename string) {
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}
	defer outFile.Close()
	_, err = outFile.Write(wb.b)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Data saved to", filename)
	}
}
