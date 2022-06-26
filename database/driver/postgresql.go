package driver

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDriver struct {
	Hostname   string
	Port       string
	DB         string
	User       string
	Password   string
	GormOption gorm.Option
}

func (o PostgresDriver) CreateConfiguration() (gorm.Dialector, gorm.Option) {
	return postgres.New(postgres.Config{
		DSN: o.createDSN(),
	}), o.GormOption
}

func (o PostgresDriver) createDSN() string {
	pattern := "host=%s user=%s password=%s dbname=%s port=%s"
	dns := fmt.Sprintf(
		pattern,
		o.Hostname,
		o.User,
		o.Password,
		o.DB,
		o.Port,
	)

	return dns
}
