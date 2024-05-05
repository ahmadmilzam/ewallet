package usecase_test

import (
	"os"
	"testing"
)

// mockEcsAPI = &mockery.EcsAPI{}
// mockFlipServer = &mockery.FlipServerModule{}
// module = cancel.NewCancelModule(mockEcsAPI, mockFlipServer)

// var testAppUsecase *usecase.AppUsecase

/*
	type AppUsecase struct {
		store  store.SQLStoreInterface
		config config.AppConfig
	}

	func NewAppUsecase(s store.SQLStoreInterface, c config.AppConfig) AppUsecaseInterface {
		return &AppUsecase{
			store:  s,
			config: c,
		}
	}
*/
func TestMain(m *testing.M) {
	// _ = config.Load("config", "../../")
	// MockSQLStoreInterface := mockery.NewSQLStoreInterface()
	// AppConfig := config.GetAppConfig()

	// sql := pgclient.New()

	// if err := sql.DB.Ping(); err != nil {
	// 	log.Fatal("cannot ping db: ", err)
	// }
	// testAppUsecase = &usecase.AppUsecase{
	// 	store:  MockSQLStoreInterface,
	// 	config: AppConfig,
	// }
	// testStore = &store.SQLStore{
	// 	DB:            sql,
	// 	QueryCommands: store.NewQueryCommands(sql),
	// }

	os.Exit(m.Run())
}
