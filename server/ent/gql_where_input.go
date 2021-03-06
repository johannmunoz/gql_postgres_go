// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/ent/manufacturer"
	"github.com/johannmunoz/gql_postgres_go/ent/predicate"
	"github.com/johannmunoz/gql_postgres_go/ent/product"
	"github.com/johannmunoz/gql_postgres_go/ent/review"
	"github.com/johannmunoz/gql_postgres_go/ent/user"
)

// ManufacturerWhereInput represents a where input for filtering Manufacturer queries.
type ManufacturerWhereInput struct {
	Not *ManufacturerWhereInput   `json:"not,omitempty"`
	Or  []*ManufacturerWhereInput `json:"or,omitempty"`
	And []*ManufacturerWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "products" edge predicates.
	HasProducts     *bool                `json:"hasProducts,omitempty"`
	HasProductsWith []*ProductWhereInput `json:"hasProductsWith,omitempty"`
}

// Filter applies the ManufacturerWhereInput filter on the ManufacturerQuery builder.
func (i *ManufacturerWhereInput) Filter(q *ManufacturerQuery) (*ManufacturerQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering manufacturers.
// An error is returned if the input is empty or invalid.
func (i *ManufacturerWhereInput) P() (predicate.Manufacturer, error) {
	var predicates []predicate.Manufacturer
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, manufacturer.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Manufacturer, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, manufacturer.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Manufacturer, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, manufacturer.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, manufacturer.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, manufacturer.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, manufacturer.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, manufacturer.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, manufacturer.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, manufacturer.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, manufacturer.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, manufacturer.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, manufacturer.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, manufacturer.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, manufacturer.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, manufacturer.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, manufacturer.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, manufacturer.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, manufacturer.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, manufacturer.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, manufacturer.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, manufacturer.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, manufacturer.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, manufacturer.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, manufacturer.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasProducts != nil {
		p := manufacturer.HasProducts()
		if !*i.HasProducts {
			p = manufacturer.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasProductsWith) > 0 {
		with := make([]predicate.Product, 0, len(i.HasProductsWith))
		for _, w := range i.HasProductsWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, manufacturer.HasProductsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("github.com/johannmunoz/gql_postgres_go/ent: empty predicate ManufacturerWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return manufacturer.And(predicates...), nil
	}
}

// ProductWhereInput represents a where input for filtering Product queries.
type ProductWhereInput struct {
	Not *ProductWhereInput   `json:"not,omitempty"`
	Or  []*ProductWhereInput `json:"or,omitempty"`
	And []*ProductWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "upc" field predicates.
	Upc             *string  `json:"upc,omitempty"`
	UpcNEQ          *string  `json:"upcNEQ,omitempty"`
	UpcIn           []string `json:"upcIn,omitempty"`
	UpcNotIn        []string `json:"upcNotIn,omitempty"`
	UpcGT           *string  `json:"upcGT,omitempty"`
	UpcGTE          *string  `json:"upcGTE,omitempty"`
	UpcLT           *string  `json:"upcLT,omitempty"`
	UpcLTE          *string  `json:"upcLTE,omitempty"`
	UpcContains     *string  `json:"upcContains,omitempty"`
	UpcHasPrefix    *string  `json:"upcHasPrefix,omitempty"`
	UpcHasSuffix    *string  `json:"upcHasSuffix,omitempty"`
	UpcEqualFold    *string  `json:"upcEqualFold,omitempty"`
	UpcContainsFold *string  `json:"upcContainsFold,omitempty"`

	// "price" field predicates.
	Price      *int  `json:"price,omitempty"`
	PriceNEQ   *int  `json:"priceNEQ,omitempty"`
	PriceIn    []int `json:"priceIn,omitempty"`
	PriceNotIn []int `json:"priceNotIn,omitempty"`
	PriceGT    *int  `json:"priceGT,omitempty"`
	PriceGTE   *int  `json:"priceGTE,omitempty"`
	PriceLT    *int  `json:"priceLT,omitempty"`
	PriceLTE   *int  `json:"priceLTE,omitempty"`

	// "manufacturer" edge predicates.
	HasManufacturer     *bool                     `json:"hasManufacturer,omitempty"`
	HasManufacturerWith []*ManufacturerWhereInput `json:"hasManufacturerWith,omitempty"`

	// "reviews" edge predicates.
	HasReviews     *bool               `json:"hasReviews,omitempty"`
	HasReviewsWith []*ReviewWhereInput `json:"hasReviewsWith,omitempty"`
}

// Filter applies the ProductWhereInput filter on the ProductQuery builder.
func (i *ProductWhereInput) Filter(q *ProductQuery) (*ProductQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering products.
// An error is returned if the input is empty or invalid.
func (i *ProductWhereInput) P() (predicate.Product, error) {
	var predicates []predicate.Product
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, product.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Product, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, product.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Product, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, product.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, product.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, product.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, product.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, product.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, product.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, product.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, product.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, product.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, product.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, product.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, product.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, product.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, product.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, product.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, product.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, product.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, product.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, product.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, product.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, product.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, product.NameContainsFold(*i.NameContainsFold))
	}
	if i.Upc != nil {
		predicates = append(predicates, product.UpcEQ(*i.Upc))
	}
	if i.UpcNEQ != nil {
		predicates = append(predicates, product.UpcNEQ(*i.UpcNEQ))
	}
	if len(i.UpcIn) > 0 {
		predicates = append(predicates, product.UpcIn(i.UpcIn...))
	}
	if len(i.UpcNotIn) > 0 {
		predicates = append(predicates, product.UpcNotIn(i.UpcNotIn...))
	}
	if i.UpcGT != nil {
		predicates = append(predicates, product.UpcGT(*i.UpcGT))
	}
	if i.UpcGTE != nil {
		predicates = append(predicates, product.UpcGTE(*i.UpcGTE))
	}
	if i.UpcLT != nil {
		predicates = append(predicates, product.UpcLT(*i.UpcLT))
	}
	if i.UpcLTE != nil {
		predicates = append(predicates, product.UpcLTE(*i.UpcLTE))
	}
	if i.UpcContains != nil {
		predicates = append(predicates, product.UpcContains(*i.UpcContains))
	}
	if i.UpcHasPrefix != nil {
		predicates = append(predicates, product.UpcHasPrefix(*i.UpcHasPrefix))
	}
	if i.UpcHasSuffix != nil {
		predicates = append(predicates, product.UpcHasSuffix(*i.UpcHasSuffix))
	}
	if i.UpcEqualFold != nil {
		predicates = append(predicates, product.UpcEqualFold(*i.UpcEqualFold))
	}
	if i.UpcContainsFold != nil {
		predicates = append(predicates, product.UpcContainsFold(*i.UpcContainsFold))
	}
	if i.Price != nil {
		predicates = append(predicates, product.PriceEQ(*i.Price))
	}
	if i.PriceNEQ != nil {
		predicates = append(predicates, product.PriceNEQ(*i.PriceNEQ))
	}
	if len(i.PriceIn) > 0 {
		predicates = append(predicates, product.PriceIn(i.PriceIn...))
	}
	if len(i.PriceNotIn) > 0 {
		predicates = append(predicates, product.PriceNotIn(i.PriceNotIn...))
	}
	if i.PriceGT != nil {
		predicates = append(predicates, product.PriceGT(*i.PriceGT))
	}
	if i.PriceGTE != nil {
		predicates = append(predicates, product.PriceGTE(*i.PriceGTE))
	}
	if i.PriceLT != nil {
		predicates = append(predicates, product.PriceLT(*i.PriceLT))
	}
	if i.PriceLTE != nil {
		predicates = append(predicates, product.PriceLTE(*i.PriceLTE))
	}

	if i.HasManufacturer != nil {
		p := product.HasManufacturer()
		if !*i.HasManufacturer {
			p = product.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasManufacturerWith) > 0 {
		with := make([]predicate.Manufacturer, 0, len(i.HasManufacturerWith))
		for _, w := range i.HasManufacturerWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, product.HasManufacturerWith(with...))
	}
	if i.HasReviews != nil {
		p := product.HasReviews()
		if !*i.HasReviews {
			p = product.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasReviewsWith) > 0 {
		with := make([]predicate.Review, 0, len(i.HasReviewsWith))
		for _, w := range i.HasReviewsWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, product.HasReviewsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("github.com/johannmunoz/gql_postgres_go/ent: empty predicate ProductWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return product.And(predicates...), nil
	}
}

// ReviewWhereInput represents a where input for filtering Review queries.
type ReviewWhereInput struct {
	Not *ReviewWhereInput   `json:"not,omitempty"`
	Or  []*ReviewWhereInput `json:"or,omitempty"`
	And []*ReviewWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "body" field predicates.
	Body             *string  `json:"body,omitempty"`
	BodyNEQ          *string  `json:"bodyNEQ,omitempty"`
	BodyIn           []string `json:"bodyIn,omitempty"`
	BodyNotIn        []string `json:"bodyNotIn,omitempty"`
	BodyGT           *string  `json:"bodyGT,omitempty"`
	BodyGTE          *string  `json:"bodyGTE,omitempty"`
	BodyLT           *string  `json:"bodyLT,omitempty"`
	BodyLTE          *string  `json:"bodyLTE,omitempty"`
	BodyContains     *string  `json:"bodyContains,omitempty"`
	BodyHasPrefix    *string  `json:"bodyHasPrefix,omitempty"`
	BodyHasSuffix    *string  `json:"bodyHasSuffix,omitempty"`
	BodyEqualFold    *string  `json:"bodyEqualFold,omitempty"`
	BodyContainsFold *string  `json:"bodyContainsFold,omitempty"`

	// "product" edge predicates.
	HasProduct     *bool                `json:"hasProduct,omitempty"`
	HasProductWith []*ProductWhereInput `json:"hasProductWith,omitempty"`

	// "author" edge predicates.
	HasAuthor     *bool             `json:"hasAuthor,omitempty"`
	HasAuthorWith []*UserWhereInput `json:"hasAuthorWith,omitempty"`
}

// Filter applies the ReviewWhereInput filter on the ReviewQuery builder.
func (i *ReviewWhereInput) Filter(q *ReviewQuery) (*ReviewQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering reviews.
// An error is returned if the input is empty or invalid.
func (i *ReviewWhereInput) P() (predicate.Review, error) {
	var predicates []predicate.Review
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, review.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Review, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, review.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Review, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, review.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, review.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, review.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, review.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, review.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, review.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, review.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, review.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, review.IDLTE(*i.IDLTE))
	}
	if i.Body != nil {
		predicates = append(predicates, review.BodyEQ(*i.Body))
	}
	if i.BodyNEQ != nil {
		predicates = append(predicates, review.BodyNEQ(*i.BodyNEQ))
	}
	if len(i.BodyIn) > 0 {
		predicates = append(predicates, review.BodyIn(i.BodyIn...))
	}
	if len(i.BodyNotIn) > 0 {
		predicates = append(predicates, review.BodyNotIn(i.BodyNotIn...))
	}
	if i.BodyGT != nil {
		predicates = append(predicates, review.BodyGT(*i.BodyGT))
	}
	if i.BodyGTE != nil {
		predicates = append(predicates, review.BodyGTE(*i.BodyGTE))
	}
	if i.BodyLT != nil {
		predicates = append(predicates, review.BodyLT(*i.BodyLT))
	}
	if i.BodyLTE != nil {
		predicates = append(predicates, review.BodyLTE(*i.BodyLTE))
	}
	if i.BodyContains != nil {
		predicates = append(predicates, review.BodyContains(*i.BodyContains))
	}
	if i.BodyHasPrefix != nil {
		predicates = append(predicates, review.BodyHasPrefix(*i.BodyHasPrefix))
	}
	if i.BodyHasSuffix != nil {
		predicates = append(predicates, review.BodyHasSuffix(*i.BodyHasSuffix))
	}
	if i.BodyEqualFold != nil {
		predicates = append(predicates, review.BodyEqualFold(*i.BodyEqualFold))
	}
	if i.BodyContainsFold != nil {
		predicates = append(predicates, review.BodyContainsFold(*i.BodyContainsFold))
	}

	if i.HasProduct != nil {
		p := review.HasProduct()
		if !*i.HasProduct {
			p = review.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasProductWith) > 0 {
		with := make([]predicate.Product, 0, len(i.HasProductWith))
		for _, w := range i.HasProductWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, review.HasProductWith(with...))
	}
	if i.HasAuthor != nil {
		p := review.HasAuthor()
		if !*i.HasAuthor {
			p = review.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasAuthorWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasAuthorWith))
		for _, w := range i.HasAuthorWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, review.HasAuthorWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("github.com/johannmunoz/gql_postgres_go/ent: empty predicate ReviewWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return review.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Not *UserWhereInput   `json:"not,omitempty"`
	Or  []*UserWhereInput `json:"or,omitempty"`
	And []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "email" field predicates.
	Email             *string  `json:"email,omitempty"`
	EmailNEQ          *string  `json:"emailNEQ,omitempty"`
	EmailIn           []string `json:"emailIn,omitempty"`
	EmailNotIn        []string `json:"emailNotIn,omitempty"`
	EmailGT           *string  `json:"emailGT,omitempty"`
	EmailGTE          *string  `json:"emailGTE,omitempty"`
	EmailLT           *string  `json:"emailLT,omitempty"`
	EmailLTE          *string  `json:"emailLTE,omitempty"`
	EmailContains     *string  `json:"emailContains,omitempty"`
	EmailHasPrefix    *string  `json:"emailHasPrefix,omitempty"`
	EmailHasSuffix    *string  `json:"emailHasSuffix,omitempty"`
	EmailEqualFold    *string  `json:"emailEqualFold,omitempty"`
	EmailContainsFold *string  `json:"emailContainsFold,omitempty"`

	// "username" field predicates.
	Username             *string  `json:"username,omitempty"`
	UsernameNEQ          *string  `json:"usernameNEQ,omitempty"`
	UsernameIn           []string `json:"usernameIn,omitempty"`
	UsernameNotIn        []string `json:"usernameNotIn,omitempty"`
	UsernameGT           *string  `json:"usernameGT,omitempty"`
	UsernameGTE          *string  `json:"usernameGTE,omitempty"`
	UsernameLT           *string  `json:"usernameLT,omitempty"`
	UsernameLTE          *string  `json:"usernameLTE,omitempty"`
	UsernameContains     *string  `json:"usernameContains,omitempty"`
	UsernameHasPrefix    *string  `json:"usernameHasPrefix,omitempty"`
	UsernameHasSuffix    *string  `json:"usernameHasSuffix,omitempty"`
	UsernameEqualFold    *string  `json:"usernameEqualFold,omitempty"`
	UsernameContainsFold *string  `json:"usernameContainsFold,omitempty"`

	// "reviews" edge predicates.
	HasReviews     *bool               `json:"hasReviews,omitempty"`
	HasReviewsWith []*ReviewWhereInput `json:"hasReviewsWith,omitempty"`
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.Email != nil {
		predicates = append(predicates, user.EmailEQ(*i.Email))
	}
	if i.EmailNEQ != nil {
		predicates = append(predicates, user.EmailNEQ(*i.EmailNEQ))
	}
	if len(i.EmailIn) > 0 {
		predicates = append(predicates, user.EmailIn(i.EmailIn...))
	}
	if len(i.EmailNotIn) > 0 {
		predicates = append(predicates, user.EmailNotIn(i.EmailNotIn...))
	}
	if i.EmailGT != nil {
		predicates = append(predicates, user.EmailGT(*i.EmailGT))
	}
	if i.EmailGTE != nil {
		predicates = append(predicates, user.EmailGTE(*i.EmailGTE))
	}
	if i.EmailLT != nil {
		predicates = append(predicates, user.EmailLT(*i.EmailLT))
	}
	if i.EmailLTE != nil {
		predicates = append(predicates, user.EmailLTE(*i.EmailLTE))
	}
	if i.EmailContains != nil {
		predicates = append(predicates, user.EmailContains(*i.EmailContains))
	}
	if i.EmailHasPrefix != nil {
		predicates = append(predicates, user.EmailHasPrefix(*i.EmailHasPrefix))
	}
	if i.EmailHasSuffix != nil {
		predicates = append(predicates, user.EmailHasSuffix(*i.EmailHasSuffix))
	}
	if i.EmailEqualFold != nil {
		predicates = append(predicates, user.EmailEqualFold(*i.EmailEqualFold))
	}
	if i.EmailContainsFold != nil {
		predicates = append(predicates, user.EmailContainsFold(*i.EmailContainsFold))
	}
	if i.Username != nil {
		predicates = append(predicates, user.UsernameEQ(*i.Username))
	}
	if i.UsernameNEQ != nil {
		predicates = append(predicates, user.UsernameNEQ(*i.UsernameNEQ))
	}
	if len(i.UsernameIn) > 0 {
		predicates = append(predicates, user.UsernameIn(i.UsernameIn...))
	}
	if len(i.UsernameNotIn) > 0 {
		predicates = append(predicates, user.UsernameNotIn(i.UsernameNotIn...))
	}
	if i.UsernameGT != nil {
		predicates = append(predicates, user.UsernameGT(*i.UsernameGT))
	}
	if i.UsernameGTE != nil {
		predicates = append(predicates, user.UsernameGTE(*i.UsernameGTE))
	}
	if i.UsernameLT != nil {
		predicates = append(predicates, user.UsernameLT(*i.UsernameLT))
	}
	if i.UsernameLTE != nil {
		predicates = append(predicates, user.UsernameLTE(*i.UsernameLTE))
	}
	if i.UsernameContains != nil {
		predicates = append(predicates, user.UsernameContains(*i.UsernameContains))
	}
	if i.UsernameHasPrefix != nil {
		predicates = append(predicates, user.UsernameHasPrefix(*i.UsernameHasPrefix))
	}
	if i.UsernameHasSuffix != nil {
		predicates = append(predicates, user.UsernameHasSuffix(*i.UsernameHasSuffix))
	}
	if i.UsernameEqualFold != nil {
		predicates = append(predicates, user.UsernameEqualFold(*i.UsernameEqualFold))
	}
	if i.UsernameContainsFold != nil {
		predicates = append(predicates, user.UsernameContainsFold(*i.UsernameContainsFold))
	}

	if i.HasReviews != nil {
		p := user.HasReviews()
		if !*i.HasReviews {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasReviewsWith) > 0 {
		with := make([]predicate.Review, 0, len(i.HasReviewsWith))
		for _, w := range i.HasReviewsWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasReviewsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("github.com/johannmunoz/gql_postgres_go/ent: empty predicate UserWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}
