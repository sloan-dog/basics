package hanoibrute

import (
	"bytes"
	"encoding/binary"

	"github.com/bobertlo/go-mpg123/mpg123"
	"github.com/gordonklaus/portaudio"
)

type AudioPlayer struct {
	src     string
	quit    chan int
	decoder *mpg123.Decoder
}

func NewAudioPlayer(src string) *AudioPlayer {
	a := AudioPlayer{
		src:  src,
		quit: make(chan int, 1),
	}

	decoder, err := mpg123.NewDecoder("")
	chk(err)
	a.decoder = decoder

	return &a
}

func (ap *AudioPlayer) Stop() {
	ap.quit <- 1
}

func (ap *AudioPlayer) Play() {
	// Replace terminate routine with some kind of message on a channel
	// fmt.Println("Playing.  Press Ctrl-C to stop.")

	// create mpg123 decoder instance and play it in goroutine
	fileName := ap.src
	chk(ap.decoder.Open(fileName))
	defer ap.decoder.Close()

	// get audio format information
	rate, channels, _ := ap.decoder.GetFormat()

	// make sure output format does not change
	ap.decoder.FormatNone()
	ap.decoder.Format(rate, channels, mpg123.ENC_SIGNED_16)

	// What does this do...
	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]int16, 8192)
	stream, err := portaudio.OpenDefaultStream(0, channels, float64(rate), len(out), &out)
	chk(err)
	defer stream.Close()

	chk(stream.Start())
	defer stream.Stop()
	for {
		audio := make([]byte, 2*len(out))
		_, err = ap.decoder.Read(audio)
		if err == mpg123.EOF {
			break
		}
		chk(err)

		chk(binary.Read(bytes.NewBuffer(audio), binary.LittleEndian, out))
		chk(stream.Write())
		select {
		// check for new messages on yee ole channel
		case <-ap.quit:
			return
		default:
		}
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
