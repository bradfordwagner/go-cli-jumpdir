package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
)

const dbfile = ".jumpdirdb.json"
const defaultWeight = 10

// Directory - structure for storing directory info
type Directory struct {
	Weight int    `json:"weight"`
	Path   string `json:"path"`
}

// Backend - Format for storing seen Directories
type Backend struct {
	Directories []Directory `json:"directories"`
}

// New - loads Backend from disk
func New() *Backend {
	b := &Backend{Directories: []Directory{}}
	f, err := os.ReadFile(dbPath())
	if err == nil {
		err = json.Unmarshal(f, b)
		if err != nil {
			panic(err)
		}
	}
	b.Sort()
	return b
}

func (b *Backend) Sort() {
	sort.Slice(b.Directories, func(i, j int) bool {
		return b.Directories[i].Path < b.Directories[j].Path
	})
}

func (b *Backend) contains(dir string) (contains bool) {
	for _, directory := range b.Directories {
		if dir == directory.Path {
			return true
		}
	}
	return
}

func (b *Backend) Insert(directories []string) {
	for _, directory := range directories {
		if !b.contains(directory) {
			b.Directories = append(b.Directories, Directory{
				Weight: defaultWeight,
				Path:   directory,
			})
		}
	}
	b.Sort()
}

func (b *Backend) Save() (err error) {
	p := dbPath()
	marshal, err := json.Marshal(b)
	if err != nil {
		return
	}
	err = os.WriteFile(p, marshal, 0755)
	if err != nil {
		return
	}
	return
}

func (b *Backend) Delete(directories []string) {
	for _, directory := range directories {
		for i, d := range b.Directories {
			if directory == d.Path {
				// move to end and reslice
				b.Directories[i] = b.Directories[len(b.Directories)-1]
				b.Directories = b.Directories[:len(b.Directories)-1]
			}
		}
	}
	b.Sort()
}

func dbPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/%s", dirname, dbfile)
}

func (b *Backend) List() (res []string) {
	sort.Slice(b.Directories, func(i, j int) bool {
		return b.Directories[i].Weight > b.Directories[j].Weight
	})
	for _, directory := range b.Directories {
		res = append(res, directory.Path)
	}
	return
}

func (b *Backend) AddWeight(directory string, weight int) (err error) {
	for i, d := range b.Directories {
		if d.Path == directory {
			b.Directories[i].Weight = b.Directories[i].Weight + weight
			return
		}
	}
	return errors.New(fmt.Sprintf("could not find directory %s", directory))
}
