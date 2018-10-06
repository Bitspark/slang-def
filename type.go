package def

import (
	"errors"
	"fmt"
)

// Generics is a specification of generics.
type Generics map[string]*Type

// Values is a specification of properties.
type Values map[string]interface{}

// Type defines structure for data.
type Type struct {
	// Description of this type in natural language (English).
	Description string `json:"description" yaml:"description"`

	// Type is either "stream", "map", "generic", "reference" or a type the target system understands.
	// Normally, the following types should be available: "string", "number", "boolean", "trigger".
	Type string `json:"type" yaml:"type"`

	// Stream is the underlying type for the stream.
	// It is ignored in case Type is not "stream".
	Stream *Type `json:"stream,omitempty" yaml:"stream,omitempty"`

	// Map is a map of underlying types.
	// It is ignored in case Type is not "map".
	Map map[string]*Type `json:"map,omitempty" yaml:"map,omitempty"`

	// Generic is a placeholder for an arbitrary type.
	// It is ignored in case Type is not "generic".
	Generic string `json:"generic,omitempty" yaml:"generic,omitempty"`

	// Reference is a placeholder for an arbitrary named type.
	// It is ignored in case Type is not "reference".
	Reference string `json:"reference,omitempty" yaml:"reference,omitempty"`

	// Generics is a specification of the generics inside a reference.
	// It is ignored in case Type is not "reference".
	Generics Generics `json:"generics,omitempty" yaml:"generics,omitempty"`
}

// Resolves the type using the given provider.
func (t Type) Resolve(provider TypeProvider, generics Generics) (Type, error) {
	resolved := Type{
		Description: t.Description,
		Type:        t.Type,
	}

	if t.Type == "stream" {
		resolvedStream, err := t.Stream.Resolve(provider, generics)
		if err != nil {
			return Type{}, err
		}
		resolved.Stream = &resolvedStream
	} else if t.Type == "map" {
		resolvedMap := make(map[string]*Type)
		for sub, subtype := range t.Map {
			resolvedSubtype, err := subtype.Resolve(provider, generics)
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
		ref, err := provider.getReference(t.Reference)
		if err != nil {
			return Type{}, err
		}
		refGens := make(Generics)
		for gen, genType := range t.Generics {
			resolvedGen, err := genType.Resolve(provider, generics)
			if err != nil {
				return Type{}, err
			}
			refGens[gen] = &resolvedGen
		}
		resolved, err = ref.Resolve(provider, refGens)
		if err != nil {
			return Type{}, err
		}
	} else {
		var err error
		resolved, err = provider.getType(t.Type)
		if err != nil {
			return Type{}, err
		}
	}

	return resolved, nil
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