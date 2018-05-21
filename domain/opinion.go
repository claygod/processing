package domain

// Processing
// Opinion
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"strconv"
)

/*
Opinion - мнение по поводу опроса.
*/
type Opinion struct {
	Unit      string
	Hash      string
	Ok        bool
	Signature *Signature
}

/*
NewOpinion - create new Opinion.
*/
func NewOpinion(unit string, hash string, ok bool) *Opinion {
	o := &Opinion{
		Unit: unit,
		Hash: hash,
		Ok:   ok,
	}
	return o
}

func (t *Opinion) AddSignature(s *Signature) {
	t.Signature = s
}

func (t *Opinion) GetHash() string {
	return t.Hash + strconv.FormatBool(t.Ok)
}
