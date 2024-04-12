package store

import (
	"context"
	"testing"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/internal/utils"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountTx(t *testing.T) {
	d := utils.RandomAccountData(entity.Account{})
	w := entity.Wallet{
		ID:        uuid.New().String(),
		AccountId: d.ID,
		Balance:   0.00,
		Type:      "CASH",
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	err := testStore.CreateAccountTx(context.Background(), d, w)

	require.NoError(t, err)
	// fmt.Println(ar, wr)
	// require.Equal(t, d.Phone, ac.Phone)
	// require.Equal(t, d.Name, ac.Name)
	// require.Equal(t, d.Email, ac.Email)
	// require.Equal(t, d.Role, ac.Role)
	// require.Equal(t, d.Status, ac.Status)
}
