package def

// A resource is an interface to the system an operator runs on.
type Resource struct {

	/* === REFERENCE + SPECIFICATION === */

	// OperatorSpecification is a reference to the definition of this resource.
	Reference string `json:"reference" yaml:"reference"`

	// Generics specifies the generics used by this resource.
	Generics Generics `json:"generics,omitempty" yaml:"generics,omitempty"`

	// Values specifies the properties used by this resource.
	Values Values `json:"values,omitempty" yaml:"values,omitempty"`

	/* === DEFINITION === */

	// Description of this resource in natural language (English).
	Description string `json:"description" yaml:"description"`

	// Properties is the definition for the structure of the values needed by this operator.
	Properties Properties `json:"properties,omitempty" yaml:"properties,omitempty"`

	// Services are operation implemented by the resource
	Services map[string]*Operation `json:"services" yaml:"services"`

	// Events are operations to be implemented by Slang
	Events map[string]*Operation `json:"events" yaml:"events"`

}

// Resolves the type using the given provider.
func (r Resource) Resolve(resourceProvider ResourceProvider, operationProvider OperationProvider, typeProvider TypeProvider, generics Generics) (Resource, error) {
	resolved := Resource{
		Description: r.Description,
	}

	if r.Reference != "" {
		var err error
		ref, err := resourceProvider.getResourceRef(r.Reference)
		if err != nil {
			return Resource{}, err
		}
		refGens := make(Generics)
		for gen, genType := range r.Generics {
			resolvedGen, err := genType.Resolve(typeProvider, generics)
			if err != nil {
				return Resource{}, err
			}
			refGens[gen] = &resolvedGen
		}
		resolved, err = ref.Resolve(resourceProvider, operationProvider, typeProvider, refGens)
		if err != nil {
			return Resource{}, err
		}
		return resolved, nil
	}

	resolvedServices := make(map[string]*Operation)
	for operation, operationDef := range r.Services {
		resolvedOperation, err := operationDef.Resolve(operationProvider, typeProvider, generics)
		if err != nil {
			return Resource{}, err
		}
		resolvedServices[operation] = &resolvedOperation
	}
	resolved.Services = resolvedServices

	resolvedEvents := make(map[string]*Operation)
	for operation, operationDef := range r.Events {
		resolvedOperation, err := operationDef.Resolve(operationProvider, typeProvider, generics)
		if err != nil {
			return Resource{}, err
		}
		resolvedEvents[operation] = &resolvedOperation
	}
	resolved.Services = resolvedEvents

	return resolved, nil
}