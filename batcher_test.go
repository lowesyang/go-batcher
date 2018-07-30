package go_batcher

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBatchMgr(t *testing.T) {
	batch := NewBatch("TEST", 10, 0, nil)
	batcher := NewBatcher()
	batcher.AddBatch(batch)
	assert.Equal(t, batcher.GetBatch("TEST"), batch)
	batcher.DelBatch("TEST")
	assert.Nil(t, batcher.GetBatch("TEST"))
}
