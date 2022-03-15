package dsp

import (
	"log"

	"github.com/gordonklaus/portaudio"
	"hodei.naiz/simplesynth/synth/generator"
)

//** for now we run just one oscillator
type DspConf struct {
	BufferSize int
	Osc        *generator.Osc
	Osc2       *generator.Osc
}

func Run(dspConf DspConf) {

	portaudio.Initialize()

	defer portaudio.Terminate()
	out := make([]float32, dspConf.BufferSize)
	//	out2 := make([]float32, dspConf.BufferSize)

	stream, err := portaudio.OpenDefaultStream(0, 2, 44100, len(out), &out)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("dsp running")
	defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}
	defer stream.Stop()

	/*we divide amplitude by amount of osc*/
	dspConf.Osc.Osc.Amplitude /= 2
	dspConf.Osc2.Osc.Amplitude /= 2
	for {
		// populate the out buffer
		if err := dspConf.Osc.Osc.Fill(dspConf.Osc.Buf); err != nil {
			log.Printf("error filling up the buffer")
		}
		if err := dspConf.Osc2.Osc.Fill(dspConf.Osc2.Buf); err != nil {
			log.Printf("error filling up the buffer")
		}

		/* 	transforms.Gain(dspConf.Osc.Buf, 20) */

		//NoteOn(*noteOn, dspConf)
		f64ToF32Copy(out, dspConf.Osc.Buf.Data, dspConf.Osc2.Buf.Data)
		//	f64ToF32Copy(out2, dspConf.Osc2.Buf.Data)

		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}

	}

}

func f64ToF32Copy(dst []float32, src []float64, src2 []float64) {
	for i := range src {
		dst[i] = float32(src[i]) + float32(src2[i])
	}
}
