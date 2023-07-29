package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	assert := assert.New(t)
	var r Register

	assert.Equal(r.Hi(), uint8(0))
	assert.Equal(r.Lo(), uint8(0))

	r.SetHi(0xff)
	assert.Equal(r.Hi(), uint8(0xff))
	assert.Equal(r.Lo(), uint8(0x00))
	assert.Equal(r, Register(0xff00))

	r.SetLo(0xff)
	assert.Equal(r.Hi(), uint8(0xff))
	assert.Equal(r.Lo(), uint8(0xff))
	assert.Equal(r, Register(0xffff))
}
