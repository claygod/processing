package repositories

// Processing
// Consensus repository (implementation)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"sync"
)

/*
ConsensusRepository - consensus store.
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

func (c *ConsensusRepository) Vote(key string, opinion bool) {
	var shift byte
	if len(key) > 0 {
		shift = []byte(key)[0]
	}
	c.store[shift].vote(key, opinion)
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

/*
voting - vote.
*/
type voting struct {
	// status  int64
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

func (v *voting) vote2(key string, opinion bool) {

}
