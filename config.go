package processing

// Processing
// Config
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"time"
)

/*
Account
*/

const defaultBlockSize int = 128 // 120 + alignment

/*
Base58
*/

const alphabet string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

/*
Node
*/

const (
	flagStop int32 = iota
	flagWork
)

const defaultAuthoritiesListPath string = "./authorities.json"
const timePauseWorkerAuthStatus time.Duration = 100 * time.Millisecond

/*
Sender
*/

/*
const (
	ReqTypeQuestion int = iota
	ReqTypeAnswer
)
*/

const timeDurationTimeout time.Duration = 10 * time.Second
