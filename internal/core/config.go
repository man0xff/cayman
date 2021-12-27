package core

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type (
	config struct {
		Categories []ConfigCategory
	}
	ConfigCategory struct {
		ID    string `yaml:"id"`
		Title string `yaml:"title"`
		Short string `yaml:"short"`
	}
)

func (c *Core) processConfigFile(file string) error {
	var config config
	err := configor.Load(&config, file)
	if err != nil {
		return err
	}

	c.categories = make(map[CategoryID]*Category, len(config.Categories))
	for _, cat := range config.Categories {
		category := &Category{
			ID:    CategoryID(cat.ID),
			Title: cat.Title,
			Short: cat.Short,
		}
		if _, ok := c.categories[category.ID]; ok {
			return fmt.Errorf("category already exists")
		}
		c.categories[category.ID] = category
	}
	return nil
}
