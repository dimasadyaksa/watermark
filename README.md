# Go Watermark Generator

## Example

### Add watermark with default font
```go
f, err := os.Open("image.jpeg")
if err != nil {
	panic(err)
}

img, _ := jpeg.Decode(f)
	
newImg, err := watermark.Draw(img,
	label.String("This is watermark", 10, 600), 
)
if err != nil {
	panic(err)
}

f, err = os.Create("output.png")
if err != nil {
	panic(err)
}

err = png.Encode(f, newImg)
if err != nil {
	panic(err)
}
```

### Add watermark with custom font
```go
f, err := os.Open("image.jpeg")
if err != nil {
	panic(err)
}

img, _ := jpeg.Decode(f)

font, err := label.FontFromFile("font-file.ttf")
if err != nil {
    panic(err)
}

dpi := 27
size := 32
face := label.FontFace(font, dpi, size)

newImg, err := watermark.Draw(img,
    label.String("This is Watermark", 10, 600,
        label.WithFontFace(face),
    ),
)
if err != nil {
	panic(err)
}

f, err = os.Create("output.png")
if err != nil {
	panic(err)
}

err = png.Encode(f, newImg)
if err != nil {
	panic(err)
}
```

### Add image watermark
```go
f, err := os.Open("image.jpeg")
if err != nil {
	panic(err)
}

img, _ := jpeg.Decode(f)

// Load watermark image
wm, err := watermark.ImageFromFile("watermark.png", watermark.ImagePNG)
if err != nil {
    panic(err)
}

newImg, err := watermark.Draw(img,
    watermark.WithImage(wm, 0, 0),
)
if err != nil {
	panic(err)
}

f, err = os.Create("output.png")
if err != nil {
	panic(err)
}

err = png.Encode(f, newImg)
if err != nil {
	panic(err)
}
```