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

func (r *queryResolver) Dos(ctx context.Context, q *string, url *string, first *int, offset *int) ([]*model.Do, error) {
	qs := *q

	preloads := GetPreloads(ctx)
	log.Println(preloads) // list of requested items.  Pass to dbase call to limit return

	if url == nil {
		temp := ""  // *string cannot be initialized
		url = &temp // in one statement
	}
	us := *url // safe to use url now if it was nil

	if first == nil {
		temp := 20    // int can not be nil, default to 20 returned values..
		first = &temp // in one statement
	}
	f := *first // safe to use url now if it was nil

	if offset == nil {
		temp := 0      // int can not be nil, start from top of list
		offset = &temp // in one statement
	}
	o := *offset // safe to use url now if it was nil

	ds, err := dbase.DescriptionCall(qs, us, f, o)
	if err != nil {
		log.Println(err)
	}

	return ds, err
}

func (r *queryResolver) Dis(ctx context.Context, q *string) ([]*model.Distribution, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
