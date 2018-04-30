// Code generated by file2byteslice. DO NOT EDIT
// (gofmt is fine after generating)

package main

import (
	"log"
	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"
)

func decodeFont(b []byte, size float64) font.Face {
	tt, err := truetype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	f := truetype.NewFace(tt, &truetype.Options{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	return f
}
