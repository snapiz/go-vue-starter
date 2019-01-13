package api_test

import (
	"log"
	"os"
	"testing"

	_ "github.com/bmizerany/pq"
	"github.com/graphql-go/graphql"
	"github.com/snapiz/go-vue-starter/server/api"
	"github.com/snapiz/go-vue-starter/server/db"
	"github.com/snapiz/go-vue-starter/server/db/models"
	"github.com/stretchr/testify/assert"
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

func TestMeQuery_AnonymousShouldNotQueryMe(t *testing.T) {
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
		Schema:        api.Schema,
		RequestString: query,
		Context:       api.NewGraphQLContext(nil, nil),
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
		Schema:        api.Schema,
		RequestString: query,
		Context:       api.NewGraphQLContext(nil, albert),
	})

	assert.Equal(t, expected, result, "Me should be albert")
}
