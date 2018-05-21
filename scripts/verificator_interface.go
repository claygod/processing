package scripts

// Processing
// Verificator (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"github.com/claygod/processing/domain"
)

/*
Verificator - transaction verification.
*/
type Verificator interface {
	Transaction(*domain.Transaction) error
	Opinion(*domain.Opinion) error
}
