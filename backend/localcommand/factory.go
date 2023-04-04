package localcommand

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"syscall"
	"time"

	"github.com/sorenisanerd/gotty/server"
)

type Options struct {
	CloseSignal  int `hcl:"close_signal" flagName:"close-signal" flagSName:"" flagDescribe:"Signal sent to the command process when gotty close it (default: SIGHUP)" default:"1"`
	CloseTimeout int `hcl:"close_timeout" flagName:"close-timeout" flagSName:"" flagDescribe:"Time in seconds to force kill process after client is disconnected (default: -1)" default:"-1"`
}

type Factory struct {
	command          string
	argv             []string
	options          *Options
	opts             []Option
	runningProcesses map[string]server.Slave
	readOnlyProcesses map[string]string
}

func NewFactory(command string, argv []string, options *Options) (*Factory, error) {
	opts := []Option{WithCloseSignal(syscall.Signal(options.CloseSignal))}
	if options.CloseTimeout >= 0 {
		opts = append(opts, WithCloseTimeout(time.Duration(options.CloseTimeout)*time.Second))
	}

	return &Factory{
		command: command,
		argv:    argv,
		options: options,
		opts:    opts,
		runningProcesses: make(map[string]server.Slave),
		readOnlyProcesses: make(map[string]string),
	}, nil
}

func (factory *Factory) Name() string {
	return "local command"
}

func (factory *Factory) New(params map[string][]string, slaveId string) (server.Slave, error) {
	if newSlaveId, ok := factory.readOnlyProcesses[slaveId]; ok {
		slaveId = newSlaveId
	}

	if slave, ok := factory.runningProcesses[slaveId]; ok {
		return slave, nil
	}

	argv := make([]string, len(factory.argv))
	copy(argv, factory.argv)
	if params["arg"] != nil && len(params["arg"]) > 0 {
		argv = append(argv, params["arg"]...)
	}

	slave, err := New(factory.command, argv, factory.opts...)
	fmt.Printf("creating a new slave process\n")
	if err == nil {
		factory.runningProcesses[slaveId] = slave
	}

	return slave, err
}

func (factory *Factory) AddReadonly(slaveId string) (string, error) {
	if alias, ok := factory.readOnlyProcesses[slaveId]; ok {
		return alias, nil
	}

	if _, ok := factory.runningProcesses[slaveId]; ok {
		newId := uuid.New().String()
		factory.readOnlyProcesses[newId] = slaveId
		return newId, nil
	}

	return "", errors.New("slaveId is not valid")

}
