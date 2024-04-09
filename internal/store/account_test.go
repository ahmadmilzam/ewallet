package store

import (
	"context"
	"testing"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	d := utils.RandomAccountData(entity.Account{})

	ac, err := testStore.CreateAccount(context.Background(), d)

	require.NoError(t, err)
	require.Equal(t, d.Phone, ac.Phone)
	require.Equal(t, d.Name, ac.Name)
	require.Equal(t, d.Email, ac.Email)
	require.Equal(t, d.Role, ac.Role)
	require.Equal(t, d.Status, ac.Status)
}
