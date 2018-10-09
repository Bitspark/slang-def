package def

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestType_ProviderWorks(t *testing.T) {
	a := assert.New(t)
	tp := createTestTypeProvider([]string{})

	tt, err := tp.dereferenceType("user")
	a.NoError(err)
	a.Equal("map", tt.Type)
}

func TestType_GoodResolve(t *testing.T) {
	a := assert.New(t)
	tp := createTestTypeProvider([]string{"goodType"})

	tr, err := Type{Type: "goodType"}.Resolve(tp, nil)
	a.NoError(err)
	a.Equal("goodType", tr.Type)
}

func TestType_BadResolve(t *testing.T) {
	a := assert.New(t)
	tp := createTestTypeProvider([]string{})

	tr, err := Type{Type: "badType"}.Resolve(tp, nil)
	a.Error(err)
	a.Equal("", tr.Type)
}

func TestType_Generics(t *testing.T) {
	a := assert.New(t)
	tp := createTestTypeProvider([]string{"boolean"})

	tt, err := tp.dereferenceType("validated")
	a.NoError(err)

	generics := make(map[string]*Type)
	generics["item"] = &Type{Type: "coolType"}

	tt, err = tt.Resolve(tp, generics)
	a.NoError(err)

	a.Equal("coolType", tt.Map["item"].Type)
}

func TestType_ReferenceGenerics(t *testing.T) {
	a := assert.New(t)
	tp := createTestTypeProvider([]string{"boolean", "string"})

	tt, err := tp.dereferenceType("validatedUser")
	a.NoError(err)

	tt, err = tt.Resolve(tp, nil)
	a.NoError(err)

	a.Equal("map", tt.Map["item"].Type)
	a.Equal("string", tt.Map["item"].Map["firstName"].Type)
}

func TestType_InlineSpec(t *testing.T) {
	a := assert.New(t)
	tp := createTestTypeProvider([]string{"boolean", "string", "number"})

	ti1, err := tp.dereferenceType("inlineSpecRef")
	a.NoError(err)

	ti2, err := tp.dereferenceType("expectedInlineSpec")
	a.NoError(err)

	generics := make(map[string]*Type)
	generics["type2"] = &Type{Type: "coolType"}

	ti1, err = ti1.Resolve(tp, generics)
	a.NoError(err)

	a.Equal(ti2, ti1)
}
