package entities

// Processing
// Block
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"crypto/sha256"
	"encoding/json"
	"sort"
)

/*
Block - contains a resource and restrictions on its use.
Только одна Condition для того, чтобы брокер/агент мог заблокировать
покупаемый ресурс, чтобы гонка за этот ресурс была минимальной.

	Transfer/перевод
Два типа ресурсов (входов), один только у чужого, и НЕ имеет Condition,
второй только у покупателя. Выходов три, один для получателя перевода, один для
комиссии брокера, один для "сдачи" (неиспользованного остатка).

	Сплит/объединение
Все ресурсы (входы) одинаковые и владелец один, сумма входов и выходов одинаковая
(учитывая и выход для брокера).

	Обмен/сделка
Два типа ресурсов, один только у чужого, и имеет Condition,
второй только у покупателя. Сумма ресурсов и покупателя достаточна для
выполнения условия продавца плюс комиссия в валюте покупателя.

Комиссия в размере процента (малого) но не меньше 1 единицы.
*/
type Block struct {
	//ResourceId   int
	//Amount       int
	OwnerAddress string
	State        ExhangeBlock // параметры блока (ресурс и его величина)
	Condition    ExhangeBlock
	ParentBlocks [][]byte
	R            []byte
	S            []byte
	Hash         []byte
}

/*
type ExhangeCondition struct {
	ResourceId   int
	Amount       int
	LimitTime    int64
	LimitAddress string
}
*/

type ExhangeBlock struct {
	ResourceId int
	Amount     int
}

/*
NewBlock - create new Block.

func NewBlock(authsCount int, myPosition int, mailingWidth int) (*Block, error) {
	//if authsCount <= myPosition {
	//	return nil, errors.New("The position is outside the list")
	//}
	return &Block{}, nil
}
*/

/*
marshalling - preparation of data for hashing
*/
func (b *Block) marshalling() ([]byte, error) {
	b.R = []byte{}
	b.S = []byte{}
	b.Hash = []byte{}
	b.sortParentBlocks()

	nb, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return nb, nil
}

/*
sortParentBlocks - locks must be sorted for determinism.
*/
func (b *Block) sortParentBlocks() {
	strs := make([]string, len(b.ParentBlocks))
	for k, v := range b.ParentBlocks {
		strs[k] = string(v)
	}
	sort.Strings(strs)
	for k, v := range strs {
		b.ParentBlocks[k] = []byte(v)
	}

}

/*
CalculateHash - calculation of the hash over all fields of the structure.
ToDo: replace the string with the byte array.

func (b *Block) CalculateHash() {
	var str string
	str += strconv.Itoa(b.ResourceId)
	str += strconv.Itoa(b.Amount)
	str += b.OwnerAddress
	str += string(b.BlockHash[:])
	str += strconv.FormatBool(b.ResizePermit)
	str += strconv.Itoa(int(b.ExhangeLimitTime))
	str += b.ExhangeLimitAdds
	for _, v := range b.ParentBlocks {
		str += string(v[:])
	}
	b.BlockHash = sha256.Sum256([]byte(str))
}
*/
