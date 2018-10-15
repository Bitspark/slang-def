package def

import (
	"fmt"
)

// Type defines structure for data.
type Type struct {
	// The name of the role this type has in its context.
	// Can make use of properties.
	Name string `json:"name" yaml:"name"`

	// Type is either "stream", "map", "generic", "reference" or a type the target system understands.
	// Normally, the following types should be available: "string", "number", "boolean", "trigger".
	// Cannot (!) make use of properties.
	Type string `json:"type" yaml:"type"`

	/* === REFERENCE + SPECIFICATION (only in case Type == "reference") === */

	// Reference is the name of a referenced type.
	// Cannot (!) make use of properties.
	Reference string `json:"reference,omitempty" yaml:"reference,omitempty"`

	// Generics is a specification of the generics inside a referenced type.
	Generics Generics `json:"generics,omitempty" yaml:"generics,omitempty"`

	// Values is a specification of the properties inside a referenced type.
	Values Values `json:"values,omitempty" yaml:"values,omitempty"`

	/* === DEFINITION === */

	// Properties of this type and its subtypes.
	// Cannot (!) make use of properties.
	Properties Properties `json:"properties" yaml:"properties"`

	// Description of this type in natural language (English).
	// Can make use of properties.
	Description string `json:"description" yaml:"description"`

	// Stream is the underlying type for the stream.
	// It is ignored in case Type is not "stream".
	Stream *Type `json:"stream,omitempty" yaml:"stream,omitempty"`

	// Map is a map of underlying types.
	// It is ignored in case Type is not "map".
	Map []*Type `json:"map,omitempty" yaml:"map,omitempty"`

	// MapTemplate specifies a map being built from a property.
	// It is ignored in case Type is not "map".
	MapTemplate MapTemplate `json:"mapTemplate,omitempty" yaml:"mapTemplate,omitempty"`

	// Generic is a placeholder for an arbitrary type.
	// It is ignored in case Type is not "generic".
	// Can make use of properties.
	Generic string `json:"generic,omitempty" yaml:"generic,omitempty"`
}

// MapTemplate
type MapTemplate struct {
	// Property, must be a stream
	Property string `json:"property" yaml:"property"`
	Index    string `json:"index" yaml:"index"`
	Value    string `json:"value" yaml:"value"`
	Expand   *Type  `json:"expand" yaml:"expand"`

	Entry *Type `json:"entry" yaml:"entry"`
}

// Generics is a specification of generics.
type Generics map[string]*Type

// Resolves the type using the given provider.
func (t Type) Resolve(typeProvider TypeProvider, generics Generics, values Values) (Type, error) {
	resolved := Type{
		Description: t.Description,
		Type:        t.Type,
	}

	if t.Type == "stream" {
		resolvedStream, err := t.Stream.Resolve(typeProvider, generics, values)
		if err != nil {
			return Type{}, err
		}
		resolved.Stream = &resolvedStream
	} else if t.Type == "map" {
		resolvedMap := []*Type{}
		for _, subtype := range t.Map {
			resolvedSubtype, err := subtype.Resolve(typeProvider, generics, values)
			if err != nil {
				return Type{}, err
			}
			resolvedMap = append(resolvedMap, &resolvedSubtype)
		}
		resolved.Map = resolvedMap
	} else if t.Type == "generic" {
		gen, ok := generics[t.Generic]
		if !ok {
			return Type{}, fmt.Errorf("generic '%s' not specified", t.Generic)
		}
		resolved = *gen
	} else if t.Type == "reference" {
		ref, err := typeProvider.dereferenceType(t.Reference)
		if err != nil {
			return Type{}, err
		}
		refGens := make(Generics)
		for gen, genType := range t.Generics {
			resolvedGen, err := genType.Resolve(typeProvider, generics, values)
			if err != nil {
				return Type{}, err
			}
			refGens[gen] = &resolvedGen
		}
		resolved, err = ref.Resolve(typeProvider, refGens, values)
		if err != nil {
			return Type{}, err
		}
	} else {
		var err error
		resolved, err = typeProvider.buildType(t.Type)
		if err != nil {
			return Type{}, err
		}
	}

	return resolved, nil
}
