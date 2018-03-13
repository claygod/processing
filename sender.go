package processing

// Processing
// Sender
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"bytes"
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"
	"time"
)

const (
	ReqTypeQuestion int = iota
	ReqTypeAnswer
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

func (s *Sender) Request() *ReqHelper {
	return &ReqHelper{send: s}
}

// send - sending request
func (s *Sender) send(req *ReqHelper) (interface{}, error) {
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
		return nil, fmt.Errorf("Response status: `%s`. %v", resp.Status, err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&req.structResponse)
	if err != nil {
		return nil, fmt.Errorf("Response status: `%s`. %v", resp.Status, err)
	}

	return req.structResponse, nil
}

/*
sendJson
Important: the answer must be then closed!
*/
func (s *Sender) sendJson(req *ReqHelper) (*http.Response, error) {
	client := &http.Client{Timeout: timeDurationTimeout}

	r, err := http.NewRequest(req.method, req.url, bytes.NewBuffer(req.buf))
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(r)
	return resp, err
}

type ReqHelper struct {
	send           *Sender
	method         string
	url            string
	buf            []byte
	structRequest  interface{}
	structResponse interface{}
}

func (r *ReqHelper) Method(method string) *ReqHelper {
	r.method = method
	return r
}

func (r *ReqHelper) Url(url string) *ReqHelper {
	r.url = url
	return r
}

func (r *ReqHelper) Buf(buf []byte) *ReqHelper {
	r.buf = buf
	return r
}

func (r *ReqHelper) StructRequest(strct interface{}) *ReqHelper {
	r.structRequest = strct
	return r
}

func (r *ReqHelper) StructResponse(strct interface{}) *ReqHelper {
	r.structResponse = strct
	return r
}

func (r *ReqHelper) Send() (interface{}, error) {
	return r.send.send(r)
}

type Request struct {
	Type        int               `json:"type"` // question or answer
	Method      int               `json:"method"`
	TimeSendReq int64             `json:"timesendreq"`
	TimeMyShift int64             `json:"timemyshift"`
	Context     map[string]string `json:"context"`
}

type Response struct {
	Status      int               `json:"status"`
	TimeSendReq int64             `json:"timesendreq"`
	TimeRecdReq int64             `json:"timerecdreq"`
	TimeSendRes int64             `json:"timesendres"`
	TimeMyShift int64             `json:"timemyshift"`
	Context     map[string]string `json:"context"`
}
