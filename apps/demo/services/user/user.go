package user

import (
	"gobpframe/apps/demo/models/user"
	"gobpframe/server"
	"gobpframe/utils/helper"
	"gobpframe/utils/logger"
)

func AutoMigrate() {
	usr := &user.User{}
	if err := server.DB().AutoMigrate(usr); err != nil {
		logger.Errorf("UserService.AutoMigrate failed, %v", err)
	}
}

func FindById(id int) *user.User {
	return FindOne(map[string]interface{}{"id": id})
}

func FindOne(filters map[string]interface{}) *user.User {
	usr := &user.User{}
	server.DB().Omit("password").Where(filters).First(usr)
	return usr
}

func Find(filters map[string]interface{}, limit, offset int) *[]user.User {
	users := &[]user.User{}
	server.DB().Omit("password").Where(filters).Offset(offset).Limit(limit).Find(&users)
	return users
}

func Create(user *user.User) *user.User {
	user.Password = helper.PasswordHash(user.Password)
	result := server.DB().Create(user)
	if result.Error != nil {
		logger.Errorf("UserService.Create error, %v", result.Error)
	}
	logger.Debugf("UserService.Create rows affected %d", result.RowsAffected)
	user.Password = ""
	return user
}

func UpdateOne(user *user.User) *user.User {
	result := server.DB().Save(user)
	if result.Error != nil {
		logger.Errorf("UserService.UpdateOne error, %v", result.Error)
	}
	user.Password = ""
	return user
}

func RemoveOne(filters map[string]interface{}) *user.User {
	usr := &user.User{}
	server.DB().Where(filters).Delete(&usr)
	usr.Password = ""
	return usr
}

func Login(username, password string) (bool, *user.User) {
	password = helper.PasswordHash(password)
	usr := FindOne(map[string]interface{}{"username": username, "password": password})
	return usr != nil && usr.ID > 0, usr
}

func ChangeStatus(id uint, status int) *user.User {
	return UpdateOne(&user.User{ID: id, Status: status})
}
