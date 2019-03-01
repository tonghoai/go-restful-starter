package handle

import (
	"errors"
	"restful-starter/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(u *models.User) (string, error) {
	var err error
	user := models.User{Username: u.Username, Password: u.Password}
	p, err := userService.GetUser(&user)
	if err != nil {
		return "", err
	}
	user = *p

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["permission"] = user.Permission
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return string(t), nil
}

func Register(u *models.User) (*models.User, error) {
	var err error
	user, err := userService.GetUser(u)
	if err != nil {
		return u, errors.New("Lỗi")
	}
	if user.ID != 0 {
		return u, errors.New("Người dùng đã tồn tại")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return u, err
	}
	u.Password = string(hash)
	res, err := userService.InsertUser(u)
	if err != nil {
		return u, err
	}
	u.Password = ""
	return res, nil
}
