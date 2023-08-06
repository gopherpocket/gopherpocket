package asm

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// LD assembles a [LD instruction]: https://rgbds.gbdev.io/docs/v0.6.1/gbz80.7#Load_Instructions
func LD[OpL, OpR Operand](lh OpL, rh OpR) *Instruction {
	var bytes, cycles int
	switch {
	case is[Reg8](lh) && is[Reg8](rh):
		bytes = 1
		cycles = 4

	case is[Reg8](lh) && is[Imm8](rh):
		bytes = 2
		cycles = 8

	case is[Reg8](lh) && is[Pointer[Reg16]](rh):
		bytes = 1
		cycles = 8

	case is[Reg16](lh) && is[Reg16](rh):
		bytes = 1
		cycles = 8

	case is[Reg16](lh) && is[Immediate16](rh):
		bytes = 3
		cycles = 12

	case is[Pointer[Reg8]](lh) && is[Reg8](rh):
		bytes = 1
		cycles = 8

	case is[Pointer[Reg16]](lh) && is[Reg8](rh):
		bytes = 1
		cycles = 8

	case is[Pointer[Reg16]](lh) && is[Imm8](rh):
		bytes = 2
		cycles = 12

	case is[Pointer[Imm16]](lh) && is[Reg8](rh):
		bytes = 3
		cycles = 16

	case is[Pointer[Imm16]](lh) && is[Reg16](rh):
		bytes = 3
		cycles = 20

	case isLDHLSPOffset(lh, rh):
		bytes = 2
		cycles = 12

	default:
		return &Instruction{
			Mnemonic: ld,
			err:      errors.New("invalid construction"),
		}
	}

	return &Instruction{
		Mnemonic: ld,
		Bytes:    bytes,
		Cycles:   cycles,
		Operands: []Operand{lh, rh},
	}
}

func (a *Assembler) ld(i int, instr *Instruction, buf *bytes.Buffer) error {

	badInstr := func(reason any) error {
		return badInstr(i, instr, reason)
	}

	illegalOperands := func() error {
		return illegalOperands(i, instr)
	}

	if len(instr.Operands) != 2 {
		return badInstr("unexpected number of operands")
	}

	lh, rh := instr.Operands[0], instr.Operands[1]

	switch {
	case instr.Bytes == 3 &&
		is[Reg16](lh) &&
		is[Imm16](rh):
		switch lh.(Reg16) {
		case BC:
			buf.WriteByte(0x01)

		case DE:
			buf.WriteByte(0x11)

		case HL:
			buf.WriteByte(0x21)

		case SP:
			buf.WriteByte(0x31)

		default:
			return illegalOperands()
		}
		binary.Write(buf, binary.LittleEndian, rh.(Imm16))

	case instr.Bytes == 1 &&
		is[Pointer[Reg16]](lh) &&
		is[Reg8](rh) &&
		rh.(Reg8) == A:
		ptr := lh.(Pointer[Reg16])
		switch {
		case ptr.Ref == BC && ptr.Delta == None:
			buf.WriteByte(0x02)

		case ptr.Ref == DE && ptr.Delta == None:
			buf.WriteByte(0x12)

		case ptr.Ref == HL && ptr.Delta == Plus:
			buf.WriteByte(0x22)

		case ptr.Ref == HL && ptr.Delta == Minus:
			buf.WriteByte(0x32)

		default:
			return illegalOperands()
		}

	case instr.Bytes == 2 &&
		is[Reg8](lh) &&
		is[Imm8](rh):
		switch lh.(Reg8) {
		case B:
			buf.WriteByte(0x06)

		case D:
			buf.WriteByte(0x16)

		case H:
			buf.WriteByte(0x26)

		case C:
			buf.WriteByte(0x0E)

		case E:
			buf.WriteByte(0x1E)

		case L:
			buf.WriteByte(0x2E)

		case A:
			buf.WriteByte(0x3E)

		default:
			return illegalOperands()
		}
		buf.WriteByte(byte(rh.(Imm8)))

	case instr.Bytes == 2 &&
		is[Pointer[Reg16]](lh) &&
		is[Imm8](rh) &&
		lh.(Pointer[Reg16]).Ref == HL &&
		lh.(Pointer[Reg16]).Delta == None:
		buf.WriteByte(0x36)
		buf.WriteByte(byte(rh.(Imm8)))

	case instr.Bytes == 1 &&
		is[Reg8](lh) &&
		lh.(Reg8) == A &&
		is[Pointer[Reg16]](rh):
		ptr := rh.(Pointer[Reg16])
		switch {
		case ptr.Ref == BC && ptr.Delta == None:
			buf.WriteByte(0x0A)

		case ptr.Ref == DE && ptr.Delta == None:
			buf.WriteByte(0x1A)

		case ptr.Ref == HL && ptr.Delta == Plus:
			buf.WriteByte(0x2A)

		case ptr.Ref == HL && ptr.Delta == Minus:
			buf.WriteByte(0x3A)

		default:
			return illegalOperands()
		}

	case instr.Bytes == 1 &&
		is[Reg8](lh) &&
		is[Reg8](rh):
		dst, src := lh.(Reg8), rh.(Reg8)

		switch {
		case dst == B && src == B:
			buf.WriteByte(0x40)

		case dst == D && src == B:
			buf.WriteByte(0x50)

		case dst == H && src == B:
			buf.WriteByte(0x60)

		case dst == B && src == C:
			buf.WriteByte(0x41)

		case dst == D && src == C:
			buf.WriteByte(0x51)

		case dst == H && src == C:
			buf.WriteByte(0x61)

		case dst == B && src == D:
			buf.WriteByte(0x42)

		case dst == D && src == D:
			buf.WriteByte(0x52)

		case dst == H && src == D:
			buf.WriteByte(0x62)

		case dst == B && src == E:
			buf.WriteByte(0x43)

		case dst == D && src == E:
			buf.WriteByte(0x53)

		case dst == H && src == E:
			buf.WriteByte(0x63)

		case dst == B && src == H:
			buf.WriteByte(0x44)

		case dst == D && src == H:
			buf.WriteByte(0x54)

		case dst == H && src == H:
			buf.WriteByte(0x64)

		case dst == B && src == L:
			buf.WriteByte(0x45)

		case dst == D && src == L:
			buf.WriteByte(0x55)

		case dst == H && src == L:
			buf.WriteByte(0x65)

		case dst == B && src == A:
			buf.WriteByte(0x47)

		case dst == D && src == A:
			buf.WriteByte(0x57)

		case dst == H && src == A:
			buf.WriteByte(0x67)

		case dst == C && src == B:
			buf.WriteByte(0x48)

		case dst == E && src == B:
			buf.WriteByte(0x58)

		case dst == L && src == B:
			buf.WriteByte(0x68)

		case dst == C && src == C:
			buf.WriteByte(0x49)

		case dst == E && src == C:
			buf.WriteByte(0x59)

		case dst == L && src == C:
			buf.WriteByte(0x69)

		case dst == C && src == D:
			buf.WriteByte(0x4A)

		case dst == E && src == D:
			buf.WriteByte(0x5A)

		case dst == L && src == D:
			buf.WriteByte(0x6A)

		case dst == C && src == E:
			buf.WriteByte(0x4B)

		case dst == E && src == E:
			buf.WriteByte(0x5B)

		case dst == L && src == E:
			buf.WriteByte(0x6B)

		case dst == C && src == H:
			buf.WriteByte(0x4C)

		case dst == E && src == H:
			buf.WriteByte(0x5C)

		case dst == L && src == H:
			buf.WriteByte(0x6C)

		case dst == C && src == L:
			buf.WriteByte(0x4D)

		case dst == E && src == L:
			buf.WriteByte(0x5D)

		case dst == L && src == L:
			buf.WriteByte(0x6D)

		case dst == C && src == A:
			buf.WriteByte(0x4F)

		case dst == E && src == A:
			buf.WriteByte(0x5F)

		case dst == L && src == A:
			buf.WriteByte(0x6F)

		case dst == A && src == B:
			buf.WriteByte(0x78)

		case dst == A && src == C:
			buf.WriteByte(0x79)

		case dst == A && src == D:
			buf.WriteByte(0x7A)

		case dst == A && src == E:
			buf.WriteByte(0x7B)

		case dst == A && src == H:
			buf.WriteByte(0x7C)

		case dst == A && src == L:
			buf.WriteByte(0x7D)

		case dst == A && src == A:
			buf.WriteByte(0x7F)

		default:
			return illegalOperands()
		}

	case instr.Bytes == 1 &&
		is[Reg8](lh) &&
		is[Pointer[Reg16]](rh) &&
		rh.(Pointer[Reg16]).Ref == HL &&
		rh.(Pointer[Reg16]).Delta == None:

		switch lh.(Reg8) {
		case B:
			buf.WriteByte(0x46)

		case C:
			buf.WriteByte(0x4E)

		case D:
			buf.WriteByte(0x56)

		case E:
			buf.WriteByte(0x5E)

		case H:
			buf.WriteByte(0x66)

		case L:
			buf.WriteByte(0x6E)

		case A:
			buf.WriteByte(0x7E)

		default:
			return illegalOperands()
		}

	case instr.Bytes == 1 &&
		is[Pointer[Reg16]](lh) &&
		is[Reg8](rh) &&
		lh.(Pointer[Reg16]).Ref == HL &&
		lh.(Pointer[Reg16]).Delta == None:
		switch rh.(Reg8) {
		case B:
			buf.WriteByte(0x70)

		case C:
			buf.WriteByte(0x71)

		case D:
			buf.WriteByte(0x72)

		case E:
			buf.WriteByte(0x73)

		case H:
			buf.WriteByte(0x74)

		case L:
			buf.WriteByte(0x75)

		case A:
			buf.WriteByte(0x77)

		default:
			return illegalOperands()
		}

	case instr.Bytes == 1 &&
		isEq(lh, Ptr(C)) &&
		isEq(rh, A):
		buf.WriteByte(0xE2)

	case instr.Bytes == 1 &&
		isEq(lh, A) &&
		isEq(rh, Ptr(C)):
		buf.WriteByte(0xF2)

	case instr.Bytes == 3 &&
		is[Pointer[Imm16]](lh) &&
		isEq(rh, A):
		buf.WriteByte(0xEA)
		binary.Write(buf, binary.LittleEndian, lh.(Pointer[Imm16]).Ref)

	case instr.Bytes == 3 &&
		isEq(lh, A) &&
		is[Pointer[Imm16]](rh):
		buf.WriteByte(0xFA)
		binary.Write(buf, binary.LittleEndian, rh.(Pointer[Imm16]).Ref)

	default:
		return illegalOperands()
	}
	return nil
}

// helper to detect if an instruction is LD HL, SP + e8
func isLDHLSPOffset(lh, rh Operand) bool {
	if !is[Register16](lh) || !is[Register16](rh) {
		return false
	}
	l := lh.(Register16)
	r := rh.(Register16)

	return l == HL && r >= spMax && r <= spMax
}
