package schema

// EnvKeyValue is the base type of Environment variable.
type EnvKeyValue struct {
	ID    uint             `json:"id,omitempty"`
	Scope EnvVariableScope `json:"scope"`
	Name  string           `json:"name"`
	Value string           `json:"value"`
}

// EnvVariableInput is based on the Lagoon API type.
type EnvVariableInput struct {
	EnvKeyValue
	Type   EnvVariableType `json:"type"`
	TypeID uint            `json:"typeId"`
}
