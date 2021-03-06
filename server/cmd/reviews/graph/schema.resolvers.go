package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/graph/model"
	"github.com/johannmunoz/gql_postgres_go/ent"
	"github.com/johannmunoz/gql_postgres_go/ent/product"
	"github.com/johannmunoz/gql_postgres_go/ent/review"
	"github.com/johannmunoz/gql_postgres_go/ent/user"
)

func (r *mutationResolver) ReviewCreate(ctx context.Context, input model.NewReview, authorID string, productID string) (*ent.Review, error) {
	authorUUID, err := uuid.Parse(authorID)
	if err != nil {
		log.Fatal(err)
	}
	author, err := r.client.User.Query().Where(user.IDEQ(authorUUID)).Only(ctx)
	if err != nil {
		log.Fatal(err)
	}
	productUUID, err := uuid.Parse(productID)
	if err != nil {
		log.Fatal(err)
	}
	product, err := r.client.Product.Query().Where(product.IDEQ(productUUID)).Only(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return r.client.Review.Create().SetBody(input.Body).SetProductID(product.ID).SetAuthorID(author.ID).Save(ctx)
}

func (r *mutationResolver) ReviewDelete(ctx context.Context, id string) (bool, error) {
	uuid, parseErr := uuid.Parse(id)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
	err := r.client.Review.DeleteOneID(uuid).Exec(ctx)

	if err != nil {
		log.Fatal(err)
		return false, nil
	}
	return true, nil
}

func (r *productResolver) Reviews(ctx context.Context, obj *ent.Product, before *ent.Cursor, after *ent.Cursor, first *int, last *int, orderBy *ent.ReviewOrder, where *ent.ReviewWhereInput) (*ent.ReviewConnection, error) {
	return r.client.Review.Query().Where(review.HasProductWith(product.IDEQ(obj.ID))).Paginate(ctx, after, first, before, last, ent.WithReviewOrder(orderBy), ent.WithReviewFilter(where.Filter))
}

func (r *queryResolver) Reviews(ctx context.Context, before *ent.Cursor, after *ent.Cursor, first *int, last *int, orderBy *ent.ReviewOrder, where *ent.ReviewWhereInput) (*ent.ReviewConnection, error) {
	return r.client.Review.Query().Paginate(ctx, after, first, before, last, ent.WithReviewOrder(orderBy), ent.WithReviewFilter(where.Filter))
}

func (r *queryResolver) Review(ctx context.Context, id string) (*ent.Review, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Fatal(err)
	}
	return r.client.Review.Query().Where(review.IDEQ(uuid)).Only(ctx)
}

func (r *reviewResolver) ID(ctx context.Context, obj *ent.Review) (string, error) {
	return obj.ID.String(), nil
}

func (r *userResolver) Reviews(ctx context.Context, obj *ent.User, before *ent.Cursor, after *ent.Cursor, first *int, last *int, orderBy *ent.ReviewOrder, where *ent.ReviewWhereInput) (*ent.ReviewConnection, error) {
	return r.client.Review.Query().Where(review.HasAuthorWith(user.IDEQ(obj.ID))).Paginate(ctx, after, first, before, last, ent.WithReviewOrder(orderBy), ent.WithReviewFilter(where.Filter))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
