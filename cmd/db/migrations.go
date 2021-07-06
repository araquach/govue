package db

import (
	"contra-design.com/govue/cmd/models"
)

func Migrate() {
	DB.AutoMigrate(models.User{}, models.Test{})
}
