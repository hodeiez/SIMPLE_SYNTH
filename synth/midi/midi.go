package midi

import (
	"log"
	"strconv"
	"strings"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"

	driver "gitlab.com/gomidi/rtmididrv"
)

type MidiMsg struct {
	Key int  //
	On  bool //

}

func IsOn(midimsg []MidiMsg) bool {

	if midimsg[len(midimsg)-1].On || !midimsg[len(midimsg)-1].On && midimsg[len(midimsg)-1].Key != midimsg[len(midimsg)-2].Key && midimsg[len(midimsg)-2].On {
		return true
	} else if !midimsg[len(midimsg)-1].On && midimsg[len(midimsg)-1].Key == midimsg[len(midimsg)-2].Key || !midimsg[len(midimsg)-1].On && !midimsg[len(midimsg)-2].On {
		return false
	}
	return false
}
func RunMidi(midimsg *MidiMsg) {
	defer func() {
		if error := recover(); error != nil {
			log.Println("NO MIDI!")
		}
	}()

	drv, err := driver.New()

	must(err)

	// make sure to close all open ports at the end
	defer drv.Close()

	ins, err := drv.Ins()
	must(err)

	outs, err := drv.Outs()
	must(err)

	in, out := ins[0], outs[0]

	must(in.Open())
	must(out.Open())

	defer in.Close()
	defer out.Close()

	rd := reader.New(
		reader.NoLogger(),
		// format every message
		reader.Each(func(pos *reader.Position, msg midi.Message) {

			//log.Printf("FIRST%s %s\n", strings.Fields(msg.String())[0], strings.Fields(msg.String())[4])
			thekey, errK := strconv.ParseInt(strings.Fields(msg.String())[4], 10, 64)
			must(errK)
			midimsg.Key = int(thekey)
			isOn := strings.Contains(strings.Fields(msg.String())[0], "channel.NoteOn")
			midimsg.On = isOn

			/* *midiMessages = append(*midiMessages, MidiMsg{int(thekey), isOn})
			if len(*midiMessages) == 10 {
				lastOne := (*midiMessages)[len(*midiMessages)-2]
				*midiMessages = []MidiMsg{lastOne, {int(thekey), isOn}}
			} */
		}),
	)

	r := rd.ListenTo(in)
	log.Print("midi started listening")

	for {

		must(r)

	}

}

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}
