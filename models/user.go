package models

import (
	log "github.com/sirupsen/logrus"
)

type User struct {
	Login
	ID        uint   `gorm:"primaryKey" json:"id"`
	TeamId    uint   `gorm:"not null;uniqueIndex" form:"team_id" json:"team_id" validate:"required"`
	Firstname string `gorm:"not null;varchar(50)" form:"first_name" json:"first_name" validate:"required"`
	Lastname  string `gorm:"not null;varchar(50)" form:"last_name" json:"last_name" validate:"required"`
	Posts     []Post `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"posts"`
}

type Login struct {
	Email    string `gorm:"not null;varchar(50)" form:"email" json:"email" validate:"required"`
	Password string `gorm:"not null;varchar(50)" form:"password" json:"password,omitempty" validate:"required"`
}

func (*User) All(teamId float64) (users []User) {
	if err := DB.Preload("Posts").Select("id, email, firstname, lastname, team_id").Where("team_id = ?", teamId).Order("id").Find(&users).Error; err != nil {
		log.Error(err)
		return nil
	}
	return users
}

func (*User) Create(user *User) (err error) {
	if err := DB.Create(&user).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (*User) Show(id string) (user User) {
	if err := DB.Preload("Posts").Select("id, firstname, lastname").First(&user, id).Error; err != nil {
		log.Error(err)
		return
	}

	return user
}

func (u *User) Delete(id string) (err error) {
	if DB.First(&u, id).Error != nil {
		log.Error(err)
		return
	}
	if err := DB.Unscoped().Delete(&u).Error; err != nil {
		log.Error(err)
		return err
	}

	return nil
}
