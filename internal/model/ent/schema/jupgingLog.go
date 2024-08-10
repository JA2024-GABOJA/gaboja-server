package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JupgingLog holds the schema definition for the JupgingLog entity.
type JupgingLog struct {
	ent.Schema
}

// Fields of the JupgingLog.
func (JupgingLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("startDate"),
		field.String("endDate"),
		field.Text("log"),
		field.Int("member_id"),
	}
}

// Edges of the JupgingLog.
func (JupgingLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("member", Member.Type).
			Ref("jupgingLog").
			Field("member_id").
			Unique().
			Required(),
	}
}
