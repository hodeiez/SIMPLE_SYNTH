package gui

import (
	"image/color"
	"strconv"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"

	"gioui.org/unit"
	"gioui.org/widget/material"
	"hodei.naiz/simplesynth/gui/components"
	"hodei.naiz/simplesynth/synth/generator"
)

func Render(w *app.Window, controller *generator.Controls) error {
	//the theme
	th := material.NewTheme(gofont.Collection())

	//fields
	var ops op.Ops
	//init

	selector := components.CreateSelector(th)
	sliders := []components.MySlider{components.Slider(th, 1, 1000.0, "A"), components.Slider(th, 1, 1000.0, "D"), components.Slider(th, 0.00, 1, "S"), components.Slider(th, 1, 1000.0, "R")}
	adsrPanel := components.SliderPanel{Sliders: sliders, PanelColor: color.NRGBA{250, 250, 50, 255}}

	marginCenter := layout.Inset{Top: unit.Dp(100),
		Bottom: unit.Dp(50),
		Right:  unit.Dp(0),
		Left:   unit.Dp(0)}

	//render
	for {

		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:

			components.SelectorCounter(selector.ButtonUp.ClickWidget, selector.ButtonDown.ClickWidget, controller.SelectorFunc)
			//	slideValue := components.SlidersAction(sliders[0].FloatWidget)

			gtx := layout.NewContext(&ops, e)
			layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return marginCenter.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceEvenly}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {

							*controller.ADSRcontrol.AttackTime = float64(sliders[0].FloatWidget.Value)
							*controller.ADSRcontrol.DecayTime = float64(sliders[1].FloatWidget.Value)
							*controller.ADSRcontrol.SustainAmp = float64(sliders[2].FloatWidget.Value)
							*controller.ADSRcontrol.ReleaseTime = float64(sliders[3].FloatWidget.Value)

							return material.Body2(th, strconv.FormatFloat(float64(sliders[0].FloatWidget.Value), 'f', 2, 64)).Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return components.ShowADSRPanel(th, gtx, adsrPanel)

						}),
					)

				})
			})
			layout.SW.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {

						return material.Label(th, unit.Dp(20), "THE SIMPLE SYNTH").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Label(th, unit.Dp(10), "by hodei").Layout(gtx)
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
