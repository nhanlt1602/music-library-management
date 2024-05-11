package services

import (
	"log"
	"music-library-management/models"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Config *models.EnvConfig

func LoadConfig() {
	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault("SERVER_PORT", "8083")
	v.SetDefault("MODE", "debug")
	v.SetConfigType("dotenv")
	v.SetConfigName(".env")
	v.AddConfigPath("./")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&Config); err != nil {
		panic(err)
	}

	if err := Config.Validate(); err != nil {
		panic(err)
	}

	// Increase the timeout duration for connecting to the MongoDB server
	options := options.Client().ApplyURI(Config.MongodbUri).SetConnectTimeout(10 * time.Second)
	err := mgm.SetDefaultConfig(nil, Config.MongodbDatabase, options)
	if err != nil {
		panic(err)
	}
}

func InitMongoDB() {
	// Code for initializing MongoDB
	// Increase the timeout duration for connecting to the MongoDB server
	options := options.Client().ApplyURI(Config.MongodbUri).SetConnectTimeout(10 * time.Second)
	err := mgm.SetDefaultConfig(nil, Config.MongodbDatabase, options)
	if err != nil {
		panic(err)
	}
	// err := mgm.SetDefaultConfig(nil, Config.MongodbDatabase, options.Client().ApplyURI(Config.MongodbUri))
	// if err != nil {
	// 	panic(err)
	// }

	log.Println("Connected to MongoDB!")
}
