package gui

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"hodei.naiz/simplesynth/gui/components"
)

func Render(w *app.Window, controller *Controls) error {
	//the theme
	th := material.NewTheme(gofont.Collection())
	//fields
	var ops op.Ops
	//init

	selector := components.CreateSelector(th)
	//render
	for {

		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:

			components.SelectorCounter(selector.ButtonUp.ClickWidget, selector.ButtonDown.ClickWidget, controller.SelectorFunc)

			gtx := layout.NewContext(&ops, e)

			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

				return material.Label(th, unit.Dp(100), "hello Synth!").Layout(gtx)

			})
			layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return components.ShowSelector(th, gtx, &selector, controller.SelectorFunc)
			})
			e.Frame(gtx.Ops)

		}
	}
}
