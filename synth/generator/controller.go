package generator

type ADSRControl struct {
	AttackTime  *float64
	DecayTime   *float64
	SustainAmp  *float64
	ReleaseTime *float64
}
type Controls struct {
	SelectorFunc *MyWaveType
	ADSRcontrol  *ADSRControl
	ShowAmp      *float64
}
