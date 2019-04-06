package schema

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/snapiz/go-vue-starter/services/main/src/db"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null"
)

var anonymous = cgo.NewContext(cgo.Context{})

var john = cgo.NewContext(cgo.Context{
	ID:           "4c9f32df-c112-4d02-a459-3493fac49ea9",
	Role:         "user",
	Email:        "john@doe.com",
	EmailHash:    "6a6c19fea4a3676970167ce51f39e6ee",
	TokenVersion: null.Int64From(1546097186),
	CreatedAt:    time.Unix(1405544146, 0),
	UpdatedAt:    null.TimeFrom(time.Unix(1405544146, 0)),
})

var albert = cgo.NewContext(cgo.Context{
	ID:           "04231367-deef-4444-bc41-529281445b5f",
	Role:         "staff",
	Email:        "albert@dupont.com",
	EmailHash:    "58a1cb05c9a75b272003cd6e4b4dc543",
	CreatedAt:    time.Unix(1405544146, 0),
	Username:     null.StringFrom("albert"),
	TokenVersion: null.Int64From(1546097186),
	Password:     null.StringFrom("$argon2id$v=19$m=65536,t=3,p=2$XEqhUPoyTftr1HQZ7/p0dA$CvX5+Et7e+QlgvwjrK2J7bFtBODhDjTuIoh5wJlDCl4"),
})

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
		Context:       anonymous,
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
		Context:       albert,
	})

	assert.Equal(t, expected, result, "Me should be albert")
}
