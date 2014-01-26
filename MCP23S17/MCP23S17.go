package MCP23S17

import (
	"github.com/luismesas/go-rpi/spi"
)

const (
	MCP23S17_MODE = 0
	MCP23S17_BPW = 8
	MCP23S17_SPEED = 10000000
)

//	Microchip's MCP23S17: A 16-Bit I/O Expander with Serial Interface.
type MCP23S17 struct{
	Device *spi.SPIDevice
	HardwareAddress byte
	
	// Controls the direction of the data I/O.
	IODIRa *MCP23S17Register
	IODIRb *MCP23S17Register

	// This register allows the user to configure the polarity on the corresponding GPIO port bits.
	IPOLa *MCP23S17Register
	IPOLb *MCP23S17Register

	// The GPINTEN register controls the interrupt-onchange feature for each pin.
	GPINTENa *MCP23S17Register
	GPINTENb *MCP23S17Register

	// The default comparison value is configured in the DEFVAL register.
	DEFVALa *MCP23S17Register
	DEFVALb *MCP23S17Register

	// The INTCON register controls how the associated pin value is compared for the interrupt-on-change feature.
	INTCONa *MCP23S17Register
	INTCONb *MCP23S17Register

	// The IOCON register contains several bits for configuring the device.
	IOCON *MCP23S17Register

	// The GPPU register controls the pull-up esistors for the port pins.
	GPPUa *MCP23S17Register
	GPPUb *MCP23S17Register

	// The INTF register reflects the interrupt condition on the port pins of any pin that is enabled for interrupts via the GPINTEN register.
	INTFa *MCP23S17Register
	INTFb *MCP23S17Register

	// The INTCAP register captures the GPIO port value at the time the interrupt occurred.
	INTCAPa *MCP23S17Register
	INTCAPb *MCP23S17Register

	// The GPIO register reflects the value on the port.
	GPIOa *MCP23S17Register
	GPIOb *MCP23S17Register

	// The OLAT register provides access to the output latches.
	OLATa *MCP23S17Register
	OLATb *MCP23S17Register
}

func NewMCP23S17(hardwareAddress uint8, bus int, chip_select int) *MCP23S17{
	mcp := new(MCP23S17)
	mcp.Device = spi.NewSPIDevice(bus, chip_select)
	mcp.HardwareAddress = hardwareAddress

	mcp.IODIRa = NewMCP23S17Register(IODIRA, mcp)
	mcp.IODIRb = NewMCP23S17Register(IODIRB, mcp)
	mcp.IPOLa = NewMCP23S17Register(IPOLA, mcp)
	mcp.IPOLb = NewMCP23S17Register(IPOLB, mcp)
	mcp.GPINTENa = NewMCP23S17Register(GPINTENA, mcp)
	mcp.GPINTENb = NewMCP23S17Register(GPINTENB, mcp)
	mcp.DEFVALa = NewMCP23S17Register(DEFVALA, mcp)
	mcp.DEFVALb = NewMCP23S17Register(DEFVALB, mcp)
	mcp.INTCONa = NewMCP23S17Register(INTCONA, mcp)
	mcp.INTCONb = NewMCP23S17Register(INTCONB, mcp)
	mcp.IOCON = NewMCP23S17Register(IOCON, mcp)
	mcp.GPPUa = NewMCP23S17Register(GPPUA, mcp)
	mcp.GPPUb = NewMCP23S17Register(GPPUB, mcp)
	mcp.INTFa = NewMCP23S17Register(INTFA, mcp)
	mcp.INTFb = NewMCP23S17Register(INTFB, mcp)
	mcp.INTCAPa = NewMCP23S17Register(INTCAPA, mcp)
	mcp.INTCAPb = NewMCP23S17Register(INTCAPB, mcp)
	mcp.GPIOa = NewMCP23S17Register(GPIOA, mcp)
	mcp.GPIOb = NewMCP23S17Register(GPIOB, mcp)
	mcp.OLATa = NewMCP23S17Register(OLATA, mcp)
	mcp.OLATb = NewMCP23S17Register(OLATB, mcp)

	return mcp
}

func (mcp *MCP23S17) Open() error{

	err := mcp.Device.Open()
	if err != nil {
		return err
	}

	err = mcp.Device.SetMode(MCP23S17_MODE)
	if err != nil {
		return err
	}

	err = mcp.Device.SetBitsPerWord(MCP23S17_BPW)
	if err != nil {
		return err
	}

	err = mcp.Device.SetSpeed(MCP23S17_SPEED)
	if err != nil {
		return err
	}

	return nil
}

func (mcp *MCP23S17) Close() error{
	return mcp.Device.Close()
}


// Returns an SPI control byte.
// The MCP23S17 is a slave SPI device. The slave address contains
// four fixed bits and three user-defined hardware address bits
// (if enabled via IOCON.HAEN) (pins A2, A1 and A0) with the
// read/write bit filling out the control byte::

// 	+--------------------+
// 	|0|1|0|0|A2|A1|A0|R/W|
// 	+--------------------+
// 	 7 6 5 4 3  2  1   0

// :param read_write_cmd: Read or write command.
// :type read_write_cmd: int	
func (mcp *MCP23S17) getSPIControlByte(read_write_cmd uint8) byte {
	board_addr_pattern := (mcp.HardwareAddress << 1) & 0xE
	rw_cmd_pattern := read_write_cmd & 0x01  // make sure it's just 1 bit long
	return 0x40 | board_addr_pattern | rw_cmd_pattern
}

// Returns the value of the address specified.
func (mcp *MCP23S17) Read(address byte) byte{
	ctrl_byte := mcp.getSPIControlByte(READ_CMD)
	data, err := mcp.Device.Send([3]byte{ctrl_byte, address, 0})
	if err != nil {
		panic("Error reading from MCP23S17: %s\n", err)
		return 0x00
	}
	if len(data) == 0 {
		return 0x00
	} else {
		return data[2]
	}
}

// Writes data to the address specified.
func (mcp *MCP23S17) Write(data byte, address byte){
	ctrl_byte := mcp.getSPIControlByte(WRITE_CMD)
	_, err := mcp.Device.Send([3]byte{ctrl_byte, address, data})
	if err != nil {
		panic("Error writing on MCP23S17: %s\n", err)
	}
}

// Returns the bit specified from the address.
func (mcp *MCP23S17) ReadBit(bit_num uint, address byte) byte{
	value := mcp.Read(address)
	return (value >> bit_num) & 1
}

// Writes the value given to the bit in the address specified.
func (mcp *MCP23S17) WriteBit(data byte, bit_num uint, address byte){
	value := mcp.Read(address)
	if data > 0 {
		value = value | (1 << bit_num) //set
	} else {
		value = value & ( 0xff ^ (1 << bit_num)) //clear
	}
	mcp.Write(value, address)
}

// Clears the interrupt flags by reading the capture register.
func (mcp *MCP23S17) ClearInterrupts(port int){
	var address byte
	address = INTCAPA
	if port == GPIOA {
		address = INTCAPB
	}
	mcp.Read(address)
}
