package scripts

// Processing
// Authority
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net/url"

	"github.com/claygod/processing/entities"
	// "time"
)

/*
Authority - an important node in the network.
*/
type Authority struct {
	Token entities.Token
	// PubKey     string `json:"pub_key"`
	Url        string `json:"url"`
	Status     int64
	lastUpdate int64
	timeShift  int64
	urlNet     *url.URL
}

/*
func (a *Authority) CheckStatus() { // ToDo: the "method" mov to Node
	// Ping
	a.Status = time.Now().Unix()
	//a.timeShift = xx
}
*/
