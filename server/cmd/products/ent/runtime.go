// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/ent/manufacturer"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/ent/product"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	manufacturerFields := schema.Manufacturer{}.Fields()
	_ = manufacturerFields
	// manufacturerDescID is the schema descriptor for id field.
	manufacturerDescID := manufacturerFields[0].Descriptor()
	// manufacturer.DefaultID holds the default value on creation for the id field.
	manufacturer.DefaultID = manufacturerDescID.Default.(func() uuid.UUID)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescPrice is the schema descriptor for price field.
	productDescPrice := productFields[3].Descriptor()
	// product.DefaultPrice holds the default value on creation for the price field.
	product.DefaultPrice = productDescPrice.Default.(int)
	// productDescID is the schema descriptor for id field.
	productDescID := productFields[0].Descriptor()
	// product.DefaultID holds the default value on creation for the id field.
	product.DefaultID = productDescID.Default.(func() uuid.UUID)
}