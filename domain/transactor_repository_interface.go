package domain

// Processing
// Transactors repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
TransactorRepository - storage transactor interface.
*/
type TransactorRepository interface {
	Create(*Transaction) (Transactor, error)
	// Write(string, *Transactor) error
	Read(string) (Transactor, error)
	Delete(string) error
}
