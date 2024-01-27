package api

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/czarro/db/sqlc"
	"github.com/czarro/logger"
	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Title        string  `json:"title" binding:"required"`
	ShortName    string  `json:"shortName" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	SmsNoti      *bool   `json:"smsNoti" binding:"required"`
	EmailNoti    *bool   `json:"emailNoti" binding:"required"`
	CallNoti     *bool   `json:"callNoti" binding:"required"`
	Image        string  `json:"image" binding:"required"`
	CurrencyId   int32   `json:"currencyId" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	CurencyValid string  `json:"curencyValid" binding:"required,customCurrency"`
	Id           int32   `json:"id,omitempty"`
}

// CreateProduct create product
func (s *Server) CreateProduct(ctx *gin.Context) {
	var req CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		CurrencyID:  req.CurrencyId,
		Price:       req.Price,
		CreateUser:  10,
	}
	logger.Info(fmt.Sprintf("arg %+v", args))
	product, err := s.store.CreateProduct(ctx, args)
	if err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	logger.Debug(fmt.Sprintf("products %+v", product))
	ctx.JSON(http.StatusOK, product)
}

// UpdateProduct update the product
func (s *Server) UpdateProduct(ctx *gin.Context) {
	var req CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		CurrencyID:  req.CurrencyId,
		Price:       req.Price,
		ModifyUser:  10,
		ID:          req.Id,
		ModifiedAt:  time.Now().UTC(),
	}
	logger.Info(fmt.Sprintf("arg %+v", args))
	product, err := s.store.UpdateProducts(ctx, args)
	fmt.Println(product)
	if err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	logger.Debug(fmt.Sprintf("products %+v", product))
	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct delete product by id
func (s *Server) DeleteProduct(ctx *gin.Context) {
	var req CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		CurrencyID:  req.CurrencyId,
		Price:       req.Price,
		ModifyUser:  10,
		ID:          req.Id,
		ModifiedAt:  time.Now().UTC(),
	}
	logger.Info(fmt.Sprintf("arg %+v", args))
	product, err := s.store.UpdateProducts(ctx, args)
	fmt.Println(product)
	if err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	logger.Debug(fmt.Sprintf("products %+v", product))
	ctx.JSON(http.StatusOK, product)
}

type listProductRequest struct {
	PageId   int32 `form:"pageId" binding:"required,min=1"`
	PageSize int32 `form:"pageSize" binding:"required,min=2,max=10"`
}

// GetListProduct get product list
func (s *Server) GetListProduct(ctx *gin.Context) {
	var req listProductRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := db.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	products, err := s.store.ListProducts(ctx, args)
	if err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	logger.Debug(fmt.Sprintf("products %+v", products))
	ctx.JSON(http.StatusOK, products)
}

type GetProductRequest struct {
	Id int32 `uri:"id" binding:"required"`
}

// GetProduct get product by id
func (s *Server) GetProduct(ctx *gin.Context) {
	var req GetProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	logger.Info(fmt.Sprintf("arg %+v", req))
	product, err := s.store.GetProduct(ctx, req.Id)
	if err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	logger.Debug(fmt.Sprintf("product %+v", product))
	ctx.JSON(http.StatusOK, product)
}
