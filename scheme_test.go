package processing

// Processing
// Scheme test
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "fmt"
	"testing"
)

func TestSchemeGetNumsToSend_Off_0(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", string(got), pair.encoded)
	authsCount := 7
	myPosition := 0
	mailingWidth := 2
	s, err := NewScheme(authsCount, myPosition, mailingWidth)

	if err != nil {
		t.Error(err)
	}
	offset := 0

	listResult, err := s.GetNumsToSend(offset)
	if err != nil {
		t.Error(err)
	}
	listTrue := []int{2, 3}
	if len(listResult) != len(listTrue) {
		t.Errorf("Do not match lengths listResult = %d, listTrue = %d", len(listResult), len(listTrue))
	} else {
		for i, _ := range listTrue {
			if listResult[i] != listTrue[i] {
				t.Errorf("In position [%d] result = %d, want %d", i, listResult[i], listTrue[i])
			}
		}
	}
}

func TestSchemeGet_Count32_Pos5_Width2_Off0(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", string(got), pair.encoded)
	authsCount := 32
	myPosition := 5
	mailingWidth := 2
	s, err := NewScheme(authsCount, myPosition, mailingWidth)

	if err != nil {
		t.Error(err)
	}
	offset := 0

	listResult, err := s.GetNumsToSend(offset)
	if err != nil {
		t.Error(err)
	}
	listTrue := []int{12, 13}
	if len(listResult) != len(listTrue) {
		t.Errorf("Do not match lengths listResult = %d, listTrue = %d", len(listResult), len(listTrue))
	} else {
		for i, _ := range listTrue {
			if listResult[i] != listTrue[i] {
				t.Errorf("In position [%d] result = %d, want %d", i, listResult[i], listTrue[i])
			}
		}
	}
}

func TestSchemeGet_Count32_Pos5_Width2_Off3(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", string(got), pair.encoded)
	authsCount := 32
	myPosition := 5
	mailingWidth := 2
	s, err := NewScheme(authsCount, myPosition, mailingWidth)

	if err != nil {
		t.Error(err)
	}
	offset := 3

	listResult, err := s.GetNumsToSend(offset)
	if err != nil {
		t.Error(err)
	}
	listTrue := []int{18, 19}
	if len(listResult) != len(listTrue) {
		t.Errorf("Do not match lengths listResult = %d, listTrue = %d", len(listResult), len(listTrue))
	} else {
		for i, _ := range listTrue {
			if listResult[i] != listTrue[i] {
				t.Errorf("In position [%d] result = %d, want %d", i, listResult[i], listTrue[i])
			}
		}
	}
}

func TestSchemeGet_Count43_Pos19_Width2_Off0(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", string(got), pair.encoded)
	authsCount := 43
	myPosition := 19
	mailingWidth := 2
	s, err := NewScheme(authsCount, myPosition, mailingWidth)

	if err != nil {
		t.Error(err)
	}
	offset := 0

	listResult, err := s.GetNumsToSend(offset)
	if err != nil {
		t.Error(err)
	}
	listTrue := []int{40, 41}
	if len(listResult) != len(listTrue) {
		t.Errorf("Do not match lengths listResult = %d, listTrue = %d", len(listResult), len(listTrue))
	} else {
		for i, _ := range listTrue {
			if listResult[i] != listTrue[i] {
				t.Errorf("In position [%d] result = %d, want %d", i, listResult[i], listTrue[i])
			}
		}
	}
}

func TestSchemeGet_Count43_Pos19_Width2_Off1(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", string(got), pair.encoded)
	authsCount := 43
	myPosition := 19
	mailingWidth := 2
	s, err := NewScheme(authsCount, myPosition, mailingWidth)

	if err != nil {
		t.Error(err)
	}
	offset := 1

	listResult, err := s.GetNumsToSend(offset)
	if err != nil {
		t.Error(err)
	}
	listTrue := []int{42}
	if len(listResult) != len(listTrue) {
		t.Errorf("Do not match lengths listResult = %d, listTrue = %d", len(listResult), len(listTrue))
	} else {
		for i, _ := range listTrue {
			if listResult[i] != listTrue[i] {
				t.Errorf("In position [%d] result = %d, want %d", i, listResult[i], listTrue[i])
			}
		}
	}
}

func TestSchemeGet_Count21_Pos0_Width16_Off0(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", string(got), pair.encoded)
	authsCount := 21
	myPosition := 0
	mailingWidth := 16
	s, err := NewScheme(authsCount, myPosition, mailingWidth)

	if err != nil {
		t.Error(err)
	}
	offset := 0

	listResult, err := s.GetNumsToSend(offset)
	if err != nil {
		t.Error(err)
	}
	listTrue := []int{16, 17, 18, 19, 20}
	if len(listResult) != len(listTrue) {
		t.Errorf("Do not match lengths listResult = %d, listTrue = %d", len(listResult), len(listTrue))
	} else {
		for i, _ := range listTrue {
			if listResult[i] != listTrue[i] {
				t.Errorf("In position [%d] result = %d, want %d", i, listResult[i], listTrue[i])
			}
		}
	}
}

func TestSchemeGet_Count15_Pos0_Width16_Off0(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", string(got), pair.encoded)
	authsCount := 15
	myPosition := 0
	mailingWidth := 16
	s, err := NewScheme(authsCount, myPosition, mailingWidth)

	if err != nil {
		t.Error(err)
	}
	offset := 0

	listResult, err := s.GetNumsToSend(offset)
	if err != nil {
		t.Error(err)
	}
	listTrue := []int{}
	if len(listResult) != len(listTrue) {
		t.Errorf("Do not match lengths listResult = %d, listTrue = %d", len(listResult), len(listTrue))
	} else {
		for i, _ := range listTrue {
			if listResult[i] != listTrue[i] {
				t.Errorf("In position [%d] result = %d, want %d", i, listResult[i], listTrue[i])
			}
		}
	}

}
