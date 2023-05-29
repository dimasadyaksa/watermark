package watermark

import (
	"errors"
	"image"
	"image/draw"
)

type Watermark struct {
	img image.Image
}

func Draw(img image.Image, drawers ...DrawFunc) (image.Image, error) {
	var drawable draw.Image

	switch v := img.(type) {
	case *image.YCbCr:
		b := v.Bounds()
		drawable = image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(drawable, drawable.Bounds(), v, b.Min, draw.Src)
	case *image.RGBA:
		drawable = v
	default:
		return nil, errors.New("invalid image")
	}

	for _, fn := range drawers {
		err := fn(drawable)
		if err != nil {
			return nil, err
		}
	}

	return drawable, nil
}
