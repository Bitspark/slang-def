package def

// A resource is an interface to the system an operator runs on.
type Resource struct {
	// Description of this resource in natural language (English).
	Description string `json:"description" yaml:"description"`

	// Operations is a map of operations this resource provides.
	Operations map[string]string `json:"operations" yaml:"operations"`
}
