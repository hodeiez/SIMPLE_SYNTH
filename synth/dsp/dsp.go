package dsp

import (
	"log"

	"github.com/gordonklaus/portaudio"
	"hodei.naiz/simplesynth/synth/generator"
)

//** for now we run just one oscillator
type DspConf struct {
	BufferSize int
	//Osc        *generator.Osc
	//Osc2       *generator.Osc
	Oscs []*generator.Osc
}

//TODO: review and fix the volume and amplitude
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
	todivide := len(dspConf.Oscs)

	for _, oscillators := range dspConf.Oscs {
		oscillators.Osc.Amplitude -= (oscillators.Osc.Amplitude / float64(todivide))

	}
	/* dspConf.Osc.Osc.Amplitude /= 8
	dspConf.Osc2.Osc.Amplitude /= 8 */
	for {
		// populate the out buffer
		for _, oscillators := range dspConf.Oscs {
			if err := oscillators.Osc.Fill(oscillators.Buf); err != nil {
				log.Printf("error filling up the buffer")
			}

		}
		/* if err := dspConf.Osc.Osc.Fill(dspConf.Osc.Buf); err != nil {
			log.Printf("error filling up the buffer")
		}
		if err := dspConf.Osc2.Osc.Fill(dspConf.Osc2.Buf); err != nil {
			log.Printf("error filling up the buffer")
		} */

		/* 	transforms.Gain(dspConf.Osc.Buf, 20) */

		//NoteOn(*noteOn, dspConf)
		f64ToF32Mixing(out, dspConf)
		//f64ToF32Copy(out, dspConf)
		//	f64ToF32Copy(out2, dspConf.Osc2.Buf.Data)

		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}

	}

}

/* func f64ToF32Copy(dst []float32, src []float64, src2 []float64) {
	for i := range src {
		dst[i] = float32(src[i]) + float32(src2[i])
	}
} */
/**************
osc[0][i]+osc[1][i]

*/

func f64ToF32Mixing(dst []float32, src DspConf) {

	for i := range src.Oscs[0].Buf.Data {
		sum := float32(0.0)
		for _, el := range src.Oscs {
			sum += float32(el.Buf.Data[i])
			dst[i] = sum

		}

	}

}
