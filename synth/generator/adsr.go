package generator

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
	loop:
		for {

			select {
			case <-voice.Quit:
				break loop
			default:

				if aTime > voice.TimeControl && voice.TimeControl < aTime+dTime {
					if aTime == 1 {

						*parameter = 0.01 //max value
					} else if aTime != 1 {
						go voice.adsrAction("INCREASE", actionType, 1/(aTime)) //1000
					}
				} else if aTime+dTime > voice.TimeControl && sAmp < *controlRate {
					go voice.adsrAction("DECREASE", actionType, 1/(dTime)) //1000

				}

				voice.TimeControl += 0.1 //0.1
				continue
			}

		}

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
