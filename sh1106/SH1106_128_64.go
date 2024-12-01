package SH1106

import (
	"fmt"
	"image"
	"image/color"

	"github.com/sandbankdisperser/go-i2c-oled/i2c"
)

type SH1106_128_64 struct {
	conn     *i2c.I2c
	vccState byte
	width    int
	height   int
	contrast int
}

func (d *SH1106_128_64) VCCState() byte {
	return d.vccState
}
func (d *SH1106_128_64) Height() int {
	return d.height
}
func (d *SH1106_128_64) Width() int {
	return d.width
}

// NewSSD1306_96_16 creates a new instance of the SSD1306_96_16 structure.
func NewSH1106_128_64(fd *i2c.I2c, vccstate byte) *SH1106_128_64 {
	return &SH1106_128_64{
		conn:   fd,
		height: 64,
		width:  128,
	}
}

func (d *SH1106_128_64) Initialize() error {
	fmt.Println("Initialize screen")

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

// Turn off OLED display
func (i *SH1106_128_64) DisplayOff() (int, error) {
	return i.conn.WriteCommand(OLED_CMD_DISPLAY_OFF)
}

// Turn on OLED display
func (i *SH1106_128_64) DisplayOn() (int, error) {
	return i.conn.WriteCommand(OLED_CMD_DISPLAY_ON)
}

func (i *SH1106_128_64) SetContrast(contrast int) error {
	var err error
	if contrast < 0 || contrast > 255 {
		return fmt.Errorf("contrast must be a values from 0 to 255")
	}
	if _, err = i.conn.WriteCommand(OLED_CMD_CONTRAST); err != nil {
		return err
	}
	_, err = i.conn.WriteCommand(byte(contrast))
	i.contrast = contrast
	return err
}

func (i *SH1106_128_64) SetDim(dim bool) error {
	contrast := i.contrast
	if !dim {
		if i.vccState == SH110X_DCDC {
			contrast = 0x9f
		} else {
			contrast = 0xCF
		}
	}
	return i.SetContrast(contrast)
}

func (i *SH1106_128_64) DrawImage(img *image.RGBA) {
	bounds := img.Bounds()
	if bounds.Max.X != i.width || i.height != bounds.Max.Y {
		panic(fmt.Sprintf("Error: Size of image is not %dx%d pixels.", i.width, i.height))
	}
	size := i.width * i.height / PIXSIZE
	data := make([]byte, size)
	for page := 0; page < i.height/8; page++ {
		for x := 0; x < i.width; x++ {
			bits := uint8(0)
			for bit := 0; bit < 8; bit++ {
				y := page*8 + 7 - bit
				if y < i.height {
					col := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
					if col.Y > 127 {
						bits = (bits << 1) | 1
					} else {
						bits = bits << 1
					}
				}
			}
			index := page*i.width + x
			data[index] = byte(bits)
		}
	}
}
