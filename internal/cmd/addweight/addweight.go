package command_addweight

import (
	"github.com/bradfordwagner/go-cli-jumpdir/internal/backend"
)

const defaultWeight = 1

type AddWeightCommand struct {
	backend *backend.Backend
}

func New() *AddWeightCommand {
	return &AddWeightCommand{
		backend: backend.New(),
	}
}

func (a *AddWeightCommand) AddWeight(directory string) (err error) {
	err = a.backend.AddWeight(directory, defaultWeight)
	if err != nil {
		return err
	}

	return a.backend.Save()
}
