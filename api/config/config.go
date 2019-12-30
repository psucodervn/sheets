package config

type LogConfig struct {
	Debug  bool `default:"false"`
	Pretty bool `default:"false"`
}

type ApiConfig struct {
	Address string `default:"0.0.0.0:8080"`
	TLS     bool   `default:"false"`
	Google  GoogleConfig
}

type GoogleConfig struct {
	SheetID         string `envconfig:"SHEET_ID" required:"true"`
	CredentialsFile string `envconfig:"CREDENTIALS_FILE" required:"true"`
}
