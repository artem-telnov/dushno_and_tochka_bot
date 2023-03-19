package main

import (
	"time"

	"github.com/artem-telnov/dushno_and_tochka_bot/internal/pkg/bot"
	"github.com/artem-telnov/dushno_and_tochka_bot/internal/pkg/log"
)

func main() {
	time.Local = time.UTC
	logger := log.NewLogger()

	logger.Info("All Rigth!")

	bot, err := bot.New(logger)

	if err != nil {
		logger.Fatal(err)
	}

	bot.StartPolling()
}
