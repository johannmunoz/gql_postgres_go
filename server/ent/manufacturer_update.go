// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/ent/manufacturer"
	"github.com/johannmunoz/gql_postgres_go/ent/predicate"
	"github.com/johannmunoz/gql_postgres_go/ent/product"
)

// ManufacturerUpdate is the builder for updating Manufacturer entities.
type ManufacturerUpdate struct {
	config
	hooks    []Hook
	mutation *ManufacturerMutation
}

// Where appends a list predicates to the ManufacturerUpdate builder.
func (mu *ManufacturerUpdate) Where(ps ...predicate.Manufacturer) *ManufacturerUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "name" field.
func (mu *ManufacturerUpdate) SetName(s string) *ManufacturerUpdate {
	mu.mutation.SetName(s)
	return mu
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (mu *ManufacturerUpdate) AddProductIDs(ids ...uuid.UUID) *ManufacturerUpdate {
	mu.mutation.AddProductIDs(ids...)
	return mu
}

// AddProducts adds the "products" edges to the Product entity.
func (mu *ManufacturerUpdate) AddProducts(p ...*Product) *ManufacturerUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddProductIDs(ids...)
}

// Mutation returns the ManufacturerMutation object of the builder.
func (mu *ManufacturerUpdate) Mutation() *ManufacturerMutation {
	return mu.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (mu *ManufacturerUpdate) ClearProducts() *ManufacturerUpdate {
	mu.mutation.ClearProducts()
	return mu
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (mu *ManufacturerUpdate) RemoveProductIDs(ids ...uuid.UUID) *ManufacturerUpdate {
	mu.mutation.RemoveProductIDs(ids...)
	return mu
}

// RemoveProducts removes "products" edges to Product entities.
func (mu *ManufacturerUpdate) RemoveProducts(p ...*Product) *ManufacturerUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *ManufacturerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ManufacturerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *ManufacturerUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *ManufacturerUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *ManufacturerUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *ManufacturerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   manufacturer.Table,
			Columns: manufacturer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: manufacturer.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manufacturer.FieldName,
		})
	}
	if mu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.ProductsTable,
			Columns: []string{manufacturer.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedProductsIDs(); len(nodes) > 0 && !mu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.ProductsTable,
			Columns: []string{manufacturer.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.ProductsTable,
			Columns: []string{manufacturer.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{manufacturer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ManufacturerUpdateOne is the builder for updating a single Manufacturer entity.
type ManufacturerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ManufacturerMutation
}

// SetName sets the "name" field.
func (muo *ManufacturerUpdateOne) SetName(s string) *ManufacturerUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (muo *ManufacturerUpdateOne) AddProductIDs(ids ...uuid.UUID) *ManufacturerUpdateOne {
	muo.mutation.AddProductIDs(ids...)
	return muo
}

// AddProducts adds the "products" edges to the Product entity.
func (muo *ManufacturerUpdateOne) AddProducts(p ...*Product) *ManufacturerUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddProductIDs(ids...)
}

// Mutation returns the ManufacturerMutation object of the builder.
func (muo *ManufacturerUpdateOne) Mutation() *ManufacturerMutation {
	return muo.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (muo *ManufacturerUpdateOne) ClearProducts() *ManufacturerUpdateOne {
	muo.mutation.ClearProducts()
	return muo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (muo *ManufacturerUpdateOne) RemoveProductIDs(ids ...uuid.UUID) *ManufacturerUpdateOne {
	muo.mutation.RemoveProductIDs(ids...)
	return muo
}

// RemoveProducts removes "products" edges to Product entities.
func (muo *ManufacturerUpdateOne) RemoveProducts(p ...*Product) *ManufacturerUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemoveProductIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *ManufacturerUpdateOne) Select(field string, fields ...string) *ManufacturerUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Manufacturer entity.
func (muo *ManufacturerUpdateOne) Save(ctx context.Context) (*Manufacturer, error) {
	var (
		err  error
		node *Manufacturer
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ManufacturerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *ManufacturerUpdateOne) SaveX(ctx context.Context) *Manufacturer {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *ManufacturerUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *ManufacturerUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *ManufacturerUpdateOne) sqlSave(ctx context.Context) (_node *Manufacturer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   manufacturer.Table,
			Columns: manufacturer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: manufacturer.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Manufacturer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, manufacturer.FieldID)
		for _, f := range fields {
			if !manufacturer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != manufacturer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manufacturer.FieldName,
		})
	}
	if muo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.ProductsTable,
			Columns: []string{manufacturer.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !muo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.ProductsTable,
			Columns: []string{manufacturer.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manufacturer.ProductsTable,
			Columns: []string{manufacturer.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Manufacturer{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{manufacturer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
