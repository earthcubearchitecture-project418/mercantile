package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/fils/ocdGraphQL/graph/generated"
	"github.com/fils/ocdGraphQL/graph/model"
	"github.com/fils/ocdGraphQL/internal/dbase"
)

func (r *mutationResolver) CreateDo(ctx context.Context, input model.NewDo) (*model.Do, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Dos(ctx context.Context, q *string) ([]*model.Do, error) {
	//panic(fmt.Errorf("not implemented"))

	qs := *q
	ds, err := dbase.DescriptionCall(qs)
	if err != nil {
		log.Println(err)
	}

	return ds, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
