package api

import (
	"fmt"
	"net/http"

	db "github.com/czarro/db/sqlc"
	"github.com/gin-gonic/gin"
)

const (
	DefaultRoleId   int32 = 100
	DefaultStatusId int32 = 1
)

type CreateUserRequest struct {
	FirstName   string `json:"firstName" binding:"required,alphanum"`
	LastName    string `json:"lastName" binding:"required,alphanum"`
	Phone       string `json:"phone" binding:"required"`
	CountryCode int32  `json:"countryCode" binding:"required"`
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
		RoleID:      DefaultRoleId,
		StatusID:    DefaultStatusId,
	}
	msg := fmt.Sprintf("arg %#+v", arg)
	fmt.Println(msg)
	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	logger.Info(fmt.Sprintf("%+v", user))
	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) GetUser(ctx *gin.Context) {
	var req getUserRequest
	logger.Info("===>>>" + ctx.Request.RequestURI)
	if err := ctx.ShouldBindUri(&req); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(req.ID)
	User, err := s.store.GetUser(ctx, req.ID)
	if err != nil {
		fmt.Println(err.Error())
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
