package converter

import (
	"image"
	"image/color"

	"github.com/aybabtme/rgbterm"
)

type Converter struct {
	Image image.Image
}

type RgbaPixel struct {
	Char byte
	R    uint8
	G    uint8
	B    uint8
	A    uint8
}

func (c *Converter) ConvertImgToString(img image.Image) string {
	size := img.Bounds()
	width := size.Max.X
	height := size.Max.Y

	resultMatrix := make([][]RgbaPixel, 0, height)

	for i := 0; i < height; i++ {
		row := make([]RgbaPixel, 0, width)

		for j := 0; j < width; j++ {
			pixel := color.NRGBAModel.Convert(img.At(j, i))
			println(pixel)
			// convertedChar := c.convertPixelToChar(pixel)
			// row = append(row, convertedChar)

		}
		resultMatrix = append(resultMatrix, row)
	}

	return "hello world"
}

func (c *Converter) convertPixelToChar(pixel color.Color) {
	// TODO
	// RGBA -> intensity 명도 계산
	// 계산된 명도로 아스키코드를 맵핑 (자체 테이블)
	// 출력 (buffer 구현)
	// 테스트 코드 ->
	// resize? + 패키징
}

func (c *Converter) getColoredString(r, g, b uint8, rawChar byte) string {
	coloredString := rgbterm.FgString(string([]byte{rawChar}), uint8(r), uint8(g), uint8(b))
	return coloredString
}
