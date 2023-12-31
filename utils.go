package main

import (
	"strconv"
	"strings"
)

// hexToRGB converts a hex color string to an RGB tuple.
func HexToRGB(hexColor string) (Color, error) {
	// Remove '#' in case it's included in the hex color string
	hexColor = strings.TrimPrefix(hexColor, "#")

	// Convert the hex color to RGB
	r, err := strconv.ParseInt(hexColor[0:2], 16, 64)
	if err != nil {
		panic(err)
	}

	g, err := strconv.ParseInt(hexColor[2:4], 16, 64)
	if err != nil {
		panic(err)
	}

	b, err := strconv.ParseInt(hexColor[4:6], 16, 64)
	if err != nil {
		panic(err)
	}

	return Color{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 1,
	}, nil
}
