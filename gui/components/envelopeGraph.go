package components

import (
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func AdsrImage(ops *op.Ops, sliders []MySlider) {
	var path clip.Path
	attackVal := float32(10) + sliders[0].FloatWidget.Value/20000000  //2000000
	decayVal := attackVal + (sliders[1].FloatWidget.Value / 20000000) // 2000000
	susamp := -20 - (sliders[2].FloatWidget.Value * 6000)
	releaseVal := attackVal + decayVal + (sliders[3].FloatWidget.Value / 20000000)

	path.Begin(ops)
	path.MoveTo(f32.Pt(0, -20))
	path.LineTo(f32.Pt(attackVal, -80))
	path.LineTo(f32.Pt(decayVal, susamp))
	path.LineTo(f32.Pt(attackVal+decayVal, susamp))
	path.LineTo(f32.Pt(releaseVal, -20))
	//color := color.NRGBA{R: 255, A: 255, B: 100}
	paint.FillShape(ops, NewMyStyles().colorAccent,
		clip.Stroke{
			Path:  path.End(),
			Width: 8,
		}.Op())

}
