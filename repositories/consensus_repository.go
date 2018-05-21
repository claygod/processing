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
	store  [256]*consensusStore // 256 arrays to reduce access competitiveness
}

/*
NewConsensusRepository - create new ConsensusRepository.
*/
func NewConsensusRepository(quorum int64) *ConsensusRepository {
	c := &ConsensusRepository{
		quorum: quorum,
	}
	for i := 0; i < 256; i++ {
		c.store[i] = newConsensusStore()
	}
	return c
}

/*
func (c *ConsensusRepository) Vote222(unit string, key string, opinion bool) (int64, error) {
	var shift byte
	if len(key) > 0 {
		shift = []byte(key)[0]
	}
	yes, no, err := c.store[shift].getConsensus(key).vote(unit, opinion)
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
*/
func (c *ConsensusRepository) Vote(opin *domain.Opinion) (int64, error) {
	var shift byte
	if len(opin.Hash) > 0 {
		shift = []byte(opin.Hash)[0]
	}
	yes, no, err := c.store[shift].getConsensus(opin.Hash).vote(opin.Unit, opin.Ok)
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
consensusStore - substore.
*/
type consensusStore struct {
	sync.Mutex
	store map[string]*consensus
}

func newConsensusStore() *consensusStore {
	v := &consensusStore{
		store: make(map[string]*consensus),
	}
	return v
}

func (v *consensusStore) getConsensus(key string) *consensus {
	v.Lock()
	c, ok := v.store[key]
	if !ok {
		c = newConsensus()
		v.store[key] = c
	}
	v.Unlock()
	return c
}

/*
consensus - vote.
*/
type consensus struct {
	// status  int64
	sync.Mutex
	yes      int64
	no       int64
	opinions map[string]bool
}

func newConsensus() *consensus {
	v := &consensus{
		opinions: make(map[string]bool),
	}
	return v
}

func (v *consensus) vote(unit string, opinion bool) (int64, int64, error) {
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
