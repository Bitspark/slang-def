package def

// A TypeProvider provides type definitions.
type TypeProvider interface {
	getTypeRef(reference string) (Type, error)
	getType(name string) (Type, error)
}

// An OperationProvider provides operation definitions.
type OperationProvider interface {
	getOperationRef(reference string) (Operation, error)
}

// A ResourceProvider provides resource definitions.
type ResourceProvider interface {
	getResourceRef(reference string) (Resource, error)
}

// An OperatorProvider provides operator definitions.
type OperatorProvider interface {
	getOperatorRef(reference string) (Operator, error)
	getOperator(reference string, generics Generics, values Values, embedding Embedding) (Operator, error)
}
