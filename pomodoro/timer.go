package main

import (
	"fmt"
	"time"
)

type TimerState int

const (
	Work TimerState = iota
	ShortBreak
	LongBreak
)

func (s TimerState) String() string {
	switch s {
	case Work:
		return "Work"
	case ShortBreak:
		return "Short Break"
	case LongBreak:
		return "Long Break"
	default:
		return "Unknown"
	}
}

type PomodoroConfig struct {
	WorkDuration          time.Duration
	ShortBreakDuration    time.Duration
	LongBreakDuration     time.Duration
	CyclesBeforeLongBreak int
}

func DefaultConfig() PomodoroConfig {
	return PomodoroConfig{
		WorkDuration:          25 * time.Minute,
		ShortBreakDuration:    5 * time.Minute,
		LongBreakDuration:     15 * time.Minute,
		CyclesBeforeLongBreak: 4,
	}
}

type PomodoroTimer struct {
	config       PomodoroConfig
	currentState TimerState
	currentCycle int
	totalCycles  int
	startTime    time.Time
	duration     time.Duration
}

func NewPomodoroTimer(config PomodoroConfig) *PomodoroTimer {
	return &PomodoroTimer{
		config:       config,
		currentState: Work,
		currentCycle: 1,
		totalCycles:  0,
	}
}

func (pt *PomodoroTimer) Start() {
	pt.startTime = time.Now()
	pt.duration = pt.GetDuration(pt.currentState)
}

func (pt *PomodoroTimer) GetDuration(ts TimerState) time.Duration {
	res := time.Duration(0)
	switch ts {
	case Work:
		res = pt.config.WorkDuration
	case ShortBreak:
		res = pt.config.ShortBreakDuration
	case LongBreak:
		res = pt.config.LongBreakDuration
		pt.totalCycles = 1
	}
	return res
}

func (pt *PomodoroTimer) GetRemainingTime() time.Duration {
	elapsed := time.Since(pt.startTime)
	remaining := pt.duration - elapsed

	if remaining <= 0 {
		return 0
	}
	return remaining
}

func (pt *PomodoroTimer) IsFinished() bool {
	return pt.GetRemainingTime() <= 0
}

func (pt *PomodoroTimer) NextState() {

	switch pt.currentState {
	case Work:
		pt.totalCycles++
		if pt.totalCycles%pt.config.CyclesBeforeLongBreak == 0 {
			pt.currentState = LongBreak
		} else {
			pt.currentState = ShortBreak
		}
	case ShortBreak, LongBreak:
		pt.currentState = Work
		if pt.currentState == Work {
			pt.currentCycle++
		}
	}
	pt.startTime = time.Now()
	pt.duration = pt.GetDuration(pt.currentState)
}

func (pt *PomodoroTimer) GetCurrentState() TimerState {
	return pt.currentState
}

func (pt *PomodoroTimer) GetCurrentCycle() int {
	return pt.currentCycle
}

func (pt *PomodoroTimer) GetTotalCycles() int {
	return pt.totalCycles
}

func (pt *PomodoroTimer) FormatRemainingTime() string {
	remaining := pt.GetRemainingTime()
	minutes := int(remaining.Minutes())
	seconds := int(remaining.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
