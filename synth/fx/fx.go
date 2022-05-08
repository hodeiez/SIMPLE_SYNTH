package fx

import (
	"log"

	"hodei.naiz/simplesynth/synth/generator"
)

type FX struct {
	PitchChan chan float64
}

func Init() FX {

	return FX{PitchChan: make(chan float64)}

}
func (fx *FX) Run(vmanager generator.VoiceManager, controllerPitch float64) {
	pitch := 440.00
	for {
		select {
		case <-fx.PitchChan:
			log.Println(<-fx.PitchChan)

			for _, o := range vmanager.Voices {

				o.Oscillator.Osc.SetFreq(o.Oscillator.Osc.Freq + (<-fx.PitchChan - o.Oscillator.BaseFreq))

			}
			for _, o := range vmanager.Voices {

				o.Oscillator.BaseFreq = <-fx.PitchChan
			}
			pitch = <-fx.PitchChan
			log.Println(<-fx.PitchChan)
		default:
			log.Println("IN DEFAULT", pitch)

		}
	}

}
