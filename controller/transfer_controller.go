package controller

import "github.com/gin-gonic/gin"

type TransferController struct{}

func NewTransferController() *TransferController {
	return &TransferController{}
}

// Send
// Send amount from user to another user
func (t *TransferController) Send(ctx *gin.Context) {

}

// Deposit
// save money to user's account
func (t *TransferController) Deposit(ctx *gin.Context) {}

// Withdraw
// withdraw money from user's account
func (t *TransferController) Withdraw(ctx *gin.Context) {

}
