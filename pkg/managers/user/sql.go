package user

import (
	"errors"
	"fmt"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/sqlstore"
)

func init() {
	bus.AddHandler("sql", GetUser)
	bus.AddHandler("sql", GetSignedInUser)

	bus.AddHandler("sql", GetToken)
	bus.AddHandler("sql", CreateToken)
	bus.AddHandler("sql", RemoveToken)
}

func GetUser(query *LoginUserQuery) error {
	user := User{
		Login: query.Username,
	}
	fmt.Println(query.Username)
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
