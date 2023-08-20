package converter

import (
	"image"
	"image/color"

	"github.com/aybabtme/rgbterm"
)

type Converter struct {
	Image image.Image
}

func (c *Converter) ConvertImgToString(img image.Image) []string {
	size := img.Bounds() //이미지의 경계값 정보를 불러온다.
	width := size.Max.X
	height := size.Max.Y

	resultMatrix := make([]string, 0, height)

	for i := 0; i < height; i++ {
		row := ""

		for j := 0; j < width; j++ {
			pixel := img.At(j, i)
			convertedChar := c.convertPixelToChar(pixel)
			row += convertedChar

		}
		resultMatrix = append(resultMatrix, row)
	}

	println(width, height)
	println(len(resultMatrix))

	return resultMatrix
}

func (c *Converter) convertPixelToChar(pixel color.Color) string {
	r, g, b, _ := pixel.RGBA()

	brightness := c.getBrightness(uint8(r), uint8(g), uint8(b))

	rawChar := c.getCharByBrightness(brightness)
	coloredChar := c.getColoredString(uint8(r), uint8(g), uint8(b), rawChar)

	return string(coloredChar)
}

func (c *Converter) getCharByBrightness(b uint8) byte {
	charIndex := " .,:;i1tfLCG08@"
	// charIndex := " .;=/+&$%@#"
	// charIndex := " .,:;ㅣㄴㄷㅁㅂ응뷃"

	i := b / 6

	k := charIndex[i]

	return k
}

func (c *Converter) getBrightness(r, g, b uint8) uint8 {
	return (r + g + b) / 3
}

func (c *Converter) getColoredString(r, g, b uint8, rawChar byte) []byte {
	coloredString := rgbterm.FgByte(rawChar, uint8(r), uint8(g), uint8(b))
	return coloredString
}
