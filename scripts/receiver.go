package scripts

// Processing
// Receiver
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "fmt"
	"github.com/claygod/processing/domain"
)

/*
Receiver
Получатель.
*/
type Receiver interface {
	ReceiveTransaction(*domain.Transaction)
}
