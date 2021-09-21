package config

import "github.com/spf13/viper"

type Message struct {
	Start    string `mapstructure:"start"`
	Expenses string `mapstructure:"expenses"`
	Support  string `mapstructure:"support"`
	Idea     string `mapstructure:"idea"`
	FAQ      string `mapstructure:"faq"`
}

type Config struct {
	TelegramToken string
	Admin         int64
	Message       Message
}

func Init() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := fromEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("message", &cfg.Message); err != nil {
		return err
	}

	return nil
}

func fromEnv(cfg *Config) error {
	if err := viper.BindEnv("bot_token"); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString("bot_token")

	if err := viper.BindEnv("admins"); err != nil {
		return err
	}
	cfg.Admin = viper.GetInt64("admins")
	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")
	return viper.ReadInConfig()
}
