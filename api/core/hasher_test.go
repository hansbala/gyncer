package core_test

import (
	"gyncer/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashString(t *testing.T) {
	data := "test@test.com"
	hash, err := core.HashString(data)
	assert.NoError(t, err)
	assert.Equal(t, "f660ab912ec121d1b1e928a0bb4bc61b15f5ad44d5efdc4e1c92a25e99b8e44a", hash)
}
