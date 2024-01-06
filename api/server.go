package api

import (
	"time"

	db "github.com/czarro/db/sqlc"
	"github.com/czarro/util"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server serve HTTP request
type Server struct {
	store  db.Store
	router *gin.Engine
}

var logger *zap.Logger

func NewServer(store db.Store, config util.Config) *Server {
	server := &Server{
		store: store,
	}
	apiPath := config.ApiServiceRoute
	router := gin.Default()
	logger, _ = zap.NewProduction()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	router.POST(apiPath+"/customers", server.CreateCustomer)
	router.GET(apiPath+"/customers/:id", server.GetCustomer)
	router.PATCH(apiPath + "/customer")

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
