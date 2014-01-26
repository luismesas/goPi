package MCP23S17

// Register addresses
const (
	IODIRA   = 0x0  // I/O direction A
	IODIRB   = 0x1  // I/O direction B
	IPOLA    = 0x2  // I/O polarity A
	IPOLB    = 0x3  // I/O polarity B
	GPINTENA = 0x4  // interupt enable A
	GPINTENB = 0x5  // interupt enable B
	DEFVALA  = 0x6  // register default value A (interupts)
	DEFVALB  = 0x7  // register default value B (interupts)
	INTCONA  = 0x8  // interupt control A
	INTCONB  = 0x9  // interupt control B
	IOCON    = 0xA  // I/O config (also 0xB)
	GPPUA    = 0xC  // port A pullups
	GPPUB    = 0xD  // port B pullups
	INTFA    = 0xE  // interupt flag A (where the interupt came from)
	INTFB    = 0xF  // interupt flag B
	INTCAPA  = 0x10 // interupt capture A (value at interupt is saved here)
	INTCAPB  = 0x11 // interupt capture B
	GPIOA    = 0x12 // port A
	GPIOB    = 0x13 // port B
	OLATA    = 0x14 // output latch A
	OLATB    = 0x15 // output latch B
)

// I/O config
const (
	BANK_OFF       = 0x00 // addressing mode
	BANK_ON        = 0x80
	INT_MIRROR_ON  = 0x40 // interupt mirror (INTa|INTb)
	INT_MIRROR_OFF = 0x00
	SEQOP_OFF      = 0x20 // incrementing address pointer
	SEQOP_ON       = 0x00
	DISSLW_ON      = 0x10 // slew rate
	DISSLW_OFF     = 0x00
	HAEN_ON        = 0x08 // hardware addressing
	HAEN_OFF       = 0x00
	ODR_ON         = 0x04 // open drain for interupts
	ODR_OFF        = 0x00
	INTPOL_HIGH    = 0x02 // interupt polarity
	INTPOL_LOW     = 0x00
)

// Commands
const (
	WRITE_CMD = 0
	READ_CMD  = 1
)

// Nibbles
const (
	LOWER_NIBBLE = 0
	UPPER_NIBBLE = 1
)
