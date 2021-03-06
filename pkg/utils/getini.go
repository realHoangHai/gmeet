package utils

import (
	"github.com/realHoangHai/gmeet-biz/pkg/log"
	"gopkg.in/ini.v1"
)

func GetIni(section, key, defaultValue string) string {
	cfg, err := ini.Load("./env/config.ini")
	if err != nil {
		log.Fatalf("Failed to read config.ini: %v", err)
	}
	if value := cfg.Section(section).Key(key).String(); value != "" {
		return value
	}
	return defaultValue
}
