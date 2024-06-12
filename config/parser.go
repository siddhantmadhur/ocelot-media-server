package config

import (
	"crypto/rand"
	"crypto/rsa"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port           int    `toml:"port"`
	SecretKey      string `toml:"secret_key"`
	FinishedWizard bool   `toml:"finished_wizard"`
	CacheDir       string `toml:"cache_dir"`
	Mutex          *sync.Mutex
}

func (c *Config) Write() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(configDir+"/ocelot.toml", 0666, os.ModeAppend)
	if err != nil {
		return err
	}
	_, err = toml.NewDecoder(f).Decode(c)
	return err
}

func (c *Config) Read() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	file, err := os.ReadFile(configDir + "/ocelot.toml")
	if err != nil {
		err = nil
		key, err := rsa.GenerateKey(rand.Reader, 32*8)
		if err != nil {
			return err
		}
		f, err := os.Create(configDir + "/ocelot.toml")
		if err != nil {
			return err
		}

		var defaultConfig = Config{
			Port:           8080,
			SecretKey:      key.PublicKey.N.Text(62),
			FinishedWizard: false,
		}
		err = toml.NewEncoder(f).Encode(defaultConfig)
		*c = defaultConfig
		return err
	}
	err = toml.Unmarshal(file, c)

	c.Mutex = &sync.Mutex{}
	return err
}
