package handlers

import (
	scribble "github.com/nanobox-io/golang-scribble"

	"github.com/p9c/okno/app/config"
)

// Handlers provides common data for our handlers
type Handlers struct {
	DB     *scribble.Driver
	Config config.Config
}
