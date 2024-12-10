package services

import (
	"picbed/daos"
	"picbed/models"
	"picbed/utils"
	"time"
)

type UserService struct{}

var UserServiceIns = &UserService{}

func (us *UserService) Create(newUser *models.User) error {
	user := &models.User{}
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
	err = daos.UserDaoIns.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) CreateFromRegister(inputModel *models.LoginDto) error {
	user := &models.User{}
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
	err = daos.UserDaoIns.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUserByUsername(username string) (*models.User, error) {
	return daos.UserDaoIns.FindByUsername(username)
}

func (us *UserService) GetUserByUserId(userId int64) (*models.User, error) {
	return daos.UserDaoIns.FindByUserId(userId)
}

func (us *UserService) DeleteUserByUserId(userId int64) error {
	return daos.UserDaoIns.SoftDelete(userId)
}

func (us *UserService) UpdateUser(updateModel *models.User) error {
	return daos.UserDaoIns.Update(updateModel)
}

func (us *UserService) GetUserList(pagination *models.Pagination) ([]*models.User, error) {
	return daos.UserDaoIns.GetUsers(pagination)
}
