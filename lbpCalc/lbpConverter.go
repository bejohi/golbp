package lbpCalc

import (
	"github.com/bejohi/golbp/model"
	"image"
	"errors"
	"image/color"
)

var pixelNeighboursY = []int{-1,-1,-1,0,1,1,10}
var pixelNeighboursX = []int{-1,0,1,1,1,0,-1,-1}

// ConvertGrayImgToLbpImg converts an Gray16 image to an lbp image (with 256) colors.
func ConvertGrayImgToLbpImg(imgWrapper *model.ImageWrapper, uniformList []byte) (*model.ImageWrapper,error){
	if grayImg, ok := (*imgWrapper).Img.(*image.Gray16); ok {
		lbpMatrix := createlbpMatrix(grayImg)
		imgWrapper.Img = createImgFromByteMatrix(lbpMatrix)
		return imgWrapper,nil
	}
	return nil, errors.New("ConvertGrayImgToLbpImg: The given image was not an Gray16 image.")
}

func createImgFromByteMatrix(matrix *[][]byte) *image.Gray{
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

func createlbpMatrix(img *image.Gray16) *[][]byte{
	width := (*img).Bounds().Max.X
	height := (*img).Bounds().Max.Y
	lbpArray := make([][]byte,height)
	for y := 0; y < height; y++{
		lbpArray[y] = make([]byte,width)
		for x := 0; x < width; x++{
			if !pointNotAtBorder(x,y,width,height){
				lbpArray[y][x] = 0
				continue
			}
			lbpArray[y][x] = calculateLbpPatternForPixel(img,x,y)
		}
	}
	return &lbpArray
}

func calculateLbpPatternForPixel(img *image.Gray16,x int,y int) byte {
	var pattern byte = 0
	for pos := uint(0); pos < 8; pos++{
		yNeighbourPos := y + pixelNeighboursY[pos]
		xNeighbourPos := x + pixelNeighboursX[pos]

		neigbourValue := (*img).Gray16At(xNeighbourPos,yNeighbourPos)
		centerValue := (*img).Gray16At(x,y)
		if neigbourValue.Y >= centerValue.Y{
			// Set the bit at the current position to 1
			pattern |= 1 << pos
		} else {
			// 0
			pattern &= ^(1 << pos)
		}


	}

	return pattern
}

func pointNotAtBorder(x int, y int, width int, height int) bool {
	if x <= 0 || y <= 0{
		return false
	}
	if x >= width -1 || y >= height -1 {
		return false
	}

	return true
}