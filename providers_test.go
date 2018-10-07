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

func (t testTypeProvider) getTypeRef(reference string) (Type, error) {
	to := Type{}

	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "types", reference+".yaml"))
	yaml.Unmarshal(typeYAMLBytes, &to)
	return to, err
}

func (t testTypeProvider) getType(name string) (Type, error) {
	if allowed, ok := t.allowedTypes[name]; !allowed || !ok {
		return Type{}, fmt.Errorf("bad type '%s'", name)
	}

	return Type{Type: name}, nil
}

func createTestTypeProvider(allowedTypes []string) TypeProvider {
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

func (t testOperationProvider) getOperationRef(reference string) (Operation, error) {
	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "operations", reference+".yaml"))

	operation := Operation{}
	yaml.Unmarshal(typeYAMLBytes, &operation)

	return operation, err
}

func createTestOperationProvider() OperationProvider {
	operation := testOperationProvider{}
	return operation
}

// RESOURCE

type testResourceProvider struct {
}

func (t testResourceProvider) getResourceRef(reference string) (Resource, error) {
	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "resources", reference+".yaml"))

	resource := Resource{}
	yaml.Unmarshal(typeYAMLBytes, &resource)

	return resource, err
}

func createTestResourceProvider() ResourceProvider {
	rp := testResourceProvider{}
	return rp
}

// OPERATOR

type builder func(Generics, Values, Embedding) (Operator, error)

type testOperatorProvider struct {
	builders map[string]builder
}

func (t testOperatorProvider) getOperatorRef(reference string) (Operator, error) {
	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "operators", reference+".yaml"))

	operator := Operator{}
	yaml.Unmarshal(typeYAMLBytes, &operator)

	return operator, err
}

func (t testOperatorProvider) getOperator(reference string, generics Generics, values Values, embedding Embedding) (Operator, error) {
	if builder, ok := t.builders[reference]; ok {
		return builder(generics, values, embedding)
	}

	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "operators", reference+".yaml"))

	operator := Operator{}
	yaml.Unmarshal(typeYAMLBytes, &operator)

	return operator, err
}

func createTestOperatorProvider() OperatorProvider {
	op := testOperatorProvider{}
	return op
}