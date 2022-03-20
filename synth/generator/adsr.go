package generator

import (
	/* "log" */
	"time"
)

type ADSR struct {
	AttackTime  float64
	DecayTime   float64
	SustainAmp  float64
	ReleaseTime float64
	ControlAmp  float64
}

//TODO: fix values and refactor to one method
func (voice *Voice) ADSRforPoly(adsrCtrl *ADSRControl) {

	a := adsrCtrl.AttackTime
	d := adsrCtrl.DecayTime
	s := adsrCtrl.SustainAmp
	r := adsrCtrl.ReleaseTime

	if voice.Midi.On {

		if *voice.TimeControl < *a && voice.Oscillator.Osc.Amplitude < 0.1 { //ATTACK

			voice.Oscillator.Osc.Amplitude += 0.1 / *a
		} else if *voice.TimeControl > *a && *voice.TimeControl < *a+*d { //DECAY

			if voice.Oscillator.Osc.Amplitude > voice.ADSR.SustainAmp {
				voice.Oscillator.Osc.Amplitude -= (1 / *d)

			}
		} else if *voice.TimeControl >= *a+*d { //SUSTAIN

			voice.Oscillator.Osc.Amplitude = *s

		}
		voice.ADSR.ControlAmp = voice.Oscillator.Osc.Amplitude
		*voice.TimeControl++
		//this goes to noteOff in voice
	} else if !voice.Midi.On {

		*voice.TimeControl = 0.0

		if voice.Oscillator.Osc.Amplitude > 0.0 && voice.ADSR.ControlAmp != 0.0 {
			voice.Oscillator.Osc.Amplitude -= (voice.ADSR.ControlAmp / *r)
		} else {
			voice.Oscillator.Osc.Amplitude = 0

		}

	}

	time.Sleep(1 * time.Nanosecond)
}
