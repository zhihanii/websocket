package websocket

import "sync"

func RegisterCloseHooks(h ...func()) {
	muCloseHooks.Lock()
	defer muCloseHooks.Unlock()
	onClose.add(h...)
}

func ResetCloseHooks() {
	muCloseHooks.Lock()
	defer muCloseHooks.Unlock()
	onClose = []func(){}
}

type Hooks []func()

func (h *Hooks) add(f ...func()) {
	*h = append(*h, f...)
}

var (
	onClose      Hooks
	muCloseHooks sync.Mutex
)
