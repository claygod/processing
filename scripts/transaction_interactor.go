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
	Send           Sender
	Verify         Verificator
	// Encrypt        entities.Encryptor
	Sign domain.Signaturer
}

func NewTransactionInteractor(
	my string,
	ch *domain.Chain,
	br domain.BlockRepository,
	cs domain.Consensus,
	tr domain.TransactorRepository,
	sr Sender,
	vr Verificator,
	// en entities.Encryptor,
	si domain.Signaturer,
) *TransactionInteractor {

	return &TransactionInteractor{
		My:             my,
		Chain:          ch, //domain.NewChain(br),
		BlockRepo:      br,
		Consensus:      cs,
		TransactorRepo: tr,
		Send:           sr,
		Verify:         vr,
		// Encrypt:        en,
		Sign: si,
	}
}

/*
AcceptClientTransaction -  подтверждение полученной от клиента транзакции (не по сети).
*/
func (t *TransactionInteractor) AcceptClientTransaction(tn *domain.Transaction) error {
	tnHash, err := tn.GetHash()
	if err != nil {
		return err
	}
	tn.AddSignature(t.Sign.Make(t.My, tnHash))
	if err := t.Verify.Transaction(tn); err != nil {
		return err
	}

	tr, err := t.TransactorRepo.Create(tn)
	if err != nil {
		return err
	}

	if err := tr.Prepare(tn); err != nil {
		t.TransactorRepo.Delete(tn.Hash)
		return err
	}
	t.Send.Transaction(tn)
	// поскольку мы рассылаем, значит наше мнение положительное
	// и мы его сразу учитываем (принимаем)
	opin := domain.NewOpinion(t.My, tn.Hash, true)
	opin.AddSignature(t.Sign.Make(t.My, opin.GetHash()))
	t.AcceptAuthorityOpinion(opin)
	return nil
}

/*
AcceptAuthorityTransaction -  подтверждение полученной по сети транзакции (не от клиента).
*/
func (t *TransactionInteractor) AcceptAuthorityTransaction(tn *domain.Transaction) error {
	// ToDo: validation and signature
	if err := t.Verify.Transaction(tn); err != nil {
		return err
	}

	tr, err := t.TransactorRepo.Create(tn)
	if err != nil {
		// тут мнение не отсылаем, т.к. наверно отсылали его раньше
		return err
	}

	opin := domain.NewOpinion(t.My, tn.Hash, false)

	if err := tr.Prepare(tn); err != nil {
		t.TransactorRepo.Delete(tn.Hash)
		opin.AddSignature(t.Sign.Make(t.My, opin.GetHash()))
		t.Send.Opinion(opin)
		return err
	}
	opin.Ok = true
	opin.AddSignature(t.Sign.Make(t.My, opin.GetHash()))
	t.Send.Opinion(opin)
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
