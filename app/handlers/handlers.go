package handlers

import (
	"github.com/1lann/cete"

	"github.com/p9c/okno/app/config"
)

// Handlers provides common data for our handlers
type Handlers struct {
	DB     *cete.DB
	Config config.Config
}
