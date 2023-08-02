package cpu

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Memory represents Gameboy Memory
type Memory struct {
	buffer [0x10000]byte
}

// NewMemory constructs a new Memory object
func NewMemory() *Memory {
	return &Memory{}
}

// WriteAt implements memoryImplements.
func (m *Memory) WriteAt(p []byte, off int64) (n int, err error) {
	start := off
	end := off + int64(len(p))
	if start < 0 || end > int64(len(m.buffer)) {
		return 0, fmt.Errorf("range [%d:%d] out of bounds", start, end)
	}
	copy(m.buffer[start:end], p)
	return len(p), nil
}

// WriteUint8At is a wrapper for writing a uint8 using [WriteAt]
func (m *Memory) WriteUint8At(v uint8, off int64) error {
	buf := [1]byte{v}
	n, err := m.WriteAt(buf[:], off)
	if n != len(buf) {
		panic("short/long write")
	}
	return err
}

// WriteUint16At is a wrapper for writing a uint16 using [WriteAt]
func (m *Memory) WriteUint16At(v uint16, off int64) error {
	buf := [2]byte{}
	binary.LittleEndian.PutUint16(buf[:], v)
	n, err := m.WriteAt(buf[:], off)
	if n != len(buf) {
		panic("short/long write")
	}
	return err
}

// ReadAt implements io.ReaderAt.
func (m *Memory) ReadAt(p []byte, off int64) (n int, err error) {
	start := off
	end := off + int64(len(p))
	if start < 0 || end > int64(len(m.buffer)) {
		return 0, fmt.Errorf("range [%d:%d] out of bounds", start, end)
	}
	copy(p, m.buffer[start:end])
	return len(p), nil
}

// ReadUint8At is a wrapper for reading a uint8 using [ReadAt]
func (m *Memory) ReadUint8At(off int64) (uint8, error) {
	var buf [1]byte
	n, err := m.ReadAt(buf[:], off)
	if err != nil {
		return 0, err
	}
	if n != len(buf) {
		panic("short/long read")
	}
	return buf[0], nil
}

// ReadUint16At is a wrapper for reading a uint16 using [ReadAt]
func (m *Memory) ReadUint16At(off int64) (uint16, error) {
	var buf [2]byte
	n, err := m.ReadAt(buf[:], off)
	if err != nil {
		return 0, err
	}
	if n != len(buf) {
		panic("short/long read")
	}
	return binary.LittleEndian.Uint16(buf[:]), nil
}

type memoryImplements interface {
	io.ReaderAt
	io.WriterAt
}

var _ memoryImplements = (*Memory)(nil)
