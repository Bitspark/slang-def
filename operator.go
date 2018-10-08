package def

// An operator is a net of instances defined by a set of instances and connections.
type Operator struct {
	// SPECIFICATION

	// Type is either "reference", "slang", "native" or "resource".
	Type OperatorType `json:"type" yaml:"type"`

	// OperatorSpecification contains information to reference and specify an operator.
	Specification OperatorSpecification `json:"specification,omitempty" yaml:"specification,omitempty"`

	// IMPLEMENTATION

	// Reference is a reference to the implementation of this operator.
	Reference OperatorReference `json:"reference" yaml:"reference"`

	// SlangOperator contains the actual implementation.
	SlangOperator SlangOperator `json:"slang,omitempty" yaml:"implementation,omitempty"`

	// SlangOperator contains the actual implementation.
	NativeOperator NativeOperator `json:"native,omitempty" yaml:"implementation,omitempty"`

	// SlangOperator contains the actual implementation.
	ResourceOperator ResourceOperator `json:"resource,omitempty" yaml:"implementation,omitempty"`

	// HANDLES

	// Handles can be referenced in connections
	Handles Handles `json:"handles" yaml:"handles"`
}

type OperatorSpecification struct {
	// Generics specifies the generics used in this operator.
	Generics Generics `json:"generics,omitempty" yaml:"generics,omitempty"`

	// Values specifies the properties used in this operator.
	Values Values `json:"values,omitempty" yaml:"values,omitempty"`

	// Embedding assigns an operator to each delegate.
	Embedding Embedding `json:"embedding,omitempty" yaml:"embedding,omitempty"`
}

type OperatorType string
type OperatorReference string

type SlangOperator struct {
	// Operation is the operation to be implemented by this operator.
	Operation Operation `json:"operation" yaml:"operation"`

	// Description of this operator in natural language (English).
	Description string `json:"description" yaml:"description"`

	// Resources is the list of the resources this operator depends on.
	Resources Resources `json:"resources,omitempty" yaml:"resources,omitempty"`

	// Properties is the definition for the structure of the values needed by this operator.
	Properties Properties `json:"properties,omitempty" yaml:"properties,omitempty"`

	// Operators is a map of all child operators inside this operator.
	Operators Operators `json:"operators,omitempty" yaml:"operators,omitempty"`

	// Connections defines the path all data takes through this operator.
	// It connects the child operators with each other and with this interface.
	Connections Connections `json:"connections,omitempty" yaml:"connections,omitempty"`
}

type NativeOperator struct {
	Native string `json:"native" yaml:"native"`
}

type ResourceOperator struct {
	Resource ResourceID `json:"resource" yaml:"resource"`

	Service string `json:"service" yaml:"service"`
}

type Handles []HandleID

type Values map[string]interface{}

// Embedding specifies how to embed a child operator into this implementation.
type Embedding struct {
	// OperatorMap maps an operator without implementation to an operator in this implementation.
	OperatorMap OperatorMap `json:"operatorMap" yaml:"operatorMap"`

	// ResourceMap maps resources in the operator to a resource in this implementation.
	ResourceMap ResourceMap `json:"resourceMap" yaml:"resourceMap"`
}

type Resources map[ResourceID]*ResourceInstance

type Properties map[string]*Type

type Operators map[OperatorID]*Operator

type Connections map[ConnectionID]*struct {
	From struct {
		Handle HandleID
		Port   InPortID
	}

	To struct {
		Handle HandleID
		Port   OutPortID
	}
}

type ResourceID string

type HandleID string

type OperatorMap map[OperatorID]OperatorID

type ResourceMap map[ResourceID]ResourceID

type ResourceInstance struct {
	Resource string `json:"resource" yaml:"resource"`

	Embedding OperatorMap `json:"embedding" yaml:"embedding"`
}

type OperatorID string

type ConnectionID string

type InPortID string

type OutPortID string
