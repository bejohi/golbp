package lbpCalc

import (
	"github.com/bejohi/golbp/model"
	"image"
	"errors"
	"github.com/bejohi/golbp/imageCalc"
)

// ConvertGrayImgToLbpImg converts an Gray16 image to an lbp image (with 256) colors.
func ConvertGrayImgToLbpImg(imgWrapper *model.ImageWrapper) (*model.ImageWrapper,error){
	if grayImg, ok := (*imgWrapper).Img.(*image.Gray16); ok {
		lbpMatrix := CreateLbpMatrix(grayImg)
		imgWrapper.Img = imageCalc.CreateGrayImgFromByteMatrix(lbpMatrix)
		return imgWrapper,nil
	}
	return nil, errors.New("ConvertGrayImgToLbpImg: The given image was not a Gray16 image.")
}

func ConvertGrayImgToUniformImg(imgWrapper *model.ImageWrapper, uniform model.LbpUniform) (*model.ImageWrapper,error){
	if grayImg, ok := (*imgWrapper).Img.(*image.Gray16); ok {
		uniformMatrix := CreateUniformMatrix(grayImg, uniform)
		imgWrapper.Img = imageCalc.CreateBinaryImageFromBoolMatrix(uniformMatrix)
		return imgWrapper,nil
	}
	return nil, errors.New("ConvertGrayImgToUniformImg: The given image was not a Gray16 image.")
}