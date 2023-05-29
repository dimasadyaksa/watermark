package watermark

import (
	"image/draw"
)

type DrawFunc func(img draw.Image) error
