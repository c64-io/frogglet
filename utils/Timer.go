package utils

import "time"

const FPS_ALPHA float32 = 0.99

type Timer struct {
	previousTime time.Time
	Elapsed      float32
	Fps          float32
	AverageFPS   float32
}

func NewTimer() *Timer {
	return &Timer{
		previousTime: time.Now(),
		AverageFPS:   240,
	}
}

func (t *Timer) Tick() {
	t.Elapsed = float32(time.Since(t.previousTime).Seconds())
	t.previousTime = time.Now()

	t.Fps = 1 / t.Elapsed

	t.AverageFPS = t.AverageFPS*FPS_ALPHA + t.Fps*(1-FPS_ALPHA)
}
