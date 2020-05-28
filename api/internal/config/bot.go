package config

type BotConfig struct {
	Telegram TelegramConfig
	Postgres PostgresConfig
}

type TelegramConfig struct {
	BotToken string `split_words:"true" required:"true"`
}
