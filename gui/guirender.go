package gui

import (
	"image"
	"image/color"

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
	selector2 := components.CreateSelector(th)
	sliders := []components.MySlider{components.Slider(th, 1, 100000000.0, "A"), components.Slider(th, 1, 100000000.0, "D"), components.Slider(th, 0.0, 0.0099, "S"), components.Slider(th, 1, 100000000.0, "R")}
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
			components.SelectorCounter(selector2.ButtonUp.ClickWidget, selector2.ButtonDown.ClickWidget, controller.SelectorFunc2)

			bindControls(controller.ADSRcontrol, sliders)
			gtx := layout.NewContext(&ops, e)
			layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return marginCenter.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

					return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceEvenly}.Layout(gtx,

						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							components.AdsrImage(&ops, sliders)
							return components.ShowADSRPanel(th, gtx, adsrPanel)

						}),
					)

				})
			})
			layout.NE.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {

							d := image.Point{Y: 100, X: 100}
							//return strokeTriangle(&ops, d, gtx)
							return layout.Dimensions{Size: d}
						}),
				)
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
			//TODO: refactor to component
			layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

				return layout.Flex{Axis: layout.Vertical, Spacing: layout.Spacing(layout.Center), WeightSum: 20}.Layout(gtx,
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Spacing: layout.Spacing(255)}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return components.ShowSelector(th, gtx, &selector, controller.SelectorFunc)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Label(th, unit.Dp(10), "Pitch").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Label(th, unit.Dp(10), "VOLUME?").Layout(gtx)
							}),
						)
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Spacing: layout.Spacing(255)}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return components.ShowSelector(th, gtx, &selector2, controller.SelectorFunc2)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Label(th, unit.Dp(10), "Pitch").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Label(th, unit.Dp(10), "VOLUME?").Layout(gtx)
							}),
						)
					}),
				)

			})

			e.Frame(gtx.Ops)

		}
	}
}
func bindControls(controller *generator.ADSRControl, sliders []components.MySlider) {
	*controller.AttackTime = float64(sliders[0].FloatWidget.Value)
	*controller.DecayTime = float64(sliders[1].FloatWidget.Value)
	*controller.SustainAmp = float64(sliders[2].FloatWidget.Value)
	*controller.ReleaseTime = float64(sliders[3].FloatWidget.Value)
}
