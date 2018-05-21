package domain

// Processing
// Consensus (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

const (
	ConsensusFills int64 = iota
	ConsensusPositive
	ConsensusNegative
)

/*
Consensus - interface.
This repository is not allowed to delete entities!

Важный момент - использование времени, т.е. когда есть транзакция
и следующая за ней и использующая её выходы. Возможно, нужно
в транзакции предусмотреть время её создания.

Возможны два варианта:
1) При каждом добавлении подтверждения транзакции просто тупо инкрементируется
счётчик и отдаётся результат, а принятие решения о выполнении транзакции снаружи.
Т.е. это тупой enumerator.
2!!) Статистика накапливается, и другой скрипт/воркер периодически заглядывает и
берет состоявшиеся консенсусы. Минусы - нужно учитывать забранные/незабранные.
Плюсы - можно применить сортировку по времени.
3) Консенсус не только считает, но и запускает на выполнение транзакции,
и соответственно, имееет ссылку на репозиторий транзакций. Это звучит плоховато
с точки зрения зависимостей.
*/
type Consensus interface { // enumerator
	// Create(*Transaction) error

	/*
		Vote - подтверждение дублирует создание, т.е. первый конфирм инициализирует
		создание подтверждающего модуля и фиксирует время создания. По этому времени
		можно будет делать выводы по удалению.
	*/
	Vote(*Opinion) int64
	//Clear() []string   // возвращает список транзакций к удалению
	//Ready() []string   // возвращает список утверждённых транзакций
	// SetDuration(int64) // устанавливаем время жизни каждого консенсуса
	SetQuorum(int64) // устанавливаем минимальный лимит подтверждений
}
