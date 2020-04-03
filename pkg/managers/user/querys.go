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
