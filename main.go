package main

import (
	"github.com/bejohi/golbp/cmd"
	"github.com/bejohi/golbp/model"
)

func main(){
	imgPath := "/home/spoof/Schreibtisch/neutral_front/007_03.jpg"
	newImgPath := "/home/spoof/Schreibtisch/neutral_front/007_03uniform.jpg"

	err := cmd.CreateUniformImgForGivenImg(imgPath,newImgPath,model.EndOfEdgeUniform{})

	if err != nil {
		panic(err)
	}
}
