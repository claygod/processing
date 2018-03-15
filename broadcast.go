package processing

// Processing
// Broadcast
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import "time"

/*
Broadcast - broad distribution of messages.
*/
type Broadcast struct {
	send *Sender
}

/*
NewBroadcast - create new Broadcast.
*/
func NewBroadcast() *Broadcast {
	return &Broadcast{}
}

func (b *Broadcast) Dispatch(groups []*Group, rh *ReqHelper) (map[string]*ReqHelper, map[string]*ReqHelper) {
	list := make(map[string]*Authority)
	listOk := make(map[string]*ReqHelper)
	listErr := make(map[string]*ReqHelper)
	// make total list
	for _, g := range groups {
		g.Range(func(k, v interface{}) bool {
			ki := k.(string)
			list[ki] = v.(*Authority)
			return true // if false, Range stops
		})
	}

	for k, a := range list {
		// ToDo: менять ли r,s для каждого ? Нет, накладно пожалуй...
		rh.req.Stamp.ExitTime = time.Now().Unix()
		rh.to = a
		rh, err := rh.Send() // b.send.send(req)
		if err != nil {
			listErr[k] = rh
		} else {
			listOk[k] = rh
		}
	}
	return listOk, listErr
}
