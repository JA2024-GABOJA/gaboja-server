// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"junction/internal/model/ent/jupginglog"
	"junction/internal/model/ent/member"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// JupgingLog is the model entity for the JupgingLog schema.
type JupgingLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StartDate holds the value of the "startDate" field.
	StartDate string `json:"startDate,omitempty"`
	// EndDate holds the value of the "endDate" field.
	EndDate string `json:"endDate,omitempty"`
	// Log holds the value of the "log" field.
	Log string `json:"log,omitempty"`
	// MemberID holds the value of the "member_id" field.
	MemberID int `json:"member_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JupgingLogQuery when eager-loading is set.
	Edges        JupgingLogEdges `json:"edges"`
	selectValues sql.SelectValues
}

// JupgingLogEdges holds the relations/edges for other nodes in the graph.
type JupgingLogEdges struct {
	// Member holds the value of the member edge.
	Member *Member `json:"member,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MemberOrErr returns the Member value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JupgingLogEdges) MemberOrErr() (*Member, error) {
	if e.Member != nil {
		return e.Member, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: member.Label}
	}
	return nil, &NotLoadedError{edge: "member"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*JupgingLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case jupginglog.FieldID, jupginglog.FieldMemberID:
			values[i] = new(sql.NullInt64)
		case jupginglog.FieldStartDate, jupginglog.FieldEndDate, jupginglog.FieldLog:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the JupgingLog fields.
func (jl *JupgingLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case jupginglog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			jl.ID = int(value.Int64)
		case jupginglog.FieldStartDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field startDate", values[i])
			} else if value.Valid {
				jl.StartDate = value.String
			}
		case jupginglog.FieldEndDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field endDate", values[i])
			} else if value.Valid {
				jl.EndDate = value.String
			}
		case jupginglog.FieldLog:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field log", values[i])
			} else if value.Valid {
				jl.Log = value.String
			}
		case jupginglog.FieldMemberID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				jl.MemberID = int(value.Int64)
			}
		default:
			jl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the JupgingLog.
// This includes values selected through modifiers, order, etc.
func (jl *JupgingLog) Value(name string) (ent.Value, error) {
	return jl.selectValues.Get(name)
}

// QueryMember queries the "member" edge of the JupgingLog entity.
func (jl *JupgingLog) QueryMember() *MemberQuery {
	return NewJupgingLogClient(jl.config).QueryMember(jl)
}

// Update returns a builder for updating this JupgingLog.
// Note that you need to call JupgingLog.Unwrap() before calling this method if this JupgingLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (jl *JupgingLog) Update() *JupgingLogUpdateOne {
	return NewJupgingLogClient(jl.config).UpdateOne(jl)
}

// Unwrap unwraps the JupgingLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (jl *JupgingLog) Unwrap() *JupgingLog {
	_tx, ok := jl.config.driver.(*txDriver)
	if !ok {
		panic("ent: JupgingLog is not a transactional entity")
	}
	jl.config.driver = _tx.drv
	return jl
}

// String implements the fmt.Stringer.
func (jl *JupgingLog) String() string {
	var builder strings.Builder
	builder.WriteString("JupgingLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", jl.ID))
	builder.WriteString("startDate=")
	builder.WriteString(jl.StartDate)
	builder.WriteString(", ")
	builder.WriteString("endDate=")
	builder.WriteString(jl.EndDate)
	builder.WriteString(", ")
	builder.WriteString("log=")
	builder.WriteString(jl.Log)
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(fmt.Sprintf("%v", jl.MemberID))
	builder.WriteByte(')')
	return builder.String()
}

// JupgingLogs is a parsable slice of JupgingLog.
type JupgingLogs []*JupgingLog
