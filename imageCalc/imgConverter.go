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
