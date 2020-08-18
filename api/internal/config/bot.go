package config

type BotConfig struct {
	Telegram TelegramConfig `split_words:"true"`
	Postgres PostgresConfig `split_words:"true"`
}

type TelegramConfig struct {
	BotToken string `split_words:"true" required:"true"`
}
