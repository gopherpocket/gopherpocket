package cpu

// Register represents a typical 16-bit register.
type Register uint16

// Hi returns the most-significant 8 bits in a register.
// For example, in the bit string representation of a register: 1111_1111_0000_0000
// the first eight "1" bits would be returned as a uint8.
func (r Register) Hi() uint8 {
	return uint8(r >> 8)
}

// SetHi sets the most-significant 8 bits in the register.
func (r *Register) SetHi(x uint8) {
	*r = (Register(x) << 8) | Register(r.Lo())
}

// Lo returns the least-significant 8 bits in a register.
// For example, in the bit string representation of a register: 0000_0000_1111_1111
// the trailing eight "1" bits would be returned as a uint8
func (r Register) Lo() uint8 {
	return uint8(r)
}

// SetLo sets the least-significant 8 bits in the register
func (r *Register) SetLo(x uint8) {
	*r = Register(x) | (Register(r.Hi()) << 8)
}

// Registers represents the entire Register file of a Gameboy CPU.
type Registers struct {
	// The high bits of AF is called the "accumulator", and the low bits are the flags.
	AF Register
	BC Register
	DE Register
	HL Register

	SP Register
	PC Register
}

// Flags returns the Flags register from the Lo bits of the AF register.
func (r *Registers) Flags() Flags {
	return Flags(r.AF.Lo())
}

// Flags represents an 8-bit flags register, found in the lower 8-bits of the AF register.
type Flags uint8

// Z is the zero flag of the Flags register, at bit 7.
// Returns true if Z == 1
func (f Flags) Z() bool {
	const bit7 = (1 << 7)
	return f&bit7 == bit7
}

// N is the subtraction flag of the Flags register, at bit 6.
// See the [Pandocs on BCD Flags]: https://gbdev.io/pandocs/CPU_Registers_and_Flags.html#the-bcd-flags-n-h
func (f Flags) N() bool {
	const bit6 = (1 << 6)
	return f&bit6 == bit6
}

// H is the half carry flag of the Flags register, at bit 5.
// See the [Pandocs on BCD Flags]: https://gbdev.io/pandocs/CPU_Registers_and_Flags.html#the-bcd-flags-n-h
func (f Flags) H() bool {
	const bit5 = (1 << 5)
	return f&bit5 == bit5
}

// C is the carry flag of the Flags register, at bit 4
func (f Flags) C() bool {
	const bit4 = (1 << 4)
	return f&bit4 == bit4
}
