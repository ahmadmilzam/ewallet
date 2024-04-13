package store

import (
	"context"
	"testing"
	"time"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/randomizer"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountTx(t *testing.T) {
	a := &entity.Account{}
	err := randomizer.RandomAccountData(a)
	require.NoError(t, err)

	now := time.Now()
	a.CreatedAt = now
	a.UpdatedAt = now

	w := &entity.Wallet{
		ID:           uuid.New().String(),
		AccountPhone: a.Phone,
		Balance:      0.00,
		Type:         "CASH",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err1 := testStore.CreateAccountTx(context.Background(), a, w)

	require.NoError(t, err1)
	// fmt.Println(ar, wr)
	// require.Equal(t, d.Phone, ac.Phone)
	// require.Equal(t, d.Name, ac.Name)
	// require.Equal(t, d.Email, ac.Email)
	// require.Equal(t, d.Role, ac.Role)
	// require.Equal(t, d.Status, ac.Status)
}
