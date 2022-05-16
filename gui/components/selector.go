package components

import (
	"gioui.org/layout"
	"gioui.org/text"

	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"hodei.naiz/simplesynth/synth/generator"
)

type Selector struct {
	ButtonUp   MyButton
	ButtonDown MyButton
	Display    material.LabelStyle
}

func CreateSelector(th *material.Theme) Selector {
	widget1 := &widget.Clickable{}
	widget2 := &widget.Clickable{}
	buttonUp := material.Button(th, widget1, "->")
	buttonDown := material.Button(th, widget2, "<-")

	setButtonStyle(&buttonUp)
	setButtonStyle(&buttonDown)

	display := material.Label(th, unit.Dp(10), "")

	return Selector{MyButton{widget1, buttonUp}, MyButton{widget2, buttonDown}, display}
}
func SelectorCounter(btnUp *widget.Clickable, btnDown *widget.Clickable, count *generator.MyWaveType) {

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

}
func ShowSelector(th *material.Theme, gtx layout.Context, raw *Selector, texto *generator.MyWaveType) layout.Dimensions {

	return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween, WeightSum: 20}.Layout(gtx,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return raw.ButtonDown.StyledButton.Layout(gtx)

		}),

		layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
			raw.Display.Text = texto.String()
			raw.Display.Alignment = text.Middle
			return raw.Display.Layout(gtx)

		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return raw.ButtonUp.StyledButton.Layout(gtx)

		}),
	)

}
func setButtonStyle(button *material.ButtonStyle) {
	button.TextSize = unit.Dp(10)

	button.Inset = layout.Inset{unit.Dp(10), unit.Dp(10), unit.Dp(10), unit.Dp(10)}

	button.Background = NewMyStyles().colorAccent
	button.CornerRadius = unit.Dp(10)

}
