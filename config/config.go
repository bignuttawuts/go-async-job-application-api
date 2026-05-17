package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Kafka    KafkaConfig
	Producer ProducerConfig
}

type ServerConfig struct {
	Port         string
	Prefork      bool
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type KafkaConfig struct {
	Brokers []string
}

type ProducerConfig struct {
	Topic string
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	return Config{
		Server: ServerConfig{
			Port:         os.Getenv("SERVER_PORT"),
			Prefork:      os.Getenv("SERVER_PREFORK") == "true",
			ReadTimeout:  parseDuration(os.Getenv("SERVER_READ_TIMEOUT")),
			WriteTimeout: parseDuration(os.Getenv("SERVER_WRITE_TIMEOUT")),
		},
		Kafka: KafkaConfig{
			Brokers: splitAndTrim(os.Getenv("KAFKA_BROKERS")),
		},
		Producer: ProducerConfig{
			Topic: os.Getenv("PRODUCER_TOPIC"),
		},
	}
}

func splitAndTrim(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

func parseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(fmt.Sprintf("Invalid duration: %s", s))
	}
	return d
}
