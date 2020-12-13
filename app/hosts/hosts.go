package hosts

import (
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/p9c/okno/app/handlers"
	"github.com/p9c/okno/app/jdb"
)

type Host struct {
	Name   string
	Slug   string
	Host   string
	Routes func(r *mux.Router)
}

func GetHosts(db *scribble.Driver) []*Host {
	return []*Host{
		okno(db),
	}
}

func okno(db *scribble.Driver) *Host {
	////////////////
	// okno.rS
	////////////////
	host := &Host{
		Name: "okno",
		Slug: "okno",
		Host: "okno.rs",
	}

	h := handlers.Handlers{jdb.NewJDB(db, host.Slug)}

	routes := func(r *mux.Router) {
		r.Headers("X-Requested-With", "XMLHttpRequest")
		r.Host(host.Host).Path("/").HandlerFunc(h.Homepage()).Name("comhttpus")
		r.Host(host.Host).Path("/{coin}").HandlerFunc(h.Homepage()).Name("comhttpus")
		//r.Host(okno).Path("/coins/").HandlerFunc(rts.Coins).Methods("GET")
		//r.Host(okno).Path("/bitnodes/").HandlerFunc(rts.BitNodes).Methods("GET")
		//r.Host(okno).Path("/bitnodes/{coin}/{nodeid}").HandlerFunc(rts.BitNode).Methods("GET")
		//
		//r.Host("{coin}." + okno).Path("/").HandlerFunc(rts.Coin).Methods("GET")
		//r.Host("{coin}." + okno).Path("/favicon.ico").HandlerFunc(rts.Ico).Name("ico")
		//
		//r.Host("api." + okno).Path("/{coin}/i").HandlerFunc(rts.ApiInfo).Name("info")
	}
	host.Routes = routes
	return host
}

//func (h *Host) ApiData() func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("Asas")
//		vars := mux.Vars(r)
//		w.WriteHeader(http.StatusOK)
//		fmt.Println("asasasassa")
//		//fmt.Fprintf(w, "Category: %v\n", vars["coin"])
//
//		//profile := Profile{"Alex", []string{"snowboarding", "programming"}}
//
//		js, err := json.Marshal(vars)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(js)
//	}
//}
