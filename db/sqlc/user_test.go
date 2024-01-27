package db

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/czarro/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) CzUser {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		RoleID:      100,
		FirstName:   util.RandomString(10),
		MiddleName:  util.RandomString(5),
		LastName:    util.RandomString(4),
		Dob:         time.Date(util.RandomInt(1900, 2024), time.February, util.RandomInt(1, 31), 0, 0, 0, 0, time.Local),
		CountryCode: int32(util.RandomInt(1, 10)),
		Phone:       fmt.Sprintf("%d", util.RandomInt64(600000000, 999999999)),
		Email:       util.RandomEmail(),
		Password:    hashedPassword,
		StatusID:    int32(util.RandomInt(1, 3)),
		CreateUser:  util.RandomInt64(1, 1000000),
	}
	User, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, User)
	require.Equal(t, arg.RoleID, User.RoleID)
	require.Equal(t, arg.FirstName, User.FirstName)
	require.Equal(t, arg.LastName, User.LastName)
	require.Equal(t, arg.Dob, User.Dob)
	require.Equal(t, arg.Phone, User.Phone)
	require.Equal(t, arg.Email, User.Email)
	require.Equal(t, arg.Password, User.Password)
	require.Equal(t, arg.StatusID, User.StatusID)

	require.NotZero(t, User.ID)
	require.NotZero(t, User.CreatedAt)
	return User
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	User := createRandomUser(t)
	sameUser, err := testStore.GetUser(context.Background(), User.ID)
	require.NoError(t, err)
	require.Equal(t, User, sameUser)
}

func TestGetListUser(t *testing.T) {
	arg := ListusersParams{
		Limit:  5,
		Offset: 1,
	}
	users, err := testStore.Listusers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)
}

func TestUpdateUser(t *testing.T) {
	User1 := createRandomUser(t)
	arg := UpdateUserParams{
		Password: util.RandomString(10),
		ID:       User1.ID,
	}
	User, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEqual(t, User1, User)
	require.NotEqual(t, User1.Password, User.Password)
}

func TestUpdateUserUsingForUpdate(t *testing.T) {
	var err error
	User := createRandomUser(t)
	arg := UpdateUserParams{
		Password: util.RandomString(10),
		ID:       User.ID,
	}
	User, err = testStore.GetUserForUpdate(context.Background(), User.ID)
	if err != nil {
		log.Println(err.Error())
	}
	udatedUser, err := testStore.UpdateUser(context.Background(), arg)
	if err != nil {
		log.Println(err)
	}
	require.NoError(t, err)
	require.NotEqual(t, User, udatedUser)
	require.NotEqual(t, User.Password, udatedUser.Password)
}
