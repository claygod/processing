package processing

// Processing
// Authority
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import "time"

/*
Authority - an important node in the network.
*/
type Authority struct {
	Id         *Id
	PubKey     string   `json:"pub_key"`
	Url        string   `json:"url"`
	Groups     []string `json:"groups_list"`
	Status     int64
	lastUpdate int64
	timeShift  int64
}

/*
NewAuthority - create new Authority.
*/
/*
func NewAuthority(id *Id, url string) *Authority {
	a := &Authority{
		Id:     id,
		Url:    url,
		Groups: make([]*Group, 0),
	}
	return a
}
*/
func (a *Authority) CheckStatus() { // ToDo: the "method" mov to Node
	// Ping
	a.Status = time.Now().Unix()
	//a.timeShift = xx
}
