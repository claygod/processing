package entities

// Processing
// Authoritys repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
AuthorityRepository - storage authorty interface.
*/
type AuthorityRepository interface {
	Create(AccountRepository, string) (int, error)
	List() []Authority
}
