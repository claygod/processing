package domain

// Processing
// Providers repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
ProviderRepository - storage provider interface.
*/
type ProviderRepository interface {
	Add(*Provider) error
	FindByResource(int) []string
}
