package config

type LogConfig struct {
	Debug  bool `default:"false" split_words:"true"`
	Pretty bool `default:"false" split_words:"true"`
}

type AuthConfig struct {
	JWTSecret string           `split_words:"true" required:"true"`
	Google    GoogleAuthConfig `split_words:"true"`
}

type ApiConfig struct {
	Address  string         `default:"0.0.0.0:80" split_words:"true"`
	TLS      bool           `default:"false" split_words:"true"`
	Postgres PostgresConfig `split_words:"true"`
	Jira     JiraConfig     `split_words:"true"`
	Wakatime WakatimeConfig `split_words:"true"`
	Auth     AuthConfig     `split_words:"true"`

	TelegramBotName string `split_words:"true" required:"true"`
}

type GoogleDocsConfig struct {
	SheetID         string `envconfig:"SHEET_ID" required:"true" split_words:"true"`
	CredentialsFile string `envconfig:"CREDENTIALS_FILE" required:"true" split_words:"true"`
}

type GoogleAuthConfig struct {
	ClientID     string `split_words:"true" required:"true"`
	ClientSecret string `split_words:"true" required:"true"`
	RedirectURL  string `split_words:"true" default:"postmessage"`
}

type PostgresConfig struct {
	Host     string `required:"true" split_words:"true"`
	Port     int    `default:"5432" split_words:"true"`
	User     string `required:"true" split_words:"true"`
	Password string `required:"true" split_words:"true"`
	DB       string `required:"true" split_words:"true"`
	SSLMode  string `split_words:"true" default:"disable"`
	Debug    bool   `default:"false" split_words:"true"`
}

type MigrationConfig struct {
	Postgres PostgresConfig `split_words:"true"`
}

type ImporterConfig struct {
	Google   GoogleDocsConfig `split_words:"true"`
	Postgres PostgresConfig   `split_words:"true"`
}

type JiraConfig struct {
	Username string `required:"true" split_words:"true"`
	Password string `required:"true" split_words:"true"`
	Host     string `required:"true" split_words:"true"`
}

type WakatimeConfig struct {
	ApiKey      string            `split_words:"true" required:"true"`
	Leaderboard string            `default:"3bd8a420-d89e-43f3-8522-2b7fd12549ba" split_words:"true"`
	MapID       map[string]string `split_words:"true" required:"true"`
}
