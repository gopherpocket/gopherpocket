// Package opcodedata provides JSON based opcode data
package opcodedata

// OpcodeData exports the root JSON document as a variable
var OpcodeData Document

type Document struct {
	Unprefixed InstructionMap `json:"unprefixed"`
	CBPrefixed InstructionMap `json:"cbprefixed"`
}

type InstructionMap map[Opcode]*InstructionInfo

type Opcode = string

type InstructionInfo struct {
	Mnemonic  string     `json:"mnemonic"`
	Bytes     int        `json:"bytes"`
	Cycles    []int      `json:"cycles"`
	Operands  []*Operand `json:"operands"`
	Immediate bool       `json:"immediate"`
	Flags     Flags      `json:"flags"`
}

type Operand struct {
	Name      string `json:"name"`
	Bytes     int    `json:"bytes"`
	Immediate bool   `json:"immediate"`
	Increment bool   `json:"increment"`
	Decrement bool   `json:"decrement"`
}

type Flags struct {
	Z string
	N string
	H string
	C string
}
