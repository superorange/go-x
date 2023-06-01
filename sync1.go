package main

import "sync"

type syncT struct {
	mu   sync.Mutex
	rl   sync.RWMutex
	data string
}

func main() {
	syncT := syncT{data: "111"}
	syncT.mu.Lock()
	defer syncT.mu.Unlock()
	syncT.data = "89123"
	println(syncT.data)

}
