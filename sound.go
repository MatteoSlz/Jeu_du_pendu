package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	ctrl        *beep.Ctrl
	speakerOnce sync.Once
	speakerErr  error
)

func initSpeaker(sampleRate beep.SampleRate) error {
	speakerOnce.Do(func() {
		speakerErr = speaker.Init(
			sampleRate,
			sampleRate.N(50*time.Millisecond),
		)
	})

	return speakerErr
}

func stopSound() {
	if ctrl == nil {
		return
	}

	speaker.Lock()
	ctrl.Streamer = nil
	ctrl.Paused = true
	speaker.Unlock()

	ctrl = nil
}

func playMusic(path string, loop bool) {
	stopSound()

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Erreur ouverture fichier :", err)
		return
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		f.Close()
		fmt.Println("Erreur décodage MP3 :", err)
		return
	}

	if err := initSpeaker(format.SampleRate); err != nil {
		streamer.Close()
		f.Close()
		fmt.Println("Erreur initialisation speaker :", err)
		return
	}

	var music beep.Streamer

	if loop {
		music = beep.Loop(-1, streamer)
	} else {
		music = streamer
	}

	ctrl = &beep.Ctrl{
		Streamer: music,
		Paused:   false,
	}

	speaker.Play(ctrl)

	fmt.Println("Musique lancée :", path)
}

func PlaySoundLooseMario() {
	playMusic("audio/LooseMario.mp3", true)
}

func PlaySoundFaux() {
	playMusic("audio/Faux.mp3", true)
}

func PlaySoundValidation() {
	playMusic("audio/Validation.mp3", true)
}
