package def

// An operator is a net of instances defined by a set of instances and connections.
type Instance struct {
	// Delegate
	Delegate bool `json:"delegate" yaml:"delegate"`

	// Specification contains information to specify an operator.
	Specification OperatorSpecification `json:"specification,omitempty" yaml:"specification,omitempty"`

	// Operator contains the actual implementation.
	Operator Operator `json:"operator,omitempty" yaml:"operator,omitempty"`
}

type OperatorSpecification struct {
	// Values specifies the properties used in this operator.
	Values Values `json:"values,omitempty" yaml:"values,omitempty"`

	// Generics specifies the generics used in this operator.
	Generics Generics `json:"generics,omitempty" yaml:"generics,omitempty"`

	// Embedding assigns an operator to each delegate.
	Embedding Embedding `json:"embedding,omitempty" yaml:"embedding,omitempty"`
}

type OperatorType string
type OperatorReference string

type Operator struct {
	// Description of this operator in natural language (English).
	Description string `json:"description" yaml:"description"`

	// Reference to an operator.
	Reference OperatorReference `json:"reference,omitempty" yaml:"reference,omitempty"`

	// Elementary
	Elementary string `json:"elementary,omitempty" yaml:"elementary,omitempty"`

	// Properties is the definition for the structure of the values needed by this operator.
	Properties Properties `json:"properties,omitempty" yaml:"properties,omitempty"`

	// Services
	Services map[ServiceName]*Operation `json:"services" yaml:"services"`

	// Delegates
	Delegates Instances `json:"delegates,omitempty" yaml:"delegates,omitempty"`

	// In case Type == "implementation"

	// Instances is a map of all child instances inside this operator.
	Instances Instances `json:"instances,omitempty" yaml:"instances,omitempty"`

	// Implementations
	Implementation map[ServiceName]*struct {
		// Handles can be referenced in connections.
		// There can be an arbitrary number of handles for each service.
		// The default service is called "main".
		Handles Handles `json:"handles" yaml:"handles"`

		// Connections defines the path all data takes through this operator.
		// It connects the child operators with each other and with this interface.
		Connections Connections `json:"connections,omitempty" yaml:"connections,omitempty"`
	} `json:"implementation,omitempty" yaml:"implementation,omitempty"`
}

type InstanceService struct {
	Delegate InstanceName `json:"delegate" yaml:"instance"`
	Instance InstanceName `json:"instance" yaml:"instance"`
	Service  ServiceName  `json:"service" yaml:"service"`
}

type Handles map[HandleID]*InstanceService

type Values map[string]interface{}

// Embedding specifies how to embed a child operator into this implementation.
type Embedding map[InstanceName]map[ServiceName]*InstanceService

type Properties map[string]*Type

type Instances map[InstanceName]*Instance

type Connections map[ConnectionID]*struct {
	Source struct {
		Handle HandleID `json:"handle" yaml:"handle"`
		Port   InPortID `json:"port" yaml:"port"`
	} `json:"source" yaml:"source"`
	Destination struct {
		Handle HandleID  `json:"handle" yaml:"handle"`
		Port   OutPortID `json:"port" yaml:"port"`
	} `json:"destination" yaml:"destination"`
}

type ServiceName string

type DelegateName string

type ResourceID string

type HandleID string

type InstanceName string

type ConnectionID string

type InPortID string

type OutPortID string
