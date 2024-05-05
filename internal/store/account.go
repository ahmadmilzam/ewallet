package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

const (
	createAccountSQL = `
	INSERT INTO accounts
	VALUES(:phone, :name, :email, :role, :status, :coa_type, :created_at, :updated_at)`
	updateAccountSQL = `
	UPDATE accounts
	SET
		phone = :phone,
		name = :name,
		email = :email,
		role = :role,
		status = :status,
		updated_at = :updated_at
	WHERE phone = :phone`
	findAccountByIdSQL           = `SELECT * FROM accounts WHERE phone = $1 LIMIT 1`
	deleteAccountByIdSQL         = `DELETE * FROM accounts WHERE phone = :phone`
	findAccountForUpdateByIdSQL  = `SELECT * FROM accounts WHERE phone = $1 LIMIT 1 FOR UPDATE`
	findAccountAndWalletsByIdSQL = `
	SELECT
		account.*,
		wallet.id "wallet.id",
		wallet.type "wallet.type",
		wallet.balance "wallet.balance",
		wallet.created_at "wallet.created_at",
		wallet.updated_at "wallet.updated_at"
	FROM accounts AS account
	JOIN wallets AS wallet ON account.phone = wallet.account_phone
	WHERE account.phone = $1`
)

func (s *QueryCommands) CreateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	_, err := s.db.NamedExecContext(ctx, createAccountSQL, account)
	if err != nil {
		return nil, fmt.Errorf("CreateAccount: %w", err)
	}

	return account, nil
}

func (s *QueryCommands) UpdateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	_, err := s.db.NamedExecContext(ctx, updateAccountSQL, account)

	if err != nil {
		return nil, fmt.Errorf("UpdateAccount: %w", err)
	}

	return account, nil
}

func (s *QueryCommands) FindAccountForUpdateById(ctx context.Context, phone string) (*entity.Account, error) {
	a := &entity.Account{}
	err := s.db.GetContext(ctx, a, findAccountForUpdateByIdSQL, phone)
	if err != nil {
		return nil, fmt.Errorf("FindAccountForUpdateByPhone: %w", err)
	}

	return a, nil
}

func (s *QueryCommands) FindAccountById(ctx context.Context, phone string) (*entity.Account, error) {
	ma := &entity.Account{}
	err := s.db.GetContext(ctx, ma, findAccountByIdSQL, phone)
	if err != nil {
		return nil, fmt.Errorf("FindAccountById: %w", err)
	}

	return ma, nil
}

func (s *QueryCommands) FindAccountAndWalletsById(ctx context.Context, phone string) ([]entity.AccountWallet, error) {
	var accWallets []entity.AccountWallet
	err := s.db.SelectContext(ctx, &accWallets, findAccountAndWalletsByIdSQL, phone)
	if err != nil {
		return nil, fmt.Errorf("FindAccountAndWalletsById: %w", err)
	}

	return accWallets, nil
}
