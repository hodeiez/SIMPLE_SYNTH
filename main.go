package main

import (
	//	"log"

	"hodei.naiz/simplesynth/gui"
	"hodei.naiz/simplesynth/synth/dsp"
	"hodei.naiz/simplesynth/synth/generator"
	"hodei.naiz/simplesynth/synth/midi"
)

func main() {

	//**********************************************setup**********************************
	bufferSize := 2048

	//--------------------controllers------------------
	count := generator.Triangle
	count2 := generator.Sine

	attackCtrl := 1000.00
	decayCtrl := 2000.00
	susCtrl := 0.08
	relCtrl := 2000.00

	amplitudeVal := 0.0
	pitch := 0.0

	adsrControl := generator.ADSRControl{AttackTime: &attackCtrl, DecayTime: &decayCtrl, SustainAmp: &susCtrl, ReleaseTime: &relCtrl}
	controller := generator.Controls{SelectorFunc2: &count2, SelectorFunc: &count, ShowAmp: &amplitudeVal, ADSRcontrol: &adsrControl, Pitch: &pitch}

	vmanager := generator.PolyInit(bufferSize, 3, controller) //4 is max polyphony

	//**********************************************gui****************************************************************
	msg := make(chan midi.MidiMsg)

	go gui.Run(&controller)

	//thread for midi

	go midi.RunMidi(msg)

	//thread for audio
	start := dsp.DspConf{BufferSize: bufferSize, VM: &vmanager}

	go dsp.Run(start)

	for {

		go generator.RunPolly(vmanager, <-msg, controller)
		go generator.SelectWave(*controller.SelectorFunc, vmanager.Voices)

	}

}
