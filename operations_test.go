package def

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperation_ProviderWorks(t *testing.T) {
	a := assert.New(t)
	op := createTestOperationProvider()

	to, err := op.getOperation("compare")
	a.NoError(err)
	a.Equal("map", to.In.Type)
}

func TestOperation_ResolveType(t *testing.T) {
	a := assert.New(t)
	op := createTestOperationProvider()
	tp := createTestTypeProvider([]string{"string", "boolean"})

	to, err := op.getOperation("validateEmail")
	a.NoError(err)

	to, err = to.Resolve(op, tp, nil)
	a.NoError(err)
	a.Equal("map", to.In.Type)
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
	a.Equal("map", to.In.Type)
}
