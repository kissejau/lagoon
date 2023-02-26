package bot

import (
	"github.com/kissejau/lagoon/models"

    "bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Handler func(model models.Message) (models.BotMessage, bool)

type Bot struct {
	url      string
	token    string
	handlers []Handler
}

func New(token string, handlers []Handler) (*Bot, error) {
	b := &Bot{
		url:      "https://api.telegram.org/bot",
		token:    token,
		handlers: handlers,
	}
	return b, nil
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
	for _, handler := range bot.handlers {
		if botMsg, fl := handler(update.Message); fl {
            // LOG
            log.Println(botMsg)
            data, err := json.Marshal(botMsg)
            if err != nil {
                return err
            }
			_, err = http.Post(bot.url+bot.token+"/sendMessage", "application/json", bytes.NewReader(data))
            if err != nil {
                return err
            }
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

func (bot Bot) GetMe() string {
	res, _ := http.Get(bot.url + bot.token + "/getMe")
	data, _ := io.ReadAll(res.Body)
	return string(data)
}
