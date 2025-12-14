package main

import "sync"

type ShardConcurrentMap[K comparable, V any] struct {
	mu        []sync.RWMutex
	m         []map[K]V
	shardNum  int
	shardFunc func(key K) int
}

func NewShardConcurrentMap[K comparable, V any](shardNum int, shardFunc func(key K) int) *ShardConcurrentMap[K, V] {
	m := new(ShardConcurrentMap[K, V])
	m.mu = make([]sync.RWMutex, shardNum)
	m.m = make([]map[K]V, shardNum)
	for idx := range m.m {
		m.m[idx] = make(map[K]V, 0)
	}
	m.shardNum = shardNum
	m.shardFunc = shardFunc
	return m
}

func (m *ShardConcurrentMap[K, V]) Load(key K) (V, bool) {
	shard := m.shardFunc(key) % m.shardNum
	m.mu[shard].RLock()
	defer m.mu[shard].RUnlock()
	v, ok := m.m[shard][key]
	return v, ok
}

func (m *ShardConcurrentMap[K, V]) Store(key K, value V) {
	shard := m.shardFunc(key) % m.shardNum
	m.mu[shard].Lock()
	defer m.mu[shard].Unlock()
	m.m[shard][key] = value
}
