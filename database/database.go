package database

import (
	"errors"

	"gorm.io/gorm"
)

var db *database

func Create(dDriver DatabaseDriver) error {
	if db != nil {
		return errors.New("database connection already established")
	}
	d := database{}
	err := d.createConn(dDriver)
	if err != nil {
		return err
	}
	db = &d
	return nil
}

func Conn() *gorm.DB {
	return db.db
}

type database struct {
	db *gorm.DB
}

func (d *database) createConn(dDriver DatabaseDriver) error {
	dialector, option := dDriver.CreateConfiguration()
	db, err := gorm.Open(dialector, option)
	if err != nil {
		return err
	}
	d.db = db
	return nil
}
