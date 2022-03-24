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
	Off bool
}

func RunMidi(val chan MidiMsg) {
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

			val <- ToMidiMsg(msg.String())

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
func ToMidiMsg(message string) MidiMsg {
	thekey, errK := strconv.ParseInt(strings.Fields(message)[4], 10, 64)
	isOff := strings.Contains(strings.Fields(message)[0], "channel.NoteOff")
	isOn := strings.Contains(strings.Fields(message)[0], "NoteOn")
	must(errK)
	return MidiMsg{Key: int(thekey), On: isOn, Off: isOff}
}
