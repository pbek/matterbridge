// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ObjectDefinition undocumented
type ObjectDefinition struct {
	// Object is the base model of ObjectDefinition
	Object
	// Attributes undocumented
	Attributes []AttributeDefinition `json:"attributes,omitempty"`
	// Metadata undocumented
	Metadata []MetadataEntry `json:"metadata,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// SupportedApis undocumented
	SupportedApis []string `json:"supportedApis,omitempty"`
}