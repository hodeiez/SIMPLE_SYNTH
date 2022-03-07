package gui

import (
	"image"
	"image/color"
	"strconv"

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

			bindControls(controller.ADSRcontrol, sliders)
			gtx := layout.NewContext(&ops, e)
			layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return marginCenter.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

					return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceEvenly}.Layout(gtx,

						layout.Rigid(func(gtx layout.Context) layout.Dimensions {

							return material.Body2(th, strconv.FormatFloat(float64(sliders[0].FloatWidget.Value), 'f', 2, 64)).Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {

							return components.ShowADSRPanel(th, gtx, adsrPanel)

						}),
					)

				})
			})
			layout.NE.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							circle := clip.Ellipse{f32.Point{sliders[2].FloatWidget.Value, sliders[3].FloatWidget.Value}, f32.Point{-1 * sliders[0].FloatWidget.Value, sliders[1].FloatWidget.Value}}.Op(gtx.Ops)
							color := color.NRGBA{R: 255, A: 255, B: 100}
							paint.FillShape(gtx.Ops, color, circle)
							d := image.Point{Y: 500}
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
func redTriangle(ops *op.Ops) {
	var path clip.Path
	path.Begin(ops)
	path.Move(f32.Pt(50, 0))
	path.Quad(f32.Pt(0, 90), f32.Pt(50, 100))
	path.Line(f32.Pt(-100, 0))
	path.Line(f32.Pt(50, -100))
	defer clip.Outline{Path: path.End()}.Op().Push(ops).Pop()
	drawRedRect(ops)
}
func drawRedRect(ops *op.Ops) {
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
