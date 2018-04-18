package entities

// Processing
// BlockValidator
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"math/big"
)

/*
BlockValidator - identifier.
*/
type BlockValidator struct {
	encoder   Encoder
	encryptor Encryptor
	accRepo   AccountRepository
}

/*
NewBlockValidator - create new BlockValidator.
*/
func NewBlockValidator(encoder Encoder, encryptor Encryptor, accRepo AccountRepository) *BlockValidator {
	b := &BlockValidator{
		encoder:   encoder,
		encryptor: encryptor,
		accRepo:   accRepo,
	}
	return b
}

/*
Validate - check a block.
*/
func (bv *BlockValidator) Validate(b Block) error {
	broker, err := bv.accRepo.Read(b.Broker)
	if err != nil {
		return err
	}
	_, err = bv.accRepo.Read(b.Owner)
	if err != nil {
		return err
	}
	hash, err := b.Marshalling()
	if err != nil {
		return err
	}
	if b.Hash != bv.encoder.Address(hash) {
		fmt.Errorf("In the hash %s block, it must be %s.", b.Hash, hash)
	}
	r := new(big.Int).SetBytes(b.R)
	s := new(big.Int).SetBytes(b.S)
	if !bv.encryptor.Verify(hash, broker.PubKey, r, s) {
		return fmt.Errorf("An error in the signature (R, S).")
	}
	return nil
}
