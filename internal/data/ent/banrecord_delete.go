// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-lynx/lynx-layout/internal/data/ent/banrecord"
	"github.com/go-lynx/lynx-layout/internal/data/ent/predicate"
)

// BanRecordDelete is the builder for deleting a BanRecord entity.
type BanRecordDelete struct {
	config
	hooks    []Hook
	mutation *BanRecordMutation
}

// Where appends a list predicates to the BanRecordDelete builder.
func (brd *BanRecordDelete) Where(ps ...predicate.BanRecord) *BanRecordDelete {
	brd.mutation.Where(ps...)
	return brd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (brd *BanRecordDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, brd.sqlExec, brd.mutation, brd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (brd *BanRecordDelete) ExecX(ctx context.Context) int {
	n, err := brd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (brd *BanRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(banrecord.Table, sqlgraph.NewFieldSpec(banrecord.FieldID, field.TypeInt64))
	if ps := brd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, brd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	brd.mutation.done = true
	return affected, err
}

// BanRecordDeleteOne is the builder for deleting a single BanRecord entity.
type BanRecordDeleteOne struct {
	brd *BanRecordDelete
}

// Where appends a list predicates to the BanRecordDelete builder.
func (brdo *BanRecordDeleteOne) Where(ps ...predicate.BanRecord) *BanRecordDeleteOne {
	brdo.brd.mutation.Where(ps...)
	return brdo
}

// Exec executes the deletion query.
func (brdo *BanRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := brdo.brd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{banrecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (brdo *BanRecordDeleteOne) ExecX(ctx context.Context) {
	if err := brdo.Exec(ctx); err != nil {
		panic(err)
	}
}
