// Package rpi is a library for interfacing with Raspberry Pi IO
package rpi

// This is the wrong url here ???
import (
	_ "github.com/luismesas/goPi/MCP23S17"
	_ "github.com/luismesas/goPi/ioctl"
	_ "github.com/luismesas/goPi/piface"
	_ "github.com/luismesas/goPi/spi"
)
