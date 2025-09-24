package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Command line flags
	workDuration := flag.Int("work", 25, "Work session duration in minutes")
	shortBreakDuration := flag.Int("short", 5, "Short break duration in minutes")
	longBreakDuration := flag.Int("long", 15, "Long break duration in minutes")
	cyclesBeforeLongBreak := flag.Int("cycles", 4, "Number of work cycles before long break")
	testMode := flag.Bool("test", false, "Test mode with shorter durations (work:5s, short:3s, long:5s)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Pomodoro Timer - Stay focused and productive!\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s -work 25 -short 5 -long 15 -cycles 4\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -test  (for quick testing)\n", os.Args[0])
	}

	flag.Parse()

	// Create configuration
	config := PomodoroConfig{
		WorkDuration:          time.Duration(*workDuration) * time.Minute,
		ShortBreakDuration:    time.Duration(*shortBreakDuration) * time.Minute,
		LongBreakDuration:     time.Duration(*longBreakDuration) * time.Minute,
		CyclesBeforeLongBreak: *cyclesBeforeLongBreak,
	}

	// Test mode for quick testing
	if *testMode {
		fmt.Println("🧪 Test mode enabled - using shorter durations")
		config.WorkDuration = 10 * time.Second
		config.ShortBreakDuration = 5 * time.Second
		config.LongBreakDuration = 8 * time.Second
	}

	// Initialize timer and audio notifier
	timer := NewPomodoroTimer(config)
	notifier := NewAudioNotifier()

	// Handle Ctrl+C gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("🍅 Pomodoro Timer Started!")
	fmt.Printf("Configuration: Work %v, Short Break %v, Long Break %v, Cycles before long break: %d\n\n",
		config.WorkDuration, config.ShortBreakDuration, config.LongBreakDuration, config.CyclesBeforeLongBreak)

	// Main timer loop
	for {
		// Start current session
		fmt.Printf("📍 Starting %s session (Cycle %d)\n", timer.GetCurrentState(), timer.GetCurrentCycle())
		notifier.PlayNotification(timer.GetCurrentState(), true)
		timer.Start()

		// Session timer loop
		notificationFlag := false
		for {
			select {
			case <-sigChan:
				fmt.Println("\n👋 Goodbye! Thanks for using Pomodoro Timer.")
				return
			default:
				if timer.IsFinished() {
					// Session completed
					fmt.Printf("\n✅ %s session complete!\n", timer.GetCurrentState())
					notifier.PlayNotification(timer.GetCurrentState(), notificationFlag)
					notificationFlag = !notificationFlag
					// Move to next state
					timer.NextState()

					notifier.PlayNotification(timer.GetCurrentState(), notificationFlag)
					notificationFlag = !notificationFlag
					// Show stats
					showStats(timer)

					fmt.Println()
					break
				}

				// Display remaining time (only if timer is running)

				fmt.Printf("\r⏰ %s - %s remaining",
					timer.GetCurrentState(),
					timer.FormatRemainingTime())

				time.Sleep(1 * time.Second)
			}
		}
	}
}

func showStats(timer *PomodoroTimer) {
	fmt.Printf("📊 Stats: Completed %d work sessions\n", timer.GetTotalCycles())

	if timer.GetCurrentState() == LongBreak {
		fmt.Println("🎉 Great job! You've earned a long break.")
	}
}
