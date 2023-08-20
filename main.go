package main

import (
	"github.com/Hokasonja/Image-to-text-converter/src/converter"
	"github.com/Hokasonja/Image-to-text-converter/src/file"
)

func main() {

	prefix := "public/images/"

	img, err := file.Upload(prefix + "test2.png")
	if err != nil {

		panic(err)
	}

	cvt := converter.Converter{Image: img}
	outputStr := cvt.ConvertImgToString(img)

	for i := 0; i < len(outputStr); i++ {
		println(outputStr[i])
	}

	// println(output.Output(outputStr))
}
