package main

import (
	/* "bufio" */
	"fmt"

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
	fmt.Println("hello synth")
	//**********************************************setup**********************************
	bufferSize := 128
	osc := generator.Oscillator(bufferSize)

	//--------------------controllers------------------
	count := generator.Triangle
	attackCtrl := 1000.00
	amplitudeVal := 0.0
	controller := gui.Controls{SelectorFunc: &count, AttackTime: &attackCtrl, ShowAmp: &amplitudeVal}
	//--------------------------------------------
	//---------------------------midi notes------------------------------
	midiMessages := []midi.MidiMsg{{Key: -1, On: false}, {Key: 0, On: false}}
	//--------------------------------------------------------------------------------
	//------------------------------ADSR-----------------------------------------------
	adsr := generator.ADSR{AttackTime: *controller.AttackTime, DecayTime: 2000.0, SustainAmp: 0.08, ReleaseTime: 2000.00, ControlAmp: 0.0}
	pos := 0.0
	//----------------------------------------------------------------------------------
	//************************************************************************************************

	//**********************************************gui****************************************************************

	go func() {

		w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(600)))

		err := gui.Render(w, &controller)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	//simple command menu thread
	/* scanner := bufio.NewScanner(os.Stdin)
	go func() {
		fmt.Println("type q if you want to quit")
		for scanner.Scan() {
			switch scanner.Text() {
			case "q":
				os.Exit(0)
			}
		}

	}() */

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
		//TODO: send all adsr Controls
		/* running  */
		/* 	log.Println(osc.Osc.Amplitude) */
		adsr.ADSR(midiMessages, &osc, &pos, controller.AttackTime)
		generator.ChangeFreq(midiMessages, &osc)
		generator.SelectWave(*controller.SelectorFunc, &osc)

		/* go run(adsr, midiMessages, osc, pos, controller) */
		//TODO: avoid log println to run adsr.ADSR

	}

}
func run(adsr generator.ADSR, midiMessages []midi.MidiMsg, osc generator.Osc, pos float64, controller gui.Controls) {
	adsr.ADSR(midiMessages, &osc, &pos, controller.AttackTime)
	generator.ChangeFreq(midiMessages, &osc)
	generator.SelectWave(*controller.SelectorFunc, &osc)
	log.Println("a")
}
