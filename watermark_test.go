package watermark_test

import (
	"github.com/dimasadyaksa/watermark"
	"github.com/dimasadyaksa/watermark/label"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	white, _ := watermark.ColorFromHex("#ff0")
	fon, err := label.FontFromFile("label/lato/Lato-Bold.ttf")
	if err != nil {
		t.Fatal(err)
	}

	face := label.FontFace(fon, 72, 80)

	newImg, err := watermark.Draw(img,
		label.String("FAHMI", 10, 200,
			label.WithFontFace(face),
			label.WithColor(white),
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create("wm-out.png")
	if err != nil {
		t.Fatal(err)
	}

	err = png.Encode(f, newImg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDrawWithFile(t *testing.T) {
	f, err := os.Open("logo.jpeg")
	if err != nil {
		t.Fatal(err)
	}

	img, _ := jpeg.Decode(f)
	wm, err := watermark.ImageFromFile("wm-out.png", watermark.ImagePNG)
	if err != nil {
		t.Fatal(err)
	}

	font, err := label.FontFromFile("font-file.ttf")
	if err != nil {
		panic(err)
	}

	face := label.FontFace(font, 27, 32)

	newImg, err := watermark.Draw(img,
		watermark.WithImage(wm, 0, 0),
		label.String("Glenn", 10, 600,
			label.WithFontFace(face),
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	f, err = os.Create("img-wm-out.png")
	if err != nil {
		t.Fatal(err)
	}

	err = png.Encode(f, newImg)
	if err != nil {
		t.Fatal(err)
	}
}
