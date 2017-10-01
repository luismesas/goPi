package main

import (
	"log"
	"time"

	"github.com/luismesas/goPi/MCP23S17"
	"github.com/luismesas/goPi/spi"
)

func main() {

	// creates a new MCP23S17 instance
	mcp := MCP23S17.NewMCP23S17(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)

	// GPIOa outputs
	gpioa := make([]*MCP23S17.MCP23S17RegisterBit, 8)
	for i := range gpioa {
		gpioa[i] = MCP23S17.NewMCP23S17RegisterBit(uint(i), MCP23S17.GPIOA, mcp)
	}

	// Connects to chip
	err := mcp.Open()
	if err != nil {
		log.Fatalf("Error connecting to chip: %s\n", err)
	}

	// Blinks output 7
	for {
		gpioa[7].Toggle()
		time.Sleep(time.Second)
	}
}
