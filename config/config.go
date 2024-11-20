package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var Instance *Config

type Config struct {
	Port    string `yaml:"port"`
	Charset string `yaml:"Charset"`
	Env     string `yaml:"Env"`

	PostgreSql struct {
		DbHost string `yaml:"db_host"`
		DbPort int    `yaml:"db_port"`
		DbUser string `yaml:"db_user"`
		DbPass string `yaml:"db_pass"`
		DbName string `yaml:"db_name"`
	} `yaml:"postgresql"`

	Redis struct {
		DbHost string `yaml:"db_host"`
		DbPass string `yaml:"db_pass"`
		DbPort string `yaml:"db_port"`
		DbName string `yaml:"db_name"`
	} `yaml:"redis"`
}

func Init(filename string) *Config {
	Instance = &Config{}
	if yamlFile, err := os.ReadFile(filename); err != nil {
		fmt.Print(err.Error())
	} else if err = yaml.Unmarshal(yamlFile, Instance); err != nil {
		fmt.Print(err.Error())
	}
	return Instance
}
