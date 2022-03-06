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

/*func CreateSliderPanel(th *material.Theme) Selector {
 widget1 := &widget.Clickable{}
	widget2 := &widget.Clickable{}
	buttonUp := material.Button(th, widget1, "->")
	buttonDown := material.Button(th, widget2, "<-")
	buttonDown.CornerRadius = unit.Dp(10)
	buttonUp.Background = color.NRGBA{0, 0, 0, 255}
	buttonDown.Background = color.NRGBA{0, 0, 0, 255}
	display := material.Label(th, unit.Dp(10), "")

	return Selector{MyButton{widget1, buttonUp}, MyButton{widget2, buttonDown}, display}

}*/
//copy pasted from selector
/* func SlidersAction(btnUp *widget.Clickable, btnDown *widget.Clickable, count *generator.MyWaveType) {

	if btnUp.Clicked() {
		if *count == generator.MyWaveTypeSize-1 {
			*count = 0
		} else {
			*count++
		}
	} else if btnDown.Clicked() {
		if *count <= 0 {
			*count = generator.MyWaveTypeSize - 1
		} else {
			*count--
		}
	}

} */
func SlidersAction(slider *widget.Float,
) float32 {
	return slider.Value
}

/* func SlidersAction2(slider []*widget.Float,
) []float32 {

	return slider.Value
} */

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
