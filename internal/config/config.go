package config

import (
	"os"
	"errors"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Topic struct {
	Name    string
	Capacity int
}

type Config struct {
	Host     string
	SwaggerURL  string
	Topics []Topic
}

func NewConfig() (Config, error) {
	var config Config
	err := godotenv.Load()
	if err != nil {
		return config, err
	}
	topics, err := getEnvAsTopics("TOPICS")
	if err != nil {
		return config, err
	}
	config = Config{
		Host:  getEnvString("HOST", "localhost:7000"),
		SwaggerURL:  getEnvString("SWAGGER_URL", "localhost:7000"),
		Topics:  topics,
	}
	return config, nil
}

func getEnvString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsTopics(key string) ([]Topic, error) {
	topics := make([]Topic, 0)
	if value, exists := os.LookupEnv(key); exists {
		if value == "" {
			return topics, errors.New("topic list is empty")
		}
		topicsWithCapacity := strings.Split(value, ",")
		for _, topicWithCapacity := range topicsWithCapacity {
			parts := strings.Split(topicWithCapacity, ":")
			if len(parts) != 2 {
				return topics, errors.New("error to split topic config by : ")
			}
			capacity, err := strconv.Atoi(parts[1])
			if err != nil {
				return topics, errors.New("error to convert capacity to int")
			}
			topics = append(topics, Topic{
				Name:    parts[0],
				Capacity: capacity,
			})
		}
	}
	return topics, nil
}
