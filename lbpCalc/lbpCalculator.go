package lbpCalc

import (
	"image"
	"github.com/bejohi/golbp/model"
)

var pixelNeighboursY = []int{-1,-1,-1,0,1,1,1,0}
var pixelNeighboursX = []int{-1,0,1,1,1,0,-1,-1}

// createUniformMatrix creates a 2d binary matrix for the given uniform struct.
// A 'true' in a cell means, that the pixel lbp pattern matches one of the given lbp- uniform numbers from the
// uniform object. This function is rly fast, since it is not necessary to create an lbp matrix in one step
// and calculate the pattern in the next one (the speed is ~ 2x as fast as the other approach).
func createUniformMatrix(img *image.Gray16, uniform model.LbpUniform)*[][]bool {
	width := (*img).Bounds().Max.X
	height := (*img).Bounds().Max.Y
	uniformArray:= make([][]bool,height)
	for y := 0; y < height; y++{
		uniformArray[y] = make([]bool,width)
		for x := 0; x < width; x++{
			if !pointNotAtBorder(x,y,width,height){
				uniformArray[y][x] = false
				continue
			}
			lbpByte := calculateLbpPatternForPixel(img,x,y)
			uniformArray[y][x] = uniform.IsByteInUniform(lbpByte)
		}
	}
	return &uniformArray
}

// CreateLbpMatrix creates a 2d matrix of bytes from a given Gray16 image. Every byte in a cell represents a lbp
// value (e.g 0b01001010).
func CreateLbpMatrix(img *image.Gray16) *[][]byte{
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

// calculateLbpPatternForPixel compares an pixel at the given position with its 8 neighbours and returns the
// lbp pattern as a byte.
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


// pointNotAtBorder returns true if a given point is not a the matrix border (which is represented with its width
// and height)
func pointNotAtBorder(x int, y int, width int, height int) bool {
	if x <= 0 || y <= 0{
		return false
	}
	if x >= width -1 || y >= height -1 {
		return false
	}

	return true
}