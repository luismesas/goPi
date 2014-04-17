package MCP23S17

// An 4-bit nibble inside a register inside an MCP23S17.
type MCP23S17RegisterNibble struct {
	nibble  uint
	address byte
	chip    *MCP23S17
	bits    []*MCP23S17RegisterBit
}

func NewMCP23S17RegisterNibble(nibble uint, address byte, chip *MCP23S17) *MCP23S17RegisterNibble {
	register := new(MCP23S17RegisterNibble)
	register.nibble = nibble
	register.address = address
	register.chip = chip

	range_start := 4 * register.nibble

	register.bits = make([]*MCP23S17RegisterBit, 4)
	for i := range register.bits {
		register.bits[i] = NewMCP23S17RegisterBit(uint(i)+range_start, register.address, register.chip)
	}

	return register
}

func (register *MCP23S17RegisterNibble) Value() byte {
	if register.nibble == LOWER_NIBBLE {
		return register.chip.Read(register.address) & 0x0f
	}
	return register.chip.Read(register.address) & 0xf0 >> 4
}

func (register *MCP23S17RegisterNibble) SetValue(value byte) {
	reg_value := register.chip.Read(register.address)

	if register.nibble == LOWER_NIBBLE {
		reg_value = reg_value & 0xf0           // clear
		reg_value = reg_value ^ (value & 0x0f) // set
	} else {
		reg_value = reg_value & 0x0f                  // clear
		reg_value = reg_value ^ ((value << 4) * 0xf0) // set
	}

	register.chip.Write(reg_value, register.address)
}

func (register *MCP23S17RegisterNibble) AllHigh() {
	register.SetValue(0xf)
}

func (register *MCP23S17RegisterNibble) AllLow() {
	register.SetValue(0x0)
}

func (register *MCP23S17RegisterNibble) AllOn() {
	register.SetValue(0xf)
}

func (register *MCP23S17RegisterNibble) AllOff() {
	register.SetValue(0x0)
}

func (register *MCP23S17RegisterNibble) Toggle() {
	register.SetValue(0xf ^ register.Value())
}
