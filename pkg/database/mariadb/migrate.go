package mariadb

import (
	"freepass-bcc/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
		&entity.SessionProposal{},
		&entity.Session{},
		&entity.SessionRegistration{},
		&entity.Feedback{},
	)

	if err != nil {
		return err
	}

	return nil
}
