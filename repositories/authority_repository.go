package repositories

// Processing
// Authority repository (implementation)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"sync"

	"github.com/claygod/processing/entities"
	"github.com/claygod/processing/scripts"
)

/*
AuthorityRepository - storage authority (implementation).
This repository is not allowed to delete scripts!
*/
type AuthorityRepository struct {
	sync.RWMutex
	tokenReposithory entities.TokenRepository
	authorities      []scripts.Authority
	indexAddress     map[string]int
	indexLink        map[string]int
	encoder          entities.Encoder
	hash             string // хэш для отсылки и верификации
}

/*
NewAuthorityRepository - create new AuthorityRepository.
*/
func NewAuthorityRepository(tr entities.TokenRepository, encoder entities.Encoder) *AuthorityRepository {
	a := &AuthorityRepository{
		tokenReposithory: tr,
		authorities:      make([]scripts.Authority, 0),
		indexAddress:     make(map[string]int),
		indexLink:        make(map[string]int),
		encoder:          encoder,
	}
	return a
}

/*
Create - create new Authority.
Return address (to public key).
*/
func (a *AuthorityRepository) Create(token entities.Token, link string) (int, error) {
	a.Lock()
	defer a.Unlock()
	if _, ok := a.indexAddress[token.Address]; ok {
		return len(a.authorities),
			fmt.Errorf("Token %s already exists in the database", token.Address)
	}
	if _, ok := a.indexLink[link]; ok {
		return len(a.authorities),
			fmt.Errorf("Link %s is already in the database and is assigned to address %s.", link, token.Address)
	}
	a.setHash()

	num := len(a.authorities)
	na := scripts.NewAuthority(token, link)
	a.authorities = append(a.authorities, na)
	a.indexAddress[token.Address] = num

	return num + 1, nil
}

/*
Read - get a token at his address.

func (a *AuthorityRepository) Read(address string) (scripts.Authority, error) {
	a.RLock()
	defer a.RUnlock()
	if tkn, ok := a.indexAddress[address]; ok {
		return a.authorities[tkn], nil
	}
	return scripts.Authority{}, fma.Errorf("Address %s not found", address)
}
*/
/*
List - get a Authority list.
*/
func (a *AuthorityRepository) List() []scripts.Authority {
	a.RLock()
	defer a.RUnlock()
	nar := make([]scripts.Authority, len(a.authorities))
	copy(nar, a.authorities)
	return nar
}

/*
Hash - get a hash.
*/
func (a *AuthorityRepository) Hash() string {
	a.RLock()
	defer a.RUnlock()
	return a.hash
}

/*
Hash - get a hash.
*/
func (a *AuthorityRepository) setHash() {
	arr := make([]string, 0, len(a.authorities))
	for _, a := range a.authorities {
		arr = append(arr, a.Token.Address)
	}
	sort.Strings(arr)

	data := make([]byte, 0)

	for _, v := range arr {
		data = append(data, []byte(v)...)
	}
	a.hash = a.encoder.Address(data)
}
