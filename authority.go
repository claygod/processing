package processing

// Processing
// Authority
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import "time"

/*
Authority - an important node in the network.
*/
type Authority struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	Status    int64
	timeShift int64
}

/*
NewAuthority - create new Authority.
*/
func NewAuthority(id string, url string) *Authority {
	a := &Authority{
		Id:  id,
		Url: url,
	}
	return a
}

func (a *Authority) CheckStatus() { // ToDo: the "method" mov to Node
	// Ping
	a.Status = time.Now().Unix()
	//a.timeShift = xx
}
