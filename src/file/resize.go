package file

import (
	"github.com/nfnt/resize"
	"image"
)

const max = 100
const min = 50

func Resize(image image.Image) image.Image {
	p := image.Bounds().Size()
	h, w := p.Y, p.X

	if h > max || w > max {
		var ratio float64

		if h > w {
			ratio = max / float64(h)
		} else {
			ratio = max / float64(w)
		}

		h, w = int(float64(h)*ratio), int(float64(w)*ratio)
	}

	if h < min || w < min {
		var ratio float64

		if h < w {
			ratio = max / float64(h)
		} else {
			ratio = max / float64(w)
		}

		h, w = int(float64(h)*ratio), int(float64(w)*ratio)
	}

	return resize.Resize(uint(h), uint(w), image, resize.Lanczos3)
}
