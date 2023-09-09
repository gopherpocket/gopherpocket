package asm

import "errors"

// INC assembles a [INC instruction]: https://rgbds.gbdev.io/docs/v0.6.1/gbz80.7/#8-bit_Arithmetic_and_Logic_Instructions
func INC[Op Operand](op Op) *Instruction {
	var bytes, cycles int
	switch {
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
			Mnemonic: inc,
			err:      errors.New("invalid construction"),
		}
	}

	return &Instruction{
		Mnemonic: inc,
		Bytes:    bytes,
		Cycles:   cycles,
		Operands: []Operand{op},
	}

}
