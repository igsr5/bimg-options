package main

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

var srcPngLocation = "img/src.png"

func main() {
	buffer, err := bimg.Read("img/src_origin.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	src := bimg.NewImage(buffer)

	extract(src, 0, 0, 3000, 2000)
}

// --------------------
// Resize
// options: [width, heigth]
// --------------------
func resize(src *bimg.Image, w int, h int) {
	buf, err := src.Resize(w, h)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write(srcPngLocation, buf)
}

// --------------------
// Extract
// options: [top, left, width, heigth]
// --------------------
func extract(src *bimg.Image, top int, left int, w int, h int) {
	buf, err := src.Extract(top, left, w, h)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write(srcPngLocation, buf)
}
