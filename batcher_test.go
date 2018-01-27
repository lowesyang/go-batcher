package go_batcher

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBatchMgr(t *testing.T) {
	batch := NewBatch("TEST", 10, 0, nil)
	batcher := NewBatcher()
	batcher.AddBatcher(batch)
	assert.Equal(t, batcher.GetBatcher("TEST"), batch)
	batcher.DelBatcher("TEST")
	assert.Nil(t, batcher.GetBatcher("TEST"))
}
