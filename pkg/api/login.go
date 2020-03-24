package api

import (
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/middleware"
	"github.com/hooone/evening/pkg/models"
	"github.com/hooone/evening/pkg/setting"
	"github.com/hooone/evening/pkg/util"
)

const (
	ViewIndex            = "index"
	LoginErrorCookieName = "login_error"
)

var getViewIndex = func() string {
	return ViewIndex
}

func (hs *HTTPServer) validateRedirectTo(redirectTo string) error {
	to, err := url.Parse(redirectTo)
	if err != nil {
		return errors.New("Invalid redirect_to cookie value")
	}
	if to.IsAbs() {
		return errors.New("Absolute urls are not allowed for redirect_to cookie value")
	}
	// when using a subUrl, the redirect_to should have a relative or absolute path that includes the subUrl, otherwise the redirect
	// will send the user to the wrong location
	if hs.Cfg.AppSubURL != "" && !strings.HasPrefix(to.Path, hs.Cfg.AppSubURL) && !strings.HasPrefix(to.Path, "/"+hs.Cfg.AppSubURL) {
		return errors.New("Invalid redirect_to cookie value")
	}
	return nil
}

func (hs *HTTPServer) cookieOptionsFromCfg() middleware.CookieOptions {
	return middleware.CookieOptions{
		Path:             hs.Cfg.AppSubURL + "/",
		Secure:           hs.Cfg.CookieSecure,
		SameSiteDisabled: hs.Cfg.CookieSameSiteDisabled,
		SameSiteMode:     hs.Cfg.CookieSameSiteMode,
	}
}

func (hs *HTTPServer) LoginView(c *models.ReqContext) {
	if _, ok := tryGetEncryptedCookie(c, LoginErrorCookieName); ok {
		middleware.DeleteCookie(c.Resp, LoginErrorCookieName, hs.cookieOptionsFromCfg)
		c.HTML(200, getViewIndex())
		return
	}

	if c.IsSignedIn {
		if redirectTo, _ := url.QueryUnescape(c.GetCookie("redirect_to")); len(redirectTo) > 0 {
			if err := hs.validateRedirectTo(redirectTo); err != nil {
				c.HTML(200, getViewIndex())
				middleware.DeleteCookie(c.Resp, "redirect_to", hs.cookieOptionsFromCfg)
				return
			}
			middleware.DeleteCookie(c.Resp, "redirect_to", hs.cookieOptionsFromCfg)
			c.Redirect(redirectTo)
			return
		}

		c.Redirect(setting.AppSubURL + "/")
		return
	}
	c.HTML(200, getViewIndex())
}
func tryGetEncryptedCookie(ctx *models.ReqContext, cookieName string) (string, bool) {
	cookie := ctx.GetCookie(cookieName)
	if cookie == "" {
		return "", false
	}

	decoded, err := hex.DecodeString(cookie)
	if err != nil {
		return "", false
	}

	decryptedError, err := util.Decrypt(decoded, setting.SecretKey)
	return string(decryptedError), err == nil
}

var validatePassword = func(providedPassword string, userPassword string, userSalt string) error {
	passwordHashed, err := util.EncodePassword(providedPassword, userSalt)
	if err != nil {
		return err
	}
	if subtle.ConstantTimeCompare([]byte(passwordHashed), []byte(userPassword)) != 1 {
		return models.ErrInvalidCredentials
	}

	return nil
}

func (hs *HTTPServer) LoginPost(c *models.ReqContext, cmd dtos.LoginCommand) Response {
	userQuery := &models.LoginUserQuery{
		ReqContext: c,
		Username:   cmd.User,
		Password:   cmd.Password,
		IpAddress:  c.Req.RemoteAddr,
	}

	if err := bus.Dispatch(userQuery); err != nil {
		if err == models.ErrUserNotFound {
			return Error(401, "Invalid username or password", err)
		}
		return Error(500, "Error while trying to authenticate user", err)
	}

	user := userQuery.User

	if err := validatePassword(userQuery.Password, user.Password, user.Salt); err != nil {
		return Error(401, "Invalid username or password", err)
	}

	hs.loginUserWithUser(user, c)

	result := CommonResult{
		Success: true,
		Data:    user.Name,
	}

	return JSON(200, result)
}

func (hs *HTTPServer) Logout(c *models.ReqContext) {
	fmt.Println("logout")
	if err := hs.AuthTokenService.RevokeToken(c.Req.Context(), c.UserToken); err != nil && err != models.ErrUserTokenNotFound {
		hs.log.Error("failed to revoke auth token", "error", err)
	}

	fmt.Println("removetoken")
	middleware.WriteSessionCookie(c, "", -1)

	fmt.Println("writesession")
	if setting.SignoutRedirectUrl != "" {
		c.Redirect(setting.SignoutRedirectUrl)
	} else {
		hs.log.Info("Successful Logout", "User", c.Email)
		fmt.Println("Redirect")
		c.Redirect(setting.AppSubURL + "/login")
	}
}

func (hs *HTTPServer) loginUserWithUser(user *models.User, c *models.ReqContext) {
	if user == nil {
		hs.log.Error("user login with nil user")
	}

	userToken, err := hs.AuthTokenService.CreateToken(c.Req.Context(), user.Id, c.RemoteAddr(), c.Req.UserAgent())
	if err != nil {
		hs.log.Error("failed to create auth token", "error", err)
	}
	hs.log.Info("Successful Login", "User", user.Login)
	middleware.WriteSessionCookie(c, userToken.UnhashedToken, hs.Cfg.LoginMaxLifetimeDays)
}
