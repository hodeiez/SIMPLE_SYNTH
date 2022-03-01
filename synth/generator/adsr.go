package generator

import (
	"log"

	"hodei.naiz/simplesynth/synth/midi"
)

type ADSR struct {
	AttackTime  float64
	DecayTime   float64
	SustainAmp  float64
	ReleaseTime float64
	ControlAmp  float64
}

func (adsr *ADSR) ADSR(midimsg []midi.MidiMsg, osc *Osc, pos *float64) {

	a := adsr.AttackTime

	d := adsr.DecayTime

	r := adsr.ReleaseTime
	//TODO: check logic to restart ADSR when two notes On, check the values conversion
	if midi.IsOn(midimsg) {

		if *pos < a {
			log.Println("ATTACK", osc.Osc.Amplitude, adsr.ControlAmp)
			val := 1 / a
			adsr.ControlAmp = osc.Osc.Amplitude
			osc.Osc.Amplitude += val
		} else if *pos > a && *pos < a+d {
			log.Println("DECAY", osc.Osc.Amplitude)
			if osc.Osc.Amplitude >= adsr.SustainAmp {
				osc.Osc.Amplitude -= (1 / d)
				adsr.ControlAmp = osc.Osc.Amplitude
			} else {
				*pos = a + d
			}
		} else if *pos >= a+d {
			log.Println("SUSTAIN", osc.Osc.Amplitude)
			osc.Osc.Amplitude = adsr.SustainAmp
			adsr.ControlAmp = osc.Osc.Amplitude
		}
		*pos++
	} else if !midi.IsOn(midimsg) {
		*pos = 0.0
		log.Println("RELEASE", osc.Osc.Amplitude)
		if osc.Osc.Amplitude > 0.0 {
			osc.Osc.Amplitude -= (1 / r)
		} else {
			osc.Osc.Amplitude = 0

		}

	}

}
