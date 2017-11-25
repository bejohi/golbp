package lbpCalc

import (
	"github.com/bejohi/golbp/model"
	"image"
	"errors"
	"image/color"
)

// ConvertGrayImgToLbpImg converts an Gray16 image to an lbp image (with 256) colors.
func ConvertGrayImgToLbpImg(imgWrapper *model.ImageWrapper) (*model.ImageWrapper,error){
	if grayImg, ok := (*imgWrapper).Img.(*image.Gray16); ok {
		lbpMatrix := CreateLbpMatrix(grayImg)
		imgWrapper.Img = CreateGrayImgFromByteMatrix(lbpMatrix)
		return imgWrapper,nil
	}
	return nil, errors.New("ConvertGrayImgToLbpImg: The given image was not a Gray16 image.")
}

//CreateGrayImgFromByteMatrix transforms a given 2d byte matrix into a gray picture and returns it.
func CreateGrayImgFromByteMatrix(matrix *[][]byte) *image.Gray{
	width := len((*matrix)[0])
	height := len(*matrix)
	grayImg := image.NewGray(image.Rect(0,0,width,height))
	for y := 0; y < height; y++{
		for x := 0; x < width; x++ {
			colorValue := color.Gray{(*matrix)[y][x]}
			grayImg.Set(x,y,colorValue)
		}
	}
	return grayImg
}

