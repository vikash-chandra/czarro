package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateCustomer(t *testing.T) {
	arg := CreateCustomerParams{
		RoleID:    1,
		FirstName: "vikash",
		LastName:  "jha",
		Dob: pgtype.Date{
			Time:  time.Date(1993, 9, 25, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
		Mobile:   "+91 886700000",
		Email:    "vikash@gmail.com",
		Password: "vikash123",
		Status:   "Active",
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
}
