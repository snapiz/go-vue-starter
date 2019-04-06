package cgo

import (
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"

	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
)

// ServiceConfig using for configure local dev server
type ServiceConfig struct {
	Name   string
	Schema graphql.Schema
	Before func(Context) Context
}

func graphqlHandler(config handler.Config, w http.ResponseWriter, r *http.Request, before func(Context) Context) {
	h := handler.New(&config)

	c := Context{
		Response: w,
		Request:  r,
	}

	if before != nil {
		c = before(c)
	}

	ctx := NewContext(c)
	h.ContextHandler(ctx, w, r)
}

// Start local dev server
func Start(services []ServiceConfig) {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	for _, service := range services {
		s.HandleFunc("/"+service.Name, func(w http.ResponseWriter, r *http.Request) {
			graphqlHandler(handler.Config{
				Schema:   &service.Schema,
				GraphiQL: true,
			}, w, r, service.Before)
		})
	}

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
