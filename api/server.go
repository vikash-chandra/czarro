package api

import (
	"time"

	db "github.com/czarro/db/sqlc"
	"github.com/czarro/logger"
	"github.com/czarro/util"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serve HTTP request
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store, config util.Config) *Server {
	server := &Server{
		store: store,
	}
	apiPath := config.ApiServiceRoute
	router := gin.Default()
	logger := logger.GetCZLogger()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, false))

	if v, ok := binding.Validator.Engine().(*validator.Validate); !ok {
		v.RegisterValidation("customCurrency", validCurrency)
	}
	// https://www.youtube.com/watch?v=GguJODC2cvI

	logger.Info("*********" + apiPath + "*********")
	router.POST(apiPath+"/users", server.CreateUser)
	router.GET(apiPath+"/users", server.GetListUser)
	router.GET(apiPath+"/users/:id", server.GetUser)
	router.POST(apiPath+"/products", server.CreateProduct)
	router.PUT(apiPath+"/products", server.UpdateProduct)
	router.GET(apiPath+"/products", server.GetListProduct)
	router.GET(apiPath+"/products/:id", server.GetProduct)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
