package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fils/ocdGraphQL/graph/generated"
	"github.com/fils/ocdGraphQL/graph/model"
	"github.com/fils/ocdGraphQL/internal/dbase"
)

func (r *mutationResolver) CreateDo(ctx context.Context, input model.NewDo) (*model.Do, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Dos(ctx context.Context, q *string, url *string) ([]*model.Do, error) {
	qs := *q

	preloads := GetPreloads(ctx)
	log.Println(preloads) // list of requested items.  Pass to dbase call to limit return

	if url == nil {
		temp := ""  // *string cannot be initialized
		url = &temp // in one statement
	}
	us := *url // safe to use url now if it was nil

	ds, err := dbase.DescriptionCall(qs, us)
	if err != nil {
		log.Println(err)
	}

	return ds, err
}

func (r *queryResolver) Dis(ctx context.Context, q *string) ([]*model.Distribution, error) {
	panic(fmt.Errorf("not implemented"))
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetRequestContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.RequestContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.SelectionSet, nil), prefixColumn)...)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)

	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
