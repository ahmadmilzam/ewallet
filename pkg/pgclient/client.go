package pgclient

import (
	"time"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	sqlxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/jmoiron/sqlx"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Client struct {
	*sqlx.DB
}

func (c *Client) Close() {
	_ = c.DB.Close()
}

func New() (*Client, error) {
	dbConfig := config.GetDBConfig()

	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithDBMPropagation(tracer.DBMPropagationModeFull), sqltrace.WithServiceName("ewallet.db"))
	db, err := sqlxtrace.Open("postgres", dbConfig.GetConnectionURI(), sqltrace.WithDBMPropagation(tracer.DBMPropagationModeFull))

	if err != nil {
		logger.ErrAttr(err)
	}

	db.SetMaxIdleConns(dbConfig.Connection.MaxIdleConn)
	db.SetMaxOpenConns(dbConfig.Connection.MaxOpenConn)
	lifeTime := time.Second * time.Duration(dbConfig.Connection.MaxLifeTimeConn)
	db.SetConnMaxLifetime(lifeTime)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Client{DB: db}, nil
}
