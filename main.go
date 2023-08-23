package main

import (
	"fmt"
	"github.com/Hokasonja/Image-to-text-converter/src/converter"
	"github.com/Hokasonja/Image-to-text-converter/src/file"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Print("Please input image file name.")
	}

	imageFileName := args[1]

	img, err := file.Upload(imageFileName)
	if err != nil {
		panic(err)
	}

	resizedImg := file.Resize(img)

	cvt := converter.Converter{Image: resizedImg}
	outputStr := cvt.ConvertImgToString(resizedImg)

	if err := file.Create(imageFileName, outputStr); err != nil {
		panic(err)
	}

	for i := 0; i < len(outputStr); i++ {
		println(outputStr[i])
	}
}
