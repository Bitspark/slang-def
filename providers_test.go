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

func (t testTypeProvider) dereferenceType(reference string) (Type, error) {
	to := Type{}

	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "types", reference+".yaml"))
	yaml.Unmarshal(typeYAMLBytes, &to)
	return to, err
}

func (t testTypeProvider) buildType(name string) (Type, error) {
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

func (t testOperationProvider) dereferenceOperation(reference string) (Operation, error) {
	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "operations", reference+".yaml"))

	operation := Operation{}
	yaml.Unmarshal(typeYAMLBytes, &operation)

	return operation, err
}

func createTestOperationProvider() OperationProvider {
	operation := testOperationProvider{}
	return operation
}

// OPERATOR

type builder func(Generics, Values, Embedding) (Instance, error)

type testOperatorProvider struct {
	builders map[string]builder
}

func (t testOperatorProvider) dereferenceOperator(reference string) (Instance, error) {
	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "operators", reference+".yaml"))

	operator := Instance{}
	yaml.Unmarshal(typeYAMLBytes, &operator)

	return operator, err
}

func (t testOperatorProvider) buildOperator(reference string, generics Generics, values Values, embedding Embedding) (Instance, error) {
	if builder, ok := t.builders[reference]; ok {
		return builder(generics, values, embedding)
	}

	typeYAMLBytes, err := ioutil.ReadFile(filepath.Join("test_tmp", "operators", reference+".yaml"))

	operator := Instance{}
	yaml.Unmarshal(typeYAMLBytes, &operator)

	return operator, err
}

func createTestOperatorProvider() OperatorProvider {
	op := testOperatorProvider{}
	return op
}
