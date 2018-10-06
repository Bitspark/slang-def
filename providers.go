package def

// An OperationProvider provides operation definitions.
type OperationProvider interface {
	getOperation(reference string) (Operation, error)
}

// A TypeProvider provides type definitions.
type TypeProvider interface {
	getType(name string) (Type, error)
	getReference(reference string) (Type, error)
}

// A ResourceProvider provides resource definitions.
type ResourceProvider interface {
	getResource(reference string) (Resource, error)
}

// An OperatorProvider provides operator definitions.
type OperatorProvider interface {
	getOperator(reference string, generics Generics, values Values, embedding Embedding) (Operator, error)
}
