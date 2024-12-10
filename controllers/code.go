package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeServerBusy
	CodeUserNotExists
	CodeUserExists
	CodePasswordInValid
	CodeInvalidParameters
	CodeServerError
	CodeResourceNotFound
	CodeUnauthorized
	CodeOperationNotPermitted
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:               "success",
	CodeServerBusy:            "业务繁忙",
	CodeUserNotExists:         "用户不存在",
	CodeUserExists:            "用户已存在",
	CodePasswordInValid:       "账号或密码错误",
	CodeInvalidParameters:     "无效的参数",
	CodeServerError:           "服务器错误",
	CodeResourceNotFound:      "资源不存在",
	CodeUnauthorized:          "鉴权失败",
	CodeOperationNotPermitted: "操作失败",
}

func (r ResCode) Msg() string {
	s, ok := codeMsgMap[r]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return s
}
