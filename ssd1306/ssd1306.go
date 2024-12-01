package ssd1306

import (
	"github.com/sandbankdisperser/go-i2c-oled/i2c"
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
const (
	I2C_SLAVE = 0x0703

	OLED_CMD                 = 0x80
	OLED_CMD_COL_ADDRESSING  = 0x21
	OLED_CMD_PAGE_ADDRESSING = 0x22
	OLED_CMD_CONTRAST        = 0x81
	OLED_CMD_START_COLUMN    = 0x00
	OLED_CMD_HIGH_COLUMN     = 0x10
	OLED_CMD_DISPLAY_OFF     = 0xAE
	OLED_CMD_DISPLAY_ON      = 0xAF

	OLED_DATA            = 0x40
	OLED_ADRESSING       = 0x21
	OLED_ADRESSING_START = 0xB0
	OLED_ADRESSING_COL   = 0x21
	OLED_END             = 0x10
	PIXSIZE              = 8
)

// sendCommands sends a sequence of command bytes to the SSD1306 device.
func sendCommands(conn *i2c.I2c, commands ...byte) error {
	for _, cmd := range commands {
		if _, err := conn.WriteCommand(cmd); err != nil {
			return err
		}
	}
	return nil
}
