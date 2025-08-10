package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Temple holds the schema definition for the Temple entity.
type Temple struct {
	ent.Schema
}

// Fields of the Temple.
func (Temple) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("寺社名（日本語）").
			NotEmpty(),
		field.String("name_en").
			Comment("寺社名（英語）").
			NotEmpty(),
		field.String("description").
			Comment("寺社の説明").
			Optional(),
		field.String("description_en").
			Comment("寺社の説明（英語）").
			Optional(),
		field.Float("latitude").
			Comment("緯度").
			Positive(),
		field.Float("longitude").
			Comment("経度").
			Positive(),
		field.String("address").
			Comment("住所").
			Optional(),
		field.String("phone").
			Comment("電話番号").
			Optional(),
		field.String("website").
			Comment("公式ウェブサイト").
			Optional(),
		field.String("instagram").
			Comment("Instagramアカウント").
			Optional(),
		field.String("twitter").
			Comment("Twitterアカウント").
			Optional(),
		field.String("opening_hours").
			Comment("開門時間").
			Optional(),
		field.String("goshuin_fee").
			Comment("御朱印料金").
			Optional(),
		field.String("goshuin_office").
			Comment("御朱印所の場所").
			Optional(),
		field.Bool("is_active").
			Comment("アクティブかどうか").
			Default(true),
	}
}

// Edges of the Temple.
func (Temple) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("goshuin_collections", GoshuinCollection.Type).
			Comment("この寺社の御朱印コレクション"),
	}
}
