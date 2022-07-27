package main

import (
	"log"
	"os"
	"time"
	"fmt"

	// Importing the libraries that are needed to play the audio file
	// go get -u github.com/faiface/beep
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	// Opens the audio file and handles the erron
	file, err := os.Open("song.mp3")
	if err != nil {
		log.Fatal(err)
	}

	// Decoding the mp3 file and setting it to the streamer variable
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initializes the speaker
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Asking the user to input the amount of time that they want the timer to wait before playing the
	// audio file.
	var a int
	fmt.Println("Enter hours: ")
	fmt.Scan(&a)

	var b int
	fmt.Println("Enter minutes: ")
	fmt.Scan(&b)

	var c int
	fmt.Println("Enter seconds: ")
	fmt.Scan(&c)

    // Creating a timer that will wait for the amount of time that the user inputs
	timer := time.NewTimer(time.Duration(a) * time.Hour + 
                           time.Duration(b) * time.Minute +
						   time.Duration(c) * time.Second)

	// speaker.Play starts the streamer but dosen't wait for it to finish
	// to fix this select{} waits for the streamer to end
	<-timer.C// This blocks the timer's channel until the duration has ended
	speaker.Play(streamer)// Play the audio file
	select{}
}
