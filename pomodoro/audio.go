package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type AudioNotifier struct{}

func NewAudioNotifier() *AudioNotifier {
	return &AudioNotifier{}
}

func (an *AudioNotifier) PlayNotification(state TimerState, isStarting bool) error {
	var message string

	if isStarting {
		switch state {
		case Work:
			message = "Time to work! Stay focused."
		case ShortBreak:
			message = "Time for a short break! Relax for 5 minutes."
		case LongBreak:
			message = "Great job! Time for a long break. Relax for 15 minutes."
		}
	} else {
		switch state {
		case Work:
			message = "Work session complete! Great job."
		case ShortBreak:
			message = "Break time is over."
		case LongBreak:
			message = "Long break is over."
		}
	}

	return an.speak(message)
}

func (an *AudioNotifier) speak(message string) error {
	// Print message to console as fallback
	fmt.Printf("🔔 %s\n", message)

	// Create context with timeout to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "say", message)
	if err := cmd.Run(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Printf("Note: Audio notification timed out\n")
		} else {
			fmt.Printf("Note: Could not play audio notification: %v\n", err)
		}
	}

	return nil
}
