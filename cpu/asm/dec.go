package asm

import "errors"

func DEC[Op Operand](op Op) *Instruction {
	var bytes, cycles int
	switch {
	case isEq(B, op):
		bytes = 1
		cycles = 4
	case isEq(D, op):
		bytes = 1
		cycles = 4
	case isEq(H, op):
		bytes = 1
		cycles = 4
	case isEq(Ptr(HL), op):
		bytes = 1
		cycles = 12
	case isEq(BC, op):
		bytes = 1
		cycles = 8
	case isEq(DE, op):
		bytes = 1
		cycles = 8
	case isEq(HL, op):
		bytes = 1
		cycles = 8
	case isEq(SP, op):
		bytes = 1
		cycles = 8
	case isEq(C, op):
		bytes = 1
		cycles = 4
	case isEq(E, op):
		bytes = 1
		cycles = 4
	case isEq(L, op):
		bytes = 1
		cycles = 4
	case isEq(A, op):
		bytes = 1
		cycles = 4

	default:
		return &Instruction{
			Mnemonic: dec,
			err:      errors.New("invalid construction"),
		}
	}

	return &Instruction{
		Mnemonic: dec,
		Bytes:    bytes,
		Cycles:   cycles,
		Operands: []Operand{op},
	}

}
