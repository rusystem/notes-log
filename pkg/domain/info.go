package domain

import (
	"errors"
	logs "github.com/rusystem/notes-log/pkg/proto"
	"time"
)

const (
	ENTITY_USER = "USER"
	ENTITY_NOTE = "NOTE"

	ACTION_CREATE   = "CREATE"
	ACTION_UPDATE   = "UPDATE"
	ACTION_GET      = "GET"
	ACTION_DELETE   = "DELETE"
	ACTION_REGISTER = "REGISTER"
	ACTION_LOGIN    = "LOGIN"
	ACTION_REFRESH  = "REFRESH"
)

var (
	entities = map[string]logs.LogRequest_Entities{
		ENTITY_USER: logs.LogRequest_USER,
		ENTITY_NOTE: logs.LogRequest_NOTE,
	}

	actions = map[string]logs.LogRequest_Actions{
		ACTION_CREATE:   logs.LogRequest_CREATE,
		ACTION_UPDATE:   logs.LogRequest_UPDATE,
		ACTION_GET:      logs.LogRequest_GET,
		ACTION_DELETE:   logs.LogRequest_DELETE,
		ACTION_REGISTER: logs.LogRequest_REGISTER,
		ACTION_LOGIN:    logs.LogRequest_LOGIN,
		ACTION_REFRESH:  logs.LogRequest_REFRESH,
	}
)

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityID  int64     `bson:"entity_id"`
	Timestamp time.Time `bson:"timestamp"`
}

func ToPbEntity(entity string) (logs.LogRequest_Entities, error) {
	val, ex := entities[entity]
	if !ex {
		return 0, errors.New("invalid entity")
	}

	return val, nil
}

func ToPbAction(action string) (logs.LogRequest_Actions, error) {
	val, ex := actions[action]
	if !ex {
		return 0, errors.New("invalid action")
	}

	return val, nil
}
