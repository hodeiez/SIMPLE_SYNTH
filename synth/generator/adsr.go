package generator

//"log"
//"log"
// "time"

type ADSR struct {
	AttackTime  float64
	DecayTime   float64
	SustainAmp  float64
	ReleaseTime float64
	ControlAmp  float64
}

//TODO: review logic and values, test adsrOnOn and adsrOnOff
func (voice *Voice) RunADSR(controller Controls, controlRate *float64, actionType string) {
	aTime := *controller.ADSRcontrol.AttackTime
	dTime := *controller.ADSRcontrol.DecayTime
	sAmp := *controller.ADSRcontrol.SustainAmp
	rTime := *controller.ADSRcontrol.ReleaseTime
	if voice.Midi.On {
		voice.Oscillator.Osc.Amplitude = 0.0
	loop:
		for {

			select {
			case <-voice.Quit:
				break loop
			default:

				if aTime > voice.TimeControl && voice.TimeControl < aTime+dTime {
					if aTime == 1 {
						voice.Oscillator.Osc.Amplitude = 0.01
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
	} else if !voice.Midi.On {
		voice.Quit <- true
		zero := 0.0
		for {
			if *controlRate <= zero {
				voice.TimeControl = 0.0
				voice.Oscillator.Osc.Reset()

				break
			}
			if voice.TimeControl < rTime {
				voice.adsrAction("DECREASE", actionType, 1/(rTime*100))
			}
			//voice.adsrAction("DECREASE", actionType, 1/(rTime*100)) //100
			voice.TimeControl += 0.1 //0.1
		}
	}

}
func (voice *Voice) RunADSROn(controller Controls, controlRate *float64, actionType string) {
	aTime := *controller.ADSRcontrol.AttackTime
	dTime := *controller.ADSRcontrol.DecayTime
	sAmp := *controller.ADSRcontrol.SustainAmp
	//rTime := *controller.ADSRcontrol.ReleaseTime
	//if voice.Midi.On {
	voice.Oscillator.Osc.Amplitude = 0.0
loop:
	for {

		select {
		case <-voice.Quit:
			break loop
		default:

			if aTime > voice.TimeControl && voice.TimeControl < aTime+dTime {
				if aTime == 1 {
					voice.Oscillator.Osc.Amplitude = 0.01
				} else if aTime != 1 {
					voice.adsrAction("INCREASE", actionType, 1/(aTime*1000)) //1000
				}
			} else if aTime+dTime > voice.TimeControl && sAmp < *controlRate {
				voice.adsrAction("DECREASE", actionType, 1/(dTime*1000)) //1000

			}

			voice.TimeControl += 0.1 //0.0001//0.1
			continue
		}

	}
	// } else if !voice.Midi.On {
	// 	voice.Quit <- true
	// 	zero := 0.0
	// 	for {
	// 		if *controlRate <= zero {
	// 			voice.TimeControl = 0.0
	// 			voice.Oscillator.Osc.Reset()

	// 			break
	// 		}
	// 		if voice.TimeControl < rTime {
	// 			voice.adsrAction("DECREASE", actionType, 1/(rTime*100))
	// 		}
	// 		//voice.adsrAction("DECREASE", actionType, 1/(rTime*100)) //100
	// 		voice.TimeControl += 0.1 //0.1
	// 	}
	// }

}
func (voice *Voice) RunADSROff(controller Controls, controlRate *float64, actionType string) {
	// aTime := *controller.ADSRcontrol.AttackTime
	// dTime := *controller.ADSRcontrol.DecayTime
	// sAmp := *controller.ADSRcontrol.SustainAmp
	rTime := *controller.ADSRcontrol.ReleaseTime
	// if voice.Midi.On {
	// 	voice.Oscillator.Osc.Amplitude = 0.0
	// loop:
	// 	for {

	// 		select {
	// 		case <-voice.Quit:
	// 			break loop
	// 		default:

	// 			if aTime > voice.TimeControl && voice.TimeControl < aTime+dTime {
	// 				if aTime == 1 {
	// 					voice.Oscillator.Osc.Amplitude = 0.01
	// 				} else if aTime != 1 {
	// 					voice.adsrAction("INCREASE", actionType, 1/(aTime*1000)) //1000
	// 				}
	// 			} else if aTime+dTime > voice.TimeControl && sAmp < *controlRate {
	// 				voice.adsrAction("DECREASE", actionType, 1/(dTime*1000)) //1000

	// 			}

	// 			voice.TimeControl += 0.1 //0.0001//0.1
	// 			continue
	// 		}

	// 	}
	// } else if !voice.Midi.On {
	voice.Quit <- true
	zero := 0.0
	for {
		if *controlRate <= zero {
			voice.TimeControl = 0.0
			voice.Oscillator.Osc.Reset()

			break
		}
		if voice.TimeControl < rTime {
			voice.adsrAction("DECREASE", actionType, 1/(rTime*100))
		}
		//voice.adsrAction("DECREASE", actionType, 1/(rTime*100)) //100
		voice.TimeControl += 0.1 //0.1
	}
	//	}

}
