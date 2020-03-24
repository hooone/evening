package sqlstore

import (
	"errors"
	"fmt"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetUser)
	bus.AddHandler("sql", GetSignedInUser)
}

func GetUser(query *models.LoginUserQuery) error {
	user := models.User{
		Login: query.Username,
	}
	fmt.Println(query.Username)
	success, err := x.Where("login = ?", query.Username).Get(&user)
	if err != nil {
		return err
	}
	if !success {
		return models.ErrUserNotFound
	}
	query.User = &user
	return nil
}
func GetSignedInUser(query *models.GetSignedInUserQuery) error {
	user := models.SignedInUser{
		Id:    query.UserId,
		Login: query.Login,
	}
	if query.UserId > 0 {
		success, err := x.Where("id = ?", query.UserId).Get(&user)
		if err != nil {
			return err
		}
		if !success {
			return errors.New("user not found")
		}
		query.Result = &user
		return nil

	}
	exists, err2 := x.Where("login = ?", query.Login).Get(&user)
	if err2 != nil {
		return err2
	}
	if !exists {
		return errors.New("user not found")
	}
	query.Result = &user
	return nil
}
