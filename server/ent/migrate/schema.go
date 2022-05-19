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
	// ReviewsColumns holds the columns for the "reviews" table.
	ReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "body", Type: field.TypeString},
		{Name: "product_reviews", Type: field.TypeUUID, Nullable: true},
		{Name: "user_reviews", Type: field.TypeUUID, Nullable: true},
	}
	// ReviewsTable holds the schema information for the "reviews" table.
	ReviewsTable = &schema.Table{
		Name:       "reviews",
		Columns:    ReviewsColumns,
		PrimaryKey: []*schema.Column{ReviewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reviews_products_reviews",
				Columns:    []*schema.Column{ReviewsColumns[2]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reviews_users_reviews",
				Columns:    []*schema.Column{ReviewsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ManufacturersTable,
		ProductsTable,
		ReviewsTable,
		UsersTable,
	}
)

func init() {
	ProductsTable.ForeignKeys[0].RefTable = ManufacturersTable
	ReviewsTable.ForeignKeys[0].RefTable = ProductsTable
	ReviewsTable.ForeignKeys[1].RefTable = UsersTable
}