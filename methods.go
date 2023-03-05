package bot

import (
    "io"
	"bytes"
	"encoding/json"
	"net/http"
    "github.com/kissejau/lagoon/models"
)

func (bot Bot) SendMessage(botMsg models.BotMessage) error {
	data, err := json.Marshal(botMsg)
	if err != nil {
		return err
	}
	_, err = http.Post(bot.url+bot.token+"/sendMessage", "application/json", bytes.NewReader(data))
    return err
}

func (bot Bot) GetMe() string {
	res, _ := http.Get(bot.url + bot.token + "/getMe")
	data, _ := io.ReadAll(res.Body)
	return string(data)
}
