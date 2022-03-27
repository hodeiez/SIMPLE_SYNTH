package components

import (
	"image/color"

	"gioui.org/layout"

	"gioui.org/widget"
	"gioui.org/widget/material"
)

type SliderPanel struct {
	Sliders    []MySlider
	PanelColor color.NRGBA
}

func SlidersAction(slider *widget.Float,
) float32 {
	return slider.Value
}

func ShowADSRPanel(th *material.Theme, gtx layout.Context, panel SliderPanel) layout.Dimensions {

	return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceEvenly}.Layout(gtx,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,

				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return panel.Sliders[0].StyledSlide.Layout(gtx)

				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {

					return panel.Sliders[0].StyledLabel.Layout(gtx)

				}))
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return panel.Sliders[1].StyledSlide.Layout(gtx)

				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {

					return panel.Sliders[1].StyledLabel.Layout(gtx)

				}))
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return panel.Sliders[2].StyledSlide.Layout(gtx)

				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {

					return panel.Sliders[2].StyledLabel.Layout(gtx)

				}))
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return panel.Sliders[3].StyledSlide.Layout(gtx)

				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {

					return panel.Sliders[3].StyledLabel.Layout(gtx)

				}))
		}),
	)
}
