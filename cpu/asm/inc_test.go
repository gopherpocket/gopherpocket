package asm

import "fmt"

func ExampleINC() {
	print := func(i *Instruction) {
		fmt.Println(i, "| Bytes:", i.Bytes, "Cycles:", i.Cycles)
	}

	print(INC(BC))
	print(INC(DE))
	print(INC(HL))
	print(INC(SP))
	print(INC(B))
	print(INC(D))
	print(INC(H))
	print(INC(Ptr(HL)))
	print(INC(C))
	print(INC(E))
	print(INC(L))
	print(INC(A))

	//Output:
	//
	// INC BC | Bytes: 1 Cycles: 8
	// INC DE | Bytes: 1 Cycles: 8
	// INC HL | Bytes: 1 Cycles: 8
	// INC SP | Bytes: 1 Cycles: 8
	// INC B | Bytes: 1 Cycles: 4
	// INC D | Bytes: 1 Cycles: 4
	// INC H | Bytes: 1 Cycles: 4
	// INC [HL] | Bytes: 1 Cycles: 12
	// INC C | Bytes: 1 Cycles: 4
	// INC E | Bytes: 1 Cycles: 4
	// INC L | Bytes: 1 Cycles: 4
	// INC A | Bytes: 1 Cycles: 4
}
