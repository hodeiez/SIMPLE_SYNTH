package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
	"hodei.naiz/simplesynth/gui"
	"hodei.naiz/simplesynth/synth/dsp"
	"hodei.naiz/simplesynth/synth/generator"
	"hodei.naiz/simplesynth/synth/midi"
)

func main() {

	//**********************************************setup**********************************
	bufferSize := 64

	//--------------------controllers------------------
	count := generator.Triangle
	attackCtrl := 1000.00
	decayCtrl := 2000.00
	susCtrl := 0.08
	relCtrl := 2000.00

	amplitudeVal := 0.0

	adsrControl := generator.ADSRControl{AttackTime: &attackCtrl, DecayTime: &decayCtrl, SustainAmp: &susCtrl, ReleaseTime: &relCtrl}
	controller := generator.Controls{SelectorFunc: &count, ShowAmp: &amplitudeVal, ADSRcontrol: &adsrControl}

	midiMessages := midi.MidiMsg{Key: -1, On: false}

	vmanager := generator.PolyInit(bufferSize, 1, controller)

	//**********************************************gui****************************************************************

	go func() {

		w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(600)), app.Title("Symple synth"))

		err := gui.Render(w, &controller)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	//thread for midi
	go func() {

		midi.RunMidi(&midiMessages)

	}()
	//thread for audio
	go func() {

		start := dsp.DspConf{BufferSize: bufferSize, VM: &vmanager}
		dsp.Run(start)
	}()

	//evaluate and execute changes

	for {

		generator.SelectWave(*controller.SelectorFunc, vmanager.Voices)
		generator.RunPolly(vmanager, midiMessages, controller)

	}

}
