package helper

import (
	"io/ioutil"
	"errors"
	"github.com/bejohi/golbp/model"
	"strings"
	"fmt"
)

func GetAllJpgPathsFromFolderPath(folderPath string) (*[]string,error) {
	filesInfos, err := ioutil.ReadDir(folderPath)
	if err != nil{
		err = errors.New("GetAllJpgPathsFromFolderPath: " + err.Error())
	}

	imgPathsArr := []string {}

	for _, fileInfo := range filesInfos {
		fileName := fileInfo.Name()
		for _, jpgEnding := range *(model.GetAcceptedFileEndingsForJpg()){
			if strings.HasSuffix(fileName,jpgEnding){
				imgPathsArr = append(imgPathsArr, fileName)
				break
			}
		}
	}

	return &imgPathsArr, nil
}
