package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
)

type AccountUsecaseInterface interface {
	CreateAccount(ctx context.Context, params CreateAccountReqParams) (entity.Account, entity.Wallet, error)
	GetAccount(ctx context.Context, phone string) (entity.Account, error)
	// ListBlogs(ctx context.Context, args ListBlogsParams) (*ListBlogsResponse, error)
}

type AccountUsecase struct {
	store entity.QueryStore
}

func NewAccountUsecase(store entity.QueryStore) AccountUsecaseInterface {
	return &AccountUsecase{
		store: store,
	}
}

func (usecase *AccountUsecase) CreateAccount(ctx context.Context, params CreateAccountReqParams) (entity.Account, entity.Wallet, error) {

	aID := uuid.New().String()
	wID := uuid.New().String()
	cAt := time.Now()
	uAt := time.Now()

	ac := entity.Account{
		ID:        aID,
		Name:      params.Name,
		Phone:     params.Phone,
		Email:     params.Email,
		Role:      "REGISTERED",
		Status:    "ACTIVE",
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}

	wl := entity.Wallet{
		ID:        wID,
		AccountId: aID,
		Balance:   0.00,
		Type:      "CASH",
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}

	err := usecase.store.CreateAccountWallet(ctx, ac, wl)

	if err != nil {
		return entity.Account{}, entity.Wallet{}, err
	}

	return ac, wl, nil
}

func (usecase *AccountUsecase) GetAccount(ctx context.Context, phone string) (entity.Account, error) {
	ac, err := usecase.store.FindAccountByPhone(ctx, phone)

	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		err = fmt.Errorf("%s: %w", httpres.GenericInternalError, err)
		return entity.Account{}, err
	}

	if err != nil {
		err = fmt.Errorf("%s: %w", httpres.GenericNotFound, err)
		return entity.Account{}, err
	}

	return ac, nil
}

// func (usecase *AccountUsecase) generateCorrelationId(max int) string {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)

// 	tNow := time.Now().UnixNano()

// 	random := r.Intn(max)
// 	return strconv.Itoa(int(tNow)) + strconv.Itoa(random)

// }

// func (usecase *AccountUsecase) CreateBlog(ctx context.Context, body string) (entity.Account, error) {

// 	blog, err := usecase.store.CreateBlog(ctx, sql.NullString{String: description, Valid: true})

// 	if err != nil {
// 		return nil, fmt.Errorf("AccountUsecase - uc.usecase.CreateBlog.: %w", err)
// 	}

// 	return &blog, nil
// }

// // ListBlogs -.
// func (usecase *AccountUsecase) ListBlogs(ctx context.Context, args intfaces.ListBlogsParams) (*intfaces.ListBlogsResponse, error) {

// 	page, err := utils.StringToInt32(args.Page)

// 	if err != nil {
// 		return nil, errors.New("enter a valid type for the pageId query parameter")
// 	}

// 	limit, err := utils.StringToInt32(args.Limit)

// 	if err != nil {
// 		return nil, errors.New("enter a valid type for the pageSize query parameter")
// 	}

// 	Limit, Offset := utils.PaginatorParams(page, limit)

// 	blogs, err := usecase.store.ListBlog(ctx, sqlc.ListBlogParams{
// 		Limit:  Limit,
// 		Offset: Offset,
// 	})

// 	if err != nil {
// 		return nil, fmt.Errorf("AccountUsecase - bank - uc.usecase.ListBlogs: %w", err)
// 	}

// 	nextPage, previousPage := utils.PaginatorPages(ctx, page, limit, len(blogs))

// 	return &intfaces.ListBlogsResponse{Blog: blogs, NextPage: nextPage, PreviousPage: previousPage}, nil
// }

// func CreateAccount(payload *AccountPayload) *Account {
// 	return
// }
