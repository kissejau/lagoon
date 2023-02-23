package models

type Update struct {
    UpdateId int        `json:"update_id"`
    Message Message     `json:"message"`
}

type Message struct {
    MessageId int       `json:"message_id"`
    Text string         `json:"text"`
    Chat Chat           `json:"chat"`
}

type Chat struct {
    Id int              `json:"id"`
    Username string     `json:"username"`
}

type Response struct {
    Result []Update       `json"result"`
}
