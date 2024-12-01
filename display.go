package goi2coled

import "image"

type Display interface {
	Initialize() error
	Width() int
	Height() int
	VCCState() byte
	DrawImage(image image.Image) error
}
