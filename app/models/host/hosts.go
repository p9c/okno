package host

import (
	"fmt"
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"net/http"
)

type Host struct {
	Name   string
	Slug   string
	Host   string
	Routes func(r *mux.Router)
}

func GetHosts(db *scribble.Driver) []*Host {
	return []*Host{
		api(db),
		admin(db),
		okno(db),

		beliRS(db),
		bitNodesNET(db),
		comhttpORG(db),
		comhttpUS(db),

		djordjeMarcetinCOM(db),
		marcetinCOM(db),

		parallelcoinIO(db),
		parallelcoinINFO(db),
		gitParallelcoinINFO(db),
		whitepaperParallelcoinINFO(db),
		docParallelcoinINFO(db),
		explorerParallelcoinINFO(db),
		logParallelcoinINFO(db),

		punqRS(db),
		solutionsRS(db),
		vesicaPiescesORG(db),
	}
}

func (h *Host) testRoutes(r *mux.Router) {
	s := r.Host(h.Host).Subrouter()
	s.Host(h.Host).Path("/").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Host: %v\n", h.Name)
		})
}
