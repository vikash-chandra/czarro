package api

import (
	"fmt"
	"net/http"

	db "github.com/czarro/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	FirstName   string `json:"first_name" binding:"required,alphanum"`
	LastName    string `json:"last_name" binding:"required,alphanum"`
	Phone       string `json:"phone" binding:"required,num"`
	CountryCode int32  `json:"country_code" binding:"required,num"`
}

func (s *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	logger.Info("===>>>" + ctx.Request.RequestURI)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	arg := db.CreateUserParams{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Phone:       req.Phone,
		CountryCode: req.CountryCode,
	}
	fmt.Println(arg)
	User, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	logger.Info(fmt.Sprintf("%+v", User))
	ctx.JSON(http.StatusOK, User)
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) GetUser(ctx *gin.Context) {
	var req getUserRequest
	logger.Info("===>>>" + ctx.Request.RequestURI)
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	User, err := s.store.GetUser(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	logger.Info(fmt.Sprintf("%+v", User))
	ctx.JSON(http.StatusOK, User)
}

// Ishu1708!
// wp_termmeta product
// wp_users
// wp_yoast_indexable
