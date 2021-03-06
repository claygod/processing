package scripts

// Processing
// Sender
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "fmt"
	"github.com/claygod/processing/domain"
)

/*
Sender - отправляет сообщения.
*/
type Sender interface {
	Transaction(*domain.Transaction)
	Opinion(*domain.Opinion)
}
