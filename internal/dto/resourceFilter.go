package dto

type ResourceFilter struct {
	FieldSelector string `query:"fieldSelector"`
}

func (rf *ResourceFilter) Validate() error {
	return nil
}
