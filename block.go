package processing

// Processing
// Block
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"crypto/sha256"
	"strconv"
)

/*
Block - contains a resource and restrictions on its use.
*/
type Block struct {
	ResourceId       int
	Amount           int
	OwnerAddress     string
	ResizePermit     bool
	ExhangeLimitTime int64
	ExhangeLimitAdds string
	ExhangeResources string
	ParentBlocks     [][32]byte
	BlockHash        [32]byte
	Signature        []byte
	// Spent        bool
}

type ExhangeConditions struct {
	From         int64
	To           int64
	Resources    []*ExhangeResource
	Addresses    []string
	ResizePermit bool
}
type ExhangeResource struct {
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
CalculateHash - calculation of the hash over all fields of the structure.
ToDo: replace the string with the byte array.
*/
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
