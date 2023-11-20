package dto

import (
	"strings"
)

type ResourceFilter struct {
	FieldSelector string `query:"fieldSelector"`
}

func (rf *ResourceFilter) Validate() error {
	return nil
}

// TODO: improve logic
func (rf *ResourceFilter) GetMetadataFilter() string {
	return strings.ReplaceAll(rf.FieldSelector, "metadata.name=", "")
}
