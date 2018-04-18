package entities

// Processing
// Fee (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Fee - commission calculation interface.
*/
type Fee interface {
	/*
		Count - calculation of fees.
		In case of an error, a negative number is returned.
	*/
	Count(int) int
}
