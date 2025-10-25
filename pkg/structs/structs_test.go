package structs

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/structs/notation_tag"
	"github.com/funtimecoding/go-library/pkg/structs/relational_tag"
	"reflect"
	"testing"
)

type Fixture1 struct {
	Name string `json:"name" gorm:"primaryKey"`
}

type Fixture2 struct {
	Age int `json:"age,omitempty" gorm:"not null"`
}

type Fixture3 struct {
	Mail string `json:"mail" gorm:"not null;default:jdoe@example.org"`
}

func TestFixture1(t *testing.T) {
	v := reflect.TypeOf(Fixture1{})
	assert.Integer(t, 1, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		n := notation_tag.FromField(f)
		assert.String(t, "name", n.Key())
		assert.False(t, n.OmitEmpty())

		r := relational_tag.FromField(f)
		assert.True(t, r.Primary())
		assert.True(t, !r.Nullable())
	}
}

func TestFixture2(t *testing.T) {
	v := reflect.TypeOf(Fixture2{})
	assert.Integer(t, 1, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		n := notation_tag.FromField(f)
		assert.String(t, "age", n.Key())
		assert.True(t, n.OmitEmpty())

		r := relational_tag.FromField(f)
		assert.False(t, r.Primary())
		assert.True(t, !r.Nullable())
	}
}

func TestFixture3(t *testing.T) {
	v := reflect.TypeOf(Fixture3{})
	assert.Integer(t, 1, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		n := notation_tag.FromField(f)
		assert.String(t, "mail", n.Key())
		assert.False(t, n.OmitEmpty())

		r := relational_tag.FromField(f)
		assert.False(t, r.Primary())
		assert.True(t, !r.Nullable())
		assert.String(t, "jdoe@example.org", r.Default())
	}
}
