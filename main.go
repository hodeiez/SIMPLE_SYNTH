package main

import (
	"hodei.naiz/simplesynth/gui"
	"hodei.naiz/simplesynth/synth/dsp"
	"hodei.naiz/simplesynth/synth/fx"
	"hodei.naiz/simplesynth/synth/generator"
	"hodei.naiz/simplesynth/synth/midi"
)

func main() {

	//**********************************************setup**********************************
	bufferSize := 2048

	//--------------------controllers------------------
	count := generator.Triangle
	count2 := generator.Sine

	attackCtrl := 50.00
	decayCtrl := 50.00
	susCtrl := 0.08
	relCtrl := 2000.00

	amplitudeVal := 0.0
	pitch := 440.0

	adsrControl := generator.ADSRControl{AttackTime: &attackCtrl, DecayTime: &decayCtrl, SustainAmp: &susCtrl, ReleaseTime: &relCtrl}
	controller := generator.Controls{SelectorFunc2: &count2, SelectorFunc: &count, ShowAmp: &amplitudeVal, ADSRcontrol: &adsrControl, Pitch: &pitch}

	vmanager := generator.PolyInit(bufferSize, 6, controller)

	//**********************************************gui****************************************************************
	msg := make(chan midi.MidiMsg)

	fx := fx.Init()
	go gui.Run(&controller, fx.PitchChan)

	go midi.RunMidi(msg)

	start := dsp.DspConf{BufferSize: bufferSize, VM: &vmanager}

	go dsp.Run(start)

	go fx.Run(vmanager, *controller.Pitch)
	for {
		go generator.RunPolly(vmanager, <-msg, controller)
		go generator.SelectWave(*controller.SelectorFunc, vmanager.Voices)

	}

}
