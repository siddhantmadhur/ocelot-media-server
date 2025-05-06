package storage

import (
	"os"

	"github.com/BurntSushi/toml"
)

type GeneralSettings struct {
	CompletedSetup bool `json:"completed_setup"`
}

type Settings struct {
	General GeneralSettings `json:"general"`
}

func GetDefaultSettings() Settings {
	return Settings{
		General: GeneralSettings{
			CompletedSetup: false,
		},
	}
}

func GetSettings() (*Settings, error) {
	persistentDir, err := GetPersistentDir()
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(persistentDir + "/settings.toml")
	if err != nil {
		if os.IsNotExist(err) {
			var defaultSettings = GetDefaultSettings()
			defaultSettings.Save()
			return &defaultSettings, nil
		}
		return nil, err
	}
	var s Settings
	_, err = toml.Decode(string(b), &s)
	return &s, err
}

func (s *Settings) Save() error {
	persistentDir, err := GetPersistentDir()
	if err != nil {
		return err
	}
	b, err := toml.Marshal(*s)
	if err != nil {
		return err
	}
	err = os.WriteFile(persistentDir+"/settings.toml", b, os.ModePerm)
	return err
}
