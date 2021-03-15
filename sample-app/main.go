package main

import (
	"log"
	"net/http"
	"sync"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mutex sync.Mutex
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.store[name]++
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
