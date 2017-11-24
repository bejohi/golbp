package lbpCalc

import (
	"github.com/bejohi/golbp/model"
	"image"
)

func ConvertGrayImgToLbpImg(imgWrapper *model.ImageWrapper, uniformList []byte) (*model.ImageWrapper,error){
	return nil, nil
}


func calculateLbpPattern(img *image.Image) *image.Image{
	width := (*img).Bounds().Max.X
	height := (*img).Bounds().Max.Y
	//lbpImg := image.NewGray(image.Rect(0,0,width,height))
	for y := 0; y < height; y++{
		for x := 0; x < width; x++{

		}
	}
}

func coordinateIsNotAtBorder(x int, y int, width int, height int) bool {
	if x <= 0 || y <= 0{
		return false
	}
	if x >= width -1 || y >= height -1 {
		return false
	}

	return true
}