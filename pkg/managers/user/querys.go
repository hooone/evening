package user

type LoginUserQuery struct {
	Username  string
	Password  string
	User      *User
	IpAddress string
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

type GetSignedInUserQuery struct {
	UserId int64
	Login  string
	OrgId  int64
	Result *SignedInUser
}
type UpdateUserLastSeenAtCommand struct {
	UserId int64
}
type CreateUserCommand struct {
	Email          string
	Login          string
	Name           string
	Company        string
	Password       string
	EmailVerified  bool
	IsAdmin        bool
	IsDisabled     bool
	SkipOrgSetup   bool
	DefaultOrgRole string

	Result User
}
