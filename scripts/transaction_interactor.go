package scripts

// Processing
// TransactionInteractor
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"github.com/claygod/processing/domain"
	// "github.com/claygod/processing/entities"
)

type TransactionInteractor struct {
	My             string
	Chain          *domain.Chain
	BlockRepo      domain.BlockRepository
	Consensus      domain.Consensus
	TransactorRepo domain.TransactorRepository
	Sender         Sender
}

func NewTransactionInteractor(
	my string,
	ch *domain.Chain,
	br domain.BlockRepository,
	cs domain.Consensus,
	tr domain.TransactorRepository,
	sr Sender,
) *TransactionInteractor {

	return &TransactionInteractor{
		My:             my,
		Chain:          ch, //domain.NewChain(br),
		BlockRepo:      br,
		Consensus:      cs,
		TransactorRepo: tr,
		Sender:         sr,
	}
}

/*
AcceptAuthorityTransaction -  подтверждение полученной по сети транзакции (не от клиента).
*/
func (t *TransactionInteractor) AcceptAuthorityTransaction(tn *domain.Transaction) error {
	// ToDo: validation and signature
	opin := domain.NewOpinion(t.My, tn.Hash, false)

	tr, err := t.TransactorRepo.Create(tn)
	if err != nil {
		// тут мнение не отсылаем, т.к. наверно отсылали его раньше
		return err
	}
	if err := tr.Prepare(tn); err != nil {
		t.TransactorRepo.Delete(tn.Hash)
		t.Sender.SendOpinion(opin)
		return err
	}
	opin.Ok = true
	t.Sender.SendOpinion(opin)
	return nil
}

/*
AcceptAuthorityOpinion

Надо определиться с логикой, например, если голосования нет, то ведь и
транзактора и транзакции нет. Хотя они могут появиться чуть позже.
*/
func (t *TransactionInteractor) AcceptAuthorityOpinion(opin *domain.Opinion) {
	// ToDo: validation

	switch t.Consensus.Vote(opin) {
	//case domain.ConsensusStateMissing:
	// вариант, когда такого голосования нет, и нету среди старых
	// Тут возможно создавать новое голосование заново (внутри консенсуса)
	case domain.ConsensusFills:
		// тут по идее ничего не нужно делать т.к. t.Consensus.Confirm уже
		// сделал нужную работу по учёту мнения
		return
	case domain.ConsensusPositive:
		t.executeTransaction(opin.Hash)
	case domain.ConsensusNegative:
		t.rollbackTransaction(opin.Hash)
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
