package cgo

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
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

// StartLambda aws lambda start
func StartLambda(fn func(*Router)) {
	lambda.Start(func(e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		r := mux.NewRouter()

		fn(&Router{
			Router: r,
		})

		accessor := core.RequestAccessor{}
		req, err := accessor.ProxyEventToHTTPRequest(e)
		w := core.NewProxyResponseWriter()

		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}

		r.ServeHTTP(w, req)

		return w.GetProxyResponse()
	})
}
