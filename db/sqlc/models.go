// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type CzCountry struct {
	ID        int32  `json:"id"`
	Iso       string `json:"iso"`
	Name      string `json:"name"`
	Nicename  string `json:"nicename"`
	Iso3      string `json:"iso3"`
	Numcode   string `json:"numcode"`
	PhoneCode int32  `json:"phone_code"`
}

type CzCurrency struct {
	ID          int32  `json:"id"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Visible     bool   `json:"visible"`
}

type CzProduct struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	ShortName   string    `json:"short_name"`
	Description string    `json:"description"`
	SmsNoti     bool      `json:"sms_noti"`
	EmailNoti   bool      `json:"email_noti"`
	CallNoti    bool      `json:"call_noti"`
	Image       string    `json:"image"`
	CurrencyID  int32     `json:"currency_id"`
	Price       float64   `json:"price"`
	StatusID    int32     `json:"status_id"`
	CreateUser  int64     `json:"create_user"`
	ModifyUser  int64     `json:"modify_user"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

type CzRole struct {
	ID         int32     `json:"id"`
	RoleName   string    `json:"role_name"`
	StatusID   int32     `json:"status_id"`
	CreateUser int64     `json:"create_user"`
	ModifyUser int64     `json:"modify_user"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Visible    bool      `json:"visible"`
}

type CzStatus struct {
	ID          int32  `json:"id"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Visible     bool   `json:"visible"`
}

type CzUser struct {
	ID                int64     `json:"id"`
	UniqueID          uuid.UUID `json:"unique_id"`
	RoleID            int32     `json:"role_id"`
	FirstName         string    `json:"first_name"`
	MiddleName        string    `json:"middle_name"`
	LastName          string    `json:"last_name"`
	Dob               time.Time `json:"dob"`
	CountryCode       int32     `json:"country_code"`
	Phone             string    `json:"phone"`
	Email             string    `json:"email"`
	Salt              string    `json:"salt"`
	Password          string    `json:"password"`
	PasswordModifedAt time.Time `json:"password_modifed_at"`
	StatusID          int32     `json:"status_id"`
	CreateUser        int64     `json:"create_user"`
	ModifyUser        int64     `json:"modify_user"`
	CreatedAt         time.Time `json:"created_at"`
	ModifiedAt        time.Time `json:"modified_at"`
}

type CzUsersAddress struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	CountryCode int32     `json:"country_code"`
	Address1    string    `json:"address1"`
	Address2    string    `json:"address2"`
	Address3    string    `json:"address3"`
	Address4    string    `json:"address4"`
	Location    string    `json:"location"`
	StatusID    int32     `json:"status_id"`
	CreateUser  int64     `json:"create_user"`
	ModifyUser  int64     `json:"modify_user"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}
