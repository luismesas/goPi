package MCP23S17

// An 8-bit register inside an MCP23S17.
type MCP23S17Register struct {
	address     byte
	chip        *MCP23S17
	lowerNibble *MCP23S17RegisterNibble
	upperNibble *MCP23S17RegisterNibble
	bits        []*MCP23S17RegisterBit
}

func NewMCP23S17Register(address byte, chip *MCP23S17) *MCP23S17Register {
	register := new(MCP23S17Register)
	register.address = address
	register.chip = chip

	register.lowerNibble = NewMCP23S17RegisterNibble(LOWER_NIBBLE, register.address, register.chip)
	register.upperNibble = NewMCP23S17RegisterNibble(UPPER_NIBBLE, register.address, register.chip)

	register.bits = make([]*MCP23S17RegisterBit, 8)
	for i := range register.bits {
		register.bits[i] = NewMCP23S17RegisterBit(uint(i), register.address, register.chip)
	}

	return register
}

func (register *MCP23S17Register) Value() byte {
	return register.chip.Read(register.address)
}

func (register *MCP23S17Register) SetValue(value byte) {
	register.chip.Write(value, register.address)
}

func (register *MCP23S17Register) AllHigh() {
	register.SetValue(0xff)
}

func (register *MCP23S17Register) AllLow() {
	register.SetValue(0x00)
}

func (register *MCP23S17Register) AllOn() {
	register.SetValue(0xff)
}

func (register *MCP23S17Register) AllOff() {
	register.SetValue(0x00)
}

func (register *MCP23S17Register) Toggle() {
	register.SetValue(0xff ^ register.Value())
}
