package config

import "os"

type Config struct {
	MongodbUri string
	MongodbDb  string
	SecretJWT  string
}

func NewConfig() *Config {

	return &Config{
		MongodbUri: os.Getenv("MONGODB_URI"),
		MongodbDb:  os.Getenv("MONGODB_DB"),
		SecretJWT:  os.Getenv("SECRET_JWT"),
	}
}
