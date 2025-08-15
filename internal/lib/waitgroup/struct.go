package waitgroup

import (
	"sync"
)

type WaitGroup struct {
	wg            sync.WaitGroup
	mu            sync.RWMutex
	activeGorutin map[string]struct{}
}

func InitWg() *WaitGroup {
	ag := make(map[string]struct{})
	return &WaitGroup{
		wg:            sync.WaitGroup{},
		mu:            sync.RWMutex{},
		activeGorutin: ag,
	}
}

func (w *WaitGroup) Add(fn string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.wg.Add(1)
	w.activeGorutin[fn] = struct{}{}
}

func (w *WaitGroup) Done(fn string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.wg.Done()
	delete(w.activeGorutin, fn)
}

func (w *WaitGroup) Wait() {
	w.wg.Wait()
}

func (w *WaitGroup) ActiveGorutin() map[string]struct{} {
	return w.activeGorutin
}
