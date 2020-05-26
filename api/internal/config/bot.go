package config

type BotConfig struct {
	Telegram TelegramConfig
	Postgres PostgresConfig
}

type TelegramConfig struct {
	Token string `split_words:"true" required:"true"`
}
