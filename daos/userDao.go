package daos

import (
	"picbed/global"
	"picbed/models"
)

var UserDaoIns = &UserDao{}

type UserDao struct{}

func (ud *UserDao) Create(user *models.User) error {
	tx := global.MysqlDB.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ud *UserDao) FindByUsername(username string) (*models.User, error) {
	var dbUser models.User
	tx := global.MysqlDB.Unscoped().First(&dbUser, "username = ?", username)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dbUser, nil
}

func (ud *UserDao) FindByUserId(userId int64) (*models.User, error) {
	user := &models.User{}
	tx := global.MysqlDB.Where("id = ?", userId).First(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (ud *UserDao) SoftDelete(userId int64) error {
	if err := global.MysqlDB.Delete(&models.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}

func (ud *UserDao) Update(user *models.User) error {
	tx := global.MysqlDB.Model(user).Updates(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ud *UserDao) GetUsers(pagination *models.Pagination) ([]*models.User, error) {
	var users []*models.User
	tx := global.MysqlDB.Model(models.User{}).Limit(pagination.GetPageSize()).Offset(pagination.GetOffset()).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
