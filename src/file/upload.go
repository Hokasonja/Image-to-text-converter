package file

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func Upload(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}
