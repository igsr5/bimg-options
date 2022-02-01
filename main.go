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

	resize(src, 3000, 1000)
}

func resize(src *bimg.Image, w int, h int) {
	buf, err := src.Resize(w, h)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	size, err := src.Size()
	if size.Width != w && size.Height != h {
		fmt.Println("The image size is invalid")
	}

	bimg.Write(srcPngLocation, buf)
}
