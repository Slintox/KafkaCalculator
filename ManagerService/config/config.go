package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		ConfigPath string      `yaml:"-" env:"-"`
		HTTP       HTTPConfig  `yaml:"http" env:"HTTP"`
		Kafka      KafkaConfig `yaml:"kafka" env:"KAFKA"`
	}

	HTTPConfig struct {
		Host string `yaml:"host" env:"HOST" env-default:":8080"`
		Port string `yaml:"port" env:"PORT" env-default:"localhost"`
	}

	KafkaConfig struct {
		URL          string              `yaml:"url" env:"URL" env-default:"calculator:9032" env-description:"Kafka url"`
		Brokers      string              `yaml:"brokers"`
		Expressions  KafkaProducerConfig `yaml:"expressions"`
		Calculations KafkaConsumerConfig `yaml:"calculations"`
	}

	KafkaProducerConfig struct {
		// Consuming
		Topic string `yaml:"topic"`
	}

	KafkaConsumerConfig struct {
		// Producing
		GroupID   string `yaml:"group_id"`
		Topic     string `yaml:"topic"`
		Partition int    `yaml:"partition"`
	}
)

func New(configName string) (*Config, error) {
	var err error
	cfg := &Config{
		ConfigPath: configName,
	}

	if err = initConfig(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func initConfig(cfg *Config) error {
	var err error
	if err := cleanenv.ReadConfig(cfg.ConfigPath, cfg); err != nil {
		return err
	}

	if err = cleanenv.ReadEnv(cfg); err != nil {
		return err
	}

	return nil
}
