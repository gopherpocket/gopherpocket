package asm

import "fmt"

func ExampleJR() {
	print := func(i *Instruction) {
		if i.CyclesAlt != 0 {
			fmt.Println(i, "| Bytes:", i.Bytes, "Cycles (Jump Taken):", i.Cycles, "Cycles (Jump not Taken):", i.CyclesAlt)
		} else {
			fmt.Println(i, "| Bytes:", i.Bytes, "Cycles:", i.Cycles)
		}
	}

	print(JR(Offset8(127)))
	print(JR(Offset8(-128)))
	print(JR(Label("foo")))

	print(JRZ(Offset8(127)))
	print(JRZ(Offset8(-128)))
	print(JRZ(Label("foo")))

	print(JRNZ(Offset8(127)))
	print(JRNZ(Offset8(-128)))
	print(JRNZ(Label("foo")))

	print(JRC(Offset8(127)))
	print(JRC(Offset8(-128)))
	print(JRC(Label("foo")))

	print(JRNC(Offset8(127)))
	print(JRNC(Offset8(-128)))
	print(JRNC(Label("foo")))

	// Output:
	//
	// JR $7F | Bytes: 2 Cycles: 12
	// JR $-80 | Bytes: 2 Cycles: 12
	// JR foo | Bytes: 2 Cycles: 12
	// JR Z, $7F | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
	// JR Z, $-80 | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
	// JR Z, foo | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
	// JR NZ, $7F | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
	// JR NZ, $-80 | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
	// JR NZ, foo | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
	// JR C, $7F | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
	// JR C, $-80 | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
	// JR C, foo | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
	// JR NC, $7F | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
	// JR NC, $-80 | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
	// JR NC, foo | Bytes: 2 Cycles (Jump Taken): 12 Cycles (Jump not Taken): 8
}
