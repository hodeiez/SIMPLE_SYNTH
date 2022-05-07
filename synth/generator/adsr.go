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
type adsrRunner func(*Voice, *float64, Controls, *float64, string)

//TODO: fix a struct for params
func midiON(voice *Voice, parameter *float64, controller Controls, controlRate *float64, actionType string) {

	if *controller.ADSRcontrol.AttackTime >= voice.TimeControl {
		if *controller.ADSRcontrol.AttackTime == 0 {

			*parameter = 0.01 //max value
		} else if *controller.ADSRcontrol.AttackTime > 0.0 && *parameter < 0.01 {
			go voice.adsrAction("INCREASE", actionType, 0.01/(*controller.ADSRcontrol.AttackTime))

		}
	}
	if *controller.ADSRcontrol.AttackTime < voice.TimeControl && *controller.ADSRcontrol.SustainAmp <= *controlRate {
		go voice.adsrAction("DECREASE", actionType, 0.01/(*controller.ADSRcontrol.DecayTime))

	}

	voice.TimeControl += 1
}
func ticker(runner adsrRunner, voice *Voice, parameter *float64, controller Controls, controlRate *float64, actionType string) {
	ticker := time.NewTicker(1 * time.Millisecond)

	go func() {
		for {
			select {
			case <-ticker.C:
				runner(voice, parameter, controller, controlRate, actionType)
			case <-voice.Quit:
				ticker.Stop()
				return
			}
		}
	}()

}
func (voice *Voice) RunADSR(parameter *float64, controller Controls, controlRate *float64, actionType string) {

	rTime := *controller.ADSRcontrol.ReleaseTime
	if voice.Midi.On {
		*parameter = 0.0 //min value
		ticker(midiON, voice, parameter, controller, controlRate, actionType)

	} else if !voice.Midi.On {

		voice.Quit <- true
		zero := 0.0 //min value

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
