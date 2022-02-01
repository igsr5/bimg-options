package main

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

func main() {
	buffer, err := bimg.Read("img/src.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	src := bimg.NewImage(buffer)

	resize(src)
}

func resize(src *bimg.Image) {
	buf, err := src.Resize(800, 600)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	size, err := src.Size()
	if size.Width != 800 && size.Height != 600 {
		fmt.Println("The image size is invalid")
	}

	bimg.Write("src.png", buf)
}
