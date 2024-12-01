package SH1106

import "github.com/sandbankdisperser/go-i2c-oled/i2c"

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
