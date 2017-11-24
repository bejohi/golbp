package imageCalc

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"github.com/bejohi/golbp/helper"
	"strings"
	"errors"
	"github.com/bejohi/golbp/model"
)

// TODO [bejohi] This functions violates the open closed principle, do it better.
// LoadImage loads an image from the harddrive, if the format is supported.
// JPEG and PNG are supported.
// Returns a pointer to an Img object which holds the image itself and also an imgage config struct.
func LoadImage(imgPath string) (*model.ImageWrapper, error){

	imgType := getImageTypeByFileName(imgPath)

	if imgType == NONE {
		return nil, errors.New("The image format is not supported!")
	}

	file, err := os.Open(imgPath)

	if err != nil {
		helper.Log.Error("LoadImage: " + err.Error())
		return nil, err
	}

	var img image.Image
	var imgError error

	if imgType == JPEG {
		img, imgError = jpeg.Decode(file)
	}

	if imgType == PNG {
		img, imgError = png.Decode(file)
	}

	if imgError != nil {
		imgError = errors.New("LoadImage: " + imgError.Error())
	}
	return &model.ImageWrapper{Img:img,Type:imgType,FullPath:imgPath}, imgError
}

// TODO [bejohi] This functions violates the open closed principle, do it better.
// SaveImg writes an image to an file on the harddrive.
// The filePath string from the struct is used as fully qualified name of the new file.
func SaveImg(img *model.ImageWrapper, imgPath string) error {
	newImgFile, err := os.Create(imgPath)
	if err != nil {
		return err
	}

	defer newImgFile.Close()

	var imgError error = nil
	if img.Type == JPEG {
		helper.Log.Info("JPEG saved at " + imgPath)
		imgError = jpeg.Encode(newImgFile,img.Img,nil)
	} else if img.Type == PNG{
		imgError = png.Encode(newImgFile,img.Img)
	} else {
		return errors.New("The image type was not supported!")
	}

	if imgError != nil {
		imgError = errors.New("SaveImage: " + imgError.Error())
	}

	return imgError

}

func getImageTypeByFileName(imgName string) string {
	if strings.HasSuffix(imgName,".jpg") || strings.HasSuffix(imgName,".jpeg"){
		return JPEG
	}
	if strings.HasSuffix(imgName,".png"){
		return PNG
	}
	return NONE
}
