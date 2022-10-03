package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		ConfigPath string      `yaml:"-" env:"-"`
		Kafka      KafkaConfig `yaml:"kafka" env:"KAFKA"`
	}

	KafkaConfig struct {
		URL          string              `yaml:"url" env:"URL" env-default:"invalid:invalid" env-description:"Kafka url"`
		Brokers      string              `yaml:"brokers" env-default:"invalid:invalid"`
		Expressions  KafkaConsumerConfig `yaml:"expressions"`
		Calculations KafkaProducerConfig `yaml:"calculations"`
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
