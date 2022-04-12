package repository

import (
	c "go-jwt/app/configs"
	m "go-jwt/app/models"
)

func InsertUser(username string, password string, email string) (*m.User, error) {
	user := m.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	if res := c.DB.Create(&user); res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func FindUserByUsername(username string) (*m.User, error) {
	var user m.User
	if res := c.DB.Where("username = ?", username).First(&user); res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func FindUserByID(id uint) (*m.User, error) {
	var user m.User
	if res := c.DB.First(&user, id); res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
