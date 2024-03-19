package bot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MyBot struct {
	Bot *tgbotapi.BotAPI
}

func BotInit() MyBot {
	token := os.Getenv("API_KEY")
	// Initialize bot with token

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return MyBot{
			Bot: nil,
		}
	}
	// Set debug mode
	log.Printf("authorized on account %s", bot.Self.UserName)
	return MyBot{
		Bot: bot,
	}

}

func (b *MyBot) SendMessage(msg error) {
	chatID := int64(-1002082837951)
	chat := tgbotapi.NewMessage(chatID, msg.Error())
	_, err := b.Bot.Send(chat)
	if err != nil {
		log.Println(err)
	}
}
