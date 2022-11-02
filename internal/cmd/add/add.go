package command_add

import (
	"github.com/bradfordwagner/go-cli-jumpdir/internal/backend"
)

type AddCommand struct {
	backend *backend.Backend
}

func New() *AddCommand {
	return &AddCommand{
		backend: backend.New(),
	}
}

func (a *AddCommand) AddDirectories(directories []string) (err error) {
	a.backend.Insert(directories)

	err = a.backend.Save()
	if err != nil {
		return err
	}
	return
}
