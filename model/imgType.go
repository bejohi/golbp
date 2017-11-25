package model

import (
	"os"
	"image"
)

// ImgType provides Methods to run image type unique operations.
type ImgType interface {
	Decode(file *os.File) (image.Image,error)
	Encode(file *os.File, img image.Image) error
	GetName() string
}