package MCP23S17

// An 8-bit register inside an MCP23S17.
type MCP23S17RegisterNeg struct{
	address byte
	chip *MCP23S17
	lowerNibble *MCP23S17RegisterNibbleNeg
	upperNibble *MCP23S17RegisterNibbleNeg
	bits []*MCP23S17RegisterBitNeg
}

func NewMCP23S17RegisterNeg(address byte, chip *MCP23S17) *MCP23S17RegisterNeg{
	register := new(MCP23S17RegisterNeg)
	register.address = address
	register.chip = chip

	register.lowerNibble = NewMCP23S17RegisterNibbleNeg(LOWER_NIBBLE, register.address, register.chip)
	register.upperNibble = NewMCP23S17RegisterNibbleNeg(UPPER_NIBBLE, register.address, register.chip)

	register.bits = make([]*MCP23S17RegisterBitNeg,8)
	for i := range(register.bits) {
		register.bits[i] = NewMCP23S17RegisterBitNeg(uint(i), register.address, register.chip)
	}

	return register
}

func (register *MCP23S17RegisterNeg) Value() byte{
	return 0xff ^ register.chip.Read(register.address)
}

func (register *MCP23S17RegisterNeg) SetValue(value byte){
	register.chip.Write(0xff ^ value, register.address)
}

func (register *MCP23S17RegisterNeg) AllHigh(){
	register.SetValue(0xff)
}

func (register *MCP23S17RegisterNeg) AllLow(){
	register.SetValue(0x00)
}

func (register *MCP23S17RegisterNeg) AllOn(){
	register.SetValue(0xff)
}

func (register *MCP23S17RegisterNeg) AllOff(){
	register.SetValue(0x00)
}

func (register *MCP23S17RegisterNeg) Toggle(){
	register.SetValue(0xff ^ register.Value())
}
