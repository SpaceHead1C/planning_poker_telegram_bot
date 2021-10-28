package bot

import "github.com/BurntSushi/toml"

type Config struct {
	BotToken    string `toml:"bot_token"`
	BotAddress  string `toml:"bot_address"`
	BotPort     string `toml:"bot_port"`
	TelegramURL string `toml:"telegram_url"`
	CertPath    string `toml:"cert_path"`
	KeyPath     string `toml:"key_path"`
	LogsPath    string `toml:"logs_path"`
}

func ParseConfig(conf *Config) error {
	if _, err := toml.DecodeFile("./configs/config.toml", conf); err != nil {
		return err
	}

	return nil
}

func NewConfig() *Config {
	return &Config{
		TelegramURL: "https://api.telegram.org/bot",
	}
}
