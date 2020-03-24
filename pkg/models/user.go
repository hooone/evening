package models

import (
	"errors"
	"time"
)

var (
	ErrUserNotFound     = errors.New("User not found")
	ErrLastGrafanaAdmin = errors.New("Cannot remove last grafana admin")
)

type User struct {
	Id       int64
	Version  int
	Email    string
	Name     string
	Login    string
	Password string
	Salt     string
	Rands    string
	Theme    string

	IsAdmin bool
	OrgId   int64

	Created time.Time
	Updated time.Time
}

type SignedInUser struct {
	Id          int64
	OrgId       int64
	Login       string
	Name        string
	Email       string
	IsAnonymous bool    `xorm:"-"`
	Teams       []int64 `xorm:"-"`
}

func (l SignedInUser) TableName() string {
	return "user"
}

type GetSignedInUserQuery struct {
	UserId int64
	Login  string
	OrgId  int64
	Result *SignedInUser
}
