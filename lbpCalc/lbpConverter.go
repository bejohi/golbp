package lbpCalc

import (
	"github.com/bejohi/golbp/model"
	"image"
	"errors"
)

// ConvertGrayImgToLbpImg converts an Gray16 image to an lbp image, with 256 colors.
func ConvertGrayImgToLbpImg(imgWrapper *model.ImageWrapper) (*model.ImageWrapper,error){
	if grayImg, ok := (*imgWrapper).Img.(*image.Gray16); ok {
		lbpMatrix := CreateLbpMatrix(grayImg)
		imgWrapper.Img = createGrayImgFromByteMatrix(lbpMatrix)
		return imgWrapper,nil
	}
	return nil, errors.New("ConvertGrayImgToLbpImg: The given image was not a Gray16 image.")
}

// ConvertGrayImgToUniformImg converts an Gray16 image to an uniform image, with 1 color.
func ConvertGrayImgToUniformImg(imgWrapper *model.ImageWrapper, uniform model.LbpUniform) (*image.Gray,error){
	if grayImg, ok := (*imgWrapper).Img.(*image.Gray16); ok {
		uniformMatrix := createUniformMatrix(grayImg, uniform)
		return createBinaryImageFromBoolMatrix(uniformMatrix),nil
	}
	return nil, errors.New("ConvertGrayImgToUniformImg: The given image was not a Gray16 image.")
}