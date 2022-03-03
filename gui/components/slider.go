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
	slider.FingerSize = unit.Dp(50)

	label := material.Label(th, unit.Dp(50), text)

	slider.Color = color.NRGBA{54, 68, 113, 200}
	slider.Float.Axis = layout.Vertical

	return MySlider{&widget, slider, label}

}
