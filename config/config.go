package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	AdapterService AdapterService
	RabbitMQ       RabbitMQ
	Mongodb        Mongodb
}

type Mongodb struct {
	ConnectionString string
	Address          string
	Username         string
	Password         string
	DbName           string
	ConnectTimeout   time.Duration
	MaxConnIdleTime  int
	MinPoolSize      uint64
	MaxPoolSize      uint64
}

type RabbitMQ struct {
	Connection   string
	EmailEvent   EmailEvent
	ConsumerName string
	ProducerName string
}

type OrderEvent struct {
	Connection string
	Exchange   string
	RoutingKey string
	Queue      string
}

type EmailEvent struct {
	Connection string
	Exchange   string
	RoutingKey string
	Queue      string
}

type CartEvent struct {
	Connection string
	Exchange   string
	RoutingKey string
	Queue      string
}

type AdapterService struct {
	UserService  UserService
	EmailService EmailService
}

type UserService struct {
	AuthURL     string
	UserURL     string
	InternalKey string
}

type ProductService struct {
	BaseURL     string
	InternalKey string
}

type EmailService struct {
	Email string
	Host  string
	Key   string
}

// Get config path for local or docker
func getDefaultConfig() string {
	return "./config/config"
}

// Load config file from given path
func NewConfig() (*Config, error) {
	config := Config{}
	path := os.Getenv("cfgPath")
	if path == "" {
		path = getDefaultConfig()
	}

	v := viper.New()

	v.SetConfigName(path)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	err := v.Unmarshal(&config)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil
}
