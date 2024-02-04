package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	mockdb "github.com/czarro/db/mock"
	db "github.com/czarro/db/sqlc"
	"github.com/czarro/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.Password)
	if err != nil {
		return false
	}

	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUser(t)
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"firstName":   user.FirstName,
				"lastName":    user.LastName,
				"phone":       user.Phone,
				"countryCode": user.CountryCode,
				"status_id":   user.StatusID,
				"password":    password,
				"otp":         util.RandomString(4),
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserParams{
					RoleID:      user.RoleID,
					FirstName:   user.FirstName,
					LastName:    user.LastName,
					Phone:       user.Phone,
					CountryCode: user.CountryCode,
					StatusID:    user.StatusID,
					Password:    user.Password,
					CreateUser:  user.CreateUser,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomUser(t *testing.T) (user db.CzUser, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.CzUser{
		RoleID:      100,
		FirstName:   util.RandomString(10),
		MiddleName:  util.RandomString(5),
		LastName:    util.RandomString(4),
		CountryCode: int32(util.RandomInt(1, 10)),
		Phone:       fmt.Sprintf("%d", util.RandomInt64(600000000, 999999999)),
		Password:    hashedPassword,
		StatusID:    int32(util.RandomInt(1, 1)),
		// CreateUser:  util.RandomInt64(1, 1000000),
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.CzUser) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser userResponse
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	require.Equal(t, user.FirstName, gotUser.FirstName)
	require.Equal(t, user.LastName, gotUser.LastName)
	require.Equal(t, user.CountryCode, gotUser.CountryCode)
	require.Equal(t, user.Phone, gotUser.Phone)
}
