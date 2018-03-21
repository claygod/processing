package processing

// Processing
// Sender
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"bytes"
	"encoding/json"
	// "errors"
	// "fmt"
	// "math/big"
	"net/http"
	// "net/url"
	"sync"
	//"sync/atomic"
	"time"
	"unsafe"
)

/*
Sender - sending requests and receiving replies.
*/
type Sender struct {
	// my     *Authority
	crypto *Crypto
}

/*
NewSender - create new Sender.
*/
func NewSender(crypto *Crypto) *Sender {
	r := &Sender{crypto: crypto}
	return r
}

func (s *Sender) Request() *ReqHelper {
	return NewReqHelper(s)
}

func (s *Sender) Broadcast(rh *ReqHelper) (*ReqHelper, error) {

	rh.status = StatusOk
	if err := s.fillStamp(&rh.reqMsg.Stamp, rh.reqMsg); err != nil {
		rh.err = err
		rh.status = ErrorFillStamp
		return rh, err
	}

	list := make(map[string]*Authority)

	// make total list
	for _, g := range rh.groups {
		g.Range(func(k, v interface{}) bool {
			ki := k.(string)
			list[ki] = v.(*Authority)
			return true // if false, Range stops
		})
	}

	var wg sync.WaitGroup
	for _, a := range list {
		rh.to = a
		rh.reqHttp.URL = a.urlNet
		wg.Add(1)
		go s.send7(rh, &wg)
	}
	wg.Wait()
	return rh, nil
}

func (s *Sender) send7(rh *ReqHelper, wg *sync.WaitGroup) int {
	defer wg.Done()
	// rh.status = StatusOk
	client := &http.Client{Timeout: timeDurationTimeout}

	resp, err := client.Do(rh.reqHttp)
	if err == nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode == StatusOk {
		rh.listOk.Store(rh.to.Id.address, rh.to)
	} else {
		rh.listErr.Store(rh.to.Id.address, rh.to)
	}

	return resp.StatusCode
}

func (s *Sender) fillStamp(st *Stamp, msg *Message) error {
	st.From = s.crypto.address
	st.ExitTime = time.Now().Unix()
	rr, ss, err := s.crypto.sign([]byte(msg.dataForVerification()))
	if err != nil {
		return err
	}
	st.R10 = rr.String()
	st.S10 = ss.String()
	return nil
}

type ReqHelper struct {
	send    *Sender
	to      *Authority
	method  string
	reqMsg  *Message
	reqJson []byte
	reqHttp *http.Request
	// res     *Message // ToDo: del
	err     error
	status  int
	groups  []*Group
	listOk  *sync.Map
	listErr *sync.Map
	counter *int32
}

func NewReqHelper(s *Sender) *ReqHelper {
	return &ReqHelper{
		send:   s,
		reqMsg: NewMessage(),
		// res:    NewMessage(), // ToDo: To initiate or not?
		err:    nil,
		groups: make([]*Group, 0, 1),
	}
}

func (r *ReqHelper) msgToJson() error {
	b, err := json.Marshal(r.reqMsg)
	if err != nil {
		return err
	}
	r.reqJson = b
	return nil
}

func (r *ReqHelper) To(a *Authority) *ReqHelper {
	r.to = a
	return r
}

func (r *ReqHelper) Method(method string) *ReqHelper {
	r.method = method
	return r
}

func (rh *ReqHelper) newHttp() (*http.Request, error) {
	r, err := http.NewRequest(rh.method, rh.to.Url, bytes.NewBuffer(rh.reqJson))
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/json")
	rh.reqHttp = r

	return r, nil
}

func (r *ReqHelper) Make() (*ReqHelper, error) {
	if err := r.msgToJson(); err != nil {
		return r, err
	}

	reqHttp, err := r.newHttp()
	if err != nil {
		r.err = err
		r.status = ErrorNewHttp
		return r, err
	}
	r.reqHttp = reqHttp
	return r, err
}

func NewMessage() *Message {
	return &Message{
		Context: make([][]byte, 0),
	}
}

type Message struct {
	Stamp Stamp `json:"stamp"`
	// Event   int               `json:"event"`
	Context [][]byte `json:"context"`
}

func (r *Message) dataForVerification() []byte {
	data := (*(*[8]byte)(unsafe.Pointer(&r.Stamp.EntryTime)))[:]
	data = append(data, (*(*[8]byte)(unsafe.Pointer(&r.Stamp.ExitTime)))[:]...)
	for _, v := range r.Context {
		data = append(data, v...)
	}
	return data
}

type Stamp struct {
	From      string `json:"from"`
	EntryTime int64  `json:"entry"`
	ExitTime  int64  `json:"exit"`
	R10       string `json:"r10"`
	S10       string `json:"s10"`
}
