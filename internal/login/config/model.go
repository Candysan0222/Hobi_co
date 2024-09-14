package config

import "gorm.io/gorm/logger"

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	LogLevel logger.LogLevel
}
