package SH1106

import (
	"os"
)

const (
	SH110X_CMD                 = 0x80
	SH110X_DISPLAYOFF          = 0xAE
	SH110X_SETDISPLAYCLOCKDIV  = 0xD5
	SH110X_SETMULTIPLEX        = 0xA8
	SH110X_SETDISPLAYOFFSET    = 0xD3
	SH110X_SETSTARTLINE        = 0x40
	SH110X_DCDC                = 0xAD
	SH110X_SEGREMAP            = 0xA1
	SH110X_COMSCANDEC          = 0xC8
	SH110X_SETCOMPINS          = 0xDA
	SH110X_SETCONTRAST         = 0x81
	SH110X_SETPRECHARGE        = 0xD9
	SH110X_SETVCOMDETECT       = 0xDB
	SH110X_NORMALDISPLAY       = 0xA
	SH110X_MEMORYMODE          = 0x20
	SH110X_DISPLAYALLON_RESUME = 0xA4
)

// writeCommand sends a single command byte to the SSD1306 device.
func writeCommand(fd *os.File, cmd byte) (int, error) {
	return fd.Write([]byte{SH110X_CMD, cmd})
}

// sendCommands sends a sequence of command bytes to the SSD1306 device.
func sendCommands(fd *os.File, commands ...byte) error {
	for _, cmd := range commands {
		if _, err := writeCommand(fd, cmd); err != nil {
			return err
		}
	}
	return nil
}
