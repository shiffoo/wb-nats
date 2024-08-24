package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/shiffoo/wb-nats-streaming/internal/logger"
)

var CONFIG Config

type Config struct {
	Env    string `yaml:"env"`
	Server struct {
		URL string `yaml:"address" env-default:"localhost:8080"`
	} `yaml:"server"`
	Nats struct {
		URL            string `yaml:"url"`
		StreamName     string `yaml:"stream-name"`
		StreamSubjects string `yaml:"stream-subjects"`
		SubjectName    string `yaml:"subject-name"`
		Subscriber     string `yaml:"subscriber"`
	} `yaml:"nats"`
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     int64  `yaml:"port"`
	} `yaml:"db"`
}

func InitConfig() error {
	logg := logger.GetLogger()
	if err := cleanenv.ReadConfig("../../config/local.yml", &CONFIG); err != nil {
		logg.Error(err)
		return err
	}
	logg.Info("Init config")
	return nil
}
