package main

import (
	"fmt"
	"sync"
)

type SyncedMap struct {
	data map[string]int
	mu   sync.Mutex
}

func NewSyncedMap() *SyncedMap { // указатель на SyncedMap
	return &SyncedMap{
		data: make(map[string]int),
	}
}

func (sm *SyncedMap) Add(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SyncedMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func (sm *SyncedMap) Exists(key string) bool { // проверяет существует ли ключ
	sm.mu.Lock()
	defer sm.mu.Unlock()
	_, exists := sm.data[key]
	return exists
}

func main() {
	m := NewSyncedMap()
	m.Add("one", 1)
	fmt.Println(m.Exists("one")) // true
	m.Delete("one")
	fmt.Println(m.Exists("one")) // false
}
