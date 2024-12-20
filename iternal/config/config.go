package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Driver   string
	SecretPassword string
}

type Config struct {
	DBConfig
}

func (c *Config) readConf() error {
	err := godotenv.Load()

	if err != err {
		log.Fatal("Error loading .env file")
	}

	c.DBConfig = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DRIVER"),
		SecretPassword: os.Getenv("SECRET_PASSWORD"),
	}

	if c.DBConfig.Host == "" ||
		c.DBConfig.Port == "" ||
		c.DBConfig.User == "" ||
		c.DBConfig.Password == "" ||
		c.DBConfig.DBName == "" ||
		c.DBConfig.Driver == "" ||
		c.DBConfig.SecretPassword == ""{
		return errors.New("missing requirement variables")
	}
	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := config.readConf()
	if err != nil {
		log.Fatal(err)
	}
	return config, nil
}
