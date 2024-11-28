package ssd1306

import (
	"os"
)

const (
	SSD1306_CMD                 = 0x80
	SSD1306_SETDISPLAYCLOCKDIV  = 0xD5
	SSD1306_DISPLAYOFF          = 0xAE
	SSD1306_SETMULTIPLEX        = 0xA8
	SSD1306_SETDISPLAYOFFSET    = 0xD3
	SSD1306_SETSTARTLINE        = 0x0
	SSD1306_CHARGEPUMP          = 0x8D
	SSD1306_MEMORYMODE          = 0x20
	SSD1306_SEGREMAP            = 0xA0
	SSD1306_COMSCANDEC          = 0xC8
	SSD1306_SETCOMPINS          = 0xDA
	SSD1306_SETCONTRAST         = 0x81
	SSD1306_SETPRECHARGE        = 0xD9
	SSD1306_SETVCOMDETECT       = 0xDB
	SSD1306_DISPLAYALLON_RESUME = 0xA4
	SSD1306_NORMALDISPLAY       = 0xA6
	SSD1306_EXTERNALVCC         = 0x1
	SSD1306_SWITCHCAPVCC        = 0x2
)

// writeCommand sends a single command byte to the SSD1306 device.
func writeCommand(fd *os.File, cmd byte) (int, error) {
	return fd.Write([]byte{SSD1306_CMD, cmd})
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
