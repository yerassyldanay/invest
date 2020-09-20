package intest

import (
	"invest/model"
	"testing"
)

func TestGetUser(t *testing.T) {
	var user = model.User{
		Id: 1,
	}
	err := user.GetByIdPreloaded(model.GetDB())
	if err != nil {
		return
	}
}