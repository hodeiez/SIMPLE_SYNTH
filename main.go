package main

import (
	"fmt"
	/* "log" */

	/* "github.com/go-audio/transforms" */
	"hodei.naiz/simplesynth/synth/dsp"
	"hodei.naiz/simplesynth/synth/generator"
	"hodei.naiz/simplesynth/synth/midi"
)

func main() {
	fmt.Println("hello synth")
	bufferSize := 512
	osc := generator.Oscillator(bufferSize)
	message := midi.MidiMsg{Key: 0, On: false}
	go func() {

		midi.RunMidi(&message)

	}()
	go func() {

		start := dsp.DspConf{BufferSize: bufferSize, Osc: osc}
		dsp.Run(start)
	}()
	for {
		generator.ChangeFreq(float64(message.Key), &osc)

	}

}
