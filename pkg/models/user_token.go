package models

import (
	"errors"
)

// Typed errors
var (
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
type LoginUserQuery struct {
	ReqContext *ReqContext
	Username   string
	Password   string
	User       *User
	IpAddress  string
}

type GetTokenQuery struct {
	HashedToken string
	Result      *UserToken
}
type CreateTokenCommand struct {
	Data   UserToken
	Result *UserToken
}
type RemoveTokenCommand struct {
	TokenId int64
	Result  int64
}

// type RevokeAuthTokenCmd struct {
// 	AuthTokenId int64 `json:"authTokenId"`
// }

// // UserTokenService are used for generating and validating user tokens
// type UserTokenService interface {
// 	LookupToken(ctx context.Context, unhashedToken string) (*UserToken, error)
// }
