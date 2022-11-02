package backend

import "os"

// Directory - structure for storing directory info
type Directory struct {
	Weight int
	Path   string
}

// Backend - Format for storing seen directories
type Backend struct {
	directories []Directory
}

// Init - loads Backend from disk
func (d *Backend) Init() {
	dirname, err := os.UserHomeDir()

}
