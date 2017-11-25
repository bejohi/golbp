package model

import (
	"os"
	"image"
	"image/jpeg"
)

type TypeJpg struct {}

func (typeJpg TypeJpg) Decode(file *os.File) (image.Image,error) {
	return jpeg.Decode(file)
}

func (typeJpg TypeJpg) Encode(file *os.File, img image.Image) error {
	return jpeg.Encode(file,img,nil)
}

func (typeJpg TypeJpg) GetName() string{
	return "JPG"
}

func GetAcceptedFileEndingsForJpg() *[]string{
	fileEndingsArr := []string {"jpg","jpeg"}
	return &fileEndingsArr
}