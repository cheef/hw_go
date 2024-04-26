package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache

	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (lru *lruCache) Set(key Key, value interface{}) bool {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	for k, item := range lru.items {
		if k == key {
			item.Value = value
			lru.queue.MoveToFront(item)
			return true
		}
	}

	if lru.capacity == len(lru.items) {
		lru.purge()
	}

	i := lru.queue.PushFront(value)
	lru.items[key] = i

	return false
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	for k, item := range lru.items {
		if k == key {
			lru.queue.MoveToFront(item)
			return item.Value, true
		}
	}

	return nil, false
}

func (lru *lruCache) purge() {
	tail := lru.queue.Back()

	for key, item := range lru.items {
		if item == tail {
			delete(lru.items, key)
			lru.queue.Remove(tail)
		}
	}
}

func (lru *lruCache) Clear() {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	lru.queue = NewList()
	lru.items = make(map[Key]*ListItem, lru.capacity)
}
