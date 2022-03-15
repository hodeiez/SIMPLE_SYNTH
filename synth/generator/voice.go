package generator

import (
	"log"

	"hodei.naiz/simplesynth/synth/midi"
)

type Voice struct {
	Oscillator *Osc
	//IsOn       bool
	Midi midi.MidiMsg
	ADSR *ADSR
}

type VoiceManager struct {
	Voices []*Voice
}
type FoundKey struct {
	Index int
	Key   int
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
			log.Println("Assigned", vManager.Voices[0], vManager.Voices[1])

		}
	}
}
func VoiceOnNoteOff(vManager VoiceManager, midimsg midi.MidiMsg) {
	foundKey := VoicesHasKey(midimsg, vManager)
	if !midimsg.On && foundKey.Index != -1 {

		vManager.Voices[foundKey.Index].Midi.On = false
		vManager.Voices[foundKey.Index].Midi.Key = -1

		log.Println("Assigned", vManager.Voices[0], vManager.Voices[1])
	}
}

func RunPolly(vManager VoiceManager, midimsg midi.MidiMsg) {
	VoiceOnNoteOn(vManager, midimsg)
	VoiceOnNoteOff(vManager, midimsg)
}
