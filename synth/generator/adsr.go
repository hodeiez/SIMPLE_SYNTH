package generator

import "log"

//import "log"

type ADSR struct {
	AttackTime  float64
	DecayTime   float64
	SustainAmp  float64
	ReleaseTime float64
	ControlAmp  float64
}

//TODO: fix this chaos, do it reusable and fix performance
func (voice *Voice) RunADSR(parameter *float64, controller Controls, controlRate *float64, actionType string) {
	aTime := *controller.ADSRcontrol.AttackTime
	dTime := *controller.ADSRcontrol.DecayTime
	sAmp := *controller.ADSRcontrol.SustainAmp
	rTime := *controller.ADSRcontrol.ReleaseTime
	if voice.Midi.On {
		*parameter = 0.0 //min value
		// voice.Oscillator.Osc.Amplitude = 0.0
	loop:
		for {

			select {
			case <-voice.Quit:
				break loop
			default:

				if aTime > voice.TimeControl && voice.TimeControl < aTime+dTime {
					if aTime == 1 {
						// voice.Oscillator.Osc.Amplitude = 0.01
						*parameter = 0.01 //max value
					} else if aTime != 1 {
						voice.adsrAction("INCREASE", actionType, 1/(aTime)) //1000
					}
				} else if aTime+dTime > voice.TimeControl && sAmp < *controlRate {
					voice.adsrAction("DECREASE", actionType, 1/(dTime)) //1000

				}

				voice.TimeControl += 0.1 //0.1
				continue
			}

		}
		//TODO: fix sustain loop
	} else if !voice.Midi.On {
		voice.Quit <- true
		zero := 0.0
		// voice.TimeControl = zero
		// *controlRate = zero
	loop2:
		for {
			if *controlRate <= zero {
				*parameter = zero
				voice.TimeControl = 0.0
				//voice.Oscillator.Osc.Reset()

				break loop2
			}
			if voice.TimeControl < rTime {
				voice.adsrAction("DECREASE", actionType, 1/(rTime*100))
			}
			//voice.adsrAction("DECREASE", actionType, 1/(rTime*100)) //100
			voice.TimeControl += 0.1 //0.1

		}
	}
	log.Printf("%f timeConrtol,", voice.TimeControl)
}
