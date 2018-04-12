package scripts

// Processing
// Authoritys repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"github.com/claygod/processing/entities"
)

/*
AuthorityRepository - storage authorty interface.
*/
type AuthorityRepository interface {
	Create(entities.TokenRepository, string) (int, error)
	List() []Authority
}
