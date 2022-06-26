package database

import "gorm.io/gorm"

type DatabaseDriver interface {
	CreateConfiguration() (gorm.Dialector, gorm.Option)
}
