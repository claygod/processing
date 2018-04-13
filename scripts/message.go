package scripts

// Processing
// Message
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
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

type Parcel struct {
	Sender   string `json:"from"`
	SentTime int64  `json:"exit"`
	Messages []*Message
	R10      string `json:"r10"`
	S10      string `json:"s10"`
}

func (p *Parcel) AddMessage(msg *Message) error {
	for _, m := range p.Messages {
		if m.Hash == msg.Hash {
			return fmt.Errorf("Message with hash %s already added.", msg.Hash)
		}
	}
	p.Messages = append(p.Messages, msg)
	return nil
}

func (p *Parcel) dataForVerification() []byte {
	var data []byte
	data = append(data, []byte(p.Sender)...)
	//data = append(data, []byte(strconv.Itoa(m.Offset))...)
	data = append(data, []byte(strconv.Itoa(int(p.SentTime)))...)
	for _, m := range p.Messages {
		data = append(data, m.dataForVerification()...)
	}
	return data
}

/*
type Stamp struct {
	From     string `json:"from"`
	SendTime int64  `json:"exit"`
	R10      string `json:"r10"`
	S10      string `json:"s10"`
}
*/
