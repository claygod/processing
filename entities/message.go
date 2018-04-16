package entities

// Processing
// Message
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"strconv"
)

type Message struct {
	From     string `json:"from"`
	Offset   int    `json:"offset"`
	SentTime int64  `json:"sent_time"`
	Body     []byte `json:"body"`
	Hash     string
}

func NewMessage(from string, offset int, sentTime int64, body []byte) *Message {
	return &Message{
		From:     from,
		Offset:   offset,
		SentTime: sentTime,
		Body:     body,
	}
}

func (m *Message) dataForVerification() []byte {
	var data []byte
	data = append(data, []byte(m.From)...)
	data = append(data, []byte(strconv.Itoa(m.Offset))...)
	data = append(data, []byte(strconv.Itoa(int(m.SentTime)))...)
	data = append(data, m.Body...)
	return data
}
