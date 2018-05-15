package domain

// Processing
// WaitTransaction (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
WaitTransaction - storage wait-transaction interface.
*/
type WaitTransaction interface {
	TransactionRepository
	Delete(string) error
}
