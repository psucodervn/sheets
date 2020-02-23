package config

type LogConfig struct {
	Debug  bool `default:"false"`
	Pretty bool `default:"false"`
}

type ApiConfig struct {
	Address string `default:"0.0.0.0:80"`
	TLS     bool   `default:"false"`
	Google  GoogleConfig
}

type GoogleConfig struct {
	SheetID         string `envconfig:"SHEET_ID" required:"true"`
	CredentialsFile string `envconfig:"CREDENTIALS_FILE" required:"true"`
	Auth            GoogleAuthConfig
}

type GoogleAuthConfig struct {
	ClientID     string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"CLIENT_SECRET" required:"true"`
	CallbackURL  string `envconfig:"CALLBACK_URL" required:"true"`
}
