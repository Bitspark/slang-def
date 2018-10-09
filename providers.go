package def

// A TypeProvider provides type definitions.
type TypeProvider interface {
	dereferenceType(reference string) (Type, error)
	buildType(name string) (Type, error)
}

// An OperationProvider provides operation definitions.
type OperationProvider interface {
	dereferenceOperation(reference string) (Operation, error)
}

// An OperatorProvider provides operator definitions.
type OperatorProvider interface {
	dereferenceOperator(reference string) (Instance, error)
	buildOperator(reference string, generics Generics, values Values, embedding Embedding) (Instance, error)
}
