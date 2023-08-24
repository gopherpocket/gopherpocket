package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssembler(t *testing.T) {
	code, err := Assemble(
		NOP(),
		LD(BC, Imm16(0xabcd)),
		LD(Ptr(BC), A),
		LD(B, Imm8(0xFF)),
		LD(Ptr(HL), Imm8(0xDE)),
		LD(A, Ptr(HL, Plus)),
		LD(L, Imm8(0xBA)),
		LD(D, L),
		LD(H, Ptr(HL)),
		LD(Ptr(HL), B),
		LD(Ptr(C), A),
		LD(Ptr(Imm16(0xFACE)), A),
		AddLabel("foo",
			NOP()),
		JR(Label("foo")),
	)
	assert.NoError(t, err)
	assert.Equal(t, []byte{
		0x00,             // NOP
		0x01, 0xcd, 0xab, // LD BC, 0xabcd
		0x02,       // LD [BC], A
		0x06, 0xff, // LD  B, $FF
		0x36, 0xde, // LD [HL], $DE
		0x2a,       // LD A, [HL+]
		0x2e, 0xba, // LD L, $BA
		0x55,             // LD D, L
		0x66,             // LD H, [HL]
		0x70,             // LD [HL], B
		0xE2,             // LD [C], A
		0xEA, 0xce, 0xfa, // LD [$FACE], A
		0x00,                  // NOP
		0x18, byte(SImm8(-2)), // JR -2
	}, code)
}
