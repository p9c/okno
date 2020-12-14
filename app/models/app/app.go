package app

import (
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/p9c/okno/app/models/host"
	"net/http"
)

type OKNO struct {
	Database *scribble.Driver
	Hosts    []*host.Host
	Server   *http.Server
}
