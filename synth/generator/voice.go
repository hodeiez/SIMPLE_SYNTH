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
	return &Voice{Oscillator: &osc, Midi: Midi, TimeControl: timeControl, ADSR: &adsr}
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
func VoiceOnNoteOn(vManager VoiceManager, midimsg midi.MidiMsg, controller Controls) {
	//log.Println(midimsg)
	foundKey := VoicesHasKey(midimsg, vManager)
	if midimsg.On && foundKey.Index == -1 {
		voiceIndex := vManager.FindFreeVoice()
		if voiceIndex != -1 {
			vManager.Voices[voiceIndex].Midi = midimsg

			ChangeFreq(vManager.Voices[voiceIndex].Midi, vManager.Voices[voiceIndex].Oscillator)
			vManager.Voices[voiceIndex].runTimeControl()
		}
	}
}
func VoiceOnNoteOff(vManager VoiceManager, midimsg midi.MidiMsg, controller Controls) {
	foundKey := VoicesHasKey(midimsg, vManager)
	if !midimsg.On && foundKey.Index != -1 {
		//allSameKeyOff(&vManager, midimsg, controller)
		vManager.Voices[foundKey.Index].Midi = midimsg

		vManager.Voices[foundKey.Index].stopControl()
	}
	//	}
}
func allSameKeyOff(vManager *VoiceManager, midimsg midi.MidiMsg, controller Controls) {

	for _, voice := range vManager.Voices {
		if voice.Midi.Key == midimsg.Key {
			voice.Midi.On = false

		}
	}
}
func RunPolly(vManager VoiceManager, midimsg midi.MidiMsg, controller Controls) {

	VoiceOnNoteOn(vManager, midimsg, controller)
	VoiceOnNoteOff(vManager, midimsg, controller)

	//	go vManager.runTimer()
	//vManager.adsrRun(controller)
	//	log.Println(*<-vManager.Voices[0].TimeControl, *<-vManager.Voices[1].TimeControl, *<-vManager.Voices[2].TimeControl, *<-vManager.Voices[3].TimeControl, *<-vManager.Voices[4].TimeControl, *<-vManager.Voices[5].TimeControl)
	//	log.Println(*vManager.Voices[0].TimeControl, *vManager.Voices[1].TimeControl, *vManager.Voices[2].TimeControl, *vManager.Voices[3].TimeControl, *vManager.Voices[4].TimeControl, *vManager.Voices[5].TimeControl)
}

func (vManager *VoiceManager) adsrRun(controller Controls) {

	//	for _, voice := range vManager.Voices {

	//	voice.ADSRforPoly(controller.ADSRcontrol)

	//	}
}
func (voice *Voice) runTimeControl() {

	//for voice.Oscillator.Osc.Amplitude < 2 {
	//voice.Oscillator.Osc.Amplitude = 0.1
	go func() { //ATTACK, call to end
		for {
			voice.Oscillator.Osc.Amplitude += 0.00001
			time.Sleep(1 * time.Millisecond)
		}
	}()
	//	time.Sleep(1 * time.Millisecond)
	//voice.TimeControl++

	//	}

}
func (voice *Voice) stopControl() {
	//	voice.TimeControl = 0.0
	voice.Oscillator.Osc.Amplitude = 0.0
}
