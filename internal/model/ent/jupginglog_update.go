// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"junction/internal/model/ent/jupginglog"
	"junction/internal/model/ent/member"
	"junction/internal/model/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JupgingLogUpdate is the builder for updating JupgingLog entities.
type JupgingLogUpdate struct {
	config
	hooks    []Hook
	mutation *JupgingLogMutation
}

// Where appends a list predicates to the JupgingLogUpdate builder.
func (jlu *JupgingLogUpdate) Where(ps ...predicate.JupgingLog) *JupgingLogUpdate {
	jlu.mutation.Where(ps...)
	return jlu
}

// SetStartDate sets the "startDate" field.
func (jlu *JupgingLogUpdate) SetStartDate(s string) *JupgingLogUpdate {
	jlu.mutation.SetStartDate(s)
	return jlu
}

// SetNillableStartDate sets the "startDate" field if the given value is not nil.
func (jlu *JupgingLogUpdate) SetNillableStartDate(s *string) *JupgingLogUpdate {
	if s != nil {
		jlu.SetStartDate(*s)
	}
	return jlu
}

// SetEndDate sets the "endDate" field.
func (jlu *JupgingLogUpdate) SetEndDate(s string) *JupgingLogUpdate {
	jlu.mutation.SetEndDate(s)
	return jlu
}

// SetNillableEndDate sets the "endDate" field if the given value is not nil.
func (jlu *JupgingLogUpdate) SetNillableEndDate(s *string) *JupgingLogUpdate {
	if s != nil {
		jlu.SetEndDate(*s)
	}
	return jlu
}

// SetLog sets the "log" field.
func (jlu *JupgingLogUpdate) SetLog(s string) *JupgingLogUpdate {
	jlu.mutation.SetLog(s)
	return jlu
}

// SetNillableLog sets the "log" field if the given value is not nil.
func (jlu *JupgingLogUpdate) SetNillableLog(s *string) *JupgingLogUpdate {
	if s != nil {
		jlu.SetLog(*s)
	}
	return jlu
}

// SetMemberID sets the "member_id" field.
func (jlu *JupgingLogUpdate) SetMemberID(i int) *JupgingLogUpdate {
	jlu.mutation.SetMemberID(i)
	return jlu
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (jlu *JupgingLogUpdate) SetNillableMemberID(i *int) *JupgingLogUpdate {
	if i != nil {
		jlu.SetMemberID(*i)
	}
	return jlu
}

// SetMember sets the "member" edge to the Member entity.
func (jlu *JupgingLogUpdate) SetMember(m *Member) *JupgingLogUpdate {
	return jlu.SetMemberID(m.ID)
}

// Mutation returns the JupgingLogMutation object of the builder.
func (jlu *JupgingLogUpdate) Mutation() *JupgingLogMutation {
	return jlu.mutation
}

// ClearMember clears the "member" edge to the Member entity.
func (jlu *JupgingLogUpdate) ClearMember() *JupgingLogUpdate {
	jlu.mutation.ClearMember()
	return jlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jlu *JupgingLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, jlu.sqlSave, jlu.mutation, jlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jlu *JupgingLogUpdate) SaveX(ctx context.Context) int {
	affected, err := jlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jlu *JupgingLogUpdate) Exec(ctx context.Context) error {
	_, err := jlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jlu *JupgingLogUpdate) ExecX(ctx context.Context) {
	if err := jlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jlu *JupgingLogUpdate) check() error {
	if jlu.mutation.MemberCleared() && len(jlu.mutation.MemberIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "JupgingLog.member"`)
	}
	return nil
}

func (jlu *JupgingLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := jlu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(jupginglog.Table, jupginglog.Columns, sqlgraph.NewFieldSpec(jupginglog.FieldID, field.TypeInt))
	if ps := jlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jlu.mutation.StartDate(); ok {
		_spec.SetField(jupginglog.FieldStartDate, field.TypeString, value)
	}
	if value, ok := jlu.mutation.EndDate(); ok {
		_spec.SetField(jupginglog.FieldEndDate, field.TypeString, value)
	}
	if value, ok := jlu.mutation.Log(); ok {
		_spec.SetField(jupginglog.FieldLog, field.TypeString, value)
	}
	if jlu.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jupginglog.MemberTable,
			Columns: []string{jupginglog.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jlu.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jupginglog.MemberTable,
			Columns: []string{jupginglog.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, jlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jupginglog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	jlu.mutation.done = true
	return n, nil
}

// JupgingLogUpdateOne is the builder for updating a single JupgingLog entity.
type JupgingLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JupgingLogMutation
}

// SetStartDate sets the "startDate" field.
func (jluo *JupgingLogUpdateOne) SetStartDate(s string) *JupgingLogUpdateOne {
	jluo.mutation.SetStartDate(s)
	return jluo
}

// SetNillableStartDate sets the "startDate" field if the given value is not nil.
func (jluo *JupgingLogUpdateOne) SetNillableStartDate(s *string) *JupgingLogUpdateOne {
	if s != nil {
		jluo.SetStartDate(*s)
	}
	return jluo
}

// SetEndDate sets the "endDate" field.
func (jluo *JupgingLogUpdateOne) SetEndDate(s string) *JupgingLogUpdateOne {
	jluo.mutation.SetEndDate(s)
	return jluo
}

// SetNillableEndDate sets the "endDate" field if the given value is not nil.
func (jluo *JupgingLogUpdateOne) SetNillableEndDate(s *string) *JupgingLogUpdateOne {
	if s != nil {
		jluo.SetEndDate(*s)
	}
	return jluo
}

// SetLog sets the "log" field.
func (jluo *JupgingLogUpdateOne) SetLog(s string) *JupgingLogUpdateOne {
	jluo.mutation.SetLog(s)
	return jluo
}

// SetNillableLog sets the "log" field if the given value is not nil.
func (jluo *JupgingLogUpdateOne) SetNillableLog(s *string) *JupgingLogUpdateOne {
	if s != nil {
		jluo.SetLog(*s)
	}
	return jluo
}

// SetMemberID sets the "member_id" field.
func (jluo *JupgingLogUpdateOne) SetMemberID(i int) *JupgingLogUpdateOne {
	jluo.mutation.SetMemberID(i)
	return jluo
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (jluo *JupgingLogUpdateOne) SetNillableMemberID(i *int) *JupgingLogUpdateOne {
	if i != nil {
		jluo.SetMemberID(*i)
	}
	return jluo
}

// SetMember sets the "member" edge to the Member entity.
func (jluo *JupgingLogUpdateOne) SetMember(m *Member) *JupgingLogUpdateOne {
	return jluo.SetMemberID(m.ID)
}

// Mutation returns the JupgingLogMutation object of the builder.
func (jluo *JupgingLogUpdateOne) Mutation() *JupgingLogMutation {
	return jluo.mutation
}

// ClearMember clears the "member" edge to the Member entity.
func (jluo *JupgingLogUpdateOne) ClearMember() *JupgingLogUpdateOne {
	jluo.mutation.ClearMember()
	return jluo
}

// Where appends a list predicates to the JupgingLogUpdate builder.
func (jluo *JupgingLogUpdateOne) Where(ps ...predicate.JupgingLog) *JupgingLogUpdateOne {
	jluo.mutation.Where(ps...)
	return jluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jluo *JupgingLogUpdateOne) Select(field string, fields ...string) *JupgingLogUpdateOne {
	jluo.fields = append([]string{field}, fields...)
	return jluo
}

// Save executes the query and returns the updated JupgingLog entity.
func (jluo *JupgingLogUpdateOne) Save(ctx context.Context) (*JupgingLog, error) {
	return withHooks(ctx, jluo.sqlSave, jluo.mutation, jluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jluo *JupgingLogUpdateOne) SaveX(ctx context.Context) *JupgingLog {
	node, err := jluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jluo *JupgingLogUpdateOne) Exec(ctx context.Context) error {
	_, err := jluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jluo *JupgingLogUpdateOne) ExecX(ctx context.Context) {
	if err := jluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jluo *JupgingLogUpdateOne) check() error {
	if jluo.mutation.MemberCleared() && len(jluo.mutation.MemberIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "JupgingLog.member"`)
	}
	return nil
}

func (jluo *JupgingLogUpdateOne) sqlSave(ctx context.Context) (_node *JupgingLog, err error) {
	if err := jluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(jupginglog.Table, jupginglog.Columns, sqlgraph.NewFieldSpec(jupginglog.FieldID, field.TypeInt))
	id, ok := jluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JupgingLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jupginglog.FieldID)
		for _, f := range fields {
			if !jupginglog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != jupginglog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jluo.mutation.StartDate(); ok {
		_spec.SetField(jupginglog.FieldStartDate, field.TypeString, value)
	}
	if value, ok := jluo.mutation.EndDate(); ok {
		_spec.SetField(jupginglog.FieldEndDate, field.TypeString, value)
	}
	if value, ok := jluo.mutation.Log(); ok {
		_spec.SetField(jupginglog.FieldLog, field.TypeString, value)
	}
	if jluo.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jupginglog.MemberTable,
			Columns: []string{jupginglog.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jluo.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jupginglog.MemberTable,
			Columns: []string{jupginglog.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &JupgingLog{config: jluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jupginglog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	jluo.mutation.done = true
	return _node, nil
}
