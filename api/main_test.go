package api

import (
	"os"
	"testing"

	db "github.com/czarro/db/sqlc"
	"github.com/czarro/logger"
	"github.com/czarro/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var config util.Config

func newTestServer(t *testing.T, store db.Store) *Server {
	server := NewServer(store, config)
	require.NotNil(t, server)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	// config = util.Config{Environment: "development"}
	logger.CZLoggerInit(config)
	os.Exit(m.Run())
}
