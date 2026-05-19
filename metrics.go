package main

import (
	"sort"
	"sync"
)

type Metrics struct {
	mu       sync.Mutex
	latency  []int
	success  int
	fail     int
	errors   map[string]int
}

func NewMetrics() *Metrics {
	return &Metrics{
		latency: make([]int, 0, 10000),
		errors:  make(map[string]int),
	}
}

func (m *Metrics) AddLatency(v int) {
	m.mu.Lock()
	m.latency = append(m.latency, v)
	m.mu.Unlock()
}

func (m *Metrics) AddSuccess() {
	m.mu.Lock()
	m.success++
	m.mu.Unlock()
}

func (m *Metrics) AddFail(reason string) {
	m.mu.Lock()
	m.fail++
	m.errors[reason]++
	m.mu.Unlock()
}

func (m *Metrics) Snapshot() (p50, p95, p99 int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.latency) == 0 {
		return 0, 0, 0
	}

	cp := append([]int(nil), m.latency...)
	sort.Ints(cp)

	p50 = cp[len(cp)*50/100]
	p95 = cp[len(cp)*95/100]
	p99 = cp[len(cp)*99/100]

	return
}
