package model

import (
	"os"
	"image"
	"errors"
)

type TypeNone struct {}

func (typePng TypeNone) Decode(file *os.File) (image.Image,error) {
	return nil, errors.New("NONE image type can not get decoded")
}

func (typePng TypeNone) Encode(file *os.File, img image.Image) error {
	return errors.New("NONE image type can not get encoded.")
}

func (typePng TypeNone) GetName() string{
	return "NONE"
}