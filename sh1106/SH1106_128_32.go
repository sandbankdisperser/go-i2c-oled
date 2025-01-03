package SH1106

import "github.com/sandbankdisperser/go-i2c-oled/i2c"

type SH1106_128_32 struct {
	conn     *i2c.I2c
	vccState byte
	width    int
	height   int
	contrast int
}

func (d *SH1106_128_32) VCCState() byte {
	return d.vccState
}
func (d *SH1106_128_32) Height() int {
	return d.height
}
func (d *SH1106_128_32) Width() int {
	return d.width
}

// NewSSD1306_96_16 creates a new instance of the SSD1306_96_16 structure.
func NewSH1106_128_32(fd *i2c.I2c, vccstate byte) *SH1106_128_32 {
	return &SH1106_128_32{
		conn:   fd,
		height: 32,
		width:  128,
	}
}

func (d *SH1106_128_32) Initialize() error {
	data := []byte{
		SH110X_DISPLAYOFF,               // 0xAE
		SH110X_SETDISPLAYCLOCKDIV, 0x80, // 0xD5, 0x80,
		SH110X_SETMULTIPLEX, 0x3F, // 0xA8, 0x3F,
		SH110X_SETDISPLAYOFFSET, 0x00, // 0xD3, 0x00,
		SH110X_SETSTARTLINE, // 0x40
		SH110X_DCDC, 0x8B,   // DC/DC on
		SH110X_SEGREMAP + 1,     // 0xA1
		SH110X_COMSCANDEC,       // 0xC8
		SH110X_SETCOMPINS, 0x12, // 0xDA, 0x12,
		SH110X_SETCONTRAST, 0xFF, // 0x81, 0xFF
		SH110X_SETPRECHARGE, 0x1F, // 0xD9, 0x1F,
		SH110X_SETVCOMDETECT, 0x40, // 0xDB, 0x40,
		0x33, // Set VPP to 9V
		SH110X_NORMALDISPLAY,
		SH110X_MEMORYMODE, 0x10, // 0x20, 0x00
		SH110X_DISPLAYALLON_RESUME,
	}

	return sendCommands(d.conn, data...)
}
