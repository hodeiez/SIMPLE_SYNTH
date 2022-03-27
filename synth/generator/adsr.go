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

//TODO: review logic and values
func (voice *Voice) RunADSR(controller Controls, controlRate *float64, actionType string) {
	if voice.Midi.On {
	loop:
		for {

			select {
			case <-voice.Quit:

				break loop
			default:

				if *controller.ADSRcontrol.AttackTime > voice.TimeControl {

					voice.adsrAction("INCREASE", actionType, 1/(*controller.ADSRcontrol.AttackTime*1000))
				} else if *controller.ADSRcontrol.AttackTime+*controller.ADSRcontrol.DecayTime > voice.TimeControl && *controller.ADSRcontrol.SustainAmp < *controlRate {
					voice.adsrAction("DECREASE", actionType, 1/(*controller.ADSRcontrol.DecayTime*1000))

				}

				voice.TimeControl += 0.1
				continue
			}

		}
	} else if !voice.Midi.On {
		voice.Quit <- true
		zero := 0.0
		for {
			if *controlRate <= zero {
				voice.TimeControl = 0.0
				break
			}
			voice.adsrAction("DECREASE", actionType, 1/(*controller.ADSRcontrol.ReleaseTime*100))
			voice.TimeControl += 0.1
		}
	}

}
