package imageCalc

import (
	"image"
	"image/color"
	"github.com/bejohi/golbp/model"
)

// TODO [bejohi] make this more abstract
// ConvertRgbaToGrayImg converts a given rgba image to an grayscale image and returns the new image.
func ConvertRgbaToGrayImg(img *model.ImageWrapper) *model.ImageWrapper {
	width := (*img).Img.Bounds().Max.X
	height := (*img).Img.Bounds().Max.Y
	grayImg := image.NewGray(image.Rect(0,0,width,height))
	for y := 0; y < height; y++{
		for x := 0; x < width; x++{
			oldPixel := (*img).Img.At(x,y)
			//r,g,b,_:= oldPixel.RGBA()
			//grayValue := float64(r) * 0.299 + float64(g) * 0.587 + float64(b) * 0.114
			pixel := color.Gray16Model.Convert(oldPixel)
			grayImg.Set(x,y,pixel)
		}
	}
	(*img).Img = grayImg
	return img
}
