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

func OnClose() {
	muCloseHooks.Lock()
	for i := range onClose {
		onClose[i]()
	}
	muCloseHooks.Unlock()
}

type Hooks []func()

func (h *Hooks) add(f ...func()) {
	*h = append(*h, f...)
}

var (
	onClose      Hooks
	muCloseHooks sync.Mutex
)
