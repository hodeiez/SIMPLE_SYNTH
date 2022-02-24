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

func RunMidi(midimsg *MidiMsg, appended *[]MidiMsg) { //*reader.Reader {

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
			midimsg.On = strings.Contains(strings.Fields(msg.String())[0], "channel.NoteOn")
			*appended = append(*appended, MidiMsg{int(thekey), midimsg.On})
			if len(*appended) == 10 {
				lastOne := (*appended)[len(*appended)-2]
				*appended = []MidiMsg{lastOne, {int(thekey), midimsg.On}}
			}
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
