package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	db, err := ConnectDB()
	assert.NotNil(t, db)
	assert.NoError(t, err)
}
