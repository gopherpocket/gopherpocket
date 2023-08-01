// Package asm implements Gameboy CPU assembly
package asm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

// Assembler is a Gameboy Z-80-like assembler.
type Assembler struct {
}

// NewAssembler constructs a new [Assembler] object.
func NewAssembler() *Assembler {
	return &Assembler{}
}

// Assemble accepts a stream of instructions and assembles them into binary bytes.
func (a *Assembler) Assemble(instrs ...*Instruction) ([]byte, error) {
	var buffer bytes.Buffer

	for i, instr := range instrs {
		if err := instr.err; err != nil {
			return nil, fmt.Errorf("bad instruction: line %d: %v", i, err)
		}
		switch instr.Mnemonic {
		case nop:
			buffer.WriteByte(0x00)

		case ld:
			switch {
			case instr.Bytes == 3 &&
				len(instr.Operands) == 2 &&
				is[Reg16](instr.Operands[0]) &&
				is[Imm16](instr.Operands[1]):
				switch instr.Operands[0].(Reg16) {
				case BC:
					buffer.WriteByte(0x01)

				case DE:
					buffer.WriteByte(0x11)

				case HL:
					buffer.WriteByte(0x21)

				case SP:
					buffer.WriteByte(0x31)

				default:
					return nil, fmt.Errorf("bad instruction: line %d: %q: unknown register-16", i, instr.String())

				}
				binary.Write(&buffer, binary.LittleEndian, instr.Operands[1].(Imm16))
			}
		default:
			return nil, fmt.Errorf("bad instruction: line %d: %q", i, instr.String())
		}
	}

	return buffer.Bytes(), nil
}

func Assemble(instrs ...*Instruction) ([]byte, error) {
	assm := NewAssembler()
	return assm.Assemble(instrs...)
}

// Instruction is a in intermediate representation of a Gameboy CPU instruction.
type Instruction struct {
	Mnemonic string
	Bytes    int
	Cycles   int
	Operands []Operand
	err      error
}

// Err returns the current error assosciated with the Instruction, if it exists.
func (i *Instruction) Err() error {
	return fmt.Errorf("%s: %v", i.Mnemonic, i.err)
}

// String implements fmt.Stringer
func (i *Instruction) String() string {
	var opStr []string
	for _, op := range i.Operands {
		opStr = append(opStr, op.String())
	}
	var builder strings.Builder
	builder.WriteString(i.Mnemonic)
	if len(opStr) > 0 {
		builder.WriteByte(' ')
		builder.WriteString(strings.Join(opStr, ", "))
	}

	return builder.String()
}

// Operand is One of: [Register, Immediate, Pointer]
type Operand interface {
	operand()
	fmt.Stringer
}

type (
	Register8   int
	Reg8        = Register8
	Register16  int
	Reg16       = Register16
	Immediate8  uint8
	Imm8        = Immediate8
	Immediate16 uint16
	Imm16       = Immediate16
)

// SImm8 is a helper that constructs an Imm8 from a signed 8 bit integer.
// This is helpful since Go will not automatically convert Signed Int8 to Unsigned Int8 from constants.
func SImm8(x int8) Imm8 {
	return signedImm[int8, Imm8](x)
}

// SImm16 is a helper that constructs an Imm16 from a signed 16 bit integer.
// This is helpful since Go will not automatically convert Signed Int8 to Unsigned Int8 from constants.
func SImm16(x int16) Imm16 {
	return signedImm[int16, Imm16](x)
}

func signedImm[T int8 | int16, R Imm8 | Imm16](x T) R {
	if x < 0 {
		return (R(x) - 1) | 0x80
	}
	return R(x)
}

// Ref is a generic type constraint for the reference a Pointer Operand can hold.
// Any Pointer can be a reference to another Operand that is either a:
// * Register8/16
// * Immediate8/16
type Ref interface {
	Register8 | Register16 | Immediate8 | Immediate16
	Operand
}

func (Register8) operand() {}

// String implements fmt.Stringer
func (r Register8) String() string {
	switch {
	case r >= A && r <= L:
		return registerStrs[r]

	default:
		return "<invalid register>"
	}
}

func (Register16) operand() {}

// String implements fmt.Stringer
func (r Register16) String() string {
	switch {
	case r >= AF && r <= PC:
		return compoundStrs[r]

	case r == SP:
		return "SP"

	case r >= spMin && r <= spMax:
		delta := int(r - SP)
		if delta > 0 {
			return "SP + " + strconv.Itoa(delta)
		} else {
			return "SP - " + strconv.Itoa(int(r-SP))
		}

	default:
		return "<invalid register>"
	}
}

func (Immediate8) operand() {}

// String implements fmt.Stringer
func (i Immediate8) String() string {
	return fmt.Sprintf("$%X", uint8(i))
}

func (Immediate16) operand() {}

// String implements fmt.Stringer
func (i Immediate16) String() string {
	return fmt.Sprintf("$%X", uint16(i))
}

// Pointer reoresents a Pointer to a Reference Operand.
// In some operations, The gameoby CPU can increment or decrement Pointer value, which can be
// represented by the Delta field.
type Pointer[R Ref] struct {
	Ref   R
	Delta Delta
}

// Delta represents whether a Register Pointer should be incremented or Decremented on use.
// This is to represent instructions like LD [HL+], A
type Delta int

const (
	Plus  Delta = -1
	Minus Delta = 1
)

// String implements fmt.Stringer
func (d Delta) String() string {
	switch d {
	case Minus:
		return "+"

	case Plus:
		return "-"

	default:
		return ""
	}
}

func (Pointer[_]) operand() {}

// String implements fmt.Stringer
func (p Pointer[_]) String() string {
	return "[" + p.Ref.String() + p.Delta.String() + "]"
}

// Ptr is a helper function to construct a generic Pointer Ref, with an Optional Delta provided in the form of
// a variadic argument list.
func Ptr[R Ref](r R, delta ...Delta) Pointer[R] {
	var d Delta
	if len(delta) > 0 {
		d = delta[0]
	}

	return Pointer[R]{
		Ref:   r,
		Delta: d,
	}
}

// 8 Bit Registers
const (
	A Register8 = iota
	F
	B
	C
	D
	E
	H
	L
)

// 16 Bit Registers
const (
	AF Register16 = iota
	BC
	DE
	HL
	PC

	// We use this special enumeration value to let us adding a signed 8 bit integer to SP as a right-hand opeand
	// This lets us express Instructions like LD HL, SP + 0x7F
	spMin            = SP - 0x80
	SP    Register16 = 0x85
	spMax            = SP + 0x7F
)

var registerStrs = []string{"A", "F", "B", "C", "D", "E", "H", "L"}
var compoundStrs = []string{"AF", "BC", "DE", "HL", "PC"}
