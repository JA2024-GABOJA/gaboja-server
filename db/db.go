package db

import (
	"entgo.io/ent/dialect"
	"fmt"
	"junction/config"
	"junction/internal/model/ent"

	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	config *config.Config
}

func NewDb(filePath string) (*ent.Client, error) {
	c := &Db{
		config: config.NewConfig(filePath),
	}

	user := c.config.Db.User
	password := c.config.Db.Pass
	host := c.config.Db.Host
	port := c.config.Db.Port
	database := c.config.Db.Database

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	return ent.Open(dialect.MySQL, dataSourceName)
}
