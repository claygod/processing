package repositories

// Processing
// Account repository (implementation)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"sync"

	"github.com/claygod/processing/entities"
)

/*
AccountRepository - accounts store.
*/
type AccountRepository struct {
	accs [256]*accountStore // 256 arrays to reduce access competitiveness
}

/*
NewUnit - create new Unit.
*/
func NewAccountRepository() AccountRepository {
	ar := AccountRepository{}
	for i := 0; i < 256; i++ {
		ar.accs[i] = newAccountStore()
	}
	return ar
}

func (a *AccountRepository) Read(str string) *entities.Account {
	var key byte
	if len(str) > 0 {
		key = []byte(str)[0]
	}
	return a.accs[key].getAccount(str)
}

/*
accountStore - accounts substore.
*/
type accountStore struct {
	sync.Mutex
	accounts map[string]*entities.Account
}

func newAccountStore() *accountStore {
	as := &accountStore{
		accounts: make(map[string]*entities.Account),
	}
	return as
}

func (as *accountStore) getAccount(key string) *entities.Account {
	as.Lock()
	a, ok := as.accounts[key]
	if !ok {
		a = entities.NewAccount()
		as.accounts[key] = a
	}
	as.Unlock()
	return a
}
