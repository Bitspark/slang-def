package def

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperation_ProviderWorks(t *testing.T) {
	a := assert.New(t)
	op := createTestOperationProvider()

	to, err := op.dereferenceOperation("compare")
	a.NoError(err)
	a.Equal("generic", to.In["a"].Type)
	a.Equal("generic", to.In["a"].Type)
}

func TestOperation_ResolveType(t *testing.T) {
	a := assert.New(t)
	op := createTestOperationProvider()
	tp := createTestTypeProvider([]string{"string", "boolean"})

	to, err := op.dereferenceOperation("validateEmail")
	a.NoError(err)

	to, err = to.Resolve(op, tp, nil)
	a.NoError(err)
	a.Equal("map", to.In["email"].Type)
	a.Equal("boolean", to.Out["valid"].Type)
}

func TestOperation_ResolveOperation(t *testing.T) {
	a := assert.New(t)
	op := createTestOperationProvider()
	tp := createTestTypeProvider([]string{"string", "boolean", "coolType"})

	generics := make(map[string]*Type)
	generics["item"] = &Type{Type: "coolType"}

	to := Operation{Reference: "compare", Generics: generics}

	var err error
	to, err = to.Resolve(op, tp, generics)
	a.NoError(err)
	a.Equal("coolType", to.In["a"].Type)
	a.Equal("coolType", to.In["b"].Type)
	a.Equal("boolean", to.Out["order"].Type)
}

func TestOperation_ResolveGeneric(t *testing.T) {
	a := assert.New(t)
	op := createTestOperationProvider()
	tp := createTestTypeProvider([]string{"string", "boolean", "coolType"})

	to := Operation{Reference: "compareUser", Generics: nil}

	var err error
	to, err = to.Resolve(op, tp, nil)
	a.NoError(err)
	a.Equal("map", to.In["a"].Type)
	a.Equal("map", to.In["b"].Type)
	a.Equal("boolean", to.Out["order"].Type)
}
