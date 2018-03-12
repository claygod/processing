package processing

// Processing
// Sender
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const timeDurationTimeout time.Duration = 10 * time.Second

/*
Sender - sending requests and receiving replies.
*/
type Sender struct {
	my *Authority
}

/*
NewSender - create new Sender.
*/
func NewSender() *Sender {
	r := &Sender{}
	return r
}

/*
// Verification - проверка на доступ с получением userId
func (a *Sender) Verification(sessionId string, cabinetId int) (bool, int) {
qJson := a.getJsonSSrequest(sessionId, cabinetId)
if qJson == nil {
return false, -1
}
userId, sCode := a.sendRequest(qJson)
if sCode == 200 {
return true, userId
}
return false, userId
}

// getJsonSSrequest - подготовка запроса в json формате
func (a *Sender) getJsonSSrequest(structRequest interface{}) []byte {
	qJson, err := json.Marshal(structRequest)
	if err != nil {
		//lgMsg := &logMessage{"error", ERRORsendHTTPrequest, Fields{"error_context": err.Error()}}
		//a.logger <- lgMsg
		return nil
	}
	return qJson
}
*/

func (s *Sender) Request() *Request {
	return &Request{send: s}
}

// send - sending request
func (s *Sender) send(req *Request) (interface{}, error) {
	client := &http.Client{Timeout: timeDurationTimeout}

	qJson, err := json.Marshal(req.structRequest)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(req.method, req.url, bytes.NewBuffer(qJson))
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&req.structResponse)
	if err != nil {
		return nil, err
	}

	return req.structResponse, nil
}

type Request struct {
	send           *Sender
	method         string
	url            string
	structRequest  interface{}
	structResponse interface{}
}

func (r *Request) Method(method string) *Request {
	r.method = method
	return r
}

func (r *Request) Url(url string) *Request {
	r.url = url
	return r
}

func (r *Request) StructRequest(strct interface{}) *Request {
	r.structRequest = strct
	return r
}

func (r *Request) StructResponse(strct interface{}) *Request {
	r.structResponse = strct
	return r
}

func (r *Request) Send() (interface{}, error) {
	return r.send.send(r)
}
