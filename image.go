package watermark

import (
	"errors"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

const (
	ImageJPG = iota
	ImagePNG
)

func WithImage(wm image.Image, x, y int) DrawFunc {
	return func(img draw.Image) error {
		offset := image.Pt(x, y)

		draw.Draw(img, img.Bounds(), img, image.Point{}, draw.Src)
		draw.Draw(img, wm.Bounds().Add(offset), wm, image.Point{}, draw.Over)

		return nil
	}
}

func ImageFromFile(path string, imgType int) (image.Image, error) {
	img, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	switch imgType {
	case ImageJPG:
		return jpeg.Decode(img)
	case ImagePNG:
		return png.Decode(img)
	}

	return nil, errors.New("invalid image type")
}
