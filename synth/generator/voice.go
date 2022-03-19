package generator

import (
	"log"

	"hodei.naiz/simplesynth/synth/midi"
)

type Voice struct {
	Oscillator  *Osc
	TimeControl *float64
	Midi        midi.MidiMsg
	ADSR        *ADSR
}

type VoiceManager struct {
	Voices []*Voice
}
type FoundKey struct {
	Index int
	Key   int
}

func setupVoice(bufferSize int, controller Controls) *Voice {
	osc := Oscillator(bufferSize)
	adsr := ADSR{AttackTime: *controller.ADSRcontrol.AttackTime, DecayTime: *controller.ADSRcontrol.DecayTime, SustainAmp: *controller.ADSRcontrol.SustainAmp, ReleaseTime: *controller.ADSRcontrol.ReleaseTime, ControlAmp: 0.01}
	osc.Osc.Amplitude = 0.01
	Midi := midi.MidiMsg{Key: -1, On: false}
	timeControl := 0.0
	return &Voice{Oscillator: &osc, Midi: Midi, TimeControl: &timeControl, ADSR: &adsr}
}
func PolyInit(bufferSize int, amountOfVoices int, controller Controls) VoiceManager {
	var voices []*Voice
	i := 0
	for i <= amountOfVoices {
		voices = append(voices, setupVoice(bufferSize, controller))

		i++
	}
	return VoiceManager{Voices: voices}
}
func (vManager *VoiceManager) FindFreeVoice() int {
	key := -1
	for index, voice := range vManager.Voices {
		if !voice.Midi.On {
			key = index
			break

		}
	}
	return key
}
func VoicesHasKey(midimsg midi.MidiMsg, vManager VoiceManager) FoundKey {
	for indx, voice := range vManager.Voices {
		if voice.Midi.Key == midimsg.Key {
			return FoundKey{indx, midimsg.Key}

		}
	}
	return FoundKey{-1, -1}
}
func VoiceOnNoteOn(vManager VoiceManager, midimsg midi.MidiMsg) {
	foundKey := VoicesHasKey(midimsg, vManager)
	if midimsg.On && foundKey.Index == -1 {
		voiceIndex := vManager.FindFreeVoice()
		if voiceIndex == -1 {
			log.Println("Not free", vManager.Voices[0], vManager.Voices[1])

		} else {
			vManager.Voices[voiceIndex].Midi = midimsg
			ChangeFreq(vManager.Voices[voiceIndex].Midi, vManager.Voices[voiceIndex].Oscillator)
			//value := float64(1000.00)
			/* if *vManager.Voices[voiceIndex].TimeControl < value {
				vManager.Voices[voiceIndex].Oscillator.Osc.Amplitude += 0.01
			} */
			log.Println("TCONTROL", vManager.Voices[voiceIndex].TimeControl)

		}
	}
}
func VoiceOnNoteOff(vManager VoiceManager, midimsg midi.MidiMsg) {
	foundKey := VoicesHasKey(midimsg, vManager)
	if !midimsg.On && foundKey.Index != -1 {

		vManager.Voices[foundKey.Index].Midi.On = false
		vManager.Voices[foundKey.Index].Midi.Key = -1
		/* if *vManager.Voices[foundKey.Index].TimeControl == 0.0 {
			vManager.Voices[foundKey.Index].Oscillator.Osc.Amplitude = 0.0000
		} */
		//	log.Println("Assigned", vManager.Voices[0], vManager.Voices[1])
	}
}

func RunPolly(vManager VoiceManager, midimsg midi.MidiMsg, controller Controls) {
	VoiceOnNoteOn(vManager, midimsg)
	VoiceOnNoteOff(vManager, midimsg)
	//vManager.timeControlCounter()
	vManager.adsrRun(controller)

	//log.Println("Assigned", vManager.Voices[0], vManager.Voices[1])
}

/*
func (vManager *VoiceManager) timeControlCounter() {
	for _, voice := range vManager.Voices {
		if voice.Midi.On {
			*voice.TimeControl++
		} else if !voice.Midi.On {
			*voice.TimeControl = 0
		}
	}
} */
func (vManager *VoiceManager) adsrRun(controller Controls) {
	for _, voice := range vManager.Voices {
		voice.ADSR.ADSRforPoly(voice.Midi, voice.Oscillator, voice.TimeControl, controller.ADSRcontrol)
	}
}
