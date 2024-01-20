package db

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/czarro/util"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) CzProduct {
	arg := CreateProductParams{
		Title:       util.RandomString(50),
		ShortName:   util.RandomString(30),
		Description: util.RandomString(10),
		SmsNoti:     true,
		EmailNoti:   false,
		CallNoti:    true,
		Image:       util.RandomString(100),
		CurrencyID:  int32(1),
		Price:       float64(util.RandomInt(1, 100)),
		StatusID:    1,
		CreateUser:  util.RandomInt64(1, 4),
	}
	product, err := testStore.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.Equal(t, arg.Title, product.Title)
	require.Equal(t, arg.ShortName, product.ShortName)
	require.Equal(t, arg.Description, product.Description)
	require.Equal(t, arg.SmsNoti, product.SmsNoti)
	require.Equal(t, arg.EmailNoti, product.EmailNoti)
	require.Equal(t, arg.CallNoti, product.CallNoti)
	require.Equal(t, arg.Image, product.Image)
	require.Equal(t, arg.CurrencyID, product.CurrencyID)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.StatusID, product.StatusID)
	require.Equal(t, arg.CreateUser, product.CreateUser)

	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)
	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	product := createRandomProduct(t)
	sameProduct, err := testStore.GetProduct(context.Background(), product.ID)
	require.NoError(t, err)
	require.Equal(t, product, sameProduct)
}

func TestGetListProducts(t *testing.T) {
	arg := ListProductsParams{
		Limit:  5,
		Offset: 1,
	}
	users, err := testStore.ListProducts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)
}

func TestUpdateProduct(t *testing.T) {
	p1 := createRandomProduct(t)
	arg := UpdateProductsParams{
		Title:       p1.Title,
		ShortName:   p1.ShortName,
		Description: p1.Description,
		SmsNoti:     p1.SmsNoti,
		EmailNoti:   p1.EmailNoti,
		CallNoti:    p1.CallNoti,
		Image:       util.RandomString(100),
		CurrencyID:  1,
		Price:       p1.Price,
		StatusID:    1,
		ModifyUser:  util.RandomInt64(1, 4),
		ModifiedAt:  time.Now().UTC(),
		ID:          p1.ID,
	}
	p2, err := testStore.UpdateProducts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEqual(t, p1, p2)
	require.NotEqual(t, p1.Image, p2.Image)
	require.NotEqual(t, p1.ModifiedAt, p2.ModifiedAt)
}

func TestUpdateProductUsingForUpdate(t *testing.T) {
	var err error
	p1 := createRandomProduct(t)
	p2, err := testStore.GetProductForUpdate(context.Background(), p1.ID)
	if err != nil {
		log.Println(err.Error())
	}
	arg := UpdateProductsParams{
		Title:       p1.Title,
		ShortName:   p1.ShortName,
		Description: p1.Description,
		SmsNoti:     p1.SmsNoti,
		EmailNoti:   p1.EmailNoti,
		CallNoti:    p1.CallNoti,
		Image:       util.RandomString(100),
		CurrencyID:  1,
		Price:       p1.Price,
		StatusID:    1,
		ModifyUser:  util.RandomInt64(1, 4),
		ModifiedAt:  time.Now().UTC(),
		ID:          p1.ID,
	}
	udatedP1, err := testStore.UpdateProducts(context.Background(), arg)
	if err != nil {
		log.Println(err)
	}
	require.NoError(t, err)
	require.NotEqual(t, p1, udatedP1)
	require.NotEqual(t, p1.Image, p2.Image)
	require.NotEqual(t, p1.ModifiedAt, p2.ModifiedAt)
}
