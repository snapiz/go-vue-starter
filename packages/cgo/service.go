package cgo

import (
	"log"
	"net/http"
	"os"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/graphql-go/graphql"

	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
)

// ServiceConfig using for configure local dev server
type ServiceConfig struct {
	Name    string
	Schema  graphql.Schema
	FetchUser func(qm.QueryMod) (interface{}, error)
}

func graphqlHandler(config handler.Config, w http.ResponseWriter, r *http.Request, fetchUser func(qm.QueryMod) (interface{}, error)) {
	h := handler.New(&config)

	c := Context{
		Response: w,
		Request:  r,
	}

	c.SetHost()
	c.SetUser(fetchUser)
	ctx := NewContext(c)
	h.ContextHandler(ctx, w, r)
}

// Start local dev server
func Start(service ServiceConfig, before func(*mux.Router)) {
	r := mux.NewRouter()
	s := r.PathPrefix("/" + service.Name).Subrouter()

	if before != nil {
		before(s)
	}

	s.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		graphqlHandler(handler.Config{
			Schema:   &service.Schema,
			GraphiQL: true,
		}, w, r, service.FetchUser)
	})

	http.Handle("/", r)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

// Start aws lambda start
/* func Start(config handler.Config) {
	lambda.Start(func(e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		accessor := core.RequestAccessor{}
		r, err := accessor.ProxyEventToHTTPRequest(e)
		w := core.NewProxyResponseWriter()

		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}

		graphqlHandler(config, w, r)

		return w.GetProxyResponse()
	})
} */
