package database

import "gorm.io/gorm"

type DatabaseConfiguration struct {
	Driver     string
	Hostname   string
	Port       string
	DB         string
	User       string
	Password   string
	Param      map[string]string
	GormOption gorm.Option
}
