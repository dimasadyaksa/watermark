package label

import (
	"bytes"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"io"
	"os"
)

func WithFontFace(f font.Face) StringOption {
	return func(opt *strOption) {
		opt.face = f
	}
}

func FontFace(f *truetype.Font, dpi, size float64) font.Face {
	return truetype.NewFace(f, &truetype.Options{
		Size: size,
		DPI:  dpi,
	})
}

func FontFromFile(path string) (*truetype.Font, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}

	ttfFont, err := truetype.Parse(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return ttfFont, nil
}
