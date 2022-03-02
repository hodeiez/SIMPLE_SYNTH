package gui

import (
	"image/color"

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
	//th.Palette = material.Palette{Bg: color.NRGBA{250, 50, 50, 255}}
	th.Bg = color.NRGBA{100, 100, 100, 255}
	th.ContrastBg = color.NRGBA{0, 0, 100, 255}
	th.ContrastFg = color.NRGBA{160, 100, 100, 255}
	th.Fg = color.NRGBA{100, 100, 100, 255}

	//fields
	var ops op.Ops
	//init

	selector := components.CreateSelector(th)
	sliders := []components.MySlider{components.Slider(th, 0, 10, "A"), components.Slider(th, 0, 10, "D"), components.Slider(th, 0, 10, "S"), components.Slider(th, 0, 10, "R")}
	adsrPanel := components.SliderPanel{Sliders: sliders, PanelColor: color.NRGBA{250, 250, 50, 255}}
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
				return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceEvenly}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {

						return material.Label(th, unit.Dp(50), "Amp+AM+FM").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Label(th, unit.Dp(20), "hello Synth!").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return components.ShowADSRPanel(th, gtx, adsrPanel)
						/* return material.Label(th, unit.Dp(50), "ADSR here").Layout(gtx) */
					}),
				)

			})
			layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return components.ShowSelector(th, gtx, &selector, controller.SelectorFunc)
			})

			e.Frame(gtx.Ops)

		}
	}
}
