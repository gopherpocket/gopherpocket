package asm

import "fmt"

func ExampleJP() {
	print := func(i *Instruction) {
		if i.CyclesAlt != 0 {
			fmt.Println(i, "| Bytes:", i.Bytes, "Cycles (Jump Taken):", i.Cycles, "Cycles (Jump not Taken):", i.CyclesAlt)
		} else {
			fmt.Println(i, "| Bytes:", i.Bytes, "Cycles:", i.Cycles)
		}
	}

	print(JP(Imm16(127)))
	print(JP(HL))
	print(JP(Label("foo")))

	print(JPZ(Imm16(127)))
	print(JPZ(Label("foo")))

	print(JPNZ(Imm16(127)))
	print(JPNZ(Label("foo")))

	print(JPC(Imm16(127)))
	print(JPC(Label("foo")))

	print(JPNC(Imm16(127)))
	print(JPNC(Label("foo")))

	// Output:
	// JP $7F | Bytes: 3 Cycles: 16
	// JP HL | Bytes: 1 Cycles: 4
	// JP foo | Bytes: 3 Cycles: 16
	// JP Z, $7F | Bytes: 3 Cycles (Jump Taken): 16 Cycles (Jump not Taken): 12
	// JP Z, foo | Bytes: 3 Cycles (Jump Taken): 16 Cycles (Jump not Taken): 12
	// JP NZ, $7F | Bytes: 3 Cycles (Jump Taken): 16 Cycles (Jump not Taken): 12
	// JP NZ, foo | Bytes: 3 Cycles (Jump Taken): 16 Cycles (Jump not Taken): 12
	// JP C, $7F | Bytes: 3 Cycles (Jump Taken): 16 Cycles (Jump not Taken): 12
	// JP C, foo | Bytes: 3 Cycles (Jump Taken): 16 Cycles (Jump not Taken): 12
	// JP NC, $7F | Bytes: 3 Cycles (Jump Taken): 16 Cycles (Jump not Taken): 12
	// JP NC, foo | Bytes: 3 Cycles (Jump Taken): 16 Cycles (Jump not Taken): 12
}
