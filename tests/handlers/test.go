package handlers

import (
	"github.com/limingxinleo/gorpc/tests/models"
	"github.com/limingxinleo/gorpc/tests"
)

type Test struct {
}

func (*Test) Version() string {
	return "1.0.0"
}

func (*Test) GetUserById(id float64) *models.User {
	db := tests.Provider.DB.GetInstance()
	user := &models.User{}
	db.First(&user)
	return user
}
