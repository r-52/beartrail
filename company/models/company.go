package models

import (
	"database/sql"

	"github.com/jaevor/go-nanoid"
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

	RealmID string
}

func (company *Company) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := nanoid.Standard(21)
	if err != nil {
		panic(err)
	}
	company.Ident = uid()
	return
}
