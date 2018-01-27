package go_batcher

import (
	"time"
	"sync"
	"fmt"
)

type Batch struct {
	name        string // Name of batcher
	mutex       sync.Mutex
	input       chan interface{}
	maxCapacity int           // Max count of channel
	timeout     time.Duration // Timeout of force calling callback func
	callback    cbFunc        // Callback function
}

func NewBatch(name string, maxCapacity int, timeout time.Duration, callback cbFunc) *Batch {
	batcher := &Batch{
		name:        name,
		maxCapacity: maxCapacity,
		timeout:     timeout,
		callback:    callback,
	}
	return batcher
}

func (bc *Batch) String() string {
	return fmt.Sprintf("Batch { name:%s, maxCapacity:%d, timeout:%s }", bc.name, bc.maxCapacity, bc.timeout)
}

// Push single data into batcher
func (bc *Batch) Push(data interface{}) {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	if bc.input == nil {
		bc.input = make(chan interface{}, bc.maxCapacity)
		go bc.run()
	}

	bc.input <- data
}

// Push batch of data into batcher
func (bc *Batch) Batch(batch []interface{}) {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	batchLen := len(batch)
	// If the length of batch is larger than maxCapacity
	if batchLen > bc.maxCapacity {
		batch = batch[:bc.maxCapacity]
	}

	if bc.input == nil {
		bc.input = make(chan interface{}, bc.maxCapacity)
		go bc.run()
	}

	for _, item := range batch {
		bc.input <- item
	}
}

func (bc *Batch) run() {
	var batch []interface{}
	timer := time.NewTimer(bc.timeout)

	for {
		select {
		case <-timer.C:
			bc.callback(batch)
			bc.close()
			return
		case item := <-bc.input:
			batch = append(batch, item)
			if len(batch) == bc.maxCapacity {
				// Callback with batch
				bc.callback(batch)
				// Init batch array
				batch = []interface{}{}
			}
		}
	}
}

func (bc *Batch) close() {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()
	if bc.input != nil {
		close(bc.input)
		bc.input = nil
	}
}
