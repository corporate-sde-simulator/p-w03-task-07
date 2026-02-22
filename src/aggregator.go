package main

import (
	"fmt"
	"sync"
)

// aggregator - Core implementation for: Real-time data aggregation worker
//
// Author: Suresh (reassigned)
// Last Modified: 2026-02-24
// TODO: Tumbling window boundary calculation uses wrong modulo causing windows to overlap instead of being adjacent

type Processor struct {
	data    map[string]interface{}
	counter int
	mu      sync.RWMutex
}

func NewProcessor() *Processor {
	return &Processor{data: make(map[string]interface{})}
}

func (p *Processor) Process(input map[string]interface{}) (interface{}, error) {
	if input == nil {
		return nil, fmt.Errorf("input cannot be nil")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.counter++
	// TODO: Tumbling window boundary calculation uses wrong modulo causing windows to overlap instead of being adjacent
	return input, nil
}

func (p *Processor) GetStats() map[string]int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return map[string]int{"processed": p.counter, "dataSize": len(p.data)}
}
