package spi

import (
	"log"
	"fmt"
	"os"
	"unsafe"
	"github.com/luismesas/go-rpi/ioctl"
)

const SPIDEV = "/dev/spidev"
const SPI_HELP_LINK = "http://piface.github.io/pifacecommon/installation.html#enable-the-spi-module"

// Defaults
const (
	SPI_HARDWARE_ADDR = 0
	SPI_BUS = 0
	SPI_CHIP = 0
	SPI_DELAY = 0
)


type SPIDevice struct{
	Bus int // 0
	Chip int // 0
	file *os.File // nil

	mode uint8
	bpw uint8
	speed uint32
}

// An SPI Device at /dev/spi<bus>.<chip_select>.
func NewSPIDevice(bus int, chipSelect int) *SPIDevice{
	spi := new(SPIDevice)
	spi.Bus = bus
	spi.Chip = chipSelect

	return spi
}

// Opens SPI device
func (spi *SPIDevice) Open() error{
	spiDevice := fmt.Sprintf("%s%d.%d", SPIDEV, spi.Bus, spi.Chip)

	var err error
	spi.file, err = os.OpenFile(spiDevice, os.O_RDWR, 0)
	// spi.file, err = os.Create(spiDevice)
	if err != nil {
		return fmt.Errorf("I can't see %s. Have you enabled the SPI module? (%s)", spiDevice, SPI_HELP_LINK)
	}

	log.Println("SPI Open")
	return nil
}

// Closes SPI device
func (spi *SPIDevice) Close() error{
	err := spi.file.Close()
	if err != nil {
		return fmt.Errorf("Error closing spi", err)
	}
	return nil
}

// Sends bytes over SPI channel and returns []byte response
func (spi *SPIDevice) Send(bytes_to_send [3]byte) ([]byte, error){
	wBuffer := bytes_to_send
	rBuffer := [3]byte{}

	// generates message
	transfer := SPI_IOC_TRANSFER{}
	transfer.txBuf = uint64( uintptr( unsafe.Pointer(&wBuffer)))
	transfer.rxBuf = uint64( uintptr( unsafe.Pointer(&rBuffer)))
	transfer.length = uint32(unsafe.Sizeof(wBuffer))
	transfer.delayUsecs = SPI_DELAY
	transfer.bitsPerWord = spi.bpw
	transfer.speedHz = spi.speed

	// sends message over SPI
	err := ioctl.IOCTL(spi.file.Fd(), SPI_IOC_MESSAGE(1), uintptr(unsafe.Pointer(&transfer)))
	if err != nil {
		return nil, fmt.Errorf("Error on sending: %s\n", err)
	}

	// generates a valid response
	ret := make([]byte, unsafe.Sizeof(rBuffer))
	for i := range(ret) {
		ret[i] = rBuffer[i]
	}

	return ret, nil
}

func (spi *SPIDevice) SetMode(mode uint8) error{
	spi.mode = mode
	err := ioctl.IOCTL(spi.file.Fd(), SPI_IOC_WR_MODE(), uintptr(unsafe.Pointer(&mode)))
	if err != nil {
		return fmt.Errorf("Error setting mode: %s\n", err)
	}
	return nil
}

func (spi *SPIDevice) SetBitsPerWord(bpw uint8) error{
	spi.bpw = bpw
	err := ioctl.IOCTL(spi.file.Fd(), SPI_IOC_WR_BITS_PER_WORD(), uintptr(unsafe.Pointer(&bpw)))
	if err != nil {
		return fmt.Errorf("Error setting bits per word: %s\n", err)
	}
	return nil
}

func (spi *SPIDevice) SetSpeed(speed uint32) error{
	spi.speed = speed
	err := ioctl.IOCTL(spi.file.Fd(), SPI_IOC_WR_MAX_SPEED_HZ(), uintptr(unsafe.Pointer(&speed)))
	if err != nil {
		return fmt.Errorf("Error setting speed: %s\n", err)
	}
	return nil
}
