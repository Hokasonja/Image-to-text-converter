package main

import (
	"fmt"

	"github.com/Hokasonja/Image-to-text-converter/src/converter"
	"github.com/Hokasonja/Image-to-text-converter/src/file"
	"github.com/Hokasonja/Image-to-text-converter/src/output"
)

func hello() string {
	return "Hello HokaSonja"
}

func plus(a, b int) int {
	return a + b
}

func main() {

	fmt.Println(hello())
	fmt.Println(plus(2, 3))

	img, err := file.Upload("test2.png")
	if err != nil {
		panic(err)
	}

	cvt := converter.Converter{Image: img}
	outputStr := cvt.ConvertImgToString(img)
	output.Output([]string{outputStr})
}
