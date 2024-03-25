package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Load(path string) error {
	cfgEnv := os.Getenv("CONFIG_ENV")
	if len(cfgEnv) == 0 {
		cfgEnv = path
	}
	fmt.Println("cfgEnv", cfgEnv)
	err := godotenv.Load(cfgEnv)

	if err != nil {
		return err
	}
	fmt.Println("S3_BUCKET", os.Getenv("S3_BUCKET"))
	return nil
}
