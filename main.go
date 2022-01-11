package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Memory struct {
	cache map[string]string
	lock  sync.Mutex
}

func (m *Memory) Get(key string) (string, error) {
	m.lock.Lock()
	result, exists := m.cache[key]
	m.lock.Unlock()
	if !exists {
		log.Println("Does not exist that key ", key)
		return "", errors.New("Not contained that key")
	}
	return result, nil
}

func (m *Memory) Set(key string, value string) (string, error) {
	m.lock.Lock()
	_, exists := m.cache[key]
	m.lock.Unlock()
	if !exists {
		m.lock.Lock()
		m.cache[key] = value
		m.lock.Unlock()

	}
	return m.cache[key], nil
}
func (m *Memory) All() (map[string]string, error) {
	m.lock.Lock()
	result := cache.cache
	m.lock.Unlock()
	return result, nil

}

func NewCache() *Memory {
	return &Memory{
		cache: make(map[string]string),
	}
}

var cache = NewCache()

func returnAllKeysValues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mapAux := map[string]string{}
	mapAux, _ = cache.All()
	jsonString, _ := json.Marshal(mapAux)
	w.Write(jsonString)
	log.Println(r, w)

}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {

	cache.Set("uno", "1")
	cache.Set("dos", "2")
	http.HandleFunc("/", echoString)
	http.HandleFunc("/cache", returnAllKeysValues)

	log.Println(http.ListenAndServe(":10000", nil))

}
