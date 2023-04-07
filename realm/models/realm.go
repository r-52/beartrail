package models

import (
	"time"

	"github.com/jaevor/go-nanoid"
	"github.com/r-52/beartrail/realm/db"
	"gorm.io/gorm"
)

type Realm struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"-"`
	Hostname  string    `gorm:"index,unique" json:"hostname"`
}

func (realm *Realm) BeforeCreate(tx *gorm.DB) (err error) {
	realm.CreatedAt = time.Now()

	nid, err := nanoid.Standard(32)
	if err != nil {
		return err
	}

	realm.ID = nid()
	return
}

func GetRealmByHostname(hostname string) (Realm, error) {
	var realm Realm
	result := db.Database.Where("hostname = ?", hostname).First(&realm)
	return realm, result.Error
}
