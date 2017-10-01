package spi

import (
	"unsafe"

	"github.com/luismesas/goPi/ioctl"
)

const SPI_IOC_MAGIC = 107

// Read of SPI mode (SPI_MODE_0..SPI_MODE_3)
func SPI_IOC_RD_MODE() uintptr {
	return ioctl.IOR(SPI_IOC_MAGIC, 1, 1)
}

// Write of SPI mode (SPI_MODE_0..SPI_MODE_3)
func SPI_IOC_WR_MODE() uintptr {
	return ioctl.IOW(SPI_IOC_MAGIC, 1, 1)
}

// Read SPI bit justification
func SPI_IOC_RD_LSB_FIRST() uintptr {
	return ioctl.IOR(SPI_IOC_MAGIC, 2, 1)
}

// Write SPI bit justification
func SPI_IOC_WR_LSB_FIRST() uintptr {
	return ioctl.IOW(SPI_IOC_MAGIC, 2, 1)
}

// Read SPI device word length (1..N)
func SPI_IOC_RD_BITS_PER_WORD() uintptr {
	return ioctl.IOR(SPI_IOC_MAGIC, 3, 1)
}

// Write SPI device word length (1..N)
func SPI_IOC_WR_BITS_PER_WORD() uintptr {
	return ioctl.IOW(SPI_IOC_MAGIC, 3, 1)
}

// Read SPI device default max speed hz
func SPI_IOC_RD_MAX_SPEED_HZ() uintptr {
	return ioctl.IOR(SPI_IOC_MAGIC, 4, 4)
}

// Write SPI device default max speed hz
func SPI_IOC_WR_MAX_SPEED_HZ() uintptr {
	return ioctl.IOW(SPI_IOC_MAGIC, 4, 4)
}

// Write custom SPI message
func SPI_IOC_MESSAGE(n uintptr) uintptr {
	return ioctl.IOW(SPI_IOC_MAGIC, 0, uintptr(SPI_MESSAGE_SIZE(n)))
}

func SPI_MESSAGE_SIZE(n uintptr) uintptr {
	if (n * unsafe.Sizeof(SPI_IOC_TRANSFER{})) < (1 << ioctl.IOC_SIZEBITS) {
		return (n * unsafe.Sizeof(SPI_IOC_TRANSFER{}))
	}
	return 0
}

type SPI_IOC_TRANSFER struct {
	txBuf       uint64
	rxBuf       uint64
	length      uint32
	speedHz     uint32
	delayUsecs  uint16
	bitsPerWord uint8
	csChange    uint8
	pad         uint32
}
