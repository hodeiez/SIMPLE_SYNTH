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

/* func (voice *Voice) ADSRforPoly(adsrCtrl *ADSRControl) {
//log.Println(voice)
a := adsrCtrl.AttackTime
d := adsrCtrl.DecayTime
s := adsrCtrl.SustainAmp
r := adsrCtrl.ReleaseTime
//log.Println(voice)
if voice.Midi.On {

	if *voice.TimeControl < *a && voice.Oscillator.Osc.Amplitude < 0.1 { //ATTACK

		voice.Oscillator.Osc.Amplitude += 0.1 / *voice.TimeControl
	} else if *voice.TimeControl > *a && *voice.TimeControl < *a+*d { //DECAY

		if voice.Oscillator.Osc.Amplitude > voice.ADSR.SustainAmp {
			voice.Oscillator.Osc.Amplitude -= (0.1 / *d)

		}
	} else if *voice.TimeControl >= *a+*d { //SUSTAIN

		voice.Oscillator.Osc.Amplitude = *s

	}
	voice.ADSR.ControlAmp = voice.Oscillator.Osc.Amplitude
	go func() {
		for {
			*voice.TimeControl++
			//	time.Sleep(1 * time.Nanosecond)
		}

	}() //*voice.TimeControl++
	//this goes to noteOff in voice
} else if !voice.Midi.On {

	*voice.TimeControl = 0.0

	if voice.Oscillator.Osc.Amplitude > 0.0 && voice.ADSR.ControlAmp != 0.0 {
		voice.Oscillator.Osc.Amplitude -= (voice.ADSR.ControlAmp / *r)
	} else {
		voice.Oscillator.Osc.Amplitude = 0

	}

} */

//time.Sleep(1 * time.Nanosecond)
//}

//TODO: fix values and refactor to one method
// func (voice *Voice) ADSRforPoly(adsrCtrl *ADSRControl) {
// 	//log.Println(voice)
// 	a := adsrCtrl.AttackTime
// 	d := adsrCtrl.DecayTime
// 	s := adsrCtrl.SustainAmp
// 	r := adsrCtrl.ReleaseTime
// 	//log.Println(voice)
// 	if voice.Midi.On {

// 		if *voice.TimeControl < *a && voice.Oscillator.Osc.Amplitude < 0.1 { //ATTACK
// 			go func() {
// 				for {
// 					*voice.TimeControl++
// 					time.Sleep(1 * time.Nanosecond)
// 					voice.Oscillator.Osc.Amplitude += (0.1 / *a)
// 					voice.ADSR.ControlAmp = voice.Oscillator.Osc.Amplitude
// 					log.Println(*voice.TimeControl)
// 				}

// 			}()
// 			//voice.Oscillator.Osc.Amplitude += (0.1 / *a)
// 		} else if *voice.TimeControl > *a && *voice.TimeControl < *a+*d { //DECAY

// 			if voice.Oscillator.Osc.Amplitude > voice.ADSR.SustainAmp {
// 				go func() {
// 					for {
// 						*voice.TimeControl++
// 						time.Sleep(1 * time.Nanosecond)
// 						voice.Oscillator.Osc.Amplitude -= (0.1 / *d)
// 						voice.ADSR.ControlAmp = voice.Oscillator.Osc.Amplitude
// 					}

// 				}()

// 				//voice.Oscillator.Osc.Amplitude -= (0.1 / *d)

// 			}
// 		} else if *voice.TimeControl >= *a+*d { //SUSTAIN
// 			go func() {
// 				for {
// 					*voice.TimeControl++
// 					time.Sleep(1 * time.Nanosecond)
// 					voice.Oscillator.Osc.Amplitude = *s
// 					voice.ADSR.ControlAmp = voice.Oscillator.Osc.Amplitude
// 				}

// 			}()
// 			//voice.Oscillator.Osc.Amplitude = *s

// 		}
// 		voice.ADSR.ControlAmp = voice.Oscillator.Osc.Amplitude

// 		//this goes to noteOff in voice
// 	} else if !voice.Midi.On {

// 		*voice.TimeControl = 0.0

// 		if voice.Oscillator.Osc.Amplitude > 0.0 && voice.ADSR.ControlAmp != 0.0 {
// 			go func() {
// 				for {
// 					*voice.TimeControl++
// 					//time.Sleep(1 * time.Nanosecond)
// 					voice.Oscillator.Osc.Amplitude -= (voice.ADSR.ControlAmp / *r)

// 				}

// 			}()
// 			//voice.Oscillator.Osc.Amplitude -= (voice.ADSR.ControlAmp / *r)
// 		} else {
// 			voice.Oscillator.Osc.Amplitude = 0

// 		}

// 	}
// 	time.Sleep(1 * time.Nanosecond)
// 	//	log.Println(*voice.TimeControl)
// 	//*voice.TimeControl++
// 	//	time.Sleep(1 * time.Nanosecond)
// }
// func (voice *Voice) ATTACK(adsrCtrl *ADSRControl) {
// 	log.Println(voice)
// 	a := adsrCtrl.AttackTime

// 	r := adsrCtrl.ReleaseTime
// 	//log.Println(voice)

// 	if voice.Midi.On {

// 		if *voice.TimeControl < *a && voice.Oscillator.Osc.Amplitude < 0.1 { //ATTACK

// 			voice.Oscillator.Osc.Amplitude += 0.1 / *a

// 			voice.ADSR.ControlAmp = voice.Oscillator.Osc.Amplitude
// 		} else if *voice.TimeControl > *a {

// 			//this goes to noteOff in voice
// 		} else if !voice.Midi.On {

// 			*voice.TimeControl = 0.0

// 			if voice.Oscillator.Osc.Amplitude > 0.0 && voice.ADSR.ControlAmp != 0.0 {
// 				voice.Oscillator.Osc.Amplitude -= (voice.ADSR.ControlAmp / *r)
// 			} else {
// 				voice.Oscillator.Osc.Amplitude = 0

// 			}

// 		}
// 	}

// 	time.Sleep(1 * time.Nanosecond)
// }

// func (voice *Voice) ADSRon(adsrCtrl *ADSRControl) {

// 	a := adsrCtrl.AttackTime
// 	//log.Println(voice)

// 	if voice.Midi.On {
// 		log.Println(*voice.TimeControl)
// 		if *voice.TimeControl < *a && voice.Oscillator.Osc.Amplitude < 1 { //ATTACK

// 			voice.Oscillator.Osc.Amplitude += 1 / *a
// 			//		voice.ADSR.ControlAmp = voice.Oscillator.Osc.Amplitude
// 		} else if *voice.TimeControl > *a {

// 			//this goes to noteOff in voice
// 		}
// 		//*voice.TimeControl++
// 		//	}
// 		go func() {
// 			for {
// 				*voice.TimeControl++
// 				time.Sleep(1 * time.Nanosecond)
// 			}

// 		}()

// 	}
// }
// func (voice *Voice) ADSRoff(adsrCtrl *ADSRControl) {

// 	voice.Oscillator.Osc.Amplitude = 0
// 	*voice.TimeControl = 0.0
// }
