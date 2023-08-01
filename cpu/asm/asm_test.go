package asm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssembler(t *testing.T) {
	code, err := Assemble(
		NOP(),
		NOP(),
		LD(BC, Imm16(0xabcd)),
	)
	assert.NoError(t, err)
	assert.Equal(t, code, []byte{0x00, 0x00, 0x01, 0xcd, 0xab})
}
