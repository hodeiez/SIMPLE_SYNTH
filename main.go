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
	bufferSize := 128
	osc := generator.Oscillator(bufferSize)

	count := generator.Triangle
	controller := gui.Controls{SelectorFunc: &count}
	midiMessages := []midi.MidiMsg{{Key: 0, On: false}, {Key: 0, On: false}}
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

	//simple command menu thread
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

		midi.RunMidi(&midiMessages)

	}()
	//thread for audio
	go func() {

		start := dsp.DspConf{BufferSize: bufferSize, Osc: &osc}
		dsp.Run(start)
	}()

	//evaluate and execute changes

	for {
		//log.Println(midiMessages)
		generator.ChangeFreq(midiMessages, &osc)
		generator.SelectWave(*controller.SelectorFunc, &osc)

	}

}
