package ssd1306

import (
	"fmt"
	"image"
	"image/color"

	"github.com/sandbankdisperser/go-i2c-oled/i2c"
)

type SSD1306_128_32 struct {
	conn *i2c.I2c

	vccState byte
	width    int
	height   int
	contrast int
}

func (d *SSD1306_128_32) VCCState() byte {
	return d.vccState
}
func (d *SSD1306_128_32) Height() int {
	return 32
}
func (d *SSD1306_128_32) Width() int {
	return 128
}

// NewSSD1306_128_32 creates a new instance of the SSD1306_128_32 structure.
func NewSSD1306_128_32(fd *i2c.I2c, vccstate byte) *SSD1306_128_32 {
	return &SSD1306_128_32{
		height:   32,
		width:    128,
		conn:     fd,
		vccState: vccstate,
	}
}

func (d *SSD1306_128_32) Initialize() error {
	fmt.Println("Initialize screen")
	data := []byte{
		SSD1306_DISPLAYOFF,         // 0xAE
		SSD1306_SETDISPLAYCLOCKDIV, // 0xD5
		0x80,                       // the suggested ratio 0x80
		SSD1306_SETMULTIPLEX,       // 0xA8
		0x1F,                       // Multiplex value for 128x32
		SSD1306_SETDISPLAYOFFSET,   // 0xD3
		0x0,                        // no offset
		SSD1306_SETSTARTLINE | 0x0, // line #0
		SSD1306_CHARGEPUMP,         // 0x8D
	}

	// Adjust charge pump settings based on vccstate.
	if d.vccState == SSD1306_EXTERNALVCC {
		data = append(data, byte(0x10)) // External Vcc
	} else {
		data = append(data, byte(0x14)) // Internal Vcc
	}

	// Additional setup commands.
	data = append(data, []byte{
		SSD1306_MEMORYMODE,     // 0x20
		0x00,                   // 0x0 act like ks0108
		SSD1306_SEGREMAP | 0x1, // Map segment 0 to column 127
		SSD1306_COMSCANDEC,     // Scan in descending order
		SSD1306_SETCOMPINS,     // 0xDA
		0x02,                   // Sequential COM pin configuration
		SSD1306_SETCONTRAST,    // 0x81
		0x8F,                   // Contrast value
		SSD1306_SETPRECHARGE,   // 0xd9
	}...)

	// Precharge period based on vccstate.
	if d.vccState == SSD1306_EXTERNALVCC {
		data = append(data, byte(0x22)) // External Vcc
	} else {
		data = append(data, byte(0xF1)) // Internal Vcc
	}

	// Final setup commands.
	data = append(data, []byte{
		SSD1306_SETVCOMDETECT,       // 0xDB
		0x40,                        // VCOM deselect level
		SSD1306_DISPLAYALLON_RESUME, // 0xA4
		SSD1306_NORMALDISPLAY,       // 0xA6
	}...)

	return sendCommands(d.conn, data...)
}
func (i *SSD1306_128_32) SetContrast(contrast int) error {
	var err error
	if contrast < 0 || contrast > 255 {
		return fmt.Errorf("Contrast must be a values from 0 to 255")
	}
	if _, err = i.conn.WriteCommand(OLED_CMD_CONTRAST); err != nil {
		return err
	}
	_, err = i.conn.WriteCommand(byte(contrast))
	i.contrast = contrast
	return err
}

func (i *SSD1306_128_32) SetDim(dim bool) error {
	contrast := i.contrast
	if !dim {
		if i.vccState == SSD1306_EXTERNALVCC {
			contrast = 0x9f
		} else {
			contrast = 0xCF
		}
	}
	return i.SetContrast(contrast)
}

func (i *SSD1306_128_32) DrawImage(img *image.RGBA) {
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
