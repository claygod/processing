package processing

// Processing
// Scheme
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
	"fmt"
	"sync"
)

/*
Scheme - message distribution scheme.
*/
type Scheme struct {
	sync.RWMutex
	authsCount   int
	myPosition   int
	mailingWidth int
	list         []*Authority
}

/*
NewScheme - create new Scheme.
*/
func NewScheme(auths []*Authority, my *Authority, mailingWidth int) (*Scheme, error) {
	list := make([]*Authority, len(auths)*2)
	copy(list, auths)
	copy(list, auths)

	p := -1
	for k, a := range list {
		if a == my {
			p = k
			break
		}
	}
	if p == -1 {
		return nil, errors.New("Your own link was not found")
	}

	return &Scheme{
		authsCount:   len(auths),
		myPosition:   p,
		mailingWidth: mailingWidth,
		list:         list,
	}, nil
}

func (s *Scheme) GetListToSend(offset int, step int) ([]*Authority, error) {
	if offset >= s.authsCount {
		return nil, fmt.Errorf("Offset %d more authorities available (%d)", offset, s.authsCount)
	}

	auths := make([]*Authority, 0, s.mailingWidth)
	position := s.myPosition + offset
	if position >= s.authsCount {
		position -= s.authsCount
	}
	base := (position + 1) * s.mailingWidth

	if base < s.authsCount {
		for i := base; i < base+s.mailingWidth; i++ {
			if i < s.authsCount {
				auths = append(auths, s.list[i+offset]) // !!
			}
		}
	}

	return auths, nil
}
