package cgo

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/relay"
	"github.com/lithammer/shortuuid"
)

// FromGlobalID takes the "global ID" created by toGlobalID, and returns the type name and ID
// used to create it.
func FromGlobalID(globalID string) *relay.ResolvedGlobalID {
	resolvedID := relay.FromGlobalID(globalID)
	id, _ := decodeID(resolvedID.ID)
	resolvedID.ID = id

	return resolvedID
}

func structoMap(v interface{}) map[string]interface{} {
	var objMap interface{}
	b, _ := json.Marshal(v)
	_ = json.Unmarshal(b, &objMap)

	return objMap.(map[string]interface{})
}

func encodeID(id interface{}) (string, error) {
	u, err := uuid.Parse(fmt.Sprintf("%v", id))

	if err != nil {
		return "", err
	}

	return shortuuid.DefaultEncoder.Encode(u), nil
}

func decodeID(id string) (string, error) {
	u, err := shortuuid.DefaultEncoder.Decode(id)

	if err != nil {
		return "", err
	}

	return u.String(), nil
}

// GlobalIDField short uuid
func GlobalIDField(name string) *graphql.Field {
	return relay.GlobalIDField(name, func(obj interface{}, info graphql.ResolveInfo, ctx context.Context) (string, error) {
		o := structoMap(obj)
		id, _ := o["id"]
		return encodeID(id)
	})
}

func getSelectedFields(selectionPath []string,
	resolveParams graphql.ResolveParams) map[string]bool {
	fields := resolveParams.Info.FieldASTs
	for _, propName := range selectionPath {
		found := false
		for _, field := range fields {
			if field.Name.Value == propName {
				selections := field.SelectionSet.Selections
				fields = make([]*ast.Field, 0)
				for _, selection := range selections {
					fields = append(fields, selection.(*ast.Field))
				}
				found = true
				break
			}
		}
		if !found {
			return map[string]bool{}
		}
	}
	collect := map[string]bool{}
	for _, field := range fields {
		collect[field.Name.Value] = true
	}
	return collect
}

func toCursor(v interface{}) string {
	o := structoMap(v)
	id, _ := o["id"]
	id, _ = encodeID(id)
	createdAt, _ := o["created_at"]
	cursor := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v|%v", id, createdAt)))
	return cursor
}

// Cursor used for node cursor
type Cursor struct {
	ID        string
	CreatedAt string
}

func fromCursor(cursor string) *Cursor {
	if cursor == "" {
		return nil
	}

	cursorDec, _ := base64.StdEncoding.DecodeString(cursor)
	sdata := strings.Split(string(cursorDec), "|")
	id, _ := decodeID(sdata[0])

	return &Cursor{
		ID:        id,
		CreatedAt: sdata[1],
	}
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// ResolveConnection resolve connection
func ResolveConnection(p graphql.ResolveParams,
	c Context,
	resolve func(mods []qm.QueryMod) ([]interface{}, error),
	toSQLFields func(fields map[string]bool) qm.QueryMod,
	orderAsc bool) (interface{}, error) {
	o1 := "asc"
	o2 := "desc"

	if !orderAsc {
		o1 = "desc"
		o2 = "asc"
	}

	args := relay.NewConnectionArguments(p.Args)
	limit := 40
	order := o1
	direction := ""
	if args.First != -1 {
		if args.First < 0 {
			c.Panic(http.StatusBadRequest, "First must be greater than 0")
		} else if args.First > 0 && args.Last < limit {
			limit = args.First
		}

	} else if args.Last != -1 {
		if args.Last < 0 {
			c.Panic(http.StatusBadRequest, "Last must be greater than 0")
		}
		order = o2
		if args.Last > 0 && args.Last < limit {
			limit = args.Last
		}
	}

	cursor := ""
	if args.Before != "" {
		cursor = string(args.Before[:])
		direction = "<"
	} else if args.After != "" {
		cursor = string(args.After[:])
		direction = ">"
	}

	fields := getSelectedFields([]string{p.Info.FieldASTs[0].Name.Value, "edges", "node"}, p)
	fields["id"] = true
	mods := []qm.QueryMod{
		toSQLFields(fields),
	}
	if c := fromCursor(cursor); c != nil {
		mods = append(mods, qm.Where("created_at "+direction+" ? or (created_at = ? and id "+direction+" ?)", c.CreatedAt, c.CreatedAt, c.ID))
	}
	mods = append(mods, qm.OrderBy(fmt.Sprintf("created_at %s, id %s", order, order)))
	mods = append(mods, qm.Limit(limit+1))
	rows, err := resolve(mods)

	if err != nil {
		return rows, err
	}

	hasPreviousPage := false
	hasNextPage := false

	if len(rows) > limit {
		if args.First != -1 {
			hasNextPage = true
		} else if args.Last != -1 {
			hasPreviousPage = true
		}
		rows = rows[:len(rows)-1]
	}

	if order == o2 {
		reverseAny(rows)
	}

	edges := []map[string]interface{}{}
	for _, row := range rows {
		edges = append(edges, map[string]interface{}{
			"cursor": toCursor(row),
			"node":   row,
		})
	}

	return map[string]interface{}{
		"pageInfo": map[string]interface{}{
			"hasPreviousPage": hasPreviousPage,
			"hasNextPage":     hasNextPage,
			"startCursor":     nil,
			"endCursor":       nil,
		},
		"edges": edges,
	}, nil
}
