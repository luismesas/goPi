/* 
Package piface is for interfacing with a PiFace Digital IO board. 

 - More Info at http://www.piface.org.uk/products/piface_digital/
 - Guides at http://prod.www.piface.org.uk/guides/
 - goPi blink example at  https://github.com/luismesas/goPi/blob/master/examples/blink_piface/blink_piface.go

Its possible to connect up to four PiFace boards to a Raspberry Pi, however some jumper
settings are required to set the hardware address. 

 - See http://prod.www.piface.org.uk/guides/howto/PiFace_Digital_Jumper_Settings/

*/
package piface



import (
	"fmt"
	"github.com/luismesas/go-rpi/MCP23S17"
)

// A PiFace Digital board.
type PiFaceDigital struct {
	mcp *MCP23S17.MCP23S17

	InputPins  []*MCP23S17.MCP23S17RegisterBitNeg
	InputPort  *MCP23S17.MCP23S17RegisterNeg
	OutputPins []*MCP23S17.MCP23S17RegisterBit
	OutputPort *MCP23S17.MCP23S17Register
	Leds       []*MCP23S17.MCP23S17RegisterBit
	Relays     []*MCP23S17.MCP23S17RegisterBit
	Switches   []*MCP23S17.MCP23S17RegisterBit
}

// Create a new PiFaceDigital Instance
func NewPiFaceDigital(hardware_addr byte, bus int, chip_select int) *PiFaceDigital {
	pfd := new(PiFaceDigital)
	pfd.mcp = MCP23S17.NewMCP23S17(hardware_addr, bus, chip_select)
	// pfd.device = interrupts.NewGPIOInterruptDevice()

	pfd.InputPins = make([]*MCP23S17.MCP23S17RegisterBitNeg, 8)
	for i := range pfd.InputPins {
		pfd.InputPins[i] = MCP23S17.NewMCP23S17RegisterBitNeg(uint(i), MCP23S17.GPIOB, pfd.mcp)
	}

	pfd.InputPort = MCP23S17.NewMCP23S17RegisterNeg(MCP23S17.GPIOB, pfd.mcp)

	pfd.OutputPins = make([]*MCP23S17.MCP23S17RegisterBit, 8)
	for i := range pfd.OutputPins {
		pfd.OutputPins[i] = MCP23S17.NewMCP23S17RegisterBit(uint(i), MCP23S17.GPIOA, pfd.mcp)
	}

	pfd.OutputPort = MCP23S17.NewMCP23S17Register(MCP23S17.GPIOA, pfd.mcp)

	pfd.Leds = make([]*MCP23S17.MCP23S17RegisterBit, 8)
	for i := range pfd.Leds {
		pfd.Leds[i] = MCP23S17.NewMCP23S17RegisterBit(uint(i), MCP23S17.GPIOA, pfd.mcp)
	}

	pfd.Relays = make([]*MCP23S17.MCP23S17RegisterBit, 2)
	for i := range pfd.Relays {
		pfd.Relays[i] = MCP23S17.NewMCP23S17RegisterBit(uint(i), MCP23S17.GPIOA, pfd.mcp)
	}

	pfd.Switches = make([]*MCP23S17.MCP23S17RegisterBit, 4)
	for i := range pfd.Switches {
		pfd.Switches[i] = MCP23S17.NewMCP23S17RegisterBit(uint(i), MCP23S17.GPIOB, pfd.mcp)
	}

	return pfd
}

// Initialize the board
func (pfd *PiFaceDigital) InitBoard() error {

	err := pfd.mcp.Open()
	if err != nil {
		return err
	}

	var ioconfig byte
	ioconfig = (MCP23S17.BANK_OFF |
		MCP23S17.INT_MIRROR_OFF |
		MCP23S17.SEQOP_OFF |
		MCP23S17.DISSLW_OFF |
		MCP23S17.HAEN_ON |
		MCP23S17.ODR_OFF |
		MCP23S17.INTPOL_LOW)

	pfd.mcp.IOCON.SetValue(ioconfig)
	if pfd.mcp.IOCON.Value() != ioconfig {
		return fmt.Errorf("No PiFace Digital board detected (hardware_addr=%d, bus=%b, chip_select=%b).", pfd.mcp.HardwareAddress, pfd.mcp.Device.Bus, pfd.mcp.Device.Chip)
	}

	pfd.mcp.GPIOa.SetValue(0)
	pfd.mcp.IODIRa.SetValue(0)    // GPIOA as outputs
	pfd.mcp.IODIRb.SetValue(0xFF) // GPIOB as inputs
	pfd.mcp.GPPUb.SetValue(0xFF)  // input pullups on
	// pfd.EnableInterrupts()

	return nil
}

// Not Implemented
func (pfd *PiFaceDigital) EnableInterrupts() error {
	return fmt.Errorf("EnableInterrupts() Not implemented")
}

// Not Implemented
func (pfd *PiFaceDigital) Open() error {
	return fmt.Errorf("Open() Not implemented")
}

// Not Implemented
func (pfd *PiFaceDigital) Close() error {
	return fmt.Errorf("Close() Not implemented")
}
