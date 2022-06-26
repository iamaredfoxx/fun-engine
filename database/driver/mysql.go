package driver

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDriver struct {
	Hostname   string
	Port       string
	DB         string
	User       string
	Password   string
	Param      map[string]string
	GormOption gorm.Option
}

func (o MysqlDriver) CreateConfiguration() (gorm.Dialector, gorm.Option) {
	return mysql.New(mysql.Config{
		DSN: o.createDSN(),
	}), o.GormOption
}

func (o MysqlDriver) createDSN() string {
	pattern := "%s:%s@tcp(%s:%s)/%s"
	dns := fmt.Sprintf(
		pattern,
		o.User,
		o.Password,
		o.Hostname,
		o.Port,
		o.DB,
	)
	if len(o.Param) > 0 {
		dns = strings.Join([]string{dns, o.createParams()}, "?")
	}
	return dns
}

func (o MysqlDriver) createParams() string {
	params := []string{}
	for i, v := range o.Param {
		params = append(params, fmt.Sprintf("%s=%s", i, v))
	}
	return strings.Join(params, "&")
}
