package repositories

// Processing
// Account repository (implementation)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"sync"

	"github.com/claygod/processing/entities"
)

/*
AccountRepository - storage account (implementation).
This repository is not allowed to delete entities!
*/
type AccountRepository struct {
	sync.RWMutex
	accounts       []entities.Account
	indexAddress map[string]int
	encoder      entities.Encoder
}

/*
NewAccountRepository - create new AccountRepository.
*/
func NewAccountRepository(encoder entities.Encoder) *AccountRepository {
	t := &AccountRepository{
		accounts:       make([]entities.Account, 0),
		indexAddress: make(map[string]int),
		encoder:      encoder,
	}
	return t
}

/*
Create - create new Account.
Return address (to public key).
*/
func (t *AccountRepository) Create(pubKey []byte) (string, error) {
	address := t.encoder.Address(pubKey)
	t.Lock()
	defer t.Unlock()
	if _, ok := t.indexAddress[address]; ok {
		return "", fmt.Errorf("The address %s for this key already exists", address)
	}
	num := len(t.accounts)
	nt := entities.NewAccount(address, pubKey)
	t.accounts = append(t.accounts, nt)
	t.indexAddress[address] = num
	return address, nil
}

/*
Read - get a account at his address.
*/
func (t *AccountRepository) Read(address string) (entities.Account, error) {
	t.RLock()
	defer t.RUnlock()
	if tkn, ok := t.indexAddress[address]; ok {
		return t.accounts[tkn], nil
	}
	return entities.Account{}, fmt.Errorf("Address %s not found", address)
}

/*
List - get a accountss list.
*/
func (t *AccountRepository) List() []entities.Account {
	t.RLock()
	defer t.RUnlock()
	ntr := make([]entities.Account, len(t.accounts))
	copy(ntr, t.accounts)
	return ntr
}
