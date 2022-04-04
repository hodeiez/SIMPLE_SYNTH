package generator

type ADSRControl struct {
	AttackTime  *float64
	DecayTime   *float64
	SustainAmp  *float64
	ReleaseTime *float64
}
type Controls struct {
	SelectorFunc  *MyWaveType
	SelectorFunc2 *MyWaveType //think on how to apply for second OSC
	ADSRcontrol   *ADSRControl
	ShowAmp       *float64
	Pitch         *float64
}
