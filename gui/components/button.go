package components

import (
	"image/color"

	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type MyButton struct {
	ClickWidget  *widget.Clickable
	StyledButton material.ButtonStyle
}

func Button(th *material.Theme, text string) MyButton {
	var widget widget.Clickable
	button := material.Button(th, &widget, text)
	button.CornerRadius = unit.Dp(10)

	button.Background = color.NRGBA{0, 0, 0, 255}
	return MyButton{&widget, button}

}
