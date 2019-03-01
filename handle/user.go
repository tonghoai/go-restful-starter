package handle

import (
	"restful-starter/models"
)

func GetUser(u *models.User) (*models.User, error) {
	var err error
	p, err := userService.GetUser(u)

	if err != nil {
		return u, err
	}

	return p, nil
}
