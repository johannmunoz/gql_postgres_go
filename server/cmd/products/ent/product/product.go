// Code generated by entc, DO NOT EDIT.

package product

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUpc holds the string denoting the upc field in the database.
	FieldUpc = "upc"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// EdgeManufacturer holds the string denoting the manufacturer edge name in mutations.
	EdgeManufacturer = "manufacturer"
	// Table holds the table name of the product in the database.
	Table = "products"
	// ManufacturerTable is the table that holds the manufacturer relation/edge.
	ManufacturerTable = "products"
	// ManufacturerInverseTable is the table name for the Manufacturer entity.
	// It exists in this package in order to avoid circular dependency with the "manufacturer" package.
	ManufacturerInverseTable = "manufacturers"
	// ManufacturerColumn is the table column denoting the manufacturer relation/edge.
	ManufacturerColumn = "manufacturer_products"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldUpc,
	FieldPrice,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "products"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"manufacturer_products",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultPrice holds the default value on creation for the "price" field.
	DefaultPrice int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)