package bot

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/mymmrac/telego"
	"go.uber.org/zap"

	th "github.com/mymmrac/telego/telegohandler"

	"github.com/artem-telnov/dushno_and_tochka_bot/internal/pkg/custompredicates"
	"github.com/artem-telnov/dushno_and_tochka_bot/internal/pkg/eventprocessor"
)

func New(logger *zap.SugaredLogger) (*Bot, error) {
	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	if telegramBotToken == "" {
		err := errors.New("Telegram Bot Token not found. Please specify TELEGRAM_BOT_TOKEN env.")
		return nil, err
	}

	newApiBot, err := telego.NewBot(telegramBotToken, telego.WithLogger(logger))

	if err != nil {
		return nil, err
	}

	botHandler := &Bot{
		bot:    newApiBot,
		logger: logger,
	}

	return botHandler, nil

}

type Bot struct {
	logger *zap.SugaredLogger
	bot    *telego.Bot
}

func (bot *Bot) StartPolling() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{}, 1)

	updates, _ := bot.bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot.bot, updates)

	go func() {
		<-sigs

		bot.logger.Info("Polling is stoping")
		bot.bot.StopLongPolling()
		bh.Stop()
		bot.logger.Info("Long polling stoped")

		done <- struct{}{}
	}()

	defer bh.Stop()
	defer bot.bot.StopLongPolling()

	bh.Use(
		func(next th.Handler) th.Handler {
			return func(bot *telego.Bot, update telego.Update) {
				go next(bot, update)
			}
		},
	)

	bh.Handle(eventprocessor.ProcessStartComand, th.CommandEqual("start"))
	bh.Handle(eventprocessor.ProcessHelpComand, th.CommandEqual("help"))
	bh.Handle(eventprocessor.ProcessProposeTaskFromMessage, th.TextEqual("Хочу предложить задачу"))
	bh.Handle(eventprocessor.ProcessGetLinkFromReply, custompredicates.IsNewProposeTask)
	bh.Handle(eventprocessor.ProcessNotSupportedComandsComand, th.AnyCommand())

	bh.Start()

	<-done
	bot.logger.Info("Long polling is done")
}
