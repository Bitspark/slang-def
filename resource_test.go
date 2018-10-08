package def

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResource_ProviderWorks(t *testing.T) {
	a := assert.New(t)
	rp := createTestResourceProvider()

	tr, err := rp.getResourceRef("filesystem")
	a.NoError(err)
	a.Equal(2, len(tr.Operations))
}

func TestResource_ResolveType(t *testing.T) {
	/*a := assert.New(t)
	op := createTestOperationProvider()
	tp := createTestTypeProvider([]string{"string", "boolean"})

	to, err := op.getOperationRef("validateEmail")
	a.NoError(err)

	to, err = to.Resolve(op, tp, nil)
	a.NoError(err)
	a.Equal("map", to.In.Type)*/
}

func TestResource_ResolveOperation(t *testing.T) {
	/*a := assert.New(t)
	op := createTestOperationProvider()
	tp := createTestTypeProvider([]string{"string", "boolean", "coolType"})

	generics := make(map[string]*Type)
	generics["item"] = &Type{Type: "coolType"}

	to := Operation{OperatorSpecification: "compare", Generics: generics}

	var err error
	to, err = to.Resolve(op, tp, generics)
	a.NoError(err)
	a.Equal("map", to.In.Type)*/
}

func TestResource_ResolveGeneric(t *testing.T) {
	/*a := assert.New(t)
	op := createTestOperationProvider()
	tp := createTestTypeProvider([]string{"string", "boolean", "coolType"})

	to := Operation{OperatorSpecification: "compareUser", Generics: nil}

	var err error
	to, err = to.Resolve(op, tp, nil)
	a.NoError(err)
	a.Equal("map", to.In.Type)*/
}
