package processing

// Processing
// Broadcast
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

// "sync"

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

func (b *Broadcast) Dispatch(groups []*Group, req *ReqHelper) (map[string]interface{}, map[string]error) {
	list := make(map[string]interface{})
	listOk := make(map[string]interface{})
	listErr := make(map[string]error)
	// make total list
	for _, g := range groups {
		g.Range(func(k, v interface{}) bool {
			ki := k.(string)
			list[ki] = v
			return true // if false, Range stops
		})
	}

	for k, v := range list {
		a := v.(*Authority)
		req.Url(a.Url)
		res, err := b.send.send(req)
		if err != nil {
			listErr[k] = err
		} else {
			listOk[k] = res
		}
	}
	return listOk, listErr
}
