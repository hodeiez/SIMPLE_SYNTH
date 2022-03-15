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
	osc := generator.Oscillator(bufferSize)

	//--------------------controllers------------------
	count := generator.Triangle
	attackCtrl := 1000.00
	decayCtrl := 2000.00
	susCtrl := 0.08
	relCtrl := 2000.00

	amplitudeVal := 0.0

	adsrControl := generator.ADSRControl{AttackTime: &attackCtrl, DecayTime: &decayCtrl, SustainAmp: &susCtrl, ReleaseTime: &relCtrl}
	controller := generator.Controls{SelectorFunc: &count, ShowAmp: &amplitudeVal, ADSRcontrol: &adsrControl}
	//--------------------------------------------
	//---------------------------midi notes------------------------------
	//midiMessages := []midi.MidiMsg{{Key: -1, On: false}, {Key: 0, On: false}}
	midiMessages := midi.MidiMsg{Key: -1, On: false}
	//--------------------------------------------------------------------------------
	//------------------------------ADSR-----------------------------------------------
	adsr := generator.ADSR{AttackTime: *controller.ADSRcontrol.AttackTime, DecayTime: *controller.ADSRcontrol.DecayTime, SustainAmp: *controller.ADSRcontrol.SustainAmp, ReleaseTime: *controller.ADSRcontrol.ReleaseTime, ControlAmp: 0.01}
	pos := 0.0
	pos2 := 0.0
	//----------------------------------------------------------------------------------
	//************************************************************************************************
	/***

	TESTING POLYPHONY*/
	osc2 := generator.Oscillator(bufferSize)
	voicesArray := []*generator.Voice{{Oscillator: &osc, Midi: midi.MidiMsg{Key: -1, On: false}}, {Oscillator: &osc2, Midi: midi.MidiMsg{Key: -1, On: false}}}
	voices := generator.VoiceManager{Voices: voicesArray}

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

		start := dsp.DspConf{BufferSize: bufferSize, Osc: &osc, Osc2: &osc2}
		dsp.Run(start)
	}()

	//evaluate and execute changes
	currentNote := midi.MidiMsg{-1, false}

	for {

		adsr.ADSR(midiMessages, &osc, &pos, controller.ADSRcontrol, &currentNote)
		adsr.ADSR(midiMessages, &osc2, &pos2, controller.ADSRcontrol, &currentNote)
		//generator.ChangeFreq(midiMessages, &osc)
		generator.SelectWave(*controller.SelectorFunc, &osc)
		generator.SelectWave(*controller.SelectorFunc, &osc2)
		generator.RunPolly(voices, midiMessages)
		//generator.VoiceOnNoteOff(voices, midiMessages)
	}

}
