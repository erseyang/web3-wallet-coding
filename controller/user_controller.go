package controller

import (
	"com.wallet/coding/domain/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(ctx *gin.Context)
}

type userController struct {
}

func NewUserController() UserController {
	return &userController{}
}

func (u *userController) TransferLog(c *gin.Context) {

}

// Login
// create user session
func (u *userController) Login(c *gin.Context) {
	fmt.Println("start...")
	_ = service.NewUserServiceImpl().Login(nil)
}

// Balance
// user check him/her wallet balance
func (u *userController) Balance(c *gin.Context) {

}
