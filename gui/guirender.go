package gui

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/op/clip"
	"gioui.org/op/paint"

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
	sliders := []components.MySlider{components.Slider(th, 1, 100000000.0, "A"), components.Slider(th, 1, 100000000.0, "D"), components.Slider(th, 0.000001, 0.0099, "S"), components.Slider(th, 1, 100000000.0, "R")}
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

			bindControls(controller.ADSRcontrol, sliders)
			gtx := layout.NewContext(&ops, e)
			layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return marginCenter.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

					return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceEvenly}.Layout(gtx,

						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							adsrImage(&ops, sliders)
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
			layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return components.ShowSelector(th, gtx, &selector, controller.SelectorFunc)
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

//TODO: move to components and fix values (add beziers?)
func adsrImage(ops *op.Ops, sliders []components.MySlider) {
	var path clip.Path
	attackVal := float32(10) + sliders[0].FloatWidget.Value/2000000
	decayVal := attackVal + (sliders[1].FloatWidget.Value / 2000000)
	susamp := -20 - (sliders[2].FloatWidget.Value * 6000)
	releaseVal := attackVal + decayVal + (sliders[3].FloatWidget.Value / 20000000)

	path.Begin(ops)
	path.MoveTo(f32.Pt(0, -20))
	path.LineTo(f32.Pt(attackVal, -80))
	path.LineTo(f32.Pt(decayVal, susamp))
	path.LineTo(f32.Pt(attackVal+decayVal, susamp))
	path.LineTo(f32.Pt(releaseVal, -20))
	color := color.NRGBA{R: 255, A: 255, B: 100}
	paint.FillShape(ops, color,
		clip.Stroke{
			Path:  path.End(),
			Width: 8,
		}.Op())

}
