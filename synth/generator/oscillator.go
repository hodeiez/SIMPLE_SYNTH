package generator

import (
	"log"
	"math"

	"os"
	"os/signal"

	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
	"hodei.naiz/simplesynth/synth/midi"
)

type Osc struct {
	amplitude   *float64
	gainControl float64
	Osc         *generator.Osc
	Buf         *audio.FloatBuffer
}

func Oscillator(bufferSize int) Osc {
	// this has to go to a preconf**************

	buf := &audio.FloatBuffer{
		Data:   make([]float64, bufferSize),
		Format: audio.FormatMono44100,
	}
	//***************************
	currentNote := 440.0
	osc := generator.NewOsc(generator.WaveSine, currentNote, buf.Format.SampleRate)
	osc.Amplitude = 1
	osc.Freq = 440.0
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	log.Println("oscillator running")
	return Osc{&osc.Amplitude, 0.0, osc, buf}

}

func ChangeFreq(midimsg midi.MidiMsg, osc *Osc) Osc {
	//TODO: use array loop when polyphony is on
	//	var NoteToPitch = make([]float64, 128)
	a := 440.0
	//for i := 0; i < 128; i++ {
	NoteToPitch := (a / 32) * (math.Pow(2, ((float64(midimsg.Key) - 9) / 12)))
	//	}
	//TODO: fix logic for noteOn noteOff
	osc.Osc.SetFreq(NoteToPitch)
	osc.Osc.Shape = generator.WaveType(generator.WaveTriangle)
	if midimsg.On {

		osc.Osc.Amplitude = 100
	} else if !midimsg.On {
		osc.Osc.Amplitude = 0
	}
	return *osc
}

func SelectWave(selector int, osc *Osc) Osc {
	switch selector {
	case 0:
		osc.Osc.Shape = generator.WaveType(generator.WaveTriangle)
	case 1:
		osc.Osc.Shape = generator.WaveType(generator.WaveSaw)
	case 2:
		osc.Osc.Shape = generator.WaveType(generator.WaveSqr)
	case 3:
		osc.Osc.Shape = generator.WaveType(generator.WaveSine)

	}
	return *osc
}
