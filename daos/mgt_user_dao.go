package daos

import (
	"picbed/global"
	"picbed/models"
)

var MgtUserDaoIns = &MgtUserDao{}

type MgtUserDao struct{}

func (ud *MgtUserDao) Create(user *models.MgtUser) error {
	tx := global.MysqlDB.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ud *MgtUserDao) FindByUsername(username string) (*models.MgtUser, error) {
	var dbUser models.MgtUser
	tx := global.MysqlDB.Unscoped().First(&dbUser, "username = ?", username)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dbUser, nil
}

func (ud *MgtUserDao) FindByUserId(userId int64) (*models.MgtUser, error) {
	user := &models.MgtUser{}
	tx := global.MysqlDB.Where("id = ?", userId).First(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (ud *MgtUserDao) SoftDelete(userId int64) error {
	if err := global.MysqlDB.Delete(&models.MgtUser{}, userId).Error; err != nil {
		return err
	}
	return nil
}

func (ud *MgtUserDao) Update(user *models.MgtUser) error {
	tx := global.MysqlDB.Model(user).Updates(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ud *MgtUserDao) GetUsers(pagination *models.Pagination) ([]*models.MgtUser, error) {
	var users []*models.MgtUser
	tx := global.MysqlDB.Model(models.MgtUser{}).Limit(pagination.GetPageSize()).Offset(pagination.GetOffset()).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
