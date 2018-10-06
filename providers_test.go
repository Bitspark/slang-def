package def

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

// TYPE

type testTypeProvider struct {
	allowedTypes map[string]bool
}

func (t testTypeProvider) getReference(reference string) (Type, error) {
	to := Type{}

	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "types", reference+".yaml"))
	yaml.Unmarshal(typeYAMLBytes, &to)
	return to, err
}

func (t testTypeProvider) getType(reference string) (Type, error) {
	if allowed, ok := t.allowedTypes[reference]; !allowed || !ok {
		return Type{}, fmt.Errorf("bad type '%s'", reference)
	}

	return Type{Type: reference}, nil
}

func createTestTypeProvider(allowedTypes []string) testTypeProvider {
	tp := testTypeProvider{}
	tp.allowedTypes = make(map[string]bool)
	for _, at := range allowedTypes {
		tp.allowedTypes[at] = true
	}
	return tp
}

// OPERATION

type testOperationProvider struct {
}

func (t testOperationProvider) getOperation(reference string) (Operation, error) {
	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "operations", reference+".yaml"))

	to := Operation{}
	yaml.Unmarshal(typeYAMLBytes, &to)

	return to, err
}

func createTestOperationProvider() testOperationProvider {
	op := testOperationProvider{}
	return op
}

// RESOURCE

type testResourceProvider struct {
}

func (t testResourceProvider) get(reference string) (Resource, error) {
	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "resources", reference+".yaml"))

	to := Resource{}
	yaml.Unmarshal(typeYAMLBytes, &to)

	return to, err
}

// OPERATOR

type builder func(Generics, Values, Embedding) (Operator, error)

type testOperatorProvider struct {
	builders map[string]builder
}

func (t testOperatorProvider) get(reference string, generics Generics, values Values, embedding Embedding) (Operator, error) {
	if builder, ok := t.builders[reference]; ok {
		return builder(generics, values, embedding)
	}

	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "operators", reference+".yaml"))

	to := Operator{}
	yaml.Unmarshal(typeYAMLBytes, &to)

	return to, err
}