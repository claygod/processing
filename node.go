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
	"sync"
	"time"
)

const defaultAuthoritiesListPath string = "./authorities.json"
const timePauseWorkerAuthStatus time.Duration = 100 * time.Millisecond

/*
Node - complete network node.
*/
type Node struct {
	my *Authority
	//authorities map[string]*Authority
	auths    sync.Map
	accounts map[string]*Account
}

/*
NewNode - create new Node.
*/
func NewNode(id string, path string) (*Node, error) {
	n := &Node{
		accounts: make(map[string]*Account),
	}
	if err := n.loadAuthList(path); err != nil {
		return nil, err
	}
	//my, ok := authorities[id] // is there any node in the list
	my, ok := n.auths.Load(id)
	if !ok {
		return nil, errors.New(fmt.Sprintf("The key %s is not found in the list", id))
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
	//var authMap = make(map[string]*Authority)

	if err := json.Unmarshal(bs, &authSlice); err != nil {
		return err
	}

	for _, a := range authSlice {
		if a.Id != "" && a.Url != "" { // Filtering out incorrect entries
			//authMap[a.Id] = &a
			n.auths.Store(a.Id, &a)
		}
	}
	return nil
}

/*
workerAuthStatus - load authorities list from file.
*/
func (n *Node) workerAuthStatus() {
	for {
		n.checkAuthStatus()
	}
}

/*
checkAuthStatus - check authorities status.
*/
func (n *Node) checkAuthStatus() {
	n.auths.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		a := v.(*Authority)
		a.CheckStatus()
		time.Sleep(timePauseWorkerAuthStatus)
		runtime.Gosched()
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
