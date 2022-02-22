package components

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Selector struct {
	Buttons []material.ButtonStyle
	Display material.LabelStyle
}

func create(th *material.Theme) Selector {
	var widget widget.Clickable
	button := material.Button(th, &widget, "-")
	button.CornerRadius = unit.Dp(10)

	button.Background = color.NRGBA{0, 0, 0, 255}
	display := material.Label(th, unit.Dp(10), "-")

	return Selector{[]material.ButtonStyle{button, button}, display}
}
func Show(th *material.Theme, gtx layout.Context) layout.Dimensions {
	raw := create(th)

	raw.Buttons[0].Text = "<-"
	raw.Buttons[1].Text = "->"
	return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return raw.Buttons[0].Layout(gtx)

		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return raw.Display.Layout(gtx)

		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return raw.Buttons[1].Layout(gtx)

		}),
	)
	/*---*/
}
