package label

import (
	"github.com/dimasadyaksa/watermark"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
)

type strOption struct {
	c    color.Color
	face font.Face
}

type StringOption func(opt *strOption)

func defaultStringOption() StringOption {
	return func(opt *strOption) {
		opt.c = color.RGBA{A: 0xff}
		opt.face = basicfont.Face7x13
	}
}

func String(label string, x, y int, options ...StringOption) watermark.DrawFunc {
	return func(img draw.Image) error {
		var opt strOption
		defaultStringOption()(&opt)

		for _, op := range options {
			op(&opt)
		}

		d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(opt.c),
			Face: opt.face,
			Dot:  fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)},
		}

		d.DrawString(label)

		return nil
	}
}
