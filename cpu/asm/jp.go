package asm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

func JP[T Addr | Register16](a16 T) *Instruction {
	if reg, ok := any(a16).(Register16); ok {
		if reg != HL {
			return &Instruction{
				err: errors.New("JP r16 can only be JP HL"),
			}
		}
		return &Instruction{
			Mnemonic: jp,
			Bytes:    1,
			Cycles:   4,
			Operands: []Operand{
				HL,
			},
		}
	}

	return &Instruction{
		Mnemonic: jp,
		Bytes:    3,
		Cycles:   16,
		Operands: []Operand{
			Operand(a16),
		},
	}
}

func JPZ[T Addr](a16 T) *Instruction {
	return &Instruction{
		Mnemonic:  jp,
		Bytes:     3,
		Cycles:    16,
		CyclesAlt: 12,
		Operands: []Operand{
			ZF,
			Operand(a16),
		},
	}
}

func JPNZ[T Addr](a16 T) *Instruction {
	return &Instruction{
		Mnemonic:  jp,
		Bytes:     3,
		Cycles:    16,
		CyclesAlt: 12,
		Operands: []Operand{
			NZF,
			Operand(a16),
		},
	}
}

func JPC[T Addr](a16 T) *Instruction {
	return &Instruction{
		Mnemonic:  jp,
		Bytes:     3,
		Cycles:    16,
		CyclesAlt: 12,
		Operands: []Operand{
			CF,
			Operand(a16),
		},
	}
}

func JPNC[T Addr](a16 T) *Instruction {
	return &Instruction{
		Mnemonic:  jp,
		Bytes:     3,
		Cycles:    16,
		CyclesAlt: 12,
		Operands: []Operand{
			NCF,
			Operand(a16),
		},
	}
}

func (a *Assembler) jp(i int, instr *Instruction, labels map[Label]int, buf *bytes.Buffer) error {
	badInstr := func(reason any) error {
		return badInstr(i, instr, reason)
	}

	illegalOperands := func() error {
		return illegalOperands(i, instr)
	}

	if len(instr.Operands) == 1 {
		var addr Imm16
		switch v := instr.Operands[0].(type) {
		case Register16:
			buf.WriteByte(0xE9)
			return nil

		case Imm16:
			addr = v

		case Label:
			label := v
			// lookup label
			labelIdx, ok := labels[label]
			if !ok {
				return fmt.Errorf("unknown label %q", label)
			}
			if labelIdx < 0 || labelIdx > 0xFFFF {
				return fmt.Errorf("label %q out of range (%d)", label, labelIdx)
			}
			addr = Imm16(addr)

		default:
			return illegalOperands()
		}
		// unconditional jump
		buf.WriteByte(0xc3)
		binary.Write(buf, binary.LittleEndian, uint16(addr))
		return nil
	}

	if len(instr.Operands) == 2 {
		lh, rh := instr.Operands[0], instr.Operands[1]
		if !is[ConditionCode](lh) || !(is[Offset8](rh) || is[Label](rh)) {
			return illegalOperands()
		}

		cc := lh.(ConditionCode)
		var addr Imm16

		switch v := rh.(type) {
		case Imm16:
			addr = v

		case Label:
			label := v
			// lookup label
			labelIdx, ok := labels[label]
			if !ok {
				return fmt.Errorf("unknown label %q", label)
			}
			if labelIdx < 0 || labelIdx > 0xFFFF {
				return fmt.Errorf("label %q out of range (%d)", label, labelIdx)
			}
			addr = Imm16(addr)

		default:
			return illegalOperands()
		}

		switch cc {
		case ZF:
			buf.WriteByte(0xCA)

		case NZF:
			buf.WriteByte(0xC2)

		case CF:
			buf.WriteByte(0xDA)

		case NCF:
			buf.WriteByte(0xD2)

		default:
			return illegalOperands()
		}

		binary.Write(buf, binary.LittleEndian, uint16(addr))

		return nil
	}

	return badInstr("unexpected number of operands")
}
