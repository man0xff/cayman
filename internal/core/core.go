package core

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Params struct {
		ConfigFile string

		DB DBParams
	}

	DBParams struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
)

type Core struct {
	db *gorm.DB

	categories map[CategoryID]*Category
}

func New(params *Params) (*Core, error) {
	c := &Core{}
	// err := c.initDB(params)
	// if err != nil {
	// 	return nil, err
	// }
	err := c.processConfigFile(params.ConfigFile)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Core) initDB(params *Params) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		params.DB.User,
		params.DB.Password,
		params.DB.Host,
		params.DB.Port,
		params.DB.Database)
	dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

func (c *Core) GetCategory(id CategoryID) (*Category, bool) {
	out, ok := c.categories[id]
	return out, ok
}

func (c *Core) ListCategories() []*Category {
	out := make([]*Category, 0, len(c.categories))
	for _, category := range c.categories {
		out = append(out, category)
	}
	return out
}
