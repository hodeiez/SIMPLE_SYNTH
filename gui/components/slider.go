package components

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type MySlider struct {
	FloatWidget *widget.Float
	StyledSlide material.SliderStyle
	StyledLabel material.LabelStyle
}

func Slider(th *material.Theme, min float32, max float32, text string) MySlider {
	var widget widget.Float
	slider := material.Slider(th, &widget, min, max)
	slider.FingerSize = unit.Dp(100)
	label := material.Label(th, unit.Dp(10), text)
	slider.Color = color.NRGBA{0, 200, 0, 255}
	slider.Float.Axis = layout.Vertical

	return MySlider{&widget, slider, label}

}
