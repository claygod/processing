package processing

// Processing
// Sender
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
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

// send - sending request
func (s *Sender) send(rh *ReqHelper) error {
	rh.status = StatusOk
	if err := s.fillStamp(&rh.req.Stamp, rh.req); err != nil {
		rh.err = err
		rh.status = ErrorFillStamp
		return err
	}

	client, r, err := s.newHttp(rh)
	if err != nil {
		rh.err = err
		rh.status = ErrorNewHttp
		return err
	}

	resp, err := client.Do(r)
	if err != nil {
		err2 := fmt.Errorf("Response status: `%s`. %v", resp.Status, err)
		rh.err = err2
		rh.status = resp.StatusCode
		return err2
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&rh.res)
	if err != nil {
		err2 := fmt.Errorf("Response status: `%s`. %v", resp.Status, err)
		rh.err = err2
		rh.status = resp.StatusCode
		return err2
	}

	if !s.checkStamp(rh.res, rh.to.PubKey) {
		rh.status = ErrorAnswerVerification
		return errors.New("The answer did not pass the verification.")
	}

	return nil
}

func (s *Sender) newHttp(rh *ReqHelper) (*http.Client, *http.Request, error) {
	qJson, err := json.Marshal(rh.req)
	if err != nil {
		return nil, nil, err
	}

	r, err := http.NewRequest(rh.method, rh.to.Url, bytes.NewBuffer(qJson))
	if err != nil {
		return nil, nil, err
	}

	r.Header.Set("Content-Type", "application/json")
	return &http.Client{Timeout: timeDurationTimeout}, r, nil
}

func (s *Sender) fillStamp(st *Stamp, msg *Message) error {
	st.From = s.crypto.address
	rr, ss, err := s.crypto.sign([]byte(msg.dataForVerification()))
	if err != nil {
		return err
	}
	st.R10 = rr.String()
	st.S10 = ss.String()
	return nil
}

func (s *Sender) checkStamp(msg *Message, pubKey string) bool {
	rr, ok := new(big.Int).SetString(msg.Stamp.R10, 10)
	if !ok {
		return false
	}
	ss, ok := new(big.Int).SetString(msg.Stamp.S10, 10)
	if !ok {
		return false
	}

	return s.crypto.verify(
		[]byte(msg.dataForVerification()),
		[]byte(pubKey),
		rr, ss)
}

type ReqHelper struct {
	send   *Sender
	to     *Authority
	method string
	req    *Message
	res    *Message
	err    error
	status int
}

func NewReqHelper(s *Sender) *ReqHelper {
	return &ReqHelper{
		send: s,
		req:  NewMessage(),
		res:  NewMessage(), // ToDo: To initiate or not?
		err:  nil,
	}
}

func (r *ReqHelper) To(a *Authority) *ReqHelper {
	r.to = a
	return r
}

func (r *ReqHelper) Method(method string) *ReqHelper {
	r.method = method
	return r
}

func (r *ReqHelper) Event(event int) *ReqHelper {
	r.req.Event = event
	return r
}

func (r *ReqHelper) Context(ctx map[string]string) *ReqHelper {
	r.req.Context = ctx
	return r
}

func (r *ReqHelper) Send() (*ReqHelper, error) {
	err := r.send.send(r)
	return r, err
}

func NewMessage() *Message {
	return &Message{
		Context: make(map[string]string),
	}
}

type Message struct {
	Stamp   Stamp             `json:"stamp"`
	Event   int               `json:"event"`
	Context map[string]string `json:"context"`
}

func (r *Message) dataForVerification() string {
	data := strconv.FormatInt(r.Stamp.EntryTime, 10)
	data += strconv.FormatInt(r.Stamp.ExitTime, 10)
	for k, v := range r.Context {
		data = data + k + v
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
