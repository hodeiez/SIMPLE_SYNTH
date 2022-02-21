package midi

import (
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	"strconv"
	"strings"

	driver "gitlab.com/gomidi/rtmididrv"
)

type MidiMsg struct {
	Key int  //
	On  bool //

}

func RunMidi(midimsg *MidiMsg) { //*reader.Reader {
	// you would take a real driver here e.g.
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
		// write every message to the out port
		reader.Each(func(pos *reader.Position, msg midi.Message) {
			//	fmt.Printf("%s %s\n", strings.Fields(msg.String())[0], strings.Fields(msg.String())[4])
			thekey, errK := strconv.ParseInt(strings.Fields(msg.String())[4], 10, 64)
			must(errK)
			midimsg.Key = int(thekey)
			midimsg.On = strings.Contains(strings.Fields(msg.String())[0], "channel.NoteOn")

		}),
	)

	r := rd.ListenTo(in)
	for {

		must(r)

	}

}

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}
