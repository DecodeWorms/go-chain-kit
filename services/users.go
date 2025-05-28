package services

import (
	"go-chain-kit/idgenerator"
	"go-chain-kit/models"
	"gorm.io/gorm"
)

type Users struct {
	db    *gorm.DB
	idGen idgenerator.IdGenerator
}

func NewUser(db *gorm.DB, idGen idgenerator.IdGenerator) *Users {
	return &Users{
		db:    db,
		idGen: idGen,
	}
}

func (u Users) Create(data *models.UserReq) (string, error) {
	rec := &models.User{
		ID:       u.idGen.Generate(),
		Username: data.UserName,
		Email:    data.Email,
	}
	// Persist the record into the db
	if err := u.db.Create(rec).Error; err != nil {
		return "", err
	}
	if rec != nil {
		return rec.ID, nil
	}
	return rec.ID, nil
}

func (u Users) UserByID(ID string) (*models.User, error) {
	var us *models.User
	if err := u.db.Where("id = ?", ID).Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}
