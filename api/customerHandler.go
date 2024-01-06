package api

import (
	"fmt"
	"net/http"

	db "github.com/czarro/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateCustomerRequest struct {
	FirstName   string `json:"first_name" binding:"required,alphanum"`
	LastName    string `json:"last_name" binding:"required,alphanum"`
	Phone       string `json:"phone" binding:"required,alphanum"`
	CountryCode string `json:"country_code" binding:"required,alphanum"`
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
	customer, err := s.store.CreateCustomer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	logger.Info(fmt.Sprintf("%+v", customer))
	ctx.JSON(http.StatusOK, customer)
}

type getCustomerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) GetCustomer(ctx *gin.Context) {
	var req getCustomerRequest
	logger.Info("===>>>" + ctx.Request.RequestURI)
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	customer, err := s.store.GetCustomer(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	logger.Info(fmt.Sprintf("%+v", customer))
	ctx.JSON(http.StatusOK, customer)
}

// Ishu1708!
// wp_termmeta product
// wp_users
// wp_yoast_indexable
