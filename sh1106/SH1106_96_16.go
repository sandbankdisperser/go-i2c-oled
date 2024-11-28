package SH1106

import "os"

type SH1106_96_16 struct {
	fd       *os.File
	vccstate byte
}

func (d *SH1106_96_16) VCCState() byte {
	return d.vccstate
}
func (d *SH1106_96_16) Height() int {
	return 16
}
func (d *SH1106_96_16) Width() int {
	return 96
}

// NewSSD1306_96_16 creates a new instance of the SSD1306_96_16 structure.
func NewSSD1306_96_16(fd *os.File, vccstate byte) *SH1106_96_16 {
	return &SH1106_96_16{
		fd:       fd,
		vccstate: vccstate,
	}
}

func (d *SH1106_96_16) Initialize() error {
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

	return sendCommands(d.fd, data...)
}
