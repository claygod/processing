package processing

// Processing
// Group
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"sync"
)

/*
Group - list of accounts.
*/
type Group struct {
	sync.Map
}

/*
NewGroup - create new Group.
*/
func NewGroup() *Group {
	return &Group{}
}

func (g *Group) StoreList(list map[string]interface{}) {
	for k, v := range list {
		g.Store(k, v)

	}
}
