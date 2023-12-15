package api

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/czarro/db/sqlc"
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

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()
	logger, _ = zap.NewProduction()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	logger.Info("my name is vikash")
	router.POST("/customers", server.CreateCustomer)
	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

type CreateCustomerRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
}

func (s *Server) CreateCustomer(ctx *gin.Context) {
	var req CreateCustomerRequest
	logger.Info("===>>>" + ctx.Request.RequestURI)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	arg := db.CreateCustomerParams{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Phone:       req.Phone,
		CountryCode: req.CountryCode,
	}
	fmt.Println(arg)
	account, err := s.store.CreateCustomer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	logger.Info(fmt.Sprintf("%+v", account))
	ctx.JSON(http.StatusOK, account)
}
