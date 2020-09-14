package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func start(playlistFilepath string) {
	f, err := os.Open(playlistFilepath + "/oaka-Close 2 U.mp3")
	if err != nil {
		fmt.Println(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Playing")
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
	// keep the program running forever
	// select {}
}
