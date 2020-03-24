package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/infra/log"
	"github.com/hooone/evening/pkg/models"
	"github.com/hooone/evening/pkg/registry"
	"github.com/hooone/evening/pkg/setting"
	"github.com/hooone/evening/pkg/util"
)

func init() {
	registry.RegisterService(&UserTokenService{})
}

var getTime = time.Now

const urgentRotateTime = 1 * time.Minute

type UserTokenService struct {
	Cfg *setting.Cfg `inject:""`
	log log.Logger
}

func (s *UserTokenService) Init() error {
	s.log = log.New("auth")
	return nil
}

//CreateToken create token and save
func (s *UserTokenService) CreateToken(ctx context.Context, userId int64, clientAddr, userAgent string) (*models.UserToken, error) {
	clientIP, err := util.ParseIPAddress(clientAddr)
	if err != nil {
		s.log.Debug("Failed to parse client IP address", "clientAddr", clientAddr, "err", err)
		clientIP = ""
	}
	token, err := util.RandomHex(16)
	if err != nil {
		return nil, err
	}

	hashedToken := hashToken(token)

	now := getTime().Unix()

	userAuthToken := models.UserToken{
		UserId:        userId,
		AuthToken:     hashedToken,
		PrevAuthToken: hashedToken,
		ClientIp:      clientIP,
		UserAgent:     userAgent,
		RotatedAt:     now,
		CreatedAt:     now,
		UpdatedAt:     now,
		SeenAt:        0,
		AuthTokenSeen: false,
	}
	cmd := models.CreateTokenCommand{
		Data: userAuthToken,
	}

	if err := bus.Dispatch(&cmd); err != nil {
		return nil, err
	}

	cmd.Result.UnhashedToken = token

	s.log.Debug("user auth token created", "tokenId", userAuthToken.Id, "userId", userAuthToken.UserId, "clientIP", userAuthToken.ClientIp, "userAgent", userAuthToken.UserAgent, "authToken", userAuthToken.AuthToken)

	return cmd.Result, err
}

func (s *UserTokenService) LookupToken(ctx context.Context, unhashedToken string) (*models.UserToken, error) {
	hashedToken := hashToken(unhashedToken)
	if setting.Env == setting.DEV {
		s.log.Debug("looking up token", "unhashed", unhashedToken, "hashed", hashedToken)
	}
	query := models.GetTokenQuery{
		HashedToken: hashedToken,
	}
	if err := bus.Dispatch(&query); err != nil {
		return nil, err
	}
	return query.Result, nil
}

func hashToken(token string) string {
	hashBytes := sha256.Sum256([]byte(token + setting.SecretKey))
	return hex.EncodeToString(hashBytes[:])
}

func (s *UserTokenService) RevokeToken(ctx context.Context, token *models.UserToken) error {
	if token == nil {
		return models.ErrUserTokenNotFound
	}

	cmd := models.RemoveTokenCommand{
		TokenId: token.Id,
	}

	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	rowsAffected := cmd.Result

	if rowsAffected == 0 {
		s.log.Debug("user auth token not found/revoked", "tokenId", token.Id, "userId", token.UserId, "clientIP", token.ClientIp, "userAgent", token.UserAgent)
		return models.ErrUserTokenNotFound
	}

	s.log.Debug("user auth token revoked", "tokenId", token.Id, "userId", token.UserId, "clientIP", token.ClientIp, "userAgent", token.UserAgent)

	return nil
}
