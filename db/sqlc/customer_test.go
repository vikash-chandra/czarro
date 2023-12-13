package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/czarro/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomCustomer(t *testing.T) Customer {
	arg := CreateCustomerParams{
		RoleID:    int32(util.RandomInt(1, 10)),
		FirstName: util.RandomString(10),
		LastName:  util.RandomString(4),
		Dob: pgtype.Date{
			Time: time.Date(int(util.RandomInt(1900, 2023)),
				time.Month(util.RandomInt(1, 12)),
				int(util.RandomInt(1, 28)),
				0, 0, 0, 0, time.UTC,
			),
			Valid: true,
		},
		Mobile:     fmt.Sprintf("+91 %d", util.RandomInt(600000000, 999999999)),
		Email:      util.RandomEmail(),
		Password:   util.RandomString(8),
		Status:     "Active",
		CreateUser: pgtype.Int4{Int32: 2, Valid: true},
	}
	customer, err := testQueries.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	require.Equal(t, arg.RoleID, customer.RoleID)
	require.Equal(t, arg.FirstName, customer.FirstName)
	require.Equal(t, arg.LastName, customer.LastName)
	require.Equal(t, arg.Dob, customer.Dob)
	require.Equal(t, arg.Mobile, customer.Mobile)
	require.Equal(t, arg.Email, customer.Email)
	require.Equal(t, arg.Password, customer.Password)
	require.Equal(t, arg.Status, customer.Status)

	require.NotZero(t, customer.ID)
	require.NotZero(t, customer.CreatedAt)
	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomer(t *testing.T) {
	customer := createRandomCustomer(t)
	sameCustomer, err := testQueries.GetCustomer(context.Background(), customer.ID)
	require.NoError(t, err)
	require.Equal(t, customer, sameCustomer)
}

func TestGetListCustomer(t *testing.T) {
	arg := ListCustomersParams{
		Limit:  5,
		Offset: 1,
	}
	customers, err := testQueries.ListCustomers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customers)
}

func TestUpdateCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	arg := UpdateCustomerParams{
		Password: util.RandomString(9),
		ID:       customer1.ID,
	}
	customer, err := testQueries.UpdateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEqual(t, customer1, customer)
	require.NotEqual(t, customer1.Password, customer.Password)
}
