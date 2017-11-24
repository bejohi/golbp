package model

import (
	"image"
)

type ImageWrapper struct {
	Img      image.Image
	FullPath string
	Type     string
}
