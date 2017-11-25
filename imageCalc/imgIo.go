package imageCalc

import (
	"image"
	"os"
	"github.com/bejohi/golbp/helper"
	"strings"
	"errors"
	"github.com/bejohi/golbp/model"
)

// LoadImage loads an image from the harddrive, if the format is supported.
// JPEG and PNG are supported.
// Returns a pointer to an Img object which holds the image itself and also an imgage config struct.
func LoadImage(imgPath string) (*model.ImageWrapper, error){
	imgType := getImageTypeByFileName(imgPath)

	file, err := os.Open(imgPath)

	if err != nil {
		helper.LogError("LoadImage: " + err.Error())
		return nil, err
	}

	defer file.Close()

	var img image.Image
	var imgError error
	img, imgError = imgType.Decode(file)

	if imgError != nil {
		imgError = errors.New("LoadImage: " + imgError.Error())
	}

	return &model.ImageWrapper{Img:img,ImageType:imgType,FullPath:imgPath}, imgError
}

// SaveImg writes an image to an file on the harddrive.
// The filePath string from the struct is used as fully qualified name of the new file.
func SaveImg(imgWrapper *model.ImageWrapper, imgPath string) error {
	newImgFile, err := os.Create(imgPath)
	if err != nil {
		err = errors.New("SaveImg: " + err.Error())
		return err
	}

	defer newImgFile.Close()

	imgError := imgWrapper.ImageType.Encode(newImgFile,imgWrapper.Img)

	if imgError != nil {
		imgError = errors.New("SaveImg: " + imgError.Error())
	}

	return imgError
}

// getImageTypeByFileName is a naively implemented approach to detect the type of an given imageName.
func getImageTypeByFileName(imgName string) model.ImgType {
	for _, jpgEnding := range *model.GetAcceptedFileEndingsForJpg() {
		if strings.HasSuffix(imgName,jpgEnding){
			return model.TypeJpg{}
		}
	}
	for _, pngEnding := range *model.GetAcceptedFileEndingsForPng() {
		if strings.HasSuffix(imgName,pngEnding){
			return model.TypePng{}
		}
	}
	return model.TypeNone{}
}
