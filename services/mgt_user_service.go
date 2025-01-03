package services

import (
	"picbed/daos"
	"picbed/models"
	"picbed/utils"
	"time"
)

type MgtUserService struct{}

func (us *MgtUserService) GetUserByUsername(param any) (*models.MgtUser, error) {
	panic("unimplemented")
}

var MgtUserServiceIns = &MgtUserService{}

func (us *MgtUserService) Create(newUser *models.MgtUser) error {
	user := &models.MgtUser{}
	now := time.Now()
	user.CreatedAt = &now
	user.UpdatedAt = &now
	salt, _ := utils.GenerateRandomSalt()
	hashPassword, err := utils.HashPassword(newUser.EncryptedPassword, salt)
	if err != nil {
		return err
	}
	user.EncryptedPassword = hashPassword
	user.Salt = salt
	user.Status = 1
	user.UserId = utils.GenerateSnowId()
	err = daos.MgtUserDaoIns.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *MgtUserService) CreateFromRegister(inputModel *models.LoginDto) error {
	user := &models.MgtUser{}
	now := time.Now()
	user.CreatedAt = &now
	user.UpdatedAt = &now
	user.Username = inputModel.Username
	salt, _ := utils.GenerateRandomSalt()
	hashPassword, err := utils.HashPassword(inputModel.Password, salt)
	if err != nil {
		return err
	}
	user.EncryptedPassword = hashPassword
	user.Salt = salt
	user.Status = 1
	user.UserId = utils.GenerateSnowId()
	err = daos.MgtUserDaoIns.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *MgtUserService) GetUserByUserId(userId int64) (*models.MgtUser, error) {
	return daos.MgtUserDaoIns.FindByUserId(userId)
}

func (us *MgtUserService) DeleteUserByUserId(userId int64) error {
	return daos.MgtUserDaoIns.SoftDelete(userId)
}

func (us *MgtUserService) UpdateUser(updateModel *models.MgtUser) error {
	return daos.MgtUserDaoIns.Update(updateModel)
}

func (us *MgtUserService) GetUserList(pagination *models.Pagination) ([]*models.MgtUser, error) {
	return daos.MgtUserDaoIns.GetUsers(pagination)
}
