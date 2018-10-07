package def

type Properties map[string]*Type

// An operation defines an interface and has a description of its semantics in natural language (English).
type Operation struct {

	/* === REFERENCE + SPECIFICATION === */

	// Reference specifies a named operation.
	Reference string `json:"reference,omitempty" yaml:"reference,omitempty"`

	// Generics specify generics used in the operation referenced in Reference.
	Generics Generics `json:"generics,omitempty" yaml:"generics,omitempty"`

	/* === DEFINITION === */

	// Description of the semantics of this operation
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// In is the definition of the structure of the data this operation accepts
	In Type `json:"in,omitempty" yaml:"in,omitempty"`

	// Out is the definition of the structure of the data this operation returns
	Out Type `json:"out,omitempty" yaml:"out,omitempty"`

}

// Resolves the type using the given provider.
func (o Operation) Resolve(operationProvider OperationProvider, typeProvider TypeProvider, generics Generics) (Operation, error) {
	var err error
	resolved := o

	if o.Reference != "" {
		var ref Operation
		ref, err = operationProvider.getOperationRef(o.Reference)
		if err != nil {
			return Operation{}, err
		}
		refGens := make(Generics)
		for gen, genType := range o.Generics {
			resolvedGen, err := genType.Resolve(typeProvider, generics)
			if err != nil {
				return Operation{}, err
			}
			refGens[gen] = &resolvedGen
		}
		resolved, err = ref.Resolve(operationProvider, typeProvider, refGens)
		if err != nil {
			return Operation{}, err
		}
		return resolved, nil
	}

	resolved.In, err = resolved.In.Resolve(typeProvider, generics)
	if err != nil {
		return Operation{}, err
	}

	resolved.Out, err = resolved.Out.Resolve(typeProvider, generics)
	if err != nil {
		return Operation{}, err
	}

	return resolved, nil
}
