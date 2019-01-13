package api

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

var nodeDefinitions = *relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
	IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
		/* c := NewContext(ctx) */
		// resolve id from global id
		/* resolvedID := relay.FromGlobalID(id) */

		// based on id and its type, return the object
		/* switch resolvedID.Type {
		case "Faction":
			return GetFaction(resolvedID.ID), nil
		case "Ship":
			return GetShip(resolvedID.ID), nil
		default:
			return nil, errors.New("Unknown node type")
		} */
		return nil, nil
	},
	TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
		// based on the type of the value, return GraphQLObjectType
		/* switch p.Value.(type) {
		case *Faction:
			return factionType
		default:
			return shipType
		} */
		return nil
	},
})
