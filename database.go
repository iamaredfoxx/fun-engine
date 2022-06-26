package funengine

import (
	"errors"
	"strings"

	"github.com/iamaredfoxx/fun-engine/database"
	"github.com/iamaredfoxx/fun-engine/database/driver"
)

func Database(config database.DatabaseConfiguration) error {

	var dDriver database.DatabaseDriver

	switch strings.ToLower(config.Driver) {
	case "mysql":
		dDriver = driver.MysqlDriver{
			Hostname:   config.Hostname,
			Port:       config.Port,
			DB:         config.DB,
			User:       config.User,
			Password:   config.Password,
			Param:      config.Param,
			GormOption: config.GormOption,
		}
	case "postgres":
		dDriver = driver.PostgresDriver{
			Hostname:   config.Hostname,
			Port:       config.Port,
			DB:         config.DB,
			User:       config.User,
			Password:   config.Password,
			GormOption: config.GormOption,
		}
	case "none":
		return nil
	default:
		return errors.New("database driver not supported")
	}

	err := database.Create(dDriver)
	if err != nil {
		return err
	}

	return nil
}
