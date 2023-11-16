package dto

type MapResource map[string]any

type GenericResource struct {
	APIVersion string           `json:"apiVersion,omitempty"`
	Kind       string           `json:"kind,omitempty"`
	Metadata   ResourceMetadata `json:"metadata,omitempty"`

	Spec   map[string]any `json:"spec,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
	Status ResourceStatus `json:"status,omitempty"`
	Items  any            `json:"items,omitempty"`
}

func (r *GenericResource) Validate() error {
	return nil
}

type ResourceMetadata struct {
	Name string `json:"name,omitempty"`
}

type ResourceStatus struct {
	Phase string `json:"phase,omitempty"`
}
