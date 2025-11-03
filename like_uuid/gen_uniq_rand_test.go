package like_uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IdGeneratorUniqueRandom_Generate(t *testing.T) {
	g := NewIdGeneratorUniqueRandom()

	t.Run("success unique", func(t *testing.T) {
		n := 1_000_000
		prevIds := make(map[Id]struct{}, n)

		for i := 0; i < n; i++ {
			id, err := g.Generate()

			assert.NotEqual(t, IdNil, id)
			assert.NoError(t, err)
			assert.False(t, func() bool { _, ok := prevIds[id]; return ok }() /* assert.NotContains -- O(n) */)

			prevIds[id] = struct{}{}
		}
	})
}
