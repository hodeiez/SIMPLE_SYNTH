package components

import (
	"image/color"

	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type MySlider struct {
	FloatWidget  *widget.Float
	StyledButton material.SliderStyle
}

func Slider(th *material.Theme, min float32, max float32) MySlider {
	var widget widget.Float
	slider := material.Slider(th, &widget, min, max)
	slider.FingerSize = unit.Dp(100)
	slider.Color = color.NRGBA{0, 200, 0, 255}
	return MySlider{&widget, slider}

}
