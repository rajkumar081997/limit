package store

import (
	"errors"

	handler "github.com/m/v2/new_server/handler"
)

type Serialize interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type config interface {
	Handler() string
}

type Store interface {
	Item(string) (string, error)
	GetId(string) (string, error)
	List(string) ([]string, error)
	Remove(string) (string, error)
	Close() error
}

func Newstore(dbname string) (Store,error) {
	switch dbname {
	case "bolt":
		return handler.Newbolt_handler(&handler.Bolt_handler{})

	default:
		return nil, errors.New("We don't have MongoDb")
	}
}
