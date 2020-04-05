package user

import (
	"errors"
	"time"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/sqlstore"
	"github.com/hooone/evening/pkg/util"
)

func init() {
	bus.AddHandler("sql", GetUser)
	bus.AddHandler("sql", GetSignedInUser)
	bus.AddHandler("sql", CreateUser)
	bus.AddHandler("sql", UpdateUserLastSeenAt)

	bus.AddHandler("sql", GetToken)
	bus.AddHandler("sql", CreateToken)
	bus.AddHandler("sql", RemoveToken)

}

func (u *SignedInUser) ShouldUpdateLastSeenAt() bool {
	return u.Id > 0 && time.Since(u.LastSeenAt) > time.Minute*5
}

func GetUser(query *LoginUserQuery) error {
	user := User{
		Login: query.Username,
	}
	success, err := sqlstore.X.Where("login = ?", query.Username).Get(&user)
	if err != nil {
		return err
	}
	if !success {
		return ErrUserNotFound
	}
	query.User = &user
	return nil
}
func GetSignedInUser(query *GetSignedInUserQuery) error {
	user := SignedInUser{
		Id:    query.UserId,
		Login: query.Login,
	}
	if query.UserId > 0 {
		success, err := sqlstore.X.Where("id = ?", query.UserId).Get(&user)
		if err != nil {
			return err
		}
		if !success {
			return errors.New("user not found")
		}
		query.Result = &user
		return nil

	}
	exists, err2 := sqlstore.X.Where("login = ?", query.Login).Get(&user)
	if err2 != nil {
		return err2
	}
	if !exists {
		return errors.New("user not found")
	}
	query.Result = &user
	return nil
}

func CreateToken(cmd *CreateTokenCommand) error {
	token := cmd.Data
	_, err := sqlstore.X.Insert(&token)
	if err != nil {
		return err
	}
	cmd.Result = &token
	return nil
}

func RemoveToken(cmd *RemoveTokenCommand) error {
	token := UserToken{Id: cmd.TokenId}
	r, err := sqlstore.X.Id(cmd.TokenId).Delete(token)
	if err != nil {
		return err
	}
	cmd.Result = r
	return nil
}

func GetToken(query *GetTokenQuery) error {
	token := UserToken{AuthToken: query.HashedToken}
	success, err := sqlstore.X.Where("auth_token = ? OR prev_auth_token = ?", query.HashedToken, query.HashedToken).Get(&token)
	if err != nil {
		return err
	}
	if !success {
		return ErrUserTokenNotFound
	}
	query.Result = &token
	return nil
}
func UpdateUserLastSeenAt(cmd *UpdateUserLastSeenAtCommand) error {
	user := SignedInUser{
		Id:         cmd.UserId,
		LastSeenAt: time.Now(),
	}
	_, err := sqlstore.X.ID(cmd.UserId).Cols("last_seen_at").Update(&user)
	return err
}

func CreateUser(cmd *CreateUserCommand) error {
	if cmd.Email == "" {
		cmd.Email = cmd.Login
	}

	// create user
	user := User{
		Email:      cmd.Email,
		Name:       cmd.Name,
		Login:      cmd.Login,
		IsAdmin:    cmd.IsAdmin,
		Created:    time.Now(),
		Updated:    time.Now(),
		LastSeenAt: time.Now().AddDate(-10, 0, 0),
	}

	salt, err := util.GetRandomString(10)
	if err != nil {
		return err
	}
	user.Salt = salt
	rands, err := util.GetRandomString(10)
	if err != nil {
		return err
	}
	user.Rands = rands

	if len(cmd.Password) > 0 {
		encodedPassword, err := util.EncodePassword(cmd.Password, user.Salt)
		if err != nil {
			return err
		}
		user.Password = encodedPassword
	}

	_, err3 := sqlstore.X.Insert(&user)
	if err3 != nil {
		return err3
	}
	user.OrgId = user.Id
	_, err2 := sqlstore.X.ID(user.Id).Cols("org_id").Update(&user)
	if err2 != nil {
		return err2
	}
	cmd.Result = user

	return nil
}
