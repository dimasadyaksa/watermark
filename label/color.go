package label

import (
	"image/color"
)

func WithColor(c color.Color) StringOption {
	return func(opt *strOption) {
		opt.c = c
	}
}
