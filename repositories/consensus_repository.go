package repositories

// Processing
// Consensus repository (implementation)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/claygod/processing/domain"
)

const (
	flagStop int32 = iota
	flagWork
)

/*
ConsensusRepository - consensus store and vote.
Репозиторий не следит за временем голосования, этот функционал лучше
передать в репо транзакторов или ещё куда, т.е. может быть так:
консенсус достигнут, но транзактора нет, значит и ничего консенсус не толкнет.
*/
type ConsensusRepository struct {
	quorum int64
	store  [256]*voteStore // 256 arrays to reduce access competitiveness
}

/*
NewConsensusRepository - create new ConsensusRepository.
*/
func NewConsensusRepository(quorum int64) *ConsensusRepository {
	c := &ConsensusRepository{
		quorum: quorum,
	}
	for i := 0; i < 256; i++ {
		c.store[i] = newVoteStore()
	}
	return c
}

func (c *ConsensusRepository) Vote(unit string, key string, opinion bool) (int64, error) {
	var shift byte
	if len(key) > 0 {
		shift = []byte(key)[0]
	}
	yes, no, err := c.store[shift].getVoting(key).vote(unit, opinion)
	if err != nil {
		return domain.ConsensusFills, err
	}
	if yes >= c.quorum {
		return domain.ConsensusPositive, nil
	}
	if no >= c.quorum {
		return domain.ConsensusNegative, nil
	}
	return domain.ConsensusFills, nil
}

func (c *ConsensusRepository) SetQuorum(q int64) {
	atomic.StoreInt64(&c.quorum, q)
}

/*
voteStore - vote substore.
*/
type voteStore struct {
	sync.Mutex
	store map[string]*voting
}

func newVoteStore() *voteStore {
	v := &voteStore{
		store: make(map[string]*voting),
	}
	return v
}

func (v *voteStore) getVoting(key string) *voting {
	v.Lock()
	x, ok := v.store[key]
	if !ok {
		x = newVoting()
		v.store[key] = x
	}
	v.Unlock()
	return x
}

/*
func (v *voteStore) vote(key string, opinion bool) {
	v.Lock()
	a, ok := v.store[key]
	if !ok {
		a = newVoting()
		v.store[key] = a
	}
	v.Unlock()
	// return a
}
*/
/*
voting - vote.
*/
type voting struct {
	// status  int64
	sync.Mutex
	yes      int64
	no       int64
	opinions map[string]bool
}

func newVoting() *voting {
	v := &voting{
		opinions: make(map[string]bool),
	}
	return v
}

func (v *voting) vote(unit string, opinion bool) (int64, int64, error) {
	v.Lock()
	defer v.Unlock()
	if _, ok := v.opinions[unit]; ok {
		return -1, -1, fmt.Errorf("Repeated voting '%s'.", unit)
	}
	if opinion {
		v.yes++
	} else {
		v.no++
	}
	return v.yes, v.no, nil
}
