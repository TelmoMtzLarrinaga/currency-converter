package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	currencyconverter "github.com/TelmoMtzLarrinaga/currency-converter/internal/cli"
	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffhelp"
	"go.uber.org/zap"
)

func main() {
	// create new `zap` production logger.
	logger := zap.Must(zap.NewProduction())

	// controls some aspect of parsing behavior.
	opts := []ff.Option{
		ff.WithEnvVarPrefix("CC"),
	}

	// root currency-converter command
	cmd := currencyconverter.NewCurrencyConverterCmd()

	// parse and run
	err := cmd.ParseAndRun(context.Background(), os.Args[1:], opts...)
	if errors.Is(err, ff.ErrUnknownFlag) || errors.Is(err, ff.ErrDuplicateFlag) || errors.Is(err, ff.ErrAlreadyParsed) || errors.Is(err, ff.ErrNoExec) {
		logger.Error("main.go: error parsing commands incoming commands", zap.Error(err))
		os.Exit(1)
	}

	if errors.Is(err, ff.ErrHelp) {
		fmt.Printf("%s\n", ffhelp.Command(cmd))
		return
	}

	logger.Info("Thanks For Using The Currency Converter.")
}
