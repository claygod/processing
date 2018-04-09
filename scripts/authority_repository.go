package scripts

// Processing
// Authoritys repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
AuthorityRepository - storage authorty interface.
*/
type AuthorityRepository interface {
	Add(*Authority) error
	UrlList() []string
	// AddressList() []string
}
