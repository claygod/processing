package processing

// Processing
// Scheme
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
	"fmt"
)

/*
Scheme - message distribution scheme.
*/
type Scheme struct {
	authsCount   int
	myPosition   int
	mailingWidth int
}

/*
NewScheme - create new Scheme.
*/
func NewScheme(authsCount int, myPosition int, mailingWidth int) (*Scheme, error) {
	if authsCount <= myPosition {
		return nil, errors.New("The position is outside the list")
	}
	return &Scheme{
		authsCount:   authsCount,
		myPosition:   myPosition,
		mailingWidth: mailingWidth,
	}, nil
}

func (s *Scheme) GetNumsToSend(offset int) ([]int, error) {
	if offset >= s.authsCount {
		return nil, fmt.Errorf("Offset %d more authorities available (%d)", offset, s.authsCount)
	}

	nums := make([]int, 0, s.mailingWidth)
	position := s.myPosition + offset
	if position >= s.authsCount {
		position -= s.authsCount
	}
	base := (position + 1) * s.mailingWidth

	if base < s.authsCount {
		for i := base; i < base+s.mailingWidth; i++ {
			if i < s.authsCount {
				n := i // + offset
				if n > s.authsCount {
					n -= s.authsCount
				}
				nums = append(nums, n) // !!
			}
		}
	}

	return nums, nil
}
