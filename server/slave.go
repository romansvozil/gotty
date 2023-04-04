package server

import (
	"github.com/sorenisanerd/gotty/webtty"
)

// Slave is webtty.Slave with some additional methods.
type Slave interface {
	webtty.Slave

	Close() error
	HasPublicReadOnly() bool
	SetHasPublicReadOnly(hasPublicReadOnly bool)
}

type Factory interface {
	Name() string
	New(params map[string][]string, slaveId string) (Slave, error)
	AddReadonly(slaveId string) (string, error)
}
