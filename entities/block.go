package entities

// Processing
// Block
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"crypto/sha256"
	"bytes"
	"encoding/gob"
	"encoding/json"
	//"sort"
)

/*
Block - contains a resource and restrictions on its use.
Только одна Condition для того, чтобы брокер/агент мог заблокировать
покупаемый ресурс, чтобы гонка за этот ресурс была минимальной.
Подписывает авторитет!

	Transfer/перевод
Один тип ресурсов (входов), все только свои, и НЕ имеют Condition. Выходов три, один для получателя перевода, один для
комиссии брокера, один для "сдачи" (неиспользованного остатка).

	Merger Сплит/объединение
Все ресурсы (входы) одинаковые и владелец один, и НЕ имеют Condition, сумма входов и выходов одинаковая
(плюс выход для брокера).

	Separation/Разделение
	...

	Обмен/сделка
Два типа ресурсов и два входа, один у чужого и имеет Condition,
второй только у покупателя и без  Condition. Сумма ресурсов и покупателя достаточна для
выполнения условия продавца плюс комиссия в валюте покупателя.

Комиссия в размере процента (малого) но не меньше 1 единицы.
*/
type Block struct {
	Owner     string
	Broker    string
	State     ExhangeBlock // параметры блока (ресурс и его величина)
	Condition ExhangeBlock
	// ParentBlocks [][]byte
	R    []byte
	S    []byte
	Hash string
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
Marshalling - preparation of data for hashing
*/
func (b *Block) Marshalling() (string, error) {
	b.R = []byte{}
	b.S = []byte{}
	b.Hash = ""
	//b.sortParentBlocks()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(b)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

/*
MarshallingJson - preparation of data for hashing
*/
func (b *Block) MarshallingJson() (string, error) {
	b.R = []byte{}
	b.S = []byte{}
	b.Hash = ""
	//b.sortParentBlocks()

	nb, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(nb), nil
}

/*
SetHash - set hash.
*/
func (b *Block) SetHash(hash string) {
	b.Hash = hash
}

/*
sortParentBlocks - locks must be sorted for determinism.

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
*/
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
