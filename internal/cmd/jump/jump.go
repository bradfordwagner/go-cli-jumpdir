package command_jump

import (
	"github.com/bradfordwagner/go-cli-jumpdir/internal/backend"
	"os"
)

const defaultAddWeight = 1

type JumpCommand struct {
	backend *backend.Backend
}

func New() *JumpCommand {
	return &JumpCommand{
		backend: backend.New(),
	}
}

func (a *JumpCommand) Jump(dir string) (err error) {
	err = a.backend.AddWeight(dir, defaultAddWeight)
	if err != nil {
		return
	}

	// write into stdout for scripties
	_, err = os.Stdout.Write([]byte(dir))
	if err != nil {
		return
	}

	err = a.backend.Save()
	if err != nil {
		return
	}
	return
}
