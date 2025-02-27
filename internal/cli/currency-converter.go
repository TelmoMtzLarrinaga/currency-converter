package currencyconverter

import (
	"context"
	"fmt"

	"github.com/TelmoMtzLarrinaga/currency-converter/internal/exchange"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/peterbourgon/ff/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewCurrencyConverterCmd() *ff.Command {
	// create the root command
	fs := ff.NewFlagSet("currency_converter")

	return &ff.Command{
		Name:      "currency_converter",
		Usage:     "cmd [FLAGS] subcmd [FLAGS] <ARGS> [<ARG>...]",
		ShortHelp: "It starts up the currency converter program",
		LongHelp:  "It will convert between a number of base currencies.\n Obtaining real time data through a third party API.\n",
		Flags:     fs,
		Subcommands: []*ff.Command{
			newExchangeSubcommand(),
		},
	}
}

func newExchangeSubcommand() *ff.Command {
	// create default logger config
	logCfg := zap.Config{
		Level:             zap.NewAtomicLevel(), // Info and above
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: true, // Disable automatic stacktrace capturing.
		DisableCaller:     true, // Disable log annotation with calling func.
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			TimeKey:     "ts",
			EncodeTime:  zapcore.ISO8601TimeEncoder, // e.g 2020-07-10 15:00:00.000,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
		},
	}

	// create exchange subcommand
	fs := ff.NewFlagSet("exchange")

	// create new exchange config
	cfg := exchange.NewExchangeConfig()

	// define the flags under exchange subcommand
	_ = fs.StringLong("config", "cc.config", "Config file (optional)")
	fs.BoolVarDefault(&cfg.Debug, '0', "debug", false, "Debug log level")

	return &ff.Command{
		Name:      "exchange",
		Usage:     "cmd [FLAGS] subcmd [FLAGS] <ARGS> [<ARG>...]",
		ShortHelp: "It shows the TUI for the currency exchange",
		LongHelp:  "The user will be able to select from a bse number of currencies shown.\n",
		Flags:     fs,
		Exec: func(_ context.Context, _ []string) error {
			if cfg.Debug {
				logCfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
			}

			// create new logger with specified allowed level
			cfg.Logger = zap.Must(logCfg.Build())

			// run the currency-converter bubble tea program
			p := tea.NewProgram(exchange.InitialModel(cfg))
			if _, err := p.Run(); err != nil {
				return fmt.Errorf("cli: error running the currency-converter bubble tea program: %w", err)
			}
			return nil
		},
	}
}
