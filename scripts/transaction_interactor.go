package scripts

// Processing
// Transaction interactor
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"

	"github.com/claygod/processing/entities"
)

/*
TransactionInteractor - transactional use cases.
*/
type TransactionInteractor struct {
	broker    entities.Authority
	encoder   entities.Encoder
	encryptor entities.Encryptor
	fee       int
}

/*
NewTransactionInteractor - create new TransactionInteractor.
*/
func NewTransactionInteractor() *TransactionInteractor { // inputs []entities.Block
	ti := &TransactionInteractor{}
	return ti
}

/*
Transfer - проверяем полученную транзакцию.
*/
func (ti *TransactionInteractor) CheckTransfer(t entities.Transaction) error {

	/*
		if err := ti.checkInputsTransfer(initiator, inputs); err != nil {
			return err
		}

		//t := entities.NewTransaction(initiator, ti.broker.Account.Address)

		if err := ti.addInputs(inputs, t); err != nil {
			return err
		}
	*/
	return nil
}

/*
checkInputsTransfer - проверка входов - должен быть один ресурс у всех входов и один владелец.
*/
func (ti *TransactionInteractor) checkInputsTransfer(initiator string, inputs []entities.Block) error {
	if len(inputs) == 0 {
		fmt.Errorf("No input resources")
	}
	resource := inputs[0].State.ResourceId
	for _, b := range inputs {
		if initiator != b.Owner {
			return fmt.Errorf("Transfer: entries belong to more than one owner - %s, %s...", initiator, b.Owner)
		}
		if resource != b.State.ResourceId {
			return fmt.Errorf("You can only transfer one resource at a time - %s, %s...", resource, b.State.ResourceId)
		}
	}
	return nil
}

/*
checkOutputsTransfer - проверка выходов - должно быть два выхода (один из них брокера), один ресурс.
*/
func (ti *TransactionInteractor) checkOutputsTransfer(initiator string, outputs []entities.Block, inAmount int) error {
	if len(outputs) != 1 {
		fmt.Errorf("Transfer is carried out only on one address.")
	}

	var brokerBlock *entities.Block
	// var recipientBlock *entities.Block
	if outputs[0].Owner == ti.broker.Account.Address && outputs[1].Owner == initiator {
		brokerBlock = &outputs[0]
		//recipientBlock = &outputs[1]
	} else if outputs[0].Owner == ti.broker.Account.Address && outputs[0].Owner == initiator {
		brokerBlock = &outputs[1]
		//recipientBlock = &outputs[0]
	} else {
		fmt.Errorf("Error in outputs.")
	}
	//receiver :=
	fullAmount := outputs[0].State.Amount + outputs[1].State.Amount
	fee := ti.toFee(fullAmount)

	if brokerBlock.State.Amount != fee {
		fmt.Errorf("The expected brokerage fee is %d, and %d is indicated.", fee, brokerBlock.State.Amount)
	}

	if brokerBlock.State.Amount != fee {
		fmt.Errorf("The expected brokerage fee is %d, and %d is indicated.", fee, brokerBlock.State.Amount)
	}

	resource := outputs[0].State.ResourceId
	for _, b := range outputs {
		if initiator != b.Owner {
			return fmt.Errorf("Transfer: entries belong to more than one owner - %s, %s...", initiator, b.Owner)
		}
		if resource != b.State.ResourceId {
			return fmt.Errorf("You can only transfer one resource at a time - %s, %s...", resource, b.State.ResourceId)
		}
	}
	return nil
}

func (ti *TransactionInteractor) toFee(amount int) int {

	return 1
}

/*
addInput - add Input.
*/
func (ti *TransactionInteractor) addInputs(inputs []entities.Block, t *entities.Transaction) error {
	for _, b := range inputs {
		if err := ti.addInput(b, t); err != nil {
			return fmt.Errorf("Hash %s is already in the list of 'Inputs'.", b.Hash)
		}
	}
	return nil
}

/*
addInput - add Input.
*/
func (ti *TransactionInteractor) addInput(b entities.Block, t *entities.Transaction) error {
	conditionCounter := 0
	if len(t.Outputs) > 0 {
		return fmt.Errorf("The outputs are already formed and inputs can not be added.")
	}
	for _, ib := range t.Inputs {
		if ib.Hash == b.Hash {
			return fmt.Errorf("Hash %s is already in the list of 'Inputs'.", ib.Hash)
		}
		if ib.Condition.ResourceId != 0 { // ToDo
			conditionCounter++
		}
	}
	if conditionCounter > 1 {
		return fmt.Errorf("It is impossible to combine several exchanges in one transaction.")
	}
	t.Inputs = append(t.Inputs, b)
	return nil
}
