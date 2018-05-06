package entities

// Processing
// Executor (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Executor - interface.
*/
type Executor interface { // enumerator
	Execute(string) error
}
