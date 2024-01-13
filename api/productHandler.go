package api

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/czarro/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Title       string `json:"title" binding:"required"`
	ShortName   string `json:"shortName" binding:"required"`
	Description string `json:"description" binding:"required"`
	SmsNoti     *bool  `json:"smsNoti" binding:"required"`
	EmailNoti   *bool  `json:"emailNoti" binding:"required"`
	CallNoti    *bool  `json:"callNoti" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Id          int32  `json:"id,omitempty"`
}

func (s *Server) CreateProduct(ctx *gin.Context) {
	var req CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	args := db.CreateProductParams{
		Title:       req.Title,
		ShortName:   req.ShortName,
		Description: req.Description,
		SmsNoti:     *req.SmsNoti,
		EmailNoti:   *req.EmailNoti,
		CallNoti:    *req.CallNoti,
		Image:       req.Image,
		CreateUser:  10,
	}
	fmt.Println(args)
	product, err := s.store.CreateProduct(ctx, args)
	fmt.Println(product)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (s *Server) UpdateProduct(ctx *gin.Context) {
	var req CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	args := db.UpdateProductsParams{
		Title:       req.Title,
		ShortName:   req.ShortName,
		Description: req.Description,
		SmsNoti:     *req.SmsNoti,
		EmailNoti:   *req.EmailNoti,
		CallNoti:    *req.CallNoti,
		Image:       req.Image,
		ModifyUser:  10,
		ID:          req.Id,
		ModifiedAt:  time.Now().UTC(),
	}
	fmt.Println(args)
	product, err := s.store.UpdateProducts(ctx, args)
	fmt.Println(product)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, product)
}
