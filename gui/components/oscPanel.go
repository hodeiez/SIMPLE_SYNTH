package components

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"hodei.naiz/simplesynth/helpers"
	"hodei.naiz/simplesynth/synth/generator"
)

type OscPanel struct {
	Margin       *layout.Inset
	WaveSelector *Selector
	WaveType     *generator.MyWaveType
	PitchSlider  MySlider
}

func NewOscPanel(selector Selector, waveType *generator.MyWaveType, slider MySlider) OscPanel {
	marginOscPanel := layout.Inset{Top: unit.Dp(0),
		Bottom: unit.Dp(0),
		Right:  unit.Dp(0),
		Left:   unit.Dp(0)}
	return OscPanel{Margin: &marginOscPanel, WaveSelector: &selector, WaveType: waveType, PitchSlider: slider}
}
func (oscPanel OscPanel) Render(th *material.Theme) layout.FlexChild {
	//selectorSize := layout.Dimensions{Size: image.Point{1, 2}}

	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return oscPanel.Margin.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceSides, WeightSum: 10}.Layout(gtx,

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					selector := oscPanel.WaveSelector
					return ShowSelector(th, gtx, selector, oscPanel.WaveType)

				}),
				layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
					label := material.Label(th, unit.Dp(10), oscPanel.PitchSlider.StyledLabel.Text)
					label.TextSize = unit.Dp(20)
					slider := oscPanel.PitchSlider
					slider.FloatWidget.Axis = layout.Horizontal

					return slider.StyledSlide.Layout(gtx)
				}),

				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {

					label := material.Label(th, unit.Dp(10), helpers.Float32ToString(oscPanel.PitchSlider.FloatWidget.Value))
					label.TextSize = unit.Dp(13)
					return label.Layout(gtx)
				}),
			)
		})
	})
}
