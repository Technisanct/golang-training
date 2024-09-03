package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang-training/storage/mongodb"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	configFile = "config.yaml"

	envKey = "TS_ENV" // env variable name to get deployment environment
)

type configuration struct {
	Env      string    `yaml:"env"`
	Secret   string    `yaml:"secret" json:"secret"`
	PortHTTP uint16    `yaml:"portHTTP"`
	PortGRPC uint16    `yaml:"portGRPC"`
	Database *database `yaml:"database" json:"database"`
	Cors     *cors     `yaml:"cors" json:"cors"`
}

type database struct {
	MongoDB *mongoDB `yaml:"mongoDB" json:"mongoDB"`
}

type cors struct {
	AllowedOrigins []string `yaml:"allowedOrigins" json:"allowedOrigins"`
	AllowedMethods []string `yaml:"allowedMethods" json:"allowedMethods"`
	AllowedHeaders []string `yaml:"allowedHeaders" json:"allowedHeaders"`
	ExposeHeaders  []string `yaml:"exposeHeaders" json:"exposeHeaders"`
}

type mongoDB struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	DBName   string `yaml:"dbName" json:"dbName"`
	Endpoint string `yaml:"endpoint" json:"endpoint"`
	Client   *mongo.Client
}

var config *configuration

func Init() {
	if err := loadConfig(); err != nil {
		panic(err)
	}

	initMongoDB()
}

func loadConfig() error {
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}

	return nil
}

// Get ... return Configuration struct with values
func Get() *configuration {
	if config == nil {
		panic("Configuration not loaded. Application crashed!")
	}
	return config
}

func initMongoDB() {
	mongoDBClient := mongodb.NewClient(
		config.Env,
		config.Database.MongoDB.Username,
		config.Database.MongoDB.Password,
		config.Database.MongoDB.DBName,
		config.Database.MongoDB.Endpoint)

	config.Database.MongoDB.Client = mongoDBClient
}
