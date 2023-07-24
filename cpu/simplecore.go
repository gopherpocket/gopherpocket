package cpu

type SimpleCore struct {
	Registers
	Memory *Memory
}

// func (c *SimpleCore) fetch() (op uint8, err error) {
// 	op, err = c.Memory.ReadUint8At(int64(c.PC))
// 	if err != nil {
// 		err = fmt.Errorf("fetching instruction @ $%X: %v", c.PC, err)
// 	}
// 	return
// }

// func (c *SimpleCore) decode(op uint8) (*asm.Instruction, error) {
// 	// prefixed
// 	var instructionMap opcodedata.InstructionMap
// 	if op == 0xCB {
// 		instructionMap = opcodedata.OpcodeData.Unprefixed
// 	} else {
// 		instructionMap = opcodedata.OpcodeData.Unprefixed
// 	}
// 	instruction, ok := instructionMap[fmt.Sprintf("0x%02X", op)]
// 	if !ok {
// 		return nil, fmt.Errorf("invalid opcode: %02X", op)
// 	}
// 	fmt.Println(instruction)
// 	panic("not finished")
// }

// func (c *SimpleCore) execute() {
// 	panic("not finished")
// }
