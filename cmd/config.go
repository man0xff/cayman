package main

import (
	"os"

	"github.com/jinzhu/configor"
)

type config struct {
}

func loadConfigFile() *config {
	if len(os.Args) < 2 {
		panic("config file must be specified")
	}
	var config config
	configor.Load(&config, os.Args[1])
	return &config
}
