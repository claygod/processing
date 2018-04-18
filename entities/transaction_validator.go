package entities

// Processing
// TransactionValidator
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"math/big"
)

/*
TransactionValidator - identifier.
*/
type TransactionValidator struct {
	encoder   Encoder
	encryptor Encryptor
	accRepo   AccountRepository
	authRepo  AuthorityRepository
	blockRepo BlockRepository
	fee       Fee
}

/*
NewTransactionValidator - create new TransactionValidator.
*/
func NewTransactionValidator(
	encoder Encoder,
	encryptor Encryptor,
	accRepo AccountRepository,
	authRepo AuthorityRepository,
	blockRepo BlockRepository,
	fee Fee) *TransactionValidator {
	b := &TransactionValidator{
		encoder:   encoder,
		encryptor: encryptor,
		accRepo:   accRepo,
		authRepo:  authRepo,
		blockRepo: blockRepo,
		fee:       fee,
	}
	return b
}

/*
Validate - проверка заполненной транзакции (API).
*/
func (bv *TransactionValidator) Validate(t *Transaction) error { // 1 fee = сотая часть процента
	switch t.Type {
	case TransactionTypeTransfer:
		bv.checkTransfer(t)
	case TransactionTypeJoin:

	case TransactionTypeSeparation:

	case TransactionTypeExchange:
	default:
		fmt.Errorf("Unsupported transaction type - %d.", t.Type)
	}
	return nil
}

/*
validateSignature - проверяем хэш и ключи и наличие инициатора и брокера.
*/
func (bv *TransactionValidator) validateSignature(t *Transaction) error {
	broker, err := bv.authRepo.Read(t.Broker)
	if err != nil {
		return err
	}
	if _, err := bv.accRepo.Read(t.Initiator); err != nil {
		return err
	}
	hash, err := t.Marshalling()
	if err != nil {
		return err
	}
	if t.Hash != bv.encoder.Address(hash) {
		fmt.Errorf("In the hash %s transaction, it must be %s.", t.Hash, hash)
	}
	r := new(big.Int).SetBytes(t.R)
	s := new(big.Int).SetBytes(t.S)
	if !bv.encryptor.Verify(hash, broker.Account.PubKey, r, s) {
		return fmt.Errorf("An error in the signature (R, S).")
	}
	return nil
}

/*
validateBlockLists - check slice blocks.
*/
func (bv *TransactionValidator) validateBlocks(blocks []Block) error {
	for _, b := range blocks {
		if err := bv.validateBlock(b); err != nil {
			return nil
		}
	}
	return nil
}

/*
validate - check a block.
*/
func (bv *TransactionValidator) validateBlock(b Block) error {
	broker, err := bv.authRepo.Read(b.Broker)
	if _, err := bv.authRepo.Read(b.Broker); err != nil {
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
	if !bv.encryptor.Verify(hash, broker.Account.PubKey, r, s) {
		return fmt.Errorf("An error in the signature (R, S).")
	}
	return nil
}

/*
func (bv *TransactionValidator) inAmount(t *Transaction) int {
	inAmount := 0
	for _, b := range t.Inputs {
		inAmount += b.State.Amount
	}
	return inAmount
}
*/
