package command_leaderboard

import (
	"fmt"
	"github.com/bradfordwagner/go-cli-jumpdir/internal/backend"
	"os"
	"sort"
)

type LeaderboardCommand struct {
	backend *backend.Backend
	reverse bool
	limit   int
}

func New(reverse bool, limit int) *LeaderboardCommand {
	return &LeaderboardCommand{
		backend: backend.New(),
		reverse: reverse,
		limit:   limit,
	}
}

func (l *LeaderboardCommand) Leaderboard() (res []string) {
	dirs := l.backend.Directories

	sort.Slice(dirs, func(i, j int) bool {
		if l.reverse {
			return dirs[i].Weight < dirs[j].Weight
		}
		return dirs[i].Weight > dirs[j].Weight
	})

	count := len(dirs)
	if l.limit > 0 && l.limit < count {
		count = l.limit
	}

	for i := 0; i < count; i++ {
		res = append(res, fmt.Sprintf("%d\t%s", dirs[i].Weight, dirs[i].Path))
	}
	return
}

func (l *LeaderboardCommand) LeaderboardToStdOut() {
	entries := l.Leaderboard()
	for _, entry := range entries {
		fmt.Fprintln(os.Stdout, entry)
	}
}
