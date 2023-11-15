package dto

type Resource map[string]any

func (r *Resource) Validate() error {
	return nil
}
