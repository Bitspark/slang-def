package def

type Resources map[string]*Resource
type Embedding map[string]string

type Operators []*Operator
type Connections map[string][]string

// An operator is a net of instances defined by a set of instances and connections.
type Operator struct {

	/* === REFERENCE + SPECIFICATION === */

	// Reference is a reference to the implementation of this operator.
	// If it is provided, Operators and Connections must not be specified.
	Reference string `json:"reference" yaml:"reference"`

	// Generics specifies the generics used in this operator.
	Generics Generics `json:"generics,omitempty" yaml:"generics,omitempty"`

	// Values specifies the properties used in this operator.
	Values Values `json:"values,omitempty" yaml:"values,omitempty"`

	// Embedding specifies the resource usage of this operator.
	Embedding Embedding `json:"embedding,omitempty" yaml:"embedding,omitempty"`

	/* === DEFINITION === */

	// Description of the implementation in this operator
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Operation is the operation implemented by this operator.
	Operation Operation `json:"operation" yaml:"operation"`

	// Properties is the definition for the structure of the values needed by this operator.
	Properties Properties `json:"properties,omitempty" yaml:"properties,omitempty"`

	// Resources is the definition of the resources this operator depends on.
	Resources Resources `json:"resources,omitempty" yaml:"resources,omitempty"`

	// Operators is a map of all child operators inside this operator.
	Operators Operators `json:"operators,omitempty" yaml:"operators,omitempty"`

	// Connections defines the path all data takes through this operator.
	// It connects the child operators with each other and with this interface.
	Connections Connections `json:"connections,omitempty" yaml:"connections,omitempty"`

}
