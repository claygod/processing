package entities

// Processing
// TransactionValidator (transfer)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
)

/*
checkTransfer - проверка заполненной транзакции 'Transfer'.
*/
func (bv *TransactionValidator) checkTransfer(t *Transaction) error {
	// проверка хэша, ключей и наличия аккаунтов брокера и инициатора
	if err := bv.validateSignature(t); err != nil {
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
	if err := bv.checkTransferInputs(t); err != nil {
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
	inAmount := 0
	for _, b := range t.Inputs {
		inAmount += b.State.Amount
	}
	// inAmount := bv.inAmount(t)
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
