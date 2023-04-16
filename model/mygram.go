package model

import (
	"MyGram/helper"
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Created_At time.Time `json:"created_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
	Updated_At time.Time `json:"updated_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
}

type User struct {
	GormModel
	Username    string    `json:"username" gorm:"not null;unique;type:varchar(255);" binding:"required"`
	Email       string    `json:"email" gorm:"not null;unique;type:varchar(255);" binding:"required"`
	Password    string    `json:"password" gorm:"not null;type:varchar(255);" binding:"required"`
	Age         int       `json:"age" gorm:"not null;"`
	Photos      []Photo   `json:"photos" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments    []Comment `json:"comments" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
	SocialMedia SocialMedia
}

type Photo struct {
	GormModel
	Title    string    `json:"title" gorm:"not null;type:varchar(255);" binding:"required"`
	Caption  string    `json:"caption" gorm:"not null;type:varchar(255);"`
	PhotoUrl string    `json:"photo_url" gorm:"not null;type:varchar(255);" binding:"required"`
	UserID   int       `json:"user_id"`
	Comments []Comment `json:"comments" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
	User     *User
}

type Comment struct {
	GormModel
	Message    string    `json:"message" gorm:"not null;type:varchar(255);" binding:"required"`
	Created_At time.Time `json:"created_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
	Updated_At time.Time `json:"updated_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
	PhotoID    int       `json:"photo_id"`
	UserID     int       `json:"user_id"`
	Photo      *Photo
	User       *User
}

type SocialMedia struct {
	GormModel
	UserID         int       `json:"user_id" gorm:"not null;unique;"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	Created_At     time.Time `json:"created_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
	Updated_At     time.Time `json:"updated_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
	User           *User
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helper.HashPass(u.Password)
	err = nil
	return
}
