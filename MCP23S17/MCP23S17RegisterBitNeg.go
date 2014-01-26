package MCP23S17

// A bit inside register inside an MCP23S17.
type MCP23S17RegisterBitNeg struct {
	bit_num uint
	address byte
	chip    *MCP23S17
}

func NewMCP23S17RegisterBitNeg(bit_num uint, address byte, chip *MCP23S17) *MCP23S17RegisterBitNeg {
	bit := new(MCP23S17RegisterBitNeg)
	bit.bit_num = bit_num
	bit.address = address
	bit.chip = chip
	return bit
}

func (register *MCP23S17RegisterBitNeg) Value() byte {
	return 1 ^ register.chip.ReadBit(register.bit_num, register.address)
}

func (register *MCP23S17RegisterBitNeg) SetValue(value byte) {
	register.chip.WriteBit(value^1, register.bit_num, register.address)
}

func (register *MCP23S17RegisterBitNeg) AllHigh() {
	register.SetValue(1)
}

func (register *MCP23S17RegisterBitNeg) AllLow() {
	register.SetValue(0)
}

func (register *MCP23S17RegisterBitNeg) AllOn() {
	register.SetValue(1)
}

func (register *MCP23S17RegisterBitNeg) AllOff() {
	register.SetValue(0)
}

func (register *MCP23S17RegisterBitNeg) Toggle() {
	register.SetValue(1 ^ register.Value())
}
