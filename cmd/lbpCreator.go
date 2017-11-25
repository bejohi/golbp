package cmd

import (
	"github.com/bejohi/golbp/imageCalc"
	"errors"
	"github.com/bejohi/golbp/lbpCalc"
)

// CreateLbpImgForGivenImg loads an given rgb image from the disk, converts it to an lbp image and saves the newly
// created image at the newImgPath location.
func CreateLbpImgForGivenImg(imgPath string, newImgPath string) error {
	imgWrapper, err := imageCalc.LoadImage(imgPath)

	if err != nil {
		err = errors.New("CreateLbpImgForGivenImg: " + err.Error())
		return err
	}

	grayImg:= imageCalc.ConvertRgbaToGrayImg(imgWrapper)

	lbpImg, lbpErr := lbpCalc.ConvertGrayImgToLbpImg(grayImg)

	if lbpErr != nil {
		lbpErr = errors.New("CreateLbpImgForGivenImg: " + lbpErr.Error())
		return lbpErr
	}

	return imageCalc.SaveImg(lbpImg, newImgPath)
}
