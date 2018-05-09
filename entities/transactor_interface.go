package entities

// Processing
// Transactor (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Transactor - transaction service.
*/
type Transactor interface {
	Prepare(*Transaction) error
	Rollback(string) error
	Execute(string) error
}
