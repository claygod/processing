package entities

// Processing
// Transaction repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
TransactionRepository - storage transaction interface.
This repository is not allowed to delete entities!
*/
type TransactionRepository interface {
	Write(Transaction) error
	Read(string) (Transaction, error) // by hash
	List() []Transaction
}
