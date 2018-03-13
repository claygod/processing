package processing

// Processing
// Node
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
	// "sync"
	"encoding/base64"
	"sync/atomic"
	"time"
)

const (
	flagStop int32 = iota
	flagWork
)

const defaultAuthoritiesListPath string = "./authorities.json"
const timePauseWorkerAuthStatus time.Duration = 100 * time.Millisecond

/*
Node - complete network node.
*/
type Node struct {
	my          *Authority
	authorities map[string]*Group
	auths       *Group
	accounts    map[string]*Account
	flags       map[string]*int32
}

/*
NewNode - create new Node.
*/
func NewNode(address string, path string) (*Node, error) {
	n := &Node{
		authorities: make(map[string]*Group),
		accounts:    make(map[string]*Account),
		auths:       NewGroup(),
	}
	if err := n.loadAuthList(path); err != nil {
		return nil, err
	}
	//my, ok := authorities[id] // is there any node in the list
	my, ok := n.auths.Load(address)
	if !ok {
		return nil, errors.New(fmt.Sprintf("The key (address) %s is not found in the list", address))
	}
	n.my = my.(*Authority)
	//n.authorities = authorities

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
		if a.Url != "" &&
			a.PubKey64 != "" &&
			len(a.Groups) != 0 { // Filtering out incorrect entries
			pubKey, err := base64.StdEncoding.DecodeString(a.PubKey64)
			if err != nil {
				continue
			}
			address := PubKeyToAddress(pubKey)
			// fmt.Println("===", address)
			n.auths.Store(address, &a)
			for _, gName := range a.Groups {
				g, ok := n.authorities[gName]
				if !ok {
					g = NewGroup()
					n.authorities[gName] = g
				}
				g.Store(address, &a)
			}
		}
	}
	return nil
}

/*
worker - universal performer in cycles.
*/
func (n *Node) worker(flag *int32, f func(*int32)) {
	// atomic.StoreInt32(flag, flagWork)
	for {
		f(flag)
		if atomic.LoadInt32(flag) == flagStop {
			return
		}
	}
}

/*
worker - universal performer in cycles with channel.
*/
func (n *Node) workerChan(flag *int32, f func(*int32, chan interface{}), ch chan interface{}) {
	// atomic.StoreInt32(flag, flagWork)
	for {
		f(flag, ch)
		if atomic.LoadInt32(flag) == flagStop {
			return
		}
	}
}

/*
checkAuthStatus - check authorities status.
*/
func (n *Node) checkAuthStatus(flag *int32) {
	n.auths.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		a := v.(*Authority)
		a.CheckStatus()
		time.Sleep(timePauseWorkerAuthStatus)
		runtime.Gosched()
		if atomic.LoadInt32(flag) == flagStop {
			return false
		}
		return true // if false, Range stops
	})
}

/*
type Authorities struct {
	mx sync.RWMutex
	m  map[string]*Authority
}

func (c *Authorities) Load(key string) (*Authority, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *Authorities) Store(key string, value *Authority) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	c.m[key] = value
}
*/
