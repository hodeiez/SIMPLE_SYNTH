package dsp

import (
	"log"

	/* "github.com/go-audio/transforms" */
	"github.com/gordonklaus/portaudio"
	"hodei.naiz/simplesynth/synth/generator"
)

//** for now we run just one oscillator
type DspConf struct {
	BufferSize int
	Osc        *generator.Osc
}

func Run(dspConf DspConf) {

	portaudio.Initialize()

	defer portaudio.Terminate()
	out := make([]float32, dspConf.BufferSize)
	stream, err := portaudio.OpenDefaultStream(0, 1, 44100, len(out), &out)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("dsp running")
	defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}
	defer stream.Stop()

	for {
		// populate the out buffer
		if err := dspConf.Osc.Osc.Fill(dspConf.Osc.Buf); err != nil {
			log.Printf("error filling up the buffer")
		}
		//	transforms.Gain(dspConf.Osc.Buf, 20)
		//NoteOn(*noteOn, dspConf)
		f64ToF32Copy(out, dspConf.Osc.Buf.Data)

		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}

	}

}

func f64ToF32Copy(dst []float32, src []float64) {
	for i := range src {
		dst[i] = float32(src[i])
	}
}
