package scripts

// Processing
// TransactionInteractor
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"github.com/claygod/processing/domain"
	// "github.com/claygod/processing/entities"
)

type TransactionInteractor struct {
	Chain          *domain.Chain
	BlockRepo      domain.BlockRepository
	Consensus      domain.Consensus
	TransactorRepo domain.TransactorRepository
	//WaitTransactions domain.WaitTransaction
}

func NewTransactionInteractor(
	ch *domain.Chain,
	br domain.BlockRepository,
	cs domain.Consensus,
	tr domain.TransactorRepository,
	//wt domain.WaitTransaction,
) *TransactionInteractor {

	return &TransactionInteractor{
		Chain:          ch, //domain.NewChain(br),
		BlockRepo:      br,
		Consensus:      cs,
		TransactorRepo: tr,
	}
}

func (t *TransactionInteractor) ToConfirm(tn *domain.Transaction) error {
	tr, err := t.TransactorRepo.Create(tn)
	if err != nil {
		return err
	}
	if err := tr.Prepare(tn); err != nil {
		t.TransactorRepo.Delete(tn.Hash)
		return err
	}
	//if err := t.WaitTransactions.Write(tn); err != nil {
	//	t.TransactorRepo.Delete(tn.Hash)
	//	return err
	//}
	return nil
}

/*
AddOpinion

Надо определиться с логикой, например, если голосования нет, то ведь и
транзактора и транзакции нет. Хотя они могут появиться чуть позже.
*/
func (t *TransactionInteractor) AddOpinion(unit string, hash string, ok bool) {

	switch t.Consensus.Vote(unit, hash, ok) {
	//case domain.ConsensusStateMissing:
	// вариант, когда такого голосования нет, и нету среди старых
	// Тут возможно создавать новое голосование заново (внутри консенсуса)
	case domain.ConsensusFills:
		// тут по идее ничего не нужно делать т.к. t.Consensus.Confirm уже
		// сделал нужную работу по учёту мнения
		return
	case domain.ConsensusPositive:
		t.executeTransaction(hash)
	case domain.ConsensusNegative:
		t.rollbackTransaction(hash)
		//case domain.ConsensusStateExpired:
		// вариант с устаревшим голосованием, по которому уже принято решение
		// - просто отбрасываем мнение как не интересующее
		//return
	}
}

func (t *TransactionInteractor) executeTransaction(hash string) error {
	tr, err := t.TransactorRepo.Read(hash)
	if err != nil {
		return err
	}
	trn, err := tr.Finish()
	if err != nil {
		t.TransactorRepo.Delete(hash)
		return err
	}
	t.TransactorRepo.Delete(hash)
	t.Chain.AddTransaction(trn)
	return nil
}

func (t *TransactionInteractor) rollbackTransaction(hash string) error {
	tr, err := t.TransactorRepo.Read(hash)
	if err != nil {
		return err
	}
	if err := tr.Rollback(); err != nil {
		t.TransactorRepo.Delete(hash)
		return err
	}
	t.TransactorRepo.Delete(hash)
	return nil
}
