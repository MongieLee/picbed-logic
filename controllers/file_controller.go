package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"picbed/utils"
	"strings"
	"time"
)

type FileController struct{}

func (fc *FileController) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ResponseWithFail(ctx, CodeInvalidParameters)
		return
	}
	parts := strings.Split(file.Filename, ".")
	fileSuffix := parts[len(parts)-1]
	saveFileName := fmt.Sprintf("%d%s.%s", time.Now().Unix(), utils.GenerateRandomChars(6), fileSuffix)
	dst := fmt.Sprintf("uploads/%s/%s", time.Now().Format("20060102"), saveFileName)
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		ResponseWithFail(ctx, CodeServerError)
	}
	returnUrl := "http://localhost:9999/" + dst
	ResponseWithSuccessMsg(ctx, "上传成功", returnUrl)
}
