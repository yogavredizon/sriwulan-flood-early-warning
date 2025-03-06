package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Database struct {
	DB Postgres `yaml:"db"`
}

type Postgres struct {
	Config map[string]string `yaml:"postgres"`
}

func FetchYAML() (string, error) {
	yamlfile, err := os.ReadFile("../properties/dev-properties.yaml")
	if err != nil {
		return "", err
	}

	var database Database
	err = yaml.Unmarshal(yamlfile, &database)
	if err != nil {
		return "", err
	}

	var encPass string = database.DB.Config["password"]
	var password string
	password, err = Decrypt(encPass, []byte(os.Getenv("key")), []byte(os.Getenv("iv")))
	if err != nil {
		return "", err
	}

	database.DB.Config["password"] = password
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", database.DB.Config["host"], database.DB.Config["port"], database.DB.Config["user"], database.DB.Config["password"], database.DB.Config["dbname"], database.DB.Config["sslmode"])

	return dsn, nil
}
