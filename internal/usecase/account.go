package usecase

import (
	"context"
	"time"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
)

type AccountUsecaseInterface interface {
	CreateAccount(ctx context.Context, params CreateAccountReqParams) (entity.Account, error)
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

func (usecase *AccountUsecase) CreateAccount(ctx context.Context, params CreateAccountReqParams) (entity.Account, error) {

	ac := entity.Account{
		ID:        uuid.New().String(),
		Name:      params.Name,
		Phone:     params.Phone,
		Role:      "BASIC",
		Status:    "ACTIVE",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	ac, err := usecase.store.CreateAccount(ctx, ac)
	if err != nil {
		return entity.Account{}, err
	}

	return ac, nil
}

func (usecase *AccountUsecase) GetAccount(ctx context.Context, phone string) (entity.Account, error) {

	ac, err := usecase.store.FindAccountByPhone(ctx, phone)
	if err != nil {
		return entity.Account{}, err
	}

	return ac, nil
}

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
