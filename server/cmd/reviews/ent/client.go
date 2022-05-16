// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/ent/migrate"

	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/ent/manufacturer"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/ent/product"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/ent/review"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Manufacturer is the client for interacting with the Manufacturer builders.
	Manufacturer *ManufacturerClient
	// Product is the client for interacting with the Product builders.
	Product *ProductClient
	// Review is the client for interacting with the Review builders.
	Review *ReviewClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Manufacturer = NewManufacturerClient(c.config)
	c.Product = NewProductClient(c.config)
	c.Review = NewReviewClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Manufacturer: NewManufacturerClient(cfg),
		Product:      NewProductClient(cfg),
		Review:       NewReviewClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Manufacturer: NewManufacturerClient(cfg),
		Product:      NewProductClient(cfg),
		Review:       NewReviewClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Manufacturer.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Manufacturer.Use(hooks...)
	c.Product.Use(hooks...)
	c.Review.Use(hooks...)
	c.User.Use(hooks...)
}

// ManufacturerClient is a client for the Manufacturer schema.
type ManufacturerClient struct {
	config
}

// NewManufacturerClient returns a client for the Manufacturer from the given config.
func NewManufacturerClient(c config) *ManufacturerClient {
	return &ManufacturerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `manufacturer.Hooks(f(g(h())))`.
func (c *ManufacturerClient) Use(hooks ...Hook) {
	c.hooks.Manufacturer = append(c.hooks.Manufacturer, hooks...)
}

// Create returns a create builder for Manufacturer.
func (c *ManufacturerClient) Create() *ManufacturerCreate {
	mutation := newManufacturerMutation(c.config, OpCreate)
	return &ManufacturerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Manufacturer entities.
func (c *ManufacturerClient) CreateBulk(builders ...*ManufacturerCreate) *ManufacturerCreateBulk {
	return &ManufacturerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Manufacturer.
func (c *ManufacturerClient) Update() *ManufacturerUpdate {
	mutation := newManufacturerMutation(c.config, OpUpdate)
	return &ManufacturerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ManufacturerClient) UpdateOne(m *Manufacturer) *ManufacturerUpdateOne {
	mutation := newManufacturerMutation(c.config, OpUpdateOne, withManufacturer(m))
	return &ManufacturerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ManufacturerClient) UpdateOneID(id uuid.UUID) *ManufacturerUpdateOne {
	mutation := newManufacturerMutation(c.config, OpUpdateOne, withManufacturerID(id))
	return &ManufacturerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Manufacturer.
func (c *ManufacturerClient) Delete() *ManufacturerDelete {
	mutation := newManufacturerMutation(c.config, OpDelete)
	return &ManufacturerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ManufacturerClient) DeleteOne(m *Manufacturer) *ManufacturerDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ManufacturerClient) DeleteOneID(id uuid.UUID) *ManufacturerDeleteOne {
	builder := c.Delete().Where(manufacturer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ManufacturerDeleteOne{builder}
}

// Query returns a query builder for Manufacturer.
func (c *ManufacturerClient) Query() *ManufacturerQuery {
	return &ManufacturerQuery{
		config: c.config,
	}
}

// Get returns a Manufacturer entity by its id.
func (c *ManufacturerClient) Get(ctx context.Context, id uuid.UUID) (*Manufacturer, error) {
	return c.Query().Where(manufacturer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ManufacturerClient) GetX(ctx context.Context, id uuid.UUID) *Manufacturer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ManufacturerClient) Hooks() []Hook {
	return c.hooks.Manufacturer
}

// ProductClient is a client for the Product schema.
type ProductClient struct {
	config
}

// NewProductClient returns a client for the Product from the given config.
func NewProductClient(c config) *ProductClient {
	return &ProductClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `product.Hooks(f(g(h())))`.
func (c *ProductClient) Use(hooks ...Hook) {
	c.hooks.Product = append(c.hooks.Product, hooks...)
}

// Create returns a create builder for Product.
func (c *ProductClient) Create() *ProductCreate {
	mutation := newProductMutation(c.config, OpCreate)
	return &ProductCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Product entities.
func (c *ProductClient) CreateBulk(builders ...*ProductCreate) *ProductCreateBulk {
	return &ProductCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Product.
func (c *ProductClient) Update() *ProductUpdate {
	mutation := newProductMutation(c.config, OpUpdate)
	return &ProductUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProductClient) UpdateOne(pr *Product) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProduct(pr))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProductClient) UpdateOneID(id uuid.UUID) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProductID(id))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Product.
func (c *ProductClient) Delete() *ProductDelete {
	mutation := newProductMutation(c.config, OpDelete)
	return &ProductDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProductClient) DeleteOne(pr *Product) *ProductDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProductClient) DeleteOneID(id uuid.UUID) *ProductDeleteOne {
	builder := c.Delete().Where(product.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProductDeleteOne{builder}
}

// Query returns a query builder for Product.
func (c *ProductClient) Query() *ProductQuery {
	return &ProductQuery{
		config: c.config,
	}
}

// Get returns a Product entity by its id.
func (c *ProductClient) Get(ctx context.Context, id uuid.UUID) (*Product, error) {
	return c.Query().Where(product.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProductClient) GetX(ctx context.Context, id uuid.UUID) *Product {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryReviews queries the reviews edge of a Product.
func (c *ProductClient) QueryReviews(pr *Product) *ReviewQuery {
	query := &ReviewQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, id),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, product.ReviewsTable, product.ReviewsColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProductClient) Hooks() []Hook {
	return c.hooks.Product
}

// ReviewClient is a client for the Review schema.
type ReviewClient struct {
	config
}

// NewReviewClient returns a client for the Review from the given config.
func NewReviewClient(c config) *ReviewClient {
	return &ReviewClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `review.Hooks(f(g(h())))`.
func (c *ReviewClient) Use(hooks ...Hook) {
	c.hooks.Review = append(c.hooks.Review, hooks...)
}

// Create returns a create builder for Review.
func (c *ReviewClient) Create() *ReviewCreate {
	mutation := newReviewMutation(c.config, OpCreate)
	return &ReviewCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Review entities.
func (c *ReviewClient) CreateBulk(builders ...*ReviewCreate) *ReviewCreateBulk {
	return &ReviewCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Review.
func (c *ReviewClient) Update() *ReviewUpdate {
	mutation := newReviewMutation(c.config, OpUpdate)
	return &ReviewUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReviewClient) UpdateOne(r *Review) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReview(r))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReviewClient) UpdateOneID(id uuid.UUID) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReviewID(id))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Review.
func (c *ReviewClient) Delete() *ReviewDelete {
	mutation := newReviewMutation(c.config, OpDelete)
	return &ReviewDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ReviewClient) DeleteOne(r *Review) *ReviewDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ReviewClient) DeleteOneID(id uuid.UUID) *ReviewDeleteOne {
	builder := c.Delete().Where(review.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReviewDeleteOne{builder}
}

// Query returns a query builder for Review.
func (c *ReviewClient) Query() *ReviewQuery {
	return &ReviewQuery{
		config: c.config,
	}
}

// Get returns a Review entity by its id.
func (c *ReviewClient) Get(ctx context.Context, id uuid.UUID) (*Review, error) {
	return c.Query().Where(review.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReviewClient) GetX(ctx context.Context, id uuid.UUID) *Review {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProduct queries the product edge of a Review.
func (c *ReviewClient) QueryProduct(r *Review) *ProductQuery {
	query := &ProductQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, id),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, review.ProductTable, review.ProductColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthor queries the author edge of a Review.
func (c *ReviewClient) QueryAuthor(r *Review) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, review.AuthorTable, review.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ReviewClient) Hooks() []Hook {
	return c.hooks.Review
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryReviews queries the reviews edge of a User.
func (c *UserClient) QueryReviews(u *User) *ReviewQuery {
	query := &ReviewQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ReviewsTable, user.ReviewsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
