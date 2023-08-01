package asm

// NOP assembles a NOP instruction, that performs no effect, and takes 4 cycles.
func NOP() *Instruction {
	return &Instruction{
		Mnemonic: nop,
		Bytes:    1,
		Cycles:   4,
		Operands: nil,
	}
}
