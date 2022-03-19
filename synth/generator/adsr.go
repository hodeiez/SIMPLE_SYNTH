package generator

import (
	/* "log" */
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

func (voice *Voice) ADSRforPoly(adsrCtrl *ADSRControl) {

	a := adsrCtrl.AttackTime
	d := adsrCtrl.DecayTime
	s := adsrCtrl.SustainAmp
	r := adsrCtrl.ReleaseTime

	if voice.Midi.On {

		if *voice.TimeControl < *a && voice.Oscillator.Osc.Amplitude < 1 { //ATTACK

			voice.Oscillator.Osc.Amplitude += 1 / *a
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
func (adsr *ADSR) ADSR(midimsg midi.MidiMsg, osc *Osc, pos *float64, adsrCtrl *ADSRControl, currentNote *midi.MidiMsg) {

	a := adsrCtrl.AttackTime
	d := adsrCtrl.DecayTime
	s := adsrCtrl.SustainAmp
	r := adsrCtrl.ReleaseTime

	if midimsg.On {
		*currentNote = midimsg
	}
	if midimsg.On {

		if *pos < *a && osc.Osc.Amplitude < 0.6 { //ATTACK

			osc.Osc.Amplitude += 0.6 / *a
		} else if *pos > *a && *pos < *a+*d { //DECAY

			if osc.Osc.Amplitude > adsr.SustainAmp {
				osc.Osc.Amplitude -= (1 / *d)

			}
		} else if *pos >= *a+*d { //SUSTAIN

			osc.Osc.Amplitude = *s

		}
		adsr.ControlAmp = osc.Osc.Amplitude
		*pos++
	} else if !midimsg.On && currentNote.Key == midimsg.Key {

		*pos = 0.0

		if osc.Osc.Amplitude > 0.0 && adsr.ControlAmp != 0.0 {
			osc.Osc.Amplitude -= (adsr.ControlAmp / *r)
		} else {
			osc.Osc.Amplitude = 0

		}

	}

	time.Sleep(1 * time.Nanosecond)
}
