// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ManufacturersColumns holds the columns for the "manufacturers" table.
	ManufacturersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
	}
	// ManufacturersTable holds the schema information for the "manufacturers" table.
	ManufacturersTable = &schema.Table{
		Name:       "manufacturers",
		Columns:    ManufacturersColumns,
		PrimaryKey: []*schema.Column{ManufacturersColumns[0]},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "upc", Type: field.TypeString},
		{Name: "price", Type: field.TypeInt, Default: 0},
		{Name: "manufacturer_products", Type: field.TypeUUID, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_manufacturers_products",
				Columns:    []*schema.Column{ProductsColumns[4]},
				RefColumns: []*schema.Column{ManufacturersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ManufacturersTable,
		ProductsTable,
	}
)

func init() {
	ProductsTable.ForeignKeys[0].RefTable = ManufacturersTable
}
