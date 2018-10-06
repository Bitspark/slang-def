package def

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProviderWorks(t *testing.T) {
	a := assert.New(t)
	tp := createTestTypeProvider([]string{})

	tt, err := tp.getReference("user")
	a.NoError(err)
	a.Equal("map", tt.Type)
}

func TestGoodResolve(t *testing.T) {
	a := assert.New(t)
	tp := createTestTypeProvider([]string{"goodType"})

	tr, err := Type{Type: "goodType"}.Resolve(tp, nil)
	a.NoError(err)
	a.Equal("goodType", tr.Type)
}

func TestBadResolve(t *testing.T) {
	a := assert.New(t)
	tp := createTestTypeProvider([]string{})

	tr, err := Type{Type: "badType"}.Resolve(tp, nil)
	a.Error(err)
	a.Equal("", tr.Type)
}

func TestGenerics(t *testing.T) {
	a := assert.New(t)
	tp := createTestTypeProvider([]string{"boolean"})

	tt, err := tp.getReference("validated")
	a.NoError(err)

	generics := make(map[string]*Type)
	generics["item"] = &Type{Type: "coolType"}

	tt, err = tt.Resolve(tp, generics)
	a.NoError(err)

	a.Equal("coolType", tt.Map["item"].Type)
}

func TestReferenceGenerics(t *testing.T) {
	a := assert.New(t)
	tp := createTestTypeProvider([]string{"boolean", "string"})

	tt, err := tp.getReference("validatedUser")
	a.NoError(err)

	tt, err = tt.Resolve(tp, nil)
	a.NoError(err)

	a.Equal("map", tt.Map["item"].Type)
	a.Equal("string", tt.Map["item"].Map["firstName"].Type)
}