package controllers

import (
	"errors"
	"picbed/models"
	"picbed/services"
	"picbed/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MgtUserController struct{}

func (muc *MgtUserController) GetMgtUserById(ctx *gin.Context) {
	inputModel := &models.MgtUser{}
	if err := ctx.ShouldBindJSON(inputModel); err != nil {
		var vErr validator.ValidationErrors
		ok := errors.As(err, &vErr)
		if !ok {
			ResponseWithFail(ctx, CodeInvalidParameters)
		}
		translate := vErr.Translate(utils.Trans)
		ResponseWithFailMsg(ctx, CodeInvalidParameters, utils.RemoveTopStruct(translate))
		return
	}
	dbUser, _ := services.MgtUserServiceIns.GetUserByUsername(inputModel.Username)
	if dbUser != nil {
		ResponseWithFail(ctx, CodeUserExists)
		return
	}
	err := services.MgtUserServiceIns.Create(inputModel)
	if err != nil {
		ResponseWithFailMsg(ctx, CodeUserExists, err.Error())
		return
	}
	ResponseWithSuccess(ctx, nil)
}
