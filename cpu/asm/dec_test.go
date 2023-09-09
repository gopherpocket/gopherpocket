package asm

import "fmt"

func ExampleDEC() {
	print := func(i *Instruction) {
		fmt.Println(i, "| Bytes:", i.Bytes, "Cycles:", i.Cycles)
	}

	print(DEC(B))
	print(DEC(D))
	print(DEC(H))
	print(DEC(Ptr(HL)))
	print(DEC(BC))
	print(DEC(DE))
	print(DEC(HL))
	print(DEC(SP))
	print(DEC(C))
	print(DEC(E))
	print(DEC(L))
	print(DEC(A))

	//Output:
	//
	// DEC B | Bytes: 1 Cycles: 4
	// DEC D | Bytes: 1 Cycles: 4
	// DEC H | Bytes: 1 Cycles: 4
	// DEC [HL] | Bytes: 1 Cycles: 12
	// DEC BC | Bytes: 1 Cycles: 8
	// DEC DE | Bytes: 1 Cycles: 8
	// DEC HL | Bytes: 1 Cycles: 8
	// DEC SP | Bytes: 1 Cycles: 8
	// DEC C | Bytes: 1 Cycles: 4
	// DEC E | Bytes: 1 Cycles: 4
	// DEC L | Bytes: 1 Cycles: 4
	// DEC A | Bytes: 1 Cycles: 4
}
