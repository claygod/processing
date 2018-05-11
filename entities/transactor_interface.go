package entities

// Processing
// Transactor (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Transactor - transaction service.
*/
type Transactor interface {
	Prepare() error
	Rollback() error
	Execute() error
}
