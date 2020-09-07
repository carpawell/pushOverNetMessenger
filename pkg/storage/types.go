package storage

import "time"

type Notification struct {
	Message string
	Status  int
	Time    time.Time
}

type MessagesStatistic struct {
	From    *time.Time `json:"fromTime"`
	Total   int64      `json:"totalMessages"`
	Succeed int64      `json:"succeedMessages"`
	Failed  int64      `json:"failedMessages"`
}
