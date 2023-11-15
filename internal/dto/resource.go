package dto

type GenericResource map[string]any

func (r *GenericResource) Validate() error {
	return nil
}
