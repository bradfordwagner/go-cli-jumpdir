package command_list

import (
	"github.com/bradfordwagner/go-cli-jumpdir/internal/backend"
	"os"
	"strings"
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

func (a *ListCommand) ListToStdOut() {
	dirs := a.backend.List()
	res := strings.Join(dirs, "\n")
	_, _ = os.Stdout.Write([]byte(res))
}
