package cgo

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Start local dev server
func Start(fn func(*Router)) {
	r := mux.NewRouter()

	fn(&Router{
		Router: r,
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
