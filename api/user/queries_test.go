package user_test

import (
	"log"
	"os"
	"reflect"
	"testing"

	_ "github.com/bmizerany/pq"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

import (
	"github.com/snapiz/go-vue-starter/api"
	"github.com/snapiz/go-vue-starter/api/schema"
	"github.com/snapiz/go-vue-starter/db"
	"github.com/snapiz/go-vue-starter/db/models"

	"github.com/volatiletech/sqlboiler/queries/qm"
)

var john *models.User
var albert *models.User

func TestMain(m *testing.M) {
	if err := db.Fixtures.Load(); err != nil {
		log.Fatal(err)
	}

	john, _ = models.Users(qm.Where("id = ?", 1)).OneG()
	albert, _ = models.Users(qm.Where("id = ?", 2)).OneG()

	os.Exit(m.Run())
}

func TestMeQuery_anonymous(t *testing.T) {
	query := `
        query MeQuery {
          me {
            id
			email
			role
			createdAt
			updatedAt
          }
        }
	  `

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"me": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema.Schema,
		RequestString: query,
		Context:       api.NewGraphQLContext(nil, nil),
	})

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}

func TestMeQuery_john(t *testing.T) {
	query := `
        query MeQuery {
          me {
            id
			email
			role
			createdAt
			updatedAt
          }
        }
	  `

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"me": map[string]interface{}{
				"id":        "VXNlcjox",
				"email":     "john@doe.com",
				"role":      "USER",
				"createdAt": "2016-01-01T12:30:12Z",
				"updatedAt": "2016-01-01T12:30:12Z",
			},
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema.Schema,
		RequestString: query,
		Context:       api.NewGraphQLContext(nil, john),
	})

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}

func TestMeQuery_albert(t *testing.T) {
	query := `
        query MeQuery {
          me {
            id
			email
			role
			createdAt
			updatedAt
          }
        }
	  `

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"me": map[string]interface{}{
				"id":        "VXNlcjoy",
				"email":     "albert@dupont.com",
				"role":      "STAFF",
				"createdAt": "2016-01-01T12:30:12Z",
				"updatedAt": nil,
			},
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema.Schema,
		RequestString: query,
		Context:       api.NewGraphQLContext(nil, albert),
	})

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
