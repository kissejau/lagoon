package bot

import (
    "lagoon/models"
    "net/http"
    "strconv"
    "io"
    "encoding/json"
    "log"
)

type Bot struct {
    url string
    token string 
}

func New(token string) (*Bot, error) {
    b := &Bot{
        url:    "https://api.telegram.org/bot",
        token:  token,
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
            offset = update.UpdateId + 1
        }
    }
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
