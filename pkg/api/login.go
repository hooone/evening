package api

import (
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/api/middleware"
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/user"
	"github.com/hooone/evening/pkg/services/navigation"
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

func (hs *HTTPServer) LoginView(c *dtos.ReqContext) {
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
func tryGetEncryptedCookie(ctx *dtos.ReqContext, cookieName string) (string, bool) {
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
		return user.ErrInvalidCredentials
	}

	return nil
}

func (hs *HTTPServer) LoginPost(c *dtos.ReqContext, cmd dtos.LoginCommand) Response {
	userQuery := &user.LoginUserQuery{
		Username:  cmd.User,
		Password:  cmd.Password,
		IpAddress: c.Req.RemoteAddr,
	}

	if err := bus.Dispatch(userQuery); err != nil {
		if err == user.ErrUserNotFound {
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

func (hs *HTTPServer) Logout(c *dtos.ReqContext) {
	if err := hs.AuthTokenService.RevokeToken(c.Req.Context(), c.UserToken); err != nil && err != user.ErrUserTokenNotFound {
		hs.log.Error("failed to revoke auth token", "error", err)
	}

	middleware.WriteSessionCookie(c, "", -1)

	if setting.SignoutRedirectUrl != "" {
		c.Redirect(setting.SignoutRedirectUrl)
	} else {
		hs.log.Info("Successful Logout", "User", c.Email)
		c.Redirect(setting.AppSubURL + "/login")
	}
}

func (hs *HTTPServer) loginUserWithUser(user *user.User, c *dtos.ReqContext) {
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

func (hs *HTTPServer) SignUpStep2(c *dtos.ReqContext, form dtos.SignUpStep2Form) Response {
	createUserCmd := user.CreateUserCommand{
		Email:    form.Username,
		Login:    form.Username,
		Name:     form.Username,
		Password: form.Password,
	}
	result := new(CommonResult)

	if len(form.Password) < 4 {
		result.Data = 1
		result.Message = "Password is missing or too short"
		result.Success = false
		return JSON(200, result)
	}

	// check if user exists
	existing := user.GetSignedInUserQuery{Login: form.Username}
	if err := bus.Dispatch(&existing); err == nil {
		result.Data = 1
		result.Message = "User with same username address already exists"
		result.Success = false
		return JSON(200, result)
	}

	// dispatch create command
	if err := bus.Dispatch(&createUserCmd); err != nil {
		result.Data = 2
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}

	if err := hs.NavigationService.CreateTreePage(navigation.Page{
		Name: "home",
		Text: "主页",
	}, createUserCmd.Result.OrgId, "zh-CN"); err != nil {
		result.Data = 3
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}

	user := &createUserCmd.Result
	hs.loginUserWithUser(user, c)

	result.Data = 0
	result.Success = true
	return JSON(200, result)
}
