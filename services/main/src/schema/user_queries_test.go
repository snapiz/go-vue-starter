package schema

import (
	"log"
	"os"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/snapiz/go-vue-starter/services/main/src/db"
	"github.com/snapiz/go-vue-starter/packages/tgo"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if err := db.Fixtures.Load(); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestMeQuery_AnonymousShouldNotQueryMe(t *testing.T) {
	query := `
        query MeQuery {
          me {
            id
			email
			role
          }
        }
	  `

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"me": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: query,
		Context:       tgo.Anonymous,
	})

	assert.Equal(t, expected, result, "Me should be nil")
}

func TestMeQuery_ShouldQueryMe(t *testing.T) {
	query := `
        query MeQuery {
          me {
            id
						email
						role
          }
        }
	  `

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"me": map[string]interface{}{
				"id":    "VXNlcjphNXRScjQ1SnU4RFYydWdZZ2JVeGo=",
				"email": "albert@dupont.com",
				"role":  "STAFF",
			},
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: query,
		Context:       tgo.Albert,
	})

	assert.Equal(t, expected, result, "Me should be albert")
}
