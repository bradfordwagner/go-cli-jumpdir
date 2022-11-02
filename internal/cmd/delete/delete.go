package command_delete

import (
	"github.com/bradfordwagner/go-cli-jumpdir/internal/backend"
)

type DeleteCommand struct {
	backend *backend.Backend
}

func New() *DeleteCommand {
	return &DeleteCommand{
		backend: backend.New(),
	}
}

func (a *DeleteCommand) DeleteDirectories(directories []string) (err error) {
	a.backend.Delete(directories)

	err = a.backend.Save()
	if err != nil {
		return err
	}
	return
}
