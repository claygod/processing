package entities

// Processing
// Blocks repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
BlockRepository - storage blocks interface.
*/
type BlockRepository interface {
	Write(*Block) error
	Read(int64) (*Block, error)
}
