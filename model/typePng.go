package model

import (
	"os"
	"image"
	"image/png"
)

type TypePng struct {}

func (typePng TypePng ) Decode(file *os.File) (image.Image,error) {
	return png.Decode(file)
}

func (typePng TypePng ) Encode(file *os.File, img image.Image) error {
	return png.Encode(file,img)
}

func (typePng TypePng) GetName() string{
	return "PNG"
}

func GetAcceptedFileEndingsForPng() *[]string{
	fileEndingsArr := []string {"png"}
	return &fileEndingsArr
}