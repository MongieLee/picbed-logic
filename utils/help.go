package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var node *snowflake.Node

// InitSnow 初始化雪花算法
func InitSnow(startTime string, machineId int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineId)
	if err != nil {
		return
	}
	return
}

// GenerateSnowId 生成雪花ID
func GenerateSnowId() int64 {
	if node == nil {
		err := InitSnow("2024-01-01", 1)
		if err != nil {
			ErrorF("InitSnow invoked failed")
		}
	}
	return node.Generate().Int64()
}

const _constSalt = "picbed-logic"

// GenerateRandomSalt 生成随机盐
func GenerateRandomSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(append(salt, []byte(_constSalt)...)), nil
}

// HashPassword 字符串哈希化
func HashPassword(password string, salt string) (string, error) {
	combined := fmt.Sprintf("%v%v", password, salt)
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(combined), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), nil
}
