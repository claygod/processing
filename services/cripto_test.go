package services

// Processing
// Cripto test
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"fmt"
	"os"
	// "strconv"
	//"testing"
	//"github.com/claygod/processing/entities"
)

/*
func TestCriptoNew(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", pair.decoded, string(got), pair.encoded)
	c, err := NewCrypto()
	if err != nil {
		t.Error("Error  - create new crypto.")
	}
	//fmt.Println("pvtkey: ", c.pvtKey.key())
	//fmt.Println("pubkey: ", c.pubKey)
	fmt.Println("address: ", c.address)
}

func TestCriptoGenArrJson(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", pair.decoded, string(got), pair.encoded)
	if _, err := forTestGenTempJson("./auths.json", 5); err != nil {
		t.Error(err)
	} else {
		forTestDelTempJson("./auths.json")
	}

}

func forTestGenTempJson(path string, num int) (map[string]*Crypto, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	re := make(map[string]*Crypto)
	out := ""

	for i := 0; i < num; i++ {
		//iStr := strconv.Itoa(i)
		cr, _ := NewCrypto()
		re[cr.address] = cr
		pubKeyStr := cr.b58.Encode(cr.pubKey) //byteKeyToString(cr.pubKey) //b58.Decode(cr.pubKey)

		//fmt.Println(pubKeyStr)
		//fmt.Println(cr.pubKey)
		//fmt.Println(cr.b58.Decode(pubKeyStr))
		//fmt.Println("#")

		str := fmt.Sprintf(`
	{
		"pub_key": "%s",
		"url": "http://localhost/%d",
		"groups_list": ["gr%d", "gr%d"]
	}`, pubKeyStr, i, i, i+1)
		if i < num-1 {
			str += ","
		}
		out += str
	}
	file.WriteString("[")
	file.WriteString(out)
	file.WriteString("\n]")
	return re, nil
}
*/
func forTestDelTempJson(path string) {
	os.Remove(path)
}
