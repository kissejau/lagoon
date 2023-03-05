package bot

import (
	"github.com/kissejau/lagoon/models"

	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Handler func(bot Bot, message models.Message)

type Bot struct {
	url      string
	token    string
	handlers []Handler
}

func New(token string) (*Bot, error) {
	b := &Bot{
		url:   "https://api.telegram.org/bot",
		token: token,
	}
	return b, nil
}

func (bot *Bot) RegistrateHandler(handler Handler) {
	bot.handlers = append(bot.handlers, handler)
}

func (bot Bot) Run() {
	offset := 0
	for {
		updates, err := bot.getUpdates(offset)
		if err != nil {
			log.Panic(err)
		}
		log.Println(updates)

		for _, update := range updates {
			err := bot.respond(update)
			offset = update.UpdateId + 1
			if err != nil {
				panic(err)
			}
		}
	}
}

func (bot Bot) respond(update models.Update) error {
	if update.Message.Text[0] == '/' {
		log.Printf("command is execed")
	} else {
		for _, handler := range bot.handlers {
			handler(bot, update.Message)
		}
	}
	return nil
}

func (bot Bot) getUpdates(offset int) ([]models.Update, error) {
	res, err := http.Get(bot.url + bot.token + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	var response models.Response

	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response.Result, nil
}

