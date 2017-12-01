package lbpCalc

import (
	"image"
	"errors"
	"github.com/bejohi/golbp/model"
)

// GetUniformImage loads an image from the disk, converts it to a gray image with binary uniform values and returns it.
func GetUniformImage(imgPath string, uniform model.LbpUniform) (*image.Gray, error){
	imgWrapper, err := loadImage(imgPath)

	if err != nil {
		err = errors.New("CreateLbpImgForGivenImg: " + err.Error())
		return nil,err
	}
	grayImg:= ConvertRgbaToGrayImg(imgWrapper)

	return ConvertGrayImgToUniformImg(grayImg,uniform)
}
