// Code generated by ent, DO NOT EDIT.

package member

import (
	"junction/internal/model/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldID, id))
}

// Sno applies equality check predicate on the "sno" field. It's identical to SnoEQ.
func Sno(v int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldSno, v))
}

// SnoEQ applies the EQ predicate on the "sno" field.
func SnoEQ(v int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldSno, v))
}

// SnoNEQ applies the NEQ predicate on the "sno" field.
func SnoNEQ(v int) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldSno, v))
}

// SnoIn applies the In predicate on the "sno" field.
func SnoIn(vs ...int) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldSno, vs...))
}

// SnoNotIn applies the NotIn predicate on the "sno" field.
func SnoNotIn(vs ...int) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldSno, vs...))
}

// SnoGT applies the GT predicate on the "sno" field.
func SnoGT(v int) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldSno, v))
}

// SnoGTE applies the GTE predicate on the "sno" field.
func SnoGTE(v int) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldSno, v))
}

// SnoLT applies the LT predicate on the "sno" field.
func SnoLT(v int) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldSno, v))
}

// SnoLTE applies the LTE predicate on the "sno" field.
func SnoLTE(v int) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldSno, v))
}

// HasJupgingLog applies the HasEdge predicate on the "jupgingLog" edge.
func HasJupgingLog() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, JupgingLogTable, JupgingLogColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJupgingLogWith applies the HasEdge predicate on the "jupgingLog" edge with a given conditions (other predicates).
func HasJupgingLogWith(preds ...predicate.JupgingLog) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := newJupgingLogStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Member) predicate.Member {
	return predicate.Member(sql.NotPredicates(p))
}
