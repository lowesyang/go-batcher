package go_batcher

import "sync"

type cbFunc func(batch []interface{})

type Batcher struct {
	batches map[string]*Batch
	mutex    sync.RWMutex
}

func NewBatcher() *Batcher {
	return &Batcher{
		batches: make(map[string]*Batch),
	}
}

func (bm *Batcher) AddBatcher(batch *Batch) {
	bm.mutex.Lock()
	bm.batches[batch.name] = batch
	bm.mutex.Unlock()
}

func (bm *Batcher) GetBatcher(name string) *Batch {
	bm.mutex.RLock()
	defer bm.mutex.RUnlock()
	return bm.batches[name]
}
