package monitorUtil

import (
	"errors"
	"sync"
	"time"
)

type Monitor interface {
	AddFile(path string, interval int, handler func()) error
	DelFile(path string) error
}

type myMonitor struct {
	sync.RWMutex
	fs Files
}

type Files map[string]*myFile

func (m *myMonitor) init() {
	m.fs = make(map[string]*myFile)
}

func NewMonitor() Monitor {
	mon := new(myMonitor)
	mon.init()
	return mon
}

func (m *myMonitor) AddFile(path string, interval int, handler func()) error {
	m.Lock()
	defer m.Unlock()
	if handler == nil {
		return errors.New("handler cannot be nil")
	}
	if _, ok := m.fs[path]; ok {
		return errors.New("already exist")
	}
	f := &myFile{
		interval: interval,
		filePath: path,
		updateOn: time.Now().Unix(),
		handle:   handler,
	}
	m.fs[path] = f
	f.init()
	return nil
}

func (m *myMonitor) DelFile(path string) error {
	m.Lock()
	defer m.Unlock()
	f, ok := m.fs[path]
	if !ok {
		return nil
	}
	f.c.Stop()
	f.c.Remove(f.id)
	delete(m.fs, path)
	return nil
}
