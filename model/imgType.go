package model

import (
	"os"
	"image"
)

const (
	NONE = "NOIMGTYPE"
	JPEG = "JPEG"
	PNG = "PNG"
)

type ImgType interface {
	Decode(file *os.File) (image.Image,error)
	Encode(file *os.File, img image.Image) error
	GetName() string
}