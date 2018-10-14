package def

import (
	"errors"
	"fmt"
)

// Type defines structure for data.
type Type struct {
	// Type is either "stream", "map", "generic", "reference" or a type the target system understands.
	// Normally, the following types should be available: "string", "number", "boolean", "trigger".
	Type string `json:"type" yaml:"type"`

	/* === REFERENCE + SPECIFICATION (only in case Type == "reference") === */

	// Reference is the name of a referenced type.
	Reference string `json:"reference,omitempty" yaml:"reference,omitempty"`

	// Generics is a specification of the generics inside a referenced type.
	Generics Generics `json:"generics,omitempty" yaml:"generics,omitempty"`

	/* === DEFINITION === */

	// Description of this type in natural language (English).
	Description string `json:"description" yaml:"description"`

	// Stream is the underlying type for the stream.
	// It is ignored in case Type is not "stream".
	Stream *Type `json:"stream,omitempty" yaml:"stream,omitempty"`

	// Map is a map of underlying types.
	// It is ignored in case Type is not "map".
	Map Map `json:"map,omitempty" yaml:"map,omitempty"`

	// Generic is a placeholder for an arbitrary type.
	// It is ignored in case Type is not "generic".
	Generic string `json:"generic,omitempty" yaml:"generic,omitempty"`
}

type Map map[string]*Type

// Generics is a specification of generics.
type Generics map[string]*Type

// Resolves the type using the given provider.
func (t Type) Resolve(typeProvider TypeProvider, generics Generics) (Type, error) {
	resolved := Type{
		Description: t.Description,
		Type:        t.Type,
	}

	if t.Type == "stream" {
		resolvedStream, err := t.Stream.Resolve(typeProvider, generics)
		if err != nil {
			return Type{}, err
		}
		resolved.Stream = &resolvedStream
	} else if t.Type == "map" {
		resolvedMap := make(map[string]*Type)
		for sub, subtype := range t.Map {
			resolvedSubtype, err := subtype.Resolve(typeProvider, generics)
			if err != nil {
				return Type{}, err
			}
			resolvedMap[sub] = &resolvedSubtype
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
			resolvedGen, err := genType.Resolve(typeProvider, generics)
			if err != nil {
				return Type{}, err
			}
			refGens[gen] = &resolvedGen
		}
		resolved, err = ref.Resolve(typeProvider, refGens)
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

func (t Map) Resolve(typeProvider TypeProvider, generics Generics) (Map, error) {
	newMap := make(Map)
	for name, sub := range t {
		newSub, err := sub.Resolve(typeProvider, generics)
		if err != nil {
			return nil, err
		}
		newMap[name] = &newSub
	}
	return newMap, nil
}

// Validate returns false iff the type uses unknown types or is malformed
func (t Type) Validate() error {
	if t.Type == "stream" {
		if t.Stream == nil {
			return errors.New("stream not defined")
		}
		return t.Stream.Validate()
	}

	if t.Type == "map" {
		if t.Map == nil {
			return errors.New("map not defined")
		}
		for _, subtype := range t.Map {
			err := subtype.Validate()
			if err != nil {
				return err
			}
		}
		return nil
	}

	if t.Type == "generic" {
		if t.Generic == "" {
			return errors.New("generic not defined")
		}
		return nil
	}

	if t.Type == "reference" {
		if t.Reference == "" {
			return errors.New("reference not defined")
		}
		return nil
	}

	return nil
}
