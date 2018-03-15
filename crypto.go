package processing

// Processing
// Crypto
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	//"crypto/hmac"
	"crypto/rand"
	//"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	//"encoding/base64"
	//"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/big"
)

/*
Crypto - work with passwords.
*/
type Crypto struct {
	limit   int64
	pubKey  []byte
	pvtKey  *keyBox
	address string
	b58     *Base58
}

/*
NewCrypto - create new Crypto.
*/
func NewCrypto() (*Crypto, error) {
	c := &Crypto{
		pubKey: make([]byte, 0),
		b58:    NewBase58(),
	}
	pubKey, pvtKey := c.genNewKeys()
	kb, err := newKeyBox(pvtKey)
	if err != nil {
		fmt.Println("Level keybox")
		return nil, err
	}
	c.pubKey = pubKey
	c.pvtKey = kb
	sh1 := sha256.Sum256(pubKey)
	sh := sha256.Sum256(sh1[:])
	c.address = c.b58.Encode(sh[0:32])
	return c, nil
}

func (c *Crypto) PubKeyToAddress(pubKey []byte) string {
	sh1 := sha256.Sum256(pubKey)
	sh2 := sha256.Sum256(sh1[:])
	return c.b58.Encode(sh2[0:32])
}

func (c *Crypto) genNewKeys() ([]byte, []byte) {
	pvt_key, _ := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	pvt_key_bytes, _ := x509.MarshalECPrivateKey(pvt_key)
	pub_key_bytes, _ := x509.MarshalPKIXPublicKey(&pvt_key.PublicKey)
	return pvt_key_bytes, pub_key_bytes
}

func (c *Crypto) sign(str []byte) (*big.Int, *big.Int, error) {
	zero := big.NewInt(0)
	pvt_key, err := x509.ParseECPrivateKey(c.pvtKey.key())
	if err != nil {
		return zero, zero, err
	}

	r, s, err := ecdsa.Sign(rand.Reader, pvt_key, str)
	if err != nil {
		return zero, zero, err
	}
	return r, s, nil
}

func (c *Crypto) verify(str []byte, pub_key_bytes []byte, r *big.Int, s *big.Int) bool {
	pub_key, err := x509.ParsePKIXPublicKey(pub_key_bytes)
	if err != nil {
		return false
	}

	switch pub_key := pub_key.(type) {
	case *ecdsa.PublicKey:
		return ecdsa.Verify(pub_key, str, r, s)
	default:
		return false
	}
}

/*
keyBox - to store the password.
*/
type keyBox struct {
	rndKey []byte // for safety, keep separate
	pass   []byte
}

func newKeyBox(pass []byte) (*keyBox, error) {
	key := make([]byte, 16)
	rand.Read(key)

	k := &keyBox{}
	k.rndKey = key
	cpass, err := k.encrypt(pass, key)
	if err != nil {
		return nil, err
	}
	k.pass = cpass
	if _, err := k.decrypt(); err != nil { // reverse check
		return nil, err
	}
	return k, nil
}

func (k *keyBox) key() []byte {
	pass, _ := k.decrypt()
	return pass
}

func (k *keyBox) encrypt(ptxt []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, ptxt, nil), nil
}

func (k *keyBox) decrypt() ([]byte, error) {
	c, err := aes.NewCipher(k.rndKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(k.pass) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := k.pass[:nonceSize], k.pass[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
