package config

import (
	"github.com/joho/godotenv"
	"os"
)

func Load(path string) error {
	cfgEnv := os.Getenv("CONFIG_ENV")
	if len(cfgEnv) == 0 {
		cfgEnv = path
	}
	err := godotenv.Load(cfgEnv)

	if err != nil {
		return err
	}

	return nil
}
