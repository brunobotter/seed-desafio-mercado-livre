package controller

import (
	"net/http"

	"github.com/brunobotter/mercado-livre/internal/domain/contract"
	"github.com/brunobotter/mercado-livre/internal/request"
	"github.com/brunobotter/mercado-livre/internal/util"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	svc contract.ServiceManager
}

func NewCategoryController(svc contract.ServiceManager) *CategoryController {
	return &CategoryController{
		svc: svc,
	}
}

func (c *CategoryController) SaveNewCategory(ctx *gin.Context) {
	var registerRequest request.SaveCategoryRequest
	if err := ctx.Bind(&registerRequest); err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Invalid Body")
		return
	}
	response, err := c.svc.InternalService().CategoryService().SaveCategory(ctx, registerRequest)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Error to register category")
		return
	}
	util.ResponderApiOk(ctx, response)
}
