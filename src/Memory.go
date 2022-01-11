package main

import (
	"errors"
	"log"
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

func (m *Memory) Delete(key ...string) (bool, error) {
	m.lock.Lock()
	for i := 0; i < len(key); i++ {
		delete(cache.cache, key[i])
	}
	m.lock.Unlock()

	return true, nil
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
