// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (m *Manufacturer) Products(ctx context.Context) ([]*Product, error) {
	result, err := m.Edges.ProductsOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryProducts().All(ctx)
	}
	return result, err
}

func (pr *Product) Manufacturer(ctx context.Context) (*Manufacturer, error) {
	result, err := pr.Edges.ManufacturerOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryManufacturer().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Product) Reviews(ctx context.Context) ([]*Review, error) {
	result, err := pr.Edges.ReviewsOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryReviews().All(ctx)
	}
	return result, err
}

func (r *Review) Product(ctx context.Context) (*Product, error) {
	result, err := r.Edges.ProductOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryProduct().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (r *Review) Author(ctx context.Context) (*User, error) {
	result, err := r.Edges.AuthorOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryAuthor().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Reviews(ctx context.Context) ([]*Review, error) {
	result, err := u.Edges.ReviewsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryReviews().All(ctx)
	}
	return result, err
}
