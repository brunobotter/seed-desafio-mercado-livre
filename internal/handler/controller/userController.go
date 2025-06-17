package controller

import (
	"net/http"

	"github.com/brunobotter/mercado-livre/internal/domain/contract"
	"github.com/brunobotter/mercado-livre/internal/request"
	"github.com/brunobotter/mercado-livre/internal/util"
	"github.com/gin-gonic/gin"
)

type UserControleler struct {
	svc contract.ServiceManager
}

func NewUserController(svc contract.ServiceManager) *UserControleler {
	return &UserControleler{
		svc: svc,
	}
}

func (c *UserControleler) RegisterNewUser(ctx *gin.Context) {
	var registerRequest request.RegisterNewUserRequest
	if err := ctx.Bind(&registerRequest); err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Invalid Body")
	}
	err := c.svc.InternalService().UserService().Register(ctx, registerRequest)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Error to register user")
	}
	util.ResponderApiOk(ctx, "User created with success!")
}
