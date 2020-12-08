package handlers

import (
	"database/sql"

	"github.com/p9c/okno/app/config"
)

// Handlers provides common data for our handlers
type Handlers struct {
	DB     *sql.DB
	Config config.Config
}
