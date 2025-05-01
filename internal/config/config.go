package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port        int
	DatabaseUrl string
}

func (c *Config) ReadEnvironment() {
	c.Port, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))

}
