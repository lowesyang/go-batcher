package go_batcher

import (
	"testing"
	"time"
	"github.com/magiconair/properties/assert"
)

func TestBatcherBySingle(t *testing.T) {
	var maxCap = 10
	var res int
	ch := make(chan int)
	batcher := NewBatch("TEST", maxCap, 5*time.Millisecond, func(batch []interface{}) {
		res = len(batch)
		ch <- res
	})
	for i := 0; i < maxCap-2; i++ {
		batcher.Push(i)
	}

	res = <-ch
	assert.Equal(t, res, maxCap-2)

	res = 0
	// The length of batch is actually larger than batch max cap
	for i := 0; i < maxCap+2; i++ {
		batcher.Push(i)
	}
	res = <-ch
	assert.Equal(t, res, maxCap)
	res = <-ch
	assert.Equal(t, res, 2)
}

func TestBatcherByBatch(t *testing.T) {
	var maxCap = 10
	var testBatch = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var res int
	ch := make(chan int)
	batcher := NewBatch("TEST", maxCap, 5*time.Millisecond, func(batch []interface{}) {
		res = len(batch)
		ch <- res
	})
	batcher.Batch(testBatch)
	assert.Equal(t, res, 0)

	res = <-ch
	assert.Equal(t, res, maxCap)

	testBatch = append(testBatch, 11, 12)
	batcher.Batch(testBatch)
	res = <-ch
	assert.Equal(t, res, maxCap)

}
