package gui

import "hodei.naiz/simplesynth/synth/generator"

type ADSRControl struct {
	AttackTime  *float64
	DecayTime   *float64
	SustainAmp  *float64
	ReleaseTime *float64
}
type Controls struct {
	SelectorFunc *generator.MyWaveType
	AttackTime   *float64
	DecayTime    *float64
	SustainAmp   *float64
	ReleaseTime  *float64
	ADSRcontrol  *ADSRControl
	ShowAmp      *float64
}
