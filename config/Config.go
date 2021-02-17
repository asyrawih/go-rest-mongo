package config

import (
	exception "github.com/hananloser/rest-fiber-mongo-2/exception"
	"github.com/joho/godotenv"
	"os"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
	// @todo Adding Implementation If Needed
}

// Get .ENV Settings
func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

// Load The .ENV files
func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	if err != nil {
		exception.PanicIfNeeded(err)
	}
	return &configImpl{}
}
