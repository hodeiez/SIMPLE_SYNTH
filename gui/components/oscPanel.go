package components

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
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
		Bottom: unit.Dp(20),
		Right:  unit.Dp(0),
		Left:   unit.Dp(0)}
	return OscPanel{Margin: &marginOscPanel, WaveSelector: &selector, WaveType: waveType, PitchSlider: slider}
}
func (oscPanel OscPanel) Render(th *material.Theme) layout.FlexChild {
	return layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
		return oscPanel.Margin.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Spacing: layout.Spacing(255), WeightSum: 20}.Layout(gtx,

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					theSelector := ShowSelector(th, gtx, oscPanel.WaveSelector, oscPanel.WaveType)
					theSelector.Size.X = 200
					return theSelector

				}),
				layout.Flexed(5, func(gtx layout.Context) layout.Dimensions {
					label := material.Label(th, unit.Dp(10), oscPanel.PitchSlider.StyledLabel.Text)
					label.TextSize = unit.Dp(20)

					return oscPanel.PitchSlider.StyledSlide.Layout(gtx)
					//return label.Layout(gtx)
				}),
				layout.Flexed(5, func(gtx layout.Context) layout.Dimensions {
					label := material.Label(th, unit.Dp(10), "VOLUME")
					label.TextSize = unit.Dp(20)
					return label.Layout(gtx)
				}),
			)
		})
	})
}
