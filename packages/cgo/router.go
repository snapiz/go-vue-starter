package cgo

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func graphqlHandler(config handler.Config, w http.ResponseWriter, r *http.Request, fetchUser func(qm.QueryMod) (interface{}, error)) {
	h := handler.New(&config)

	c := Context{
		Response: w,
		Request:  r,
	}

	c.SetHost()
	c.FetchUser(fetchUser)

	h.ContextHandler(NewContext(c), w, r)
}

func restHandler(handlerFunc func(context Context) (map[string]interface{}, error), w http.ResponseWriter, r *http.Request, fetchUser func(qm.QueryMod) (interface{}, error)) {
	defer func() {
		if err := recover(); err != nil {

			jsonData, jerr := json.Marshal(map[string]interface{}{
				"message": err,
			})

			if jerr != nil {
				http.Error(w, jerr.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}()

	c := Context{
		Response: w,
		Request:  r,
	}

	c.SetHost()
	c.FetchUser(fetchUser)
	c.Params = mux.Vars(r)

	for k := range r.URL.Query() {
		c.Params[k] = r.URL.Query().Get(k)
	}

	resp, err := handlerFunc(c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// RouteConfig using for configure router
type RouteConfig struct {
	Path        string
	Schema      *graphql.Schema
	HandlerFunc func(context Context) (map[string]interface{}, error)
}

// Router using for configure server
type Router struct {
	Router    *mux.Router
	FetchUser func(qm.QueryMod) (interface{}, error)
}

//Add route to router
func (r *Router) Add(config *RouteConfig) *mux.Route {
	return r.Router.HandleFunc(config.Path, func(w http.ResponseWriter, req *http.Request) {
		if config.Schema != nil {
			graphqlHandler(handler.Config{
				Schema:   config.Schema,
				GraphiQL: true,
			}, w, req, r.FetchUser)
		} else {
			restHandler(config.HandlerFunc, w, req, r.FetchUser)
		}
	})
}
