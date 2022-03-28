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
	gainControl float64
	Osc         *generator.Osc
	Buf         *audio.FloatBuffer
}

type MyWaveType int64

const (
	Triangle MyWaveType = iota
	Saw
	Square
	Sine
	MyWaveTypeSize
)

func (s MyWaveType) String() string {

	switch s {
	case Triangle:
		return "Triangle"
	case Saw:
		return "Saw"
	case Sine:
		return "Sine"
	case Square:
		return "Square"
	}

	return "-"
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
	return Osc{0.0, osc, buf}

}

func ChangeFreq(midimsg midi.MidiMsg, osc *Osc) Osc {

	a := 440.0

	NoteToPitch := (a / 32) * (math.Pow(2, ((float64(midimsg.Key) - 9) / 12)))

	if midimsg.On {
		osc.Osc.SetFreq(NoteToPitch)
	}
	return *osc
}

func SelectWave(selector MyWaveType, voices []*Voice) {
	for _, o := range voices {
		switch selector {
		case 0:
			o.Oscillator.Osc.Shape = generator.WaveType(generator.WaveTriangle)
			o.Oscillator.Osc.Shape = generator.WaveType(generator.WaveSine)
		case 1:
			o.Oscillator.Osc.Shape = generator.WaveType(generator.WaveSaw)
			o.Oscillator2.Osc.Shape = generator.WaveType(generator.WaveTriangle)
		case 2:
			o.Oscillator2.Osc.Shape = generator.WaveType(generator.WaveSqr)
			o.Oscillator2.Osc.Shape = generator.WaveType(generator.WaveSaw)
		case 3:
			o.Oscillator2.Osc.Shape = generator.WaveType(generator.WaveSine)
			o.Oscillator.Osc.Shape = generator.WaveType(generator.WaveSqr)

		}

	}
}
