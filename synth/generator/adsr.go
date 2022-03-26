package generator

import (
//"log"
// "log"
// "time"
)

type ADSR struct {
	AttackTime  float64
	DecayTime   float64
	SustainAmp  float64
	ReleaseTime float64
	ControlAmp  float64
}

//TODO: make dynamic to accept other parameters to apply (cutoff, pitch,...)
func (voice *Voice) RunADSR(controller Controls) {
	if voice.Midi.On {
	loop:
		for {

			select {
			case <-voice.Quit:

				break loop
			default:

				if *controller.ADSRcontrol.AttackTime > voice.TimeControl {

					voice.adsrAction("INCREASE_AMP", 1/(*controller.ADSRcontrol.AttackTime*1000))
				} else if *controller.ADSRcontrol.AttackTime+*controller.ADSRcontrol.DecayTime > voice.TimeControl && *controller.ADSRcontrol.SustainAmp < voice.Oscillator.Osc.Amplitude {
					voice.adsrAction("DECREASE_AMP", 1/(*controller.ADSRcontrol.DecayTime*1000))

				}

				voice.TimeControl += 0.1
				continue
			}

		}
	} else if !voice.Midi.On {
		voice.Quit <- true
		for {
			if voice.Oscillator.Osc.Amplitude <= 0 {
				voice.TimeControl = 0.0
				break
			}
			voice.adsrAction("DECREASE_AMP", 1/(*controller.ADSRcontrol.ReleaseTime*100))
			voice.TimeControl += 0.1
		}
	}

}
