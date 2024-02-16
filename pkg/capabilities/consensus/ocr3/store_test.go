package consensus

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

func TestOCR3Store(t *testing.T) {
	ctx := tests.Context(t)
	rea := time.Second
	n := time.Now()
	fc := clockwork.NewFakeClockAt(n)

	s := newStore(rea, fc)
	rid := uuid.New().String()
	req := &request{
		WorkflowExecutionID: rid,
		ExpiresAt:           n.Add(10 * time.Second),
	}

	t.Run("add", func(t *testing.T) {
		err := s.add(ctx, req)
		require.NoError(t, err)
	})

	t.Run("get", func(t *testing.T) {
		got, err := s.get(ctx, rid)
		require.NoError(t, err)
		assert.Equal(t, req, got)
	})

	t.Run("evict", func(t *testing.T) {
		wasPresent := s.evict(ctx, rid)
		assert.True(t, wasPresent)
		assert.Len(t, s.requests, 0)

		// evicting doesn't remove from the list of requestIDs
		assert.Len(t, s.requestIDs, 1)
	})

	t.Run("firstN, evicts removed items", func(t *testing.T) {
		r, err := s.firstN(ctx, 1)
		assert.Nil(t, r)
		assert.ErrorContains(t, err, "queue is empty")
		assert.Len(t, s.requestIDs, 0)
	})

	t.Run("firstN, zero batch size", func(t *testing.T) {
		_, err := s.firstN(ctx, 0)
		assert.ErrorContains(t, err, "batchsize cannot be 0")
	})

	t.Run("firstN, batchSize larger than queue", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			err := s.add(ctx, &request{WorkflowExecutionID: uuid.New().String(), ExpiresAt: n.Add(1 * time.Hour)})
			require.NoError(t, err)
		}
		items, err := s.firstN(ctx, 100)
		require.NoError(t, err)
		assert.Len(t, items, 10)
	})
}
