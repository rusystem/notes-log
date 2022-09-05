package domain

import (
	"time"
)

type LogItem struct {
	Entity    string    `bson:"entity" json:"entity"`
	Action    string    `bson:"action" json:"action"`
	EntityID  int64     `bson:"entity_id" json:"entityID"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
}
