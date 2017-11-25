package model

import (
	"image"
)

// ImageWrapper stores an image.Image itself and some additional information about it.
type ImageWrapper struct {
	Img      	image.Image
	FullPath 	string
	ImageType	ImgType
}
