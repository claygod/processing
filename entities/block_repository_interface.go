package entities

// Processing
// Block repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
BlockRepository - storage block interface.
This repository is not allowed to delete entities!
*/
type BlockRepository interface {
	Write(Block) error
	Read(string) (Block, error) // by hash
	List() []Block
}
