package domain

// Processing
// Opinion
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Opinion - мнение по поводу опроса.
*/
type Opinion struct {
	Unit string
	Hash string
	Ok   bool
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
