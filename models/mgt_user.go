package models

import (
	"time"

	"gorm.io/gorm"
)

type MgtUser struct {
	Id                int64          `json:"id" gorm:"primaryKey" binding:"required"`
	UserId            int64          `json:"userId" gorm:"column:user_id"`
	Username          string         `json:"userName" gorm:"column:username;not null" binding:"required"`
	Nickname          *string        `json:"nickname" gorm:"column:nickname"`
	EncryptedPassword string         `json:"-" gorm:"column:encrypted_password;not null"`
	Salt              string         `json:"-" gorm:"column:salt"`
	Status            int8           `json:"status" gorm:"column:status" binding:"required"`
	Avatar            *string        `json:"avatar" gorm:"column:avatar"`
	CreatedAt         *time.Time     `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt         *time.Time     `json:"-" gorm:"column:updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

func (u MgtUser) TableName() string {
	return "sys_mgt_user"
}
