package sqlstore

import (
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetToken)
	bus.AddHandler("sql", CreateToken)
	bus.AddHandler("sql", RemoveToken)
}

func CreateToken(cmd *models.CreateTokenCommand) error {
	token := cmd.Data
	_, err := x.Insert(&token)
	if err != nil {
		return err
	}
	cmd.Result = &token
	return nil
}

func RemoveToken(cmd *models.RemoveTokenCommand) error {
	token := models.UserToken{Id: cmd.TokenId}
	r, err := x.Id(cmd.TokenId).Delete(token)
	if err != nil {
		return err
	}
	cmd.Result = r
	return nil
}

func GetToken(query *models.GetTokenQuery) error {
	token := models.UserToken{AuthToken: query.HashedToken}
	success, err := x.Where("auth_token = ? OR prev_auth_token = ?", query.HashedToken, query.HashedToken).Get(&token)
	if err != nil {
		return err
	}
	if !success {
		return models.ErrUserTokenNotFound
	}
	query.Result = &token
	return nil
}
