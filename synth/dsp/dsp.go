package dsp

import (
	"log"

	"github.com/gordonklaus/portaudio"
	"hodei.naiz/simplesynth/synth/generator"
)

//** for now we run just one oscillator
type DspConf struct {
	BufferSize int

	VM *generator.VoiceManager
}

//TODO: review and fix the volume and amplitude
func Run(dspConf DspConf) {

	portaudio.Initialize()

	defer portaudio.Terminate()
	out := make([]float32, dspConf.BufferSize)

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

	for {
		// populate the out buffer

		for _, oscillators := range dspConf.VM.Voices {
			if err := oscillators.Oscillator.Osc.Fill(oscillators.Oscillator.Buf); err != nil {
				log.Printf("error filling up the buffer")
			}

		}

		f64ToF32Mixing(out, dspConf)

		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}

	}

}

func f64ToF32Mixing(dst []float32, src DspConf) {

	for i := range src.VM.Voices[0].Oscillator.Buf.Data {
		sum := float32(0.0)
		for _, el := range src.VM.Voices {
			sum += float32(el.Oscillator.Buf.Data[i])
			dst[i] = sum

		}

	}

}
