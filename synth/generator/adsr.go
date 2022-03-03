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

	if midi.IsOn(midimsg) {

		if *pos < a && osc.Osc.Amplitude < 1 { //ATTACK
			log.Println("Attack")

			osc.Osc.Amplitude += 1 / a
		} else if *pos > a && *pos < a+d { //DECAY
			log.Println("Decay")
			if osc.Osc.Amplitude >= adsr.SustainAmp {
				osc.Osc.Amplitude -= (1 / d)

			}
		} else if *pos >= a+d { //SUSTAIN
			log.Println("Sustain")
			osc.Osc.Amplitude = adsr.SustainAmp

		}
		adsr.ControlAmp = osc.Osc.Amplitude
		*pos++
	} else if !midi.IsOn(midimsg) {
		*pos = 0.0

		log.Println("RELEASE")
		if osc.Osc.Amplitude > 0.0 && adsr.ControlAmp != 0.0 {
			osc.Osc.Amplitude -= (adsr.ControlAmp / r)
		} else {
			osc.Osc.Amplitude = 0

		}

	}

}
