package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/czarro/db/sqlc"
	"github.com/czarro/logger"
	"github.com/gin-gonic/gin"
)

const (
	DefaultRoleId   int32 = 100 // user
	DefaultStatusId int32 = 1   // active
)

type CreateUserRequest struct {
	FirstName   string `json:"firstName" binding:"required,alphanum"`
	LastName    string `json:"lastName" binding:"required,alphanum"`
	Phone       string `json:"phone" binding:"required"`
	CountryCode int32  `json:"countryCode" binding:"required"`
}

func (s *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	logger.Info(ctx.Request.RequestURI)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
	logger.Info(msg)
	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	logger.Debug(fmt.Sprintf("%+v", user))
	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) GetUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	logger.Info(fmt.Sprintf("id is %+v", req.ID))
	user, err := s.store.GetUser(ctx, req.ID)
	if err != nil {
		logger.Error(err.Error())
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	logger.Debug(fmt.Sprintf("%+v", user))
	ctx.JSON(http.StatusOK, user)
}

type listUsersRequest struct {
	PageId   int32 `form:"pageId" binding:"required,min=1"`
	PageSize int32 `form:"pageSize"`
}

func (s *Server) GetListUser(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	logger.Info(fmt.Sprintf("req is %+v", req))
	arg := db.ListusersParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	logger.Info(fmt.Sprintf("arg is %+v", arg))
	users, err := s.store.Listusers(ctx, arg)
	if err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	logger.Debug(fmt.Sprintf("%+v", users))
	ctx.JSON(http.StatusOK, users)
}

// Ishu1708!
// wp_termmeta product
// wp_users
// wp_yoast_indexable
