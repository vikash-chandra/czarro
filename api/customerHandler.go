package api

import (
	"fmt"
	"net/http"

	db "github.com/czarro/db/sqlc"
	"github.com/gin-gonic/gin"
)

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
