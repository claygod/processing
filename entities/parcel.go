package entities

// Processing
// Parcel
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"strconv"
)

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
