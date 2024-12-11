package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"picbed/models"
	"picbed/services"
	"picbed/utils"
)

type AuthController struct{}

func (uc *AuthController) Register(ctx *gin.Context) {
	inputModel := &models.RegisterDto{}
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
	dbUser, _ := services.UserServiceIns.GetUserByUsername(inputModel.Username)
	if dbUser != nil {
		ResponseWithFail(ctx, CodeUserExists)
		return
	}
	err := services.UserServiceIns.CreateFromRegister(&models.LoginDto{Username: inputModel.Username, Password: inputModel.Password})
	if err != nil {
		ResponseWithFail(ctx, CodeServerBusy)
		return
	}
	ResponseWithSuccess(ctx, nil)
}

func (uc *AuthController) Login(ctx *gin.Context) {
	inputModel := &models.LoginDto{}
	if err := ctx.ShouldBindJSON(inputModel); err != nil {
		var vErr validator.ValidationErrors
		ok := errors.As(err, &vErr)
		if !ok {
			ResponseWithFailMsg(ctx, CodeInvalidParameters, err.Error())
		}
		translate := vErr.Translate(utils.Trans)
		ResponseWithFailMsg(ctx, CodeInvalidParameters, utils.RemoveTopStruct(translate))
		return
	}
	dbUser, _ := services.UserServiceIns.GetUserByUsername(inputModel.Username)
	if dbUser == nil {
		ResponseWithFail(ctx, CodePasswordInValid)
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.EncryptedPassword), []byte(inputModel.Password+dbUser.Salt))
	if err != nil {
		ResponseWithFail(ctx, CodePasswordInValid)
		return
	}
	token, err := utils.GenerateAccessToken(dbUser)
	if err != nil {
		ResponseWithFailMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	refreshToken, err := utils.GenerateRefreshToken(dbUser)
	if err != nil {
		ResponseWithFailMsg(ctx, CodeServerBusy, err.Error())
		return
	}
	ResponseWithSuccess(ctx, gin.H{
		"token":        token,
		"refreshToken": refreshToken,
		"userInfo":     dbUser,
	})
}

func (uc *AuthController) RefreshToken(ctx *gin.Context) {
	value, exists := ctx.Get("userInfo")
	if !exists {
		ResponseWithUnAuthorized(ctx)
		return
	}
	claims, ok := value.(*utils.CustomClaims)
	if !ok {
		ResponseWithUnAuthorized(ctx)
		return
	}

	err := claims.Valid()
	if err != nil {
		ResponseWithUnAuthorized(ctx)
	}
	user := &models.User{UserId: claims.UserId, Username: claims.Username}
	token, err := utils.GenerateAccessToken(user)
	if err != nil {
		ResponseWithFailMsg(ctx, CodeServerBusy, err.Error())
	}
	refreshToken, err := utils.GenerateRefreshToken(user)
	if err != nil {
		ResponseWithFailMsg(ctx, CodeServerBusy, err.Error())
	}
	ResponseWithSuccess(ctx, gin.H{
		"token":        token,
		"refreshToken": refreshToken,
	})
}
