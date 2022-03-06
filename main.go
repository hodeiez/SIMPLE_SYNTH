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
	//IN GOAUDIO osc.go file comment //fmt.Println(o.CurrentPhaseAngle) in line 117
	//**********************************************setup**********************************
	bufferSize := 128
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
	midiMessages := []midi.MidiMsg{{Key: -1, On: false}, {Key: 0, On: false}}
	//--------------------------------------------------------------------------------
	//------------------------------ADSR-----------------------------------------------
	adsr := generator.ADSR{AttackTime: *controller.ADSRcontrol.AttackTime, DecayTime: *controller.ADSRcontrol.DecayTime, SustainAmp: *controller.ADSRcontrol.SustainAmp, ReleaseTime: *controller.ADSRcontrol.ReleaseTime, ControlAmp: 0.01}
	pos := 0.0
	//----------------------------------------------------------------------------------
	//************************************************************************************************

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

		start := dsp.DspConf{BufferSize: bufferSize, Osc: &osc}
		dsp.Run(start)
	}()

	//evaluate and execute changes

	for {

		adsr.ADSR(midiMessages, &osc, &pos, controller.ADSRcontrol)
		generator.ChangeFreq(midiMessages, &osc)
		generator.SelectWave(*controller.SelectorFunc, &osc)

	}

}
