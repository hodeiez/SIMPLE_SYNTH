package generator

import (
	//"log"
	//"log"
	"time"
)

type ADSR struct {
	AttackTime  float64
	DecayTime   float64
	SustainAmp  float64
	ReleaseTime float64
	ControlAmp  float64
}

func (voice *Voice) RunADSR(parameter *float64, controller Controls, controlRate *float64, actionType string) {
	aTime := *controller.ADSRcontrol.AttackTime
	dTime := *controller.ADSRcontrol.DecayTime
	sAmp := *controller.ADSRcontrol.SustainAmp
	rTime := *controller.ADSRcontrol.ReleaseTime
	if voice.Midi.On {
		*parameter = 0.0 //min value
		ticker := time.NewTicker(1 * time.Millisecond)
		//TODO: abstract timer to a function
		go func() {
			for {
				select {
				case <-ticker.C:
					//	log.Println(aTime)
					if aTime >= voice.TimeControl { // && voice.TimeControl < aTime {
						if aTime == 0 {

							*parameter = 0.01 //max value
						} else if aTime > 0.0 && *parameter < 0.01 {
							go voice.adsrAction("INCREASE", actionType, 0.01/(aTime))

						}
					}
					if aTime < voice.TimeControl && sAmp <= *controlRate {
						go voice.adsrAction("DECREASE", actionType, 0.01/(dTime))

					}

					voice.TimeControl += 1 //0.1
				case <-voice.Quit:
					ticker.Stop()
					return
				}
			}
		}()

	} else if !voice.Midi.On {

		voice.Quit <- true
		zero := 0.0

	loop2:
		for {
			if *controlRate <= zero {
				*parameter = zero
				voice.TimeControl = zero

				break loop2
			} else if voice.TimeControl < rTime && voice.TimeControl > zero {
				go voice.adsrAction("DECREASE", actionType, 1/(rTime))
			}
			voice.TimeControl += 0.1 //0.1

		}
	}
}
