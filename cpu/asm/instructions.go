package asm

const (
	nop = "NOP"
	ld  = "LD"
)

// NOP assembles a NOP instruction, that performs no effect, and takes 4 cycles.
func NOP() *Instruction {
	return &Instruction{
		Mnemonic: nop,
		Bytes:    1,
		Cycles:   4,
		Operands: nil,
	}
}

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
		panic("invalid construction of LD")
	}

	return &Instruction{
		Mnemonic: ld,
		Bytes:    bytes,
		Cycles:   cycles,
		Operands: []Operand{lh, rh},
	}
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

// helper to check if an operand is of a specific concrete type
func is[T any, Op Operand](op Op) bool {
	_, ok := any(op).(T)
	return ok
}
