package config

import (
	"github.com/restore/payment/handler"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Configuration struct {
	Payment handler.PaymentConfig `yaml:"payment"`
}

var config Configuration

func Init() {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
}

func NewPaymentConfig() *handler.PaymentConfig {
	return &config.Payment
}
