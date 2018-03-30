package processing

// Processing
// Node
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
)

/*
Node - complete network node.
*/
type Node struct {
	my          *Authority
	authorities map[string]*Authority
	ids         map[string]*Id
	flags       map[string]*int32
	cripto      *Crypto
}

/*
NewNode - create new Node.
*/
func NewNode(address string, path string) (*Node, error) {
	cr, err := NewCrypto()
	if err != nil {
		return nil, err
	}

	n := &Node{
		authorities: make(map[string]*Authority),
		ids:         make(map[string]*Id),
		cripto:      cr,
	}

	if err := n.loadAuthList(path); err != nil {
		return nil, err
	}
	my, ok := n.authorities[address]

	if !ok {
		/*
			for k, v := range n.authorities {
				fmt.Println("? ", k)
				fmt.Println("! ", v)
			}
		*/
		return nil, errors.New(fmt.Sprintf("The key (address) %s is not found in the list", address))
	}
	n.my = my
	return n, nil
}

/*
loadAuthList - load authorities list from file.
*/
func (n *Node) loadAuthList(path string) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var authSlice []Authority
	if err := json.Unmarshal(bs, &authSlice); err != nil {
		return err
	}

	for _, a := range authSlice {
		if a.Url != "" && a.PubKey != "" {

			urlNet, err := url.ParseRequestURI(a.Url)
			if err != nil {
				continue // ToDo: log OR return err
			}
			a.urlNet = urlNet
			n.authorities[a.PubKey] = &a
		}
	}
	return nil
}

/*
worker - universal performer in cycles.

func (n *Node) worker(flag *int32, f func(*int32)) {
	// atomic.StoreInt32(flag, flagWork)
	for {
		f(flag)
		if atomic.LoadInt32(flag) == flagStop {
			return
		}
	}
}
*/
/*
worker - universal performer in cycles with channel.

func (n *Node) workerChan(flag *int32, f func(*int32, chan interface{}), ch chan interface{}) {
	// atomic.StoreInt32(flag, flagWork)
	for {
		f(flag, ch)
		if atomic.LoadInt32(flag) == flagStop {
			return
		}
	}
}
*/
