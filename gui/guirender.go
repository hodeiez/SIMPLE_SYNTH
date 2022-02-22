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

func Render(w *app.Window) error {
	//the theme
	th := material.NewTheme(gofont.Collection())
	//fields
	var ops op.Ops

	for {

		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

				return material.Label(th, unit.Dp(100), "hello Synth!").Layout(gtx)

			})
			layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return components.Show(th, gtx)
			})
			e.Frame(gtx.Ops)

		}
	}
}
