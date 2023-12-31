package asm

import "reflect"

const (
	unknown = "<unknown>"
	nop     = "NOP"
	ld      = "LD"
	inc     = "INC"
	dec     = "DEC"
)

// helper to check if an operand is of a specific concrete type
func is[T any, Op Operand](op Op) bool {
	_, ok := any(op).(T)
	return ok
}

func isEq(l Operand, r Operand) bool {
	return reflect.DeepEqual(l, r)
}
