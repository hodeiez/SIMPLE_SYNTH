package generator

import (
	"log"

	"hodei.naiz/simplesynth/synth/midi"
)

type ADSR struct {
	AttackTime  float64
	DecayTime   float64
	SustainTime float64
	ReleaseTime float64
}

func (adsr *ADSR) ADSR(midimsg []midi.MidiMsg, osc *Osc, pos *float64) {

	a := adsr.AttackTime

	d := adsr.DecayTime

	if *pos < a && midimsg[len(midimsg)-1].On {
		log.Println("attack")
		osc.Osc.Amplitude = *pos / a
		*pos++
	} else if !midimsg[len(midimsg)-1].On && midimsg[len(midimsg)-1].Key == midimsg[len(midimsg)-2].Key {
		log.Println("decay")
		if osc.Osc.Amplitude > 0.0 {
			osc.Osc.Amplitude = 1 - (*pos / d)
			*pos++
		} else if osc.Osc.Amplitude <= 0.0 {
			*pos = 0.0
			osc.Osc.Amplitude = 0.0
		}

		log.Println(*pos)
	}
}
