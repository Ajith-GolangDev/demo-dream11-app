package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
}

type Wallet struct {
	gorm.Model
	UserID  uint    `json:"user_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User    User    `gorm:"foreignKey:UserID"`
	Balance float64 `json:"balance" gorm:"type:decimal(10,2);not null"`
}

type Contest struct {
	gorm.Model
	Name     string  `json:"name" gorm:"type:varchar(100);not null"`
	EntryFee float64 `json:"entry_fee" gorm:"type:decimal(10,2);not null"`
}

type Player struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(100);not null"`
	Team        string  `json:"team" gorm:"type:varchar(100);not null"`
	CreditScore float64 `json:"credit_score" gorm:"type:decimal(10,2);not null"`
}

type UserTeam struct {
	gorm.Model
	UserID    uint    `json:"user_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User      User    `gorm:"foreignKey:UserID"`
	ContestID uint    `json:"contest_id" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Contest   Contest `gorm:"foreignKey:ContestID"`
	PlayerIDs string  `json:"player_ids" gorm:"type:text;not null"` // Comma separated player IDs
}
