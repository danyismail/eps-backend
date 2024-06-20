package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/structs"
	"eps-backend/utils"
	"errors"

	"github.com/google/uuid"
)

type UserConstruct struct {
	db db.DBConnection
}

func NewUserStore(db db.DBConnection) *UserConstruct {
	return &UserConstruct{
		db: db,
	}
}

func (c *UserConstruct) Create(userRequest structs.CreateUser) (user *model.User, err error) {
	//check email already exist
	var count int64
	result := c.db.DigiEps.Table("users").Where("email = ?", userRequest.Email).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	if count > 0 {
		return nil, errors.New("email already exist")
	}

	//set hash password
	hash, err := utils.GeneratePassword(userRequest.Password)
	if err != nil {
		return nil, err
	}

	newUser := model.User{
		ID:       model.CustomUUID(uuid.New()),
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: hash,
		Role:     userRequest.Role,
	}

	if err := c.db.DigiEps.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (c *UserConstruct) GetAll(page, view int) (users []model.User, err error) {
	//count data selected
	var count int64
	result := c.db.DigiEps.Find(&model.User{}).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	if page == 0 {
		page = 1
	}
	if view == 0 {
		view = 10
	}
	offset := (page - 1) * view
	hasil := c.db.DigiEps.Debug().Limit(view).Offset(offset).Find(&users)
	if hasil.Error != nil {
		return nil, hasil.Error
	}
	return users, nil
}

func (c *UserConstruct) GetUser(param map[string]interface{}) (result *model.User, err error) {
	var user model.User
	if err := c.db.DigiEps.Debug().Model(&user).Where(param).Scan(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *UserConstruct) Delete(id string) error {
	result := c.db.DigiEps.Debug().Exec("DELETE FROM users WHERE id = ?", id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data found")
	}
	return nil
}

func (c *UserConstruct) Count() int64 {
	var total int64
	result := c.db.DigiEps.Debug().Table("users").Count(&total)
	if result.Error != nil {
		return 0
	}
	return total
}
