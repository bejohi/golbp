package imageCalc

import (
	"image"
	"image/color"
	"github.com/bejohi/golbp/model"
)

// ConvertRgbaToGrayImg converts a given rgba image to an grayscale image and returns the new image.
func ConvertRgbaToGrayImg(img *model.ImageWrapper) *model.ImageWrapper {
	width := (*img).Img.Bounds().Max.X
	height := (*img).Img.Bounds().Max.Y
	grayImg := image.NewGray16(image.Rect(0,0,width,height))
	for y := 0; y < height; y++{
		for x := 0; x < width; x++{
			oldPixel := (*img).Img.At(x,y)
			pixel := color.Gray16Model.Convert(oldPixel)
			grayImg.Set(x,y,pixel)
		}
	}
	(*img).Img = grayImg
	return img
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