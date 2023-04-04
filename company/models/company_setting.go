package models

import "gorm.io/gorm"

type CompanySetting struct {
	gorm.Model

	CompanyID uint
}
