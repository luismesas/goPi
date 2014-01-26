package MCP23S17

// A negated 4-bit nibble inside a register inside an MCP23S17.
type MCP23S17RegisterNibbleNeg struct {
	nibble  uint
	address byte
	chip    *MCP23S17
	bits    []*MCP23S17RegisterBitNeg
}

func NewMCP23S17RegisterNibbleNeg(nibble uint, address byte, chip *MCP23S17) *MCP23S17RegisterNibbleNeg {
	register := new(MCP23S17RegisterNibbleNeg)
	register.nibble = nibble
	register.address = address
	register.chip = chip

	range_start := 4 * register.nibble

	register.bits = make([]*MCP23S17RegisterBitNeg, 4)
	for i := range register.bits {
		register.bits[i] = NewMCP23S17RegisterBitNeg(uint(i)+range_start, register.address, register.chip)
	}

	return register
}

func (register *MCP23S17RegisterNibbleNeg) Value() byte {
	var value byte
	if register.nibble == LOWER_NIBBLE {
		value = register.chip.Read(register.address) & 0x0f
	} else {
		value = register.chip.Read(register.address) & 0xf0 >> 4
	}

	return 0xf ^ value
}

func (register *MCP23S17RegisterNibbleNeg) SetValue(value byte) {
	reg_value := register.chip.Read(register.address)

	if register.nibble == LOWER_NIBBLE {
		reg_value = reg_value & 0xf0                // clear
		reg_value = reg_value ^ (value&0x0f ^ 0x0f) // set
	} else {
		reg_value = reg_value & 0x0f                     // clear
		reg_value = reg_value ^ ((value<<4)*0xf0 ^ 0xf0) // set
	}

	register.chip.Write(reg_value, register.address)
}

func (register *MCP23S17RegisterNibbleNeg) AllHigh() {
	register.SetValue(0xf)
}

func (register *MCP23S17RegisterNibbleNeg) AllLow() {
	register.SetValue(0x0)
}

func (register *MCP23S17RegisterNibbleNeg) AllOn() {
	register.SetValue(0xf)
}

func (register *MCP23S17RegisterNibbleNeg) AllOff() {
	register.SetValue(0x0)
}

func (register *MCP23S17RegisterNibbleNeg) Toggle() {
	register.SetValue(0xf ^ register.Value())
}
