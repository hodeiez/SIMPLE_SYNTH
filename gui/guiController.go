package gui

import "hodei.naiz/simplesynth/synth/generator"

type Controls struct {
	SelectorFunc *generator.MyWaveType
	AttackTime   *float64
	DecayTime    *float64
	SustainAmp   *float64
	ReleaseTime  *float64
}
