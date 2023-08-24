package asm

import (
	"bytes"
	"fmt"
)

type Addr interface {
	Offset8 | Label
	Operand
}

func JR[T Addr](e8 T) *Instruction {
	return &Instruction{
		Mnemonic: jr,
		Bytes:    2,
		Cycles:   12,
		Operands: []Operand{
			Operand(e8),
		},
	}
}

func JRZ[T Addr](e8 T) *Instruction {
	return &Instruction{
		Mnemonic:  jr,
		Bytes:     2,
		Cycles:    12,
		CyclesAlt: 8,
		Operands: []Operand{
			ZF,
			Operand(e8),
		},
	}
}

func JRNZ[T Addr](e8 T) *Instruction {
	return &Instruction{
		Mnemonic:  jr,
		Bytes:     2,
		Cycles:    12,
		CyclesAlt: 8,
		Operands: []Operand{
			NZF,
			Operand(e8),
		},
	}
}

func JRC[T Addr](e8 T) *Instruction {
	return &Instruction{
		Mnemonic:  jr,
		Bytes:     2,
		Cycles:    12,
		CyclesAlt: 8,
		Operands: []Operand{
			CF,
			Operand(e8),
		},
	}
}

func JRNC[T Addr](e8 T) *Instruction {
	return &Instruction{
		Mnemonic:  jr,
		Bytes:     2,
		Cycles:    12,
		CyclesAlt: 8,
		Operands: []Operand{
			NCF,
			Operand(e8),
		},
	}
}

func (a *Assembler) jr(i int, instr *Instruction, labels map[Label]int, buf *bytes.Buffer) error {
	badInstr := func(reason any) error {
		return badInstr(i, instr, reason)
	}

	illegalOperands := func() error {
		return illegalOperands(i, instr)
	}

	if len(instr.Operands) == 1 {
		var off Offset8
		switch v := instr.Operands[0].(type) {
		case Offset8:
			off = v

		case Label:
			label := v
			// lookup label
			labelIdx, ok := labels[label]
			if !ok {
				return fmt.Errorf("unknown label %q", label)
			}
			curPos := buf.Len()
			delta := labelIdx - (curPos + instr.Bytes)
			if delta > 127 || delta < -128 {
				return fmt.Errorf("label %q out of range (%d)", label, delta)
			}
			off = Offset8(delta)

		default:
			return illegalOperands()
		}
		// unconditional jump
		buf.WriteByte(0x18)
		buf.WriteByte(byte(off))
		return nil
	}

	if len(instr.Operands) == 2 {
		lh, rh := instr.Operands[0], instr.Operands[1]
		if !is[ConditionCode](lh) || !is[Offset8](rh) {
			return illegalOperands()
		}

		cc := lh.(ConditionCode)
		off := rh.(Offset8)

		switch cc {
		case ZF:
			buf.WriteByte(0x28)

		case NZF:
			buf.WriteByte(0x20)

		case CF:
			buf.WriteByte(0x38)

		case NCF:
			buf.WriteByte(0x30)

		default:
			return illegalOperands()
		}

		buf.WriteByte(byte(off))

		return nil
	}

	return badInstr("unexpected number of operands")
}
