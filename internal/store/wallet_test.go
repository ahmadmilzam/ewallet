package store_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWalletById(t *testing.T) {
	id := "001"
	_, err := testStore.FindWalletById(context.Background(), id)
	assert.NoError(t, err)
}
