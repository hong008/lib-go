package timeUtil

import "time"

type Timer interface {
	Remain(unit string) int32
	Stop()
	After(t time.Duration, handler func())
	Every(t time.Duration, handle func())
}

type myTimer struct {
	timer   *time.Timer
	endTime time.Time
	handler func()
}

func NewTimer() Timer {
	return new(myTimer)
}

func (m *myTimer) Remain(unit string) int32 {
	r := m.endTime.Sub(time.Now())
	switch unit {
	case "H":
		return int32(r.Hours())
	case "M":
		return int32(r.Minutes())
	default:
		return int32(r.Seconds())
	}
}

func (m *myTimer) Stop() {
	if m.timer == nil {
		return
	}
	m.timer.Stop()
	m.timer = nil
}

func (m *myTimer) After(t time.Duration, handler func()) {
	m.endTime = time.Now().Add(t)
	m.handler = handler
	time.AfterFunc(t, handler)
}

func (m *myTimer) Every(t time.Duration, handle func()) {
	ticker := time.NewTicker(t)
	go func() {
		for range ticker.C {
			handle()
		}
	}()
}
