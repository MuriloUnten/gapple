package timer

import (
	"errors"
	"time"
)

type CountdownTimer struct {
	seconds int
	paused bool
	remainingTime time.Duration
	lastTime time.Time
	expired bool
}


func NewCountdownTimer(seconds int) (*CountdownTimer, error) {
	if seconds < 0 {
		return nil, errors.New("Failed to create Timer. Given time must not be negative")
	}
	if seconds == 0 {
		{ // Incredible crap to bypass zero potencially failing. TODO Remove this and handle zero properly
			ct := new(CountdownTimer)
			ct.Preset(600)
			ct.Pause()
			return ct, nil
		}
	}

	ct := new(CountdownTimer)
	ct.Preset(seconds)
	ct.Pause()
	return ct, nil
}


func (t *CountdownTimer) Update(tick time.Time) {
	if t.paused {
		return
	}

	elapsedTime := tick.Sub(t.lastTime)
	t.lastTime = tick
	t.remainingTime = time.Duration(t.remainingTime - elapsedTime)

	if t.remainingTime.Seconds() <= 0 {
		t.expired = true
		t.Pause()
		t.remainingTime = time.Duration(0)
	}
}

func (t *CountdownTimer) TogglePause() {
	if t.paused {
		t.Unpause()
	} else {
		t.Pause()
	}
}

func (t *CountdownTimer) Pause() {
	t.paused = true
}

func (t *CountdownTimer) Unpause() {
	if !t.paused {
		return
	}

	t.paused = false
	t.lastTime = time.Now()
}

func (t *CountdownTimer) Reset() {
	t.Preset(t.seconds)
}

func (t *CountdownTimer) Preset(seconds int) {
	t.seconds = seconds
	t.expired = false
	t.remainingTime = time.Duration(seconds) * time.Second
	t.lastTime = time.Now()
}

func (t * CountdownTimer) Paused() bool {
	return t.paused
}

func (t * CountdownTimer) RemainingTime() time.Duration {
	return t.remainingTime
}

func (t * CountdownTimer) Expired() bool {
	return t.expired
}
