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
func NewTransactionValidator(encoder Encoder, encryptor Encryptor,
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
Validate - проверка заполненной транзакции.
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
checkTransfer - проверка заполненной транзакции 'Transfer'.
*/
func (bv *TransactionValidator) checkTransfer(t *Transaction) error {
	// проверка на наличие аккаунтов брокера и инициатора
	if _, err := bv.accRepo.Read(t.Initiator); err != nil {
		return err
	}
	if _, err := bv.authRepo.Read(t.Broker); err != nil {
		return err
	}
	// проверка на количество входов/выходов
	if len(t.Inputs) == 0 || len(t.Outputs) != 2 {
		fmt.Errorf("Few inputs (%d) and outputs (%d)", len(t.Inputs), len(t.Outputs))
	}
	// проверка на совпадение входных и выходных ресурсов
	if t.Inputs[0].State.ResourceId != t.Outputs[0].State.ResourceId {
		fmt.Errorf("Input resources (%d) do not coincide with output resources (%d).",
			t.Inputs[0].State.ResourceId, t.Outputs[0].State.ResourceId)
	}
	// проверка входов
	if err := t.checkTransferInputs(); err != nil {
		return err
	}
	// проверка выходов
	if err := bv.checkTransferOutputs(t); err != nil {
		return err
	}
	return nil
}

/*
checkTransferInputs - проверка входов заполненной транзакции 'Transfer'.
*/
func (bv *TransactionValidator) checkTransferInputs(t *Transaction) error {
	resource := t.Inputs[0].State.ResourceId
	for _, b := range t.Inputs {
		if t.Initiator != b.Owner {
			return fmt.Errorf("Transfer: entries belong to more than one owner - %s, %s...", t.Initiator, b.Owner)
		}
		if resource != b.State.ResourceId {
			return fmt.Errorf("You can only transfer one resource at a time - %s, %s...", resource, b.State.ResourceId)
		}
		if b.Condition.Amount != 0 || b.Condition.ResourceId != 0 {
			return fmt.Errorf("In the transfer you can not use the inputs available for exchange.")
		}
	}
	return nil
}

/*
checkTransferOutputs - проверка выходов заполненной транзакции 'Transfer'.
*/
func (bv *TransactionValidator) checkTransferOutputs(t *Transaction) error {
	// проверка на совпадение входных и выходных ресурсов
	resource := t.Outputs[0].State.ResourceId
	for _, b := range t.Outputs {
		// проверка корректности самих блоков
		if err := bv.validateBlock(b); err != nil {
			return err
		}
		// совпадают ли брокеры транзакции и блоков
		if t.Broker != b.Broker {
			return fmt.Errorf("Brokers of transaction (%s) and block (%s) do not match.", t.Broker, b.Broker)
		}
		// выходные ресурсы должны быть одинаковыми
		if resource != b.State.ResourceId {
			return fmt.Errorf("You can only transfer one resource at a time - %s, %s...", resource, b.State.ResourceId)
		}
		// обменные условия должны быть обнулены
		if b.Condition.Amount != 0 || b.Condition.ResourceId != 0 {
			return fmt.Errorf("In the transfer you can not use the outputs available for exchange.")
		}
	}
	// поиск брокера и инициатора (отправителя)
	var brokerBlock *Block
	// var recipientBlock *Block
	if t.Outputs[0].Owner == t.Broker && t.Outputs[1].Owner == t.Initiator {
		brokerBlock = &t.Outputs[0]
		//recipientBlock = &t.Outputs[1]
	} else if t.Outputs[0].Owner == t.Broker && t.Outputs[0].Owner == t.Initiator {
		brokerBlock = &t.Outputs[1]
		//recipientBlock = &t.Outputs[0]
	} else {
		fmt.Errorf("Error in outputs.")
	}
	// проверка на общую сумму
	outAmount := t.Outputs[0].State.Amount + t.Outputs[1].State.Amount
	inAmount := t.inAmount()
	if outAmount != inAmount {
		fmt.Errorf("Do not match the amount of inputs (%d) and outputs (%d).", inAmount, outAmount)
	}
	// проверка на комиссионные брокеру
	feeAmount := bv.fee.Count(outAmount)
	if feeAmount < 0 {
		fmt.Errorf("Error calculating transaction fees. (%d).", feeAmount)
	}
	// feeAmount := int(fee.Count(uint64(outAmount))) //t.toFee(outAmount, fee)
	if brokerBlock.State.Amount != feeAmount {
		fmt.Errorf("The expected brokerage fee is %d, and %d is indicated.", feeAmount, brokerBlock.State.Amount)
	}
	// проверка суммы получателю не делается, она косвенно в проверках на сумму и на комиссионные
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
