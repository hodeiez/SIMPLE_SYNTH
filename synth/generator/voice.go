package generator

import (
	"time"

	"hodei.naiz/simplesynth/synth/midi"
)

type Voice struct {
	Oscillator  *Osc
	TimeControl float64
	Midi        midi.MidiMsg
	ADSR        *ADSR
	Quit        chan bool
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
	adsr := ADSR{AttackTime: *controller.ADSRcontrol.AttackTime, DecayTime: *controller.ADSRcontrol.DecayTime,
		SustainAmp: *controller.ADSRcontrol.SustainAmp, ReleaseTime: *controller.ADSRcontrol.ReleaseTime, ControlAmp: 0.01}

	osc.Osc.Amplitude = 0.0
	Midi := midi.MidiMsg{Key: -1, On: false}
	timeControl := 0.0
	quit := make(chan bool, 1)
	//

	return &Voice{Oscillator: &osc, Midi: Midi, TimeControl: timeControl, ADSR: &adsr, Quit: quit}
}
func PolyInit(bufferSize int, amountOfVoices int, controller Controls) VoiceManager {
	var voices []*Voice
	i := 0
	for i < amountOfVoices {
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
		time.Sleep(1 * time.Nanosecond)
		if voice.Midi.Key == midimsg.Key {
			return FoundKey{indx, midimsg.Key}

		}
	}
	return FoundKey{-1, -1}
}
func VoiceOnNoteOn(vManager VoiceManager, midimsg midi.MidiMsg, controller Controls, basePitch float64) {

	foundKey := VoicesHasKey(midimsg, vManager)
	if midimsg.On && foundKey.Index == -1 {
		voiceIndex := vManager.FindFreeVoice()
		if voiceIndex != -1 {

			vManager.Voices[voiceIndex].Midi = midimsg

			ChangeFreq(vManager.Voices[voiceIndex].Midi, vManager.Voices[voiceIndex].Oscillator, basePitch)
			vManager.Voices[voiceIndex].RunADSR(&vManager.Voices[voiceIndex].Oscillator.Osc.Amplitude, controller, &vManager.Voices[voiceIndex].Oscillator.Osc.Amplitude, "AMP")

		}
	}
}
func VoiceOnNoteOff(vManager VoiceManager, midimsg midi.MidiMsg, controller Controls) {
	foundKey := VoicesHasKey(midimsg, vManager)
	if !midimsg.On && foundKey.Index != -1 {

		vManager.Voices[foundKey.Index].Midi = midimsg
		vManager.Voices[foundKey.Index].Midi.Key = -1

		vManager.Voices[foundKey.Index].RunADSR(&vManager.Voices[foundKey.Index].Oscillator.Osc.Amplitude, controller, &vManager.Voices[foundKey.Index].Oscillator.Osc.Amplitude, "AMP")

	}

}

func RunPolly(vManager VoiceManager, midimsg midi.MidiMsg, controller Controls, basePitch float64) {

	VoiceOnNoteOn(vManager, midimsg, controller, basePitch)
	VoiceOnNoteOff(vManager, midimsg, controller)

}

func (voice *Voice) increaseAmp(amp float64) {
	voice.Oscillator.Osc.Amplitude += amp
}
func (voice *Voice) decreaseAmp(amp float64) {
	voice.Oscillator.Osc.Amplitude -= amp
}

//TODO: refactor types (enums?)
func (voice *Voice) adsrAction(selector string, actionType string, rate float64) {
	switch selector {
	case "INCREASE":
		switch actionType {
		case "AMP":
			voice.increaseAmp(rate)
		}
	case "DECREASE":
		switch actionType {
		case "AMP":
			voice.decreaseAmp(rate)
		}
	}
}
