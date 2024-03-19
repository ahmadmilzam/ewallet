package sqlclient_test

import (
	"testing"

	"github.com/ahmadmilzam/ewallet/pkg/sqlclient"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	db := sqlclient.New()
	defer db.Close()

	assert.NotNil(t, db)
}
