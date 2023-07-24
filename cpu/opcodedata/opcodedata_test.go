package opcodedata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert.NotZero(t, OpcodeData)
}
