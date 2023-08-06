package asm

import "fmt"

func ExampleLD() {
	print := func(i *Instruction) {
		fmt.Println(i, "| Bytes:", i.Bytes, "Cycles:", i.Cycles)
	}

	print(LD(BC, Imm16(0xFFFF)))
	print(LD(Ptr(BC), A))
	print(LD(B, SImm8(-128)))
	print(LD(Ptr(Imm16(0xFF00)), SP))
	print(LD(C, Imm8(0xFF)))
	print(LD(DE, Imm16(0xDEAD)))
	print(LD(Ptr(DE), A))
	print(LD(D, Imm8(0xFF)))
	print(LD(A, Ptr(DE)))
	print(LD(E, Imm8(0x4)))
	print(LD(HL, Imm16(0xFACE)))
	print(LD(Ptr(HL, Plus), A))
	print(LD(H, Imm8(0xF)))
	print(LD(A, Ptr(HL, Plus)))
	print(LD(L, Imm8(0xD)))
	print(LD(Ptr(HL, Minus), A))
	print(LD(Ptr(HL), Imm8(0xC)))
	print(LD(B, B))
	print(LD(B, Ptr(HL)))
	print(LD(Ptr(Imm16(0xFF00)), A))
	print(LD(HL, SP+8))
	print(LD(SP, HL))

	// Output:
	//
	// LD BC, $FFFF | Bytes: 3 Cycles: 12
	// LD [BC], A | Bytes: 1 Cycles: 8
	// LD B, $FF | Bytes: 2 Cycles: 8
	// LD [$FF00], SP | Bytes: 3 Cycles: 20
	// LD C, $FF | Bytes: 2 Cycles: 8
	// LD DE, $DEAD | Bytes: 3 Cycles: 12
	// LD [DE], A | Bytes: 1 Cycles: 8
	// LD D, $FF | Bytes: 2 Cycles: 8
	// LD A, [DE] | Bytes: 1 Cycles: 8
	// LD E, $4 | Bytes: 2 Cycles: 8
	// LD HL, $FACE | Bytes: 3 Cycles: 12
	// LD [HL+], A | Bytes: 1 Cycles: 8
	// LD H, $F | Bytes: 2 Cycles: 8
	// LD A, [HL+] | Bytes: 1 Cycles: 8
	// LD L, $D | Bytes: 2 Cycles: 8
	// LD [HL-], A | Bytes: 1 Cycles: 8
	// LD [HL], $C | Bytes: 2 Cycles: 12
	// LD B, B | Bytes: 1 Cycles: 4
	// LD B, [HL] | Bytes: 1 Cycles: 8
	// LD [$FF00], A | Bytes: 3 Cycles: 16
	// LD HL, SP + 8 | Bytes: 1 Cycles: 8
	// LD SP, HL | Bytes: 1 Cycles: 8
}
