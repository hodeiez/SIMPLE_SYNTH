package generator

import (
	"time"

	"hodei.naiz/simplesynth/synth/midi"
)

type ADSR struct {
	AttackTime  float64
	DecayTime   float64
	SustainAmp  float64
	ReleaseTime float64
	ControlAmp  float64
}

func (adsr *ADSR) ADSR(midimsg []midi.MidiMsg, osc *Osc, pos *float64, adsrCtrl *ADSRControl) {

	a := adsrCtrl.AttackTime
	d := adsrCtrl.DecayTime
	s := adsrCtrl.SustainAmp
	r := adsrCtrl.ReleaseTime

	if midi.IsOn(midimsg) {

		if *pos < *a && osc.Osc.Amplitude < 1 { //ATTACK

			osc.Osc.Amplitude += 1 / *a
		} else if *pos > *a && *pos < *a+*d { //DECAY

			if osc.Osc.Amplitude > adsr.SustainAmp {
				osc.Osc.Amplitude -= (1 / *d)

			}
		} else if *pos >= *a+*d { //SUSTAIN

			osc.Osc.Amplitude = *s

		}
		adsr.ControlAmp = osc.Osc.Amplitude
		*pos++
	} else if !midi.IsOn(midimsg) {
		*pos = 0.0

		if osc.Osc.Amplitude > 0.0 && adsr.ControlAmp != 0.0 {
			osc.Osc.Amplitude -= (adsr.ControlAmp / *r)
		} else {
			osc.Osc.Amplitude = 0

		}

	}
	time.Sleep(1 * time.Nanosecond)
}
