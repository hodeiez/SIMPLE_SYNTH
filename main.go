package main

import (
	"bufio"
	"fmt"

	"gioui.org/app"
	"hodei.naiz/simplesynth/gui"
	"hodei.naiz/simplesynth/synth/dsp"
	"hodei.naiz/simplesynth/synth/generator"
	"hodei.naiz/simplesynth/synth/midi"
	"log"
	"os"
)

func main() {
	fmt.Println("hello synth")
	//setup**********************************
	bufferSize := 512
	osc := generator.Oscillator(bufferSize)
	message := midi.MidiMsg{Key: 0, On: false}
	count := generator.Triangle
	controller := gui.Controls{SelectorFunc: &count}
	//TODO: fix latency midi message->osc
	//************************************************************************************************
	//gui****************************************************************

	go func() {

		w := app.NewWindow()
		err := gui.Render(w, &controller)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	//simple menu thread this will be for gui
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		fmt.Println("type q if you want to quit")
		for scanner.Scan() {
			switch scanner.Text() {
			case "q":
				os.Exit(0)
			}
		}

	}()

	//thread for midi
	go func() {

		midi.RunMidi(&message)

	}()
	//thread for audio
	go func() {

		start := dsp.DspConf{BufferSize: bufferSize, Osc: osc}
		dsp.Run(start)
	}()

	//evaluate and execute changes

	for {

		generator.ChangeFreq(message, &osc)
		generator.SelectWave(*controller.SelectorFunc, &osc)

	}

}
