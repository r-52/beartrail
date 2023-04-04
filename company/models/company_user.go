package models

import "gorm.io/gorm"

type CompanyUser struct {
	gorm.Model

	CompanyID uint
	UserID    uint
}
