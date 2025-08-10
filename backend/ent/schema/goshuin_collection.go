package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// GoshuinCollection holds the schema definition for the GoshuinCollection entity.
type GoshuinCollection struct {
	ent.Schema
}

// Fields of the GoshuinCollection.
func (GoshuinCollection) Fields() []ent.Field {
	return []ent.Field{
		field.Int("temple_id").
			Comment("寺社ID").
			Positive(),
		field.String("image_url").
			Comment("御朱印の画像URL").
			Optional(),
		field.String("notes").
			Comment("メモ").
			Optional(),
		field.Time("collected_at").
			Comment("収集日時").
			Default(time.Now),
		field.Time("created_at").
			Comment("作成日時").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Comment("更新日時").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the GoshuinCollection.
func (GoshuinCollection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("temple", Temple.Type).
			Ref("goshuin_collections").
			Field("temple_id").
			Unique().
			Required().
			Comment("この御朱印が属する寺社"),
	}
}
