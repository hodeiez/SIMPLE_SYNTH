package gui

import (
	"image"
	"image/color"
	"log"
	"os"

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

func Render(w *app.Window, controller *generator.Controls, test chan float64) error {
	//the theme
	th := material.NewTheme(gofont.Collection())

	//fields
	var ops op.Ops
	//init
	selector := components.CreateSelector(th)
	//A->100000000.0 //D->100000000.0
	sliders := []components.MySlider{components.Slider(th, 1, 1000000000.0, "A"), components.Slider(th, 1000000, 1000000000.0 /* 100000000000.0 */, "D"), components.Slider(th, 0.0, 0.0099, "S"), components.Slider(th, 1, 500000000.0, "R")}
	adsrPanel := components.SliderPanel{Sliders: sliders, PanelColor: color.NRGBA{250, 250, 50, 255}}

	marginCenter := layout.Inset{Top: unit.Dp(100),
		Bottom: unit.Dp(50),
		Right:  unit.Dp(0),
		Left:   unit.Dp(0)}
	marginOscPanels := layout.Inset{Top: unit.Dp(0),
		Bottom: unit.Dp(0),
		Right:  unit.Dp(0),
		Left:   unit.Dp(30)}

	slider := components.Slider(th, -60.0, 0.0, "pitch")
	oscPanel1 := components.NewOscPanel(selector, controller.SelectorFunc, slider)
	//
	//render

	for {

		//close(test)
		e := <-w.Events()
		switch e := e.(type) {

		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			if slider.FloatWidget.Changed() {
				test <- float64(slider.FloatWidget.Value)

				// select {
				// case test <- float64(slider.FloatWidget.Value):
				// 	log.Println("sending")
				// default:
				// 	log.Println("is default")
				// }
			}
			components.SelectorCounter(oscPanel1.WaveSelector.ButtonUp.ClickWidget, oscPanel1.WaveSelector.ButtonDown.ClickWidget, oscPanel1.WaveType)

			bindControls(controller.ADSRcontrol, sliders, controller, slider)

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
						return material.Label(th, unit.Dp(20), "SIMPLE SYNTH").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Label(th, unit.Dp(10), "by hodei").Layout(gtx)
					}),
				)

			})

			layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return marginOscPanels.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical, Spacing: layout.Spacing(layout.Center), WeightSum: 20}.Layout(gtx,
						oscPanel1.Render(th),
					)

				})
			})
			e.Frame(gtx.Ops)

		}

	}
}
func bindControls(controller *generator.ADSRControl, sliders []components.MySlider, pitch *generator.Controls, pitchSlide components.MySlider) {
	*controller.AttackTime = float64(sliders[0].FloatWidget.Value)
	*controller.DecayTime = float64(sliders[1].FloatWidget.Value)
	*controller.SustainAmp = float64(sliders[2].FloatWidget.Value)
	*controller.ReleaseTime = float64(sliders[3].FloatWidget.Value)
	*pitch.Pitch = float64(pitchSlide.FloatWidget.Value)

}
func Run(controller *generator.Controls, test chan float64) {
	w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(600)), app.Title("Symple synth"))

	err := Render(w, controller, test)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
