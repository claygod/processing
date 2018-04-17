package entities

// Processing
// Fee
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Fee - commission calculation.
	measure - how much will be one part of the resulting number
	tax - payment amount from measure
*/
type Fee struct {
	measure int
	tax     int
}

/*
NewFee - create new Fee.
*/
func NewFee(measure int, tax int) *Fee {
	if measure < 100 || tax < 1 {
		return nil
	}
	f := &Fee{
		measure: measure,
		tax:     tax,
	}
	return f
}

/*
Count - calculation of fees.
In case of an error, a negative number is returned.
*/
func (f *Fee) Count(amount int) int {
	if amount < 0 {
		return amount
	}
	res := amount * f.tax / f.measure
	if res == 0 {
		res = 1
	}
	return res
}
