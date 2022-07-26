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

	// Asking the user to input a number of seconds and then it is storing that number in the variable
	// seconds
	var seconds int
	fmt.Println("Enter seconds: ")
	fmt.Scan(&seconds)

	// A for loop that counts down from the number of seconds the user inputs
	for i := seconds; i > 0; i-- {
		time.Sleep(1*time.Second)
		fmt.Println(i)

		if i == 1 {// Playing the audio file when the countdown reaches 1
			speaker.Play(streamer)
			select{}
		}
	}
}