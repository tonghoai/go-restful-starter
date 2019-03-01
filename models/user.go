package models

import (
	"restful-starter/db"
	shared "restful-starter/shared/model"
)

type User shared.User
type UserService shared.UserService

func (us *UserService) GetUser(u *User) (*User, error) {
	var user User
	var query = db.GetMysql().Conn
	if u.ID != 0 {
		query = query.Where("id = ?", u.ID)
	}
	if len(u.Username) != 0 {
		query = query.Where("username = ?", u.Username)
	}
	query.First(&user)
	return &user, nil
}

func (us *UserService) InsertUser(u *User) (*User, error) {
	db.GetMysql().Conn.Create(u)
	return u, nil
}
