package scripts

// Processing
// TransactionInteractor
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"github.com/claygod/processing/domain"
	// "github.com/claygod/processing/entities"
)

type TransactionInteractor struct {
	Chain            *domain.Chain
	BlockRepo        domain.BlockRepository
	Consensus        domain.Consensus
	TransactorRepo   domain.TransactorRepository
	WaitTransactions domain.WaitTransaction
}

func NewTransactionInteractor(
	ch *domain.Chain,
	br domain.BlockRepository,
	cs domain.Consensus,
	tr domain.TransactorRepository,
	wt domain.WaitTransaction,
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
	if err := tr.Prepare(); err != nil {
		t.TransactorRepo.Delete(tn.Hash)
		return err
	}
	if err := t.WaitTransactions.Write(tn); err != nil {
		t.TransactorRepo.Delete(tn.Hash)
		return err
	}
	return nil
}

func (t *TransactionInteractor) AddOpinion(hash string, ok bool) {

	switch t.Consensus.Confirm(hash, ok) {
	case domain.ConsensusStateMissing:
		//
	case domain.ConsensusStateFills:
		//
	case domain.ConsensusStatePositive:
		t.ExecuteTransaction(hash)
	case domain.ConsensusStateNegative:
		t.RollbackTransaction(hash)
	}
}

func (t *TransactionInteractor) ExecuteTransaction(hash string) error {
	tr, err := t.TransactorRepo.Read(hash)
	if err != nil {
		return err
	}
	if err := tr.Execute(); err != nil {
		t.TransactorRepo.Delete(hash)
		return err
	}
	t.TransactorRepo.Delete(hash)
	tn, err := t.WaitTransactions.Read(hash)
	if err != nil {
		return err // ToDo: этот вариант недопустим, если есть транзактор, должна быть и транзакция
	}
	t.Chain.AddTransaction(tn)
	return nil
}

func (t *TransactionInteractor) RollbackTransaction(hash string) error {
	tr, err := t.TransactorRepo.Read(hash)
	if err != nil {
		return err
	}
	if err := tr.Rollback(); err != nil {
		t.TransactorRepo.Delete(hash)
		return err
	}
	t.TransactorRepo.Delete(hash)

	if err := t.WaitTransactions.Delete(hash); err != nil {
		return err // ToDo: этот вариант недопустим, если есть транзактор, должна быть и транзакция
	}
	//t.Chain.AddTransaction(tn)
	return nil
}
