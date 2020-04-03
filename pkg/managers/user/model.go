package user

import (
	"errors"
	"time"
)

// Typed errors
var (
	ErrUserNotFound = errors.New("User not found")

	ErrUserTokenNotFound = errors.New("user token not found")

	ErrEmailNotAllowed       = errors.New("Required email domain not fulfilled")
	ErrInvalidCredentials    = errors.New("Invalid Username or Password")
	ErrNoEmail               = errors.New("Login provider didn't return an email address")
	ErrProviderDeniedRequest = errors.New("Login provider denied login request")
	ErrSignUpNotAllowed      = errors.New("Signup is not allowed for this adapter")
	ErrTooManyLoginAttempts  = errors.New("Too many consecutive incorrect login attempts for user. Login for user temporarily blocked")
	ErrPasswordEmpty         = errors.New("No password provided")
	ErrUserDisabled          = errors.New("User is disabled")
	ErrAbsoluteRedirectTo    = errors.New("Absolute urls are not allowed for redirect_to cookie value")
	ErrInvalidRedirectTo     = errors.New("Invalid redirect_to cookie value")
)

type User struct {
	Id       int64
	Version  int
	Email    string
	Name     string
	Login    string
	Password string
	Salt     string
	Rands    string
	Theme    string

	IsAdmin bool
	OrgId   int64

	Created time.Time
	Updated time.Time
}

type SignedInUser struct {
	Id          int64
	OrgId       int64
	Login       string
	Name        string
	Email       string
	IsAnonymous bool    `xorm:"-"`
	Teams       []int64 `xorm:"-"`
}

func (l SignedInUser) TableName() string {
	return "user"
}

// UserToken represents a user token
type UserToken struct {
	Id            int64
	UserId        int64
	AuthToken     string
	PrevAuthToken string
	UserAgent     string
	ClientIp      string
	AuthTokenSeen bool
	SeenAt        int64
	RotatedAt     int64
	CreatedAt     int64
	UpdatedAt     int64
	UnhashedToken string `xorm:"-"`
}
