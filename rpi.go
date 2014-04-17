
// Package rpi is a library for interfacing with Raspberry Pi IO
package rpi

// This is the wrong url here ??? 
import (
	_ "github.com/luismesas/go-rpi/MCP23S17"
	_ "github.com/luismesas/go-rpi/ioctl"
	_ "github.com/luismesas/go-rpi/spi"
	_ "github.com/luismesas/go-rpi/piface"
)
