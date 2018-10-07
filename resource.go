package def

// A resource is an interface to the system an operator runs on.
type Resource struct {

	/* === REFERENCE === */

	// Reference is a reference to the definition of this resource.
	Reference string `json:"reference" yaml:"reference"`

	/* === DEFINITION === */

	// Description of this resource in natural language (English).
	Description string `json:"description" yaml:"description"`

	// Operations is a map of operations this resource provides.
	Operations map[string]*Operation `json:"operations" yaml:"operations"`

}

// Resolves the type using the given provider.
func (r Resource) Resolve(resourceProvider ResourceProvider, operationProvider OperationProvider, typeProvider TypeProvider, generics Generics) (Resource, error) {
	resolved := Resource{
		Description: r.Description,
	}

	if r.Reference == "" {
		resolved.Operations = r.Operations
	} else {
		var err error
		resolved, err = resourceProvider.getResourceRef(r.Reference)
		if err != nil {
			return Resource{}, err
		}
	}

	resolvedOperations := make(map[string]*Operation)
	for operation, operationDef := range r.Operations {
		resolvedOperation, err := operationDef.Resolve(operationProvider, typeProvider, generics)
		if err != nil {
			return Resource{}, err
		}
		resolvedOperations[operation] = &resolvedOperation
	}
	resolved.Operations = resolvedOperations

	return resolved, nil
}