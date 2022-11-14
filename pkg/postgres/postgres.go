package postgres

import (
	"fmt"

	"github.com/ichetiva/todo-golang/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s timezone=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		"disable",
		"Europe/Moscow",
	)
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
