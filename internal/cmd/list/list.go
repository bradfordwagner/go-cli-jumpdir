package command_list

import (
	"github.com/bradfordwagner/go-cli-jumpdir/internal/backend"
)

type ListCommand struct {
	backend *backend.Backend
}

func New() *ListCommand {
	return &ListCommand{
		backend: backend.New(),
	}
}

func (a *ListCommand) List() (res []string) {
	return a.backend.List()
}
