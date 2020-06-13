package controllers

import (
	"app/domain"
	"app/interfaces/database"
	"app/usecase"
	"strconv"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.)  {

}
