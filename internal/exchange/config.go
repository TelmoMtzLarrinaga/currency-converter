package exchange

import "go.uber.org/zap"

type Config struct {
	Logger   *zap.Logger
	Debug    bool
	Quantity float64
	ID       string
}

func NewExchangeConfig() *Config {
	return &Config{
		Debug: false,
	}
}
