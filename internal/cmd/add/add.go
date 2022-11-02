package command_add

import "github.com/sirupsen/logrus"

type AddCommand struct {
}

func New() *AddCommand {
	return &AddCommand{}
}

func (a *AddCommand) Run() {
	logrus.Info("hello from run command")
}
