package db

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/czarro/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomCustomer(t *testing.T) Customer {
	arg := CreateCustomerParams{
		RoleID:     pgtype.Int4{Int32: int32(100), Valid: true},
		FirstName:  util.RandomString(10),
		MiddleName: util.RandomString(5),
		LastName:   util.RandomString(4),
		Dob: pgtype.Date{
			Time: time.Date(int(util.RandomInt(1900, 2023)),
				time.Month(util.RandomInt(1, 12)),
				int(util.RandomInt(1, 28)),
				0, 0, 0, 0, time.UTC,
			),
			Valid: true,
		},
		CountryCode: "+91",
		Phone:       fmt.Sprintf("%d", util.RandomInt(600000000, 999999999)),
		Email:       pgtype.Text{String: util.RandomEmail(), Valid: true},
		Password:    pgtype.Text{String: util.RandomString(8), Valid: true},
		StatusID:    pgtype.Int4{Int32: 1, Valid: true},
		CreateUser:  int32(12),
	}
	customer, err := testStore.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	require.Equal(t, arg.RoleID, customer.RoleID)
	require.Equal(t, arg.FirstName, customer.FirstName)
	require.Equal(t, arg.LastName, customer.LastName)
	require.Equal(t, arg.Dob, customer.Dob)
	require.Equal(t, arg.Phone, customer.Phone)
	require.Equal(t, arg.Email, customer.Email)
	require.Equal(t, arg.Password, customer.Password)
	require.Equal(t, arg.StatusID, customer.StatusID)

	require.NotZero(t, customer.ID)
	require.NotZero(t, customer.CreatedAt)
	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomer(t *testing.T) {
	customer := createRandomCustomer(t)
	sameCustomer, err := testStore.GetCustomer(context.Background(), customer.ID)
	require.NoError(t, err)
	require.Equal(t, customer, sameCustomer)
}

func TestGetListCustomer(t *testing.T) {
	arg := ListCustomersParams{
		Limit:  5,
		Offset: 1,
	}
	customers, err := testStore.ListCustomers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customers)
}

func TestUpdateCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	arg := UpdateCustomerParams{
		Password: pgtype.Text{String: util.RandomString(9), Valid: true},
		ID:       customer1.ID,
	}
	customer, err := testStore.UpdateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEqual(t, customer1, customer)
	require.NotEqual(t, customer1.Password, customer.Password)
}

func TestUpdateCustomerUsingForUpdate(t *testing.T) {
	var err error
	customer := createRandomCustomer(t)
	arg := UpdateCustomerParams{
		Password: pgtype.Text{String: util.RandomString(9), Valid: true},
		ID:       customer.ID,
	}
	customer, err = testStore.GetCustomerForUpdate(context.Background(), customer.ID)
	if err != nil {
		log.Println(err.Error())
	}
	udatedCustomer, err := testStore.UpdateCustomer(context.Background(), arg)
	if err != nil {
		log.Println(err)
	}
	require.NoError(t, err)
	require.NotEqual(t, customer, udatedCustomer)
	require.NotEqual(t, customer.Password, udatedCustomer.Password)
}
