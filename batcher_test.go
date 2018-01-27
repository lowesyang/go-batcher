package go_batcher

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestBatchMgr(t *testing.T) {
	batch := NewBatch("TEST", 10, 0, nil)
	batcher := NewBatcher()
	batcher.AddBatcher(batch)
	assert.Equal(t, batcher.GetBatcher("TEST"), batch)
}
