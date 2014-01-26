package MCP23S17

// A bit inside register inside an MCP23S17.
type MCP23S17RegisterBit struct {
	bit_num uint
	address byte
	chip    *MCP23S17
}

func NewMCP23S17RegisterBit(bit_num uint, address byte, chip *MCP23S17) *MCP23S17RegisterBit {
	bit := new(MCP23S17RegisterBit)
	bit.bit_num = bit_num
	bit.address = address
	bit.chip = chip
	return bit
}

func (register *MCP23S17RegisterBit) Value() byte {
	return register.chip.ReadBit(register.bit_num, register.address)
}

func (register *MCP23S17RegisterBit) SetValue(value byte) {
	register.chip.WriteBit(value, register.bit_num, register.address)
}

func (register *MCP23S17RegisterBit) AllHigh() {
	register.SetValue(1)
}

func (register *MCP23S17RegisterBit) AllLow() {
	register.SetValue(0)
}

func (register *MCP23S17RegisterBit) AllOn() {
	register.SetValue(1)
}

func (register *MCP23S17RegisterBit) AllOff() {
	register.SetValue(0)
}

func (register *MCP23S17RegisterBit) Toggle() {
	register.SetValue(1 ^ register.Value())
}
