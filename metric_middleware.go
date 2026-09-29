package main

import (
	"fmt"
	"td_redis/store"
	"time"
)

type MetricMiddleware struct {
	getCalls    int
	setCalls    int
	deleteCalls int
	lenCalls    int
	inner       store.Storer

	getMisses       int
	totalGetLatency time.Duration
	totalSetLatency time.Duration
}

func NewMetricMiddleware(inner store.Storer) *MetricMiddleware {
	return &MetricMiddleware{
		inner: inner,
	}
}

func (m *MetricMiddleware) Get(key string) (string, error) {
	start := time.Now()
	val, err := m.inner.Get(key)
	if err != nil {
		m.getMisses++
	}
	m.getCalls++
	m.totalGetLatency += time.Since(start)
	return val, err
}

func (m *MetricMiddleware) Set(key string, value string) error {
	start := time.Now()
	err := m.inner.Set(key, value)
	m.setCalls++ //side effect

	m.totalSetLatency += time.Since(start)
	return err
}

func (m *MetricMiddleware) Delete(key string) {
	m.deleteCalls++
	m.inner.Delete(key)
}

func (m *MetricMiddleware) Len() int {
	m.deleteCalls++
	return m.inner.Len()
}

func (m *MetricMiddleware) Keys() []string {
	return m.inner.Keys()
}

func (m *MetricMiddleware) Report() {
	fmt.Printf("GET calls: %d. (missed %d)\n", m.getCalls, m.getMisses)
	fmt.Printf("SET calls: %d.", m.setCalls)
	fmt.Printf("DELETE calls: %d.", m.deleteCalls)
	fmt.Printf("LEN calls: %d.", m.lenCalls)

	if m.getCalls > 0 {
		avg := m.totalGetLatency / time.Duration(m.getCalls)
		fmt.Printf("Average get latency: %v.", avg)
	}
	if m.setCalls > 0 {
		avg := m.totalSetLatency / time.Duration(m.setCalls)
		fmt.Printf("Average get latency: %v.", avg)
	}
}
