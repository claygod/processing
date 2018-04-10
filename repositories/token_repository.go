package repositories

// Processing
// Token repository (implementation)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"sync"

	"github.com/claygod/processing/entities"
)

/*
TokenRepository - storage token (implementation).
This repository is not allowed to delete entities!
*/
type TokenRepository struct {
	sync.RWMutex
	tokens       []entities.Token
	indexAddress map[string]int
	encoder      entities.Encoder
}

/*
NewTokenRepository - create new TokenRepository.
*/
func NewTokenRepository(encoder entities.Encoder) *TokenRepository {
	t := &TokenRepository{
		tokens:       make([]entities.Token, 0),
		indexAddress: make(map[string]int),
		encoder:      encoder,
	}
	return t
}

/*
Create - create new Token.
Return address (to public key).
*/
func (t *TokenRepository) Create(pubKey []byte) (string, error) {
	address := t.encoder.Address(pubKey)
	t.Lock()
	defer t.Unlock()
	if _, ok := t.indexAddress[address]; ok {
		return "", fmt.Errorf("The address %s for this key already exists", address)
	}
	num := len(t.tokens)
	nt := entities.NewToken(address, pubKey)
	t.tokens = append(t.tokens, nt)
	t.indexAddress[address] = num
	return address, nil
}

/*
Read - get a token at his address.
*/
func (t *TokenRepository) Read(address string) (entities.Token, error) {
	t.RLock()
	defer t.RUnlock()
	if tkn, ok := t.indexAddress[address]; ok {
		return t.tokens[tkn], nil
	}
	return entities.Token{}, fmt.Errorf("Address %s not found", address)
}

/*
List - get a tokens list.
*/
func (t *TokenRepository) List() []entities.Token {
	t.RLock()
	defer t.RUnlock()
	ntr := make([]entities.Token, len(t.tokens))
	copy(ntr, t.tokens)
	return ntr
}
