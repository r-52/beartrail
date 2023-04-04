package models

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	Ident string `gorm:"index"`
	Url   sql.NullString
	Name  string

	CompanyUsers     []CompanyUser
	CompanySetting   CompanySetting
	CompanySettingID uint

	RealmID uint
}

func (company *Company) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	company.Ident = uuid.String()
	return
}
