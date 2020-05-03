package config

type LogConfig struct {
	Debug  bool `default:"false"`
	Pretty bool `default:"false"`
}

type ApiConfig struct {
	Address string `default:"0.0.0.0:80"`
	TLS     bool   `default:"false"`
	//Google   GoogleDocsConfig
	Postgres PostgresConfig
	Jira     JiraConfig
	Wakatime WakatimeConfig
}

type GoogleDocsConfig struct {
	SheetID         string `envconfig:"SHEET_ID" required:"true"`
	CredentialsFile string `envconfig:"CREDENTIALS_FILE" required:"true"`
	// Auth            GoogleAuthConfig
}

type GoogleAuthConfig struct {
	ClientID     string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"CLIENT_SECRET" required:"true"`
	CallbackURL  string `envconfig:"CALLBACK_URL" required:"true"`
}

type PostgresConfig struct {
	Host     string `required:"true"`
	Port     int    `default:"5432"`
	User     string `required:"true"`
	Password string `required:"true"`
	DB       string `required:"true"`
	SSLMode  string `default:"disable"`
	Debug    bool   `default:"false"`
}

type MigrationConfig struct {
	Postgres PostgresConfig
}

type ImporterConfig struct {
	Google   GoogleDocsConfig
	Postgres PostgresConfig
}

type JiraConfig struct {
	Username string `required:"true"`
	Password string `required:"true"`
	Host     string `required:"true"`
}

type WakatimeConfig struct {
	ApiKey      string            `split_words:"true" required:"true"`
	Leaderboard string            `default:"3bd8a420-d89e-43f3-8522-2b7fd12549ba"`
	MapID       map[string]string `split_words:"true" required:"true"`
}
