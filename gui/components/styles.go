package components

import "image/color"

type MyStyles struct {
	colorAccent color.NRGBA
}

func NewMyStyles() *MyStyles {
	colorAccent := color.NRGBA{R: 255, A: 255, B: 100}
	return &MyStyles{colorAccent: colorAccent}
}
