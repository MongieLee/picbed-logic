package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"picbed/models"
	"picbed/services"
	"picbed/utils"
	"strconv"
)

type UserController struct{}

// CreateUser 创建用户
func (uc *UserController) CreateUser(ctx *gin.Context) {
	inputModel := &models.User{}
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
	if dbUser != nil {
		ResponseWithFail(ctx, CodeUserExists)
		return
	}
	err := services.UserServiceIns.Create(inputModel)
	if err != nil {
		ResponseWithFailMsg(ctx, CodeUserExists, err.Error())
		return
	}
	ResponseWithSuccess(ctx, nil)
}

// DeleteUser 删除u用户
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var params map[string]interface{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ResponseWithFailMsg(ctx, CodeInvalidParameters, err.Error())
		return
	}
	userIdFloat, ok := params["userId"].(float64)
	if !ok {
		ResponseWithFail(ctx, CodeInvalidParameters)
		return
	}
	userId := int64(userIdFloat)
	_, err := services.UserServiceIns.GetUserByUserId(userId)
	if err != nil {
		ResponseWithFailMsg(ctx, CodeUserNotExists, err.Error())
		return
	}
	err = services.UserServiceIns.DeleteUserByUserId(userId)
	if err != nil {
		ResponseWithFailMsg(ctx, CodeServerError, err.Error())
		return
	}
	ResponseWithSuccess(ctx, nil)
}

// UpdateUser 更新用户
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		var vErr validator.ValidationErrors
		ok := errors.As(err, &vErr)
		if !ok {
			ResponseWithFailMsg(ctx, CodeInvalidParameters, err.Error())
			return
		}
		translate := vErr.Translate(utils.Trans)
		ResponseWithFailMsg(ctx, CodeInvalidParameters, utils.RemoveTopStruct(translate))
		return
	}
	err := services.UserServiceIns.UpdateUser(&user)
	if err != nil {
		ResponseWithFailMsg(ctx, CodeServerError, err.Error())
		return
	}
	ResponseWithSuccess(ctx, nil)
}

// GetUserById 根据用户id获取用户
func (uc *UserController) GetUserById(ctx *gin.Context) {
	userIdStr := ctx.Query("id")
	if userIdStr == "" {
		ResponseWithFail(ctx, CodeInvalidParameters)
		return
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		ResponseWithFail(ctx, CodeInvalidParameters)
		return
	}
	dbUser, err := services.UserServiceIns.GetUserByUserId(userId)
	if err != nil {
		ResponseWithFail(ctx, CodeUserNotExists)
		return
	}
	ResponseWithSuccess(ctx, dbUser)
}

// GetUsers 获取用户列表
func (uc *UserController) GetUsers(ctx *gin.Context) {
	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		pageSize = 10
	}
	if pageSize < 1 {
		pageSize = 10
	}
	pageNum, err := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	if err != nil {
		pageNum = 1
	}
	if pageNum < 1 {
		pageNum = 1
	}
	pagination := &models.Pagination{
		PageNum:  pageNum,
		PageSize: pageSize,
	}
	list, err := services.UserServiceIns.GetUserList(pagination)
	if err != nil {
		ResponseWithFail(ctx, CodeServerBusy)
		return
	}
	ResponseWithSuccess(ctx, list)
	return
}
