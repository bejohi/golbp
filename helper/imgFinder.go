package helper

import (
	"io/ioutil"
	"errors"
	"github.com/bejohi/golbp/model"
	"strings"
)

func GetAllJpgPathsFromFolderPath(folderPath string) (*[]string,error) {
	filesInfos, err := ioutil.ReadDir(folderPath)
	if err != nil{
		err = errors.New("GetAllJpgPathsFromFolderPath: " + err.Error())
	}

	var imgPathsArr []string

	for _, fileInfo := range filesInfos {
		fileName := fileInfo.Name()
		for _, jpgEnding := range *(model.GetAcceptedFileEndingsForJpg()){
			if strings.HasSuffix(fileName,jpgEnding){
				imgPathsArr = append(imgPathsArr, folderPath +"/" + fileName)
				break
			}
		}
	}

	return &imgPathsArr, nil
}

// AppendStringToImgName is a very dangerous (!) and naively implemented approach to append a substring to a filename.
func AppendStringToFileName(imgPath string, appendString string) string {
	rightPartOfDot := strings.Split(imgPath,".")[0]
	separatedBySlashes := strings.Split(rightPartOfDot,"/")
	imgName := separatedBySlashes[len(separatedBySlashes)-1]
	return strings.Replace(imgPath,imgName,imgName + appendString,1)
}
