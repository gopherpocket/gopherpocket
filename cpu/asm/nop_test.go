package asm

import "fmt"

func ExampleNOP() {
	instr := NOP()

	fmt.Println(instr)
	// Output: NOP
}
